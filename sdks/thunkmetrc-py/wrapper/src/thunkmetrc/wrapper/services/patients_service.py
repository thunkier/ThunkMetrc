from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class PatientsService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_patients(self, body: Optional[List['PatientsCreatePatientsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Adds new patients to a specified Facility.
        Permissions Required:
        - Manage Patients
        POST /patients/v2
        Parameters:
        - body (List['PatientsCreatePatientsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.patients.create_patients(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_patients_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Removes a Patient, identified by an Id, from a specified Facility.
        Permissions Required:
        - Manage Patients
        DELETE /patients/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.patients.delete_patients_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_active_patients(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Patient']:
        """
        Retrieves a list of active patients for a specified Facility.
        Permissions Required:
        - Manage Patients
        GET /patients/v2/active
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.patients.get_active_patients(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_patients_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'Patient':
        """
        Retrieves a Patient by Id.
        Permissions Required:
        - Manage Patients
        GET /patients/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.patients.get_patients_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def update_patients(self, body: Optional[List['PatientsUpdatePatientsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates Patient information for a specified Facility.
        Permissions Required:
        - Manage Patients
        PUT /patients/v2
        Parameters:
        - body (List['PatientsUpdatePatientsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.patients.update_patients(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

