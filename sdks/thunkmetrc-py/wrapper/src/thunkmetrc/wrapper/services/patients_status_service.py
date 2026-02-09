from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class PatientsStatusService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def get_patients_statuses_by_patient_license_number(self, patient_license_number: str, body: Any = None, license_number: Optional[str] = None) -> List['PatientsStatus']:
        """
        Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
        Permissions Required:
        - Lookup Patients
        GET /patients/v2/statuses/{patientLicenseNumber}
        Parameters:
        - patient_license_number (str): Path parameter patientLicenseNumber
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.patients_status.get_patients_statuses_by_patient_license_number(patient_license_number, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

