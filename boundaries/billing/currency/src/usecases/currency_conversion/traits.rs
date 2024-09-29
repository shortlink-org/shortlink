use std::fmt::Debug;
use crate::domain::exchange_rate::entities::ExchangeRate;
use async_trait::async_trait;

/// Trait defining the interface for currency conversion.
#[async_trait]
pub trait ICurrencyConversionUseCase: Send + Sync + Debug {
    /// Retrieves historical exchange rates for the specified parameters.
    async fn get_historical_rates(
        &self,
        base_currency: &str,
        target_currency: &str,
        start_date: &str,
        end_date: &str,
    ) -> Option<Vec<ExchangeRate>>;
}
