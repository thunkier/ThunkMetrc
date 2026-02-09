use thunkmetrc_client::MetrcClient;
use serde_json::Value;
use std::error::Error;
use std::sync::Arc;
#[allow(unused_imports)]
use futures::Stream;
use crate::ratelimiter::MetrcRateLimiter;
#[allow(unused_imports)]
use crate::utils::iterate_pages;
#[allow(unused_imports)]
use crate::models::*;

pub struct RetailIdService {
    client: MetrcClient,
    rate_limiter: Arc<MetrcRateLimiter>,
}

impl RetailIdService {
    pub fn new(client: MetrcClient, rate_limiter: Arc<MetrcRateLimiter>) -> Self {
        Self {
            client,
            rate_limiter,
        }
    }

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
    pub async fn create_retail_id_associate(&self, license_number: Option<String>, body: Option<Vec<CreateAssociateRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.retail_id().create_retail_id_associate(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: WriteResponse = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
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
    pub async fn create_retail_id_generate(&self, license_number: Option<String>, body: Option<CreateGenerateRequest>) -> std::result::Result<Option<Vec<WriteResponse>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.retail_id().create_retail_id_generate(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Vec<WriteResponse> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
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
    pub async fn create_retail_id_merge(&self, license_number: Option<String>, body: Option<CreateMergeRequest>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.retail_id().create_retail_id_merge(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: WriteResponse = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
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
    pub async fn create_retail_id_packages_info(&self, license_number: Option<String>, body: Option<CreatePackagesInfoRequest>) -> std::result::Result<Option<Vec<WriteResponse>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.retail_id().create_retail_id_packages_info(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Vec<WriteResponse> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetAllotment
    /// Retrieves the available Retail Item ID quota for a facility.
    /// Permissions Required:
    /// - Download Product Labels
    /// - Manage Product Labels
    /// - Manage Tag Orders
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_retail_id_allotment(&self, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<Vec<Allotment>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.retail_id().get_retail_id_allotment(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Vec<Allotment> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
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
    pub async fn get_retail_id_receive_by_label(&self, label: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<Vec<Receive>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let label = label.to_string();
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let label = label.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.retail_id().get_retail_id_receive_by_label(&label, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Vec<Receive> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
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
    pub async fn get_retail_id_receive_qr_by_short_code(&self, short_code: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<Vec<ReceiveQrByShortCode>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let short_code = short_code.to_string();
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let short_code = short_code.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.retail_id().get_retail_id_receive_qr_by_short_code(&short_code, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Vec<ReceiveQrByShortCode> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }
}

