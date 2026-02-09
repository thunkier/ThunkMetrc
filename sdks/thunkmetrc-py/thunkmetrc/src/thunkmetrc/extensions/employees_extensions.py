from typing import Generator, List, Dict, Tuple, TYPE_CHECKING
import datetime
from .metrc_extensions import sync_parallel
from thunkmetrc.wrapper.ratelimiter import MetrcRateLimiter

if TYPE_CHECKING:
    from thunkmetrc.wrapper.wrapper import MetrcWrapper
    from thunkmetrc.wrapper.models import *

class EmployeesExtensions:
    @staticmethod
    async def iterate_get_employees(
        service,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Employee']:
        """
        Iterates through all pages of GetEmployees.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_employees(
                license_number=license_number,
                page_size=page_size,
                page_number=page
            )
            if not response.data:
                break
            all_results.extend(response.data)
            if not response.meta or response.meta.total_pages <= page:
                break
            page += 1
        return all_results
