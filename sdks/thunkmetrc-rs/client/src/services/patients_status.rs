use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct PatientsStatusClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> PatientsStatusClient<'a> {
    /// GET GetPatientsStatusesByPatientLicenseNumber
    /// Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
    /// Permissions Required:
    /// - Lookup Patients
    /// Parameters:
    /// - patient_license_number (str): Path parameter patientLicenseNumber
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_patients_statuses_by_patient_license_number(&self, patient_license_number: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/patients/v2/statuses/{}", urlencoding::encode(patient_license_number).as_ref());
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
}

