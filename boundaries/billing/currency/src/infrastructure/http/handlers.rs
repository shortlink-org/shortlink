use crate::domain::exchange_rate::entities::{Currency, ExchangeRate};
use crate::usecases::currency_conversion::traits::ICurrencyConversionUseCase;
use crate::usecases::exchange_rate::traits::IRateFetcherUseCase;
use rust_decimal::Decimal;
use serde::{Deserialize, Serialize};
use std::sync::Arc;
use opentelemetry::{global, KeyValue};
use opentelemetry::metrics::{Counter, Meter};
use tracing::{info, instrument};
use utoipa::{IntoParams, ToSchema};
use warp::reply::Json;
use rust_decimal_macros::dec;

// Request query parameters
#[derive(Deserialize, ToSchema, IntoParams, Debug)]
pub struct ExchangeRateQuery {
    base_currency: String,
    target_currency: String,
}

#[derive(Deserialize, ToSchema, IntoParams, Debug)]
pub struct HistoricalRateQuery {
    base_currency: String,
    target_currency: String,
    start_date: String,
    end_date: String,
}

// Response structures
#[derive(Serialize, ToSchema)]
pub struct ExchangeRateResponse {
    base_currency: String,
    target_currency: String,
    exchange_rate: Decimal,
    timestamp: String,
}

#[derive(Serialize, ToSchema)]
pub struct HistoricalRateResponse {
    date: String,
    exchange_rate: Decimal,
}

#[utoipa::path(
    get,
    path = "/rates/current",
    params(ExchangeRateQuery),
    responses(
        (status = 200, description = "Success", body = ExchangeRateResponse),
        (status = 404, description = "Not Found")
    )
)]
// Handler for current exchange rate
#[instrument]
pub async fn get_current_exchange_rate(
    query: ExchangeRateQuery,
    rate_fetcher: Arc<dyn IRateFetcherUseCase>,
) -> Result<Json, warp::Rejection> {
    info!(
        "Fetching current exchange rate for {} to {}",
        query.base_currency, query.target_currency
    );

    if let Some(rate) = rate_fetcher
        .fetch_rate(&query.base_currency, &query.target_currency)
        .await
    {
        let response = ExchangeRateResponse {
            base_currency: rate.from.code,
            target_currency: rate.to.code,
            exchange_rate: rate.rate,
            timestamp: "2024-09-12T12:00:00Z".to_string(), // Mocked timestamp
        };
        Ok(warp::reply::json(&response))
    } else {
        Err(warp::reject::not_found())
    }
}

#[utoipa::path(
    get,
    path = "/rates/historical",
    params(HistoricalRateQuery),
    responses(
        (status = 200, description = "Success", body = [HistoricalRateResponse]),
        (status = 404, description = "Not Found")
    )
)]
// Handler for historical exchange rates
#[instrument]
pub async fn get_historical_exchange_rate(
    query: HistoricalRateQuery,
    _conversion_service: Arc<dyn ICurrencyConversionUseCase>,
) -> Result<Json, warp::Rejection> {
    info!(
        "Fetching historical exchange rates for {} to {} from {} to {}",
        query.base_currency, query.target_currency, query.start_date, query.end_date
    );

    let response = vec![
        HistoricalRateResponse {
            date: "2024-01-01".to_string(),
            exchange_rate: Decimal::new(84, 2), // 0.84
        },
        HistoricalRateResponse {
            date: "2024-01-02".to_string(),
            exchange_rate: Decimal::new(85, 2), // 0.85
        },
    ];

    Ok(warp::reply::json(&response))
}

#[derive(Debug)]
struct Metrics {
    request_counter: Counter<u64>,
    error_counter: Counter<u64>,
}

impl Metrics {
    fn new(meter: &Meter) -> Self {
        Self {
            request_counter: meter.u64_counter("requests_total").with_description("Total number of requests").init(),
            error_counter: meter.u64_counter("errors_total").with_description("Total number of errors").init(),
        }
    }

    fn increment_requests(&self) {
        self.request_counter.add(1, &[KeyValue::new("service.name", "currency-service")]);
    }

    fn increment_errors(&self) {
        self.error_counter.add(1, &[KeyValue::new("service.name", "currency-service")]);
    }
}

// In your main function or service initialization
fn init_metrics() -> Metrics {
    let meter = global::meter("currency-service");
    Metrics::new(&meter)
}
