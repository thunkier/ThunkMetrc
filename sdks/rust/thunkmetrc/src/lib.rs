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
    ) -> Result<Vec<Value>, Box<dyn Error>> {
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
        let mut page = 1;
        let page_size = 20; // STRICT LIMIT

        loop {
            // Arguments: (lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, body)
            // Note: The generated wrapper might have different signature order or optional vs required.
            // Based on Go/Python, usually: query params first? or Path?
            // Rust generated wrapper uses `Option` for most args.
            // Let's assume generated signature. I'll need to check `wrapper` source to be 100% sure of order, 
            // but usually it's alphabetical or definition order.
            // Checking: packages_get_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body, ...)
            
            // Using named arguments would be nicer but Rust uses positional.
            // I'll take a guess at the signature based on other SDKs, then fix if compile fails.
            // Usually: (path args..., body, query args...)
            // GetActiveV2 has NO path args.
            // So: body, query args? 
            // Or query args... ?
            // I'll assume: body, last_modified_end, last_modified_start, license_number, page_number, page_size
            
            let result_opt = self.wrapper.packages_get_active_v2(
                Some(end_str.clone()),
                Some(start_str.clone()),
                Some(license_number.to_string()),
                Some(page.to_string()),
                Some(page_size.to_string()),
                None, // body
            ).await?;

            if let Some(val) = result_opt {
                if let Some(obj) = val.as_object() {
                    if let Some(data) = obj.get("Data").and_then(|d| d.as_array()) {
                        for item in data {
                            all_packages.push(item.clone());
                        }
                    } else {
                        // If no Data or empty, break
                        break;
                    }

                    if let Some(total_pages_val) = obj.get("TotalPages").and_then(|v| v.as_i64()) {
                       if page >= total_pages_val {
                           break;
                       }
                    } else {
                        break;
                    }
                } else {
                    break;
                }
            } else {
                break;
            }

            page += 1;
        }

        Ok(all_packages)
    }
}
