use serde::{Serialize, Deserialize};
use rust_decimal::Decimal;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ExchangeRate {
    pub from: Currency,
    pub to: Currency,
    pub rate: Decimal,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Currency {
    pub code: String,
    pub symbol: String,
}

impl ExchangeRate {
    pub fn new(from: Currency, to: Currency, rate: Decimal) -> Self {
        Self { from, to, rate }
    }
}
