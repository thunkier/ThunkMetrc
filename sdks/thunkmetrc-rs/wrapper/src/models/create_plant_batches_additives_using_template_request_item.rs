use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreatePlantBatchesAdditivesUsingTemplateRequestItem {
    #[serde(rename = "ActualDate", skip_serializing_if = "Option::is_none")]
    pub actual_date: Option<String>,
    #[serde(rename = "AdditivesTemplateName", skip_serializing_if = "Option::is_none")]
    pub additives_template_name: Option<String>,
    #[serde(rename = "PlantBatchName", skip_serializing_if = "Option::is_none")]
    pub plant_batch_name: Option<String>,
    #[serde(rename = "Rate", skip_serializing_if = "Option::is_none")]
    pub rate: Option<String>,
    #[serde(rename = "TotalAmountApplied", skip_serializing_if = "Option::is_none")]
    pub total_amount_applied: Option<i64>,
    #[serde(rename = "TotalAmountUnitOfMeasure", skip_serializing_if = "Option::is_none")]
    pub total_amount_unit_of_measure: Option<String>,
    #[serde(rename = "Volume", skip_serializing_if = "Option::is_none")]
    pub volume: Option<String>,
}
