use serde::{Deserialize, Serialize};
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct AdditivesTemplatesAdditivesTemplateActiveIngredientsItem {
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "Percentage", skip_serializing_if = "Option::is_none")]
    pub percentage: Option<f64>,
}
