use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateLocationsRequestItem {
    #[serde(rename = "LocationTypeName", skip_serializing_if = "Option::is_none")]
    pub location_type_name: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
}
