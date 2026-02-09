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

pub struct UnitsOfMeasureService {
    client: MetrcClient,
    rate_limiter: Arc<MetrcRateLimiter>,
}

impl UnitsOfMeasureService {
    pub fn new(client: MetrcClient, rate_limiter: Arc<MetrcRateLimiter>) -> Self {
        Self {
            client,
            rate_limiter,
        }
    }

    /// GET GetActiveUnitsOfMeasure
    /// Retrieves all active units of measure.
    /// Parameters:
    pub async fn get_active_units_of_measure(&self, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<UnitOfMeasure>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let body_val = body_val.clone();

                async move {
                    client.units_of_measure().get_active_units_of_measure(body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<UnitOfMeasure> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetInactiveUnitsOfMeasure
    /// Retrieves all inactive units of measure.
    /// Parameters:
    pub async fn get_inactive_units_of_measure(&self, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<UnitOfMeasure>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let body_val = body_val.clone();

                async move {
                    client.units_of_measure().get_inactive_units_of_measure(body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<UnitOfMeasure> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }
}

