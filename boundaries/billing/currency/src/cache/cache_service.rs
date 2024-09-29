use crate::domain::exchange_rate::entities::ExchangeRate;
use std::collections::HashMap;
use std::error::Error;
use std::sync::Arc;
use tokio::sync::Mutex;

#[derive(Default, Clone, Debug)]
pub struct CacheService {
    store: Arc<Mutex<HashMap<String, ExchangeRate>>>,
}

impl CacheService {
    pub fn new() -> Self {
        Self {
            store: Arc::new(Mutex::new(HashMap::new())),
        }
    }

    /// Retrieves an exchange rate from the cache.
    pub async fn get_rate(&self, from: &str, to: &str) -> Option<ExchangeRate> {
        let key = format!("exchange_rate:{}:{}", from, to);
        let store = self.store.lock().await;
        store.get(&key).cloned()
    }

    /// Stores an exchange rate in the cache.
    pub async fn set_rate(&self, rate: &ExchangeRate) -> Result<(), Box<dyn Error + Send + Sync>> {
        let key = format!("exchange_rate:{}:{}", rate.from.code, rate.to.code);
        let mut store = self.store.lock().await;
        store.insert(key, rate.clone());
        Ok(())
    }
}
