use rust_decimal::Decimal;

#[derive(Debug, Clone)]
pub struct ExchangeRate {
    pub from_currency: Currency,
    pub to_currency: Currency,
    pub rate: Decimal,
}

#[derive(Debug, Clone)]
pub struct Currency {
    pub code: String,
    pub symbol: String,
}

impl ExchangeRate {
    pub fn new(from_currency: Currency, to_currency: Currency, rate: Decimal) -> Self {
        Self {
            from_currency,
            to_currency,
            rate,
        }
    }
}
