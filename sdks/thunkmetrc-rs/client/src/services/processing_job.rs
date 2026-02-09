use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct ProcessingJobClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> ProcessingJobClient<'a> {
    /// POST CreateAdjustProcessingJob
    /// Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_adjust_processing_job(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/adjust");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// POST CreateJobTypes
    /// Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_processing_job_job_types(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/jobtypes");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// POST CreateProcessingJobPackages
    /// Creates packages from processing jobs at a Facility, including optional location and note assignments.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_processing_job_packages(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/createpackages");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// DELETE DeleteJobTypeById
    /// Archives a Processing Job Type at a Facility, making it inactive for future use.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_processing_job_job_type_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/jobtypes/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// DELETE DeleteProcessingJobById
    /// Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_processing_job_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::DELETE, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// PUT FinishProcessingJob
    /// Completes processing jobs at a Facility by recording final notes and waste measurements.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn finish_processing_job_processing_job(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/finish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetActiveJobTypes
    /// Retrieves a list of all active processing job types defined within a Facility.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_processing_job_active_job_types(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/jobtypes/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
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
    /// GET GetActiveProcessingJob
    /// Retrieves active processing jobs at a specified Facility.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_active_processing_job(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/active");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
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
    /// GET GetInactiveJobTypes
    /// Retrieves a list of all inactive processing job types defined within a Facility.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_processing_job_inactive_job_types(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/jobtypes/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
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
    /// GET GetInactiveProcessingJob
    /// Retrieves inactive processing jobs at a specified Facility.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_inactive_processing_job(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/inactive");
        let mut query_params = Vec::new();
        if let Some(p) = last_modified_end {
            query_params.push(format!("lastModifiedEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = last_modified_start {
            query_params.push(format!("lastModifiedStart={}", urlencoding::encode(&p)));
        }
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
    /// GET GetJobTypesAttributes
    /// Retrieves all processing job attributes available for a Facility.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_processing_job_job_types_attributes(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/jobtypes/attributes");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetJobTypesCategories
    /// Retrieves all processing job categories available for a specified Facility.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_processing_job_job_types_categories(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/jobtypes/categories");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetProcessingJobById
    /// Retrieves a ProcessingJob by Id.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_processing_job_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/{}", urlencoding::encode(id).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// POST StartProcessingJob
    /// Initiates new processing jobs at a Facility, including job details and associated packages.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn start_processing_job_processing_job(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/start");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// PUT UnfinishProcessingJob
    /// Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn unfinish_processing_job_processing_job(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/unfinish");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// PUT UpdateJobTypes
    /// Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_processing_job_job_types(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/processing/v2/jobtypes");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::PUT, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
}

