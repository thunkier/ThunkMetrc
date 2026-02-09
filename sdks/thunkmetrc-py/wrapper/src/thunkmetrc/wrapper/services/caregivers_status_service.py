from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class CaregiversStatusService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def get_caregivers_status_by_caregiver_license_number(self, caregiver_license_number: str, body: Any = None, license_number: Optional[str] = None) -> List['CaregiversStatus']:
        """
        Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
        Permissions Required:
        - Lookup Caregivers
        GET /caregivers/v2/status/{caregiverLicenseNumber}
        Parameters:
        - caregiver_license_number (str): Path parameter caregiverLicenseNumber
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.caregivers_status.get_caregivers_status_by_caregiver_license_number(caregiver_license_number, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

