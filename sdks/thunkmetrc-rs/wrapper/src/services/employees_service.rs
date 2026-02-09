use thunkmetrc_client::MetrcClient;
use serde_json::Value;
use std::error::Error;
use std::sync::Arc;
#[allow(unused_imports)]
use futures::Stream;
use crate::ratelimiter::MetrcRateLimiter;
#[allow(unused_imports)]
use crate::utils::iterate_pages;
#[allow(unused_imports)]
use crate::models::*;

pub struct EmployeesService {
    client: MetrcClient,
    rate_limiter: Arc<MetrcRateLimiter>,
}

impl EmployeesService {
    pub fn new(client: MetrcClient, rate_limiter: Arc<MetrcRateLimiter>) -> Self {
        Self {
            client,
            rate_limiter,
        }
    }

    /// GET GetEmployees
    /// Retrieves a list of employees for a specified Facility.
    /// Permissions Required:
    /// - Manage Employees
    /// - View Employees
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_employees(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<Employee>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let page_number = page_number.clone();
                let page_size = page_size.clone();
                let body_val = body_val.clone();

                async move {
                    client.employees().get_employees(license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<Employee> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetPermissions
    /// Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
    /// Permissions Required:
    /// - Manage Employees
    /// Parameters:
    /// - employee_license_number (Option<String>): Filter by employeeLicenseNumber
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_employees_permissions(&self, employee_license_number: Option<String>, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let employee_license_number = employee_license_number.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.employees().get_employees_permissions(employee_license_number, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: serde_json::Value = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }
}

