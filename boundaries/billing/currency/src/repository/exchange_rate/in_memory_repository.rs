use crate::domain::exchange_rate::entities::ExchangeRate;
use crate::repository::exchange_rate::repository::ExchangeRateRepository;
use async_trait::async_trait;
use std::collections::HashMap;
use std::fmt::{Debug, Formatter};
use std::sync::Mutex;
use std::sync::Arc;

pub struct InMemoryExchangeRateRepository {
    rates: Arc<Mutex<HashMap<String, ExchangeRate>>>,  // Store historical rates in memory
}

impl InMemoryExchangeRateRepository {
    pub fn new() -> Self {
        Self {
            rates: Arc::new(Mutex::new(HashMap::new())),
        }
    }

    fn make_key(&self, base_currency: &str, target_currency: &str, date: &str) -> String {
        format!("{}:{}:{}", base_currency, target_currency, date)
    }
}

impl Debug for InMemoryExchangeRateRepository {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "InMemoryExchangeRateRepository")
    }
}

#[async_trait]
impl ExchangeRateRepository for InMemoryExchangeRateRepository {
    async fn save_rate(&self, rate: &ExchangeRate) -> Result<(), Box<dyn std::error::Error + Send + Sync>> {
        let key = self.make_key(&rate.from.code, &rate.to.code, "current");
        self.rates.lock().unwrap().insert(key, rate.clone());
        Ok(())
    }

    async fn get_rate(&self, from: &str, to: &str) -> Option<ExchangeRate> {
        let key = self.make_key(from, to, "current");
        self.rates.lock().unwrap().get(&key).cloned()
    }

    async fn get_historical_rates(
        &self,
        base_currency: &str,
        target_currency: &str,
        start_date: &str,
        end_date: &str,
    ) -> Option<Vec<ExchangeRate>> {
        let rates = self.rates.lock().unwrap();
        let mut result = Vec::new();

        // Just an example: iterate over rates to find ones matching the date range
        for (key, rate) in rates.iter() {
            // For simplicity, assume the key contains the date information
            if key.contains(base_currency) && key.contains(target_currency) {
                result.push(rate.clone());
            }
        }

        if result.is_empty() {
            None
        } else {
            Some(result)
        }
    }
}
