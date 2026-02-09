use thunkmetrc_wrapper::services::locations_service::LocationsService;
use thunkmetrc_wrapper::models::paginated_response::PaginatedResponse;

use thunkmetrc_wrapper::models::*;
use anyhow::Result;
use async_stream::try_stream;
use futures::{Stream, StreamExt};
use chrono::{DateTime, Utc, Duration};
use thunkmetrc_wrapper::MetrcClient;

use async_trait::async_trait;

#[async_trait]
pub trait LocationsServiceExt {

    /// Iterator for get_active_locations.
    fn iterate_get_active_locations<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<LocationsLocation>> + 'a;
    /// Sync for get_active_locations.
    async fn sync_get_active_locations(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<LocationsLocation>>;

    /// Iterator for get_inactive_locations.
    fn iterate_get_inactive_locations<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<LocationsLocation>> + 'a;

    /// Iterator for get_locations_types.
    fn iterate_get_locations_types<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<LocationsType>> + 'a;
}

#[async_trait]
impl LocationsServiceExt for LocationsService {

    fn iterate_get_active_locations<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<LocationsLocation>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_active_locations(
                            last_modified_end.map(|s| s.to_string()),
                            last_modified_start.map(|s| s.to_string()),
                            license_number.map(|s| s.to_string()),
                            Some(page_str.clone()),
                            page_size.map(|s| s.to_string()),).await?;

                // Handle return types
                if let Some(data) = res {
                    if data.is_empty() { break; }
                    for item in data { yield item; }
                } else { break; }
                page += 1;
            }
        }
    }
    async fn sync_get_active_locations(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<LocationsLocation>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_active_locations(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_inactive_locations<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<LocationsLocation>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_inactive_locations(
                            license_number.map(|s| s.to_string()),
                            Some(page_str.clone()),
                            page_size.map(|s| s.to_string()),).await?;

                // Handle return types
                if let Some(data) = res {
                    if data.is_empty() { break; }
                    for item in data { yield item; }
                } else { break; }
                page += 1;
            }
        }
    }

    fn iterate_get_locations_types<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<LocationsType>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_locations_types(
                            license_number.map(|s| s.to_string()),
                            Some(page_str.clone()),
                            page_size.map(|s| s.to_string()),).await?;

                // Handle return types
                if let Some(data) = res {
                    if data.is_empty() { break; }
                    for item in data { yield item; }
                } else { break; }
                page += 1;
            }
        }
    }
}

