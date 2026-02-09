use thunkmetrc_wrapper::services::sales_service::SalesService;
use thunkmetrc_wrapper::models::paginated_response::PaginatedResponse;

use thunkmetrc_wrapper::models::*;
use anyhow::Result;
use async_stream::try_stream;
use futures::{Stream, StreamExt};
use chrono::{DateTime, Utc, Duration};
use thunkmetrc_wrapper::MetrcClient;

use async_trait::async_trait;

#[async_trait]
pub trait SalesServiceExt {

    /// Iterator for get_active_deliveries.
    fn iterate_get_active_deliveries<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<ActiveDelivery>> + 'a;
    /// Sync for get_active_deliveries.
    async fn sync_get_active_deliveries(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<ActiveDelivery>>;

    /// Iterator for get_active_deliveries_retailer.
    fn iterate_get_active_deliveries_retailer<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<ActiveDeliveriesRetailer>> + 'a;
    /// Sync for get_active_deliveries_retailer.
    async fn sync_get_active_deliveries_retailer(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<ActiveDeliveriesRetailer>>;

    /// Iterator for get_active_receipts.
    fn iterate_get_active_receipts<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<ActiveReceipt>> + 'a;
    /// Sync for get_active_receipts.
    async fn sync_get_active_receipts(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<ActiveReceipt>>;

    /// Iterator for get_deliveries_return_reasons.
    fn iterate_get_deliveries_return_reasons<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<DeliveriesReturnReason>> + 'a;

    /// Iterator for get_inactive_deliveries.
    fn iterate_get_inactive_deliveries<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<InactiveDelivery>> + 'a;
    /// Sync for get_inactive_deliveries.
    async fn sync_get_inactive_deliveries(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<InactiveDelivery>>;

    /// Iterator for get_inactive_deliveries_retailer.
    fn iterate_get_inactive_deliveries_retailer<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<InactiveDeliveriesRetailer>> + 'a;
    /// Sync for get_inactive_deliveries_retailer.
    async fn sync_get_inactive_deliveries_retailer(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<InactiveDeliveriesRetailer>>;

    /// Iterator for get_inactive_receipts.
    fn iterate_get_inactive_receipts<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<InactiveReceipt>> + 'a;
    /// Sync for get_inactive_receipts.
    async fn sync_get_inactive_receipts(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<InactiveReceipt>>;
}

#[async_trait]
impl SalesServiceExt for SalesService {

    fn iterate_get_active_deliveries<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<ActiveDelivery>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_active_deliveries(
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
    async fn sync_get_active_deliveries(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<ActiveDelivery>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_active_deliveries(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_active_deliveries_retailer<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<ActiveDeliveriesRetailer>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_active_deliveries_retailer(
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
    async fn sync_get_active_deliveries_retailer(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<ActiveDeliveriesRetailer>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_active_deliveries_retailer(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_active_receipts<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<ActiveReceipt>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_active_receipts(
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
    async fn sync_get_active_receipts(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<ActiveReceipt>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_active_receipts(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_deliveries_return_reasons<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<DeliveriesReturnReason>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_deliveries_return_reasons(
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

    fn iterate_get_inactive_deliveries<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<InactiveDelivery>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_inactive_deliveries(
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
    async fn sync_get_inactive_deliveries(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<InactiveDelivery>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_inactive_deliveries(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_inactive_deliveries_retailer<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<InactiveDeliveriesRetailer>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_inactive_deliveries_retailer(
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
    async fn sync_get_inactive_deliveries_retailer(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<InactiveDeliveriesRetailer>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_inactive_deliveries_retailer(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_inactive_receipts<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<InactiveReceipt>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_inactive_receipts(
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
    async fn sync_get_inactive_receipts(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<InactiveReceipt>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_inactive_receipts(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }
}

