use super::external_rate_provider::ExternalRateProvider;
use crate::cache::cache_service::CacheService;
use crate::domain::exchange_rate::entities::ExchangeRate;
use crate::repository::exchange_rate::repository::ExchangeRateRepository;
use crate::usecases::exchange_rate::traits::IRateFetcherUseCase;
use async_trait::async_trait;
use std::error::Error;
use std::sync::Arc;
use tokio::time::{sleep, Duration};
use tracing::{error, info};

/// Use case for fetching exchange rates from external providers.
pub struct RateFetcherUseCase {
    pub repository: Arc<dyn ExchangeRateRepository + Send + Sync>,
    pub cache: Arc<CacheService>,
    pub providers: Vec<Arc<dyn ExternalRateProvider + Send + Sync>>,
    pub max_retries: usize,
}

impl RateFetcherUseCase {
    /// Creates a new instance of `RateFetcherUseCase`.
    ///
    /// # Arguments
    ///
    /// * `repository` - An implementation of `ExchangeRateRepository`.
    /// * `cache` - The cache service for storing exchange rates.
    /// * `providers` - A list of external rate providers.
    /// * `max_retries` - Maximum number of retry attempts for fetching rates.
    pub fn new(
        repository: Arc<dyn ExchangeRateRepository + Send + Sync>,
        cache: Arc<CacheService>,
        providers: Vec<Arc<dyn ExternalRateProvider + Send + Sync>>,
        max_retries: usize,
    ) -> Self {
        Self {
            repository,
            cache,
            providers,
            max_retries,
        }
    }

    /// Fetches the exchange rate for the specified currency pair.
    ///
    /// # Arguments
    ///
    /// * `from` - The source currency code (e.g., "USD").
    /// * `to` - The target currency code (e.g., "EUR").
    ///
    /// # Returns
    ///
    /// * `Some(ExchangeRate)` if successfully fetched.
    /// * `None` if all providers fail.
    pub async fn fetch_rate(&self, from: &str, to: &str) -> Option<ExchangeRate> {
        // Step 1: Check Cache
        if let Some(rate) = self.cache.get_rate(from, to).await {
            info!("Cache hit for {} to {}", from, to);
            return Some(rate);
        }

        info!("Cache miss for {} to {}", from, to);

        // Step 2: Fetch from Providers with Fallback
        for provider in &self.providers {
            for attempt in 1..=self.max_retries {
                match provider.fetch_rate(from, to).await {
                    Ok(rate) => {
                        info!(
                            "Fetched rate from provider {} on attempt {}: {:?}",
                            provider.type_id(),
                            attempt,
                            rate
                        );
                        // Step 3: Save using `save_rate` method which returns Result
                        if let Err(e) = self.save_rate(rate.clone()).await {
                            error!("Failed to save rate: {}", e);
                        }
                        return Some(rate);
                    }
                    Err(e) => {
                        error!(
                            "Error fetching rate from provider {} on attempt {}: {}",
                            provider.type_id(),
                            attempt,
                            e
                        );
                        if attempt < self.max_retries {
                            let backoff = 2_u64.pow(attempt as u32);
                            info!("Retrying after {} seconds...", backoff);
                            sleep(Duration::from_secs(backoff)).await;
                        }
                    }
                }
            }
        }

        // If all providers and retries fail
        error!(
            "Failed to fetch exchange rate for {} to {} after {} attempts",
            from, to, self.max_retries
        );
        None
    }

    /// Saves an exchange rate directly to the repository and cache.
    ///
    /// # Arguments
    ///
    /// * `rate` - The exchange rate to save.
    pub async fn save_rate(&self, rate: ExchangeRate) -> Result<(), Box<dyn Error + Send + Sync>> {
        // Store in repository
        self.repository.save_rate(&rate).await;
        // Store in cache
        self.cache.set_rate(&rate).await?;
        Ok(())
    }
}

/// Helper trait to identify the provider type.
#[async_trait]
pub trait ProviderIdentifiable {
    fn type_id(&self) -> &'static str;
}

#[async_trait]
impl ProviderIdentifiable for dyn ExternalRateProvider + Send + Sync {
    fn type_id(&self) -> &'static str {
        // This method should be overridden by concrete providers.
        "UnknownProvider"
    }
}

impl<T: ExternalRateProvider + Send + Sync> ProviderIdentifiable for T {
    fn type_id(&self) -> &'static str {
        std::any::type_name::<T>()
    }
}

#[async_trait]
impl IRateFetcherUseCase for RateFetcherUseCase {
    async fn fetch_rate(&self, from: &str, to: &str) -> Option<ExchangeRate> {
        self.fetch_rate(from, to).await
    }

    async fn save_rate(&self, rate: ExchangeRate) -> Result<(), Box<dyn Error + Send + Sync>> {
        self.repository.save_rate(&rate).await;
        self.cache.set_rate(&rate).await
    }
}
