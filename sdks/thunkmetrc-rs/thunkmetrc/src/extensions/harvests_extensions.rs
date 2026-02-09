use thunkmetrc_wrapper::services::harvests_service::HarvestsService;
use thunkmetrc_wrapper::models::paginated_response::PaginatedResponse;

use thunkmetrc_wrapper::models::*;
use anyhow::Result;
use async_stream::try_stream;
use futures::{Stream, StreamExt};
use chrono::{DateTime, Utc, Duration};
use thunkmetrc_wrapper::MetrcClient;

use async_trait::async_trait;

#[async_trait]
pub trait HarvestsServiceExt {

    /// Iterator for get_active_harvests.
    fn iterate_get_active_harvests<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Harvest>> + 'a;
    /// Sync for get_active_harvests.
    async fn sync_get_active_harvests(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Harvest>>;

    /// Iterator for get_harvests_waste.
    fn iterate_get_harvests_waste<'a>(
        &'a self,
        harvest_id: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<HarvestsWaste>> + 'a;

    /// Iterator for get_inactive_harvests.
    fn iterate_get_inactive_harvests<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Harvest>> + 'a;
    /// Sync for get_inactive_harvests.
    async fn sync_get_inactive_harvests(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Harvest>>;

    /// Iterator for get_on_hold_harvests.
    fn iterate_get_on_hold_harvests<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Harvest>> + 'a;
    /// Sync for get_on_hold_harvests.
    async fn sync_get_on_hold_harvests(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Harvest>>;

    /// Iterator for get_waste_types.
    fn iterate_get_waste_types<'a>(
        &'a self,
        page_size: Option<&str>
    ) -> impl Stream<Item = Result<WasteType>> + 'a;
}

#[async_trait]
impl HarvestsServiceExt for HarvestsService {

    fn iterate_get_active_harvests<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Harvest>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_active_harvests(
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
    async fn sync_get_active_harvests(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Harvest>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_active_harvests(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_harvests_waste<'a>(
        &'a self,
        harvest_id: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<HarvestsWaste>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_harvests_waste(
                            harvest_id.map(|s| s.to_string()),
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

    fn iterate_get_inactive_harvests<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Harvest>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_inactive_harvests(
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
    async fn sync_get_inactive_harvests(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Harvest>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_inactive_harvests(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_on_hold_harvests<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Harvest>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_on_hold_harvests(
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
    async fn sync_get_on_hold_harvests(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Harvest>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_on_hold_harvests(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_waste_types<'a>(
        &'a self,
        page_size: Option<&str>
    ) -> impl Stream<Item = Result<WasteType>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_waste_types(
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

