mod domain;
mod usecases;
mod repository;
mod cache;

use rust_decimal_macros::dec;
use domain::exchange_rate::entities::{Currency, ExchangeRate};
use domain::currency_conversion::entities::Amount;
use usecases::exchange_rate::fetcher::RateFetcherUseCase;
use usecases::currency_conversion::converter::CurrencyConversionUseCase;
use repository::exchange_rate::in_memory_repository::InMemoryExchangeRateRepository;

#[tokio::main]
async fn main() {
    // Create repository, cache, and use cases
    let exchange_rate_repository = InMemoryExchangeRateRepository::new();
    let rate_fetcher_use_case = RateFetcherUseCase::new(exchange_rate_repository);
    let currency_conversion_use_case = CurrencyConversionUseCase::new(&rate_fetcher_use_case); // Pass reference

    // Example rate to save
    let usd_to_eur_rate = ExchangeRate::new(
        Currency { code: "USD".to_string(), symbol: "$".to_string() },
        Currency { code: "EUR".to_string(), symbol: "€".to_string() },
        dec!(0.85), // Example rate
    );

    rate_fetcher_use_case.save_rate(usd_to_eur_rate).await;

    // Example amount to convert
    let amount = Amount {
        currency: "USD".to_string(),
        value: dec!(100.0), // Use Decimal for the value
    };

    // Perform conversion
    if let Some(converted) = currency_conversion_use_case.convert(amount, "EUR").await {
        println!("Converted amount: {} {}", converted.value, converted.currency);
    } else {
        println!("Conversion failed.");
    }
}
