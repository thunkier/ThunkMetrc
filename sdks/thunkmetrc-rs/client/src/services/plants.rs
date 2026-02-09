use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct PlantsClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> PlantsClient<'a> {
    /// POST CreateAdditivesByLocation
    /// Records additive usage for plants based on their location within a specified Facility.
    /// Permissions Required:
    /// - Manage Plants
    /// - Manage Plants Additives
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_plants_additives_by_location(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/additives/bylocation");
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
    /// POST CreateAdditivesByLocationUsingTemplate
    /// Records additive usage for plants by location using a predefined additive template at a specified Facility.
    /// Permissions Required:
    /// - Manage Plants Additives
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_plants_additives_by_location_using_template(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/additives/bylocation/usingtemplate");
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
    /// POST CreateManicure
    /// Creates harvest product records from plant batches at a specified Facility.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// - Manicure/Harvest Veg/Flower Plants
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_plants_manicure(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/manicure");
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
    /// POST CreatePlantBatchPackages
    /// Creates packages from plant batches at a specified Facility.
    /// Permissions Required:
    /// - View Immature Plants
    /// - Manage Immature Plants Inventory
    /// - View Veg/Flower Plants
    /// - Manage Veg/Flower Plants Inventory
    /// - View Packages
    /// - Create/Submit/Discontinue Packages
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_plants_plant_batch_packages(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/plantbatch/packages");
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
    /// POST CreatePlantsAdditives
    /// Records additive usage details applied to specific plants at a Facility.
    /// Permissions Required:
    /// - Manage Plants Additives
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_plants_additives(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/additives");
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
    /// POST CreatePlantsAdditivesUsingTemplate
    /// Records additive usage for plants using predefined additive templates at a specified Facility.
    /// Permissions Required:
    /// - Manage Plants Additives
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_plants_additives_using_template(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/additives/usingtemplate");
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
    /// POST CreatePlantsPlantings
    /// Creates new plant batches at a specified Facility from existing plant data.
    /// Permissions Required:
    /// - View Immature Plants
    /// - Manage Immature Plants Inventory
    /// - View Veg/Flower Plants
    /// - Manage Veg/Flower Plants Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_plants_plantings(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/plantings");
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
    /// POST CreatePlantsWaste
    /// Records waste events for plants at a Facility, including method, reason, and location details.
    /// Permissions Required:
    /// - Manage Plants Waste
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_plants_waste(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/waste");
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
    /// DELETE DeletePlants
    /// Removes plants from a Facility’s inventory while recording the reason for their disposal.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// - Destroy Veg/Flower Plants
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_plants(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2");
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
    /// GET GetAdditives
    /// Retrieves additive records applied to plants at a specified Facility.
    /// Permissions Required:
    /// - View/Manage Plants Additives
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_plants_additives(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/additives");
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
    /// GET GetAdditivesTypes
    /// Retrieves a list of all plant additive types defined within a Facility.
    /// Parameters:
    pub async fn get_plants_additives_types(&self, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let path = format!("/plants/v2/additives/types");

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetFloweringPlants
    /// Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_flowering_plants(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/flowering");
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
    /// GET GetGrowthPhases
    /// Retrieves the list of growth phases supported by a specified Facility.
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_plants_growth_phases(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/growthphases");
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
    /// GET GetInactivePlants
    /// Retrieves inactive plants at a specified Facility.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_inactive_plants(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/inactive");
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
    /// GET GetMotherInactivePlants
    /// Retrieves inactive mother-phase plants at a specified Facility.
    /// Permissions Required:
    /// - View Mother Plants
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_mother_inactive_plants(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/mother/inactive");
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
    /// GET GetMotherOnHoldPlants
    /// Retrieves mother-phase plants currently marked as on hold at a specified Facility.
    /// Permissions Required:
    /// - View Mother Plants
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_mother_on_hold_plants(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/mother/onhold");
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
    /// GET GetMotherPlants
    /// Retrieves mother-phase plants at a specified Facility.
    /// Permissions Required:
    /// - View Mother Plants
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_mother_plants(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/mother");
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
    /// GET GetOnHoldPlants
    /// Retrieves plants that are currently on hold at a specified Facility.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_on_hold_plants(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/onhold");
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
    /// GET GetPlantsById
    /// Retrieves a Plant by Id.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_plants_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/{}", urlencoding::encode(id).as_ref());
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
    /// GET GetPlantsByLabel
    /// Retrieves a Plant by label.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// Parameters:
    /// - label (str): Path parameter label
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_plants_by_label(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/{}", urlencoding::encode(label).as_ref());
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
    /// GET GetPlantsWaste
    /// Retrieves a list of recorded plant waste events for a specific Facility.
    /// Permissions Required:
    /// - View Plants Waste
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_plants_waste(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/waste");
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
    /// GET GetPlantsWasteMethods
    /// Retrieves a list of all available plant waste methods for use within a Facility.
    /// Parameters:
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_plants_waste_methods(&self, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/waste/methods/all");
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
    /// GET GetPlantsWasteReasons
    /// Retriveves available reasons for recording mature plant waste at a specified Facility.
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_plants_waste_reasons(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/waste/reasons");
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
    /// GET GetVegetativePlants
    /// Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_vegetative_plants(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/vegetative");
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
    /// GET GetWasteById
    /// Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
    /// Permissions Required:
    /// - View Plants Waste
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_plants_waste_by_id(&self, id: &str, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/waste/{}/plant", urlencoding::encode(id).as_ref());
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
    /// GET GetWastePackageById
    /// Retrieves a list of package records linked to the specified plantWasteId for a given facility.
    /// Permissions Required:
    /// - View Plants Waste
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_plants_waste_package_by_id(&self, id: &str, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/waste/{}/package", urlencoding::encode(id).as_ref());
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
    /// PUT UpdateAdjustPlants
    /// Adjusts the recorded count of plants at a specified Facility.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// - Manage Veg/Flower Plants Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_adjust_plants(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/adjust");
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
    /// PUT UpdateGrowthPhase
    /// Changes the growth phases of plants within a specified Facility.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// - Manage Veg/Flower Plants Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_plants_growth_phase(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/growthphase");
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
    /// PUT UpdateHarvest
    /// Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// - Manicure/Harvest Veg/Flower Plants
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_plants_harvest(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/harvest");
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
    /// PUT UpdateMerge
    /// Merges multiple plant groups into a single group within a Facility.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// - Manicure/Harvest Veg/Flower Plants
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_plants_merge(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/merge");
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
    /// PUT UpdatePlantsLocation
    /// Moves plant batches to new locations within a specified Facility.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// - Manage Veg/Flower Plants Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_plants_location(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/location");
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
    /// PUT UpdatePlantsStrain
    /// Updates the strain information for plants within a Facility.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// - Manage Veg/Flower Plants Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_plants_strain(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/strain");
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
    /// PUT UpdatePlantsTag
    /// Replaces existing plant tags with new tags for plants within a Facility.
    /// Permissions Required:
    /// - View Veg/Flower Plants
    /// - Manage Veg/Flower Plants Inventory
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_plants_tag(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/tag");
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
    /// PUT UpdateSplit
    /// Splits an existing plant group into multiple groups within a Facility.
    /// Permissions Required:
    /// - View Plant
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_plants_split(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/plants/v2/split");
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

