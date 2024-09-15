use crate::domain::exchange_rate::entities::ExchangeRate;
use crate::repository::exchange_rate::repository::ExchangeRateRepository;
use async_trait::async_trait;
use std::collections::HashMap;
use std::sync::Mutex;

pub struct InMemoryExchangeRateRepository {
    store: Mutex<HashMap<(String, String), ExchangeRate>>,
}

impl InMemoryExchangeRateRepository {
    pub fn new() -> Self {
        InMemoryExchangeRateRepository {
            store: Mutex::new(HashMap::new()),
        }
    }
}

#[async_trait]
impl ExchangeRateRepository for InMemoryExchangeRateRepository {
    async fn get_rate(&self, from: &str, to: &str) -> Option<ExchangeRate> {
        let store = self.store.lock().unwrap();
        store.get(&(from.to_string(), to.to_string())).cloned()
    }

    async fn save_rate(&self, rate: &ExchangeRate) {
        let mut store = self.store.lock().unwrap();
        store.insert(
            (rate.from_currency.code.clone(), rate.to_currency.code.clone()),
            rate.clone(),
        );
    }
}
