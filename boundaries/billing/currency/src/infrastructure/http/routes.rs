use super::handlers::{get_current_exchange_rate, get_historical_exchange_rate};
use crate::usecases::currency_conversion::converter::traits::ICurrencyConversionUseCase;
use crate::usecases::exchange_rate::fetcher::traits::IRateFetcherUseCase;
use std::sync::Arc;
use warp::Filter;

/// Defines the API routes for exchange rates.
pub fn api(
    rate_fetcher: Arc<dyn IRateFetcherUseCase>,
    conversion_service: Arc<dyn ICurrencyConversionUseCase>,
) -> impl Filter<Extract = impl warp::Reply, Error = warp::Rejection> + Clone {
    // Current exchange rate route
    let current_rate = warp::path!("rates" / "current")
        .and(warp::get())
        .and(warp::query::<super::handlers::ExchangeRateQuery>())
        .and(with_rate_fetcher(rate_fetcher.clone()))
        .and_then(get_current_exchange_rate);

    // Historical exchange rates route
    let historical_rate = warp::path!("rates" / "historical")
        .and(warp::get())
        .and(warp::query::<super::handlers::HistoricalRateQuery>())
        .and(with_conversion_service(conversion_service.clone()))
        .and_then(get_historical_exchange_rate);

    // Combine the routes using `or` to allow either route to be matched
    current_rate.or(historical_rate)
}

/// Injects the `RateFetcherUseCase` into the Warp filter chain.
///
/// # Arguments
///
/// * `rate_fetcher` - An `Arc` pointing to a trait object implementing `IRateFetcherUseCase`.
fn with_rate_fetcher(
    rate_fetcher: Arc<dyn IRateFetcherUseCase>,
) -> impl Filter<Extract = (Arc<dyn IRateFetcherUseCase>,), Error = std::convert::Infallible> + Clone {
    warp::any().map(move || rate_fetcher.clone())
}

/// Injects the `CurrencyConversionUseCase` into the Warp filter chain.
///
/// # Arguments
///
/// * `conversion_service` - An `Arc` pointing to a trait object implementing `ICurrencyConversionUseCase`.
fn with_conversion_service(
    conversion_service: Arc<dyn ICurrencyConversionUseCase>,
) -> impl Filter<Extract = (Arc<dyn ICurrencyConversionUseCase>,), Error = std::convert::Infallible> + Clone
{
    warp::any().map(move || conversion_service.clone())
}
