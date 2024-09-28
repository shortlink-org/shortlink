use super::external_rate_provider::ExternalRateProvider;
use crate::domain::exchange_rate::entities::{Currency, ExchangeRate};
use async_trait::async_trait;
use rust_decimal::Decimal;
use std::error::Error;

/// Mock implementation of the Bloomberg exchange rate provider.
pub struct MockBloombergProvider;

impl MockBloombergProvider {
    pub fn new() -> Self {
        Self
    }
}

#[async_trait]
impl ExternalRateProvider for MockBloombergProvider {
    async fn fetch_rate(
        &self,
        from: &str,
        to: &str,
    ) -> Result<ExchangeRate, Box<dyn Error + Send + Sync>> {
        // Simulate fetching data with predefined mock rates.
        let rate = match (from, to) {
            ("USD", "EUR") => Decimal::new(85, 2), // Represents 0.85
            ("EUR", "GBP") => Decimal::new(75, 2), // Represents 0.75
            _ => Decimal::new(100, 2),             // Represents 1.00
        };

        Ok(ExchangeRate::new(
            Currency {
                code: from.to_string(),
                symbol: match from {
                    "USD" => "$".to_string(),
                    "EUR" => "€".to_string(),
                    "GBP" => "£".to_string(),
                    _ => "".to_string(),
                },
            },
            Currency {
                code: to.to_string(),
                symbol: match to {
                    "USD" => "$".to_string(),
                    "EUR" => "€".to_string(),
                    "GBP" => "£".to_string(),
                    _ => "".to_string(),
                },
            },
            rate,
        ))
    }
}
