use rust_decimal::Decimal;

pub struct Amount {
    pub currency: String,
    pub value: Decimal,
}

pub struct ConvertedAmount {
    pub currency: String,
    pub value: Decimal,
}

impl ConvertedAmount {
    pub fn new(currency: String, value: Decimal) -> Self {
        Self { currency, value }
    }
}
