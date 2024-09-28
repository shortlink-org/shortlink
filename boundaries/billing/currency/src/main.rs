mod domain;
mod usecases;
mod repository;
mod cache;
mod infrastructure;

use std::sync::Arc;
use warp::Filter;
use rust_decimal_macros::dec;
use domain::exchange_rate::entities::{Currency, ExchangeRate};
use usecases::exchange_rate::fetcher::RateFetcherUseCase;
use usecases::currency_conversion::converter::CurrencyConversionUseCase;
use repository::exchange_rate::in_memory_repository::InMemoryExchangeRateRepository;
use infrastructure::http::routes::api;
use utoipa::OpenApi;
use tracing_subscriber::fmt::format::FmtSpan;
use tracing_subscriber::EnvFilter;
use tracing::info;

// Import the ExchangeRateRepository trait
use repository::exchange_rate::repository::ExchangeRateRepository;
use repository::exchange_rate::redis_repository::RedisExchangeRateRepository;

// Import dotenvy
use dotenvy::dotenv;
use std::env;

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
    // Load environment variables from .env file
    dotenv().ok();

    // Initialize tracing subscriber with log level from environment
    let log_level = env::var("LOG_LEVEL").unwrap_or_else(|_| "info".to_string());
    tracing_subscriber::fmt()
        .with_env_filter(EnvFilter::new(log_level))
        .with_span_events(FmtSpan::CLOSE)
        .init();

    // Retrieve Redis URL from environment
    let redis_url = env::var("REDIS_URL").expect("REDIS_URL must be set in .env");

    // Create Redis repository and use cases
    let exchange_rate_repository = Arc::new(
        RedisExchangeRateRepository::new(&redis_url)
            .await
            .expect("Failed to connect to Redis"),
    );

    // Create repository and use cases
    let rate_fetcher_use_case = Arc::new(RateFetcherUseCase::new(
        exchange_rate_repository.clone() as Arc<dyn ExchangeRateRepository + Send + Sync>,
    ));
    let _exchange_rate_repository = Arc::new(InMemoryExchangeRateRepository::new());
    let currency_conversion_use_case = Arc::new(CurrencyConversionUseCase::new(rate_fetcher_use_case.clone()));

    // Example rate to save
    let usd_to_eur_rate = ExchangeRate::new(
        Currency {
            code: "USD".to_string(),
            symbol: "$".to_string(),
        },
        Currency {
            code: "EUR".to_string(),
            symbol: "€".to_string(),
        },
        dec!(0.85), // Example rate
    );

    rate_fetcher_use_case.save_rate(usd_to_eur_rate).await;

    // Generate OpenAPI specification
    let openapi = ApiDoc::openapi();

    // Serve OpenAPI JSON at `/api-docs/openapi.json`
    let openapi_filter = warp::path!("api-docs" / "openapi.json")
        .map(move || warp::reply::json(&openapi));

    // Retrieve server host and port from environment
    let server_host = env::var("SERVER_HOST").unwrap_or_else(|_| "127.0.0.1".to_string());
    let server_port: u16 = env::var("SERVER_PORT")
        .unwrap_or_else(|_| "3030".to_string())
        .parse()
        .expect("SERVER_PORT must be a valid u16");

    // Set up the HTTP server with the API routes and OpenAPI filter
    let routes = api(rate_fetcher_use_case, currency_conversion_use_case)
        .or(openapi_filter)
        .with(warp::trace::request());

    warp::serve(routes)
        .run(([127, 0, 0, 1], server_port))
        .await;
}
