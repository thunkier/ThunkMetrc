from typing import Generator, List, Dict, Tuple, TYPE_CHECKING
import datetime
from .metrc_extensions import sync_parallel
from thunkmetrc.wrapper.ratelimiter import MetrcRateLimiter

if TYPE_CHECKING:
    from thunkmetrc.wrapper.wrapper import MetrcWrapper
    from thunkmetrc.wrapper.models import *

class PlantsExtensions:
    @staticmethod
    async def iterate_get_plants_additives(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Additive']:
        """
        Iterates through all pages of GetAdditives.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_plants_additives(
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
    async def iterate_get_flowering_plants(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Plant']:
        """
        Iterates through all pages of GetFloweringPlants.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_flowering_plants(
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
    async def iterate_get_inactive_plants(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Plant']:
        """
        Iterates through all pages of GetInactivePlants.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_inactive_plants(
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
    async def iterate_get_mother_inactive_plants(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Mother']:
        """
        Iterates through all pages of GetMotherInactivePlants.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_mother_inactive_plants(
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
    async def iterate_get_mother_on_hold_plants(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Mother']:
        """
        Iterates through all pages of GetMotherOnHoldPlants.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_mother_on_hold_plants(
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
    async def iterate_get_mother_plants(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Mother']:
        """
        Iterates through all pages of GetMotherPlants.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_mother_plants(
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
    async def iterate_get_on_hold_plants(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Plant']:
        """
        Iterates through all pages of GetOnHoldPlants.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_on_hold_plants(
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
    async def iterate_get_plants_waste(
        service,
        license_number: str = None,
        page_size: str = None,
    ) -> List['PlantsWaste']:
        """
        Iterates through all pages of GetPlantsWaste.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_plants_waste(
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
    async def iterate_get_plants_waste_methods(
        service,
        page_size: str = None,
    ) -> List['WasteMethod']:
        """
        Iterates through all pages of GetPlantsWasteMethods.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_plants_waste_methods(
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
    async def iterate_get_plants_waste_reasons(
        service,
        license_number: str = None,
        page_size: str = None,
    ) -> List['WasteReason']:
        """
        Iterates through all pages of GetPlantsWasteReasons.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_plants_waste_reasons(
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
    async def iterate_get_vegetative_plants(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Plant']:
        """
        Iterates through all pages of GetVegetativePlants.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_vegetative_plants(
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
    async def iterate_get_plants_waste_by_id(
        service,
        license_number: str = None,
        page_size: str = None,
    ) -> List['PlantsWaste']:
        """
        Iterates through all pages of GetWasteById.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_plants_waste_by_id(
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
    async def iterate_get_plants_waste_package_by_id(
        service,
        license_number: str = None,
        page_size: str = None,
    ) -> List['WastePackage']:
        """
        Iterates through all pages of GetWastePackageById.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_plants_waste_package_by_id(
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
    async def sync_get_plants_additives(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['Additive']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PlantsExtensions.iterate_get_plants_additives(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_plants_additives_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['Additive']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.plants
             return await PlantsExtensions.sync_get_plants_additives(
                 wrapper.plants,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_flowering_plants(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['Plant']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PlantsExtensions.iterate_get_flowering_plants(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_flowering_plants_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['Plant']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.plants
             return await PlantsExtensions.sync_get_flowering_plants(
                 wrapper.plants,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_inactive_plants(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['Plant']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PlantsExtensions.iterate_get_inactive_plants(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_inactive_plants_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['Plant']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.plants
             return await PlantsExtensions.sync_get_inactive_plants(
                 wrapper.plants,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_mother_inactive_plants(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['Mother']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PlantsExtensions.iterate_get_mother_inactive_plants(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_mother_inactive_plants_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['Mother']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.plants
             return await PlantsExtensions.sync_get_mother_inactive_plants(
                 wrapper.plants,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_mother_on_hold_plants(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['Mother']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PlantsExtensions.iterate_get_mother_on_hold_plants(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_mother_on_hold_plants_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['Mother']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.plants
             return await PlantsExtensions.sync_get_mother_on_hold_plants(
                 wrapper.plants,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_mother_plants(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['Mother']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PlantsExtensions.iterate_get_mother_plants(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_mother_plants_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['Mother']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.plants
             return await PlantsExtensions.sync_get_mother_plants(
                 wrapper.plants,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_on_hold_plants(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['Plant']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PlantsExtensions.iterate_get_on_hold_plants(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_on_hold_plants_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['Plant']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.plants
             return await PlantsExtensions.sync_get_on_hold_plants(
                 wrapper.plants,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_vegetative_plants(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['Plant']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await PlantsExtensions.iterate_get_vegetative_plants(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_vegetative_plants_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['Plant']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.plants
             return await PlantsExtensions.sync_get_vegetative_plants(
                 wrapper.plants,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
