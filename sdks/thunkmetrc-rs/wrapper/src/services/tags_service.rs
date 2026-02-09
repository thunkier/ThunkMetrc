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

pub struct TagsService {
    client: MetrcClient,
    rate_limiter: Arc<MetrcRateLimiter>,
}

impl TagsService {
    pub fn new(client: MetrcClient, rate_limiter: Arc<MetrcRateLimiter>) -> Self {
        Self {
            client,
            rate_limiter,
        }
    }

    /// GET GetAvailablePackage
    /// Returns a list of available package tags. NOTE: This is a premium endpoint.
    /// Permissions Required:
    /// - Manage Tags
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_tags_available_package(&self, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<Vec<Tag>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.tags().get_tags_available_package(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Vec<Tag> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetAvailablePlant
    /// Returns a list of available plant tags. NOTE: This is a premium endpoint.
    /// Permissions Required:
    /// - Manage Tags
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_tags_available_plant(&self, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<Vec<Tag>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.tags().get_tags_available_plant(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Vec<Tag> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetStagedTags
    /// Returns a list of staged tags. NOTE: This is a premium endpoint.
    /// Permissions Required:
    /// - Manage Tags
    /// - RetailId.AllowPackageStaging Key Value enabled
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_staged_tags(&self, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<Vec<Staged>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.tags().get_staged_tags(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Vec<Staged> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }
}

