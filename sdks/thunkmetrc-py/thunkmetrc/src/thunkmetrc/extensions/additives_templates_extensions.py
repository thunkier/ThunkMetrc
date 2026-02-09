from typing import Generator, List, Dict, Tuple, TYPE_CHECKING
import datetime
from .metrc_extensions import sync_parallel
from thunkmetrc.wrapper.ratelimiter import MetrcRateLimiter

if TYPE_CHECKING:
    from thunkmetrc.wrapper.wrapper import MetrcWrapper
    from thunkmetrc.wrapper.models import *

class AdditivesTemplatesExtensions:
    @staticmethod
    async def iterate_get_active_additives_templates(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['AdditivesTemplate']:
        """
        Iterates through all pages of GetActiveAdditivesTemplates.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_active_additives_templates(
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
    async def iterate_get_inactive_additives_templates(
        service,
        license_number: str = None,
        page_size: str = None,
    ) -> List['AdditivesTemplate']:
        """
        Iterates through all pages of GetInactiveAdditivesTemplates.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_inactive_additives_templates(
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
    async def sync_get_active_additives_templates(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['AdditivesTemplate']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await AdditivesTemplatesExtensions.iterate_get_active_additives_templates(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_active_additives_templates_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['AdditivesTemplate']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.additives_templates
             return await AdditivesTemplatesExtensions.sync_get_active_additives_templates(
                 wrapper.additives_templates,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
