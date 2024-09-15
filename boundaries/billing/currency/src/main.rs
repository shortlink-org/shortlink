mod domain;
mod usecases;
mod repository;
mod cache;
mod infrastructure;

use std::sync::Arc;
use warp::Filter;
use rust_decimal_macros::dec;
use domain::exchange_rate::entities::{Currency, ExchangeRate};
use crate::repository::exchange_rate::repository::ExchangeRateRepository;
use usecases::exchange_rate::fetcher::RateFetcherUseCase;
use usecases::currency_conversion::converter::CurrencyConversionUseCase;
use repository::exchange_rate::in_memory_repository::InMemoryExchangeRateRepository;
use infrastructure::http::routes::api;
use utoipa::OpenApi;

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
    // Create repository and use cases
    let exchange_rate_repository = Arc::new(InMemoryExchangeRateRepository::new());
    let rate_fetcher_use_case = Arc::new(RateFetcherUseCase::new(exchange_rate_repository as Arc<dyn ExchangeRateRepository + Send + Sync>));
    let currency_conversion_use_case = Arc::new(CurrencyConversionUseCase::new(rate_fetcher_use_case.clone()));

    // Example rate to save
    let usd_to_eur_rate = ExchangeRate::new(
        Currency { code: "USD".to_string(), symbol: "$".to_string() },
        Currency { code: "EUR".to_string(), symbol: "€".to_string() },
        dec!(0.85), // Example rate
    );

    rate_fetcher_use_case.save_rate(usd_to_eur_rate).await;

    // Generate OpenAPI specification
    let openapi = ApiDoc::openapi();

    // Serve OpenAPI JSON at `/api-docs/openapi.json`
    let openapi_filter = warp::path!("api-docs" / "openapi.json")
        .map(move || warp::reply::json(&openapi));

    // Set up the HTTP server with the API routes
    let routes = api(rate_fetcher_use_case, currency_conversion_use_case)
        .or(openapi_filter);

    warp::serve(routes)
        .run(([127, 0, 0, 1], 3030))
        .await;
}
