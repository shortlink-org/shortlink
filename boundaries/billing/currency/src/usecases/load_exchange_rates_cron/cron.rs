use tokio_cron_scheduler::{Job, JobScheduler};
use std::sync::Arc;
use tracing::{info, error, instrument};
use crate::usecases::exchange_rate::RateFetcherUseCase;

/// Periodic currency update job.
#[instrument]
pub(crate) async fn run_currency_update_job(rate_fetcher_use_case: Arc<RateFetcherUseCase>) {
    info!("Starting exchange rate update job...");

    let from_currency = "USD";
    let to_currency = "EUR";

    match rate_fetcher_use_case.fetch_rate(from_currency, to_currency).await {
        Some(exchange_rate) => {
            info!(
                "Successfully fetched exchange rate for {} to {}: {}",
                from_currency, to_currency, exchange_rate.rate
            );
            // Additional logic to store in cache/database if needed
        }
        None => {
            error!(
                "Failed to fetch exchange rate for {} to {}",
                from_currency, to_currency
            );
        }
    }
}
