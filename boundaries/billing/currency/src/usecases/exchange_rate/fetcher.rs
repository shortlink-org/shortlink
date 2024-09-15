use std::sync::Arc;
use crate::repository::exchange_rate::repository::ExchangeRateRepository;

pub struct RateFetcherUseCase {
    pub repository: Arc<dyn ExchangeRateRepository + Send + Sync>, // Use Arc<dyn Trait>
}

impl RateFetcherUseCase {
    pub fn new(repository: Arc<dyn ExchangeRateRepository + Send + Sync>) -> Self {
        Self { repository }
    }

    pub async fn fetch_rate(&self, from: &str, to: &str) -> Option<crate::domain::exchange_rate::entities::ExchangeRate> {
        self.repository.get_rate(from, to).await
    }

    pub async fn save_rate(&self, rate: crate::domain::exchange_rate::entities::ExchangeRate) {
        self.repository.save_rate(&rate).await;
    }
}
