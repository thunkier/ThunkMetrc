from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class PatientCheckInsService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_patient_check_ins(self, body: Optional[List['PatientCheckInsCreatePatientCheckInsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records patient check-ins for a specified Facility.
        Permissions Required:
        - ManagePatientsCheckIns
        POST /patient-checkins/v2
        Parameters:
        - body (List['PatientCheckInsCreatePatientCheckInsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.patient_check_ins.create_patient_check_ins(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_patient_check_ins_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Archives a Patient Check-In, identified by its Id, for a specified Facility.
        Permissions Required:
        - ManagePatientsCheckIns
        DELETE /patient-checkins/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.patient_check_ins.delete_patient_check_ins_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_patient_check_ins_locations(self, body: Any = None) -> List['PatientCheckInsLocation']:
        """
        Retrieves a list of Patient Check-In locations.
        GET /patient-checkins/v2/locations
        Parameters:
        """
        async def op():
            return await self.client.patient_check_ins.get_patient_check_ins_locations(body=body,
            )
        return await self.limiter.execute(None, True, op)

    async def get_patient_check_ins(self, body: Any = None, checkin_date_end: Optional[str] = None, checkin_date_start: Optional[str] = None, license_number: Optional[str] = None) -> PaginatedResponse['PatientCheckIn']:
        """
        Retrieves a list of patient check-ins for a specified Facility.
        Permissions Required:
        - ManagePatientsCheckIns
        GET /patient-checkins/v2
        Parameters:
        - checkin_date_end (str, optional): Filter by checkinDateEnd
        - checkin_date_start (str, optional): Filter by checkinDateStart
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.patient_check_ins.get_patient_check_ins(body=body,
                checkin_date_end=checkin_date_end,
                checkin_date_start=checkin_date_start,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def update_patient_check_ins(self, body: Optional[List['PatientCheckInsUpdatePatientCheckInsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates patient check-ins for a specified Facility.
        Permissions Required:
        - ManagePatientsCheckIns
        PUT /patient-checkins/v2
        Parameters:
        - body (List['PatientCheckInsUpdatePatientCheckInsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.patient_check_ins.update_patient_check_ins(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

