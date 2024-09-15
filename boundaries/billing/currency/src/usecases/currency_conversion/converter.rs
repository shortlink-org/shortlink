use crate::repository::exchange_rate::repository::ExchangeRateRepository;
use crate::domain::currency_conversion::entities::{Amount, ConvertedAmount};
use crate::usecases::exchange_rate::fetcher::RateFetcherUseCase;

pub struct CurrencyConversionUseCase<'a, R: ExchangeRateRepository> {
    pub rate_fetcher: &'a RateFetcherUseCase<R>, // Use a reference here
}

impl<'a, R: ExchangeRateRepository> CurrencyConversionUseCase<'a, R> {
    pub fn new(rate_fetcher: &'a RateFetcherUseCase<R>) -> Self {
        Self { rate_fetcher }
    }

    pub async fn convert(&self, amount: Amount, to_currency: &str) -> Option<ConvertedAmount> {
        if let Some(rate) = self.rate_fetcher.fetch_rate(&amount.currency, to_currency).await {
            let converted_value = amount.value * rate.rate;
            Some(ConvertedAmount::new(to_currency.to_string(), converted_value))
        } else {
            None
        }
    }
}
