use crate::domain::currency_conversion::entities::{Amount, ConvertedAmount};
use crate::usecases::exchange_rate::fetcher::RateFetcherUseCase;
use std::sync::Arc;
use async_trait::async_trait;
use crate::usecases::currency_conversion::converter::traits::ICurrencyConversionUseCase;
use crate::domain::exchange_rate::entities::ExchangeRate;

pub struct CurrencyConversionUseCase {
    pub rate_fetcher: Arc<RateFetcherUseCase>,
}

impl CurrencyConversionUseCase {
    pub fn new(rate_fetcher: Arc<RateFetcherUseCase>) -> Self {
        Self { rate_fetcher }
    }

    pub async fn convert(&self, amount: Amount, to_currency: &str) -> Option<ConvertedAmount> {
        if let Some(rate) = self
            .rate_fetcher
            .fetch_rate(&amount.currency, to_currency)
            .await
        {
            let converted_value = amount.value * rate.rate;
            Some(ConvertedAmount::new(
                to_currency.to_string(),
                converted_value,
            ))
        } else {
            None
        }
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
        self.get_historical_rates(base_currency, target_currency, start_date, end_date).await
    }
}