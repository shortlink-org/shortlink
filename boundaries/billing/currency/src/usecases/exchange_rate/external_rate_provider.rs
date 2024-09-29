use crate::domain::exchange_rate::entities::ExchangeRate;
use async_trait::async_trait;
use std::error::Error;
use std::fmt::Debug;

/// Trait defining the interface for external exchange rate providers.
#[async_trait]
pub trait ExternalRateProvider: Send + Sync + Debug {
    /// Fetches the exchange rate for the given currency pair.
    ///
    /// # Arguments
    ///
    /// * `from` - The source currency code (e.g., "USD").
    /// * `to` - The target currency code (e.g., "EUR").
    ///
    /// # Returns
    ///
    /// * `Ok(ExchangeRate)` if the rate is successfully fetched.
    /// * `Err(Box<dyn Error + Send + Sync>)` if an error occurs.
    async fn fetch_rate(
        &self,
        from: &str,
        to: &str,
    ) -> Result<ExchangeRate, Box<dyn Error + Send + Sync>>;
}
