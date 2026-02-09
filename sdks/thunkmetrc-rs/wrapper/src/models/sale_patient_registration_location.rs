use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SalePatientRegistrationLocation {
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
}
