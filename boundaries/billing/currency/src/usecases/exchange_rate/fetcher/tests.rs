#[cfg(test)]
mod tests {
    use super::*;
    use crate::cache::cache_service::CacheService;
    use crate::domain::exchange_rate::entities::{Currency, ExchangeRate};
    use crate::repository::exchange_rate::in_memory_repository::InMemoryExchangeRateRepository;
    use crate::usecases::exchange_rate::fetcher::external_rate_provider::ExternalRateProvider;
    use crate::usecases::exchange_rate::fetcher::mock_bloomberg_provider::MockBloombergProvider;
    use crate::usecases::exchange_rate::fetcher::mock_yahoo_provider::MockYahooProvider;
    use rust_decimal_macros::dec;
    use std::error::Error;
    use std::sync::Arc;
    use tokio::sync::Mutex;

    #[tokio::test]
    async fn test_fetch_rate_cache_hit() {
        // Setup
        let repository = Arc::new(InMemoryExchangeRateRepository::new());
        let cache = Arc::new(CacheService::new());
        let bloomberg = Arc::new(MockBloombergProvider::new());
        let yahoo = Arc::new(MockYahooProvider::new());

        // Insert a rate into the cache
        let cached_rate = ExchangeRate::new(
            Currency {
                code: "USD".to_string(),
                symbol: "$".to_string(),
            },
            Currency {
                code: "EUR".to_string(),
                symbol: "€".to_string(),
            },
            dec!(0.85),
        );
        cache.set_rate(&cached_rate).await.unwrap();

        let use_case =
            RateFetcherUseCase::new(repository.clone(), cache.clone(), vec![bloomberg, yahoo], 3);

        // Execute
        let rate = use_case.fetch_rate("USD", "EUR").await;

        // Verify
        assert!(rate.is_some());
        let rate = rate.unwrap();
        assert_eq!(rate.rate, dec!(0.85));
    }

    #[tokio::test]
    async fn test_fetch_rate_cache_miss_bloomberg_success() {
        // Setup
        let repository = Arc::new(InMemoryExchangeRateRepository::new());
        let cache = Arc::new(CacheService::new());
        let bloomberg = Arc::new(MockBloombergProvider::new());
        let yahoo = Arc::new(MockYahooProvider::new());

        let use_case =
            RateFetcherUseCase::new(repository.clone(), cache.clone(), vec![bloomberg, yahoo], 3);

        // Execute
        let rate = use_case.fetch_rate("USD", "EUR").await;

        // Verify
        assert!(rate.is_some());
        let rate = rate.unwrap();
        assert_eq!(rate.rate, dec!(0.85));

        // Ensure it's stored in cache and repository
        let cached_rate = cache.get_rate("USD", "EUR").await.unwrap();
        assert_eq!(cached_rate.rate, dec!(0.85));

        let db_rate = repository.get_rate("USD", "EUR").await.unwrap();
        assert_eq!(db_rate.rate, dec!(0.85));
    }

    #[tokio::test]
    async fn test_fetch_rate_cache_miss_bloomberg_fail_yahoo_success() {
        // Setup
        struct FailingProvider;

        #[async_trait]
        impl ExternalRateProvider for FailingProvider {
            async fn fetch_rate(
                &self,
                _from: &str,
                _to: &str,
            ) -> Result<ExchangeRate, Box<dyn Error + Send + Sync>> {
                Err("Provider failure".into())
            }
        }

        let repository = Arc::new(InMemoryExchangeRateRepository::new());
        let cache = Arc::new(CacheService::new());
        let failing_provider = Arc::new(FailingProvider);
        let yahoo = Arc::new(MockYahooProvider::new());

        let use_case = RateFetcherUseCase::new(
            repository.clone(),
            cache.clone(),
            vec![failing_provider, yahoo],
            3,
        );

        // Execute
        let rate = use_case.fetch_rate("USD", "EUR").await;

        // Verify
        assert!(rate.is_some());
        let rate = rate.unwrap();
        assert_eq!(rate.rate, dec!(0.86)); // Yahoo's mock rate

        // Ensure it's stored in cache and repository
        let cached_rate = cache.get_rate("USD", "EUR").await.unwrap();
        assert_eq!(cached_rate.rate, dec!(0.86));

        let db_rate = repository.get_rate("USD", "EUR").await.unwrap();
        assert_eq!(db_rate.rate, dec!(0.86));
    }

    #[tokio::test]
    async fn test_fetch_rate_all_providers_fail() {
        // Setup
        struct FailingProvider;

        #[async_trait]
        impl ExternalRateProvider for FailingProvider {
            async fn fetch_rate(
                &self,
                _from: &str,
                _to: &str,
            ) -> Result<ExchangeRate, Box<dyn Error + Send + Sync>> {
                Err("Provider failure".into())
            }
        }

        let repository = Arc::new(InMemoryExchangeRateRepository::new());
        let cache = Arc::new(CacheService::new());
        let failing_provider1 = Arc::new(FailingProvider);
        let failing_provider2 = Arc::new(FailingProvider);

        let use_case = RateFetcherUseCase::new(
            repository.clone(),
            cache.clone(),
            vec![failing_provider1, failing_provider2],
            2,
        );

        // Execute
        let rate = use_case.fetch_rate("USD", "EUR").await;

        // Verify
        assert!(rate.is_none());

        // Ensure it's not stored in cache or repository
        assert!(cache.get_rate("USD", "EUR").await.is_none());
        assert!(repository.get_rate("USD", "EUR").await.is_none());
    }
}
