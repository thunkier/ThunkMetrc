use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct AdditivesTemplatesUpdateAdditivesTemplatesRequestItemActiveIngredientsItem {
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "percentage", skip_serializing_if = "Option::is_none")]
    pub percentage: Option<f64>,
}
