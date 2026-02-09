use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct CaregiversStatusClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> CaregiversStatusClient<'a> {
    /// GET GetCaregiversStatusByCaregiverLicenseNumber
    /// Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
    /// Permissions Required:
    /// - Lookup Caregivers
    /// Parameters:
    /// - caregiver_license_number (str): Path parameter caregiverLicenseNumber
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_caregivers_status_by_caregiver_license_number(&self, caregiver_license_number: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/caregivers/v2/status/{}", urlencoding::encode(caregiver_license_number).as_ref());
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

