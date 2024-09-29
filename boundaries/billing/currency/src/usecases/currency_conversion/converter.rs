use std::fmt::{Debug, Formatter};
use crate::domain::currency_conversion::entities::{Amount, ConvertedAmount};
use crate::domain::exchange_rate::entities::ExchangeRate;
use crate::usecases::currency_conversion::traits::ICurrencyConversionUseCase;
use crate::usecases::exchange_rate::RateFetcherUseCase;
use async_trait::async_trait;
use std::sync::Arc;
use rust_decimal::Decimal;
use tracing::{error, info};

pub struct CurrencyConversionUseCase {
    pub rate_fetcher: Arc<RateFetcherUseCase>,
    pub divergence_threshold: f64, // Divergence threshold percentage, e.g., 0.5%
    pub bloomberg_weight: f64,     // Weight for Bloomberg data, e.g., 0.7
    pub yahoo_weight: f64,         // Weight for Yahoo data, e.g., 0.3
}

impl CurrencyConversionUseCase {
    /// Creates a new instance of CurrencyConversionUseCase.
    pub fn new(
        rate_fetcher: Arc<RateFetcherUseCase>,
        divergence_threshold: f64,
        bloomberg_weight: f64,
        yahoo_weight: f64,
    ) -> Self {
        Self {
            rate_fetcher,
            divergence_threshold,
            bloomberg_weight,
            yahoo_weight,
        }
    }

    /// Converts an amount using the weighted average approach to handle discrepancies.
    pub async fn convert(&self, amount: Amount, to_currency: &str) -> Option<ConvertedAmount> {
        let bloomberg_rate = self
            .rate_fetcher
            .fetch_rate_from_provider("Bloomberg", &amount.currency, to_currency)
            .await;

        let yahoo_rate = self
            .rate_fetcher
            .fetch_rate_from_provider("Yahoo", &amount.currency, to_currency)
            .await;

        if let (Some(bloomberg_rate), Some(yahoo_rate)) = (bloomberg_rate, yahoo_rate) {
            // Check for divergence
            let rate_difference = (bloomberg_rate.rate - yahoo_rate.rate).abs();
            let average_rate = (bloomberg_rate.rate + yahoo_rate.rate) / Decimal::from(2); // Divide by Decimal::from(2)

            let divergence = (rate_difference / average_rate) * Decimal::from(100); // Convert the percentage

            if divergence > Decimal::from_f64_retain(self.divergence_threshold).unwrap_or(Decimal::ZERO) {
                error!(
                    "Divergence of {}% exceeds the threshold of {}%. Triggering alert.",
                    divergence, self.divergence_threshold
                );
                // Trigger an alert for manual intervention (log and return error)
                return None;
            }

            // Apply weighted average
            let weighted_average_rate = (bloomberg_rate.rate * Decimal::from_f64_retain(self.bloomberg_weight).unwrap_or(Decimal::ZERO))
                + (yahoo_rate.rate * Decimal::from_f64_retain(self.yahoo_weight).unwrap_or(Decimal::ZERO));

            let converted_value = amount.value * weighted_average_rate;

            info!(
                "Converted amount using weighted average: {} to {}, rate: {}",
                amount.currency, to_currency, weighted_average_rate
            );

            Some(ConvertedAmount::new(to_currency.to_string(), converted_value))
        } else {
            None
        }
    }
}

impl Debug for CurrencyConversionUseCase {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("CurrencyConversionUseCase")
            .field("rate_fetcher", &self.rate_fetcher)  // Assuming rate_fetcher implements Debug
            .field("divergence_threshold", &self.divergence_threshold)
            .field("bloomberg_weight", &self.bloomberg_weight)
            .field("yahoo_weight", &self.yahoo_weight)
            .finish()
    }
}

#[async_trait]
impl ICurrencyConversionUseCase for CurrencyConversionUseCase {
    async fn get_historical_rates(
        &self,
        base_currency: &str,
        target_currency: &str,
        start_date: &str,
        end_date: &str,
    ) -> Option<Vec<ExchangeRate>> {
        self.rate_fetcher
            .get_historical_rates(base_currency, target_currency, start_date, end_date)
            .await
    }
}
