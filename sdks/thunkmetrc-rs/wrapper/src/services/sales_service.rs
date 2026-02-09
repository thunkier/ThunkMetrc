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

pub struct SalesService {
    client: MetrcClient,
    rate_limiter: Arc<MetrcRateLimiter>,
}

impl SalesService {
    pub fn new(client: MetrcClient, rate_limiter: Arc<MetrcRateLimiter>) -> Self {
        Self {
            client,
            rate_limiter,
        }
    }

    /// POST CreateDeliveries
    /// Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    /// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_sales_deliveries(&self, license_number: Option<String>, body: Option<Vec<CreateDeliveriesRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().create_sales_deliveries(license_number, body_val.as_ref()
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

    /// POST CreateDeliveriesRetailerDepart
    /// Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    /// - Industry/Facility Type/Retailer Delivery
    /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    /// - Manage Retailer Delivery
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_sales_deliveries_retailer_depart(&self, license_number: Option<String>, body: Option<Vec<CreateDeliveriesRetailerDepartRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().create_sales_deliveries_retailer_depart(license_number, body_val.as_ref()
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

    /// POST CreateDeliveriesRetailerEnd
    /// Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    /// - Industry/Facility Type/Retailer Delivery
    /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    /// - Manage Retailer Delivery
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_sales_deliveries_retailer_end(&self, license_number: Option<String>, body: Option<Vec<CreateDeliveriesRetailerEndRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().create_sales_deliveries_retailer_end(license_number, body_val.as_ref()
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

    /// POST CreateDeliveriesRetailerRestock
    /// Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    /// - Industry/Facility Type/Retailer Delivery
    /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    /// - Manage Retailer Delivery
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_sales_deliveries_retailer_restock(&self, license_number: Option<String>, body: Option<Vec<CreateDeliveriesRetailerRestockRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().create_sales_deliveries_retailer_restock(license_number, body_val.as_ref()
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

    /// POST CreateReceipts
    /// Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Sales (Write)
    /// - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
    /// - Industry/Facility Type/Advanced Sales
    /// - WebApi Sales Read Write State (All or WriteOnly)
    /// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_sales_receipts(&self, license_number: Option<String>, body: Option<Vec<CreateReceiptsRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().create_sales_receipts(license_number, body_val.as_ref()
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

    /// POST CreateSalesDeliveriesRetailer
    /// Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    /// - Industry/Facility Type/Retailer Delivery
    /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    /// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    /// - Manage Retailer Delivery
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_sales_deliveries_retailer(&self, license_number: Option<String>, body: Option<Vec<CreateSalesDeliveriesRetailerRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().create_sales_deliveries_retailer(license_number, body_val.as_ref()
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

    /// DELETE DeleteDeliveryById
    /// Voids a sales delivery for a Facility using the provided License Number and delivery Id.
    /// Permissions Required:
    /// - Manage Sales Delivery
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_sales_delivery_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let id = id.to_string();
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let id = id.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().delete_sales_delivery_by_id(&id, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: serde_json::Value = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// DELETE DeleteDeliveryRetailerById
    /// Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    /// - Industry/Facility Type/Retailer Delivery
    /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    /// - Manage Retailer Delivery
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_sales_delivery_retailer_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let id = id.to_string();
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let id = id.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().delete_sales_delivery_retailer_by_id(&id, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: serde_json::Value = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// DELETE DeleteReceiptById
    /// Archives a sales receipt for a Facility using the provided License Number and receipt Id.
    /// Permissions Required:
    /// - Manage Sales
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_sales_receipt_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let id = id.to_string();
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let id = id.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().delete_sales_receipt_by_id(&id, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: serde_json::Value = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetActiveDeliveries
    /// Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
    /// Permissions Required:
    /// - View Sales Delivery
    /// - Manage Sales Delivery
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_sales_active_deliveries(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<ActiveDelivery>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let last_modified_end = last_modified_end.clone();
                let last_modified_start = last_modified_start.clone();
                let license_number = license_number.clone();
                let page_number = page_number.clone();
                let page_size = page_size.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_active_deliveries(last_modified_end, last_modified_start, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<ActiveDelivery> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetActiveDeliveriesRetailer
    /// Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
    /// Permissions Required:
    /// - View Retailer Delivery
    /// - Manage Retailer Delivery
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_sales_active_deliveries_retailer(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<ActiveDeliveriesRetailer>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let last_modified_end = last_modified_end.clone();
                let last_modified_start = last_modified_start.clone();
                let license_number = license_number.clone();
                let page_number = page_number.clone();
                let page_size = page_size.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_active_deliveries_retailer(last_modified_end, last_modified_start, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<ActiveDeliveriesRetailer> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetActiveReceipts
    /// Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
    /// Permissions Required:
    /// - View Sales
    /// - Manage Sales
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_sales_active_receipts(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<ActiveReceipt>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let last_modified_end = last_modified_end.clone();
                let last_modified_start = last_modified_start.clone();
                let license_number = license_number.clone();
                let page_number = page_number.clone();
                let page_size = page_size.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_active_receipts(last_modified_end, last_modified_start, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<ActiveReceipt> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetCounties
    /// Returns a list of counties available for sales deliveries.
    /// Parameters:
    pub async fn get_sales_counties(&self, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<County>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_counties(body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<County> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetCustomerTypes
    /// Returns a list of customer types.
    /// Parameters:
    pub async fn get_sales_customer_types(&self, body: Option<&Value>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_customer_types(body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: serde_json::Value = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetDeliveriesReturnReasons
    /// Returns a list of return reasons for sales deliveries based on the provided License Number.
    /// Permissions Required:
    /// - Sales Delivery
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_sales_deliveries_return_reasons(&self, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<DeliveriesReturnReason>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let page_number = page_number.clone();
                let page_size = page_size.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_deliveries_return_reasons(license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<DeliveriesReturnReason> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetDeliveryRetailerById
    /// Retrieves a retailer delivery record by its ID, with an optional License Number.
    /// Permissions Required:
    /// - View Retailer Delivery
    /// - Manage Retailer Delivery
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_sales_delivery_retailer_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<DeliveryRetailer>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let id = id.to_string();
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let id = id.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_delivery_retailer_by_id(&id, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: DeliveryRetailer = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetInactiveDeliveries
    /// Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
    /// Permissions Required:
    /// - View Sales Delivery
    /// - Manage Sales Delivery
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_sales_inactive_deliveries(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<InactiveDelivery>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let last_modified_end = last_modified_end.clone();
                let last_modified_start = last_modified_start.clone();
                let license_number = license_number.clone();
                let page_number = page_number.clone();
                let page_size = page_size.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_inactive_deliveries(last_modified_end, last_modified_start, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<InactiveDelivery> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetInactiveDeliveriesRetailer
    /// Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
    /// Permissions Required:
    /// - View Retailer Delivery
    /// - Manage Retailer Delivery
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_sales_inactive_deliveries_retailer(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<InactiveDeliveriesRetailer>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let last_modified_end = last_modified_end.clone();
                let last_modified_start = last_modified_start.clone();
                let license_number = license_number.clone();
                let page_number = page_number.clone();
                let page_size = page_size.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_inactive_deliveries_retailer(last_modified_end, last_modified_start, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<InactiveDeliveriesRetailer> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetInactiveReceipts
    /// Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
    /// Permissions Required:
    /// - View Sales
    /// - Manage Sales
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_sales_inactive_receipts(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<InactiveReceipt>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let last_modified_end = last_modified_end.clone();
                let last_modified_start = last_modified_start.clone();
                let license_number = license_number.clone();
                let page_number = page_number.clone();
                let page_size = page_size.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_inactive_receipts(last_modified_end, last_modified_start, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<InactiveReceipt> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetPatientRegistrationLocations
    /// Returns a list of valid Patient registration locations for sales.
    /// Parameters:
    pub async fn get_sales_patient_registration_locations(&self, body: Option<&Value>) -> std::result::Result<Option<Vec<PatientRegistrationLocation>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_patient_registration_locations(body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Vec<PatientRegistrationLocation> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetPaymentTypes
    /// Returns a list of available payment types for the specified License Number.
    /// Permissions Required:
    /// - View Sales Delivery
    /// - Manage Sales Delivery
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_sales_payment_types(&self, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_payment_types(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: serde_json::Value = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetReceiptById
    /// Retrieves a sales receipt by its Id, with an optional License Number.
    /// Permissions Required:
    /// - View Sales
    /// - Manage Sales
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_sales_receipt_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<ActiveReceipt>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let id = id.to_string();
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let id = id.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_receipt_by_id(&id, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: ActiveReceipt = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetReceiptsExternalByExternalNumber
    /// Retrieves a Sales Receipt by its external number, with an optional License Number.
    /// Permissions Required:
    /// - View Sales
    /// - Manage Sales
    /// Parameters:
    /// - external_number (str): Path parameter externalNumber
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_sales_receipts_external_by_external_number(&self, external_number: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<Vec<ReceiptsExternalByExternalNumber>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let external_number = external_number.to_string();
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let external_number = external_number.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_receipts_external_by_external_number(&external_number, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: Vec<ReceiptsExternalByExternalNumber> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetSalesDeliveryById
    /// Retrieves a sales delivery record by its Id, with an optional License Number.
    /// Permissions Required:
    /// - View Sales Delivery
    /// - Manage Sales Delivery
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_sales_delivery_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<SalesDelivery>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let id = id.to_string();
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let id = id.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().get_sales_delivery_by_id(&id, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: SalesDelivery = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// PUT UpdateDeliveries
    /// Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// Permissions Required:
    /// - Manage Sales Delivery
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_sales_deliveries(&self, license_number: Option<String>, body: Option<Vec<UpdateDeliveriesRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().update_sales_deliveries(license_number, body_val.as_ref()
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

    /// PUT UpdateDeliveriesComplete
    /// Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
    /// Permissions Required:
    /// - Manage Sales Delivery
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_sales_deliveries_complete(&self, license_number: Option<String>, body: Option<Vec<UpdateDeliveriesCompleteRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().update_sales_deliveries_complete(license_number, body_val.as_ref()
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

    /// PUT UpdateDeliveriesHub
    /// Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// Permissions Required:
    /// - Manage Sales Delivery, Manage Sales Delivery Hub
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_sales_deliveries_hub(&self, license_number: Option<String>, body: Option<Vec<UpdateDeliveriesHubRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().update_sales_deliveries_hub(license_number, body_val.as_ref()
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

    /// PUT UpdateDeliveriesHubAccept
    /// Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
    /// Permissions Required:
    /// - Manage Sales Delivery Hub
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_sales_deliveries_hub_accept(&self, license_number: Option<String>, body: Option<Vec<UpdateDeliveriesHubAcceptRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().update_sales_deliveries_hub_accept(license_number, body_val.as_ref()
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

    /// PUT UpdateDeliveriesHubDepart
    /// Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
    /// Permissions Required:
    /// - Manage Sales Delivery Hub
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_sales_deliveries_hub_depart(&self, license_number: Option<String>, body: Option<Vec<UpdateDeliveriesHubDepartRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().update_sales_deliveries_hub_depart(license_number, body_val.as_ref()
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

    /// PUT UpdateDeliveriesHubVerifyID
    /// Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
    /// Permissions Required:
    /// - Manage Sales Delivery Hub
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_sales_deliveries_hub_verify_id(&self, license_number: Option<String>, body: Option<Vec<UpdateDeliveriesHubVerifyIDRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().update_sales_deliveries_hub_verify_id(license_number, body_val.as_ref()
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

    /// PUT UpdateDeliveriesRetailer
    /// Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// Permissions Required:
    /// - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
    /// - Industry/Facility Type/Retailer Delivery
    /// - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
    /// - WebApi Sales Deliveries Read Write State (All or WriteOnly)
    /// - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
    /// - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
    /// - Manage Retailer Delivery
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_sales_deliveries_retailer(&self, license_number: Option<String>, body: Option<Vec<UpdateDeliveriesRetailerRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().update_sales_deliveries_retailer(license_number, body_val.as_ref()
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

    /// PUT UpdateReceipts
    /// Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
    /// Permissions Required:
    /// - Manage Sales
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_sales_receipts(&self, license_number: Option<String>, body: Option<Vec<UpdateReceiptsRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().update_sales_receipts(license_number, body_val.as_ref()
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

    /// PUT UpdateReceiptsFinalize
    /// Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
    /// Permissions Required:
    /// - Manage Sales
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_sales_receipts_finalize(&self, license_number: Option<String>, body: Option<Vec<UpdateReceiptsFinalizeRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().update_sales_receipts_finalize(license_number, body_val.as_ref()
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

    /// PUT UpdateReceiptsUnfinalize
    /// Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
    /// Permissions Required:
    /// - Manage Sales
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_sales_receipts_unfinalize(&self, license_number: Option<String>, body: Option<Vec<UpdateReceiptsUnfinalizeRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.sales().update_sales_receipts_unfinalize(license_number, body_val.as_ref()
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
}

