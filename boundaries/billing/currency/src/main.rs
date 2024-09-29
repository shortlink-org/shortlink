mod cache;
mod domain;
mod infrastructure;
mod repository;
mod usecases;

use std::env;
use std::sync::Arc;
use deadpool_redis::Metrics;
use dotenvy::dotenv;
use tokio_cron_scheduler::{Job, JobScheduler};
use tracing::{error, instrument};
use tracing_subscriber::prelude::*;
use tracing_subscriber::{fmt, EnvFilter, Registry};
use opentelemetry::{global, KeyValue};
use opentelemetry::metrics::MetricsError::Config;
use opentelemetry::trace::TracerProvider;
use opentelemetry_otlp::WithExportConfig;
use tracing_opentelemetry::OpenTelemetryLayer;
use utoipa::OpenApi;
use warp::{Filter, Rejection, Reply};

// Import OpenTelemetry SDK trace module
use opentelemetry_sdk::{trace as sdktrace, Resource};
use opentelemetry_sdk::runtime::Tokio;
use serde::de::Error;
use tracing_subscriber::fmt::format::FmtSpan;
use spandoc::spandoc;
use tracing::log::info;
// Import service modules
use crate::cache::CacheService;
use crate::infrastructure::http::routes::api;
use crate::repository::exchange_rate::redis_repository::RedisExchangeRateRepository;
use crate::repository::exchange_rate::repository::ExchangeRateRepository;
use crate::usecases::currency_conversion::converter::CurrencyConversionUseCase;
use crate::usecases::exchange_rate::mock_bloomberg_provider::MockBloombergProvider;
use crate::usecases::exchange_rate::mock_yahoo_provider::MockYahooProvider;
use crate::usecases::exchange_rate::RateFetcherUseCase;
use crate::usecases::load_exchange_rates_cron::cron::run_currency_update_job;

#[derive(OpenApi)]
#[openapi(
    paths(
        infrastructure::http::handlers::get_current_exchange_rate,
        infrastructure::http::handlers::get_historical_exchange_rate,
    ),
    components(schemas(
        infrastructure::http::handlers::ExchangeRateQuery,
        infrastructure::http::handlers::HistoricalRateQuery,
        infrastructure::http::handlers::ExchangeRateResponse,
        infrastructure::http::handlers::HistoricalRateResponse
    )),
    tags(
        (name = "Currency", description = "Currency conversion and exchange rates API")
    )
)]
struct ApiDoc;

#[tokio::main]
#[spandoc]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Load environment variables
    dotenv().ok();

    // Initialize tracing and OpenTelemetry
    init_tracing()?;

    // Initialize services (repositories, caches, external providers, etc.)
    let (rate_fetcher_use_case, currency_conversion_use_case) = init_services().await;

    // Initialize Metrics (this should match your actual metrics usage)
    let metrics = Arc::new(Metrics::default());

    // Initialize scheduler and job
    let mut scheduler = init_scheduler(rate_fetcher_use_case.clone()).await;

    // Generate OpenAPI specification
    let openapi = ApiDoc::openapi();
    let openapi_filter = warp::path!("api-docs" / "openapi.json")
        .map(move || warp::reply::json(&openapi));

    // Serve the API routes
    serve_api(rate_fetcher_use_case, currency_conversion_use_case, openapi_filter, metrics).await;

    // Wait for a shutdown signal and stop the scheduler
    tokio::signal::ctrl_c().await.unwrap();
    scheduler.shutdown().await.unwrap();
    info!("Shutting down");

    // Shutdown the tracer
    global::shutdown_tracer_provider();

    Ok(())
}

/// Initialize OpenTelemetry tracing using the OTLP exporter over gRPC (Tempo)
fn init_tracing() -> Result<(), opentelemetry::trace::TraceError> {
    // Get the OTLP endpoint from the environment variable or use a default
    let otlp_endpoint = env::var("OTLP_ENDPOINT")
        .unwrap_or_else(|_| "http://localhost:4317".to_string());

    // Create the OTLP exporter and get a TracerProvider
    let tracer_provider = opentelemetry_otlp::new_pipeline()
        .tracing()
        .with_exporter(
            opentelemetry_otlp::new_exporter()
                .tonic() // Use gRPC
                .with_endpoint(otlp_endpoint),
        )
        .with_trace_config(
            sdktrace::Config::default()
                .with_resource(Resource::new(vec![
                    KeyValue::new("service.name", "currency-service"),
                ])),
        )
        .install_batch(Tokio)?;

    // Set the global tracer provider
    opentelemetry::global::set_tracer_provider(tracer_provider.clone());

    // Get a tracer from the tracer provider
    let tracer = tracer_provider.tracer("currency-service");

    // Create a tracing layer with the configured tracer
    let telemetry = tracing_opentelemetry::layer().with_tracer(tracer);

    // Create a tracing subscriber for console and file output
    let subscriber = Registry::default()
        .with(EnvFilter::from_default_env())
        .with(telemetry);

    tracing::subscriber::set_global_default(subscriber)
        .expect("Failed to set the tracing subscriber");

    Ok(())
}

/// Initialize the services (repositories, caches, providers, etc.)
async fn init_services() -> (Arc<RateFetcherUseCase>, Arc<CurrencyConversionUseCase>) {
    // Retrieve Redis URL from environment
    let redis_url = env::var("REDIS_URL").expect("REDIS_URL must be set in .env");

    // Create Redis repository
    let exchange_rate_repository = Arc::new(
        RedisExchangeRateRepository::new(&redis_url)
            .await
            .expect("Failed to connect to Redis"),
    );

    // Initialize cache service
    let cache_service = Arc::new(CacheService::new());

    // Initialize external providers
    let bloomberg_provider = Arc::new(MockBloombergProvider::new());
    let yahoo_provider = Arc::new(MockYahooProvider::new());

    // Load weighted average approach configuration from environment variables or use defaults
    let divergence_threshold: f64 = env::var("DIVERGENCE_THRESHOLD")
        .unwrap_or_else(|_| "0.5".to_string()) // Default to 0.5%
        .parse()
        .expect("DIVERGENCE_THRESHOLD must be a valid number");

    let bloomberg_weight: f64 = env::var("BLOOMBERG_WEIGHT")
        .unwrap_or_else(|_| "0.7".to_string()) // Default to 0.7
        .parse()
        .expect("BLOOMBERG_WEIGHT must be a valid number");

    let yahoo_weight: f64 = env::var("YAHOO_WEIGHT")
        .unwrap_or_else(|_| "0.3".to_string()) // Default to 0.3
        .parse()
        .expect("YAHOO_WEIGHT must be a valid number");

    // Create RateFetcherUseCase
    let rate_fetcher_use_case = Arc::new(RateFetcherUseCase::new(
        exchange_rate_repository.clone() as Arc<dyn ExchangeRateRepository + Send + Sync>,
        cache_service.clone(),
        vec![bloomberg_provider, yahoo_provider],
        3, // max_retries
    ));

    // Create CurrencyConversionUseCase with weighted average approach
    let currency_conversion_use_case = Arc::new(CurrencyConversionUseCase::new(
        rate_fetcher_use_case.clone(),
        divergence_threshold,
        bloomberg_weight,
        yahoo_weight,
    ));

    (rate_fetcher_use_case, currency_conversion_use_case)
}

/// Initialize the scheduler and add the currency update job.
async fn init_scheduler(rate_fetcher_use_case: Arc<RateFetcherUseCase>) -> JobScheduler {
    let scheduler = JobScheduler::new().await.unwrap();

    // Define a cron job to run every hour (adjust as needed)
    let job = Job::new_async("0 0 * * * *", move |_uuid, _l| {
        let rate_fetcher_use_case = rate_fetcher_use_case.clone();
        Box::pin(async move {
            run_currency_update_job(rate_fetcher_use_case).await;
        })
    })
        .unwrap();

    scheduler.add(job).await.unwrap();
    scheduler.start().await.unwrap();

    scheduler
}

/// Set up and run the HTTP API server.
async fn serve_api(
    rate_fetcher_use_case: Arc<RateFetcherUseCase>,
    currency_conversion_use_case: Arc<CurrencyConversionUseCase>,
    openapi_filter: impl Filter<Extract = impl warp::Reply, Error = warp::Rejection> + Clone + Send + Sync + 'static,
    metrics: Arc<Metrics>,
) {
    // Retrieve server host and port from environment
    let server_host = env::var("SERVER_HOST").unwrap_or_else(|_| "127.0.0.1".to_string());
    let server_port: u16 = env::var("SERVER_PORT")
        .unwrap_or_else(|_| "3030".to_string())
        .parse()
        .expect("SERVER_PORT must be a valid u16");

    // Set up the HTTP server with the API routes and OpenAPI filter
    let routes = api(rate_fetcher_use_case, currency_conversion_use_case, metrics)
        .or(openapi_filter)
        .with(warp::trace::request());

    info!("Starting HTTP server at http://{}:{}", server_host, server_port);

    warp::serve(routes).run(([127, 0, 0, 1], server_port)).await;
}
