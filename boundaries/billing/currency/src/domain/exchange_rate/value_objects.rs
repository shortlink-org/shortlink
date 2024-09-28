use crate::domain::exchange_rate::entities::Currency;

pub struct CurrencyPair {
    pub base: Currency,
    pub quote: Currency,
}
