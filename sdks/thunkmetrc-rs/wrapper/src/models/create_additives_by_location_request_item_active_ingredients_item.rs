use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreateAdditivesByLocationRequestItemActiveIngredientsItem {
    #[serde(rename = "Name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "Percentage", skip_serializing_if = "Option::is_none")]
    pub percentage: Option<i64>,
}
