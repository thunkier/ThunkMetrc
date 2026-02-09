use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct UpdateUseByDateRequest {
    #[serde(rename = "ExpirationDate", skip_serializing_if = "Option::is_none")]
    pub expiration_date: Option<String>,
    #[serde(rename = "Label", skip_serializing_if = "Option::is_none")]
    pub label: Option<String>,
    #[serde(rename = "SellByDate", skip_serializing_if = "Option::is_none")]
    pub sell_by_date: Option<String>,
    #[serde(rename = "UseByDate", skip_serializing_if = "Option::is_none")]
    pub use_by_date: Option<String>,
}
