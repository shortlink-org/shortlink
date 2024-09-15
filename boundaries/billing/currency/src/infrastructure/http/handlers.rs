use warp::reply::Json;
use serde::{Deserialize, Serialize};
use rust_decimal::Decimal;
use std::sync::Arc;
use crate::usecases::exchange_rate::fetcher::RateFetcherUseCase;
use crate::usecases::currency_conversion::converter::CurrencyConversionUseCase;
use utoipa::{ToSchema, IntoParams};

// Request query parameters
#[derive(Deserialize, ToSchema, IntoParams)]
pub struct ExchangeRateQuery {
    base_currency: String,
    target_currency: String,
}

#[derive(Deserialize, ToSchema, IntoParams)]
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
pub async fn get_current_exchange_rate(
    query: ExchangeRateQuery,
    rate_fetcher: Arc<RateFetcherUseCase>,
) -> Result<Json, warp::Rejection> {
    if let Some(rate) = rate_fetcher.fetch_rate(&query.base_currency, &query.target_currency).await {
        let response = ExchangeRateResponse {
            base_currency: rate.from_currency.code,
            target_currency: rate.to_currency.code,
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
pub async fn get_historical_exchange_rate(
    query: HistoricalRateQuery,
    conversion_service: Arc<CurrencyConversionUseCase>,
) -> Result<Json, warp::Rejection> {
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
