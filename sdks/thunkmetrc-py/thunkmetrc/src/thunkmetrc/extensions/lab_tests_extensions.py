from typing import Generator, List, Dict, Tuple, TYPE_CHECKING
import datetime
from .metrc_extensions import sync_parallel
from thunkmetrc.wrapper.ratelimiter import MetrcRateLimiter

if TYPE_CHECKING:
    from thunkmetrc.wrapper.wrapper import MetrcWrapper
    from thunkmetrc.wrapper.models import *

class LabTestsExtensions:
    @staticmethod
    async def iterate_get_lab_tests_batches(
        service,
        page_size: str = None,
    ) -> List['Batch']:
        """
        Iterates through all pages of GetBatches.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_lab_tests_batches(
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
    @staticmethod
    async def iterate_get_lab_tests_types(
        service,
        page_size: str = None,
    ) -> List['LabTestsType']:
        """
        Iterates through all pages of GetLabTestsTypes.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_lab_tests_types(
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
    @staticmethod
    async def iterate_get_lab_tests_results(
        service,
        license_number: str = None,
        package_id: str = None,
        page_size: str = None,
    ) -> List['Result']:
        """
        Iterates through all pages of GetResults.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_lab_tests_results(
                license_number=license_number,
                package_id=package_id,
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
