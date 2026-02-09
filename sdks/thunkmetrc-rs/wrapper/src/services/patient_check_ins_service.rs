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

pub struct PatientCheckInsService {
    client: MetrcClient,
    rate_limiter: Arc<MetrcRateLimiter>,
}

impl PatientCheckInsService {
    pub fn new(client: MetrcClient, rate_limiter: Arc<MetrcRateLimiter>) -> Self {
        Self {
            client,
            rate_limiter,
        }
    }

    /// POST CreatePatientCheckIns
    /// Records patient check-ins for a specified Facility.
    /// Permissions Required:
    /// - ManagePatientsCheckIns
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_patient_check_ins(&self, license_number: Option<String>, body: Option<Vec<CreatePatientCheckInsRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.patient_check_ins().create_patient_check_ins(license_number, body_val.as_ref()
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

    /// DELETE DeletePatientCheckInsById
    /// Archives a Patient Check-In, identified by its Id, for a specified Facility.
    /// Permissions Required:
    /// - ManagePatientsCheckIns
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_patient_check_ins_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
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
                    client.patient_check_ins().delete_patient_check_ins_by_id(&id, license_number, body_val.as_ref()
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

    /// GET GetLocations
    /// Retrieves a list of Patient Check-In locations.
    /// Parameters:
    pub async fn get_patient_check_ins_locations(&self, body: Option<&Value>) -> std::result::Result<Option<Vec<PatientCheckInsLocation>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let body_val = body_val.clone();

                async move {
                    client.patient_check_ins().get_patient_check_ins_locations(body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Vec<PatientCheckInsLocation> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetPatientCheckIns
    /// Retrieves a list of patient check-ins for a specified Facility.
    /// Permissions Required:
    /// - ManagePatientsCheckIns
    /// Parameters:
    /// - checkin_date_end (Option<String>): Filter by checkinDateEnd
    /// - checkin_date_start (Option<String>): Filter by checkinDateStart
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_patient_check_ins(&self, checkin_date_end: Option<String>, checkin_date_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<PatientCheckIn>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let checkin_date_end = checkin_date_end.clone();
                let checkin_date_start = checkin_date_start.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.patient_check_ins().get_patient_check_ins(checkin_date_end, checkin_date_start, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<PatientCheckIn> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// PUT UpdatePatientCheckIns
    /// Updates patient check-ins for a specified Facility.
    /// Permissions Required:
    /// - ManagePatientsCheckIns
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_patient_check_ins(&self, license_number: Option<String>, body: Option<Vec<UpdatePatientCheckInsRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.patient_check_ins().update_patient_check_ins(license_number, body_val.as_ref()
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

