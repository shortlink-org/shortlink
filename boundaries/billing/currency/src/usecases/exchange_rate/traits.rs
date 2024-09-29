use crate::domain::exchange_rate::entities::ExchangeRate;
use async_trait::async_trait;
use std::error::Error;
use std::fmt::Debug;

/// Trait defining the interface for fetching exchange rates.
#[async_trait]
pub trait IRateFetcherUseCase: Send + Sync + Debug {
    /// Fetches the exchange rate for the specified currency pair.
    async fn fetch_rate(&self, from: &str, to: &str) -> Option<ExchangeRate>;

    /// Saves an exchange rate to the repository and cache.
    async fn save_rate(&self, rate: ExchangeRate) -> Result<(), Box<dyn Error + Send + Sync>>;
    async fn get_historical_rates(
        &self,
        base_currency: &str,
        target_currency: &str,
        start_date: &str,
        end_date: &str,
    ) -> Option<Vec<ExchangeRate>>;
}
