use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateOutgoingTemplatesRequestItemDestinationsItemPackagesItem {
    #[serde(rename = "GrossUnitOfWeightName", skip_serializing_if = "Option::is_none")]
    pub gross_unit_of_weight_name: Option<String>,
    #[serde(rename = "GrossWeight", skip_serializing_if = "Option::is_none")]
    pub gross_weight: Option<f64>,
    #[serde(rename = "PackageLabel", skip_serializing_if = "Option::is_none")]
    pub package_label: Option<String>,
    #[serde(rename = "WholesalePrice", skip_serializing_if = "Option::is_none")]
    pub wholesale_price: Option<String>,
}
