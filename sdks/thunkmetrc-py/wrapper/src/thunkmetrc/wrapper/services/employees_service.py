from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class EmployeesService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def get_employees(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Employee']:
        """
        Retrieves a list of employees for a specified Facility.
        Permissions Required:
        - Manage Employees
        - View Employees
        GET /employees/v2
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.employees.get_employees(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_employees_permissions(self, body: Any = None, employee_license_number: Optional[str] = None, license_number: Optional[str] = None) -> Any:
        """
        Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
        Permissions Required:
        - Manage Employees
        GET /employees/v2/permissions
        Parameters:
        - employee_license_number (str, optional): Filter by employeeLicenseNumber
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.employees.get_employees_permissions(body=body,
                employee_license_number=employee_license_number,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

