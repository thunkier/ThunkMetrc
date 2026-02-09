use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct EmployeesClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> EmployeesClient<'a> {
    /// GET GetEmployees
    /// Retrieves a list of employees for a specified Facility.
    /// Permissions Required:
    /// - Manage Employees
    /// - View Employees
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_employees(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/employees/v2");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_number {
            query_params.push(format!("pageNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = page_size {
            query_params.push(format!("pageSize={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetPermissions
    /// Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
    /// Permissions Required:
    /// - Manage Employees
    /// Parameters:
    /// - employee_license_number (Option<String>): Filter by employeeLicenseNumber
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_employees_permissions(&self, employee_license_number: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/employees/v2/permissions");
        let mut query_params = Vec::new();
        if let Some(p) = employee_license_number {
            query_params.push(format!("employeeLicenseNumber={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
}

