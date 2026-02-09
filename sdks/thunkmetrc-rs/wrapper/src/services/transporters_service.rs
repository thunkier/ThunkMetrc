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

pub struct TransportersService {
    client: MetrcClient,
    rate_limiter: Arc<MetrcRateLimiter>,
}

impl TransportersService {
    pub fn new(client: MetrcClient, rate_limiter: Arc<MetrcRateLimiter>) -> Self {
        Self {
            client,
            rate_limiter,
        }
    }

    /// POST CreateDrivers
    /// Creates new driver records for a Facility.
    /// Permissions Required:
    /// - Manage Transporters
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_transporters_drivers(&self, license_number: Option<String>, body: Option<Vec<CreateDriversRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.transporters().create_transporters_drivers(license_number, body_val.as_ref()
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

    /// POST CreateVehicles
    /// Creates new vehicle records for a Facility.
    /// Permissions Required:
    /// - Manage Transporters
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_transporters_vehicles(&self, license_number: Option<String>, body: Option<Vec<CreateVehiclesRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.transporters().create_transporters_vehicles(license_number, body_val.as_ref()
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

    /// DELETE DeleteDriverById
    /// Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
    /// Permissions Required:
    /// - Manage Transporters
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_transporters_driver_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
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
                    client.transporters().delete_transporters_driver_by_id(&id, license_number, body_val.as_ref()
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

    /// DELETE DeleteVehicleById
    /// Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
    /// Permissions Required:
    /// - Manage Transporters
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_transporters_vehicle_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
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
                    client.transporters().delete_transporters_vehicle_by_id(&id, license_number, body_val.as_ref()
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

    /// GET GetDriverById
    /// Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
    /// Permissions Required:
    /// - Transporters
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_transporters_driver_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<TransportersDriver>, Box<dyn Error + Send + Sync>> {
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
                    client.transporters().get_transporters_driver_by_id(&id, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: TransportersDriver = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetDrivers
    /// Retrieves a list of drivers for a Facility.
    /// Permissions Required:
    /// - Transporters
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_transporters_drivers(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<TransportersDriver>>, Box<dyn Error + Send + Sync>> {
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
                    client.transporters().get_transporters_drivers(license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<TransportersDriver> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetVehicleById
    /// Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
    /// Permissions Required:
    /// - Transporters
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_transporters_vehicle_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<TransportersVehicle>, Box<dyn Error + Send + Sync>> {
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
                    client.transporters().get_transporters_vehicle_by_id(&id, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: TransportersVehicle = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetVehicles
    /// Retrieves a list of vehicles for a Facility.
    /// Permissions Required:
    /// - Transporters
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_transporters_vehicles(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<TransportersVehicle>>, Box<dyn Error + Send + Sync>> {
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
                    client.transporters().get_transporters_vehicles(license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<TransportersVehicle> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// PUT UpdateDrivers
    /// Updates existing driver records for a Facility.
    /// Permissions Required:
    /// - Manage Transporters
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_transporters_drivers(&self, license_number: Option<String>, body: Option<Vec<UpdateDriversRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.transporters().update_transporters_drivers(license_number, body_val.as_ref()
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

    /// PUT UpdateVehicles
    /// Updates existing vehicle records for a facility.
    /// Permissions Required:
    /// - Manage Transporters
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_transporters_vehicles(&self, license_number: Option<String>, body: Option<Vec<UpdateVehiclesRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.transporters().update_transporters_vehicles(license_number, body_val.as_ref()
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

