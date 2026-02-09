use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct PatientCheckInsClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> PatientCheckInsClient<'a> {
    /// POST CreatePatientCheckIns
    /// Records patient check-ins for a specified Facility.
    /// Permissions Required:
    /// - ManagePatientsCheckIns
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_patient_check_ins(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/patient-checkins/v2");
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
    /// DELETE DeletePatientCheckInsById
    /// Archives a Patient Check-In, identified by its Id, for a specified Facility.
    /// Permissions Required:
    /// - ManagePatientsCheckIns
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_patient_check_ins_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/patient-checkins/v2/{}", urlencoding::encode(id).as_ref());
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
    /// GET GetLocations
    /// Retrieves a list of Patient Check-In locations.
    /// Parameters:
    pub async fn get_patient_check_ins_locations(&self, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let path = format!("/patient-checkins/v2/locations");

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetPatientCheckIns
    /// Retrieves a list of patient check-ins for a specified Facility.
    /// Permissions Required:
    /// - ManagePatientsCheckIns
    /// Parameters:
    /// - checkin_date_end (Option<String>): Filter by checkinDateEnd
    /// - checkin_date_start (Option<String>): Filter by checkinDateStart
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_patient_check_ins(&self, checkin_date_end: Option<String>, checkin_date_start: Option<String>, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/patient-checkins/v2");
        let mut query_params = Vec::new();
        if let Some(p) = checkin_date_end {
            query_params.push(format!("checkinDateEnd={}", urlencoding::encode(&p)));
        }
        if let Some(p) = checkin_date_start {
            query_params.push(format!("checkinDateStart={}", urlencoding::encode(&p)));
        }
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// PUT UpdatePatientCheckIns
    /// Updates patient check-ins for a specified Facility.
    /// Permissions Required:
    /// - ManagePatientsCheckIns
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_patient_check_ins(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/patient-checkins/v2");
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

