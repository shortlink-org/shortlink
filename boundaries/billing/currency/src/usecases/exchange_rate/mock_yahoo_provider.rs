use super::external_rate_provider::ExternalRateProvider;
use crate::domain::exchange_rate::entities::{Currency, ExchangeRate};
use async_trait::async_trait;
use rust_decimal::Decimal;
use std::error::Error;
use std::fmt::{Debug, Formatter};

/// Mock implementation of the Yahoo exchange rate provider.
pub struct MockYahooProvider;

impl MockYahooProvider {
    pub fn new() -> Self {
        Self
    }
}

impl Debug for MockYahooProvider {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "MockYahooProvider")
    }
}

#[async_trait]
impl ExternalRateProvider for MockYahooProvider {
    async fn fetch_rate(
        &self,
        from: &str,
        to: &str,
    ) -> Result<ExchangeRate, Box<dyn Error + Send + Sync>> {
        // Simulate fetching data with different predefined mock rates.
        let rate = match (from, to) {
            ("USD", "EUR") => Decimal::new(86, 2), // Represents 0.86
            ("EUR", "GBP") => Decimal::new(76, 2), // Represents 0.76
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
