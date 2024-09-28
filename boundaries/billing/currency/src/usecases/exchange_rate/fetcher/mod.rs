pub mod external_rate_provider;
pub mod mock_bloomberg_provider;
pub mod mock_yahoo_provider;
pub mod rate_fetcher_use_case;
pub mod traits;

// Re-export RateFetcherUseCase for easier access
pub use rate_fetcher_use_case::RateFetcherUseCase;
pub use traits::IRateFetcherUseCase;

#[cfg(test)]
pub mod tests;