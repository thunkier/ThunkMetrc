use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct HarvestsClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> HarvestsClient<'a> {
    /// POST CreateHarvestsPackages
    /// Creates packages from harvested products for a specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Manage Harvests
    /// - View Packages
    /// - Create/Submit/Discontinue Packages
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_harvests_packages(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/packages");
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
    /// POST CreateHarvestsWaste
    /// Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
    /// Permissions Required:
    /// - View Harvests
    /// - Manage Harvests
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_harvests_waste(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/waste");
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
    /// POST CreatePackagesTesting
    /// Creates packages for testing from harvested products for a specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Manage Harvests
    /// - View Packages
    /// - Create/Submit/Discontinue Packages
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_harvests_packages_testing(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/packages/testing");
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
    /// DELETE DeleteWasteById
    /// Discontinues a specific harvest waste record by Id for the specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Discontinue Harvest Waste
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_harvests_waste_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/waste/{}", urlencoding::encode(id).as_ref());
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
    /// PUT FinishHarvests
    /// Marks one or more harvests as finished for the specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Finish/Discontinue Harvests
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn finish_harvests_harvests(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/finish");
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
    /// GET GetActiveHarvests
    /// Retrieves a list of active harvests for a specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_active_harvests(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/active");
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
    /// GET GetHarvestsById
    /// Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
    /// Permissions Required:
    /// - View Harvests
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_harvests_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/{}", urlencoding::encode(id).as_ref());
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
    /// GET GetHarvestsWaste
    /// Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
    /// Permissions Required:
    /// - View Harvests
    /// Parameters:
    /// - harvest_id (Option<String>): Filter by harvestId
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_harvests_waste(&self, harvest_id: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/waste");
        let mut query_params = Vec::new();
        if let Some(p) = harvest_id {
            query_params.push(format!("harvestId={}", urlencoding::encode(&p)));
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
    /// GET GetInactiveHarvests
    /// Retrieves a list of inactive harvests for a specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_inactive_harvests(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/inactive");
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
    /// GET GetOnHoldHarvests
    /// Retrieves a list of harvests on hold for a specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_on_hold_harvests(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/onhold");
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
    /// GET GetWasteTypes
    /// Retrieves a list of Waste types for harvests.
    /// Parameters:
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_harvests_waste_types(&self, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/waste/types");
        let mut query_params = Vec::new();
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
    /// PUT UnfinishHarvests
    /// Reopens one or more previously finished harvests for the specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Finish/Discontinue Harvests
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn unfinish_harvests_harvests(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/unfinish");
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
    /// PUT UpdateHarvestsLocation
    /// Updates the Location of Harvest for a specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Manage Harvests
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_harvests_location(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/location");
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
    /// PUT UpdateRename
    /// Renames one or more harvests for the specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Manage Harvests
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_harvests_rename(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/rename");
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
    /// PUT UpdateRestoreHarvestedPlants
    /// Restores previously harvested plants to their original state for the specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Finish/Discontinue Harvests
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_harvests_restore_harvested_plants(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/harvests/v2/restore/harvestedplants");
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

