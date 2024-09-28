#[cfg(test)]
mod tests {
    use super::*;
    use crate::cache::CacheService;
    use crate::domain::exchange_rate::entities::{Currency, ExchangeRate};
    use rust_decimal_macros::dec;

    #[tokio::test]
    async fn test_set_and_get_rate() {
        let cache = CacheService::new();
        let rate = ExchangeRate::new(
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

        cache.set_rate(&rate).await.unwrap();
        let fetched_rate = cache.get_rate("USD", "EUR").await;
        assert!(fetched_rate.is_some());
        assert_eq!(fetched_rate.unwrap().rate, dec!(0.85));
    }

    #[tokio::test]
    async fn test_get_nonexistent_rate() {
        let cache = CacheService::new();
        let fetched_rate = cache.get_rate("GBP", "JPY").await;
        assert!(fetched_rate.is_none());
    }

    #[tokio::test]
    async fn test_update_rate() {
        let cache = CacheService::new();
        let rate1 = ExchangeRate::new(
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
        cache.set_rate(&rate1).await.unwrap();

        let rate2 = ExchangeRate::new(
            Currency {
                code: "USD".to_string(),
                symbol: "$".to_string(),
            },
            Currency {
                code: "EUR".to_string(),
                symbol: "€".to_string(),
            },
            dec!(0.90),
        );
        cache.set_rate(&rate2).await.unwrap();

        let fetched_rate = cache.get_rate("USD", "EUR").await.unwrap();
        assert_eq!(fetched_rate.rate, dec!(0.90));
    }

    #[tokio::test]
    async fn test_multiple_rates() {
        let cache = CacheService::new();
        let rate_usd_eur = ExchangeRate::new(
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
        let rate_gbp_jpy = ExchangeRate::new(
            Currency {
                code: "GBP".to_string(),
                symbol: "£".to_string(),
            },
            Currency {
                code: "JPY".to_string(),
                symbol: "¥".to_string(),
            },
            dec!(150.0),
        );

        cache.set_rate(&rate_usd_eur).await.unwrap();
        cache.set_rate(&rate_gbp_jpy).await.unwrap();

        let fetched_usd_eur = cache.get_rate("USD", "EUR").await.unwrap();
        let fetched_gbp_jpy = cache.get_rate("GBP", "JPY").await.unwrap();

        assert_eq!(fetched_usd_eur.rate, dec!(0.85));
        assert_eq!(fetched_gbp_jpy.rate, dec!(150.0));
    }
}
