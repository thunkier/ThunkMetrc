from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class ProcessingJobService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_adjust_processing_job(self, body: Optional[List['ProcessingJobCreateAdjustProcessingJobRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
        Permissions Required:
        - Manage Processing Job
        POST /processing/v2/adjust
        Parameters:
        - body (List['ProcessingJobCreateAdjustProcessingJobRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.processing_job.create_adjust_processing_job(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_processing_job_job_types(self, body: Optional[List['ProcessingJobCreateJobTypesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
        Permissions Required:
        - Manage Processing Job
        POST /processing/v2/jobtypes
        Parameters:
        - body (List['ProcessingJobCreateJobTypesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.processing_job.create_processing_job_job_types(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_processing_job_packages(self, body: Optional[List['ProcessingJobCreateProcessingJobPackagesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates packages from processing jobs at a Facility, including optional location and note assignments.
        Permissions Required:
        - Manage Processing Job
        POST /processing/v2/createpackages
        Parameters:
        - body (List['ProcessingJobCreateProcessingJobPackagesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.processing_job.create_processing_job_packages(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_processing_job_job_type_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Archives a Processing Job Type at a Facility, making it inactive for future use.
        Permissions Required:
        - Manage Processing Job
        DELETE /processing/v2/jobtypes/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.processing_job.delete_processing_job_job_type_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_processing_job_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
        Permissions Required:
        - Manage Processing Job
        DELETE /processing/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.processing_job.delete_processing_job_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def finish_processing_job_processing_job(self, body: Optional[List['ProcessingJobFinishProcessingJobRequest']] = None, license_number: Optional[str] = None) -> Any:
        """
        Completes processing jobs at a Facility by recording final notes and waste measurements.
        Permissions Required:
        - Manage Processing Job
        PUT /processing/v2/finish
        Parameters:
        - body (List['ProcessingJobFinishProcessingJobRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.processing_job.finish_processing_job_processing_job(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_processing_job_active_job_types(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['ActiveJobType']:
        """
        Retrieves a list of all active processing job types defined within a Facility.
        Permissions Required:
        - Manage Processing Job
        GET /processing/v2/jobtypes/active
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.processing_job.get_processing_job_active_job_types(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_active_processing_job(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['ProcessingJob']:
        """
        Retrieves active processing jobs at a specified Facility.
        Permissions Required:
        - Manage Processing Job
        GET /processing/v2/active
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.processing_job.get_active_processing_job(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_processing_job_inactive_job_types(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['InactiveJobType']:
        """
        Retrieves a list of all inactive processing job types defined within a Facility.
        Permissions Required:
        - Manage Processing Job
        GET /processing/v2/jobtypes/inactive
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.processing_job.get_processing_job_inactive_job_types(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_inactive_processing_job(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['ProcessingJob']:
        """
        Retrieves inactive processing jobs at a specified Facility.
        Permissions Required:
        - Manage Processing Job
        GET /processing/v2/inactive
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.processing_job.get_inactive_processing_job(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_processing_job_job_types_attributes(self, body: Any = None, license_number: Optional[str] = None) -> PaginatedResponse['JobTypesAttribute']:
        """
        Retrieves all processing job attributes available for a Facility.
        Permissions Required:
        - Manage Processing Job
        GET /processing/v2/jobtypes/attributes
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.processing_job.get_processing_job_job_types_attributes(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_processing_job_job_types_categories(self, body: Any = None, license_number: Optional[str] = None) -> PaginatedResponse['JobTypesCategory']:
        """
        Retrieves all processing job categories available for a specified Facility.
        Permissions Required:
        - Manage Processing Job
        GET /processing/v2/jobtypes/categories
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.processing_job.get_processing_job_job_types_categories(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_processing_job_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'ProcessingJob':
        """
        Retrieves a ProcessingJob by Id.
        Permissions Required:
        - Manage Processing Job
        GET /processing/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.processing_job.get_processing_job_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def start_processing_job_processing_job(self, body: Optional[List['ProcessingJobStartProcessingJobRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Initiates new processing jobs at a Facility, including job details and associated packages.
        Permissions Required:
        - Manage Processing Job
        POST /processing/v2/start
        Parameters:
        - body (List['ProcessingJobStartProcessingJobRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.processing_job.start_processing_job_processing_job(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def unfinish_processing_job_processing_job(self, body: Optional[List['ProcessingJobUnfinishProcessingJobRequest']] = None, license_number: Optional[str] = None) -> Any:
        """
        Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
        Permissions Required:
        - Manage Processing Job
        PUT /processing/v2/unfinish
        Parameters:
        - body (List['ProcessingJobUnfinishProcessingJobRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.processing_job.unfinish_processing_job_processing_job(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_processing_job_job_types(self, body: Optional[List['ProcessingJobUpdateJobTypesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
        Permissions Required:
        - Manage Processing Job
        PUT /processing/v2/jobtypes
        Parameters:
        - body (List['ProcessingJobUpdateJobTypesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.processing_job.update_processing_job_job_types(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

