use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantsWaste {
    #[serde(rename = "GrowthPhase", skip_serializing_if = "Option::is_none")]
    pub growth_phase: Option<i64>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "LocationId", skip_serializing_if = "Option::is_none")]
    pub location_id: Option<i64>,
    #[serde(rename = "LocationName", skip_serializing_if = "Option::is_none")]
    pub location_name: Option<String>,
    #[serde(rename = "PlantId", skip_serializing_if = "Option::is_none")]
    pub plant_id: Option<i64>,
    #[serde(rename = "PlantWasteId", skip_serializing_if = "Option::is_none")]
    pub plant_waste_id: Option<i64>,
    #[serde(rename = "StateName", skip_serializing_if = "Option::is_none")]
    pub state_name: Option<String>,
    #[serde(rename = "SublocationId", skip_serializing_if = "Option::is_none")]
    pub sublocation_id: Option<i64>,
    #[serde(rename = "SublocationName", skip_serializing_if = "Option::is_none")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "WasteAmount", skip_serializing_if = "Option::is_none")]
    pub waste_amount: Option<i64>,
    #[serde(rename = "WasteUnitOfMeasureAbbreviation", skip_serializing_if = "Option::is_none")]
    pub waste_unit_of_measure_abbreviation: Option<String>,
}
