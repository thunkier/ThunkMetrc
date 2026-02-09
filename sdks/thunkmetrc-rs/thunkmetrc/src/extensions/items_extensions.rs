use thunkmetrc_wrapper::services::items_service::ItemsService;
use thunkmetrc_wrapper::models::paginated_response::PaginatedResponse;

use thunkmetrc_wrapper::models::*;
use anyhow::Result;
use async_stream::try_stream;
use futures::{Stream, StreamExt};
use chrono::{DateTime, Utc, Duration};
use thunkmetrc_wrapper::MetrcClient;

use async_trait::async_trait;

#[async_trait]
pub trait ItemsServiceExt {

    /// Iterator for get_active_items.
    fn iterate_get_active_items<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Item>> + 'a;
    /// Sync for get_active_items.
    async fn sync_get_active_items(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Item>>;

    /// Iterator for get_brands.
    fn iterate_get_brands<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Brand>> + 'a;

    /// Iterator for get_categories.
    fn iterate_get_categories<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Category>> + 'a;

    /// Iterator for get_inactive_items.
    fn iterate_get_inactive_items<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Item>> + 'a;
}

#[async_trait]
impl ItemsServiceExt for ItemsService {

    fn iterate_get_active_items<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Item>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_active_items(
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
    async fn sync_get_active_items(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Item>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_active_items(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_brands<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Brand>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_brands(
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

    fn iterate_get_categories<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Category>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_categories(
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

    fn iterate_get_inactive_items<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Item>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_inactive_items(
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

