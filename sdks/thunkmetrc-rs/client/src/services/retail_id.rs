use crate::client::MetrcClient;
use serde_json::Value;
use std::error::Error;

pub struct RetailIdClient<'a> {
    pub(crate) client: &'a MetrcClient,
}

impl<'a> RetailIdClient<'a> {
    /// POST CreateAssociate
    /// Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
    /// - WebApi Retail ID Read Write State (All or WriteOnly)
    /// - Industry/View Packages
    /// - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_retail_id_associate(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/retailid/v2/associate");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// POST CreateGenerate
    /// Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
    /// - WebApi Retail ID Read Write State (All or WriteOnly)
    /// - Industry/View Packages
    /// - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_retail_id_generate(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/retailid/v2/generate");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// POST CreateMerge
    /// Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
    /// - WebApi Retail ID Read Write State (All or WriteOnly)
    /// - Key Value Settings/Retail ID Merge Packages Enabled
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_retail_id_merge(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/retailid/v2/merge");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// POST CreatePackagesInfo
    /// Retrieves Package information for given list of Package labels.
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
    /// - WebApi Retail ID Read Write State (All or WriteOnly)
    /// - Industry/View Packages
    /// - Admin/Employees/Packages Page/Product Labels(Manage)
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_retail_id_packages_info(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/retailid/v2/packages/info");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::POST, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetAllotment
    /// Retrieves the available Retail Item ID quota for a facility.
    /// Permissions Required:
    /// - Download Product Labels
    /// - Manage Product Labels
    /// - Manage Tag Orders
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_retail_id_allotment(&self, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/retailid/v2/allotment");
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetReceiveByLabel
    /// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Manage RetailId
    /// - WebApi Retail ID Read Write State (All or ReadOnly)
    /// - Industry/View Packages
    /// - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
    /// Parameters:
    /// - label (str): Path parameter label
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_retail_id_receive_by_label(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/retailid/v2/receive/{}", urlencoding::encode(label).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
    /// GET GetReceiveQrByShortCode
    /// Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Manage RetailId
    /// - WebApi Retail ID Read Write State (All or ReadOnly)
    /// - Industry/View Packages
    /// - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
    /// Parameters:
    /// - short_code (str): Path parameter shortCode
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_retail_id_receive_qr_by_short_code(&self, short_code: &str, license_number: Option<String>, body: Option<&Value>) -> Result<Option<Value>, Box<dyn Error + Send + Sync>> {
        let mut path = format!("/retailid/v2/receive/qr/{}", urlencoding::encode(short_code).as_ref());
        let mut query_params = Vec::new();
        if let Some(p) = license_number {
            query_params.push(format!("licenseNumber={}", urlencoding::encode(&p)));
        }
        if !query_params.is_empty() {
            path.push_str("?");
            path.push_str(&query_params.join("&"));
        }

        self.client.send(reqwest::Method::GET, &path, body.map(|b| serde_json::to_value(b).unwrap()).as_ref()).await
    }
}

