from typing import Generator, List, Dict, Tuple, TYPE_CHECKING
import datetime
from .metrc_extensions import sync_parallel
from thunkmetrc.wrapper.ratelimiter import MetrcRateLimiter

if TYPE_CHECKING:
    from thunkmetrc.wrapper.wrapper import MetrcWrapper
    from thunkmetrc.wrapper.models import *

class PackagesExtensions:
    @staticmethod
    async def iterate_get_active_packages(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['PackagesPackage']:
        """
        Iterates through all pages of GetActivePackages.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_active_packages(
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
    async def iterate_get_packages_adjust_reasons(
        service,
        license_number: str = None,
        page_size: str = None,
    ) -> List['AdjustReason']:
        """
        Iterates through all pages of GetAdjustReasons.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_packages_adjust_reasons(
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
    async def iterate_get_packages_adjustments(
        service,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Adjustment']:
        """
        Iterates through all pages of GetAdjustments.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_packages_adjustments(
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
    async def iterate_get_in_transit_packages(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['InTransit']:
        """
        Iterates through all pages of GetInTransitPackages.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_in_transit_packages(
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
    async def iterate_get_inactive_packages(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['PackagesPackage']:
        """
        Iterates through all pages of GetInactivePackages.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_inactive_packages(
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
    async def iterate_get_packages_lab_samples(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['PackagesPackage']:
        """
        Iterates through all pages of GetLabSamples.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_packages_lab_samples(
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
    async def iterate_get_on_hold_packages(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['PackagesPackage']:
        """
        Iterates through all pages of GetOnHoldPackages.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_on_hold_packages(
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
    async def iterate_get_packages_transferred(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['PackagesPackage']:
        """
        Iterates through all pages of GetTransferred.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_packages_transferred(
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
    async def sync_get_active_packages(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['PackagesPackage']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PackagesExtensions.iterate_get_active_packages(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_active_packages_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['PackagesPackage']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.packages
             return await PackagesExtensions.sync_get_active_packages(
                 wrapper.packages,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_in_transit_packages(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['InTransit']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PackagesExtensions.iterate_get_in_transit_packages(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_in_transit_packages_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['InTransit']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.packages
             return await PackagesExtensions.sync_get_in_transit_packages(
                 wrapper.packages,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_inactive_packages(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['PackagesPackage']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PackagesExtensions.iterate_get_inactive_packages(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_inactive_packages_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['PackagesPackage']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.packages
             return await PackagesExtensions.sync_get_inactive_packages(
                 wrapper.packages,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_packages_lab_samples(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['PackagesPackage']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PackagesExtensions.iterate_get_packages_lab_samples(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_packages_lab_samples_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['PackagesPackage']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.packages
             return await PackagesExtensions.sync_get_packages_lab_samples(
                 wrapper.packages,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_on_hold_packages(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['PackagesPackage']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PackagesExtensions.iterate_get_on_hold_packages(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_on_hold_packages_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['PackagesPackage']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.packages
             return await PackagesExtensions.sync_get_on_hold_packages(
                 wrapper.packages,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_packages_transferred(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['PackagesPackage']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PackagesExtensions.iterate_get_packages_transferred(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_packages_transferred_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['PackagesPackage']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.packages
             return await PackagesExtensions.sync_get_packages_transferred(
                 wrapper.packages,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
