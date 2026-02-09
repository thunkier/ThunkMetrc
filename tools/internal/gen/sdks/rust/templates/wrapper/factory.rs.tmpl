use std::sync::Arc;
use crate::{MetrcClient, MetrcRateLimiter, RateLimiterConfig, MetrcWrapper};

pub struct MetrcFactory {
    shared_limiter: Arc<MetrcRateLimiter>,
}

impl MetrcFactory {
    pub fn new(max_concurrent_requests: usize) -> Self {
        let config = RateLimiterConfig { max_concurrent_requests };
        MetrcFactory {
            shared_limiter: Arc::new(MetrcRateLimiter::new(Some(config))),
        }
    }

    pub fn create(&self, base_url: &str, vendor_key: &str, user_key: &str) -> MetrcWrapper {
        let client = MetrcClient::new(base_url, vendor_key, user_key);
        MetrcWrapper::new_with_limiter(client, self.shared_limiter.clone())
    }
}
