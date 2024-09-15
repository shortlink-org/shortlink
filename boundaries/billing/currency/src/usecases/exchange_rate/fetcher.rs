use crate::domain::exchange_rate::entities::{ExchangeRate};
use crate::repository::exchange_rate::repository::ExchangeRateRepository;

pub struct RateFetcherUseCase<R: ExchangeRateRepository> {
    pub repository: R,
}

impl<R: ExchangeRateRepository> RateFetcherUseCase<R> {
    pub fn new(repository: R) -> Self {
        Self { repository }
    }

    pub async fn fetch_rate(&self, from: &str, to: &str) -> Option<ExchangeRate> {
        // Fetch rate from repository or external API
        self.repository.get_rate(from, to).await
    }

    pub async fn save_rate(&self, rate: ExchangeRate) {
        self.repository.save_rate(&rate).await;
    }
}
