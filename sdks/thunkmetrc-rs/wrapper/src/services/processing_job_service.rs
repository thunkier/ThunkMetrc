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

pub struct ProcessingJobService {
    client: MetrcClient,
    rate_limiter: Arc<MetrcRateLimiter>,
}

impl ProcessingJobService {
    pub fn new(client: MetrcClient, rate_limiter: Arc<MetrcRateLimiter>) -> Self {
        Self {
            client,
            rate_limiter,
        }
    }

    /// POST CreateAdjustProcessingJob
    /// Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_adjust_processing_job(&self, license_number: Option<String>, body: Option<Vec<CreateAdjustProcessingJobRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.processing_job().create_adjust_processing_job(license_number, body_val.as_ref()
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

    /// POST CreateJobTypes
    /// Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_processing_job_job_types(&self, license_number: Option<String>, body: Option<Vec<CreateJobTypesRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.processing_job().create_processing_job_job_types(license_number, body_val.as_ref()
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

    /// POST CreateProcessingJobPackages
    /// Creates packages from processing jobs at a Facility, including optional location and note assignments.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn create_processing_job_packages(&self, license_number: Option<String>, body: Option<Vec<CreateProcessingJobPackagesRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.processing_job().create_processing_job_packages(license_number, body_val.as_ref()
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

    /// DELETE DeleteJobTypeById
    /// Archives a Processing Job Type at a Facility, making it inactive for future use.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_processing_job_job_type_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
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
                    client.processing_job().delete_processing_job_job_type_by_id(&id, license_number, body_val.as_ref()
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

    /// DELETE DeleteProcessingJobById
    /// Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn delete_processing_job_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
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
                    client.processing_job().delete_processing_job_by_id(&id, license_number, body_val.as_ref()
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

    /// PUT FinishProcessingJob
    /// Completes processing jobs at a Facility by recording final notes and waste measurements.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn finish_processing_job_processing_job(&self, license_number: Option<String>, body: Option<Vec<FinishProcessingJobRequestItem>>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.processing_job().finish_processing_job_processing_job(license_number, body_val.as_ref()
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

    /// GET GetActiveJobTypes
    /// Retrieves a list of all active processing job types defined within a Facility.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_processing_job_active_job_types(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<ActiveJobType>>, Box<dyn Error + Send + Sync>> {
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
                    client.processing_job().get_processing_job_active_job_types(last_modified_end, last_modified_start, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<ActiveJobType> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetActiveProcessingJob
    /// Retrieves active processing jobs at a specified Facility.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_active_processing_job(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<ProcessingJob>>, Box<dyn Error + Send + Sync>> {
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
                    client.processing_job().get_active_processing_job(last_modified_end, last_modified_start, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<ProcessingJob> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetInactiveJobTypes
    /// Retrieves a list of all inactive processing job types defined within a Facility.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_processing_job_inactive_job_types(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<InactiveJobType>>, Box<dyn Error + Send + Sync>> {
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
                    client.processing_job().get_processing_job_inactive_job_types(last_modified_end, last_modified_start, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<InactiveJobType> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetInactiveProcessingJob
    /// Retrieves inactive processing jobs at a specified Facility.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - last_modified_end (Option<String>): Filter by lastModifiedEnd
    /// - last_modified_start (Option<String>): Filter by lastModifiedStart
    /// - license_number (Option<String>): Filter by licenseNumber
    /// - page_number (Option<String>): Filter by pageNumber
    /// - page_size (Option<String>): Filter by pageSize
    pub async fn get_inactive_processing_job(&self, last_modified_end: Option<String>, last_modified_start: Option<String>, license_number: Option<String>, page_number: Option<String>, page_size: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<ProcessingJob>>, Box<dyn Error + Send + Sync>> {
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
                    client.processing_job().get_inactive_processing_job(last_modified_end, last_modified_start, license_number, page_number, page_size, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<ProcessingJob> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetJobTypesAttributes
    /// Retrieves all processing job attributes available for a Facility.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_processing_job_job_types_attributes(&self, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<JobTypesAttribute>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.processing_job().get_processing_job_job_types_attributes(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<JobTypesAttribute> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetJobTypesCategories
    /// Retrieves all processing job categories available for a specified Facility.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_processing_job_job_types_categories(&self, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<PaginatedResponse<JobTypesCategory>>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,true,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.processing_job().get_processing_job_job_types_categories(license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: PaginatedResponse<JobTypesCategory> = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// GET GetProcessingJobById
    /// Retrieves a ProcessingJob by Id.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - id (str): Path parameter id
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn get_processing_job_by_id(&self, id: &str, license_number: Option<String>, body: Option<&Value>) -> std::result::Result<Option<ProcessingJob>, Box<dyn Error + Send + Sync>> {
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
                    client.processing_job().get_processing_job_by_id(&id, license_number, body_val.as_ref()
                    ).await.map_err(|e| e as Box<dyn Error + Send + Sync>)
                }
            }
        ).await?;

        if let Some(v) = resp_val {
            let typed: ProcessingJob = serde_json::from_value(v)?;
            Ok(Some(typed))
        } else {
            Ok(None)
        }
    }

    /// POST StartProcessingJob
    /// Initiates new processing jobs at a Facility, including job details and associated packages.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn start_processing_job_processing_job(&self, license_number: Option<String>, body: Option<Vec<StartProcessingJobRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.processing_job().start_processing_job_processing_job(license_number, body_val.as_ref()
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

    /// PUT UnfinishProcessingJob
    /// Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn unfinish_processing_job_processing_job(&self, license_number: Option<String>, body: Option<Vec<UnfinishProcessingJobRequestItem>>) -> std::result::Result<Option<serde_json::Value>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.processing_job().unfinish_processing_job_processing_job(license_number, body_val.as_ref()
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

    /// PUT UpdateJobTypes
    /// Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
    /// Permissions Required:
    /// - Manage Processing Job
    /// Parameters:
    /// - body (Option<&Value>): Request body
    /// - license_number (Option<String>): Filter by licenseNumber
    pub async fn update_processing_job_job_types(&self, license_number: Option<String>, body: Option<Vec<UpdateJobTypesRequestItem>>) -> std::result::Result<Option<WriteResponse>, Box<dyn Error + Send + Sync>> {
        let body_val = if let Some(b) = body { Some(serde_json::to_value(b)?) } else { None };
        let client = self.client.clone();
        
        let body_val = body_val.clone();

        let resp_val = self.rate_limiter.execute(None,false,
            move || {
                let client = client.clone();
                let license_number = license_number.clone();
                let body_val = body_val.clone();

                async move {
                    client.processing_job().update_processing_job_job_types(license_number, body_val.as_ref()
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

