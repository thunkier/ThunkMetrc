use std::collections::HashMap;
use std::sync::Arc;
use tokio::sync::{Mutex, Semaphore};
use std::time::{Duration, Instant};

#[derive(Clone)]
pub struct RateLimiterConfig {
    pub enabled: bool,
    pub max_get_per_second_per_facility: f64,
    pub max_get_per_second_integrator: f64,
    pub max_concurrent_get_per_facility: usize,
    pub max_concurrent_get_integrator: usize,
    pub max_retries: usize,
}

impl Default for RateLimiterConfig {
    fn default() -> Self {
        Self {
            enabled: false,
            max_get_per_second_per_facility: 50.0,
            max_get_per_second_integrator: 150.0,
            max_concurrent_get_per_facility: 10,
            max_concurrent_get_integrator: 30,
            max_retries: 5,
        }
    }
}

pub struct TokenBucket {
    rate: f64,
    capacity: f64,
    tokens: f64,
    last_refill: Instant,
}

impl TokenBucket {
    pub fn new(rate: f64, capacity: f64) -> Self {
        Self {
            rate,
            capacity,
            tokens: capacity,
            last_refill: Instant::now(),
        }
    }

    fn refill(&mut self) {
        let now = Instant::now();
        let elapsed = now.duration_since(self.last_refill).as_secs_f64();
        self.tokens = (self.tokens + elapsed * self.rate).min(self.capacity);
        self.last_refill = now;
    }

    pub fn try_consume(&mut self, amount: f64) -> Option<Duration> {
        self.refill();
        if self.tokens >= amount {
            self.tokens -= amount;
            None
        } else {
            let missing = amount - self.tokens;
            Some(Duration::from_secs_f64(missing / self.rate))
        }
    }
}

pub struct MetrcRateLimiter {
    config: RateLimiterConfig,
    integrator_rate: Mutex<TokenBucket>,
    integrator_sem: Arc<Semaphore>,
    facility_rates: Mutex<HashMap<String, Arc<Mutex<TokenBucket>>>>,
    facility_sems: Mutex<HashMap<String, Arc<Semaphore>>>,
}

impl MetrcRateLimiter {
    pub fn new(config: Option<RateLimiterConfig>) -> Self {
        let config = config.unwrap_or_default();
        Self {
            integrator_rate: Mutex::new(TokenBucket::new(config.max_get_per_second_integrator, config.max_get_per_second_integrator)),
            integrator_sem: Arc::new(Semaphore::new(config.max_concurrent_get_integrator)),
            facility_rates: Mutex::new(HashMap::new()),
            facility_sems: Mutex::new(HashMap::new()),
            config,
        }
    }

    pub async fn execute<F, Fut, T>(&self, facility: Option<&str>, is_get: bool, op: F) -> Result<T, Box<dyn std::error::Error + Send + Sync>>
    where
        F: Fn() -> Fut,
        Fut: std::future::Future<Output = Result<T, Box<dyn std::error::Error + Send + Sync>>>,
    {
        if !self.config.enabled || !is_get {
            return op().await;
        }

        // 1. Integrator Semaphore
        let _sem_permit = self.integrator_sem.acquire().await.map_err(|e| format!("Semaphore closed: {:?}", e)).unwrap();

        // 2. Facility Semaphore
        let _fac_permit = if let Some(f) = facility {
             let sem = {
                 let mut sems = self.facility_sems.lock().await;
                 sems.entry(f.to_string())
                     .or_insert_with(|| Arc::new(Semaphore::new(self.config.max_concurrent_get_per_facility)))
                     .clone()
             };
             Some(sem.acquire_owned().await.map_err(|e| format!("Semaphore closed: {:?}", e)).unwrap())
        } else {
             None
        };

        // 3. Global Rate
        loop {
            let mut bucket = self.integrator_rate.lock().await;
            if let Some(wait) = bucket.try_consume(1.0) {
                 drop(bucket);
                 tokio::time::sleep(wait).await;
            } else {
                 break;
            }
        }

        // 4. Facility Rate
        if let Some(f) = facility {
            let bucket_arc = {
                let mut rates = self.facility_rates.lock().await;
                rates.entry(f.to_string())
                    .or_insert_with(|| Arc::new(Mutex::new(TokenBucket::new(self.config.max_get_per_second_per_facility, self.config.max_get_per_second_per_facility))))
                    .clone()
            };
            loop {
                 let mut bucket = bucket_arc.lock().await;
                 if let Some(wait) = bucket.try_consume(1.0) {
                      drop(bucket);
                      tokio::time::sleep(wait).await;
                 } else {
                      break;
                 }
            }
        }

        // 5. Retry Loop with max retries
        let mut retry_count = 0;
        loop {
             let res = op().await;
             match res {
                 Ok(v) => return Ok(v),
                 Err(e) => {
                      if retry_count >= self.config.max_retries {
                          return Err(e);
                      }
                      retry_count += 1;

                      // Check for ApiError to respect Retry-After
                      // e is Box<dyn Error + Send + Sync>
                      if let Some(api_err) = e.downcast_ref::<thunkmetrc_client::ApiError>() {
                          if api_err.status == reqwest::StatusCode::TOO_MANY_REQUESTS {
                               if let Some(retry_after) = api_err.headers.get("Retry-After") {
                                   if let Ok(val_str) = retry_after.to_str() {
                                       if let Ok(secs) = val_str.parse::<u64>() {
                                           tokio::time::sleep(Duration::from_secs(secs)).await;
                                           continue;
                                       }
                                   }
                               }
                               // Default 429 wait
                               tokio::time::sleep(Duration::from_secs(1)).await;
                               continue;
                          }
                          // Handle 5xx with exponential backoff
                          if api_err.status.is_server_error() {
                              tokio::time::sleep(Duration::from_millis(500)).await;
                              continue; 
                          }
                      }
                      
                      let msg = format!("{:?}", e);
                      if msg.contains("429") {
                           tokio::time::sleep(Duration::from_secs(1)).await;
                           continue;
                      }
                      return Err(e);
                 }
             }
        }
    }
}
