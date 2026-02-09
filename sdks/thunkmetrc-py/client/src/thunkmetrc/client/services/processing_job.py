from typing import Any, Optional, Dict, List, Union
import urllib.parse

class ProcessingJobClient:
    def __init__(self, client):
        self.client = client
    def create_adjust_processing_job(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
        Permissions Required:
        - Manage Processing Job
        POST CreateAdjustProcessingJob
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/processing/v2/adjust"
        return self.client._send("POST", path, body, query_params)
    def create_processing_job_job_types(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
        Permissions Required:
        - Manage Processing Job
        POST CreateJobTypes
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/processing/v2/jobtypes"
        return self.client._send("POST", path, body, query_params)
    def create_processing_job_packages(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates packages from processing jobs at a Facility, including optional location and note assignments.
        Permissions Required:
        - Manage Processing Job
        POST CreateProcessingJobPackages
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/processing/v2/createpackages"
        return self.client._send("POST", path, body, query_params)
    def delete_processing_job_job_type_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Archives a Processing Job Type at a Facility, making it inactive for future use.
        Permissions Required:
        - Manage Processing Job
        DELETE DeleteJobTypeById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/processing/v2/jobtypes/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def delete_processing_job_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
        Permissions Required:
        - Manage Processing Job
        DELETE DeleteProcessingJobById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/processing/v2/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def finish_processing_job_processing_job(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Completes processing jobs at a Facility by recording final notes and waste measurements.
        Permissions Required:
        - Manage Processing Job
        PUT FinishProcessingJob
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/processing/v2/finish"
        return self.client._send("PUT", path, body, query_params)
    def get_processing_job_active_job_types(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of all active processing job types defined within a Facility.
        Permissions Required:
        - Manage Processing Job
        GET GetActiveJobTypes
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if last_modified_end is not None:
            query_params["lastModifiedEnd"] = last_modified_end
        if last_modified_start is not None:
            query_params["lastModifiedStart"] = last_modified_start
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/processing/v2/jobtypes/active"
        return self.client._send("GET", path, body, query_params)
    def get_active_processing_job(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves active processing jobs at a specified Facility.
        Permissions Required:
        - Manage Processing Job
        GET GetActiveProcessingJob
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if last_modified_end is not None:
            query_params["lastModifiedEnd"] = last_modified_end
        if last_modified_start is not None:
            query_params["lastModifiedStart"] = last_modified_start
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/processing/v2/active"
        return self.client._send("GET", path, body, query_params)
    def get_processing_job_inactive_job_types(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of all inactive processing job types defined within a Facility.
        Permissions Required:
        - Manage Processing Job
        GET GetInactiveJobTypes
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if last_modified_end is not None:
            query_params["lastModifiedEnd"] = last_modified_end
        if last_modified_start is not None:
            query_params["lastModifiedStart"] = last_modified_start
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/processing/v2/jobtypes/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_inactive_processing_job(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves inactive processing jobs at a specified Facility.
        Permissions Required:
        - Manage Processing Job
        GET GetInactiveProcessingJob
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if last_modified_end is not None:
            query_params["lastModifiedEnd"] = last_modified_end
        if last_modified_start is not None:
            query_params["lastModifiedStart"] = last_modified_start
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/processing/v2/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_processing_job_job_types_attributes(self, license_number: Optional[str] = None) -> Any:
        """
        Retrieves all processing job attributes available for a Facility.
        Permissions Required:
        - Manage Processing Job
        GET GetJobTypesAttributes
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/processing/v2/jobtypes/attributes"
        return self.client._send("GET", path, body, query_params)
    def get_processing_job_job_types_categories(self, license_number: Optional[str] = None) -> Any:
        """
        Retrieves all processing job categories available for a specified Facility.
        Permissions Required:
        - Manage Processing Job
        GET GetJobTypesCategories
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/processing/v2/jobtypes/categories"
        return self.client._send("GET", path, body, query_params)
    def get_processing_job_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a ProcessingJob by Id.
        Permissions Required:
        - Manage Processing Job
        GET GetProcessingJobById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/processing/v2/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def start_processing_job_processing_job(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Initiates new processing jobs at a Facility, including job details and associated packages.
        Permissions Required:
        - Manage Processing Job
        POST StartProcessingJob
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/processing/v2/start"
        return self.client._send("POST", path, body, query_params)
    def unfinish_processing_job_processing_job(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
        Permissions Required:
        - Manage Processing Job
        PUT UnfinishProcessingJob
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/processing/v2/unfinish"
        return self.client._send("PUT", path, body, query_params)
    def update_processing_job_job_types(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
        Permissions Required:
        - Manage Processing Job
        PUT UpdateJobTypes
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/processing/v2/jobtypes"
        return self.client._send("PUT", path, body, query_params)

