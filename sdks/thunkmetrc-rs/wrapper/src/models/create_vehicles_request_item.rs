use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateVehiclesRequestItem {
    #[serde(rename = "LicensePlateNumber", skip_serializing_if = "Option::is_none")]
    pub license_plate_number: Option<String>,
    #[serde(rename = "Make", skip_serializing_if = "Option::is_none")]
    pub make: Option<String>,
    #[serde(rename = "Model", skip_serializing_if = "Option::is_none")]
    pub model: Option<String>,
}
