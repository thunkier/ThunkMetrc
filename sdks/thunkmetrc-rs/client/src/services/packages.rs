use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct PackagesClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> PackagesClient<'a> {
    /// POST CreateAdjustPackages
    /// Records a list of adjustments for packages at a specific Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Packages Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_adjust_packages(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/adjust");
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
    /// POST CreatePackagesPackages
    /// Creates new packages for a specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Create/Submit/Discontinue Packages
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_packages_packages(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2");
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
    /// POST CreatePackagesPlantings
    /// Creates new plantings from packages for a specified Facility.
    /// Permissions Required:
    /// - View Immature Plants
    /// - Manage Immature Plants
    /// - View Packages
    /// - Manage Packages Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_packages_plantings(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/plantings");
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
    /// POST CreateTesting
    /// Creates new packages for testing for a specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Create/Submit/Discontinue Packages
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_packages_testing(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/testing");
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
    /// DELETE DeletePackagesById
    /// Discontinues a Package at a specific Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Create/Submit/Discontinue Packages
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_packages_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/{}", urlencoding::encode(id).as_ref());
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
    /// PUT FinishPackages
    /// Updates a list of packages as finished for a specific Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Packages Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn finish_packages_packages(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/finish");
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
    /// PUT FinishedgoodFlag
    /// Flags one or more Packages at the specified Facility as Finished Goods.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Packages Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn finishedgood_flag_packages(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/finishedgood/flag");
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
    /// PUT FinishedgoodUnflag
    /// Removes the Finished Good flag one or more Packages at the specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Packages Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn finishedgood_unflag_packages(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/finishedgood/unflag");
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
    /// GET GetActivePackages
    /// Retrieves a list of active packages for a specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_active_packages(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/active");
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
    /// GET GetAdjustReasons
    /// Retrieves a list of adjustment reasons for packages at a specified Facility.
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_packages_adjust_reasons(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/adjust/reasons");
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
    /// GET GetAdjustments
    /// Retrieves the Package Adjustments for a Facility
    /// Permissions Required:
    /// - View Packages
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_packages_adjustments(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/adjustments");
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
    /// GET GetInTransitPackages
    /// Retrieves a list of packages in transit for a specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_in_transit_packages(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/intransit");
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
    /// GET GetInactivePackages
    /// Retrieves a list of inactive packages for a specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_inactive_packages(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/inactive");
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
    /// GET GetLabSamples
    /// Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_packages_lab_samples(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/labsamples");
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
    /// GET GetOnHoldPackages
    /// Retrieves a list of packages on hold for a specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_on_hold_packages(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/onhold");
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
    /// GET GetPackagesById
    /// Retrieves a Package by its Id.
    /// Permissions Required:
    /// - View Packages
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_packages_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/{}", urlencoding::encode(id).as_ref());
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
    /// GET GetPackagesByLabel
    /// Retrieves a Package by its label.
    /// Permissions Required:
    /// - View Packages
    /// Parameters:
    /// - label (str): Path parameter label
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_packages_by_label(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/{}", urlencoding::encode(label).as_ref());
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
    /// GET GetPackagesTypes
    /// Retrieves a list of available Package types.
    /// Parameters:
    pub async fn get_packages_types(&self, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let path = format!("/packages/v2/types");

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetSourceHarvestById
    /// Retrieves the source harvests for a Package by its Id.
    /// Permissions Required:
    /// - View Package Source Harvests
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_packages_source_harvest_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/{}/source/harvests", urlencoding::encode(id).as_ref());
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
    /// GET GetTransferred
    /// Retrieves a list of transferred packages for a specific Facility.
    /// Permissions Required:
    /// - View Packages
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_packages_transferred(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/transferred");
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
    /// PUT UnfinishPackages
    /// Updates a list of packages as unfinished for a specific Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Packages Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn unfinish_packages_packages(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/unfinish");
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
    /// PUT UpdateAdjustPackages
    /// Set the final quantity for a Package.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Packages Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_adjust_packages(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/adjust");
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
    /// PUT UpdateDecontaminate
    /// Updates the Product decontaminate information for a list of packages at a specific Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Packages Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_packages_decontaminate(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/decontaminate");
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
    /// PUT UpdateDonationFlag
    /// Flags one or more packages for donation at the specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Packages Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_packages_donation_flag(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/donation/flag");
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
    /// PUT UpdateDonationUnflag
    /// Removes the donation flag from one or more packages at the specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Packages Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_packages_donation_unflag(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/donation/unflag");
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
    /// PUT UpdateExternalid
    /// Updates the external identifiers for one or more packages at the specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Package Inventory
    /// - External Id Enabled
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_packages_externalid(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/externalid");
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
    /// PUT UpdateItem
    /// Updates the associated Item for one or more packages at the specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Create/Submit/Discontinue Packages
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_packages_item(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/item");
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
    /// PUT UpdateLabtestsRequired
    /// Updates the list of required lab test batches for one or more packages at the specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Create/Submit/Discontinue Packages
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_packages_labtests_required(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/labtests/required");
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
    /// PUT UpdateNote
    /// Updates notes associated with one or more packages for the specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Packages Inventory
    /// - Manage Package Notes
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_packages_note(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/note");
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
    /// PUT UpdatePackagesLocation
    /// Updates the Location and Sublocation for one or more packages at the specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Create/Submit/Discontinue Packages
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_packages_location(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/location");
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
    /// PUT UpdateRemediate
    /// Updates a list of Product remediations for packages at a specific Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Packages Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_packages_remediate(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/remediate");
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
    /// PUT UpdateTradeSampleFlag
    /// Flags or unflags one or more packages at the specified Facility as trade samples.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Packages Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_packages_trade_sample_flag(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/tradesample/flag");
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
    /// PUT UpdateTradeSampleUnflag
    /// Removes the trade sample flag from one or more packages at the specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Manage Packages Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_packages_trade_sample_unflag(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/tradesample/unflag");
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
    /// PUT UpdateUseByDate
    /// Updates the use-by date for one or more packages at the specified Facility.
    /// Permissions Required:
    /// - View Packages
    /// - Create/Submit/Discontinue Packages
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_packages_use_by_date(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/packages/v2/usebydate");
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

