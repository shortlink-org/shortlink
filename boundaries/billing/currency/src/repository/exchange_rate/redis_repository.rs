use crate::domain::exchange_rate::entities::ExchangeRate;
use crate::repository::exchange_rate::repository::ExchangeRateRepository;
use async_trait::async_trait;
use deadpool_redis::{Config, Pool, Runtime};
use redis::cmd;
use serde_json;

pub struct RedisExchangeRateRepository {
    pool: Pool,
}

impl RedisExchangeRateRepository {
    pub async fn new(redis_url: &str) -> Result<Self, deadpool_redis::CreatePoolError> {
        // Initialize the deadpool Redis configuration
        let cfg = Config::from_url(redis_url);

        // Create the connection pool
        let pool = cfg.create_pool(Some(Runtime::Tokio1))?;

        Ok(Self { pool })
    }
}

#[async_trait]
impl ExchangeRateRepository for RedisExchangeRateRepository {
    async fn get_rate(&self, from: &str, to: &str) -> Option<ExchangeRate> {
        let key = format!("exchange_rate:{}:{}", from, to);

        // Get a connection from the pool
        let mut conn = self.pool.get().await.ok()?;

        // Execute the GET command
        let value: String = cmd("GET").arg(&key).query_async(&mut conn).await.ok()?;

        // Deserialize the JSON string into ExchangeRate
        serde_json::from_str::<ExchangeRate>(&value).ok()
    }

    async fn save_rate(&self, rate: &ExchangeRate) {
        let key = format!("exchange_rate:{}:{}", rate.from.code, rate.to.code);

        if let Ok(value) = serde_json::to_string(rate) {
            // Get a connection from the pool
            if let Ok(mut conn) = self.pool.get().await {
                // Execute the SET command
                let _: Result<(), _> = cmd("SET").arg(&key).arg(value).query_async(&mut conn).await;
                // Optionally, handle the result or log errors
            }
        }
    }
}
