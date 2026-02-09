use serde::{Deserialize, Serialize};
#[allow(unused_imports)]
use super::*;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct CreatePhotoRequestItem {
    #[serde(rename = "EncodedImageBase64", skip_serializing_if = "Option::is_none")]
    pub encoded_image_base64: Option<String>,
    #[serde(rename = "FileName", skip_serializing_if = "Option::is_none")]
    pub file_name: Option<String>,
}
