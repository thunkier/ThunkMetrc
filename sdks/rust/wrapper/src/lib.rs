pub use thunkmetrc_client::MetrcClient;
use serde_json::Value;
use serde::{Deserialize, Serialize};
use std::error::Error;
use std::sync::Arc;
mod ratelimiter;
use ratelimiter::{MetrcRateLimiter, RateLimiterConfig};

pub struct MetrcWrapper {
    client: MetrcClient,
    rate_limiter: Arc<MetrcRateLimiter>,
}

impl MetrcWrapper {
    pub fn new(client: MetrcClient) -> Self {
        MetrcWrapper {
            client,
            rate_limiter: Arc::new(MetrcRateLimiter::new(Some(RateLimiterConfig::default()))),
        }
    }

    pub fn new_with_config(client: MetrcClient, config: RateLimiterConfig) -> Self {
        MetrcWrapper {
            client,
            rate_limiter: Arc::new(MetrcRateLimiter::new(Some(config))),
        }
    }

    /// POST CreateAdditives V1
    /// Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plants_create_additives_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreateAdditivesV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_additives_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdditives V2
    /// Records additive usage details applied to specific plants at a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plants_create_additives_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreateAdditivesV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_additives_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdditivesBylocation V1
    /// Permissions Required:
    ///   - Manage Plants
    ///   - Manage Plants Additives
    ///
    pub async fn plants_create_additives_bylocation_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreateAdditivesBylocationV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_additives_bylocation_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdditivesBylocation V2
    /// Records additive usage for plants based on their location within a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants
    ///   - Manage Plants Additives
    ///
    pub async fn plants_create_additives_bylocation_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreateAdditivesBylocationV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_additives_bylocation_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdditivesBylocationUsingtemplate V2
    /// Records additive usage for plants by location using a predefined additive template at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plants_create_additives_bylocation_usingtemplate_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_additives_bylocation_usingtemplate_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdditivesUsingtemplate V2
    /// Records additive usage for plants using predefined additive templates at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plants_create_additives_usingtemplate_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreateAdditivesUsingtemplateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_additives_usingtemplate_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateChangegrowthphases V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_create_changegrowthphases_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreateChangegrowthphasesV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_changegrowthphases_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateHarvestplants V1
    /// NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manicure/Harvest Veg/Flower Plants
    ///
    pub async fn plants_create_harvestplants_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreateHarvestplantsV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_harvestplants_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateManicure V2
    /// Creates harvest product records from plant batches at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manicure/Harvest Veg/Flower Plants
    ///
    pub async fn plants_create_manicure_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreateManicureV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_manicure_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateManicureplants V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manicure/Harvest Veg/Flower Plants
    ///
    pub async fn plants_create_manicureplants_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreateManicureplantsV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_manicureplants_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateMoveplants V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_create_moveplants_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreateMoveplantsV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_moveplants_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePlantbatchPackage V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn plants_create_plantbatch_package_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreatePlantbatchPackageV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_plantbatch_package_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePlantbatchPackage V2
    /// Creates packages from plant batches at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn plants_create_plantbatch_package_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreatePlantbatchPackageV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_plantbatch_package_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePlantings V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_create_plantings_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreatePlantingsV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_plantings_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePlantings V2
    /// Creates new plant batches at a specified Facility from existing plant data.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_create_plantings_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreatePlantingsV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_plantings_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateWaste V1
    /// Permissions Required:
    ///   - Manage Plants Waste
    ///
    pub async fn plants_create_waste_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreateWasteV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_waste_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateWaste V2
    /// Records waste events for plants at a Facility, including method, reason, and location details.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Waste
    ///
    pub async fn plants_create_waste_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsCreateWasteV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_create_waste_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Destroy Veg/Flower Plants
    ///
    pub async fn plants_delete_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_delete_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE Delete V2
    /// Removes plants from a Facility’s inventory while recording the reason for their disposal.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Destroy Veg/Flower Plants
    ///
    pub async fn plants_delete_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_delete_v2(license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_get_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V2
    /// Retrieves a Plant by Id.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_get_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetAdditives V1
    /// Permissions Required:
    ///   - View/Manage Plants Additives
    ///
    pub async fn plants_get_additives_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_get_additives_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetAdditives V2
    /// Retrieves additive records applied to plants at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View/Manage Plants Additives
    ///
    pub async fn plants_get_additives_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plants_get_additives_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetAdditivesTypes V1
    /// Permissions Required:
    ///   -
    ///
    pub async fn plants_get_additives_types_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.plants_get_additives_types_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetAdditivesTypes V2
    /// Retrieves a list of all plant additive types defined within a Facility.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_additives_types_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.plants_get_additives_types_v2(no, body.as_ref()).await }
        }).await
    }

    /// GET GetByLabel V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_by_label_v1(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let label = label.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let label = label.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_get_by_label_v1(&label, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetByLabel V2
    /// Retrieves a Plant by label.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_by_label_v2(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let label = label.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let label = label.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_get_by_label_v2(&label, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetFlowering V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_flowering_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_get_flowering_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetFlowering V2
    /// Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_flowering_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plants_get_flowering_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetGrowthPhases V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_growth_phases_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_get_growth_phases_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetGrowthPhases V2
    /// Retrieves the list of growth phases supported by a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_growth_phases_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_get_growth_phases_v2(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_get_inactive_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V2
    /// Retrieves inactive plants at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plants_get_inactive_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetMother V2
    /// Retrieves mother-phase plants at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Mother Plants
    ///
    pub async fn plants_get_mother_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plants_get_mother_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetMotherInactive V2
    /// Retrieves inactive mother-phase plants at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Mother Plants
    ///
    pub async fn plants_get_mother_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plants_get_mother_inactive_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetMotherOnhold V2
    /// Retrieves mother-phase plants currently marked as on hold at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Mother Plants
    ///
    pub async fn plants_get_mother_onhold_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plants_get_mother_onhold_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetOnhold V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_onhold_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_get_onhold_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetOnhold V2
    /// Retrieves plants that are currently on hold at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_onhold_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plants_get_onhold_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetVegetative V1
    /// Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_vegetative_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_get_vegetative_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetVegetative V2
    /// Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///
    pub async fn plants_get_vegetative_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plants_get_vegetative_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetWaste V2
    /// Retrieves a list of recorded plant waste events for a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Plants Waste
    ///
    pub async fn plants_get_waste_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plants_get_waste_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetWasteMethodsAll V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_waste_methods_all_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.plants_get_waste_methods_all_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetWasteMethodsAll V2
    /// Retrieves a list of all available plant waste methods for use within a Facility.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_waste_methods_all_v2(&self, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plants_get_waste_methods_all_v2(page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetWastePackage V2
    /// Retrieves a list of package records linked to the specified plantWasteId for a given facility.
    /// 
    ///   Permissions Required:
    ///   - View Plants Waste
    ///
    pub async fn plants_get_waste_package_v2(&self, id: &str, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plants_get_waste_package_v2(&id, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetWastePlant V2
    /// Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
    /// 
    ///   Permissions Required:
    ///   - View Plants Waste
    ///
    pub async fn plants_get_waste_plant_v2(&self, id: &str, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plants_get_waste_plant_v2(&id, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetWasteReasons V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_waste_reasons_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plants_get_waste_reasons_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetWasteReasons V2
    /// Retriveves available reasons for recording mature plant waste at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn plants_get_waste_reasons_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plants_get_waste_reasons_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// PUT UpdateAdjust V2
    /// Adjusts the recorded count of plants at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_update_adjust_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsUpdateAdjustV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_update_adjust_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateGrowthphase V2
    /// Changes the growth phases of plants within a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_update_growthphase_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsUpdateGrowthphaseV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_update_growthphase_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateHarvest V2
    /// Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manicure/Harvest Veg/Flower Plants
    ///
    pub async fn plants_update_harvest_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsUpdateHarvestV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_update_harvest_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateLocation V2
    /// Moves plant batches to new locations within a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_update_location_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsUpdateLocationV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_update_location_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateMerge V2
    /// Merges multiple plant groups into a single group within a Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manicure/Harvest Veg/Flower Plants
    ///
    pub async fn plants_update_merge_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsUpdateMergeV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_update_merge_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateSplit V2
    /// Splits an existing plant group into multiple groups within a Facility.
    /// 
    ///   Permissions Required:
    ///   - View Plant
    ///
    pub async fn plants_update_split_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsUpdateSplitV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_update_split_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateStrain V2
    /// Updates the strain information for plants within a Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_update_strain_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsUpdateStrainV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_update_strain_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateTag V2
    /// Replaces existing plant tags with new tags for plants within a Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plants_update_tag_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantsUpdateTagV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plants_update_tag_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Create V2
    /// Creates new additive templates for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Additives
    ///
    pub async fn additives_templates_create_v2(&self, license_number: Option<String>, body: Option<&Vec<AdditivesTemplatesCreateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.additives_templates_create_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// GET Get V2
    /// Retrieves an Additive Template by its Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Additives
    ///
    pub async fn additives_templates_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.additives_templates_get_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active additive templates for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Additives
    ///
    pub async fn additives_templates_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.additives_templates_get_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive additive templates for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Additives
    ///
    pub async fn additives_templates_get_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.additives_templates_get_inactive_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// PUT Update V2
    /// Updates existing additive templates for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Additives
    ///
    pub async fn additives_templates_update_v2(&self, license_number: Option<String>, body: Option<&Vec<AdditivesTemplatesUpdateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.additives_templates_update_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// GET GetByCaregiverLicenseNumber V1
    /// Data returned by this endpoint is cached for up to one minute.
    /// 
    ///   Permissions Required:
    ///   - Lookup Caregivers
    ///
    pub async fn caregivers_status_get_by_caregiver_license_number_v1(&self, caregiver_license_number: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let caregiver_license_number = caregiver_license_number.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let caregiver_license_number = caregiver_license_number.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.caregivers_status_get_by_caregiver_license_number_v1(&caregiver_license_number, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetByCaregiverLicenseNumber V2
    /// Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
    /// 
    ///   Permissions Required:
    ///   - Lookup Caregivers
    ///
    pub async fn caregivers_status_get_by_caregiver_license_number_v2(&self, caregiver_license_number: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let caregiver_license_number = caregiver_license_number.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let caregiver_license_number = caregiver_license_number.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.caregivers_status_get_by_caregiver_license_number_v2(&caregiver_license_number, license_number, body.as_ref()).await }
        }).await
    }

    /// POST CreateFinish V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Finish/Discontinue Harvests
    ///
    pub async fn harvests_create_finish_v1(&self, license_number: Option<String>, body: Option<&Vec<HarvestsCreateFinishV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_create_finish_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePackage V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn harvests_create_package_v1(&self, license_number: Option<String>, body: Option<&Vec<HarvestsCreatePackageV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_create_package_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePackage V2
    /// Creates packages from harvested products for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn harvests_create_package_v2(&self, license_number: Option<String>, body: Option<&Vec<HarvestsCreatePackageV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_create_package_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePackageTesting V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn harvests_create_package_testing_v1(&self, license_number: Option<String>, body: Option<&Vec<HarvestsCreatePackageTestingV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_create_package_testing_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePackageTesting V2
    /// Creates packages for testing from harvested products for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn harvests_create_package_testing_v2(&self, license_number: Option<String>, body: Option<&Vec<HarvestsCreatePackageTestingV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_create_package_testing_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateRemoveWaste V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///
    pub async fn harvests_create_remove_waste_v1(&self, license_number: Option<String>, body: Option<&Vec<HarvestsCreateRemoveWasteV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_create_remove_waste_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateUnfinish V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Finish/Discontinue Harvests
    ///
    pub async fn harvests_create_unfinish_v1(&self, license_number: Option<String>, body: Option<&Vec<HarvestsCreateUnfinishV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_create_unfinish_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateWaste V2
    /// Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///
    pub async fn harvests_create_waste_v2(&self, license_number: Option<String>, body: Option<&Vec<HarvestsCreateWasteV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_create_waste_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE DeleteWaste V2
    /// Discontinues a specific harvest waste record by Id for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Discontinue Harvest Waste
    ///
    pub async fn harvests_delete_waste_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.harvests_delete_waste_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.harvests_get_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V2
    /// Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.harvests_get_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.harvests_get_active_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active harvests for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.harvests_get_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V1
    /// Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.harvests_get_inactive_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive harvests for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.harvests_get_inactive_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetOnhold V1
    /// Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_onhold_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.harvests_get_onhold_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetOnhold V2
    /// Retrieves a list of harvests on hold for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_onhold_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.harvests_get_onhold_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetWaste V2
    /// Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///
    pub async fn harvests_get_waste_v2(&self, harvest_id: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let harvest_id = harvest_id.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.harvests_get_waste_v2(harvest_id, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetWasteTypes V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn harvests_get_waste_types_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.harvests_get_waste_types_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetWasteTypes V2
    /// Retrieves a list of Waste types for harvests.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn harvests_get_waste_types_v2(&self, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.harvests_get_waste_types_v2(page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// PUT UpdateFinish V2
    /// Marks one or more harvests as finished for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Finish/Discontinue Harvests
    ///
    pub async fn harvests_update_finish_v2(&self, license_number: Option<String>, body: Option<&Vec<HarvestsUpdateFinishV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_update_finish_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateLocation V2
    /// Updates the Location of Harvest for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///
    pub async fn harvests_update_location_v2(&self, license_number: Option<String>, body: Option<&Vec<HarvestsUpdateLocationV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_update_location_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateMove V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///
    pub async fn harvests_update_move_v1(&self, license_number: Option<String>, body: Option<&Vec<HarvestsUpdateMoveV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_update_move_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateRename V1
    /// Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///
    pub async fn harvests_update_rename_v1(&self, license_number: Option<String>, body: Option<&Vec<HarvestsUpdateRenameV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_update_rename_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateRename V2
    /// Renames one or more harvests for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Manage Harvests
    ///
    pub async fn harvests_update_rename_v2(&self, license_number: Option<String>, body: Option<&Vec<HarvestsUpdateRenameV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_update_rename_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateRestoreHarvestedPlants V2
    /// Restores previously harvested plants to their original state for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Finish/Discontinue Harvests
    ///
    pub async fn harvests_update_restore_harvested_plants_v2(&self, license_number: Option<String>, body: Option<&Vec<HarvestsUpdateRestoreHarvestedPlantsV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_update_restore_harvested_plants_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateUnfinish V2
    /// Reopens one or more previously finished harvests for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Harvests
    ///   - Finish/Discontinue Harvests
    ///
    pub async fn harvests_update_unfinish_v2(&self, license_number: Option<String>, body: Option<&Vec<HarvestsUpdateUnfinishV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.harvests_update_unfinish_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateRecord V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_create_record_v1(&self, license_number: Option<String>, body: Option<&Vec<LabTestsCreateRecordV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.lab_tests_create_record_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateRecord V2
    /// Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_create_record_v2(&self, license_number: Option<String>, body: Option<&Vec<LabTestsCreateRecordV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.lab_tests_create_record_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// GET GetBatches V2
    /// Retrieves a list of Lab Test batches.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn lab_tests_get_batches_v2(&self, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.lab_tests_get_batches_v2(page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetLabtestdocument V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_get_labtestdocument_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.lab_tests_get_labtestdocument_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetLabtestdocument V2
    /// Retrieves a specific Lab Test result document by its Id for a given Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_get_labtestdocument_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.lab_tests_get_labtestdocument_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetResults V1
    /// Permissions Required:
    ///   - View Packages
    ///
    pub async fn lab_tests_get_results_v1(&self, license_number: Option<String>, package_id: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let package_id = package_id.clone();
            let body = body.clone();
            async move { client.lab_tests_get_results_v1(license_number, package_id, body.as_ref()).await }
        }).await
    }

    /// GET GetResults V2
    /// Retrieves Lab Test results for a specified Package.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_get_results_v2(&self, license_number: Option<String>, package_id: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let package_id = package_id.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.lab_tests_get_results_v2(license_number, package_id, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetStates V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn lab_tests_get_states_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.lab_tests_get_states_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetStates V2
    /// Returns a list of all lab testing states.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn lab_tests_get_states_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.lab_tests_get_states_v2(no, body.as_ref()).await }
        }).await
    }

    /// GET GetTypes V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn lab_tests_get_types_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.lab_tests_get_types_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetTypes V2
    /// Returns a list of Lab Test types.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn lab_tests_get_types_v2(&self, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.lab_tests_get_types_v2(page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// PUT UpdateLabtestdocument V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_update_labtestdocument_v1(&self, license_number: Option<String>, body: Option<&Vec<LabTestsUpdateLabtestdocumentV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.lab_tests_update_labtestdocument_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateLabtestdocument V2
    /// Updates one or more documents for previously submitted lab tests.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_update_labtestdocument_v2(&self, license_number: Option<String>, body: Option<&Vec<LabTestsUpdateLabtestdocumentV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.lab_tests_update_labtestdocument_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateResultRelease V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_update_result_release_v1(&self, license_number: Option<String>, body: Option<&Vec<LabTestsUpdateResultReleaseV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.lab_tests_update_result_release_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateResultRelease V2
    /// Releases Lab Test results for one or more packages.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn lab_tests_update_result_release_v2(&self, license_number: Option<String>, body: Option<&Vec<LabTestsUpdateResultReleaseV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.lab_tests_update_result_release_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateDelivery V1
    /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_create_delivery_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateDeliveryV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_delivery_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateDelivery V2
    /// Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    ///
    pub async fn sales_create_delivery_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateDeliveryV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_delivery_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateDeliveryRetailer V1
    /// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateDeliveryRetailerV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_delivery_retailer_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateDeliveryRetailer V2
    /// Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateDeliveryRetailerV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_delivery_retailer_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateDeliveryRetailerDepart V1
    /// Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_depart_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateDeliveryRetailerDepartV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_delivery_retailer_depart_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateDeliveryRetailerDepart V2
    /// Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_depart_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateDeliveryRetailerDepartV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_delivery_retailer_depart_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateDeliveryRetailerEnd V1
    /// Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_end_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateDeliveryRetailerEndV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_delivery_retailer_end_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateDeliveryRetailerEnd V2
    /// Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_end_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateDeliveryRetailerEndV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_delivery_retailer_end_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateDeliveryRetailerRestock V1
    /// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_restock_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateDeliveryRetailerRestockV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_delivery_retailer_restock_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateDeliveryRetailerRestock V2
    /// Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_restock_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateDeliveryRetailerRestockV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_delivery_retailer_restock_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateDeliveryRetailerSale V1
    /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_create_delivery_retailer_sale_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateDeliveryRetailerSaleV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_delivery_retailer_sale_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateDeliveryRetailerSale V2
    /// Records sales deliveries originating from a retailer delivery for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    ///
    pub async fn sales_create_delivery_retailer_sale_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateDeliveryRetailerSaleV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_delivery_retailer_sale_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateReceipt V1
    /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_create_receipt_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateReceiptV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_receipt_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateReceipt V2
    /// Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales (Write)
    ///   - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
    ///   - Industry/Facility Type/Advanced Sales
    ///   - WebApi Sales Read Write State (All or WriteOnly)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    ///
    pub async fn sales_create_receipt_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesCreateReceiptV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_receipt_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateTransactionByDate V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_create_transaction_by_date_v1(&self, date: &str, license_number: Option<String>, body: Option<&Vec<SalesCreateTransactionByDateV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let date = date.to_string();
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let date = date.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_create_transaction_by_date_v1(&date, license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE DeleteDelivery V1
    /// Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_delete_delivery_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_delete_delivery_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE DeleteDelivery V2
    /// Voids a sales delivery for a Facility using the provided License Number and delivery Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_delete_delivery_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_delete_delivery_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE DeleteDeliveryRetailer V1
    /// Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_delete_delivery_retailer_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_delete_delivery_retailer_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE DeleteDeliveryRetailer V2
    /// Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_delete_delivery_retailer_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_delete_delivery_retailer_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE DeleteReceipt V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_delete_receipt_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_delete_receipt_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE DeleteReceipt V2
    /// Archives a sales receipt for a Facility using the provided License Number and receipt Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales
    ///
    pub async fn sales_delete_receipt_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_delete_receipt_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetCounties V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn sales_get_counties_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.sales_get_counties_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetCounties V2
    /// Returns a list of counties available for sales deliveries.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn sales_get_counties_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.sales_get_counties_v2(no, body.as_ref()).await }
        }).await
    }

    /// GET GetCustomertypes V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn sales_get_customertypes_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.sales_get_customertypes_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetCustomertypes V2
    /// Returns a list of customer types.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn sales_get_customertypes_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.sales_get_customertypes_v2(no, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveriesActive V1
    /// Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_get_deliveries_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let sales_date_end = sales_date_end.clone();
            let sales_date_start = sales_date_start.clone();
            let body = body.clone();
            async move { client.sales_get_deliveries_active_v1(last_modified_end, last_modified_start, license_number, sales_date_end, sales_date_start, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveriesActive V2
    /// Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
    /// 
    ///   Permissions Required:
    ///   - View Sales Delivery
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_get_deliveries_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let sales_date_end = sales_date_end.clone();
            let sales_date_start = sales_date_start.clone();
            let body = body.clone();
            async move { client.sales_get_deliveries_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, sales_date_end, sales_date_start, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveriesInactive V1
    /// Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_get_deliveries_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let sales_date_end = sales_date_end.clone();
            let sales_date_start = sales_date_start.clone();
            let body = body.clone();
            async move { client.sales_get_deliveries_inactive_v1(last_modified_end, last_modified_start, license_number, sales_date_end, sales_date_start, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveriesInactive V2
    /// Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
    /// 
    ///   Permissions Required:
    ///   - View Sales Delivery
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_get_deliveries_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let sales_date_end = sales_date_end.clone();
            let sales_date_start = sales_date_start.clone();
            let body = body.clone();
            async move { client.sales_get_deliveries_inactive_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, sales_date_end, sales_date_start, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveriesRetailerActive V1
    /// Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_get_deliveries_retailer_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_deliveries_retailer_active_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveriesRetailerActive V2
    /// Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
    /// 
    ///   Permissions Required:
    ///   - View Retailer Delivery
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_get_deliveries_retailer_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.sales_get_deliveries_retailer_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveriesRetailerInactive V1
    /// Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_get_deliveries_retailer_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_deliveries_retailer_inactive_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveriesRetailerInactive V2
    /// Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
    /// 
    ///   Permissions Required:
    ///   - View Retailer Delivery
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_get_deliveries_retailer_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.sales_get_deliveries_retailer_inactive_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveriesReturnreasons V1
    /// Permissions Required:
    ///   -
    ///
    pub async fn sales_get_deliveries_returnreasons_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_deliveries_returnreasons_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveriesReturnreasons V2
    /// Returns a list of return reasons for sales deliveries based on the provided License Number.
    /// 
    ///   Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_get_deliveries_returnreasons_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.sales_get_deliveries_returnreasons_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetDelivery V1
    /// Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_get_delivery_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_delivery_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetDelivery V2
    /// Retrieves a sales delivery record by its Id, with an optional License Number.
    /// 
    ///   Permissions Required:
    ///   - View Sales Delivery
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_get_delivery_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_delivery_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveryRetailer V1
    /// Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_get_delivery_retailer_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_delivery_retailer_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveryRetailer V2
    /// Retrieves a retailer delivery record by its ID, with an optional License Number.
    /// 
    ///   Permissions Required:
    ///   - View Retailer Delivery
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_get_delivery_retailer_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_delivery_retailer_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetPatientRegistrationsLocations V1
    /// Permissions Required:
    ///   -
    ///
    pub async fn sales_get_patient_registrations_locations_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.sales_get_patient_registrations_locations_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetPatientRegistrationsLocations V2
    /// Returns a list of valid Patient registration locations for sales.
    /// 
    ///   Permissions Required:
    ///   -
    ///
    pub async fn sales_get_patient_registrations_locations_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.sales_get_patient_registrations_locations_v2(no, body.as_ref()).await }
        }).await
    }

    /// GET GetPaymenttypes V1
    /// Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_get_paymenttypes_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_paymenttypes_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetPaymenttypes V2
    /// Returns a list of available payment types for the specified License Number.
    /// 
    ///   Permissions Required:
    ///   - View Sales Delivery
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_get_paymenttypes_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_paymenttypes_v2(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetReceipt V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_get_receipt_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_receipt_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetReceipt V2
    /// Retrieves a sales receipt by its Id, with an optional License Number.
    /// 
    ///   Permissions Required:
    ///   - View Sales
    ///   - Manage Sales
    ///
    pub async fn sales_get_receipt_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_receipt_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetReceiptsActive V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_get_receipts_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let sales_date_end = sales_date_end.clone();
            let sales_date_start = sales_date_start.clone();
            let body = body.clone();
            async move { client.sales_get_receipts_active_v1(last_modified_end, last_modified_start, license_number, sales_date_end, sales_date_start, body.as_ref()).await }
        }).await
    }

    /// GET GetReceiptsActive V2
    /// Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
    /// 
    ///   Permissions Required:
    ///   - View Sales
    ///   - Manage Sales
    ///
    pub async fn sales_get_receipts_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let sales_date_end = sales_date_end.clone();
            let sales_date_start = sales_date_start.clone();
            let body = body.clone();
            async move { client.sales_get_receipts_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, sales_date_end, sales_date_start, body.as_ref()).await }
        }).await
    }

    /// GET GetReceiptsExternalByExternalNumber V2
    /// Retrieves a Sales Receipt by its external number, with an optional License Number.
    /// 
    ///   Permissions Required:
    ///   - View Sales
    ///   - Manage Sales
    ///
    pub async fn sales_get_receipts_external_by_external_number_v2(&self, external_number: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let external_number = external_number.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let external_number = external_number.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_receipts_external_by_external_number_v2(&external_number, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetReceiptsInactive V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_get_receipts_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let sales_date_end = sales_date_end.clone();
            let sales_date_start = sales_date_start.clone();
            let body = body.clone();
            async move { client.sales_get_receipts_inactive_v1(last_modified_end, last_modified_start, license_number, sales_date_end, sales_date_start, body.as_ref()).await }
        }).await
    }

    /// GET GetReceiptsInactive V2
    /// Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
    /// 
    ///   Permissions Required:
    ///   - View Sales
    ///   - Manage Sales
    ///
    pub async fn sales_get_receipts_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, sales_date_end: Option<String>, sales_date_start: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let sales_date_end = sales_date_end.clone();
            let sales_date_start = sales_date_start.clone();
            let body = body.clone();
            async move { client.sales_get_receipts_inactive_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, sales_date_end, sales_date_start, body.as_ref()).await }
        }).await
    }

    /// GET GetTransactions V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_get_transactions_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_transactions_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetTransactionsBySalesDateStartAndSalesDateEnd V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_get_transactions_by_sales_date_start_and_sales_date_end_v1(&self, sales_date_start: &str, sales_date_end: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let sales_date_start = sales_date_start.to_string();
        let sales_date_end = sales_date_end.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let sales_date_start = sales_date_start.clone();
            let sales_date_end = sales_date_end.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sales_get_transactions_by_sales_date_start_and_sales_date_end_v1(&sales_date_start, &sales_date_end, license_number, body.as_ref()).await }
        }).await
    }

    /// PUT UpdateDelivery V1
    /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_update_delivery_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDelivery V2
    /// Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_update_delivery_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDeliveryComplete V1
    /// Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_update_delivery_complete_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryCompleteV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_complete_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDeliveryComplete V2
    /// Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery
    ///
    pub async fn sales_update_delivery_complete_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryCompleteV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_complete_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDeliveryHub V1
    /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Sales Delivery
    ///
    pub async fn sales_update_delivery_hub_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryHubV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_hub_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDeliveryHub V2
    /// Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery, Manage Sales Delivery Hub
    ///
    pub async fn sales_update_delivery_hub_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryHubV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_hub_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDeliveryHubAccept V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_update_delivery_hub_accept_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryHubAcceptV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_hub_accept_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDeliveryHubAccept V2
    /// Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery Hub
    ///
    pub async fn sales_update_delivery_hub_accept_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryHubAcceptV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_hub_accept_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDeliveryHubDepart V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_update_delivery_hub_depart_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryHubDepartV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_hub_depart_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDeliveryHubDepart V2
    /// Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery Hub
    ///
    pub async fn sales_update_delivery_hub_depart_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryHubDepartV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_hub_depart_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDeliveryHubVerifyID V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_update_delivery_hub_verify_id_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryHubVerifyIdV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_hub_verify_id_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDeliveryHubVerifyID V2
    /// Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales Delivery Hub
    ///
    pub async fn sales_update_delivery_hub_verify_id_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryHubVerifyIdV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_hub_verify_id_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDeliveryRetailer V1
    /// Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Retailer Delivery
    ///
    pub async fn sales_update_delivery_retailer_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryRetailerV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_retailer_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDeliveryRetailer V2
    /// Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    ///   - Industry/Facility Type/Retailer Delivery
    ///   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    ///   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    ///   - Manage Retailer Delivery
    ///
    pub async fn sales_update_delivery_retailer_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateDeliveryRetailerV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_delivery_retailer_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateReceipt V1
    /// Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_update_receipt_v1(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateReceiptV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_receipt_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateReceipt V2
    /// Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales
    ///
    pub async fn sales_update_receipt_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateReceiptV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_receipt_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateReceiptFinalize V2
    /// Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales
    ///
    pub async fn sales_update_receipt_finalize_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateReceiptFinalizeV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_receipt_finalize_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateReceiptUnfinalize V2
    /// Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
    /// 
    ///   Permissions Required:
    ///   - Manage Sales
    ///
    pub async fn sales_update_receipt_unfinalize_v2(&self, license_number: Option<String>, body: Option<&Vec<SalesUpdateReceiptUnfinalizeV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_receipt_unfinalize_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateTransactionByDate V1
    /// Permissions Required:
    ///   - Sales
    ///
    pub async fn sales_update_transaction_by_date_v1(&self, date: &str, license_number: Option<String>, body: Option<&Vec<SalesUpdateTransactionByDateV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let date = date.to_string();
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let date = date.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sales_update_transaction_by_date_v1(&date, license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateDriver V2
    /// Creates new driver records for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transporters
    ///
    pub async fn transporters_create_driver_v2(&self, license_number: Option<String>, body: Option<&Vec<TransportersCreateDriverV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.transporters_create_driver_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateVehicle V2
    /// Creates new vehicle records for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transporters
    ///
    pub async fn transporters_create_vehicle_v2(&self, license_number: Option<String>, body: Option<&Vec<TransportersCreateVehicleV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.transporters_create_vehicle_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE DeleteDriver V2
    /// Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Transporters
    ///
    pub async fn transporters_delete_driver_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.transporters_delete_driver_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE DeleteVehicle V2
    /// Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Transporters
    ///
    pub async fn transporters_delete_vehicle_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.transporters_delete_vehicle_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetDriver V2
    /// Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
    /// 
    ///   Permissions Required:
    ///   - Transporters
    ///
    pub async fn transporters_get_driver_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.transporters_get_driver_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetDrivers V2
    /// Retrieves a list of drivers for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Transporters
    ///
    pub async fn transporters_get_drivers_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transporters_get_drivers_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetVehicle V2
    /// Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
    /// 
    ///   Permissions Required:
    ///   - Transporters
    ///
    pub async fn transporters_get_vehicle_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.transporters_get_vehicle_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetVehicles V2
    /// Retrieves a list of vehicles for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Transporters
    ///
    pub async fn transporters_get_vehicles_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transporters_get_vehicles_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// PUT UpdateDriver V2
    /// Updates existing driver records for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transporters
    ///
    pub async fn transporters_update_driver_v2(&self, license_number: Option<String>, body: Option<&Vec<TransportersUpdateDriverV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.transporters_update_driver_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateVehicle V2
    /// Updates existing vehicle records for a facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transporters
    ///
    pub async fn transporters_update_vehicle_v2(&self, license_number: Option<String>, body: Option<&Vec<TransportersUpdateVehicleV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.transporters_update_vehicle_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// GET GetAll V1
    /// This endpoint provides a list of facilities for which the authenticated user has access.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn facilities_get_all_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.facilities_get_all_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetAll V2
    /// This endpoint provides a list of facilities for which the authenticated user has access.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn facilities_get_all_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.facilities_get_all_v2(no, body.as_ref()).await }
        }).await
    }

    /// POST Create V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_create_v1(&self, license_number: Option<String>, body: Option<&Vec<PackagesCreateV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_create_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Create V2
    /// Creates new packages for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_create_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesCreateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_create_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdjust V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_adjust_v1(&self, license_number: Option<String>, body: Option<&Vec<PackagesCreateAdjustV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_create_adjust_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdjust V2
    /// Records a list of adjustments for packages at a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_adjust_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesCreateAdjustV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_create_adjust_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateChangeItem V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_create_change_item_v1(&self, license_number: Option<String>, body: Option<&Vec<PackagesCreateChangeItemV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_create_change_item_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateChangeLocation V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_create_change_location_v1(&self, license_number: Option<String>, body: Option<&Vec<PackagesCreateChangeLocationV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_create_change_location_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateFinish V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_finish_v1(&self, license_number: Option<String>, body: Option<&Vec<PackagesCreateFinishV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_create_finish_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePlantings V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_plantings_v1(&self, license_number: Option<String>, body: Option<&Vec<PackagesCreatePlantingsV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_create_plantings_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePlantings V2
    /// Creates new plantings from packages for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_plantings_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesCreatePlantingsV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_create_plantings_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateRemediate V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_remediate_v1(&self, license_number: Option<String>, body: Option<&Vec<PackagesCreateRemediateV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_create_remediate_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateTesting V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_create_testing_v1(&self, license_number: Option<String>, body: Option<&Vec<PackagesCreateTestingV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_create_testing_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateTesting V2
    /// Creates new packages for testing for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_create_testing_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesCreateTestingV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_create_testing_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateUnfinish V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_create_unfinish_v1(&self, license_number: Option<String>, body: Option<&Vec<PackagesCreateUnfinishV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_create_unfinish_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE Delete V2
    /// Discontinues a Package at a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.packages_delete_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.packages_get_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V2
    /// Retrieves a Package by its Id.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.packages_get_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.packages_get_active_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active packages for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.packages_get_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetAdjustReasons V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn packages_get_adjust_reasons_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.packages_get_adjust_reasons_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetAdjustReasons V2
    /// Retrieves a list of adjustment reasons for packages at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn packages_get_adjust_reasons_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.packages_get_adjust_reasons_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetByLabel V1
    /// Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_by_label_v1(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let label = label.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let label = label.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.packages_get_by_label_v1(&label, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetByLabel V2
    /// Retrieves a Package by its label.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_by_label_v2(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let label = label.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let label = label.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.packages_get_by_label_v2(&label, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V1
    /// Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.packages_get_inactive_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive packages for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.packages_get_inactive_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetIntransit V2
    /// Retrieves a list of packages in transit for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_intransit_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.packages_get_intransit_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetLabsamples V2
    /// Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_labsamples_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.packages_get_labsamples_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetOnhold V1
    /// Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_onhold_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.packages_get_onhold_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetOnhold V2
    /// Retrieves a list of packages on hold for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_onhold_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.packages_get_onhold_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetSourceHarvest V2
    /// Retrieves the source harvests for a Package by its Id.
    /// 
    ///   Permissions Required:
    ///   - View Package Source Harvests
    ///
    pub async fn packages_get_source_harvest_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.packages_get_source_harvest_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetTransferred V2
    /// Retrieves a list of transferred packages for a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///
    pub async fn packages_get_transferred_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.packages_get_transferred_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetTypes V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn packages_get_types_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.packages_get_types_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetTypes V2
    /// Retrieves a list of available Package types.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn packages_get_types_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.packages_get_types_v2(no, body.as_ref()).await }
        }).await
    }

    /// PUT UpdateAdjust V2
    /// Set the final quantity for a Package.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_adjust_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateAdjustV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_adjust_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateChangeNote V1
    /// Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///   - Manage Package Notes
    ///
    pub async fn packages_update_change_note_v1(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateChangeNoteV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_change_note_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDecontaminate V2
    /// Updates the Product decontaminate information for a list of packages at a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_decontaminate_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateDecontaminateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_decontaminate_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDonationFlag V2
    /// Flags one or more packages for donation at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_donation_flag_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateDonationFlagV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_donation_flag_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateDonationUnflag V2
    /// Removes the donation flag from one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_donation_unflag_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateDonationUnflagV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_donation_unflag_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateExternalid V2
    /// Updates the external identifiers for one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Package Inventory
    ///   - External Id Enabled
    ///
    pub async fn packages_update_externalid_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateExternalidV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_externalid_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateFinish V2
    /// Updates a list of packages as finished for a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_finish_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateFinishV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_finish_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateItem V2
    /// Updates the associated Item for one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_update_item_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateItemV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_item_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateLabTestRequired V2
    /// Updates the list of required lab test batches for one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_update_lab_test_required_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateLabTestRequiredV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_lab_test_required_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateLocation V2
    /// Updates the Location and Sublocation for one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_update_location_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateLocationV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_location_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateNote V2
    /// Updates notes associated with one or more packages for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///   - Manage Package Notes
    ///
    pub async fn packages_update_note_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateNoteV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_note_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateRemediate V2
    /// Updates a list of Product remediations for packages at a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_remediate_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateRemediateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_remediate_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateTradesampleFlag V2
    /// Flags or unflags one or more packages at the specified Facility as trade samples.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_tradesample_flag_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateTradesampleFlagV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_tradesample_flag_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateTradesampleUnflag V2
    /// Removes the trade sample flag from one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_tradesample_unflag_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateTradesampleUnflagV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_tradesample_unflag_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateUnfinish V2
    /// Updates a list of packages as unfinished for a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Manage Packages Inventory
    ///
    pub async fn packages_update_unfinish_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateUnfinishV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_unfinish_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateUsebydate V2
    /// Updates the use-by date for one or more packages at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn packages_update_usebydate_v2(&self, license_number: Option<String>, body: Option<&Vec<PackagesUpdateUsebydateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.packages_update_usebydate_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Create V1
    /// Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_create_v1(&self, license_number: Option<String>, body: Option<&Vec<PatientCheckInsCreateV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.patient_check_ins_create_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Create V2
    /// Records patient check-ins for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_create_v2(&self, license_number: Option<String>, body: Option<&Vec<PatientCheckInsCreateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.patient_check_ins_create_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_delete_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.patient_check_ins_delete_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE Delete V2
    /// Archives a Patient Check-In, identified by its Id, for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.patient_check_ins_delete_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetAll V1
    /// Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_get_all_v1(&self, checkin_date_end: Option<String>, checkin_date_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let checkin_date_end = checkin_date_end.clone();
            let checkin_date_start = checkin_date_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.patient_check_ins_get_all_v1(checkin_date_end, checkin_date_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetAll V2
    /// Retrieves a list of patient check-ins for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_get_all_v2(&self, checkin_date_end: Option<String>, checkin_date_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let checkin_date_end = checkin_date_end.clone();
            let checkin_date_start = checkin_date_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.patient_check_ins_get_all_v2(checkin_date_end, checkin_date_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetLocations V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn patient_check_ins_get_locations_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.patient_check_ins_get_locations_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetLocations V2
    /// Retrieves a list of Patient Check-In locations.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn patient_check_ins_get_locations_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.patient_check_ins_get_locations_v2(no, body.as_ref()).await }
        }).await
    }

    /// PUT Update V1
    /// Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_update_v1(&self, license_number: Option<String>, body: Option<&Vec<PatientCheckInsUpdateV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.patient_check_ins_update_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT Update V2
    /// Updates patient check-ins for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - ManagePatientsCheckIns
    ///
    pub async fn patient_check_ins_update_v2(&self, license_number: Option<String>, body: Option<&Vec<PatientCheckInsUpdateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.patient_check_ins_update_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn units_of_measure_get_active_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.units_of_measure_get_active_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V2
    /// Retrieves all active units of measure.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn units_of_measure_get_active_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.units_of_measure_get_active_v2(no, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V2
    /// Retrieves all inactive units of measure.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn units_of_measure_get_inactive_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.units_of_measure_get_inactive_v2(no, body.as_ref()).await }
        }).await
    }

    /// GET GetAll V2
    /// Retrieves all available waste methods.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn waste_methods_get_all_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.waste_methods_get_all_v2(no, body.as_ref()).await }
        }).await
    }

    /// GET GetAll V1
    /// Permissions Required:
    ///   - Manage Employees
    ///
    pub async fn employees_get_all_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.employees_get_all_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetAll V2
    /// Retrieves a list of employees for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Employees
    ///   - View Employees
    ///
    pub async fn employees_get_all_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.employees_get_all_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetPermissions V2
    /// Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Employees
    ///
    pub async fn employees_get_permissions_v2(&self, employee_license_number: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let employee_license_number = employee_license_number.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.employees_get_permissions_v2(employee_license_number, license_number, body.as_ref()).await }
        }).await
    }

    /// POST Create V1
    /// NOTE: To include a photo with an item, first use POST /items/v1/photo to POST the photo, and then use the returned ID in the request body in this endpoint.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_v1(&self, license_number: Option<String>, body: Option<&Vec<ItemsCreateV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.items_create_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Create V2
    /// Creates one or more new products for the specified Facility. NOTE: To include a photo with an item, first use POST /items/v2/photo to POST the photo, and then use the returned Id in the request body in this endpoint.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_v2(&self, license_number: Option<String>, body: Option<&Vec<ItemsCreateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.items_create_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateBrand V2
    /// Creates one or more new item brands for the specified Facility identified by the License Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_brand_v2(&self, license_number: Option<String>, body: Option<&Vec<ItemsCreateBrandV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.items_create_brand_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateFile V2
    /// Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_file_v2(&self, license_number: Option<String>, body: Option<&Vec<ItemsCreateFileV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.items_create_file_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePhoto V1
    /// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_photo_v1(&self, license_number: Option<String>, body: Option<&Vec<ItemsCreatePhotoV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.items_create_photo_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePhoto V2
    /// This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_photo_v2(&self, license_number: Option<String>, body: Option<&Vec<ItemsCreatePhotoV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.items_create_photo_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateUpdate V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_create_update_v1(&self, license_number: Option<String>, body: Option<&Vec<ItemsCreateUpdateV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.items_create_update_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_delete_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.items_delete_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE Delete V2
    /// Archives the specified Product by Id for the given Facility License Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.items_delete_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE DeleteBrand V2
    /// Archives the specified Item Brand by Id for the given Facility License Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_delete_brand_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.items_delete_brand_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.items_get_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V2
    /// Retrieves detailed information about a specific Item by Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.items_get_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_active_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.items_get_active_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V2
    /// Returns a list of active items for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.items_get_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetBrands V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_brands_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.items_get_brands_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetBrands V2
    /// Retrieves a list of active item brands for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_brands_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.items_get_brands_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetCategories V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn items_get_categories_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.items_get_categories_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetCategories V2
    /// Retrieves a list of item categories.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn items_get_categories_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.items_get_categories_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetFile V2
    /// Retrieves a file by its Id for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_file_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.items_get_file_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_inactive_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.items_get_inactive_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive items for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_inactive_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.items_get_inactive_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetPhoto V1
    /// Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_photo_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.items_get_photo_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetPhoto V2
    /// Retrieves an image by its Id for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_get_photo_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.items_get_photo_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// PUT Update V2
    /// Updates one or more existing products for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_update_v2(&self, license_number: Option<String>, body: Option<&Vec<ItemsUpdateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.items_update_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateBrand V2
    /// Updates one or more existing item brands for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Items
    ///
    pub async fn items_update_brand_v2(&self, license_number: Option<String>, body: Option<&Vec<ItemsUpdateBrandV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.items_update_brand_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Create V1
    /// Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_create_v1(&self, license_number: Option<String>, body: Option<&Vec<LocationsCreateV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.locations_create_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Create V2
    /// Creates new locations for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_create_v2(&self, license_number: Option<String>, body: Option<&Vec<LocationsCreateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.locations_create_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateUpdate V1
    /// Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_create_update_v1(&self, license_number: Option<String>, body: Option<&Vec<LocationsCreateUpdateV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.locations_create_update_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_delete_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.locations_delete_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE Delete V2
    /// Archives a specified Location, identified by its Id, for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.locations_delete_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.locations_get_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V2
    /// Retrieves a Location by its Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.locations_get_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_active_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.locations_get_active_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active locations for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.locations_get_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive locations for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_inactive_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.locations_get_inactive_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetTypes V1
    /// Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_types_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.locations_get_types_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetTypes V2
    /// Retrieves a list of active location types for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_get_types_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.locations_get_types_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// PUT Update V2
    /// Updates existing locations for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn locations_update_v2(&self, license_number: Option<String>, body: Option<&Vec<LocationsUpdateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.locations_update_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Create V2
    /// Adds new patients to a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_create_v2(&self, license_number: Option<String>, body: Option<&Vec<PatientsCreateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.patients_create_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdd V1
    /// Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_create_add_v1(&self, license_number: Option<String>, body: Option<&Vec<PatientsCreateAddV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.patients_create_add_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateUpdate V1
    /// Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_create_update_v1(&self, license_number: Option<String>, body: Option<&Vec<PatientsCreateUpdateV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.patients_create_update_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_delete_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.patients_delete_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE Delete V2
    /// Removes a Patient, identified by an Id, from a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.patients_delete_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.patients_get_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V2
    /// Retrieves a Patient by Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.patients_get_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_get_active_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.patients_get_active_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active patients for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_get_active_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.patients_get_active_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// PUT Update V2
    /// Updates Patient information for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Patients
    ///
    pub async fn patients_update_v2(&self, license_number: Option<String>, body: Option<&Vec<PatientsUpdateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.patients_update_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAssociate V2
    /// Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly)
    ///   - Industry/View Packages
    ///   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
    ///
    pub async fn retail_id_create_associate_v2(&self, license_number: Option<String>, body: Option<&Vec<RetailIdCreateAssociateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.retail_id_create_associate_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateGenerate V2
    /// Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly)
    ///   - Industry/View Packages
    ///   - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
    ///
    pub async fn retail_id_create_generate_v2(&self, license_number: Option<String>, body: Option<&RetailIdCreateGenerateV2Request>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.retail_id_create_generate_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateMerge V2
    /// Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly)
    ///   - Key Value Settings/Retail ID Merge Packages Enabled
    ///
    pub async fn retail_id_create_merge_v2(&self, license_number: Option<String>, body: Option<&RetailIdCreateMergeV2Request>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.retail_id_create_merge_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePackageInfo V2
    /// Retrieves Package information for given list of Package labels.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
    ///   - WebApi Retail ID Read Write State (All or WriteOnly)
    ///   - Industry/View Packages
    ///   - Admin/Employees/Packages Page/Product Labels(Manage)
    ///
    pub async fn retail_id_create_package_info_v2(&self, license_number: Option<String>, body: Option<&RetailIdCreatePackageInfoV2Request>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.retail_id_create_package_info_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// GET GetReceiveByLabel V2
    /// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Manage RetailId
    ///   - WebApi Retail ID Read Write State (All or ReadOnly)
    ///   - Industry/View Packages
    ///   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
    ///
    pub async fn retail_id_get_receive_by_label_v2(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let label = label.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let label = label.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.retail_id_get_receive_by_label_v2(&label, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetReceiveQrByShortCode V2
    /// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
    /// 
    ///   Permissions Required:
    ///   - External Sources(ThirdPartyVendorV2)/Manage RetailId
    ///   - WebApi Retail ID Read Write State (All or ReadOnly)
    ///   - Industry/View Packages
    ///   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
    ///
    pub async fn retail_id_get_receive_qr_by_short_code_v2(&self, short_code: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let short_code = short_code.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let short_code = short_code.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.retail_id_get_receive_qr_by_short_code_v2(&short_code, license_number, body.as_ref()).await }
        }).await
    }

    /// POST CreateIntegratorSetup V2
    /// This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn sandbox_create_integrator_setup_v2(&self, user_key: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let user_key = user_key.clone();
            let body_val = body_val.clone();
            async move { client.sandbox_create_integrator_setup_v2(user_key, body_val.as_ref()).await }
        }).await
    }

    /// POST Create V1
    /// Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_create_v1(&self, license_number: Option<String>, body: Option<&Vec<StrainsCreateV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.strains_create_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Create V2
    /// Creates new strain records for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_create_v2(&self, license_number: Option<String>, body: Option<&Vec<StrainsCreateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.strains_create_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateUpdate V1
    /// Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_create_update_v1(&self, license_number: Option<String>, body: Option<&Vec<StrainsCreateUpdateV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.strains_create_update_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_delete_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.strains_delete_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE Delete V2
    /// Archives an existing strain record for a Facility
    /// 
    ///   Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.strains_delete_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.strains_get_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V2
    /// Retrieves a Strain record by its Id, with an optional license number.
    /// 
    ///   Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.strains_get_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_get_active_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.strains_get_active_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.strains_get_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_get_inactive_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.strains_get_inactive_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// PUT Update V2
    /// Updates existing strain records for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Strains
    ///
    pub async fn strains_update_v2(&self, license_number: Option<String>, body: Option<&Vec<StrainsUpdateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.strains_update_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// GET GetPackageAvailable V2
    /// Returns a list of available package tags. NOTE: This is a premium endpoint.
    /// 
    ///   Permissions Required:
    ///   - Manage Tags
    ///
    pub async fn tags_get_package_available_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.tags_get_package_available_v2(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetPlantAvailable V2
    /// Returns a list of available plant tags. NOTE: This is a premium endpoint.
    /// 
    ///   Permissions Required:
    ///   - Manage Tags
    ///
    pub async fn tags_get_plant_available_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.tags_get_plant_available_v2(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetStaged V2
    /// Returns a list of staged tags. NOTE: This is a premium endpoint.
    /// 
    ///   Permissions Required:
    ///   - Manage Tags
    ///   - RetailId.AllowPackageStaging Key Value enabled
    ///
    pub async fn tags_get_staged_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.tags_get_staged_v2(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetStatusesByPatientLicenseNumber V1
    /// Data returned by this endpoint is cached for up to one minute.
    /// 
    ///   Permissions Required:
    ///   - Lookup Patients
    ///
    pub async fn patients_status_get_statuses_by_patient_license_number_v1(&self, patient_license_number: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let patient_license_number = patient_license_number.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let patient_license_number = patient_license_number.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.patients_status_get_statuses_by_patient_license_number_v1(&patient_license_number, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetStatusesByPatientLicenseNumber V2
    /// Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
    /// 
    ///   Permissions Required:
    ///   - Lookup Patients
    ///
    pub async fn patients_status_get_statuses_by_patient_license_number_v2(&self, patient_license_number: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let patient_license_number = patient_license_number.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let patient_license_number = patient_license_number.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.patients_status_get_statuses_by_patient_license_number_v2(&patient_license_number, license_number, body.as_ref()).await }
        }).await
    }

    /// POST CreateAdditives V1
    /// Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plant_batches_create_additives_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreateAdditivesV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_additives_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdditives V2
    /// Records Additive usage details for plant batches at a specific Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plant_batches_create_additives_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreateAdditivesV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_additives_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdditivesUsingtemplate V2
    /// Records Additive usage for plant batches at a Facility using predefined additive templates.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Additives
    ///
    pub async fn plant_batches_create_additives_usingtemplate_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreateAdditivesUsingtemplateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_additives_usingtemplate_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdjust V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///
    pub async fn plant_batches_create_adjust_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreateAdjustV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_adjust_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdjust V2
    /// Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///
    pub async fn plant_batches_create_adjust_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreateAdjustV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_adjust_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateChangegrowthphase V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plant_batches_create_changegrowthphase_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreateChangegrowthphaseV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_changegrowthphase_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateGrowthphase V2
    /// Updates the growth phase of plants at a specified Facility based on tracking information.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plant_batches_create_growthphase_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreateGrowthphaseV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_growthphase_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePackage V2
    /// Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn plant_batches_create_package_v2(&self, is_from_mother_plant: Option<String>, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreatePackageV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let is_from_mother_plant = is_from_mother_plant.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_package_v2(is_from_mother_plant, license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePackageFrommotherplant V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn plant_batches_create_package_frommotherplant_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreatePackageFrommotherplantV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_package_frommotherplant_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePackageFrommotherplant V2
    /// Creates packages from mother plants at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn plant_batches_create_package_frommotherplant_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreatePackageFrommotherplantV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_package_frommotherplant_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreatePlantings V2
    /// Creates new plantings for a Facility by generating plant batches based on provided planting details.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///
    pub async fn plant_batches_create_plantings_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreatePlantingsV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_plantings_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateSplit V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///
    pub async fn plant_batches_create_split_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreateSplitV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_split_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateSplit V2
    /// Splits an existing Plant Batch into multiple groups at the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///
    pub async fn plant_batches_create_split_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreateSplitV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_split_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateWaste V1
    /// Permissions Required:
    ///   - Manage Plants Waste
    ///
    pub async fn plant_batches_create_waste_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreateWasteV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_waste_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateWaste V2
    /// Records waste information for plant batches based on the submitted data for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Plants Waste
    ///
    pub async fn plant_batches_create_waste_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreateWasteV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_create_waste_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Createpackages V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///   - View Packages
    ///   - Create/Submit/Discontinue Packages
    ///
    pub async fn plant_batches_createpackages_v1(&self, is_from_mother_plant: Option<String>, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreatepackagesV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let is_from_mother_plant = is_from_mother_plant.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_createpackages_v1(is_from_mother_plant, license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Createplantings V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants Inventory
    ///
    pub async fn plant_batches_createplantings_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesCreateplantingsV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_createplantings_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///   - Destroy Immature Plants
    ///
    pub async fn plant_batches_delete_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plant_batches_delete_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE Delete V2
    /// Completes the destruction of plant batches based on the provided input data.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Destroy Immature Plants
    ///
    pub async fn plant_batches_delete_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plant_batches_delete_v2(license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plant_batches_get_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V2
    /// Retrieves a Plant Batch by Id.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plant_batches_get_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_get_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plant_batches_get_active_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plant_batches_get_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_get_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plant_batches_get_inactive_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_get_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plant_batches_get_inactive_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetTypes V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn plant_batches_get_types_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.plant_batches_get_types_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetTypes V2
    /// Retrieves a list of plant batch types.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn plant_batches_get_types_v2(&self, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plant_batches_get_types_v2(page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetWaste V2
    /// Retrieves waste details associated with plant batches at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Plants Waste
    ///
    pub async fn plant_batches_get_waste_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.plant_batches_get_waste_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetWasteReasons V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn plant_batches_get_waste_reasons_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plant_batches_get_waste_reasons_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetWasteReasons V2
    /// Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn plant_batches_get_waste_reasons_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.plant_batches_get_waste_reasons_v2(license_number, body.as_ref()).await }
        }).await
    }

    /// PUT UpdateLocation V2
    /// Moves one or more plant batches to new locations with in a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Immature Plants
    ///   - Manage Immature Plants
    ///
    pub async fn plant_batches_update_location_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesUpdateLocationV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_update_location_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateMoveplantbatches V1
    /// Permissions Required:
    ///   - View Immature Plants
    ///
    pub async fn plant_batches_update_moveplantbatches_v1(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesUpdateMoveplantbatchesV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_update_moveplantbatches_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateName V2
    /// Renames plant batches at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plant_batches_update_name_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesUpdateNameV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_update_name_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateStrain V2
    /// Changes the strain of plant batches at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plant_batches_update_strain_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesUpdateStrainV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_update_strain_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateTag V2
    /// Replaces tags for plant batches at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - View Veg/Flower Plants
    ///   - Manage Veg/Flower Plants Inventory
    ///
    pub async fn plant_batches_update_tag_v2(&self, license_number: Option<String>, body: Option<&Vec<PlantBatchesUpdateTagV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.plant_batches_update_tag_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdjust V1
    /// Permissions Required:
    ///   - ManageProcessingJobs
    ///
    pub async fn processing_jobs_create_adjust_v1(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsCreateAdjustV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_create_adjust_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateAdjust V2
    /// Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_create_adjust_v2(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsCreateAdjustV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_create_adjust_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateJobtypes V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_create_jobtypes_v1(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsCreateJobtypesV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_create_jobtypes_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateJobtypes V2
    /// Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_create_jobtypes_v2(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsCreateJobtypesV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_create_jobtypes_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateStart V1
    /// Permissions Required:
    ///   - ManageProcessingJobs
    ///
    pub async fn processing_jobs_create_start_v1(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsCreateStartV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_create_start_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateStart V2
    /// Initiates new processing jobs at a Facility, including job details and associated packages.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_create_start_v2(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsCreateStartV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_create_start_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Createpackages V1
    /// Permissions Required:
    ///   - ManageProcessingJobs
    ///
    pub async fn processing_jobs_createpackages_v1(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsCreatepackagesV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_createpackages_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Createpackages V2
    /// Creates packages from processing jobs at a Facility, including optional location and note assignments.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_createpackages_v2(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsCreatepackagesV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_createpackages_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE Delete V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_delete_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_delete_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE Delete V2
    /// Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_delete_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE DeleteJobtypes V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_delete_jobtypes_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_delete_jobtypes_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE DeleteJobtypes V2
    /// Archives a Processing Job Type at a Facility, making it inactive for future use.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_delete_jobtypes_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_delete_jobtypes_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V2
    /// Retrieves a ProcessingJob by Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_active_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V2
    /// Retrieves active processing jobs at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_inactive_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V2
    /// Retrieves inactive processing jobs at a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_inactive_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetJobtypesActive V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_active_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_jobtypes_active_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetJobtypesActive V2
    /// Retrieves a list of all active processing job types defined within a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_jobtypes_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetJobtypesAttributes V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_attributes_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_jobtypes_attributes_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetJobtypesAttributes V2
    /// Retrieves all processing job attributes available for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_attributes_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_jobtypes_attributes_v2(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetJobtypesCategories V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_categories_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_jobtypes_categories_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetJobtypesCategories V2
    /// Retrieves all processing job categories available for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_categories_v2(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_jobtypes_categories_v2(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetJobtypesInactive V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_inactive_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_jobtypes_inactive_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetJobtypesInactive V2
    /// Retrieves a list of all inactive processing job types defined within a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_get_jobtypes_inactive_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.processing_jobs_get_jobtypes_inactive_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// PUT UpdateFinish V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_update_finish_v1(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsUpdateFinishV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_update_finish_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateFinish V2
    /// Completes processing jobs at a Facility by recording final notes and waste measurements.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_update_finish_v2(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsUpdateFinishV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_update_finish_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateJobtypes V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_update_jobtypes_v1(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsUpdateJobtypesV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_update_jobtypes_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateJobtypes V2
    /// Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_update_jobtypes_v2(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsUpdateJobtypesV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_update_jobtypes_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateUnfinish V1
    /// Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_update_unfinish_v1(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsUpdateUnfinishV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_update_unfinish_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateUnfinish V2
    /// Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
    /// 
    ///   Permissions Required:
    ///   - Manage Processing Job
    ///
    pub async fn processing_jobs_update_unfinish_v2(&self, license_number: Option<String>, body: Option<&Vec<ProcessingJobsUpdateUnfinishV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.processing_jobs_update_unfinish_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST Create V2
    /// Creates new sublocation records for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn sublocations_create_v2(&self, license_number: Option<String>, body: Option<&Vec<SublocationsCreateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sublocations_create_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE Delete V2
    /// Archives an existing Sublocation record for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn sublocations_delete_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sublocations_delete_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET Get V2
    /// Retrieves a Sublocation by its Id, with an optional license number.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn sublocations_get_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.sublocations_get_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetActive V2
    /// Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn sublocations_get_active_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.sublocations_get_active_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetInactive V2
    /// Retrieves a list of inactive sublocations for the specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn sublocations_get_inactive_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.sublocations_get_inactive_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// PUT Update V2
    /// Updates existing sublocation records for a specified Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Locations
    ///
    pub async fn sublocations_update_v2(&self, license_number: Option<String>, body: Option<&Vec<SublocationsUpdateV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.sublocations_update_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateExternalIncoming V1
    /// Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_create_external_incoming_v1(&self, license_number: Option<String>, body: Option<&Vec<TransfersCreateExternalIncomingV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.transfers_create_external_incoming_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateExternalIncoming V2
    /// Creates external incoming shipment plans for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///
    pub async fn transfers_create_external_incoming_v2(&self, license_number: Option<String>, body: Option<&Vec<TransfersCreateExternalIncomingV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.transfers_create_external_incoming_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateTemplates V1
    /// Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_create_templates_v1(&self, license_number: Option<String>, body: Option<&Vec<TransfersCreateTemplatesV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.transfers_create_templates_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// POST CreateTemplatesOutgoing V2
    /// Creates new transfer templates for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///
    pub async fn transfers_create_templates_outgoing_v2(&self, license_number: Option<String>, body: Option<&Vec<TransfersCreateTemplatesOutgoingV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.transfers_create_templates_outgoing_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// DELETE DeleteExternalIncoming V1
    /// Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_delete_external_incoming_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.transfers_delete_external_incoming_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE DeleteExternalIncoming V2
    /// Voids an external incoming shipment plan for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///
    pub async fn transfers_delete_external_incoming_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.transfers_delete_external_incoming_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE DeleteTemplates V1
    /// Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_delete_templates_v1(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.transfers_delete_templates_v1(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// DELETE DeleteTemplatesOutgoing V2
    /// Archives a transfer template for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///
    pub async fn transfers_delete_templates_outgoing_v2(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let id = id.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.transfers_delete_templates_outgoing_v2(&id, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveriesPackagesStates V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn transfers_get_deliveries_packages_states_v1(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.transfers_get_deliveries_packages_states_v1(no, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveriesPackagesStates V2
    /// Returns a list of available shipment Package states.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn transfers_get_deliveries_packages_states_v2(&self, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.transfers_get_deliveries_packages_states_v2(no, body.as_ref()).await }
        }).await
    }

    /// GET GetDelivery V1
    /// Please note: that the {id} parameter above represents a Shipment Plan ID.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_delivery_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.transfers_get_delivery_v1(&id, no, body.as_ref()).await }
        }).await
    }

    /// GET GetDelivery V2
    /// Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_delivery_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_delivery_v2(&id, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveryPackage V1
    /// Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_delivery_package_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.transfers_get_delivery_package_v1(&id, no, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveryPackage V2
    /// Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_delivery_package_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_delivery_package_v2(&id, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveryPackageRequiredlabtestbatches V1
    /// Please note: The {id} parameter above represents a Transfer Delivery Package ID, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_delivery_package_requiredlabtestbatches_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.transfers_get_delivery_package_requiredlabtestbatches_v1(&id, no, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveryPackageRequiredlabtestbatches V2
    /// Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_delivery_package_requiredlabtestbatches_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_delivery_package_requiredlabtestbatches_v2(&id, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveryPackageWholesale V1
    /// Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_delivery_package_wholesale_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.transfers_get_delivery_package_wholesale_v1(&id, no, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveryPackageWholesale V2
    /// Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_delivery_package_wholesale_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_delivery_package_wholesale_v2(&id, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveryTransporters V1
    /// Please note: that the {id} parameter above represents a Shipment Delivery ID.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_delivery_transporters_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.transfers_get_delivery_transporters_v1(&id, no, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveryTransporters V2
    /// Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_delivery_transporters_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_delivery_transporters_v2(&id, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveryTransportersDetails V1
    /// Please note: The {id} parameter above represents a Shipment Delivery ID.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_delivery_transporters_details_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.transfers_get_delivery_transporters_details_v1(&id, no, body.as_ref()).await }
        }).await
    }

    /// GET GetDeliveryTransportersDetails V2
    /// Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_delivery_transporters_details_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_delivery_transporters_details_v2(&id, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetHub V2
    /// Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_hub_v2(&self, estimated_arrival_end: Option<String>, estimated_arrival_start: Option<String>, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let estimated_arrival_end = estimated_arrival_end.clone();
            let estimated_arrival_start = estimated_arrival_start.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_hub_v2(estimated_arrival_end, estimated_arrival_start, last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetIncoming V1
    /// Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_incoming_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.transfers_get_incoming_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetIncoming V2
    /// Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_incoming_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_incoming_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetOutgoing V1
    /// Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_outgoing_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.transfers_get_outgoing_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetOutgoing V2
    /// Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_outgoing_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_outgoing_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetRejected V1
    /// Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_rejected_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.transfers_get_rejected_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetRejected V2
    /// Retrieves a list of shipments with rejected packages for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///   - View Transfers
    ///
    pub async fn transfers_get_rejected_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_rejected_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetTemplates V1
    /// Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_get_templates_v1(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.transfers_get_templates_v1(last_modified_end, last_modified_start, license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetTemplatesDelivery V1
    /// Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_get_templates_delivery_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.transfers_get_templates_delivery_v1(&id, no, body.as_ref()).await }
        }).await
    }

    /// GET GetTemplatesDeliveryPackage V1
    /// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_get_templates_delivery_package_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.transfers_get_templates_delivery_package_v1(&id, no, body.as_ref()).await }
        }).await
    }

    /// GET GetTemplatesDeliveryTransporters V1
    /// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_get_templates_delivery_transporters_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.transfers_get_templates_delivery_transporters_v1(&id, no, body.as_ref()).await }
        }).await
    }

    /// GET GetTemplatesDeliveryTransportersDetails V1
    /// Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_get_templates_delivery_transporters_details_v1(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.transfers_get_templates_delivery_transporters_details_v1(&id, no, body.as_ref()).await }
        }).await
    }

    /// GET GetTemplatesOutgoing V2
    /// Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///   - View Transfer Templates
    ///
    pub async fn transfers_get_templates_outgoing_v2(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let last_modified_end = last_modified_end.clone();
            let last_modified_start = last_modified_start.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_templates_outgoing_v2(last_modified_end, last_modified_start, license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetTemplatesOutgoingDelivery V2
    /// Retrieves a list of deliveries associated with a specific transfer template.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///   - View Transfer Templates
    ///
    pub async fn transfers_get_templates_outgoing_delivery_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_templates_outgoing_delivery_v2(&id, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetTemplatesOutgoingDeliveryPackage V2
    /// Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///   - View Transfer Templates
    ///
    pub async fn transfers_get_templates_outgoing_delivery_package_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_templates_outgoing_delivery_package_v2(&id, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetTemplatesOutgoingDeliveryTransporters V2
    /// Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///   - View Transfer Templates
    ///
    pub async fn transfers_get_templates_outgoing_delivery_transporters_v2(&self, id: &str, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_templates_outgoing_delivery_transporters_v2(&id, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// GET GetTemplatesOutgoingDeliveryTransportersDetails V2
    /// Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///   - View Transfer Templates
    ///
    pub async fn transfers_get_templates_outgoing_delivery_transporters_details_v2(&self, id: &str, no: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let id = id.to_string();
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let id = id.clone();
            let no = no.clone();
            let body = body.clone();
            async move { client.transfers_get_templates_outgoing_delivery_transporters_details_v2(&id, no, body.as_ref()).await }
        }).await
    }

    /// GET GetTypes V1
    /// Permissions Required:
    ///   - None
    ///
    pub async fn transfers_get_types_v1(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body = body.clone();
            async move { client.transfers_get_types_v1(license_number, body.as_ref()).await }
        }).await
    }

    /// GET GetTypes V2
    /// Retrieves a list of available transfer types for a Facility based on its license number.
    /// 
    ///   Permissions Required:
    ///   - None
    ///
    pub async fn transfers_get_types_v2(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error>> {
        let client = self.client.clone();
        let body = body.cloned();
        self.rate_limiter.execute(None, true, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let page_number = page_number.clone();
            let page_size = page_size.clone();
            let body = body.clone();
            async move { client.transfers_get_types_v2(license_number, page_number, page_size, body.as_ref()).await }
        }).await
    }

    /// PUT UpdateExternalIncoming V1
    /// Permissions Required:
    ///   - Transfers
    ///
    pub async fn transfers_update_external_incoming_v1(&self, license_number: Option<String>, body: Option<&Vec<TransfersUpdateExternalIncomingV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.transfers_update_external_incoming_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateExternalIncoming V2
    /// Updates external incoming shipment plans for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfers
    ///
    pub async fn transfers_update_external_incoming_v2(&self, license_number: Option<String>, body: Option<&Vec<TransfersUpdateExternalIncomingV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.transfers_update_external_incoming_v2(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateTemplates V1
    /// Permissions Required:
    ///   - Transfer Templates
    ///
    pub async fn transfers_update_templates_v1(&self, license_number: Option<String>, body: Option<&Vec<TransfersUpdateTemplatesV1RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.transfers_update_templates_v1(license_number, body_val.as_ref()).await }
        }).await
    }

    /// PUT UpdateTemplatesOutgoing V2
    /// Updates existing transfer templates for a Facility.
    /// 
    ///   Permissions Required:
    ///   - Manage Transfer Templates
    ///
    pub async fn transfers_update_templates_outgoing_v2(&self, license_number: Option<String>, body: Option<&Vec<TransfersUpdateTemplatesOutgoingV2RequestItem>>) -> Result<Option<Value>, Box<dyn Error>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        let body_val = body_val.clone();
        self.rate_limiter.execute(None, false, move || {
            let client = client.clone();
            let license_number = license_number.clone();
            let body_val = body_val.clone();
            async move { client.transfers_update_templates_outgoing_v2(license_number, body_val.as_ref()).await }
        }).await
    }

}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateAdditivesV1RequestItem {
    #[serde(rename = "ActiveIngredients")]
    pub active_ingredients: Option<Vec<PlantsCreateAdditivesV1RequestItemActiveIngredients>>,
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "AdditiveType")]
    pub additive_type: Option<String>,
    #[serde(rename = "ApplicationDevice")]
    pub application_device: Option<String>,
    #[serde(rename = "EpaRegistrationNumber")]
    pub epa_registration_number: Option<String>,
    #[serde(rename = "PlantLabels")]
    pub plant_labels: Option<Vec<String>>,
    #[serde(rename = "ProductSupplier")]
    pub product_supplier: Option<String>,
    #[serde(rename = "ProductTradeName")]
    pub product_trade_name: Option<String>,
    #[serde(rename = "TotalAmountApplied")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "TotalAmountUnitOfMeasure")]
    pub total_amount_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateAdditivesV1RequestItemActiveIngredients {
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "Percentage")]
    pub percentage: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateAdditivesV2RequestItem {
    #[serde(rename = "ActiveIngredients")]
    pub active_ingredients: Option<Vec<PlantsCreateAdditivesV2RequestItemActiveIngredients>>,
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "AdditiveType")]
    pub additive_type: Option<String>,
    #[serde(rename = "ApplicationDevice")]
    pub application_device: Option<String>,
    #[serde(rename = "EpaRegistrationNumber")]
    pub epa_registration_number: Option<String>,
    #[serde(rename = "PlantLabels")]
    pub plant_labels: Option<Vec<String>>,
    #[serde(rename = "ProductSupplier")]
    pub product_supplier: Option<String>,
    #[serde(rename = "ProductTradeName")]
    pub product_trade_name: Option<String>,
    #[serde(rename = "TotalAmountApplied")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "TotalAmountUnitOfMeasure")]
    pub total_amount_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateAdditivesV2RequestItemActiveIngredients {
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "Percentage")]
    pub percentage: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateAdditivesBylocationV1RequestItem {
    #[serde(rename = "ActiveIngredients")]
    pub active_ingredients: Option<Vec<PlantsCreateAdditivesBylocationV1RequestItemActiveIngredients>>,
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "AdditiveType")]
    pub additive_type: Option<String>,
    #[serde(rename = "ApplicationDevice")]
    pub application_device: Option<String>,
    #[serde(rename = "EpaRegistrationNumber")]
    pub epa_registration_number: Option<String>,
    #[serde(rename = "LocationName")]
    pub location_name: Option<String>,
    #[serde(rename = "ProductSupplier")]
    pub product_supplier: Option<String>,
    #[serde(rename = "ProductTradeName")]
    pub product_trade_name: Option<String>,
    #[serde(rename = "SublocationName")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "TotalAmountApplied")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "TotalAmountUnitOfMeasure")]
    pub total_amount_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateAdditivesBylocationV1RequestItemActiveIngredients {
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "Percentage")]
    pub percentage: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateAdditivesBylocationV2RequestItem {
    #[serde(rename = "ActiveIngredients")]
    pub active_ingredients: Option<Vec<PlantsCreateAdditivesBylocationV2RequestItemActiveIngredients>>,
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "AdditiveType")]
    pub additive_type: Option<String>,
    #[serde(rename = "ApplicationDevice")]
    pub application_device: Option<String>,
    #[serde(rename = "EpaRegistrationNumber")]
    pub epa_registration_number: Option<String>,
    #[serde(rename = "LocationName")]
    pub location_name: Option<String>,
    #[serde(rename = "ProductSupplier")]
    pub product_supplier: Option<String>,
    #[serde(rename = "ProductTradeName")]
    pub product_trade_name: Option<String>,
    #[serde(rename = "SublocationName")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "TotalAmountApplied")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "TotalAmountUnitOfMeasure")]
    pub total_amount_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateAdditivesBylocationV2RequestItemActiveIngredients {
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "Percentage")]
    pub percentage: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "AdditivesTemplateName")]
    pub additives_template_name: Option<String>,
    #[serde(rename = "LocationName")]
    pub location_name: Option<String>,
    #[serde(rename = "Rate")]
    pub rate: Option<String>,
    #[serde(rename = "SublocationName")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "TotalAmountApplied")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "TotalAmountUnitOfMeasure")]
    pub total_amount_unit_of_measure: Option<String>,
    #[serde(rename = "Volume")]
    pub volume: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateAdditivesUsingtemplateV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "AdditivesTemplateName")]
    pub additives_template_name: Option<String>,
    #[serde(rename = "PlantLabels")]
    pub plant_labels: Option<Vec<String>>,
    #[serde(rename = "Rate")]
    pub rate: Option<String>,
    #[serde(rename = "TotalAmountApplied")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "TotalAmountUnitOfMeasure")]
    pub total_amount_unit_of_measure: Option<String>,
    #[serde(rename = "Volume")]
    pub volume: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateChangegrowthphasesV1RequestItem {
    #[serde(rename = "GrowthDate")]
    pub growth_date: Option<String>,
    #[serde(rename = "GrowthPhase")]
    pub growth_phase: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "NewLocation")]
    pub new_location: Option<String>,
    #[serde(rename = "NewSublocation")]
    pub new_sublocation: Option<String>,
    #[serde(rename = "NewTag")]
    pub new_tag: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateHarvestplantsV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "DryingLocation")]
    pub drying_location: Option<String>,
    #[serde(rename = "DryingSublocation")]
    pub drying_sublocation: Option<String>,
    #[serde(rename = "HarvestName")]
    pub harvest_name: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "Plant")]
    pub plant: Option<String>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "Weight")]
    pub weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateManicureV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "DryingLocation")]
    pub drying_location: Option<String>,
    #[serde(rename = "DryingSublocation")]
    pub drying_sublocation: Option<String>,
    #[serde(rename = "HarvestName")]
    pub harvest_name: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "Plant")]
    pub plant: Option<String>,
    #[serde(rename = "PlantCount")]
    pub plant_count: Option<i64>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "Weight")]
    pub weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateManicureplantsV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "DryingLocation")]
    pub drying_location: Option<String>,
    #[serde(rename = "DryingSublocation")]
    pub drying_sublocation: Option<String>,
    #[serde(rename = "HarvestName")]
    pub harvest_name: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "Plant")]
    pub plant: Option<String>,
    #[serde(rename = "PlantCount")]
    pub plant_count: Option<i64>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "Weight")]
    pub weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateMoveplantsV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreatePlantbatchPackageV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Count")]
    pub count: Option<i64>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PackageTag")]
    pub package_tag: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatchType")]
    pub plant_batch_type: Option<String>,
    #[serde(rename = "PlantLabel")]
    pub plant_label: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreatePlantbatchPackageV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Count")]
    pub count: Option<i64>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PackageTag")]
    pub package_tag: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatchType")]
    pub plant_batch_type: Option<String>,
    #[serde(rename = "PlantLabel")]
    pub plant_label: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreatePlantingsV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "LocationName")]
    pub location_name: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatchName")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "PlantBatchType")]
    pub plant_batch_type: Option<String>,
    #[serde(rename = "PlantCount")]
    pub plant_count: Option<i64>,
    #[serde(rename = "PlantLabel")]
    pub plant_label: Option<String>,
    #[serde(rename = "StrainName")]
    pub strain_name: Option<String>,
    #[serde(rename = "SublocationName")]
    pub sublocation_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreatePlantingsV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "LocationName")]
    pub location_name: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatchName")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "PlantBatchType")]
    pub plant_batch_type: Option<String>,
    #[serde(rename = "PlantCount")]
    pub plant_count: Option<i64>,
    #[serde(rename = "PlantLabel")]
    pub plant_label: Option<String>,
    #[serde(rename = "StrainName")]
    pub strain_name: Option<String>,
    #[serde(rename = "SublocationName")]
    pub sublocation_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateWasteV1RequestItem {
    #[serde(rename = "LocationName")]
    pub location_name: Option<String>,
    #[serde(rename = "MixedMaterial")]
    pub mixed_material: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PlantLabels")]
    pub plant_labels: Option<Vec<Value>>,
    #[serde(rename = "ReasonName")]
    pub reason_name: Option<String>,
    #[serde(rename = "SublocationName")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "UnitOfMeasureName")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteDate")]
    pub waste_date: Option<String>,
    #[serde(rename = "WasteMethodName")]
    pub waste_method_name: Option<String>,
    #[serde(rename = "WasteWeight")]
    pub waste_weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsCreateWasteV2RequestItem {
    #[serde(rename = "LocationName")]
    pub location_name: Option<String>,
    #[serde(rename = "MixedMaterial")]
    pub mixed_material: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PlantLabels")]
    pub plant_labels: Option<Vec<Value>>,
    #[serde(rename = "ReasonName")]
    pub reason_name: Option<String>,
    #[serde(rename = "SublocationName")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "UnitOfMeasureName")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteDate")]
    pub waste_date: Option<String>,
    #[serde(rename = "WasteMethodName")]
    pub waste_method_name: Option<String>,
    #[serde(rename = "WasteWeight")]
    pub waste_weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsUpdateAdjustV2RequestItem {
    #[serde(rename = "AdjustCount")]
    pub adjust_count: Option<i64>,
    #[serde(rename = "AdjustReason")]
    pub adjust_reason: Option<String>,
    #[serde(rename = "AdjustmentDate")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "ReasonNote")]
    pub reason_note: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsUpdateGrowthphaseV2RequestItem {
    #[serde(rename = "GrowthDate")]
    pub growth_date: Option<String>,
    #[serde(rename = "GrowthPhase")]
    pub growth_phase: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "NewLocation")]
    pub new_location: Option<String>,
    #[serde(rename = "NewSublocation")]
    pub new_sublocation: Option<String>,
    #[serde(rename = "NewTag")]
    pub new_tag: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsUpdateHarvestV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "DryingLocation")]
    pub drying_location: Option<String>,
    #[serde(rename = "DryingSublocation")]
    pub drying_sublocation: Option<String>,
    #[serde(rename = "HarvestName")]
    pub harvest_name: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "Plant")]
    pub plant: Option<String>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "Weight")]
    pub weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsUpdateLocationV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsUpdateMergeV2RequestItem {
    #[serde(rename = "MergeDate")]
    pub merge_date: Option<String>,
    #[serde(rename = "SourcePlantGroupLabel")]
    pub source_plant_group_label: Option<String>,
    #[serde(rename = "TargetPlantGroupLabel")]
    pub target_plant_group_label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsUpdateSplitV2RequestItem {
    #[serde(rename = "PlantCount")]
    pub plant_count: Option<i64>,
    #[serde(rename = "SourcePlantLabel")]
    pub source_plant_label: Option<String>,
    #[serde(rename = "SplitDate")]
    pub split_date: Option<String>,
    #[serde(rename = "StrainLabel")]
    pub strain_label: Option<String>,
    #[serde(rename = "TagLabel")]
    pub tag_label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsUpdateStrainV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "StrainId")]
    pub strain_id: Option<i64>,
    #[serde(rename = "StrainName")]
    pub strain_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantsUpdateTagV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "NewTag")]
    pub new_tag: Option<String>,
    #[serde(rename = "ReplaceDate")]
    pub replace_date: Option<String>,
    #[serde(rename = "TagId")]
    pub tag_id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct AdditivesTemplatesCreateV2RequestItem {
    #[serde(rename = "ActiveIngredients")]
    pub active_ingredients: Option<Vec<AdditivesTemplatesCreateV2RequestItemActiveIngredients>>,
    #[serde(rename = "AdditiveType")]
    pub additive_type: Option<String>,
    #[serde(rename = "ApplicationDevice")]
    pub application_device: Option<String>,
    #[serde(rename = "EpaRegistrationNumber")]
    pub epa_registration_number: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "ProductSupplier")]
    pub product_supplier: Option<String>,
    #[serde(rename = "ProductTradeName")]
    pub product_trade_name: Option<String>,
    #[serde(rename = "RestrictiveEntryIntervalQuantityDescription")]
    pub restrictive_entry_interval_quantity_description: Option<String>,
    #[serde(rename = "RestrictiveEntryIntervalTimeDescription")]
    pub restrictive_entry_interval_time_description: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct AdditivesTemplatesCreateV2RequestItemActiveIngredients {
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "Percentage")]
    pub percentage: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct AdditivesTemplatesUpdateV2RequestItem {
    #[serde(rename = "ActiveIngredients")]
    pub active_ingredients: Option<Vec<AdditivesTemplatesUpdateV2RequestItemActiveIngredients>>,
    #[serde(rename = "AdditiveType")]
    pub additive_type: Option<String>,
    #[serde(rename = "ApplicationDevice")]
    pub application_device: Option<String>,
    #[serde(rename = "EpaRegistrationNumber")]
    pub epa_registration_number: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "ProductSupplier")]
    pub product_supplier: Option<String>,
    #[serde(rename = "ProductTradeName")]
    pub product_trade_name: Option<String>,
    #[serde(rename = "RestrictiveEntryIntervalQuantityDescription")]
    pub restrictive_entry_interval_quantity_description: Option<String>,
    #[serde(rename = "RestrictiveEntryIntervalTimeDescription")]
    pub restrictive_entry_interval_time_description: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct AdditivesTemplatesUpdateV2RequestItemActiveIngredients {
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "Percentage")]
    pub percentage: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsCreateFinishV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsCreatePackageV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "DecontaminateProduct")]
    pub decontaminate_product: Option<bool>,
    #[serde(rename = "DecontaminationDate")]
    pub decontamination_date: Option<String>,
    #[serde(rename = "DecontaminationSteps")]
    pub decontamination_steps: Option<String>,
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Ingredients")]
    pub ingredients: Option<Vec<HarvestsCreatePackageV1RequestItemIngredients>>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsProductionBatch")]
    pub is_production_batch: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "LabTestStageId")]
    pub lab_test_stage_id: Option<i64>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "ProcessingJobTypeId")]
    pub processing_job_type_id: Option<i64>,
    #[serde(rename = "ProductRequiresDecontamination")]
    pub product_requires_decontamination: Option<bool>,
    #[serde(rename = "ProductRequiresRemediation")]
    pub product_requires_remediation: Option<bool>,
    #[serde(rename = "ProductionBatchNumber")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "RemediateProduct")]
    pub remediate_product: Option<bool>,
    #[serde(rename = "RemediationDate")]
    pub remediation_date: Option<String>,
    #[serde(rename = "RemediationMethodId")]
    pub remediation_method_id: Option<i64>,
    #[serde(rename = "RemediationSteps")]
    pub remediation_steps: Option<String>,
    #[serde(rename = "RequiredLabTestBatches")]
    pub required_lab_test_batches: Option<Vec<Value>>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsCreatePackageV1RequestItemIngredients {
    #[serde(rename = "HarvestId")]
    pub harvest_id: Option<i64>,
    #[serde(rename = "HarvestName")]
    pub harvest_name: Option<String>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "Weight")]
    pub weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsCreatePackageV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "DecontaminateProduct")]
    pub decontaminate_product: Option<bool>,
    #[serde(rename = "DecontaminationDate")]
    pub decontamination_date: Option<String>,
    #[serde(rename = "DecontaminationSteps")]
    pub decontamination_steps: Option<String>,
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Ingredients")]
    pub ingredients: Option<Vec<HarvestsCreatePackageV2RequestItemIngredients>>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsProductionBatch")]
    pub is_production_batch: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "LabTestStageId")]
    pub lab_test_stage_id: Option<i64>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "ProcessingJobTypeId")]
    pub processing_job_type_id: Option<i64>,
    #[serde(rename = "ProductRequiresDecontamination")]
    pub product_requires_decontamination: Option<bool>,
    #[serde(rename = "ProductRequiresRemediation")]
    pub product_requires_remediation: Option<bool>,
    #[serde(rename = "ProductionBatchNumber")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "RemediateProduct")]
    pub remediate_product: Option<bool>,
    #[serde(rename = "RemediationDate")]
    pub remediation_date: Option<String>,
    #[serde(rename = "RemediationMethodId")]
    pub remediation_method_id: Option<i64>,
    #[serde(rename = "RemediationSteps")]
    pub remediation_steps: Option<String>,
    #[serde(rename = "RequiredLabTestBatches")]
    pub required_lab_test_batches: Option<Vec<Value>>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsCreatePackageV2RequestItemIngredients {
    #[serde(rename = "HarvestId")]
    pub harvest_id: Option<i64>,
    #[serde(rename = "HarvestName")]
    pub harvest_name: Option<String>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "Weight")]
    pub weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsCreatePackageTestingV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "DecontaminateProduct")]
    pub decontaminate_product: Option<bool>,
    #[serde(rename = "DecontaminationDate")]
    pub decontamination_date: Option<String>,
    #[serde(rename = "DecontaminationSteps")]
    pub decontamination_steps: Option<String>,
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Ingredients")]
    pub ingredients: Option<Vec<HarvestsCreatePackageTestingV1RequestItemIngredients>>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsProductionBatch")]
    pub is_production_batch: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "LabTestStageId")]
    pub lab_test_stage_id: Option<i64>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "ProcessingJobTypeId")]
    pub processing_job_type_id: Option<i64>,
    #[serde(rename = "ProductRequiresDecontamination")]
    pub product_requires_decontamination: Option<bool>,
    #[serde(rename = "ProductRequiresRemediation")]
    pub product_requires_remediation: Option<bool>,
    #[serde(rename = "ProductionBatchNumber")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "RemediateProduct")]
    pub remediate_product: Option<bool>,
    #[serde(rename = "RemediationDate")]
    pub remediation_date: Option<String>,
    #[serde(rename = "RemediationMethodId")]
    pub remediation_method_id: Option<i64>,
    #[serde(rename = "RemediationSteps")]
    pub remediation_steps: Option<String>,
    #[serde(rename = "RequiredLabTestBatches")]
    pub required_lab_test_batches: Option<Vec<Value>>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsCreatePackageTestingV1RequestItemIngredients {
    #[serde(rename = "HarvestId")]
    pub harvest_id: Option<i64>,
    #[serde(rename = "HarvestName")]
    pub harvest_name: Option<String>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "Weight")]
    pub weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsCreatePackageTestingV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "DecontaminateProduct")]
    pub decontaminate_product: Option<bool>,
    #[serde(rename = "DecontaminationDate")]
    pub decontamination_date: Option<String>,
    #[serde(rename = "DecontaminationSteps")]
    pub decontamination_steps: Option<String>,
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Ingredients")]
    pub ingredients: Option<Vec<HarvestsCreatePackageTestingV2RequestItemIngredients>>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsProductionBatch")]
    pub is_production_batch: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "LabTestStageId")]
    pub lab_test_stage_id: Option<i64>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "ProcessingJobTypeId")]
    pub processing_job_type_id: Option<i64>,
    #[serde(rename = "ProductRequiresDecontamination")]
    pub product_requires_decontamination: Option<bool>,
    #[serde(rename = "ProductRequiresRemediation")]
    pub product_requires_remediation: Option<bool>,
    #[serde(rename = "ProductionBatchNumber")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "RemediateProduct")]
    pub remediate_product: Option<bool>,
    #[serde(rename = "RemediationDate")]
    pub remediation_date: Option<String>,
    #[serde(rename = "RemediationMethodId")]
    pub remediation_method_id: Option<i64>,
    #[serde(rename = "RemediationSteps")]
    pub remediation_steps: Option<String>,
    #[serde(rename = "RequiredLabTestBatches")]
    pub required_lab_test_batches: Option<Vec<Value>>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsCreatePackageTestingV2RequestItemIngredients {
    #[serde(rename = "HarvestId")]
    pub harvest_id: Option<i64>,
    #[serde(rename = "HarvestName")]
    pub harvest_name: Option<String>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "Weight")]
    pub weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsCreateRemoveWasteV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "WasteType")]
    pub waste_type: Option<String>,
    #[serde(rename = "WasteWeight")]
    pub waste_weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsCreateUnfinishV1RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsCreateWasteV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "UnitOfWeight")]
    pub unit_of_weight: Option<String>,
    #[serde(rename = "WasteType")]
    pub waste_type: Option<String>,
    #[serde(rename = "WasteWeight")]
    pub waste_weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsUpdateFinishV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsUpdateLocationV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "DryingLocation")]
    pub drying_location: Option<String>,
    #[serde(rename = "DryingSublocation")]
    pub drying_sublocation: Option<String>,
    #[serde(rename = "HarvestName")]
    pub harvest_name: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsUpdateMoveV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "DryingLocation")]
    pub drying_location: Option<String>,
    #[serde(rename = "DryingSublocation")]
    pub drying_sublocation: Option<String>,
    #[serde(rename = "HarvestName")]
    pub harvest_name: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsUpdateRenameV1RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "NewName")]
    pub new_name: Option<String>,
    #[serde(rename = "OldName")]
    pub old_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsUpdateRenameV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "NewName")]
    pub new_name: Option<String>,
    #[serde(rename = "OldName")]
    pub old_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsUpdateRestoreHarvestedPlantsV2RequestItem {
    #[serde(rename = "HarvestId")]
    pub harvest_id: Option<i64>,
    #[serde(rename = "PlantTags")]
    pub plant_tags: Option<Vec<String>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HarvestsUpdateUnfinishV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LabTestsCreateRecordV1RequestItem {
    #[serde(rename = "DocumentFileBase64")]
    pub document_file_base64: Option<String>,
    #[serde(rename = "DocumentFileName")]
    pub document_file_name: Option<String>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "ResultDate")]
    pub result_date: Option<String>,
    #[serde(rename = "Results")]
    pub results: Option<Vec<LabTestsCreateRecordV1RequestItemResults>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LabTestsCreateRecordV1RequestItemResults {
    #[serde(rename = "LabTestTypeName")]
    pub lab_test_type_name: Option<String>,
    #[serde(rename = "Notes")]
    pub notes: Option<String>,
    #[serde(rename = "Passed")]
    pub passed: Option<bool>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LabTestsCreateRecordV2RequestItem {
    #[serde(rename = "DocumentFileBase64")]
    pub document_file_base64: Option<String>,
    #[serde(rename = "DocumentFileName")]
    pub document_file_name: Option<String>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "ResultDate")]
    pub result_date: Option<String>,
    #[serde(rename = "Results")]
    pub results: Option<Vec<LabTestsCreateRecordV2RequestItemResults>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LabTestsCreateRecordV2RequestItemResults {
    #[serde(rename = "LabTestTypeName")]
    pub lab_test_type_name: Option<String>,
    #[serde(rename = "Notes")]
    pub notes: Option<String>,
    #[serde(rename = "Passed")]
    pub passed: Option<bool>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LabTestsUpdateLabtestdocumentV1RequestItem {
    #[serde(rename = "DocumentFileBase64")]
    pub document_file_base64: Option<String>,
    #[serde(rename = "DocumentFileName")]
    pub document_file_name: Option<String>,
    #[serde(rename = "LabTestResultId")]
    pub lab_test_result_id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LabTestsUpdateLabtestdocumentV2RequestItem {
    #[serde(rename = "DocumentFileBase64")]
    pub document_file_base64: Option<String>,
    #[serde(rename = "DocumentFileName")]
    pub document_file_name: Option<String>,
    #[serde(rename = "LabTestResultId")]
    pub lab_test_result_id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LabTestsUpdateResultReleaseV1RequestItem {
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LabTestsUpdateResultReleaseV2RequestItem {
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryV1RequestItem {
    #[serde(rename = "ConsumerId")]
    pub consumer_id: Option<i64>,
    #[serde(rename = "DriverEmployeeId")]
    pub driver_employee_id: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientAddressCity")]
    pub recipient_address_city: Option<String>,
    #[serde(rename = "RecipientAddressCounty")]
    pub recipient_address_county: Option<String>,
    #[serde(rename = "RecipientAddressPostalCode")]
    pub recipient_address_postal_code: Option<String>,
    #[serde(rename = "RecipientAddressState")]
    pub recipient_address_state: Option<String>,
    #[serde(rename = "RecipientAddressStreet1")]
    pub recipient_address_street1: Option<String>,
    #[serde(rename = "RecipientAddressStreet2")]
    pub recipient_address_street2: Option<String>,
    #[serde(rename = "RecipientName")]
    pub recipient_name: Option<String>,
    #[serde(rename = "RecipientZoneId")]
    pub recipient_zone_id: Option<i64>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesCreateDeliveryV1RequestItemTransactions>>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryV1RequestItemTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryV2RequestItem {
    #[serde(rename = "ConsumerId")]
    pub consumer_id: Option<i64>,
    #[serde(rename = "DriverEmployeeId")]
    pub driver_employee_id: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientAddressCity")]
    pub recipient_address_city: Option<String>,
    #[serde(rename = "RecipientAddressCounty")]
    pub recipient_address_county: Option<String>,
    #[serde(rename = "RecipientAddressPostalCode")]
    pub recipient_address_postal_code: Option<String>,
    #[serde(rename = "RecipientAddressState")]
    pub recipient_address_state: Option<String>,
    #[serde(rename = "RecipientAddressStreet1")]
    pub recipient_address_street1: Option<String>,
    #[serde(rename = "RecipientAddressStreet2")]
    pub recipient_address_street2: Option<String>,
    #[serde(rename = "RecipientName")]
    pub recipient_name: Option<String>,
    #[serde(rename = "RecipientZoneId")]
    pub recipient_zone_id: Option<i64>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesCreateDeliveryV2RequestItemTransactions>>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryV2RequestItemTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerV1RequestItem {
    #[serde(rename = "DateTime")]
    pub date_time: Option<String>,
    #[serde(rename = "Destinations")]
    pub destinations: Option<Vec<SalesCreateDeliveryRetailerV1RequestItemDestinations>>,
    #[serde(rename = "DriverEmployeeId")]
    pub driver_employee_id: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<SalesCreateDeliveryRetailerV1RequestItemPackages>>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerV1RequestItemDestinations {
    #[serde(rename = "ConsumerId")]
    pub consumer_id: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "RecipientAddressCity")]
    pub recipient_address_city: Option<String>,
    #[serde(rename = "RecipientAddressCounty")]
    pub recipient_address_county: Option<String>,
    #[serde(rename = "RecipientAddressPostalCode")]
    pub recipient_address_postal_code: Option<String>,
    #[serde(rename = "RecipientAddressState")]
    pub recipient_address_state: Option<String>,
    #[serde(rename = "RecipientAddressStreet1")]
    pub recipient_address_street1: Option<String>,
    #[serde(rename = "RecipientAddressStreet2")]
    pub recipient_address_street2: Option<String>,
    #[serde(rename = "RecipientName")]
    pub recipient_name: Option<String>,
    #[serde(rename = "RecipientZoneId")]
    pub recipient_zone_id: Option<String>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesCreateDeliveryRetailerV1RequestItemDestinationsTransactions>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerV1RequestItemDestinationsTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerV1RequestItemPackages {
    #[serde(rename = "DateTime")]
    pub date_time: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "TotalPrice")]
    pub total_price: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerV2RequestItem {
    #[serde(rename = "DateTime")]
    pub date_time: Option<String>,
    #[serde(rename = "Destinations")]
    pub destinations: Option<Vec<SalesCreateDeliveryRetailerV2RequestItemDestinations>>,
    #[serde(rename = "DriverEmployeeId")]
    pub driver_employee_id: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<SalesCreateDeliveryRetailerV2RequestItemPackages>>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerV2RequestItemDestinations {
    #[serde(rename = "ConsumerId")]
    pub consumer_id: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "RecipientAddressCity")]
    pub recipient_address_city: Option<String>,
    #[serde(rename = "RecipientAddressCounty")]
    pub recipient_address_county: Option<String>,
    #[serde(rename = "RecipientAddressPostalCode")]
    pub recipient_address_postal_code: Option<String>,
    #[serde(rename = "RecipientAddressState")]
    pub recipient_address_state: Option<String>,
    #[serde(rename = "RecipientAddressStreet1")]
    pub recipient_address_street1: Option<String>,
    #[serde(rename = "RecipientAddressStreet2")]
    pub recipient_address_street2: Option<String>,
    #[serde(rename = "RecipientName")]
    pub recipient_name: Option<String>,
    #[serde(rename = "RecipientZoneId")]
    pub recipient_zone_id: Option<String>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesCreateDeliveryRetailerV2RequestItemDestinationsTransactions>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerV2RequestItemDestinationsTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerV2RequestItemPackages {
    #[serde(rename = "DateTime")]
    pub date_time: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "TotalPrice")]
    pub total_price: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerDepartV1RequestItem {
    #[serde(rename = "RetailerDeliveryId")]
    pub retailer_delivery_id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerDepartV2RequestItem {
    #[serde(rename = "RetailerDeliveryId")]
    pub retailer_delivery_id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerEndV1RequestItem {
    #[serde(rename = "ActualArrivalDateTime")]
    pub actual_arrival_date_time: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<SalesCreateDeliveryRetailerEndV1RequestItemPackages>>,
    #[serde(rename = "RetailerDeliveryId")]
    pub retailer_delivery_id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerEndV1RequestItemPackages {
    #[serde(rename = "EndQuantity")]
    pub end_quantity: Option<i64>,
    #[serde(rename = "EndUnitOfMeasure")]
    pub end_unit_of_measure: Option<String>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerEndV2RequestItem {
    #[serde(rename = "ActualArrivalDateTime")]
    pub actual_arrival_date_time: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<SalesCreateDeliveryRetailerEndV2RequestItemPackages>>,
    #[serde(rename = "RetailerDeliveryId")]
    pub retailer_delivery_id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerEndV2RequestItemPackages {
    #[serde(rename = "EndQuantity")]
    pub end_quantity: Option<i64>,
    #[serde(rename = "EndUnitOfMeasure")]
    pub end_unit_of_measure: Option<String>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerRestockV1RequestItem {
    #[serde(rename = "DateTime")]
    pub date_time: Option<String>,
    #[serde(rename = "Destinations")]
    pub destinations: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<SalesCreateDeliveryRetailerRestockV1RequestItemPackages>>,
    #[serde(rename = "RetailerDeliveryId")]
    pub retailer_delivery_id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerRestockV1RequestItemPackages {
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "RemoveCurrentPackage")]
    pub remove_current_package: Option<bool>,
    #[serde(rename = "TotalPrice")]
    pub total_price: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerRestockV2RequestItem {
    #[serde(rename = "DateTime")]
    pub date_time: Option<String>,
    #[serde(rename = "Destinations")]
    pub destinations: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<SalesCreateDeliveryRetailerRestockV2RequestItemPackages>>,
    #[serde(rename = "RetailerDeliveryId")]
    pub retailer_delivery_id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerRestockV2RequestItemPackages {
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "RemoveCurrentPackage")]
    pub remove_current_package: Option<bool>,
    #[serde(rename = "TotalPrice")]
    pub total_price: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerSaleV1RequestItem {
    #[serde(rename = "ConsumerId")]
    pub consumer_id: Option<i64>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientAddressCity")]
    pub recipient_address_city: Option<String>,
    #[serde(rename = "RecipientAddressCounty")]
    pub recipient_address_county: Option<String>,
    #[serde(rename = "RecipientAddressPostalCode")]
    pub recipient_address_postal_code: Option<String>,
    #[serde(rename = "RecipientAddressState")]
    pub recipient_address_state: Option<String>,
    #[serde(rename = "RecipientAddressStreet1")]
    pub recipient_address_street1: Option<String>,
    #[serde(rename = "RecipientAddressStreet2")]
    pub recipient_address_street2: Option<String>,
    #[serde(rename = "RecipientName")]
    pub recipient_name: Option<String>,
    #[serde(rename = "RecipientZoneId")]
    pub recipient_zone_id: Option<i64>,
    #[serde(rename = "RetailerDeliveryId")]
    pub retailer_delivery_id: Option<i64>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesCreateDeliveryRetailerSaleV1RequestItemTransactions>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerSaleV1RequestItemTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerSaleV2RequestItem {
    #[serde(rename = "ConsumerId")]
    pub consumer_id: Option<i64>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientAddressCity")]
    pub recipient_address_city: Option<String>,
    #[serde(rename = "RecipientAddressCounty")]
    pub recipient_address_county: Option<String>,
    #[serde(rename = "RecipientAddressPostalCode")]
    pub recipient_address_postal_code: Option<String>,
    #[serde(rename = "RecipientAddressState")]
    pub recipient_address_state: Option<String>,
    #[serde(rename = "RecipientAddressStreet1")]
    pub recipient_address_street1: Option<String>,
    #[serde(rename = "RecipientAddressStreet2")]
    pub recipient_address_street2: Option<String>,
    #[serde(rename = "RecipientName")]
    pub recipient_name: Option<String>,
    #[serde(rename = "RecipientZoneId")]
    pub recipient_zone_id: Option<i64>,
    #[serde(rename = "RetailerDeliveryId")]
    pub retailer_delivery_id: Option<i64>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesCreateDeliveryRetailerSaleV2RequestItemTransactions>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateDeliveryRetailerSaleV2RequestItemTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateReceiptV1RequestItem {
    #[serde(rename = "CaregiverLicenseNumber")]
    pub caregiver_license_number: Option<String>,
    #[serde(rename = "ExternalReceiptNumber")]
    pub external_receipt_number: Option<String>,
    #[serde(rename = "IdentificationMethod")]
    pub identification_method: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PatientRegistrationLocationId")]
    pub patient_registration_location_id: Option<i64>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesCreateReceiptV1RequestItemTransactions>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateReceiptV1RequestItemTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateReceiptV2RequestItem {
    #[serde(rename = "CaregiverLicenseNumber")]
    pub caregiver_license_number: Option<String>,
    #[serde(rename = "ExternalReceiptNumber")]
    pub external_receipt_number: Option<String>,
    #[serde(rename = "IdentificationMethod")]
    pub identification_method: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PatientRegistrationLocationId")]
    pub patient_registration_location_id: Option<i64>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesCreateReceiptV2RequestItemTransactions>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateReceiptV2RequestItemTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesCreateTransactionByDateV1RequestItem {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryV1RequestItem {
    #[serde(rename = "ConsumerId")]
    pub consumer_id: Option<i64>,
    #[serde(rename = "DriverEmployeeId")]
    pub driver_employee_id: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientAddressCity")]
    pub recipient_address_city: Option<String>,
    #[serde(rename = "RecipientAddressCounty")]
    pub recipient_address_county: Option<String>,
    #[serde(rename = "RecipientAddressPostalCode")]
    pub recipient_address_postal_code: Option<String>,
    #[serde(rename = "RecipientAddressState")]
    pub recipient_address_state: Option<String>,
    #[serde(rename = "RecipientAddressStreet1")]
    pub recipient_address_street1: Option<String>,
    #[serde(rename = "RecipientAddressStreet2")]
    pub recipient_address_street2: Option<String>,
    #[serde(rename = "RecipientName")]
    pub recipient_name: Option<String>,
    #[serde(rename = "RecipientZoneId")]
    pub recipient_zone_id: Option<String>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesUpdateDeliveryV1RequestItemTransactions>>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryV1RequestItemTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryV2RequestItem {
    #[serde(rename = "ConsumerId")]
    pub consumer_id: Option<i64>,
    #[serde(rename = "DriverEmployeeId")]
    pub driver_employee_id: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientAddressCity")]
    pub recipient_address_city: Option<String>,
    #[serde(rename = "RecipientAddressCounty")]
    pub recipient_address_county: Option<String>,
    #[serde(rename = "RecipientAddressPostalCode")]
    pub recipient_address_postal_code: Option<String>,
    #[serde(rename = "RecipientAddressState")]
    pub recipient_address_state: Option<String>,
    #[serde(rename = "RecipientAddressStreet1")]
    pub recipient_address_street1: Option<String>,
    #[serde(rename = "RecipientAddressStreet2")]
    pub recipient_address_street2: Option<String>,
    #[serde(rename = "RecipientName")]
    pub recipient_name: Option<String>,
    #[serde(rename = "RecipientZoneId")]
    pub recipient_zone_id: Option<String>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesUpdateDeliveryV2RequestItemTransactions>>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryV2RequestItemTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryCompleteV1RequestItem {
    #[serde(rename = "AcceptedPackages")]
    pub accepted_packages: Option<Vec<String>>,
    #[serde(rename = "ActualArrivalDateTime")]
    pub actual_arrival_date_time: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "PaymentType")]
    pub payment_type: Option<String>,
    #[serde(rename = "ReturnedPackages")]
    pub returned_packages: Option<Vec<SalesUpdateDeliveryCompleteV1RequestItemReturnedPackages>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryCompleteV1RequestItemReturnedPackages {
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "ReturnQuantityVerified")]
    pub return_quantity_verified: Option<i64>,
    #[serde(rename = "ReturnReason")]
    pub return_reason: Option<String>,
    #[serde(rename = "ReturnReasonNote")]
    pub return_reason_note: Option<String>,
    #[serde(rename = "ReturnUnitOfMeasure")]
    pub return_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryCompleteV2RequestItem {
    #[serde(rename = "AcceptedPackages")]
    pub accepted_packages: Option<Vec<String>>,
    #[serde(rename = "ActualArrivalDateTime")]
    pub actual_arrival_date_time: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "PaymentType")]
    pub payment_type: Option<String>,
    #[serde(rename = "ReturnedPackages")]
    pub returned_packages: Option<Vec<SalesUpdateDeliveryCompleteV2RequestItemReturnedPackages>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryCompleteV2RequestItemReturnedPackages {
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "ReturnQuantityVerified")]
    pub return_quantity_verified: Option<i64>,
    #[serde(rename = "ReturnReason")]
    pub return_reason: Option<String>,
    #[serde(rename = "ReturnReasonNote")]
    pub return_reason_note: Option<String>,
    #[serde(rename = "ReturnUnitOfMeasure")]
    pub return_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryHubV1RequestItem {
    #[serde(rename = "DriverEmployeeId")]
    pub driver_employee_id: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "TransporterFacilityId")]
    pub transporter_facility_id: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryHubV2RequestItem {
    #[serde(rename = "DriverEmployeeId")]
    pub driver_employee_id: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "TransporterFacilityId")]
    pub transporter_facility_id: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryHubAcceptV1RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryHubAcceptV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryHubDepartV1RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryHubDepartV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryHubVerifyIdV1RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "PaymentType")]
    pub payment_type: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryHubVerifyIdV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "PaymentType")]
    pub payment_type: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryRetailerV1RequestItem {
    #[serde(rename = "DateTime")]
    pub date_time: Option<String>,
    #[serde(rename = "Destinations")]
    pub destinations: Option<Vec<SalesUpdateDeliveryRetailerV1RequestItemDestinations>>,
    #[serde(rename = "DriverEmployeeId")]
    pub driver_employee_id: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<SalesUpdateDeliveryRetailerV1RequestItemPackages>>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryRetailerV1RequestItemDestinations {
    #[serde(rename = "ConsumerId")]
    pub consumer_id: Option<String>,
    #[serde(rename = "DriverEmployeeId")]
    pub driver_employee_id: Option<i64>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientAddressCity")]
    pub recipient_address_city: Option<String>,
    #[serde(rename = "RecipientAddressCounty")]
    pub recipient_address_county: Option<String>,
    #[serde(rename = "RecipientAddressPostalCode")]
    pub recipient_address_postal_code: Option<String>,
    #[serde(rename = "RecipientAddressState")]
    pub recipient_address_state: Option<String>,
    #[serde(rename = "RecipientAddressStreet1")]
    pub recipient_address_street1: Option<String>,
    #[serde(rename = "RecipientAddressStreet2")]
    pub recipient_address_street2: Option<String>,
    #[serde(rename = "RecipientName")]
    pub recipient_name: Option<String>,
    #[serde(rename = "RecipientZoneId")]
    pub recipient_zone_id: Option<String>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesUpdateDeliveryRetailerV1RequestItemDestinationsTransactions>>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryRetailerV1RequestItemDestinationsTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryRetailerV1RequestItemPackages {
    #[serde(rename = "DateTime")]
    pub date_time: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "TotalPrice")]
    pub total_price: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryRetailerV2RequestItem {
    #[serde(rename = "DateTime")]
    pub date_time: Option<String>,
    #[serde(rename = "Destinations")]
    pub destinations: Option<Vec<SalesUpdateDeliveryRetailerV2RequestItemDestinations>>,
    #[serde(rename = "DriverEmployeeId")]
    pub driver_employee_id: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<SalesUpdateDeliveryRetailerV2RequestItemPackages>>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryRetailerV2RequestItemDestinations {
    #[serde(rename = "ConsumerId")]
    pub consumer_id: Option<String>,
    #[serde(rename = "DriverEmployeeId")]
    pub driver_employee_id: Option<i64>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientAddressCity")]
    pub recipient_address_city: Option<String>,
    #[serde(rename = "RecipientAddressCounty")]
    pub recipient_address_county: Option<String>,
    #[serde(rename = "RecipientAddressPostalCode")]
    pub recipient_address_postal_code: Option<String>,
    #[serde(rename = "RecipientAddressState")]
    pub recipient_address_state: Option<String>,
    #[serde(rename = "RecipientAddressStreet1")]
    pub recipient_address_street1: Option<String>,
    #[serde(rename = "RecipientAddressStreet2")]
    pub recipient_address_street2: Option<String>,
    #[serde(rename = "RecipientName")]
    pub recipient_name: Option<String>,
    #[serde(rename = "RecipientZoneId")]
    pub recipient_zone_id: Option<String>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesUpdateDeliveryRetailerV2RequestItemDestinationsTransactions>>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryRetailerV2RequestItemDestinationsTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateDeliveryRetailerV2RequestItemPackages {
    #[serde(rename = "DateTime")]
    pub date_time: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "TotalPrice")]
    pub total_price: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateReceiptV1RequestItem {
    #[serde(rename = "CaregiverLicenseNumber")]
    pub caregiver_license_number: Option<String>,
    #[serde(rename = "ExternalReceiptNumber")]
    pub external_receipt_number: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "IdentificationMethod")]
    pub identification_method: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PatientRegistrationLocationId")]
    pub patient_registration_location_id: Option<i64>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesUpdateReceiptV1RequestItemTransactions>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateReceiptV1RequestItemTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateReceiptV2RequestItem {
    #[serde(rename = "CaregiverLicenseNumber")]
    pub caregiver_license_number: Option<String>,
    #[serde(rename = "ExternalReceiptNumber")]
    pub external_receipt_number: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "IdentificationMethod")]
    pub identification_method: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PatientRegistrationLocationId")]
    pub patient_registration_location_id: Option<i64>,
    #[serde(rename = "SalesCustomerType")]
    pub sales_customer_type: Option<String>,
    #[serde(rename = "SalesDateTime")]
    pub sales_date_time: Option<String>,
    #[serde(rename = "Transactions")]
    pub transactions: Option<Vec<SalesUpdateReceiptV2RequestItemTransactions>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateReceiptV2RequestItemTransactions {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateReceiptFinalizeV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateReceiptUnfinalizeV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SalesUpdateTransactionByDateV1RequestItem {
    #[serde(rename = "CityTax")]
    pub city_tax: Option<String>,
    #[serde(rename = "CountyTax")]
    pub county_tax: Option<String>,
    #[serde(rename = "DiscountAmount")]
    pub discount_amount: Option<String>,
    #[serde(rename = "ExciseTax")]
    pub excise_tax: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "MunicipalTax")]
    pub municipal_tax: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Price")]
    pub price: Option<String>,
    #[serde(rename = "QrCodes")]
    pub qr_codes: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SalesTax")]
    pub sales_tax: Option<String>,
    #[serde(rename = "SubTotal")]
    pub sub_total: Option<String>,
    #[serde(rename = "TotalAmount")]
    pub total_amount: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransportersCreateDriverV2RequestItem {
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EmployeeId")]
    pub employee_id: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransportersCreateVehicleV2RequestItem {
    #[serde(rename = "LicensePlateNumber")]
    pub license_plate_number: Option<String>,
    #[serde(rename = "Make")]
    pub make: Option<String>,
    #[serde(rename = "Model")]
    pub model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransportersUpdateDriverV2RequestItem {
    #[serde(rename = "DriversLicenseNumber")]
    pub drivers_license_number: Option<String>,
    #[serde(rename = "EmployeeId")]
    pub employee_id: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransportersUpdateVehicleV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<String>,
    #[serde(rename = "LicensePlateNumber")]
    pub license_plate_number: Option<String>,
    #[serde(rename = "Make")]
    pub make: Option<String>,
    #[serde(rename = "Model")]
    pub model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Ingredients")]
    pub ingredients: Option<Vec<PackagesCreateV1RequestItemIngredients>>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsProductionBatch")]
    pub is_production_batch: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "LabTestStageId")]
    pub lab_test_stage_id: Option<i64>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "ProcessingJobTypeId")]
    pub processing_job_type_id: Option<i64>,
    #[serde(rename = "ProductRequiresRemediation")]
    pub product_requires_remediation: Option<bool>,
    #[serde(rename = "ProductionBatchNumber")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "RequiredLabTestBatches")]
    pub required_lab_test_batches: Option<bool>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
    #[serde(rename = "UseSameItem")]
    pub use_same_item: Option<bool>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateV1RequestItemIngredients {
    #[serde(rename = "Package")]
    pub package: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Ingredients")]
    pub ingredients: Option<Vec<PackagesCreateV2RequestItemIngredients>>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsProductionBatch")]
    pub is_production_batch: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "LabTestStageId")]
    pub lab_test_stage_id: Option<i64>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "ProcessingJobTypeId")]
    pub processing_job_type_id: Option<i64>,
    #[serde(rename = "ProductRequiresRemediation")]
    pub product_requires_remediation: Option<bool>,
    #[serde(rename = "ProductionBatchNumber")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "RequiredLabTestBatches")]
    pub required_lab_test_batches: Option<bool>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
    #[serde(rename = "UseSameItem")]
    pub use_same_item: Option<bool>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateV2RequestItemIngredients {
    #[serde(rename = "Package")]
    pub package: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateAdjustV1RequestItem {
    #[serde(rename = "AdjustmentDate")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "AdjustmentReason")]
    pub adjustment_reason: Option<String>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "ReasonNote")]
    pub reason_note: Option<String>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateAdjustV2RequestItem {
    #[serde(rename = "AdjustmentDate")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "AdjustmentReason")]
    pub adjustment_reason: Option<String>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "ReasonNote")]
    pub reason_note: Option<String>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateChangeItemV1RequestItem {
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateChangeLocationV1RequestItem {
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "MoveDate")]
    pub move_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateFinishV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreatePlantingsV1RequestItem {
    #[serde(rename = "LocationName")]
    pub location_name: Option<String>,
    #[serde(rename = "PackageAdjustmentAmount")]
    pub package_adjustment_amount: Option<i64>,
    #[serde(rename = "PackageAdjustmentUnitOfMeasureName")]
    pub package_adjustment_unit_of_measure_name: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatchName")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "PlantBatchType")]
    pub plant_batch_type: Option<String>,
    #[serde(rename = "PlantCount")]
    pub plant_count: Option<i64>,
    #[serde(rename = "PlantedDate")]
    pub planted_date: Option<String>,
    #[serde(rename = "StrainName")]
    pub strain_name: Option<String>,
    #[serde(rename = "SublocationName")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "UnpackagedDate")]
    pub unpackaged_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreatePlantingsV2RequestItem {
    #[serde(rename = "LocationName")]
    pub location_name: Option<String>,
    #[serde(rename = "PackageAdjustmentAmount")]
    pub package_adjustment_amount: Option<i64>,
    #[serde(rename = "PackageAdjustmentUnitOfMeasureName")]
    pub package_adjustment_unit_of_measure_name: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatchName")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "PlantBatchType")]
    pub plant_batch_type: Option<String>,
    #[serde(rename = "PlantCount")]
    pub plant_count: Option<i64>,
    #[serde(rename = "PlantedDate")]
    pub planted_date: Option<String>,
    #[serde(rename = "StrainName")]
    pub strain_name: Option<String>,
    #[serde(rename = "SublocationName")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "UnpackagedDate")]
    pub unpackaged_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateRemediateV1RequestItem {
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "RemediationDate")]
    pub remediation_date: Option<String>,
    #[serde(rename = "RemediationMethodName")]
    pub remediation_method_name: Option<String>,
    #[serde(rename = "RemediationSteps")]
    pub remediation_steps: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateTestingV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Ingredients")]
    pub ingredients: Option<Vec<PackagesCreateTestingV1RequestItemIngredients>>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsProductionBatch")]
    pub is_production_batch: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "LabTestStageId")]
    pub lab_test_stage_id: Option<i64>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "ProcessingJobTypeId")]
    pub processing_job_type_id: Option<i64>,
    #[serde(rename = "ProductRequiresRemediation")]
    pub product_requires_remediation: Option<bool>,
    #[serde(rename = "ProductionBatchNumber")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "RequiredLabTestBatches")]
    pub required_lab_test_batches: Option<bool>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
    #[serde(rename = "UseSameItem")]
    pub use_same_item: Option<bool>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateTestingV1RequestItemIngredients {
    #[serde(rename = "Package")]
    pub package: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateTestingV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Ingredients")]
    pub ingredients: Option<Vec<PackagesCreateTestingV2RequestItemIngredients>>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsProductionBatch")]
    pub is_production_batch: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "LabTestStageId")]
    pub lab_test_stage_id: Option<i64>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "ProcessingJobTypeId")]
    pub processing_job_type_id: Option<i64>,
    #[serde(rename = "ProductRequiresRemediation")]
    pub product_requires_remediation: Option<bool>,
    #[serde(rename = "ProductionBatchNumber")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "RequiredLabTestBatches")]
    pub required_lab_test_batches: Option<Vec<String>>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
    #[serde(rename = "UseSameItem")]
    pub use_same_item: Option<bool>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateTestingV2RequestItemIngredients {
    #[serde(rename = "Package")]
    pub package: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesCreateUnfinishV1RequestItem {
    #[serde(rename = "Label")]
    pub label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateAdjustV2RequestItem {
    #[serde(rename = "AdjustmentDate")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "AdjustmentReason")]
    pub adjustment_reason: Option<String>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "ReasonNote")]
    pub reason_note: Option<String>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateChangeNoteV1RequestItem {
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateDecontaminateV2RequestItem {
    #[serde(rename = "DecontaminationDate")]
    pub decontamination_date: Option<String>,
    #[serde(rename = "DecontaminationMethodName")]
    pub decontamination_method_name: Option<String>,
    #[serde(rename = "DecontaminationSteps")]
    pub decontamination_steps: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateDonationFlagV2RequestItem {
    #[serde(rename = "Label")]
    pub label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateDonationUnflagV2RequestItem {
    #[serde(rename = "Label")]
    pub label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateExternalidV2RequestItem {
    #[serde(rename = "ExternalId")]
    pub external_id: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateFinishV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateItemV2RequestItem {
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateLabTestRequiredV2RequestItem {
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "RequiredLabTestBatches")]
    pub required_lab_test_batches: Option<Vec<String>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateLocationV2RequestItem {
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "MoveDate")]
    pub move_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateNoteV2RequestItem {
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateRemediateV2RequestItem {
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "RemediationDate")]
    pub remediation_date: Option<String>,
    #[serde(rename = "RemediationMethodName")]
    pub remediation_method_name: Option<String>,
    #[serde(rename = "RemediationSteps")]
    pub remediation_steps: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateTradesampleFlagV2RequestItem {
    #[serde(rename = "Label")]
    pub label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateTradesampleUnflagV2RequestItem {
    #[serde(rename = "Label")]
    pub label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateUnfinishV2RequestItem {
    #[serde(rename = "Label")]
    pub label: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PackagesUpdateUsebydateV2RequestItem {
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PatientCheckInsCreateV1RequestItem {
    #[serde(rename = "CheckInDate")]
    pub check_in_date: Option<String>,
    #[serde(rename = "CheckInLocationId")]
    pub check_in_location_id: Option<i64>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "RegistrationExpiryDate")]
    pub registration_expiry_date: Option<String>,
    #[serde(rename = "RegistrationStartDate")]
    pub registration_start_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PatientCheckInsCreateV2RequestItem {
    #[serde(rename = "CheckInDate")]
    pub check_in_date: Option<String>,
    #[serde(rename = "CheckInLocationId")]
    pub check_in_location_id: Option<i64>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "RegistrationExpiryDate")]
    pub registration_expiry_date: Option<String>,
    #[serde(rename = "RegistrationStartDate")]
    pub registration_start_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PatientCheckInsUpdateV1RequestItem {
    #[serde(rename = "CheckInDate")]
    pub check_in_date: Option<String>,
    #[serde(rename = "CheckInLocationId")]
    pub check_in_location_id: Option<i64>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "RegistrationExpiryDate")]
    pub registration_expiry_date: Option<String>,
    #[serde(rename = "RegistrationStartDate")]
    pub registration_start_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PatientCheckInsUpdateV2RequestItem {
    #[serde(rename = "CheckInDate")]
    pub check_in_date: Option<String>,
    #[serde(rename = "CheckInLocationId")]
    pub check_in_location_id: Option<i64>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "RegistrationExpiryDate")]
    pub registration_expiry_date: Option<String>,
    #[serde(rename = "RegistrationStartDate")]
    pub registration_start_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ItemsCreateV1RequestItem {
    #[serde(rename = "AdministrationMethod")]
    pub administration_method: Option<String>,
    #[serde(rename = "Allergens")]
    pub allergens: Option<String>,
    #[serde(rename = "Description")]
    pub description: Option<String>,
    #[serde(rename = "GlobalProductName")]
    pub global_product_name: Option<String>,
    #[serde(rename = "ItemBrand")]
    pub item_brand: Option<String>,
    #[serde(rename = "ItemCategory")]
    pub item_category: Option<String>,
    #[serde(rename = "ItemIngredients")]
    pub item_ingredients: Option<String>,
    #[serde(rename = "LabelImageFileSystemIds")]
    pub label_image_file_system_ids: Option<String>,
    #[serde(rename = "LabelPhotoDescription")]
    pub label_photo_description: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "NumberOfDoses")]
    pub number_of_doses: Option<String>,
    #[serde(rename = "PackagingImageFileSystemIds")]
    pub packaging_image_file_system_ids: Option<String>,
    #[serde(rename = "PackagingPhotoDescription")]
    pub packaging_photo_description: Option<String>,
    #[serde(rename = "ProcessingJobCategoryName")]
    pub processing_job_category_name: Option<String>,
    #[serde(rename = "ProcessingJobTypeName")]
    pub processing_job_type_name: Option<String>,
    #[serde(rename = "ProductImageFileSystemIds")]
    pub product_image_file_system_ids: Option<String>,
    #[serde(rename = "ProductPDFFileSystemIds")]
    pub product_pdf_file_system_ids: Option<String>,
    #[serde(rename = "ProductPhotoDescription")]
    pub product_photo_description: Option<String>,
    #[serde(rename = "PublicIngredients")]
    pub public_ingredients: Option<String>,
    #[serde(rename = "ServingSize")]
    pub serving_size: Option<String>,
    #[serde(rename = "Strain")]
    pub strain: Option<String>,
    #[serde(rename = "SupplyDurationDays")]
    pub supply_duration_days: Option<i64>,
    #[serde(rename = "UnitCbdAContent")]
    pub unit_cbd_a_content: Option<f64>,
    #[serde(rename = "UnitCbdAContentDose")]
    pub unit_cbd_a_content_dose: Option<f64>,
    #[serde(rename = "UnitCbdAContentDoseUnitOfMeasure")]
    pub unit_cbd_a_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdAContentUnitOfMeasure")]
    pub unit_cbd_a_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdAPercent")]
    pub unit_cbd_a_percent: Option<f64>,
    #[serde(rename = "UnitCbdContent")]
    pub unit_cbd_content: Option<f64>,
    #[serde(rename = "UnitCbdContentDose")]
    pub unit_cbd_content_dose: Option<f64>,
    #[serde(rename = "UnitCbdContentDoseUnitOfMeasure")]
    pub unit_cbd_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdContentUnitOfMeasure")]
    pub unit_cbd_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdPercent")]
    pub unit_cbd_percent: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAContent")]
    pub unit_thc_a_content: Option<f64>,
    #[serde(rename = "UnitThcAContentDose")]
    pub unit_thc_a_content_dose: Option<f64>,
    #[serde(rename = "UnitThcAContentDoseUnitOfMeasure")]
    pub unit_thc_a_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAContentUnitOfMeasure")]
    pub unit_thc_a_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAPercent")]
    pub unit_thc_a_percent: Option<f64>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentDose")]
    pub unit_thc_content_dose: Option<f64>,
    #[serde(rename = "UnitThcContentDoseUnitOfMeasure")]
    pub unit_thc_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitVolume")]
    pub unit_volume: Option<f64>,
    #[serde(rename = "UnitVolumeUnitOfMeasure")]
    pub unit_volume_unit_of_measure: Option<String>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ItemsCreateV2RequestItem {
    #[serde(rename = "AdministrationMethod")]
    pub administration_method: Option<String>,
    #[serde(rename = "Allergens")]
    pub allergens: Option<String>,
    #[serde(rename = "Description")]
    pub description: Option<String>,
    #[serde(rename = "GlobalProductName")]
    pub global_product_name: Option<String>,
    #[serde(rename = "ItemBrand")]
    pub item_brand: Option<String>,
    #[serde(rename = "ItemCategory")]
    pub item_category: Option<String>,
    #[serde(rename = "ItemIngredients")]
    pub item_ingredients: Option<String>,
    #[serde(rename = "LabelImageFileSystemIds")]
    pub label_image_file_system_ids: Option<String>,
    #[serde(rename = "LabelPhotoDescription")]
    pub label_photo_description: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "NumberOfDoses")]
    pub number_of_doses: Option<String>,
    #[serde(rename = "PackagingImageFileSystemIds")]
    pub packaging_image_file_system_ids: Option<String>,
    #[serde(rename = "PackagingPhotoDescription")]
    pub packaging_photo_description: Option<String>,
    #[serde(rename = "ProcessingJobCategoryName")]
    pub processing_job_category_name: Option<String>,
    #[serde(rename = "ProcessingJobTypeName")]
    pub processing_job_type_name: Option<String>,
    #[serde(rename = "ProductImageFileSystemIds")]
    pub product_image_file_system_ids: Option<String>,
    #[serde(rename = "ProductPDFFileSystemIds")]
    pub product_pdf_file_system_ids: Option<String>,
    #[serde(rename = "ProductPhotoDescription")]
    pub product_photo_description: Option<String>,
    #[serde(rename = "PublicIngredients")]
    pub public_ingredients: Option<String>,
    #[serde(rename = "ServingSize")]
    pub serving_size: Option<String>,
    #[serde(rename = "Strain")]
    pub strain: Option<String>,
    #[serde(rename = "SupplyDurationDays")]
    pub supply_duration_days: Option<i64>,
    #[serde(rename = "UnitCbdAContent")]
    pub unit_cbd_a_content: Option<f64>,
    #[serde(rename = "UnitCbdAContentDose")]
    pub unit_cbd_a_content_dose: Option<f64>,
    #[serde(rename = "UnitCbdAContentDoseUnitOfMeasure")]
    pub unit_cbd_a_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdAContentUnitOfMeasure")]
    pub unit_cbd_a_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdAPercent")]
    pub unit_cbd_a_percent: Option<f64>,
    #[serde(rename = "UnitCbdContent")]
    pub unit_cbd_content: Option<f64>,
    #[serde(rename = "UnitCbdContentDose")]
    pub unit_cbd_content_dose: Option<f64>,
    #[serde(rename = "UnitCbdContentDoseUnitOfMeasure")]
    pub unit_cbd_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdContentUnitOfMeasure")]
    pub unit_cbd_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdPercent")]
    pub unit_cbd_percent: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAContent")]
    pub unit_thc_a_content: Option<f64>,
    #[serde(rename = "UnitThcAContentDose")]
    pub unit_thc_a_content_dose: Option<f64>,
    #[serde(rename = "UnitThcAContentDoseUnitOfMeasure")]
    pub unit_thc_a_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAContentUnitOfMeasure")]
    pub unit_thc_a_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAPercent")]
    pub unit_thc_a_percent: Option<f64>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentDose")]
    pub unit_thc_content_dose: Option<f64>,
    #[serde(rename = "UnitThcContentDoseUnitOfMeasure")]
    pub unit_thc_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitVolume")]
    pub unit_volume: Option<f64>,
    #[serde(rename = "UnitVolumeUnitOfMeasure")]
    pub unit_volume_unit_of_measure: Option<String>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ItemsCreateBrandV2RequestItem {
    #[serde(rename = "Name")]
    pub name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ItemsCreateFileV2RequestItem {
    #[serde(rename = "EncodedImageBase64")]
    pub encoded_image_base64: Option<String>,
    #[serde(rename = "FileName")]
    pub file_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ItemsCreatePhotoV1RequestItem {
    #[serde(rename = "EncodedImageBase64")]
    pub encoded_image_base64: Option<String>,
    #[serde(rename = "FileName")]
    pub file_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ItemsCreatePhotoV2RequestItem {
    #[serde(rename = "EncodedImageBase64")]
    pub encoded_image_base64: Option<String>,
    #[serde(rename = "FileName")]
    pub file_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ItemsCreateUpdateV1RequestItem {
    #[serde(rename = "AdministrationMethod")]
    pub administration_method: Option<String>,
    #[serde(rename = "Allergens")]
    pub allergens: Option<String>,
    #[serde(rename = "Description")]
    pub description: Option<String>,
    #[serde(rename = "GlobalProductName")]
    pub global_product_name: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "ItemBrand")]
    pub item_brand: Option<String>,
    #[serde(rename = "ItemCategory")]
    pub item_category: Option<String>,
    #[serde(rename = "ItemIngredients")]
    pub item_ingredients: Option<String>,
    #[serde(rename = "LabelImageFileSystemIds")]
    pub label_image_file_system_ids: Option<String>,
    #[serde(rename = "LabelPhotoDescription")]
    pub label_photo_description: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "NumberOfDoses")]
    pub number_of_doses: Option<String>,
    #[serde(rename = "PackagingImageFileSystemIds")]
    pub packaging_image_file_system_ids: Option<String>,
    #[serde(rename = "PackagingPhotoDescription")]
    pub packaging_photo_description: Option<String>,
    #[serde(rename = "ProcessingJobCategoryName")]
    pub processing_job_category_name: Option<String>,
    #[serde(rename = "ProcessingJobTypeName")]
    pub processing_job_type_name: Option<String>,
    #[serde(rename = "ProductImageFileSystemIds")]
    pub product_image_file_system_ids: Option<String>,
    #[serde(rename = "ProductPDFFileSystemIds")]
    pub product_pdf_file_system_ids: Option<String>,
    #[serde(rename = "ProductPhotoDescription")]
    pub product_photo_description: Option<String>,
    #[serde(rename = "PublicIngredients")]
    pub public_ingredients: Option<String>,
    #[serde(rename = "ServingSize")]
    pub serving_size: Option<String>,
    #[serde(rename = "Strain")]
    pub strain: Option<String>,
    #[serde(rename = "SupplyDurationDays")]
    pub supply_duration_days: Option<i64>,
    #[serde(rename = "UnitCbdAContent")]
    pub unit_cbd_a_content: Option<f64>,
    #[serde(rename = "UnitCbdAContentDose")]
    pub unit_cbd_a_content_dose: Option<f64>,
    #[serde(rename = "UnitCbdAContentDoseUnitOfMeasure")]
    pub unit_cbd_a_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdAContentUnitOfMeasure")]
    pub unit_cbd_a_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdAPercent")]
    pub unit_cbd_a_percent: Option<f64>,
    #[serde(rename = "UnitCbdContent")]
    pub unit_cbd_content: Option<f64>,
    #[serde(rename = "UnitCbdContentDose")]
    pub unit_cbd_content_dose: Option<f64>,
    #[serde(rename = "UnitCbdContentDoseUnitOfMeasure")]
    pub unit_cbd_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdContentUnitOfMeasure")]
    pub unit_cbd_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdPercent")]
    pub unit_cbd_percent: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAContent")]
    pub unit_thc_a_content: Option<f64>,
    #[serde(rename = "UnitThcAContentDose")]
    pub unit_thc_a_content_dose: Option<f64>,
    #[serde(rename = "UnitThcAContentDoseUnitOfMeasure")]
    pub unit_thc_a_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAContentUnitOfMeasure")]
    pub unit_thc_a_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAPercent")]
    pub unit_thc_a_percent: Option<f64>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentDose")]
    pub unit_thc_content_dose: Option<f64>,
    #[serde(rename = "UnitThcContentDoseUnitOfMeasure")]
    pub unit_thc_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitVolume")]
    pub unit_volume: Option<f64>,
    #[serde(rename = "UnitVolumeUnitOfMeasure")]
    pub unit_volume_unit_of_measure: Option<String>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ItemsUpdateV2RequestItem {
    #[serde(rename = "AdministrationMethod")]
    pub administration_method: Option<String>,
    #[serde(rename = "Allergens")]
    pub allergens: Option<String>,
    #[serde(rename = "Description")]
    pub description: Option<String>,
    #[serde(rename = "GlobalProductName")]
    pub global_product_name: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "ItemBrand")]
    pub item_brand: Option<String>,
    #[serde(rename = "ItemCategory")]
    pub item_category: Option<String>,
    #[serde(rename = "ItemIngredients")]
    pub item_ingredients: Option<String>,
    #[serde(rename = "LabelImageFileSystemIds")]
    pub label_image_file_system_ids: Option<String>,
    #[serde(rename = "LabelPhotoDescription")]
    pub label_photo_description: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "NumberOfDoses")]
    pub number_of_doses: Option<String>,
    #[serde(rename = "PackagingImageFileSystemIds")]
    pub packaging_image_file_system_ids: Option<String>,
    #[serde(rename = "PackagingPhotoDescription")]
    pub packaging_photo_description: Option<String>,
    #[serde(rename = "ProcessingJobCategoryName")]
    pub processing_job_category_name: Option<String>,
    #[serde(rename = "ProcessingJobTypeName")]
    pub processing_job_type_name: Option<String>,
    #[serde(rename = "ProductImageFileSystemIds")]
    pub product_image_file_system_ids: Option<String>,
    #[serde(rename = "ProductPDFFileSystemIds")]
    pub product_pdf_file_system_ids: Option<String>,
    #[serde(rename = "ProductPhotoDescription")]
    pub product_photo_description: Option<String>,
    #[serde(rename = "PublicIngredients")]
    pub public_ingredients: Option<String>,
    #[serde(rename = "ServingSize")]
    pub serving_size: Option<String>,
    #[serde(rename = "Strain")]
    pub strain: Option<String>,
    #[serde(rename = "SupplyDurationDays")]
    pub supply_duration_days: Option<i64>,
    #[serde(rename = "UnitCbdAContent")]
    pub unit_cbd_a_content: Option<f64>,
    #[serde(rename = "UnitCbdAContentDose")]
    pub unit_cbd_a_content_dose: Option<f64>,
    #[serde(rename = "UnitCbdAContentDoseUnitOfMeasure")]
    pub unit_cbd_a_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdAContentUnitOfMeasure")]
    pub unit_cbd_a_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdAPercent")]
    pub unit_cbd_a_percent: Option<f64>,
    #[serde(rename = "UnitCbdContent")]
    pub unit_cbd_content: Option<f64>,
    #[serde(rename = "UnitCbdContentDose")]
    pub unit_cbd_content_dose: Option<f64>,
    #[serde(rename = "UnitCbdContentDoseUnitOfMeasure")]
    pub unit_cbd_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdContentUnitOfMeasure")]
    pub unit_cbd_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitCbdPercent")]
    pub unit_cbd_percent: Option<f64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAContent")]
    pub unit_thc_a_content: Option<f64>,
    #[serde(rename = "UnitThcAContentDose")]
    pub unit_thc_a_content_dose: Option<f64>,
    #[serde(rename = "UnitThcAContentDoseUnitOfMeasure")]
    pub unit_thc_a_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAContentUnitOfMeasure")]
    pub unit_thc_a_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcAPercent")]
    pub unit_thc_a_percent: Option<f64>,
    #[serde(rename = "UnitThcContent")]
    pub unit_thc_content: Option<f64>,
    #[serde(rename = "UnitThcContentDose")]
    pub unit_thc_content_dose: Option<f64>,
    #[serde(rename = "UnitThcContentDoseUnitOfMeasure")]
    pub unit_thc_content_dose_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcContentUnitOfMeasure")]
    pub unit_thc_content_unit_of_measure: Option<String>,
    #[serde(rename = "UnitThcPercent")]
    pub unit_thc_percent: Option<f64>,
    #[serde(rename = "UnitVolume")]
    pub unit_volume: Option<f64>,
    #[serde(rename = "UnitVolumeUnitOfMeasure")]
    pub unit_volume_unit_of_measure: Option<String>,
    #[serde(rename = "UnitWeight")]
    pub unit_weight: Option<f64>,
    #[serde(rename = "UnitWeightUnitOfMeasure")]
    pub unit_weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ItemsUpdateBrandV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LocationsCreateV1RequestItem {
    #[serde(rename = "LocationTypeName")]
    pub location_type_name: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LocationsCreateV2RequestItem {
    #[serde(rename = "LocationTypeName")]
    pub location_type_name: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LocationsCreateUpdateV1RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "LocationTypeName")]
    pub location_type_name: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct LocationsUpdateV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "LocationTypeName")]
    pub location_type_name: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PatientsCreateV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "ConcentrateOuncesAllowed")]
    pub concentrate_ounces_allowed: Option<i64>,
    #[serde(rename = "FlowerOuncesAllowed")]
    pub flower_ounces_allowed: Option<i64>,
    #[serde(rename = "HasSalesLimitExemption")]
    pub has_sales_limit_exemption: Option<bool>,
    #[serde(rename = "InfusedOuncesAllowed")]
    pub infused_ounces_allowed: Option<i64>,
    #[serde(rename = "LicenseEffectiveEndDate")]
    pub license_effective_end_date: Option<String>,
    #[serde(rename = "LicenseEffectiveStartDate")]
    pub license_effective_start_date: Option<String>,
    #[serde(rename = "LicenseNumber")]
    pub license_number: Option<String>,
    #[serde(rename = "MaxConcentrateThcPercentAllowed")]
    pub max_concentrate_thc_percent_allowed: Option<i64>,
    #[serde(rename = "MaxFlowerThcPercentAllowed")]
    pub max_flower_thc_percent_allowed: Option<i64>,
    #[serde(rename = "RecommendedPlants")]
    pub recommended_plants: Option<i64>,
    #[serde(rename = "RecommendedSmokableQuantity")]
    pub recommended_smokable_quantity: Option<i64>,
    #[serde(rename = "ThcOuncesAllowed")]
    pub thc_ounces_allowed: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PatientsCreateAddV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "ConcentrateOuncesAllowed")]
    pub concentrate_ounces_allowed: Option<i64>,
    #[serde(rename = "FlowerOuncesAllowed")]
    pub flower_ounces_allowed: Option<i64>,
    #[serde(rename = "HasSalesLimitExemption")]
    pub has_sales_limit_exemption: Option<bool>,
    #[serde(rename = "InfusedOuncesAllowed")]
    pub infused_ounces_allowed: Option<i64>,
    #[serde(rename = "LicenseEffectiveEndDate")]
    pub license_effective_end_date: Option<String>,
    #[serde(rename = "LicenseEffectiveStartDate")]
    pub license_effective_start_date: Option<String>,
    #[serde(rename = "LicenseNumber")]
    pub license_number: Option<String>,
    #[serde(rename = "MaxConcentrateThcPercentAllowed")]
    pub max_concentrate_thc_percent_allowed: Option<i64>,
    #[serde(rename = "MaxFlowerThcPercentAllowed")]
    pub max_flower_thc_percent_allowed: Option<i64>,
    #[serde(rename = "RecommendedPlants")]
    pub recommended_plants: Option<i64>,
    #[serde(rename = "RecommendedSmokableQuantity")]
    pub recommended_smokable_quantity: Option<i64>,
    #[serde(rename = "ThcOuncesAllowed")]
    pub thc_ounces_allowed: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PatientsCreateUpdateV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "ConcentrateOuncesAllowed")]
    pub concentrate_ounces_allowed: Option<i64>,
    #[serde(rename = "FlowerOuncesAllowed")]
    pub flower_ounces_allowed: Option<i64>,
    #[serde(rename = "HasSalesLimitExemption")]
    pub has_sales_limit_exemption: Option<bool>,
    #[serde(rename = "InfusedOuncesAllowed")]
    pub infused_ounces_allowed: Option<i64>,
    #[serde(rename = "LicenseEffectiveEndDate")]
    pub license_effective_end_date: Option<String>,
    #[serde(rename = "LicenseEffectiveStartDate")]
    pub license_effective_start_date: Option<String>,
    #[serde(rename = "LicenseNumber")]
    pub license_number: Option<String>,
    #[serde(rename = "MaxConcentrateThcPercentAllowed")]
    pub max_concentrate_thc_percent_allowed: Option<i64>,
    #[serde(rename = "MaxFlowerThcPercentAllowed")]
    pub max_flower_thc_percent_allowed: Option<i64>,
    #[serde(rename = "NewLicenseNumber")]
    pub new_license_number: Option<String>,
    #[serde(rename = "RecommendedPlants")]
    pub recommended_plants: Option<i64>,
    #[serde(rename = "RecommendedSmokableQuantity")]
    pub recommended_smokable_quantity: Option<i64>,
    #[serde(rename = "ThcOuncesAllowed")]
    pub thc_ounces_allowed: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PatientsUpdateV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "ConcentrateOuncesAllowed")]
    pub concentrate_ounces_allowed: Option<i64>,
    #[serde(rename = "FlowerOuncesAllowed")]
    pub flower_ounces_allowed: Option<i64>,
    #[serde(rename = "HasSalesLimitExemption")]
    pub has_sales_limit_exemption: Option<bool>,
    #[serde(rename = "InfusedOuncesAllowed")]
    pub infused_ounces_allowed: Option<i64>,
    #[serde(rename = "LicenseEffectiveEndDate")]
    pub license_effective_end_date: Option<String>,
    #[serde(rename = "LicenseEffectiveStartDate")]
    pub license_effective_start_date: Option<String>,
    #[serde(rename = "LicenseNumber")]
    pub license_number: Option<String>,
    #[serde(rename = "MaxConcentrateThcPercentAllowed")]
    pub max_concentrate_thc_percent_allowed: Option<i64>,
    #[serde(rename = "MaxFlowerThcPercentAllowed")]
    pub max_flower_thc_percent_allowed: Option<i64>,
    #[serde(rename = "NewLicenseNumber")]
    pub new_license_number: Option<String>,
    #[serde(rename = "RecommendedPlants")]
    pub recommended_plants: Option<i64>,
    #[serde(rename = "RecommendedSmokableQuantity")]
    pub recommended_smokable_quantity: Option<i64>,
    #[serde(rename = "ThcOuncesAllowed")]
    pub thc_ounces_allowed: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct RetailIdCreateAssociateV2RequestItem {
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "QrUrls")]
    pub qr_urls: Option<Vec<String>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct RetailIdCreateGenerateV2Request {
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct RetailIdCreateMergeV2Request {
    #[serde(rename = "packageLabels")]
    pub package_labels: Option<Vec<String>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct RetailIdCreatePackageInfoV2Request {
    #[serde(rename = "packageLabels")]
    pub package_labels: Option<Vec<String>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct StrainsCreateV1RequestItem {
    #[serde(rename = "CbdLevel")]
    pub cbd_level: Option<f64>,
    #[serde(rename = "IndicaPercentage")]
    pub indica_percentage: Option<f64>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "SativaPercentage")]
    pub sativa_percentage: Option<f64>,
    #[serde(rename = "TestingStatus")]
    pub testing_status: Option<String>,
    #[serde(rename = "ThcLevel")]
    pub thc_level: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct StrainsCreateV2RequestItem {
    #[serde(rename = "CbdLevel")]
    pub cbd_level: Option<f64>,
    #[serde(rename = "IndicaPercentage")]
    pub indica_percentage: Option<f64>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "SativaPercentage")]
    pub sativa_percentage: Option<f64>,
    #[serde(rename = "TestingStatus")]
    pub testing_status: Option<String>,
    #[serde(rename = "ThcLevel")]
    pub thc_level: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct StrainsCreateUpdateV1RequestItem {
    #[serde(rename = "CbdLevel")]
    pub cbd_level: Option<f64>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "IndicaPercentage")]
    pub indica_percentage: Option<f64>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "SativaPercentage")]
    pub sativa_percentage: Option<f64>,
    #[serde(rename = "TestingStatus")]
    pub testing_status: Option<String>,
    #[serde(rename = "ThcLevel")]
    pub thc_level: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct StrainsUpdateV2RequestItem {
    #[serde(rename = "CbdLevel")]
    pub cbd_level: Option<f64>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "IndicaPercentage")]
    pub indica_percentage: Option<f64>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "SativaPercentage")]
    pub sativa_percentage: Option<f64>,
    #[serde(rename = "TestingStatus")]
    pub testing_status: Option<String>,
    #[serde(rename = "ThcLevel")]
    pub thc_level: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateAdditivesV1RequestItem {
    #[serde(rename = "ActiveIngredients")]
    pub active_ingredients: Option<Vec<PlantBatchesCreateAdditivesV1RequestItemActiveIngredients>>,
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "AdditiveType")]
    pub additive_type: Option<String>,
    #[serde(rename = "ApplicationDevice")]
    pub application_device: Option<String>,
    #[serde(rename = "EpaRegistrationNumber")]
    pub epa_registration_number: Option<String>,
    #[serde(rename = "PlantBatchName")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "ProductSupplier")]
    pub product_supplier: Option<String>,
    #[serde(rename = "ProductTradeName")]
    pub product_trade_name: Option<String>,
    #[serde(rename = "TotalAmountApplied")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "TotalAmountUnitOfMeasure")]
    pub total_amount_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateAdditivesV1RequestItemActiveIngredients {
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "Percentage")]
    pub percentage: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateAdditivesV2RequestItem {
    #[serde(rename = "ActiveIngredients")]
    pub active_ingredients: Option<Vec<PlantBatchesCreateAdditivesV2RequestItemActiveIngredients>>,
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "AdditiveType")]
    pub additive_type: Option<String>,
    #[serde(rename = "ApplicationDevice")]
    pub application_device: Option<String>,
    #[serde(rename = "EpaRegistrationNumber")]
    pub epa_registration_number: Option<String>,
    #[serde(rename = "PlantBatchName")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "ProductSupplier")]
    pub product_supplier: Option<String>,
    #[serde(rename = "ProductTradeName")]
    pub product_trade_name: Option<String>,
    #[serde(rename = "TotalAmountApplied")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "TotalAmountUnitOfMeasure")]
    pub total_amount_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateAdditivesV2RequestItemActiveIngredients {
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "Percentage")]
    pub percentage: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateAdditivesUsingtemplateV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "AdditivesTemplateName")]
    pub additives_template_name: Option<String>,
    #[serde(rename = "PlantBatchName")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "Rate")]
    pub rate: Option<String>,
    #[serde(rename = "TotalAmountApplied")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "TotalAmountUnitOfMeasure")]
    pub total_amount_unit_of_measure: Option<String>,
    #[serde(rename = "Volume")]
    pub volume: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateAdjustV1RequestItem {
    #[serde(rename = "AdjustmentDate")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "AdjustmentReason")]
    pub adjustment_reason: Option<String>,
    #[serde(rename = "PlantBatchName")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "ReasonNote")]
    pub reason_note: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateAdjustV2RequestItem {
    #[serde(rename = "AdjustmentDate")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "AdjustmentReason")]
    pub adjustment_reason: Option<String>,
    #[serde(rename = "PlantBatchName")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "ReasonNote")]
    pub reason_note: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateChangegrowthphaseV1RequestItem {
    #[serde(rename = "Count")]
    pub count: Option<i64>,
    #[serde(rename = "CountPerPlant")]
    pub count_per_plant: Option<String>,
    #[serde(rename = "GrowthDate")]
    pub growth_date: Option<String>,
    #[serde(rename = "GrowthPhase")]
    pub growth_phase: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "NewLocation")]
    pub new_location: Option<String>,
    #[serde(rename = "NewSublocation")]
    pub new_sublocation: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "StartingTag")]
    pub starting_tag: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateGrowthphaseV2RequestItem {
    #[serde(rename = "Count")]
    pub count: Option<i64>,
    #[serde(rename = "CountPerPlant")]
    pub count_per_plant: Option<String>,
    #[serde(rename = "GrowthDate")]
    pub growth_date: Option<String>,
    #[serde(rename = "GrowthPhase")]
    pub growth_phase: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "NewLocation")]
    pub new_location: Option<String>,
    #[serde(rename = "NewSublocation")]
    pub new_sublocation: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "StartingTag")]
    pub starting_tag: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreatePackageV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Count")]
    pub count: Option<i64>,
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatch")]
    pub plant_batch: Option<String>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreatePackageFrommotherplantV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Count")]
    pub count: Option<i64>,
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatch")]
    pub plant_batch: Option<String>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreatePackageFrommotherplantV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Count")]
    pub count: Option<i64>,
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatch")]
    pub plant_batch: Option<String>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreatePlantingsV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Count")]
    pub count: Option<i64>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "SourcePlantBatches")]
    pub source_plant_batches: Option<String>,
    #[serde(rename = "Strain")]
    pub strain: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Type")]
    pub type_: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateSplitV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Count")]
    pub count: Option<i64>,
    #[serde(rename = "GroupName")]
    pub group_name: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatch")]
    pub plant_batch: Option<String>,
    #[serde(rename = "Strain")]
    pub strain: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateSplitV2RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Count")]
    pub count: Option<i64>,
    #[serde(rename = "GroupName")]
    pub group_name: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatch")]
    pub plant_batch: Option<String>,
    #[serde(rename = "Strain")]
    pub strain: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateWasteV1RequestItem {
    #[serde(rename = "MixedMaterial")]
    pub mixed_material: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PlantBatchName")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "ReasonName")]
    pub reason_name: Option<String>,
    #[serde(rename = "UnitOfMeasureName")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteDate")]
    pub waste_date: Option<String>,
    #[serde(rename = "WasteMethodName")]
    pub waste_method_name: Option<String>,
    #[serde(rename = "WasteWeight")]
    pub waste_weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateWasteV2RequestItem {
    #[serde(rename = "MixedMaterial")]
    pub mixed_material: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PlantBatchName")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "ReasonName")]
    pub reason_name: Option<String>,
    #[serde(rename = "UnitOfMeasureName")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteDate")]
    pub waste_date: Option<String>,
    #[serde(rename = "WasteMethodName")]
    pub waste_method_name: Option<String>,
    #[serde(rename = "WasteWeight")]
    pub waste_weight: Option<f64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreatepackagesV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Count")]
    pub count: Option<i64>,
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "IsDonation")]
    pub is_donation: Option<bool>,
    #[serde(rename = "IsTradeSample")]
    pub is_trade_sample: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatch")]
    pub plant_batch: Option<String>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesCreateplantingsV1RequestItem {
    #[serde(rename = "ActualDate")]
    pub actual_date: Option<String>,
    #[serde(rename = "Count")]
    pub count: Option<i64>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "SourcePlantBatches")]
    pub source_plant_batches: Option<String>,
    #[serde(rename = "Strain")]
    pub strain: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Type")]
    pub type_: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesUpdateLocationV2RequestItem {
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "MoveDate")]
    pub move_date: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesUpdateMoveplantbatchesV1RequestItem {
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "MoveDate")]
    pub move_date: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesUpdateNameV2RequestItem {
    #[serde(rename = "Group")]
    pub group: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "NewGroup")]
    pub new_group: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesUpdateStrainV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "StrainId")]
    pub strain_id: Option<i64>,
    #[serde(rename = "StrainName")]
    pub strain_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlantBatchesUpdateTagV2RequestItem {
    #[serde(rename = "Group")]
    pub group: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "NewTag")]
    pub new_tag: Option<String>,
    #[serde(rename = "ReplaceDate")]
    pub replace_date: Option<String>,
    #[serde(rename = "TagId")]
    pub tag_id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsCreateAdjustV1RequestItem {
    #[serde(rename = "AdjustmentDate")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "AdjustmentNote")]
    pub adjustment_note: Option<String>,
    #[serde(rename = "AdjustmentReason")]
    pub adjustment_reason: Option<String>,
    #[serde(rename = "CountUnitOfMeasureName")]
    pub count_unit_of_measure_name: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<ProcessingJobsCreateAdjustV1RequestItemPackages>>,
    #[serde(rename = "VolumeUnitOfMeasureName")]
    pub volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "WeightUnitOfMeasureName")]
    pub weight_unit_of_measure_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsCreateAdjustV1RequestItemPackages {
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsCreateAdjustV2RequestItem {
    #[serde(rename = "AdjustmentDate")]
    pub adjustment_date: Option<String>,
    #[serde(rename = "AdjustmentNote")]
    pub adjustment_note: Option<String>,
    #[serde(rename = "AdjustmentReason")]
    pub adjustment_reason: Option<String>,
    #[serde(rename = "CountUnitOfMeasureName")]
    pub count_unit_of_measure_name: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<ProcessingJobsCreateAdjustV2RequestItemPackages>>,
    #[serde(rename = "VolumeUnitOfMeasureName")]
    pub volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "WeightUnitOfMeasureName")]
    pub weight_unit_of_measure_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsCreateAdjustV2RequestItemPackages {
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsCreateJobtypesV1RequestItem {
    #[serde(rename = "Attributes")]
    pub attributes: Option<Vec<String>>,
    #[serde(rename = "Category")]
    pub category: Option<String>,
    #[serde(rename = "Description")]
    pub description: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "ProcessingSteps")]
    pub processing_steps: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsCreateJobtypesV2RequestItem {
    #[serde(rename = "Attributes")]
    pub attributes: Option<Vec<String>>,
    #[serde(rename = "Category")]
    pub category: Option<String>,
    #[serde(rename = "Description")]
    pub description: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "ProcessingSteps")]
    pub processing_steps: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsCreateStartV1RequestItem {
    #[serde(rename = "CountUnitOfMeasure")]
    pub count_unit_of_measure: Option<String>,
    #[serde(rename = "JobName")]
    pub job_name: Option<String>,
    #[serde(rename = "JobType")]
    pub job_type: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<ProcessingJobsCreateStartV1RequestItemPackages>>,
    #[serde(rename = "StartDate")]
    pub start_date: Option<String>,
    #[serde(rename = "VolumeUnitOfMeasure")]
    pub volume_unit_of_measure: Option<String>,
    #[serde(rename = "WeightUnitOfMeasure")]
    pub weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsCreateStartV1RequestItemPackages {
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsCreateStartV2RequestItem {
    #[serde(rename = "CountUnitOfMeasure")]
    pub count_unit_of_measure: Option<String>,
    #[serde(rename = "JobName")]
    pub job_name: Option<String>,
    #[serde(rename = "JobType")]
    pub job_type: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<ProcessingJobsCreateStartV2RequestItemPackages>>,
    #[serde(rename = "StartDate")]
    pub start_date: Option<String>,
    #[serde(rename = "VolumeUnitOfMeasure")]
    pub volume_unit_of_measure: Option<String>,
    #[serde(rename = "WeightUnitOfMeasure")]
    pub weight_unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsCreateStartV2RequestItemPackages {
    #[serde(rename = "Label")]
    pub label: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsCreatepackagesV1RequestItem {
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "FinishDate")]
    pub finish_date: Option<String>,
    #[serde(rename = "FinishNote")]
    pub finish_note: Option<String>,
    #[serde(rename = "FinishProcessingJob")]
    pub finish_processing_job: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "JobName")]
    pub job_name: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PackageDate")]
    pub package_date: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "ProductionBatchNumber")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
    #[serde(rename = "WasteCountQuantity")]
    pub waste_count_quantity: Option<String>,
    #[serde(rename = "WasteCountUnitOfMeasureName")]
    pub waste_count_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteVolumeQuantity")]
    pub waste_volume_quantity: Option<String>,
    #[serde(rename = "WasteVolumeUnitOfMeasureName")]
    pub waste_volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteWeightQuantity")]
    pub waste_weight_quantity: Option<String>,
    #[serde(rename = "WasteWeightUnitOfMeasureName")]
    pub waste_weight_unit_of_measure_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsCreatepackagesV2RequestItem {
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "FinishDate")]
    pub finish_date: Option<String>,
    #[serde(rename = "FinishNote")]
    pub finish_note: Option<String>,
    #[serde(rename = "FinishProcessingJob")]
    pub finish_processing_job: Option<bool>,
    #[serde(rename = "Item")]
    pub item: Option<String>,
    #[serde(rename = "JobName")]
    pub job_name: Option<String>,
    #[serde(rename = "Location")]
    pub location: Option<String>,
    #[serde(rename = "Note")]
    pub note: Option<String>,
    #[serde(rename = "PackageDate")]
    pub package_date: Option<String>,
    #[serde(rename = "PatientLicenseNumber")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "ProductionBatchNumber")]
    pub production_batch_number: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "Sublocation")]
    pub sublocation: Option<String>,
    #[serde(rename = "Tag")]
    pub tag: Option<String>,
    #[serde(rename = "UnitOfMeasure")]
    pub unit_of_measure: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
    #[serde(rename = "WasteCountQuantity")]
    pub waste_count_quantity: Option<String>,
    #[serde(rename = "WasteCountUnitOfMeasureName")]
    pub waste_count_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteVolumeQuantity")]
    pub waste_volume_quantity: Option<String>,
    #[serde(rename = "WasteVolumeUnitOfMeasureName")]
    pub waste_volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteWeightQuantity")]
    pub waste_weight_quantity: Option<String>,
    #[serde(rename = "WasteWeightUnitOfMeasureName")]
    pub waste_weight_unit_of_measure_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsUpdateFinishV1RequestItem {
    #[serde(rename = "FinishDate")]
    pub finish_date: Option<String>,
    #[serde(rename = "FinishNote")]
    pub finish_note: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "TotalCountWaste")]
    pub total_count_waste: Option<String>,
    #[serde(rename = "TotalVolumeWaste")]
    pub total_volume_waste: Option<String>,
    #[serde(rename = "TotalWeightWaste")]
    pub total_weight_waste: Option<i64>,
    #[serde(rename = "WasteCountUnitOfMeasureName")]
    pub waste_count_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteVolumeUnitOfMeasureName")]
    pub waste_volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteWeightUnitOfMeasureName")]
    pub waste_weight_unit_of_measure_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsUpdateFinishV2RequestItem {
    #[serde(rename = "FinishDate")]
    pub finish_date: Option<String>,
    #[serde(rename = "FinishNote")]
    pub finish_note: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "TotalCountWaste")]
    pub total_count_waste: Option<String>,
    #[serde(rename = "TotalVolumeWaste")]
    pub total_volume_waste: Option<String>,
    #[serde(rename = "TotalWeightWaste")]
    pub total_weight_waste: Option<i64>,
    #[serde(rename = "WasteCountUnitOfMeasureName")]
    pub waste_count_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteVolumeUnitOfMeasureName")]
    pub waste_volume_unit_of_measure_name: Option<String>,
    #[serde(rename = "WasteWeightUnitOfMeasureName")]
    pub waste_weight_unit_of_measure_name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsUpdateJobtypesV1RequestItem {
    #[serde(rename = "Attributes")]
    pub attributes: Option<Vec<String>>,
    #[serde(rename = "CategoryName")]
    pub category_name: Option<String>,
    #[serde(rename = "Description")]
    pub description: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "ProcessingSteps")]
    pub processing_steps: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsUpdateJobtypesV2RequestItem {
    #[serde(rename = "Attributes")]
    pub attributes: Option<Vec<String>>,
    #[serde(rename = "CategoryName")]
    pub category_name: Option<String>,
    #[serde(rename = "Description")]
    pub description: Option<String>,
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "ProcessingSteps")]
    pub processing_steps: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsUpdateUnfinishV1RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ProcessingJobsUpdateUnfinishV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SublocationsCreateV2RequestItem {
    #[serde(rename = "Name")]
    pub name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SublocationsUpdateV2RequestItem {
    #[serde(rename = "Id")]
    pub id: Option<i64>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateExternalIncomingV1RequestItem {
    #[serde(rename = "Destinations")]
    pub destinations: Option<Vec<TransfersCreateExternalIncomingV1RequestItemDestinations>>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "ShipperAddress1")]
    pub shipper_address1: Option<String>,
    #[serde(rename = "ShipperAddress2")]
    pub shipper_address2: Option<String>,
    #[serde(rename = "ShipperAddressCity")]
    pub shipper_address_city: Option<String>,
    #[serde(rename = "ShipperAddressPostalCode")]
    pub shipper_address_postal_code: Option<String>,
    #[serde(rename = "ShipperAddressState")]
    pub shipper_address_state: Option<String>,
    #[serde(rename = "ShipperLicenseNumber")]
    pub shipper_license_number: Option<String>,
    #[serde(rename = "ShipperMainPhoneNumber")]
    pub shipper_main_phone_number: Option<String>,
    #[serde(rename = "ShipperName")]
    pub shipper_name: Option<String>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateExternalIncomingV1RequestItemDestinations {
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "GrossUnitOfWeightId")]
    pub gross_unit_of_weight_id: Option<i64>,
    #[serde(rename = "GrossWeight")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<TransfersCreateExternalIncomingV1RequestItemDestinationsPackages>>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientLicenseNumber")]
    pub recipient_license_number: Option<String>,
    #[serde(rename = "TransferTypeName")]
    pub transfer_type_name: Option<String>,
    #[serde(rename = "Transporters")]
    pub transporters: Option<Vec<TransfersCreateExternalIncomingV1RequestItemDestinationsTransporters>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateExternalIncomingV1RequestItemDestinationsPackages {
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "ExternalId")]
    pub external_id: Option<String>,
    #[serde(rename = "GrossUnitOfWeightName")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "GrossWeight")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "ItemName")]
    pub item_name: Option<String>,
    #[serde(rename = "PackagedDate")]
    pub packaged_date: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "UnitOfMeasureName")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
    #[serde(rename = "WholesalePrice")]
    pub wholesale_price: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateExternalIncomingV1RequestItemDestinationsTransporters {
    #[serde(rename = "DriverLayoverLeg")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "IsLayover")]
    pub is_layover: Option<bool>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransporterDetails")]
    pub transporter_details: Option<String>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateExternalIncomingV2RequestItem {
    #[serde(rename = "Destinations")]
    pub destinations: Option<Vec<TransfersCreateExternalIncomingV2RequestItemDestinations>>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "ShipperAddress1")]
    pub shipper_address1: Option<String>,
    #[serde(rename = "ShipperAddress2")]
    pub shipper_address2: Option<String>,
    #[serde(rename = "ShipperAddressCity")]
    pub shipper_address_city: Option<String>,
    #[serde(rename = "ShipperAddressPostalCode")]
    pub shipper_address_postal_code: Option<String>,
    #[serde(rename = "ShipperAddressState")]
    pub shipper_address_state: Option<String>,
    #[serde(rename = "ShipperLicenseNumber")]
    pub shipper_license_number: Option<String>,
    #[serde(rename = "ShipperMainPhoneNumber")]
    pub shipper_main_phone_number: Option<String>,
    #[serde(rename = "ShipperName")]
    pub shipper_name: Option<String>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateExternalIncomingV2RequestItemDestinations {
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "GrossUnitOfWeightId")]
    pub gross_unit_of_weight_id: Option<i64>,
    #[serde(rename = "GrossWeight")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<TransfersCreateExternalIncomingV2RequestItemDestinationsPackages>>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientLicenseNumber")]
    pub recipient_license_number: Option<String>,
    #[serde(rename = "TransferTypeName")]
    pub transfer_type_name: Option<String>,
    #[serde(rename = "Transporters")]
    pub transporters: Option<Vec<TransfersCreateExternalIncomingV2RequestItemDestinationsTransporters>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateExternalIncomingV2RequestItemDestinationsPackages {
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "ExternalId")]
    pub external_id: Option<String>,
    #[serde(rename = "GrossUnitOfWeightName")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "GrossWeight")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "ItemName")]
    pub item_name: Option<String>,
    #[serde(rename = "PackagedDate")]
    pub packaged_date: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "UnitOfMeasureName")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
    #[serde(rename = "WholesalePrice")]
    pub wholesale_price: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateExternalIncomingV2RequestItemDestinationsTransporters {
    #[serde(rename = "DriverLayoverLeg")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "IsLayover")]
    pub is_layover: Option<bool>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransporterDetails")]
    pub transporter_details: Option<Vec<TransfersCreateExternalIncomingV2RequestItemDestinationsTransportersTransporterDetails>>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateExternalIncomingV2RequestItemDestinationsTransportersTransporterDetails {
    #[serde(rename = "DriverLayoverLeg")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateTemplatesV1RequestItem {
    #[serde(rename = "Destinations")]
    pub destinations: Option<Vec<TransfersCreateTemplatesV1RequestItemDestinations>>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateTemplatesV1RequestItemDestinations {
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<TransfersCreateTemplatesV1RequestItemDestinationsPackages>>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientLicenseNumber")]
    pub recipient_license_number: Option<String>,
    #[serde(rename = "TransferTypeName")]
    pub transfer_type_name: Option<String>,
    #[serde(rename = "Transporters")]
    pub transporters: Option<Vec<TransfersCreateTemplatesV1RequestItemDestinationsTransporters>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateTemplatesV1RequestItemDestinationsPackages {
    #[serde(rename = "GrossUnitOfWeightName")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "GrossWeight")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "WholesalePrice")]
    pub wholesale_price: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateTemplatesV1RequestItemDestinationsTransporters {
    #[serde(rename = "DriverLayoverLeg")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "IsLayover")]
    pub is_layover: Option<bool>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransporterDetails")]
    pub transporter_details: Option<String>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateTemplatesOutgoingV2RequestItem {
    #[serde(rename = "Destinations")]
    pub destinations: Option<Vec<TransfersCreateTemplatesOutgoingV2RequestItemDestinations>>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateTemplatesOutgoingV2RequestItemDestinations {
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<TransfersCreateTemplatesOutgoingV2RequestItemDestinationsPackages>>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientLicenseNumber")]
    pub recipient_license_number: Option<String>,
    #[serde(rename = "TransferTypeName")]
    pub transfer_type_name: Option<String>,
    #[serde(rename = "Transporters")]
    pub transporters: Option<Vec<TransfersCreateTemplatesOutgoingV2RequestItemDestinationsTransporters>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateTemplatesOutgoingV2RequestItemDestinationsPackages {
    #[serde(rename = "GrossUnitOfWeightName")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "GrossWeight")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "WholesalePrice")]
    pub wholesale_price: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateTemplatesOutgoingV2RequestItemDestinationsTransporters {
    #[serde(rename = "DriverLayoverLeg")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "IsLayover")]
    pub is_layover: Option<bool>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransporterDetails")]
    pub transporter_details: Option<Vec<TransfersCreateTemplatesOutgoingV2RequestItemDestinationsTransportersTransporterDetails>>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersCreateTemplatesOutgoingV2RequestItemDestinationsTransportersTransporterDetails {
    #[serde(rename = "DriverLayoverLeg")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateExternalIncomingV1RequestItem {
    #[serde(rename = "Destinations")]
    pub destinations: Option<Vec<TransfersUpdateExternalIncomingV1RequestItemDestinations>>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "ShipperAddress1")]
    pub shipper_address1: Option<String>,
    #[serde(rename = "ShipperAddress2")]
    pub shipper_address2: Option<String>,
    #[serde(rename = "ShipperAddressCity")]
    pub shipper_address_city: Option<String>,
    #[serde(rename = "ShipperAddressPostalCode")]
    pub shipper_address_postal_code: Option<String>,
    #[serde(rename = "ShipperAddressState")]
    pub shipper_address_state: Option<String>,
    #[serde(rename = "ShipperLicenseNumber")]
    pub shipper_license_number: Option<String>,
    #[serde(rename = "ShipperMainPhoneNumber")]
    pub shipper_main_phone_number: Option<String>,
    #[serde(rename = "ShipperName")]
    pub shipper_name: Option<String>,
    #[serde(rename = "TransferId")]
    pub transfer_id: Option<i64>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateExternalIncomingV1RequestItemDestinations {
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "GrossUnitOfWeightId")]
    pub gross_unit_of_weight_id: Option<i64>,
    #[serde(rename = "GrossWeight")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<TransfersUpdateExternalIncomingV1RequestItemDestinationsPackages>>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientLicenseNumber")]
    pub recipient_license_number: Option<String>,
    #[serde(rename = "TransferDestinationId")]
    pub transfer_destination_id: Option<i64>,
    #[serde(rename = "TransferTypeName")]
    pub transfer_type_name: Option<String>,
    #[serde(rename = "Transporters")]
    pub transporters: Option<Vec<TransfersUpdateExternalIncomingV1RequestItemDestinationsTransporters>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateExternalIncomingV1RequestItemDestinationsPackages {
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "ExternalId")]
    pub external_id: Option<String>,
    #[serde(rename = "GrossUnitOfWeightName")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "GrossWeight")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "ItemName")]
    pub item_name: Option<String>,
    #[serde(rename = "PackagedDate")]
    pub packaged_date: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "UnitOfMeasureName")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
    #[serde(rename = "WholesalePrice")]
    pub wholesale_price: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateExternalIncomingV1RequestItemDestinationsTransporters {
    #[serde(rename = "DriverLayoverLeg")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "IsLayover")]
    pub is_layover: Option<bool>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransporterDetails")]
    pub transporter_details: Option<String>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateExternalIncomingV2RequestItem {
    #[serde(rename = "Destinations")]
    pub destinations: Option<Vec<TransfersUpdateExternalIncomingV2RequestItemDestinations>>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "ShipperAddress1")]
    pub shipper_address1: Option<String>,
    #[serde(rename = "ShipperAddress2")]
    pub shipper_address2: Option<String>,
    #[serde(rename = "ShipperAddressCity")]
    pub shipper_address_city: Option<String>,
    #[serde(rename = "ShipperAddressPostalCode")]
    pub shipper_address_postal_code: Option<String>,
    #[serde(rename = "ShipperAddressState")]
    pub shipper_address_state: Option<String>,
    #[serde(rename = "ShipperLicenseNumber")]
    pub shipper_license_number: Option<String>,
    #[serde(rename = "ShipperMainPhoneNumber")]
    pub shipper_main_phone_number: Option<String>,
    #[serde(rename = "ShipperName")]
    pub shipper_name: Option<String>,
    #[serde(rename = "TransferId")]
    pub transfer_id: Option<i64>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateExternalIncomingV2RequestItemDestinations {
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "GrossUnitOfWeightId")]
    pub gross_unit_of_weight_id: Option<i64>,
    #[serde(rename = "GrossWeight")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<TransfersUpdateExternalIncomingV2RequestItemDestinationsPackages>>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientLicenseNumber")]
    pub recipient_license_number: Option<String>,
    #[serde(rename = "TransferDestinationId")]
    pub transfer_destination_id: Option<i64>,
    #[serde(rename = "TransferTypeName")]
    pub transfer_type_name: Option<String>,
    #[serde(rename = "Transporters")]
    pub transporters: Option<Vec<TransfersUpdateExternalIncomingV2RequestItemDestinationsTransporters>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateExternalIncomingV2RequestItemDestinationsPackages {
    #[serde(rename = "ExpirationDate")]
    pub expiration_date: Option<String>,
    #[serde(rename = "ExternalId")]
    pub external_id: Option<String>,
    #[serde(rename = "GrossUnitOfWeightName")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "GrossWeight")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "ItemName")]
    pub item_name: Option<String>,
    #[serde(rename = "PackagedDate")]
    pub packaged_date: Option<String>,
    #[serde(rename = "Quantity")]
    pub quantity: Option<i64>,
    #[serde(rename = "SellByDate")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "UnitOfMeasureName")]
    pub unit_of_measure_name: Option<String>,
    #[serde(rename = "UseByDate")]
    pub use_by_date: Option<String>,
    #[serde(rename = "WholesalePrice")]
    pub wholesale_price: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateExternalIncomingV2RequestItemDestinationsTransporters {
    #[serde(rename = "DriverLayoverLeg")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "IsLayover")]
    pub is_layover: Option<bool>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransporterDetails")]
    pub transporter_details: Option<Vec<TransfersUpdateExternalIncomingV2RequestItemDestinationsTransportersTransporterDetails>>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateExternalIncomingV2RequestItemDestinationsTransportersTransporterDetails {
    #[serde(rename = "DriverLayoverLeg")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateTemplatesV1RequestItem {
    #[serde(rename = "Destinations")]
    pub destinations: Option<Vec<TransfersUpdateTemplatesV1RequestItemDestinations>>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransferTemplateId")]
    pub transfer_template_id: Option<i64>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateTemplatesV1RequestItemDestinations {
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<TransfersUpdateTemplatesV1RequestItemDestinationsPackages>>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientLicenseNumber")]
    pub recipient_license_number: Option<String>,
    #[serde(rename = "TransferDestinationId")]
    pub transfer_destination_id: Option<i64>,
    #[serde(rename = "TransferTypeName")]
    pub transfer_type_name: Option<String>,
    #[serde(rename = "Transporters")]
    pub transporters: Option<Vec<TransfersUpdateTemplatesV1RequestItemDestinationsTransporters>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateTemplatesV1RequestItemDestinationsPackages {
    #[serde(rename = "GrossUnitOfWeightName")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "GrossWeight")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "WholesalePrice")]
    pub wholesale_price: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateTemplatesV1RequestItemDestinationsTransporters {
    #[serde(rename = "DriverLayoverLeg")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "IsLayover")]
    pub is_layover: Option<bool>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransporterDetails")]
    pub transporter_details: Option<String>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateTemplatesOutgoingV2RequestItem {
    #[serde(rename = "Destinations")]
    pub destinations: Option<Vec<TransfersUpdateTemplatesOutgoingV2RequestItemDestinations>>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "Name")]
    pub name: Option<String>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransferTemplateId")]
    pub transfer_template_id: Option<i64>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateTemplatesOutgoingV2RequestItemDestinations {
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "InvoiceNumber")]
    pub invoice_number: Option<String>,
    #[serde(rename = "Packages")]
    pub packages: Option<Vec<TransfersUpdateTemplatesOutgoingV2RequestItemDestinationsPackages>>,
    #[serde(rename = "PlannedRoute")]
    pub planned_route: Option<String>,
    #[serde(rename = "RecipientLicenseNumber")]
    pub recipient_license_number: Option<String>,
    #[serde(rename = "TransferDestinationId")]
    pub transfer_destination_id: Option<i64>,
    #[serde(rename = "TransferTypeName")]
    pub transfer_type_name: Option<String>,
    #[serde(rename = "Transporters")]
    pub transporters: Option<Vec<TransfersUpdateTemplatesOutgoingV2RequestItemDestinationsTransporters>>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateTemplatesOutgoingV2RequestItemDestinationsPackages {
    #[serde(rename = "GrossUnitOfWeightName")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "GrossWeight")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "PackageLabel")]
    pub package_label: Option<String>,
    #[serde(rename = "WholesalePrice")]
    pub wholesale_price: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateTemplatesOutgoingV2RequestItemDestinationsTransporters {
    #[serde(rename = "DriverLayoverLeg")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "EstimatedArrivalDateTime")]
    pub estimated_arrival_date_time: Option<String>,
    #[serde(rename = "EstimatedDepartureDateTime")]
    pub estimated_departure_date_time: Option<String>,
    #[serde(rename = "IsLayover")]
    pub is_layover: Option<bool>,
    #[serde(rename = "PhoneNumberForQuestions")]
    pub phone_number_for_questions: Option<String>,
    #[serde(rename = "TransporterDetails")]
    pub transporter_details: Option<Vec<TransfersUpdateTemplatesOutgoingV2RequestItemDestinationsTransportersTransporterDetails>>,
    #[serde(rename = "TransporterFacilityLicenseNumber")]
    pub transporter_facility_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TransfersUpdateTemplatesOutgoingV2RequestItemDestinationsTransportersTransporterDetails {
    #[serde(rename = "DriverLayoverLeg")]
    pub driver_layover_leg: Option<String>,
    #[serde(rename = "DriverLicenseNumber")]
    pub driver_license_number: Option<String>,
    #[serde(rename = "DriverName")]
    pub driver_name: Option<String>,
    #[serde(rename = "DriverOccupationalLicenseNumber")]
    pub driver_occupational_license_number: Option<String>,
    #[serde(rename = "VehicleLicensePlateNumber")]
    pub vehicle_license_plate_number: Option<String>,
    #[serde(rename = "VehicleMake")]
    pub vehicle_make: Option<String>,
    #[serde(rename = "VehicleModel")]
    pub vehicle_model: Option<String>,
}

