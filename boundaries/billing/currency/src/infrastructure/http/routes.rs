use warp::Filter;
use std::sync::Arc;
use super::handlers::{get_current_exchange_rate, get_historical_exchange_rate};
use crate::usecases::exchange_rate::fetcher::RateFetcherUseCase;
use crate::usecases::currency_conversion::converter::CurrencyConversionUseCase;

pub fn api(
    rate_fetcher: Arc<RateFetcherUseCase>,
    conversion_service: Arc<CurrencyConversionUseCase>,
) -> impl Filter<Extract = impl warp::Reply, Error = warp::Rejection> + Clone {
    // Current exchange rate route
    let current_rate = warp::path!("rates" / "current")
        .and(warp::get())
        .and(warp::query::<super::handlers::ExchangeRateQuery>())
        .and(with_rate_fetcher(rate_fetcher.clone()))
        .and_then(|query, rate_fetcher| get_current_exchange_rate(query, rate_fetcher));

    // Historical exchange rates route
    let historical_rate = warp::path!("rates" / "historical")
        .and(warp::get())
        .and(warp::query::<super::handlers::HistoricalRateQuery>())
        .and(with_conversion_service(conversion_service.clone()))
        .and_then(|query, conversion_service| get_historical_exchange_rate(query, conversion_service));

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
) -> impl Filter<Extract = (Arc<CurrencyConversionUseCase>,), Error = std::convert::Infallible> + Clone {
    warp::any().map(move || conversion_service.clone())
}
