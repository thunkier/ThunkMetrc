use thunkmetrc_wrapper::services::employees_service::EmployeesService;
use thunkmetrc_wrapper::models::paginated_response::PaginatedResponse;

use thunkmetrc_wrapper::models::*;
use anyhow::Result;
use async_stream::try_stream;
use futures::{Stream, StreamExt};
use chrono::{DateTime, Utc, Duration};
use thunkmetrc_wrapper::MetrcClient;

use async_trait::async_trait;

#[async_trait]
pub trait EmployeesServiceExt {

    /// Iterator for get_employees.
    fn iterate_get_employees<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Employee>> + 'a;
}

#[async_trait]
impl EmployeesServiceExt for EmployeesService {

    fn iterate_get_employees<'a>(
        &'a self,
        license_number: Option<&str>, page_size: Option<&str>
    ) -> impl Stream<Item = Result<Employee>> + 'a {
        try_stream! {
            let mut page = 1;
            loop {
                let page_str = page.to_string();
                let res = self.get_employees(
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

