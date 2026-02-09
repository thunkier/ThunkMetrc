use thunkmetrc_wrapper::services::lab_tests_service::LabTestsService;
use thunkmetrc_wrapper::models::paginated_response::PaginatedResponse;

use thunkmetrc_wrapper::models::*;
use anyhow::Result;
use async_stream::try_stream;
use futures::{Stream, StreamExt};
use chrono::{DateTime, Utc, Duration};
use thunkmetrc_wrapper::MetrcClient;

use async_trait::async_trait;

#[async_trait]
pub trait LabTestsServiceExt {

    /// Iterator for get_batches.
    fn iterate_get_batches<'a>(
        &'a self,
        page_size: Option<&str>
    ) -> impl Stream<Item = Result<Batch>> + 'a;

    /// Iterator for get_lab_tests_types.
    fn iterate_get_lab_tests_types<'a>(
        &'a self,
        page_size: Option<&str>
    ) -> impl Stream<Item = Result<LabTestsType>> + 'a;

    /// Iterator for get_results.
    fn iterate_get_results<'a>(
        &'a self,
        license_number: Option<&str>, package_id: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Result>> + 'a;
}

#[async_trait]
impl LabTestsServiceExt for LabTestsService {

    fn iterate_get_batches<'a>(
        &'a self,
        page_size: Option<&str>
    ) -> impl Stream<Item = Result<Batch>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_batches(
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

    fn iterate_get_lab_tests_types<'a>(
        &'a self,
        page_size: Option<&str>
    ) -> impl Stream<Item = Result<LabTestsType>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_lab_tests_types(
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

    fn iterate_get_results<'a>(
        &'a self,
        license_number: Option<&str>, package_id: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Result>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_results(
                            license_number.map(|s| s.to_string()),
                            package_id.map(|s| s.to_string()),
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

