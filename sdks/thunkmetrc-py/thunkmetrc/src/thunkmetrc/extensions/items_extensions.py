from typing import Generator, List, Dict, Tuple, TYPE_CHECKING
import datetime
from .metrc_extensions import sync_parallel
from thunkmetrc.wrapper.ratelimiter import MetrcRateLimiter

if TYPE_CHECKING:
    from thunkmetrc.wrapper.wrapper import MetrcWrapper
    from thunkmetrc.wrapper.models import *

class ItemsExtensions:
    @staticmethod
    async def iterate_get_active_items(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Item']:
        """
        Iterates through all pages of GetActiveItems.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_active_items(
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
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
    @staticmethod
    async def iterate_get_items_brands(
        service,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Brand']:
        """
        Iterates through all pages of GetBrands.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_items_brands(
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
    @staticmethod
    async def iterate_get_items_categories(
        service,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Category']:
        """
        Iterates through all pages of GetCategories.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_items_categories(
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
    @staticmethod
    async def iterate_get_inactive_items(
        service,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Item']:
        """
        Iterates through all pages of GetInactiveItems.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_inactive_items(
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
    
    @staticmethod
    async def sync_get_active_items(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['Item']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await ItemsExtensions.iterate_get_active_items(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_active_items_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['Item']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.items
             return await ItemsExtensions.sync_get_active_items(
                 wrapper.items,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
