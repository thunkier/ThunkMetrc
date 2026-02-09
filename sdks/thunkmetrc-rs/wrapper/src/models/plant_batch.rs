use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct PlantBatch {
    #[serde(rename = "DestroyedCount", skip_serializing_if = "Option::is_none")]
    pub destroyed_count: Option<i64>,
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
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "LocationId", skip_serializing_if = "Option::is_none")]
    pub location_id: Option<i64>,
    #[serde(rename = "LocationName", skip_serializing_if = "Option::is_none")]
    pub location_name: Option<String>,
    #[serde(rename = "LocationTypeName", skip_serializing_if = "Option::is_none")]
    pub location_type_name: Option<String>,
    #[serde(rename = "MultiPlantBatch", skip_serializing_if = "Option::is_none")]
    pub multi_plant_batch: Option<bool>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "PackagedCount", skip_serializing_if = "Option::is_none")]
    pub packaged_count: Option<i64>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatchTypeId", skip_serializing_if = "Option::is_none")]
    pub plant_batch_type_id: Option<i64>,
    #[serde(rename = "PlantBatchTypeName", skip_serializing_if = "Option::is_none")]
    pub plant_batch_type_name: Option<String>,
    #[serde(rename = "PlantedDate", skip_serializing_if = "Option::is_none")]
    pub planted_date: Option<String>,
    #[serde(rename = "SourcePackageId", skip_serializing_if = "Option::is_none")]
    pub source_package_id: Option<i64>,
    #[serde(rename = "SourcePackageLabel", skip_serializing_if = "Option::is_none")]
    pub source_package_label: Option<String>,
    #[serde(rename = "SourcePlantBatchIds", skip_serializing_if = "Option::is_none")]
    pub source_plant_batch_ids: Option<Vec<serde_json::Value>>,
    #[serde(rename = "SourcePlantBatchNames", skip_serializing_if = "Option::is_none")]
    pub source_plant_batch_names: Option<String>,
    #[serde(rename = "SourcePlantId", skip_serializing_if = "Option::is_none")]
    pub source_plant_id: Option<i64>,
    #[serde(rename = "SourcePlantLabel", skip_serializing_if = "Option::is_none")]
    pub source_plant_label: Option<String>,
    #[serde(rename = "StrainId", skip_serializing_if = "Option::is_none")]
    pub strain_id: Option<i64>,
    #[serde(rename = "StrainName", skip_serializing_if = "Option::is_none")]
    pub strain_name: Option<String>,
    #[serde(rename = "SublocationId", skip_serializing_if = "Option::is_none")]
    pub sublocation_id: Option<i64>,
    #[serde(rename = "SublocationName", skip_serializing_if = "Option::is_none")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "TrackedCount", skip_serializing_if = "Option::is_none")]
    pub tracked_count: Option<i64>,
    #[serde(rename = "UntrackedCount", skip_serializing_if = "Option::is_none")]
    pub untracked_count: Option<i64>,
}
