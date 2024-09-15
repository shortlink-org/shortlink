use crate::domain::exchange_rate::entities::Currency;

#[derive(Debug, Clone)]
pub struct CurrencyPair {
    pub base: Currency,
    pub quote: Currency,
}
