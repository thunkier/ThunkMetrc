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

pub struct CaregiversStatusService {
    client: MetrcClient,
    rate_limiter: Arc<MetrcRateLimiter>,
}

impl CaregiversStatusService {
    pub fn new(client: MetrcClient, rate_limiter: Arc<MetrcRateLimiter>) -> Self {
        Self {
            client,
            rate_limiter,
        }
    }

    /// GET GetCaregiversStatusByCaregiverLicenseNumber
    /// Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
    /// Permissions Required:
    /// - Lookup Caregivers
    /// Parameters:
    /// - caregiver_license_number (str): Path parameter caregiverLicenseNumber
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_caregivers_status_by_caregiver_license_number(&self, caregiver_license_number: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<Vec<CaregiversStatus>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let caregiver_license_number = caregiver_license_number.to_string();
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let caregiver_license_number = caregiver_license_number.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.caregivers_status().get_caregivers_status_by_caregiver_license_number(&caregiver_license_number, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Vec<CaregiversStatus> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }
}

