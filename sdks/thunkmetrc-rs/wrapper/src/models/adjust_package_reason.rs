use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct AdjustPackageReason {
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "RequiresImmatureWasteWeight", skip_serializing_if = "Option::is_none")]
    pub requires_immature_waste_weight: Option<bool>,
    #[serde(rename = "RequiresMatureWasteWeight", skip_serializing_if = "Option::is_none")]
    pub requires_mature_waste_weight: Option<bool>,
    #[serde(rename = "RequiresNote", skip_serializing_if = "Option::is_none")]
    pub requires_note: Option<bool>,
    #[serde(rename = "RequiresWasteWeight", skip_serializing_if = "Option::is_none")]
    pub requires_waste_weight: Option<bool>,
}
