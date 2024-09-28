use crate::domain::exchange_rate::entities::ExchangeRate;
use crate::repository::exchange_rate::repository::ExchangeRateRepository;
use async_trait::async_trait;
use deadpool_redis::{Config, Pool, Runtime};
use redis::cmd;
use serde::de::Error;
use serde_json;
use chrono::NaiveDate;
use std::error::Error as StdError;

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

    fn make_historical_key(
        &self,
        base_currency: &str,
        target_currency: &str,
        date: &str,
    ) -> String {
        format!("historical_exchange_rate:{}:{}:{}", base_currency, target_currency, date)
    }
}

#[async_trait]
impl ExchangeRateRepository for RedisExchangeRateRepository {
    async fn save_rate(&self, rate: &ExchangeRate) -> Result<(), Box<dyn StdError + Send + Sync>> {
        let key = format!("exchange_rate:{}:{}", rate.from.code, rate.to.code);

        if let Ok(value) = serde_json::to_string(rate) {
            // Get a connection from the pool
            let mut conn = self.pool.get().await.map_err(|e| {
                Box::new(e) as Box<dyn StdError + Send + Sync>
            })?;

            // Execute the SET command and handle any errors
            cmd("SET")
                .arg(&key)
                .arg(value)
                .query_async(&mut conn)
                .await
                .map_err(|e| Box::new(e) as Box<dyn StdError + Send + Sync>)?;

            Ok(())
        } else {
            Err(Box::new(serde_json::Error::custom("Failed to serialize rate")))
        }
    }

    async fn get_rate(&self, from: &str, to: &str) -> Option<ExchangeRate> {
        let key = format!("exchange_rate:{}:{}", from, to);

        // Get a connection from the pool
        let mut conn = self.pool.get().await.ok()?;

        // Execute the GET command
        let value: String = cmd("GET").arg(&key).query_async(&mut conn).await.ok()?;

        // Deserialize the JSON string into ExchangeRate
        serde_json::from_str::<ExchangeRate>(&value).ok()
    }

    async fn get_historical_rates(
        &self,
        base_currency: &str,
        target_currency: &str,
        start_date: &str,
        end_date: &str,
    ) -> Option<Vec<ExchangeRate>> {
        let mut conn = self.pool.get().await.ok()?;
        let mut rates = Vec::new();

        // Parse the start_date and end_date as NaiveDate (from the chrono crate)
        let start = NaiveDate::parse_from_str(start_date, "%Y-%m-%d").ok()?;
        let end = NaiveDate::parse_from_str(end_date, "%Y-%m-%d").ok()?;

        // Iterate over the range of dates using NaiveDate's iter_days() method
        for date in start.iter_days().take_while(|&d| d <= end) {
            let key = self.make_historical_key(base_currency, target_currency, &date.to_string());

            // Fetch the rate for the current date
            if let Ok(Some(value)) = cmd("GET").arg(&key).query_async::<Option<String>>(&mut conn).await {
                if let Ok(rate) = serde_json::from_str::<ExchangeRate>(&value) {
                    rates.push(rate);
                }
            }
        }

        if rates.is_empty() {
            None
        } else {
            Some(rates)
        }
    }
}
