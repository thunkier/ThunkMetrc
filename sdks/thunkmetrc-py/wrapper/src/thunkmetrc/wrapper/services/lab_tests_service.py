from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class LabTestsService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_lab_tests_record(self, body: Optional[List['LabTestsCreateRecordRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        POST /labtests/v2/record
        Parameters:
        - body (List['LabTestsCreateRecordRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.lab_tests.create_lab_tests_record(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_lab_tests_batches(self, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Batch']:
        """
        Retrieves a list of Lab Test batches.
        GET /labtests/v2/batches
        Parameters:
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.lab_tests.get_lab_tests_batches(body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_lab_tests_lab_test_document_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a specific Lab Test result document by its Id for a given Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        GET /labtests/v2/labtestdocument/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.lab_tests.get_lab_tests_lab_test_document_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_lab_tests_types(self, body: Any = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['LabTestsType']:
        """
        Returns a list of Lab Test types.
        GET /labtests/v2/types
        Parameters:
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.lab_tests.get_lab_tests_types(body=body,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_lab_tests_results(self, body: Any = None, license_number: Optional[str] = None, package_id: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Result']:
        """
        Retrieves Lab Test results for a specified Package.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        GET /labtests/v2/results
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - package_id (str, optional): Filter by packageId
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.lab_tests.get_lab_tests_results(body=body,
                license_number=license_number,
                package_id=package_id,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_lab_tests_states(self, body: Any = None) -> Any:
        """
        Returns a list of all lab testing states.
        GET /labtests/v2/states
        Parameters:
        """
        async def op():
            return await self.client.lab_tests.get_lab_tests_states(body=body,
            )
        return await self.limiter.execute(None, True, op)

    async def update_lab_tests_lab_test_document(self, body: Optional[List['LabTestsUpdateLabTestDocumentRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates one or more documents for previously submitted lab tests.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT /labtests/v2/labtestdocument
        Parameters:
        - body (List['LabTestsUpdateLabTestDocumentRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.lab_tests.update_lab_tests_lab_test_document(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_lab_tests_results_release(self, body: Optional[List['LabTestsUpdateResultsReleaseRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Releases Lab Test results for one or more packages.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT /labtests/v2/results/release
        Parameters:
        - body (List['LabTestsUpdateResultsReleaseRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.lab_tests.update_lab_tests_results_release(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

