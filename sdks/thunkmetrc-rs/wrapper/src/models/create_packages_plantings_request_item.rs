use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreatePackagesPlantingsRequestItem {
    #[serde(rename = "LocationName", skip_serializing_if = "Option::is_none")]
    pub location_name: Option<String>,
    #[serde(rename = "PackageAdjustmentAmount", skip_serializing_if = "Option::is_none")]
    pub package_adjustment_amount: Option<i64>,
    #[serde(rename = "PackageAdjustmentUnitOfMeasureName", skip_serializing_if = "Option::is_none")]
    pub package_adjustment_unit_of_measure_name: Option<String>,
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "PatientLicenseNumber", skip_serializing_if = "Option::is_none")]
    pub patient_license_number: Option<String>,
    #[serde(rename = "PlantBatchName", skip_serializing_if = "Option::is_none")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "PlantBatchType", skip_serializing_if = "Option::is_none")]
    pub plant_batch_type: Option<String>,
    #[serde(rename = "PlantCount", skip_serializing_if = "Option::is_none")]
    pub plant_count: Option<i64>,
    #[serde(rename = "PlantedDate", skip_serializing_if = "Option::is_none")]
    pub planted_date: Option<String>,
    #[serde(rename = "StrainName", skip_serializing_if = "Option::is_none")]
    pub strain_name: Option<String>,
    #[serde(rename = "SublocationName", skip_serializing_if = "Option::is_none")]
    pub sublocation_name: Option<String>,
    #[serde(rename = "UnpackagedDate", skip_serializing_if = "Option::is_none")]
    pub unpackaged_date: Option<String>,
}
