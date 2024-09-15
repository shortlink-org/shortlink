use crate::domain::exchange_rate::entities::ExchangeRate;
use std::collections::HashMap;
use std::sync::Mutex;

pub struct CacheService {
    cache: Mutex<HashMap<String, ExchangeRate>>,
}

impl CacheService {
    pub fn new() -> Self {
        Self {
            cache: Mutex::new(HashMap::new()),
        }
    }

    pub fn get(&self, key: &str) -> Option<ExchangeRate> {
        let cache = self.cache.lock().unwrap();
        cache.get(key).cloned()
    }

    pub fn save(&self, key: String, rate: ExchangeRate) {
        let mut cache = self.cache.lock().unwrap();
        cache.insert(key, rate);
    }
}
