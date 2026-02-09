use thunkmetrc_wrapper::MetrcWrapper;
use chrono::{DateTime, Utc, Duration};
use serde_json::Value;
use std::error::Error;

pub struct ThunkMetrc<'a> {
    pub wrapper: &'a MetrcWrapper,
}

impl<'a> ThunkMetrc<'a> {
    pub fn new(wrapper: &'a MetrcWrapper) -> Self {
        Self { wrapper }
    }

    /// Sync active packages inventory.
    /// 
    /// # Arguments
    ///
    /// * `license_number` - Facility license number.
    /// * `last_known_sync` - Optional last sync time.
    /// * `buffer_minutes` - Buffer in minutes (default 5).
    pub async fn active_packages_inventory_sync(
        &self,
        license_number: &str,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
    ) -> Result<Vec<Value>, Box<dyn Error + Send + Sync>> {
        let end_time = Utc::now();
        let start_time = match last_known_sync {
            Some(t) => t - Duration::minutes(buffer_minutes),
            None => end_time - Duration::hours(24),
        };

        // Format: RFC3339 (ISO 8601)
        // Metrc handles RFC3339 correctly.
        let start_str = start_time.to_rfc3339();
        let end_str = end_time.to_rfc3339();

        let mut all_packages = Vec::new();
        let mut page_number: i64 = 1;
        let page_size: i64 = 20;

        loop {
            let result = self.wrapper.packages.get_active_packages(
                Some(end_str.clone()),
                Some(start_str.clone()),
                Some(license_number.to_string()),
                Some(page_number.to_string()),
                Some(page_size.to_string()),
                None,
            ).await?;

            let page = match result {
                Some(page) => page,
                None => break,
            };

            let items = page.data.unwrap_or_default();
            if items.is_empty() {
                break;
            }

            for item in items {
                all_packages.push(serde_json::to_value(item)?);
            }

            page_number += 1;
        }

        Ok(all_packages)
    }
}
