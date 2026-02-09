use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct TransportersCreateTransporterVehiclesRequestItem {
    #[serde(rename = "license_plate_number", skip_serializing_if = "Option::is_none")]
    pub license_plate_number: Option<String>,
    #[serde(rename = "make", skip_serializing_if = "Option::is_none")]
    pub make: Option<String>,
    #[serde(rename = "model", skip_serializing_if = "Option::is_none")]
    pub model: Option<String>,
}
