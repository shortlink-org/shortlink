#[cfg(test)]
mod tests {
    use super::*;
    use crate::usecases::currency_conversion::converter::traits::ICurrencyConversionUseCase;
    use crate::usecases::exchange_rate::fetcher::traits::IRateFetcherUseCase;
    use crate::usecases::exchange_rate::fetcher::RateFetcherUseCase;
    use crate::repository::exchange_rate::in_memory_repository::InMemoryExchangeRateRepository;
    use crate::cache::CacheService;
    use crate::domain::exchange_rate::entities::{Currency, ExchangeRate};
    use rust_decimal_macros::dec;
    use serde_json::json;
    use warp::http::StatusCode;
    use warp::test::request;
    use std::sync::Arc;
    use async_trait::async_trait;
    use std::error::Error;
    use crate::infrastructure::http::routes::api;
    use crate::usecases::currency_conversion::converter::converter::CurrencyConversionUseCase;

    /// Mock implementation of `IRateFetcherUseCase`
    struct MockRateFetcherUseCase;

    #[async_trait]
    impl IRateFetcherUseCase for MockRateFetcherUseCase {
        async fn fetch_rate(&self, from: &str, to: &str) -> Option<ExchangeRate> {
            if from == "USD" && to == "EUR" {
                Some(ExchangeRate::new(
                    Currency { code: "USD".to_string(), symbol: "$".to_string() },
                    Currency { code: "EUR".to_string(), symbol: "€".to_string() },
                    dec!(0.85),
                ))
            } else {
                None
            }
        }

        async fn save_rate(&self, _rate: ExchangeRate) -> Result<(), Box<dyn Error + Send + Sync>> {
            Ok(())
        }
    }

    /// Mock implementation of `ICurrencyConversionUseCase`
    struct MockCurrencyConversionUseCase;

    #[async_trait]
    impl ICurrencyConversionUseCase for MockCurrencyConversionUseCase {
        async fn get_historical_rates(
            &self,
            _base_currency: &str,
            _target_currency: &str,
            _start_date: &str,
            _end_date: &str,
        ) -> Option<Vec<ExchangeRate>> {
            Some(vec![
                ExchangeRate::new(
                    Currency { code: "USD".to_string(), symbol: "$".to_string() },
                    Currency { code: "EUR".to_string(), symbol: "€".to_string() },
                    dec!(0.84),
                ),
                ExchangeRate::new(
                    Currency { code: "USD".to_string(), symbol: "$".to_string() },
                    Currency { code: "EUR".to_string(), symbol: "€".to_string() },
                    dec!(0.85),
                ),
            ])
        }
    }

    #[tokio::test]
    async fn test_get_current_exchange_rate_success() {
        // Setup
        let rate_fetcher = Arc::new(MockRateFetcherUseCase);
        let conversion_service = Arc::new(MockCurrencyConversionUseCase);

        let api_filter = api(rate_fetcher.clone(), conversion_service.clone());

        // Execute
        let resp = request()
            .method("GET")
            .path("/rates/current?base_currency=USD&target_currency=EUR")
            .reply(&api_filter)
            .await;

        // Verify
        assert_eq!(resp.status(), StatusCode::OK);
        let expected = json!({
            "base_currency": "USD",
            "target_currency": "EUR",
            "exchange_rate": "0.85",
            "timestamp": "2024-09-12T12:00:00Z" // Ensure your handler provides this or adjust accordingly
        });
        let body: serde_json::Value = serde_json::from_slice(resp.body()).unwrap();
        assert_eq!(body, expected);
    }

    #[tokio::test]
    async fn test_get_current_exchange_rate_not_found() {
        // Setup
        let rate_fetcher = Arc::new(MockRateFetcherUseCase);
        let conversion_service = Arc::new(MockCurrencyConversionUseCase);

        let api_filter = api(rate_fetcher.clone(), conversion_service.clone());

        // Execute
        let resp = request()
            .method("GET")
            .path("/rates/current?base_currency=GBP&target_currency=JPY")
            .reply(&api_filter)
            .await;

        // Verify
        assert_eq!(resp.status(), StatusCode::NOT_FOUND);
    }

    #[tokio::test]
    async fn test_get_historical_exchange_rate_success() {
        // Setup
        let repository = Arc::new(InMemoryExchangeRateRepository::new());
        let cache = Arc::new(CacheService::new());
        let rate_fetcher = Arc::new(RateFetcherUseCase::new(
            repository.clone(),
            cache.clone(),
            vec![], // Add mock providers if necessary
            3,      // max_retries
        ));
        let conversion_service = Arc::new(CurrencyConversionUseCase::new(rate_fetcher.clone()));

        // Insert historical rates into the repository
        let historical_rates = vec![
            ExchangeRate::new(
                Currency {
                    code: "USD".to_string(),
                    symbol: "$".to_string(),
                },
                Currency {
                    code: "EUR".to_string(),
                    symbol: "€".to_string(),
                },
                dec!(0.84),
            ),
            ExchangeRate::new(
                Currency {
                    code: "USD".to_string(),
                    symbol: "$".to_string(),
                },
                Currency {
                    code: "EUR".to_string(),
                    symbol: "€".to_string(),
                },
                dec!(0.85),
            ),
        ];

        // Insert into the in-memory repository
        for rate in historical_rates.iter() {
            rate_fetcher.save_rate(rate.clone()).await.unwrap();
        }

        let api_filter = api(rate_fetcher.clone(), conversion_service.clone());

        // Execute
        let resp = request()
            .method("GET")
            .path("/rates/historical?base_currency=USD&target_currency=EUR&start_date=2024-01-01&end_date=2024-01-02")
            .reply(&api_filter)
            .await;

        // Verify
        assert_eq!(resp.status(), StatusCode::OK);
        let expected = json!([
            {
                "date": "2024-01-01",
                "exchange_rate": "0.84"
            },
            {
                "date": "2024-01-02",
                "exchange_rate": "0.85"
            }
        ]);
        let body: serde_json::Value = serde_json::from_slice(resp.body()).unwrap();
        assert_eq!(body, expected);
    }

    #[tokio::test]
    async fn test_get_historical_exchange_rate_not_found() {
        // Setup
        let repository = Arc::new(InMemoryExchangeRateRepository::new());
        let cache = Arc::new(CacheService::new());
        let rate_fetcher = Arc::new(RateFetcherUseCase::new(
            repository.clone(),
            cache.clone(),
            vec![], // Add mock providers if necessary
            3,      // max_retries
        ));
        let conversion_service = Arc::new(CurrencyConversionUseCase::new(rate_fetcher.clone()));

        let api_filter = api(rate_fetcher.clone(), conversion_service.clone());

        // Execute
        let resp = request()
            .method("GET")
            .path("/rates/historical?base_currency=USD&target_currency=EUR&start_date=2024-01-01&end_date=2024-01-02")
            .reply(&api_filter)
            .await;

        // Verify
        assert_eq!(resp.status(), StatusCode::NOT_FOUND);
    }
}
