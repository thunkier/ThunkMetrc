use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct SourceHarvest {
    #[serde(rename = "HarvestStartDate", skip_serializing_if = "Option::is_none")]
    pub harvest_start_date: Option<String>,
    #[serde(rename = "HarvestedByFacilityLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub harvested_by_facility_license_number: Option<String>,
    #[serde(rename = "HarvestedByFacilityName", skip_serializing_if = "Option::is_none")]
    pub harvested_by_facility_name: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
}
