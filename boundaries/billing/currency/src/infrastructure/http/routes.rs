use super::handlers::{get_current_exchange_rate, get_historical_exchange_rate};
use crate::usecases::currency_conversion::converter::CurrencyConversionUseCase;
use crate::usecases::exchange_rate::fetcher::RateFetcherUseCase;
use std::sync::Arc;
use warp::Filter;

pub fn api(
    rate_fetcher: Arc<RateFetcherUseCase>,
    conversion_service: Arc<CurrencyConversionUseCase>,
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

    // Combine the routes
    current_rate.or(historical_rate)
}

fn with_rate_fetcher(
    rate_fetcher: Arc<RateFetcherUseCase>,
) -> impl Filter<Extract = (Arc<RateFetcherUseCase>,), Error = std::convert::Infallible> + Clone {
    warp::any().map(move || rate_fetcher.clone())
}

fn with_conversion_service(
    conversion_service: Arc<CurrencyConversionUseCase>,
) -> impl Filter<Extract = (Arc<CurrencyConversionUseCase>,), Error = std::convert::Infallible> + Clone
{
    warp::any().map(move || conversion_service.clone())
}
