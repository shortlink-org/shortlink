use crate::domain::exchange_rate::entities::ExchangeRate;
use async_trait::async_trait;

#[async_trait]
pub trait ExchangeRateRepository: Send + Sync {
    async fn save_rate(&self, rate: &ExchangeRate) -> Result<(), Box<dyn std::error::Error + Send + Sync>>;

    async fn get_rate(&self, from: &str, to: &str) -> Option<ExchangeRate>;

    async fn get_historical_rates(
        &self,
        base_currency: &str,
        target_currency: &str,
        start_date: &str,
        end_date: &str,
    ) -> Option<Vec<ExchangeRate>>;
}
