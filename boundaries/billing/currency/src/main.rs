mod cache;
mod domain;
mod infrastructure;
mod repository;
mod usecases;

use std::convert::Infallible;
use infrastructure::http::routes::api;
use repository::exchange_rate::redis_repository::RedisExchangeRateRepository;
use std::sync::Arc;
use tracing::{error, info};
use tracing_subscriber::fmt::format::FmtSpan;
use tracing_subscriber::EnvFilter;
use usecases::currency_conversion::ICurrencyConversionUseCase;
use usecases::exchange_rate::RateFetcherUseCase;
use utoipa::OpenApi;
use warp::{Filter, Rejection};
use tokio_cron_scheduler::{Job, JobScheduler};
use dotenvy::dotenv;
use std::env;
use std::iter::Map;
use tracing_subscriber::filter::combinator::And;
use warp::path::Exact;
use warp::reply::Json;
use crate::cache::CacheService;
use crate::repository::exchange_rate::in_memory_repository::InMemoryExchangeRateRepository;
use crate::repository::exchange_rate::repository::ExchangeRateRepository;
use crate::usecases::currency_conversion::converter::CurrencyConversionUseCase;
use usecases::exchange_rate::mock_bloomberg_provider::MockBloombergProvider;
use usecases::exchange_rate::mock_yahoo_provider::MockYahooProvider;

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
async fn main() {
    // Load environment variables
    dotenv().ok();
    init_tracing();

    // Initialize services (repositories, caches, external providers, etc.)
    let (rate_fetcher_use_case, currency_conversion_use_case) = init_services().await;

    // Initialize scheduler and job
    let mut scheduler = init_scheduler(rate_fetcher_use_case.clone()).await;

    // Generate OpenAPI specification
    let openapi = ApiDoc::openapi();
    let openapi_filter = warp::path!("api-docs" / "openapi.json")
        .map(move || warp::reply::json(&openapi));

    // Serve the API routes
    serve_api(rate_fetcher_use_case, currency_conversion_use_case).await;

    // Wait for a shutdown signal and stop the scheduler
    tokio::signal::ctrl_c().await.unwrap();
    scheduler.shutdown().await.unwrap();
    println!("Shutting down");
}

/// Initialize environment logging using tracing.
fn init_tracing() {
    let log_level = env::var("LOG_LEVEL").unwrap_or_else(|_| "info".to_string());
    tracing_subscriber::fmt()
        .with_env_filter(EnvFilter::new(log_level))
        .with_span_events(FmtSpan::CLOSE)
        .init();
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

    // Define a cron job to run every hour (adjust for testing if needed)
    let job = Job::new_async("0 * * * * *", move |_uuid, _l| {
        let rate_fetcher_use_case = rate_fetcher_use_case.clone();
        Box::pin(async move {
            run_currency_update_job(rate_fetcher_use_case).await;
        })
    }).unwrap();

    scheduler.add(job).await.unwrap();
    scheduler.start().await.unwrap();

    scheduler
}

/// Periodic currency update job.
async fn run_currency_update_job(rate_fetcher_use_case: Arc<RateFetcherUseCase>) {
    info!("Starting exchange rate update job...");

    let from_currency = "USD"; // Example currencies
    let to_currency = "EUR";

    match rate_fetcher_use_case.fetch_rate(from_currency, to_currency).await {
        Some(exchange_rate) => {
            info!(
                "Successfully fetched exchange rate for {} to {}: {}",
                from_currency, to_currency, exchange_rate.rate
            );
        }
        None => {
            error!(
                "Failed to fetch exchange rate for {} to {}",
                from_currency, to_currency
            );
        }
    }
}

/// Set up and run the HTTP API server.
async fn serve_api(
    rate_fetcher_use_case: Arc<RateFetcherUseCase>,
    currency_conversion_use_case: Arc<CurrencyConversionUseCase>,
) {
    // Retrieve server host and port from environment
    let server_host = env::var("SERVER_HOST").unwrap_or_else(|_| "127.0.0.1".to_string());
    let server_port: u16 = env::var("SERVER_PORT")
        .unwrap_or_else(|_| "3030".to_string())
        .parse()
        .expect("SERVER_PORT must be a valid u16");

    // Set up the HTTP server with the API routes and OpenAPI filter
    let routes = api(rate_fetcher_use_case, currency_conversion_use_case)
        .with(warp::trace::request());

    warp::serve(routes).run(([127, 0, 0, 1], server_port)).await;
}
