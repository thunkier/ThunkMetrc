use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Harvest {
    #[serde(rename = "ArchivedDate", skip_serializing_if = "Option::is_none")]
    pub archived_date: Option<String>,
    #[serde(rename = "CurrentWeight", skip_serializing_if = "Option::is_none")]
    pub current_weight: Option<f64>,
    #[serde(rename = "DryingLocationId", skip_serializing_if = "Option::is_none")]
    pub drying_location_id: Option<i64>,
    #[serde(rename = "DryingLocationName", skip_serializing_if = "Option::is_none")]
    pub drying_location_name: Option<String>,
    #[serde(rename = "DryingLocationTypeName", skip_serializing_if = "Option::is_none")]
    pub drying_location_type_name: Option<String>,
    #[serde(rename = "DryingSublocationId", skip_serializing_if = "Option::is_none")]
    pub drying_sublocation_id: Option<i64>,
    #[serde(rename = "DryingSublocationName", skip_serializing_if = "Option::is_none")]
    pub drying_sublocation_name: Option<String>,
    #[serde(rename = "FinishedDate", skip_serializing_if = "Option::is_none")]
    pub finished_date: Option<String>,
    #[serde(rename = "HarvestStartDate", skip_serializing_if = "Option::is_none")]
    pub harvest_start_date: Option<String>,
    #[serde(rename = "HarvestType", skip_serializing_if = "Option::is_none")]
    pub harvest_type: Option<String>,
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
    #[serde(rename = "LabTestingState", skip_serializing_if = "Option::is_none")]
    pub lab_testing_state: Option<String>,
    #[serde(rename = "LabTestingStateDate", skip_serializing_if = "Option::is_none")]
    pub lab_testing_state_date: Option<String>,
    #[serde(rename = "LastModified", skip_serializing_if = "Option::is_none")]
    pub last_modified: Option<String>,
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "PackageCount", skip_serializing_if = "Option::is_none")]
    pub package_count: Option<i64>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantCount", skip_serializing_if = "Option::is_none")]
    pub plant_count: Option<i64>,
    #[serde(rename = "SourceStrainCount", skip_serializing_if = "Option::is_none")]
    pub source_strain_count: Option<i64>,
    #[serde(rename = "SourceStrainNames", skip_serializing_if = "Option::is_none")]
    pub source_strain_names: Option<String>,
    #[serde(rename = "Strains", skip_serializing_if = "Option::is_none")]
    pub strains: Option<Vec<serde_json::Value>>,
    #[serde(rename = "TotalPackagedWeight", skip_serializing_if = "Option::is_none")]
    pub total_packaged_weight: Option<f64>,
    #[serde(rename = "TotalRestoredWeight", skip_serializing_if = "Option::is_none")]
    pub total_restored_weight: Option<f64>,
    #[serde(rename = "TotalWasteWeight", skip_serializing_if = "Option::is_none")]
    pub total_waste_weight: Option<f64>,
    #[serde(rename = "TotalWetWeight", skip_serializing_if = "Option::is_none")]
    pub total_wet_weight: Option<f64>,
    #[serde(rename = "UnitOfWeightName", skip_serializing_if = "Option::is_none")]
    pub unit_of_weight_name: Option<String>,
}
