use thunkmetrc_wrapper::services::transfers_service::TransfersService;
use thunkmetrc_wrapper::models::paginated_response::PaginatedResponse;

use thunkmetrc_wrapper::models::*;
use anyhow::Result;
use async_stream::try_stream;
use futures::{Stream, StreamExt};
use chrono::{DateTime, Utc, Duration};
use thunkmetrc_wrapper::MetrcClient;

use async_trait::async_trait;

#[async_trait]
pub trait TransfersServiceExt {

    /// Iterator for get_delivery_package_by_id.
    fn iterate_get_delivery_package_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<DeliveryPackage>> + 'a;

    /// Iterator for get_delivery_package_requiredlabtestbatch_by_id.
    fn iterate_get_delivery_package_requiredlabtestbatch_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<DeliveryPackageRequiredlabtestbatch>> + 'a;

    /// Iterator for get_delivery_package_wholesale_by_id.
    fn iterate_get_delivery_package_wholesale_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<DeliveryPackageWholesale>> + 'a;

    /// Iterator for get_delivery_transporter_by_id.
    fn iterate_get_delivery_transporter_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<DeliveryTransporter>> + 'a;

    /// Iterator for get_delivery_transporter_detail_by_id.
    fn iterate_get_delivery_transporter_detail_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<DeliveryTransporterDetail>> + 'a;

    /// Iterator for get_hub.
    fn iterate_get_hub<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Hub>> + 'a;
    /// Sync for get_hub.
    async fn sync_get_hub(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Hub>>;

    /// Iterator for get_incoming_transfers.
    fn iterate_get_incoming_transfers<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Transfer>> + 'a;
    /// Sync for get_incoming_transfers.
    async fn sync_get_incoming_transfers(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Transfer>>;

    /// Iterator for get_outgoing_template_delivery_by_id.
    fn iterate_get_outgoing_template_delivery_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<TemplateDelivery>> + 'a;

    /// Iterator for get_outgoing_template_delivery_package_by_id.
    fn iterate_get_outgoing_template_delivery_package_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<TemplateDeliveryPackage>> + 'a;

    /// Iterator for get_outgoing_template_delivery_transporter_by_id.
    fn iterate_get_outgoing_template_delivery_transporter_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<TemplateDeliveryTransporter>> + 'a;

    /// Iterator for get_outgoing_templates.
    fn iterate_get_outgoing_templates<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Template>> + 'a;
    /// Sync for get_outgoing_templates.
    async fn sync_get_outgoing_templates(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Template>>;

    /// Iterator for get_outgoing_transfers.
    fn iterate_get_outgoing_transfers<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Transfer>> + 'a;
    /// Sync for get_outgoing_transfers.
    async fn sync_get_outgoing_transfers(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Transfer>>;

    /// Iterator for get_rejected_transfers.
    fn iterate_get_rejected_transfers<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Transfer>> + 'a;

    /// Iterator for get_transfers_delivery_by_id.
    fn iterate_get_transfers_delivery_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<TransfersDelivery>> + 'a;

    /// Iterator for get_transfers_types.
    fn iterate_get_transfers_types<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<TransfersType>> + 'a;
}

#[async_trait]
impl TransfersServiceExt for TransfersService {

    fn iterate_get_delivery_package_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<DeliveryPackage>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_delivery_package_by_id(id, 
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

    fn iterate_get_delivery_package_requiredlabtestbatch_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<DeliveryPackageRequiredlabtestbatch>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_delivery_package_requiredlabtestbatch_by_id(id, 
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

    fn iterate_get_delivery_package_wholesale_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<DeliveryPackageWholesale>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_delivery_package_wholesale_by_id(id, 
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

    fn iterate_get_delivery_transporter_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<DeliveryTransporter>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_delivery_transporter_by_id(id, 
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

    fn iterate_get_delivery_transporter_detail_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<DeliveryTransporterDetail>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_delivery_transporter_detail_by_id(id, 
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

    fn iterate_get_hub<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Hub>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_hub(
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
    async fn sync_get_hub(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Hub>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_hub(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_incoming_transfers<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Transfer>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_incoming_transfers(
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
    async fn sync_get_incoming_transfers(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Transfer>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_incoming_transfers(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_outgoing_template_delivery_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<TemplateDelivery>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_outgoing_template_delivery_by_id(id, 
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

    fn iterate_get_outgoing_template_delivery_package_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<TemplateDeliveryPackage>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_outgoing_template_delivery_package_by_id(id, 
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

    fn iterate_get_outgoing_template_delivery_transporter_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<TemplateDeliveryTransporter>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_outgoing_template_delivery_transporter_by_id(id, 
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

    fn iterate_get_outgoing_templates<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Template>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_outgoing_templates(
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
    async fn sync_get_outgoing_templates(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Template>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_outgoing_templates(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_outgoing_transfers<'a>(
        &'a self,
        last_modified_end: Option<&str>, last_modified_start: Option<&str>, license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Transfer>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_outgoing_transfers(
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
    async fn sync_get_outgoing_transfers(
        &self,
        last_known_sync: Option<DateTime<Utc>>,
        buffer_minutes: i64,
            license_number: Option<&str>,
            page_size: Option<&str>,
    ) -> Result<Vec<Transfer>> {
        let end = Utc::now();
        let mut start = end - Duration::days(1);
        if let Some(lks) = last_known_sync {
            start = lks - Duration::minutes(buffer_minutes);
        }

        let start_str = start.to_rfc3339();
        let end_str = end.to_rfc3339();
        
        let mut results = Vec::new();
        let mut stream = self.iterate_get_outgoing_transfers(Some(&end_str),Some(&start_str),
            license_number,
            page_size,
            // body?
        );
        
        while let Some(item_res) = stream.next().await {
            results.push(item_res?);
        }
        
        Ok(results)
    }

    fn iterate_get_rejected_transfers<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Transfer>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_rejected_transfers(
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

    fn iterate_get_transfers_delivery_by_id<'a>(
        &'a self,
        id: &str, page_size: Option<&str>
    ) -> impl Stream<Item = Result<TransfersDelivery>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_transfers_delivery_by_id(id, 
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

    fn iterate_get_transfers_types<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<TransfersType>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_transfers_types(
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

