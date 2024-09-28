#[cfg(test)]
mod tests {
    use crate::domain::exchange_rate::entities::{Currency, ExchangeRate};
    use crate::infrastructure::http::routes::api;
    use crate::usecases::currency_conversion::traits::ICurrencyConversionUseCase;
    use crate::usecases::exchange_rate::traits::IRateFetcherUseCase;
    use rust_decimal::Decimal;
    use serde_json::json;
    use std::sync::Arc;
    use warp::http::StatusCode;
    use warp::test::request;
    use warp::Filter;

    /// Mock implementation of `IRateFetcherUseCase`
    struct MockRateFetcherUseCase;

    #[async_trait::async_trait]
    impl IRateFetcherUseCase for MockRateFetcherUseCase {
        /// Returns a predefined exchange rate for USD to EUR, else None
        async fn fetch_rate(&self, from: &str, to: &str) -> Option<ExchangeRate> {
            if from.eq_ignore_ascii_case("USD") && to.eq_ignore_ascii_case("EUR") {
                Some(ExchangeRate::new(
                    Currency {
                        code: "USD".to_string(),
                        symbol: "$".to_string(),
                    },
                    Currency {
                        code: "EUR".to_string(),
                        symbol: "€".to_string(),
                    },
                    Decimal::new(85, 2), // 0.85
                ))
            } else {
                None
            }
        }

        /// Mock save_rate does nothing
        async fn save_rate(
            &self,
            _rate: ExchangeRate,
        ) -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
            Ok(())
        }
    }

    /// Mock implementation of `ICurrencyConversionUseCase`
    struct MockCurrencyConversionUseCase;

    #[async_trait::async_trait]
    impl ICurrencyConversionUseCase for MockCurrencyConversionUseCase {
        /// Returns predefined historical rates for USD to EUR, else None
        async fn get_historical_rates(
            &self,
            base_currency: &str,
            target_currency: &str,
            start_date: &str,
            end_date: &str,
        ) -> Option<Vec<ExchangeRate>> {
            if base_currency.eq_ignore_ascii_case("USD")
                && target_currency.eq_ignore_ascii_case("EUR")
            {
                Some(vec![
                    ExchangeRate::new(
                        Currency {
                            code: "USD".to_string(),
                            symbol: "$".to_string(),
                        },
                        Currency {
                            code: "EUR".to_string(),
                            symbol: "€".to_string(),
                        },
                        Decimal::new(84, 2), // 0.84
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
                        Decimal::new(85, 2), // 0.85
                    ),
                ])
            } else {
                None
            }
        }
    }

    /// Helper function to inject `IRateFetcherUseCase` mock
    fn with_rate_fetcher(
        rate_fetcher: Arc<dyn IRateFetcherUseCase>,
    ) -> impl warp::Filter<Extract = (Arc<dyn IRateFetcherUseCase>,), Error = std::convert::Infallible>
           + Clone {
        warp::any().map(move || rate_fetcher.clone())
    }

    /// Helper function to inject `ICurrencyConversionUseCase` mock
    fn with_conversion_service(
        conversion_service: Arc<dyn ICurrencyConversionUseCase>,
    ) -> impl warp::Filter<
        Extract = (Arc<dyn ICurrencyConversionUseCase>,),
        Error = std::convert::Infallible,
    > + Clone {
        warp::any().map(move || conversion_service.clone())
    }

    #[tokio::test]
    async fn test_get_current_exchange_rate_success() {
        // Setup mocks
        let mock_rate_fetcher = Arc::new(MockRateFetcherUseCase) as Arc<dyn IRateFetcherUseCase>;
        let mock_conversion_service =
            Arc::new(MockCurrencyConversionUseCase) as Arc<dyn ICurrencyConversionUseCase>;

        // Build API filter
        let api_filter = api(mock_rate_fetcher.clone(), mock_conversion_service.clone());

        // Execute request
        let resp = request()
            .method("GET")
            .path("/rates/current?base_currency=USD&target_currency=EUR")
            .reply(&api_filter)
            .await;

        // Verify response
        assert_eq!(resp.status(), StatusCode::OK);
        let expected = json!({
            "base_currency": "USD",
            "target_currency": "EUR",
            "exchange_rate": "0.85",
            "timestamp": "2024-09-12T12:00:00Z" // Adjust as needed
        });
        let body: serde_json::Value = serde_json::from_slice(resp.body()).unwrap();
        assert_eq!(body, expected);
    }

    #[tokio::test]
    async fn test_get_current_exchange_rate_not_found() {
        // Setup mocks
        let mock_rate_fetcher = Arc::new(MockRateFetcherUseCase) as Arc<dyn IRateFetcherUseCase>;
        let mock_conversion_service =
            Arc::new(MockCurrencyConversionUseCase) as Arc<dyn ICurrencyConversionUseCase>;

        // Build API filter
        let api_filter = api(mock_rate_fetcher.clone(), mock_conversion_service.clone());

        // Execute request with non-existent currency pair
        let resp = request()
            .method("GET")
            .path("/rates/current?base_currency=GBP&target_currency=JPY")
            .reply(&api_filter)
            .await;

        // Verify response
        assert_eq!(resp.status(), StatusCode::NOT_FOUND);
    }

    #[tokio::test]
    async fn test_get_historical_exchange_rate_success() {
        // Setup mocks
        let mock_rate_fetcher = Arc::new(MockRateFetcherUseCase) as Arc<dyn IRateFetcherUseCase>;
        let mock_conversion_service =
            Arc::new(MockCurrencyConversionUseCase) as Arc<dyn ICurrencyConversionUseCase>;

        // Build API filter
        let api_filter = api(mock_rate_fetcher.clone(), mock_conversion_service.clone());

        // Execute request
        let resp = request()
            .method("GET")
            .path("/rates/historical?base_currency=USD&target_currency=EUR&start_date=2024-01-01&end_date=2024-01-02")
            .reply(&api_filter)
            .await;

        // Verify response
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
    async fn test_invalid_currency_code() {
        // Setup mocks
        let mock_rate_fetcher = Arc::new(MockRateFetcherUseCase) as Arc<dyn IRateFetcherUseCase>;
        let mock_conversion_service =
            Arc::new(MockCurrencyConversionUseCase) as Arc<dyn ICurrencyConversionUseCase>;

        // Build API filter
        let api_filter = api(mock_rate_fetcher.clone(), mock_conversion_service.clone());

        // Execute request with invalid currency codes
        let resp = request()
            .method("GET")
            .path("/rates/current?base_currency=INVALID&target_currency=EUR")
            .reply(&api_filter)
            .await;

        // Verify response
        assert_eq!(resp.status(), StatusCode::NOT_FOUND);
    }
}
