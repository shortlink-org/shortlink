use crate::domain::exchange_rate::entities::ExchangeRate;

#[async_trait::async_trait]
pub trait ExchangeRateRepository {
    async fn get_rate(&self, from: &str, to: &str) -> Option<ExchangeRate>;
    async fn save_rate(&self, rate: &ExchangeRate);
}