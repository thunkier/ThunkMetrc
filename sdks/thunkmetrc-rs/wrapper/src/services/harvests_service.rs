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

pub struct HarvestsService {
    client: MetrcClient,
    rate_limiter: Arc<MetrcRateLimiter>,
}

impl HarvestsService {
    pub fn new(client: MetrcClient, rate_limiter: Arc<MetrcRateLimiter>) -> Self {
        Self {
            client,
            rate_limiter,
        }
    }

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
    pub async fn create_harvests_packages(&self, license_number: Option<String>, body: Option<Vec<CreateHarvestsPackagesRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().create_harvests_packages(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: WriteResponse = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// POST CreateHarvestsWaste
    /// Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
    /// Permissions Required:
    /// - View Harvests
    /// - Manage Harvests
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_harvests_waste(&self, license_number: Option<String>, body: Option<Vec<CreateHarvestsWasteRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().create_harvests_waste(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: WriteResponse = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
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
    pub async fn create_harvests_packages_testing(&self, license_number: Option<String>, body: Option<Vec<CreatePackagesTestingRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().create_harvests_packages_testing(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: WriteResponse = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// DELETE DeleteWasteById
    /// Discontinues a specific harvest waste record by Id for the specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Discontinue Harvest Waste
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_harvests_waste_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let id = id.to_string();
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let id = id.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().delete_harvests_waste_by_id(&id, license_number, body_val.as_ref()
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

    /// PUT FinishHarvests
    /// Marks one or more harvests as finished for the specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Finish/Discontinue Harvests
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn finish_harvests_harvests(&self, license_number: Option<String>, body: Option<Vec<FinishHarvestsRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().finish_harvests_harvests(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: WriteResponse = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
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
    pub async fn get_active_harvests(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<Harvest>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let last_modified_end = last_modified_end.clone();
                let last_modified_start = last_modified_start.clone();
                let license_number = license_number.clone();
                let page_number = page_number.clone();
                let page_size = page_size.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().get_active_harvests(last_modified_end, last_modified_start, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<Harvest> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetHarvestsById
    /// Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
    /// Permissions Required:
    /// - View Harvests
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_harvests_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<Harvest>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let id = id.to_string();
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let id = id.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().get_harvests_by_id(&id, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Harvest = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
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
    pub async fn get_harvests_waste(&self, harvest_id: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<HarvestsWaste>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let harvest_id = harvest_id.clone();
                let license_number = license_number.clone();
                let page_number = page_number.clone();
                let page_size = page_size.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().get_harvests_waste(harvest_id, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<HarvestsWaste> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
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
    pub async fn get_inactive_harvests(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<Harvest>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let last_modified_end = last_modified_end.clone();
                let last_modified_start = last_modified_start.clone();
                let license_number = license_number.clone();
                let page_number = page_number.clone();
                let page_size = page_size.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().get_inactive_harvests(last_modified_end, last_modified_start, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<Harvest> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
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
    pub async fn get_on_hold_harvests(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<Harvest>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let last_modified_end = last_modified_end.clone();
                let last_modified_start = last_modified_start.clone();
                let license_number = license_number.clone();
                let page_number = page_number.clone();
                let page_size = page_size.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().get_on_hold_harvests(last_modified_end, last_modified_start, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<Harvest> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetWasteTypes
    /// Retrieves a list of Waste types for harvests.
    /// Parameters:
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_harvests_waste_types(&self, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<WasteType>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let page_number = page_number.clone();
                let page_size = page_size.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().get_harvests_waste_types(page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<WasteType> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// PUT UnfinishHarvests
    /// Reopens one or more previously finished harvests for the specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Finish/Discontinue Harvests
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn unfinish_harvests_harvests(&self, license_number: Option<String>, body: Option<Vec<UnfinishHarvestsRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().unfinish_harvests_harvests(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: WriteResponse = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// PUT UpdateHarvestsLocation
    /// Updates the Location of Harvest for a specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Manage Harvests
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_harvests_location(&self, license_number: Option<String>, body: Option<Vec<UpdateHarvestsLocationRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().update_harvests_location(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: WriteResponse = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// PUT UpdateRename
    /// Renames one or more harvests for the specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Manage Harvests
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_harvests_rename(&self, license_number: Option<String>, body: Option<Vec<UpdateRenameRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().update_harvests_rename(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: WriteResponse = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// PUT UpdateRestoreHarvestedPlants
    /// Restores previously harvested plants to their original state for the specified Facility.
    /// Permissions Required:
    /// - View Harvests
    /// - Finish/Discontinue Harvests
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_harvests_restore_harvested_plants(&self, license_number: Option<String>, body: Option<Vec<UpdateRestoreHarvestedPlantsRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.harvests().update_harvests_restore_harvested_plants(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: WriteResponse = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }
}

