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

pub struct PatientsStatusService {
    client: MetrcClient,
    rate_limiter: Arc<MetrcRateLimiter>,
}

impl PatientsStatusService {
    pub fn new(client: MetrcClient, rate_limiter: Arc<MetrcRateLimiter>) -> Self {
        Self {
            client,
            rate_limiter,
        }
    }

    /// GET GetPatientsStatusesByPatientLicenseNumber
    /// Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
    /// Permissions Required:
    /// - Lookup Patients
    /// Parameters:
    /// - patient_license_number (str): Path parameter patientLicenseNumber
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_patients_statuses_by_patient_license_number(&self, patient_license_number: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<Vec<PatientsStatus>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let patient_license_number = patient_license_number.to_string();
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let patient_license_number = patient_license_number.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.patients_status().get_patients_statuses_by_patient_license_number(&patient_license_number, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Vec<PatientsStatus> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }
}

