from typing import Generator, List, Dict, Tuple, TYPE_CHECKING
import datetime
from .metrc_extensions import sync_parallel
from thunkmetrc.wrapper.ratelimiter import MetrcRateLimiter

if TYPE_CHECKING:
    from thunkmetrc.wrapper.wrapper import MetrcWrapper
    from thunkmetrc.wrapper.models import *

class TransfersExtensions:
    @staticmethod
    async def iterate_get_transfers_delivery_package_by_id(
        service,
        page_size: str = None,
    ) -> List['DeliveryPackage']:
        """
        Iterates through all pages of GetDeliveryPackageById.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_transfers_delivery_package_by_id(
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
    async def iterate_get_transfers_delivery_package_requiredlabtestbatch_by_id(
        service,
        page_size: str = None,
    ) -> List['DeliveryPackageRequiredlabtestbatch']:
        """
        Iterates through all pages of GetDeliveryPackageRequiredlabtestbatchById.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_transfers_delivery_package_requiredlabtestbatch_by_id(
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
    async def iterate_get_transfers_delivery_package_wholesale_by_id(
        service,
        page_size: str = None,
    ) -> List['DeliveryPackageWholesale']:
        """
        Iterates through all pages of GetDeliveryPackageWholesaleById.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_transfers_delivery_package_wholesale_by_id(
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
    async def iterate_get_transfers_delivery_transporter_by_id(
        service,
        page_size: str = None,
    ) -> List['DeliveryTransporter']:
        """
        Iterates through all pages of GetDeliveryTransporterById.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_transfers_delivery_transporter_by_id(
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
    async def iterate_get_transfers_delivery_transporter_detail_by_id(
        service,
        page_size: str = None,
    ) -> List['DeliveryTransporterDetail']:
        """
        Iterates through all pages of GetDeliveryTransporterDetailById.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_transfers_delivery_transporter_detail_by_id(
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
    async def iterate_get_transfers_hub(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Hub']:
        """
        Iterates through all pages of GetHub.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_transfers_hub(
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
    async def iterate_get_incoming_transfers(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Transfer']:
        """
        Iterates through all pages of GetIncomingTransfers.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_incoming_transfers(
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
    async def iterate_get_transfers_outgoing_template_delivery_by_id(
        service,
        page_size: str = None,
    ) -> List['TemplateDelivery']:
        """
        Iterates through all pages of GetOutgoingTemplateDeliveryById.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_transfers_outgoing_template_delivery_by_id(
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
    async def iterate_get_transfers_outgoing_template_delivery_package_by_id(
        service,
        page_size: str = None,
    ) -> List['TemplateDeliveryPackage']:
        """
        Iterates through all pages of GetOutgoingTemplateDeliveryPackageById.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_transfers_outgoing_template_delivery_package_by_id(
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
    async def iterate_get_transfers_outgoing_template_delivery_transporter_by_id(
        service,
        page_size: str = None,
    ) -> List['TemplateDeliveryTransporter']:
        """
        Iterates through all pages of GetOutgoingTemplateDeliveryTransporterById.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_transfers_outgoing_template_delivery_transporter_by_id(
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
    async def iterate_get_transfers_outgoing_templates(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Template']:
        """
        Iterates through all pages of GetOutgoingTemplates.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_transfers_outgoing_templates(
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
    async def iterate_get_outgoing_transfers(
        service,
        last_modified_end: str = None,
        last_modified_start: str = None,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Transfer']:
        """
        Iterates through all pages of GetOutgoingTransfers.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_outgoing_transfers(
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
    async def iterate_get_rejected_transfers(
        service,
        license_number: str = None,
        page_size: str = None,
    ) -> List['Transfer']:
        """
        Iterates through all pages of GetRejectedTransfers.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_rejected_transfers(
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
    async def iterate_get_transfers_delivery_by_id(
        service,
        page_size: str = None,
    ) -> List['TransfersDelivery']:
        """
        Iterates through all pages of GetTransfersDeliveryById.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_transfers_delivery_by_id(
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
    async def iterate_get_transfers_types(
        service,
        license_number: str = None,
        page_size: str = None,
    ) -> List['TransfersType']:
        """
        Iterates through all pages of GetTransfersTypes.
        """
        all_results = []
        page = 1
        while True:
            response = await service.get_transfers_types(
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
    async def sync_get_transfers_hub(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['Hub']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await TransfersExtensions.iterate_get_transfers_hub(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_transfers_hub_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['Hub']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.transfers
             return await TransfersExtensions.sync_get_transfers_hub(
                 wrapper.transfers,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_incoming_transfers(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['Transfer']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await TransfersExtensions.iterate_get_incoming_transfers(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_incoming_transfers_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['Transfer']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.transfers
             return await TransfersExtensions.sync_get_incoming_transfers(
                 wrapper.transfers,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_transfers_outgoing_templates(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['Template']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await TransfersExtensions.iterate_get_transfers_outgoing_templates(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_transfers_outgoing_templates_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['Template']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.transfers
             return await TransfersExtensions.sync_get_transfers_outgoing_templates(
                 wrapper.transfers,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
    
    @staticmethod
    async def sync_get_outgoing_transfers(
        service,
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        license_number: str
    ) -> List['Transfer']:
        """
        Syncs data based on time window.
        """
        end = datetime.datetime.now(datetime.timezone.utc)
        start = end - datetime.timedelta(days=1)
        if last_known_sync:
            start = last_known_sync - datetime.timedelta(minutes=buffer_minutes)

        start_str = start.isoformat()
        end_str = end.isoformat()

        return await TransfersExtensions.iterate_get_outgoing_transfers(
            service,
            last_modified_start=start_str,
            last_modified_end=end_str,
            license_number=license_number
        )

    @staticmethod
    async def sync_get_outgoing_transfers_parallel(
        targets: List[Tuple['MetrcWrapper', str]],
        last_known_sync: datetime.datetime,
        buffer_minutes: int,
        concurrency_limit: int = 20
    ) -> Dict[str, List['Transfer']]:
        """
        Syncs data in parallel across multiple targets.
        """
        async def op(wrapper: 'MetrcWrapper', limiter: MetrcRateLimiter, license_number: str):
             # Access service from wrapper
             # wrapper.transfers
             return await TransfersExtensions.sync_get_outgoing_transfers(
                 wrapper.transfers,
                 last_known_sync,
                 buffer_minutes,
                 license_number
             )
        
        return await sync_parallel(targets, concurrency_limit, op)
