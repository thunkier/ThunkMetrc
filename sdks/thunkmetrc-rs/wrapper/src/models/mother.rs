use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Mother {
    #[serde(rename = "ClonedCount", skip_serializing_if = "Option::is_none")]
    pub cloned_count: Option<i64>,
    #[serde(rename = "DescendedCount", skip_serializing_if = "Option::is_none")]
    pub descended_count: Option<i64>,
    #[serde(rename = "DestroyedByUserName", skip_serializing_if = "Option::is_none")]
    pub destroyed_by_user_name: Option<String>,
    #[serde(rename = "DestroyedDate", skip_serializing_if = "Option::is_none")]
    pub destroyed_date: Option<String>,
    #[serde(rename = "DestroyedNote", skip_serializing_if = "Option::is_none")]
    pub destroyed_note: Option<String>,
    #[serde(rename = "FloweringDate", skip_serializing_if = "Option::is_none")]
    pub flowering_date: Option<String>,
    #[serde(rename = "GroupTagTypeMax", skip_serializing_if = "Option::is_none")]
    pub group_tag_type_max: Option<i64>,
    #[serde(rename = "GrowthPhase", skip_serializing_if = "Option::is_none")]
    pub growth_phase: Option<String>,
    #[serde(rename = "HarvestCount", skip_serializing_if = "Option::is_none")]
    pub harvest_count: Option<i64>,
    #[serde(rename = "HarvestId", skip_serializing_if = "Option::is_none")]
    pub harvest_id: Option<i64>,
    #[serde(rename = "HarvestedDate", skip_serializing_if = "Option::is_none")]
    pub harvested_date: Option<String>,
    #[serde(rename = "HarvestedUnitOfWeightAbbreviation", skip_serializing_if = "Option::is_none")]
    pub harvested_unit_of_weight_abbreviation: Option<String>,
    #[serde(rename = "HarvestedUnitOfWeightName", skip_serializing_if = "Option::is_none")]
    pub harvested_unit_of_weight_name: Option<String>,
    #[serde(rename = "HarvestedWetWeight", skip_serializing_if = "Option::is_none")]
    pub harvested_wet_weight: Option<f64>,
    #[serde(rename = "Id", skip_serializing_if = "Option::is_none")]
    pub id: Option<i64>,
    #[serde(rename = "IsOnHold", skip_serializing_if = "Option::is_none")]
    pub is_on_hold: Option<bool>,
    #[serde(rename = "IsOnInvestigation", skip_serializing_if = "Option::is_none")]
    pub is_on_investigation: Option<bool>,
    #[serde(rename = "IsOnInvestigationHold", skip_serializing_if = "Option::is_none")]
    pub is_on_investigation_hold: Option<bool>,
    #[serde(rename = "IsOnInvestigationRecall", skip_serializing_if = "Option::is_none")]
    pub is_on_investigation_recall: Option<bool>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "LocationId", skip_serializing_if = "Option::is_none")]
    pub location_id: Option<i64>,
    #[serde(rename = "LocationName", skip_serializing_if = "Option::is_none")]
    pub location_name: Option<String>,
    #[serde(rename = "LocationTypeName", skip_serializing_if = "Option::is_none")]
    pub location_type_name: Option<String>,
    #[serde(rename = "MotherPlantDate", skip_serializing_if = "Option::is_none")]
    pub mother_plant_date: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatchId", skip_serializing_if = "Option::is_none")]
    pub plant_batch_id: Option<i64>,
    #[serde(rename = "PlantBatchName", skip_serializing_if = "Option::is_none")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "PlantBatchTypeId", skip_serializing_if = "Option::is_none")]
    pub plant_batch_type_id: Option<i64>,
    #[serde(rename = "PlantBatchTypeName", skip_serializing_if = "Option::is_none")]
    pub plant_batch_type_name: Option<String>,
    #[serde(rename = "PlantedDate", skip_serializing_if = "Option::is_none")]
    pub planted_date: Option<String>,
    #[serde(rename = "State", skip_serializing_if = "Option::is_none")]
    pub state: Option<String>,
    #[serde(rename = "StrainId", skip_serializing_if = "Option::is_none")]
    pub strain_id: Option<i64>,
    #[serde(rename = "StrainName", skip_serializing_if = "Option::is_none")]
    pub strain_name: Option<String>,
    #[serde(rename = "SublocationId", skip_serializing_if = "Option::is_none")]
    pub sublocation_id: Option<i64>,
    #[serde(rename = "SublocationName", skip_serializing_if = "Option::is_none")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "SurvivedCount", skip_serializing_if = "Option::is_none")]
    pub survived_count: Option<i64>,
    #[serde(rename = "TagTypeMax", skip_serializing_if = "Option::is_none")]
    pub tag_type_max: Option<i64>,
    #[serde(rename = "VegetativeDate", skip_serializing_if = "Option::is_none")]
    pub vegetative_date: Option<String>,
}
