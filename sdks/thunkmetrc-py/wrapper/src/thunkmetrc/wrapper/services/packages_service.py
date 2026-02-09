from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class PackagesService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_adjust_packages(self, body: Optional[List['PackagesCreateAdjustPackagesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Records a list of adjustments for packages at a specific Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        POST /packages/v2/adjust
        Parameters:
        - body (List['PackagesCreateAdjustPackagesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.create_adjust_packages(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_packages_packages(self, body: Optional[List['PackagesCreatePackagesPackagesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates new packages for a specified Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        POST /packages/v2
        Parameters:
        - body (List['PackagesCreatePackagesPackagesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.create_packages_packages(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_packages_plantings(self, body: Optional[List['PackagesCreatePackagesPlantingsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates new plantings from packages for a specified Facility.
        Permissions Required:
        - View Immature Plants
        - Manage Immature Plants
        - View Packages
        - Manage Packages Inventory
        POST /packages/v2/plantings
        Parameters:
        - body (List['PackagesCreatePackagesPlantingsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.create_packages_plantings(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_packages_testing(self, body: Optional[List['PackagesCreateTestingRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates new packages for testing for a specified Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        POST /packages/v2/testing
        Parameters:
        - body (List['PackagesCreateTestingRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.create_packages_testing(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_packages_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Discontinues a Package at a specific Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        DELETE /packages/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.delete_packages_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def finish_packages_packages(self, body: Optional[List['PackagesFinishPackagesRequest']] = None, license_number: Optional[str] = None) -> Any:
        """
        Updates a list of packages as finished for a specific Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT /packages/v2/finish
        Parameters:
        - body (List['PackagesFinishPackagesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.finish_packages_packages(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def finishedgood_flag_packages(self, body: Optional[List['PackagesFinishedgoodFlagRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Flags one or more Packages at the specified Facility as Finished Goods.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT /packages/v2/finishedgood/flag
        Parameters:
        - body (List['PackagesFinishedgoodFlagRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.finishedgood_flag_packages(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def finishedgood_unflag_packages(self, body: Optional[List['PackagesFinishedgoodUnflagRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Removes the Finished Good flag one or more Packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT /packages/v2/finishedgood/unflag
        Parameters:
        - body (List['PackagesFinishedgoodUnflagRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.finishedgood_unflag_packages(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_active_packages(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['PackagesPackage']:
        """
        Retrieves a list of active packages for a specified Facility.
        Permissions Required:
        - View Packages
        GET /packages/v2/active
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.packages.get_active_packages(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_packages_adjust_reasons(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['AdjustReason']:
        """
        Retrieves a list of adjustment reasons for packages at a specified Facility.
        GET /packages/v2/adjust/reasons
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.packages.get_packages_adjust_reasons(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_packages_adjustments(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Adjustment']:
        """
        Retrieves the Package Adjustments for a Facility
        Permissions Required:
        - View Packages
        GET /packages/v2/adjustments
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.packages.get_packages_adjustments(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_in_transit_packages(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['InTransit']:
        """
        Retrieves a list of packages in transit for a specified Facility.
        Permissions Required:
        - View Packages
        GET /packages/v2/intransit
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.packages.get_in_transit_packages(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_inactive_packages(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['PackagesPackage']:
        """
        Retrieves a list of inactive packages for a specified Facility.
        Permissions Required:
        - View Packages
        GET /packages/v2/inactive
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.packages.get_inactive_packages(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_packages_lab_samples(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['PackagesPackage']:
        """
        Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
        Permissions Required:
        - View Packages
        GET /packages/v2/labsamples
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.packages.get_packages_lab_samples(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_on_hold_packages(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['PackagesPackage']:
        """
        Retrieves a list of packages on hold for a specified Facility.
        Permissions Required:
        - View Packages
        GET /packages/v2/onhold
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.packages.get_on_hold_packages(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_packages_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'PackagesPackage':
        """
        Retrieves a Package by its Id.
        Permissions Required:
        - View Packages
        GET /packages/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.get_packages_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_packages_by_label(self, label: str, body: Any = None, license_number: Optional[str] = None) -> List['PackagesPackage']:
        """
        Retrieves a Package by its label.
        Permissions Required:
        - View Packages
        GET /packages/v2/{label}
        Parameters:
        - label (str): Path parameter label
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.get_packages_by_label(label, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_packages_types(self, body: Any = None) -> Any:
        """
        Retrieves a list of available Package types.
        GET /packages/v2/types
        Parameters:
        """
        async def op():
            return await self.client.packages.get_packages_types(body=body,
            )
        return await self.limiter.execute(None, True, op)

    async def get_packages_source_harvest_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'SourceHarvest':
        """
        Retrieves the source harvests for a Package by its Id.
        Permissions Required:
        - View Package Source Harvests
        GET /packages/v2/{id}/source/harvests
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.get_packages_source_harvest_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_packages_transferred(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['PackagesPackage']:
        """
        Retrieves a list of transferred packages for a specific Facility.
        Permissions Required:
        - View Packages
        GET /packages/v2/transferred
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.packages.get_packages_transferred(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def unfinish_packages_packages(self, body: Optional[List['PackagesUnfinishPackagesRequest']] = None, license_number: Optional[str] = None) -> Any:
        """
        Updates a list of packages as unfinished for a specific Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT /packages/v2/unfinish
        Parameters:
        - body (List['PackagesUnfinishPackagesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.unfinish_packages_packages(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_adjust_packages(self, body: Optional[List['PackagesUpdateAdjustPackagesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Set the final quantity for a Package.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT /packages/v2/adjust
        Parameters:
        - body (List['PackagesUpdateAdjustPackagesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.update_adjust_packages(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_packages_decontaminate(self, body: Optional[List['PackagesUpdateDecontaminateRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates the Product decontaminate information for a list of packages at a specific Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT /packages/v2/decontaminate
        Parameters:
        - body (List['PackagesUpdateDecontaminateRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.update_packages_decontaminate(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_packages_donation_flag(self, body: Optional[List['PackagesUpdateDonationFlagRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Flags one or more packages for donation at the specified Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT /packages/v2/donation/flag
        Parameters:
        - body (List['PackagesUpdateDonationFlagRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.update_packages_donation_flag(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_packages_donation_unflag(self, body: Optional[List['PackagesUpdateDonationUnflagRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Removes the donation flag from one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT /packages/v2/donation/unflag
        Parameters:
        - body (List['PackagesUpdateDonationUnflagRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.update_packages_donation_unflag(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_packages_externalid(self, body: Optional[List['PackagesUpdateExternalidRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates the external identifiers for one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Manage Package Inventory
        - External Id Enabled
        PUT /packages/v2/externalid
        Parameters:
        - body (List['PackagesUpdateExternalidRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.update_packages_externalid(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_packages_item(self, body: Optional[List['PackagesUpdateItemRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates the associated Item for one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        PUT /packages/v2/item
        Parameters:
        - body (List['PackagesUpdateItemRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.update_packages_item(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_packages_labtests_required(self, body: Optional[List['PackagesUpdateLabtestsRequiredRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates the list of required lab test batches for one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        PUT /packages/v2/labtests/required
        Parameters:
        - body (List['PackagesUpdateLabtestsRequiredRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.update_packages_labtests_required(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_packages_note(self, body: Optional[List['PackagesUpdateNoteRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates notes associated with one or more packages for the specified Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        - Manage Package Notes
        PUT /packages/v2/note
        Parameters:
        - body (List['PackagesUpdateNoteRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.update_packages_note(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_packages_location(self, body: Optional[List['PackagesUpdatePackagesLocationRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates the Location and Sublocation for one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        PUT /packages/v2/location
        Parameters:
        - body (List['PackagesUpdatePackagesLocationRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.update_packages_location(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_packages_remediate(self, body: Optional[List['PackagesUpdateRemediateRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates a list of Product remediations for packages at a specific Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT /packages/v2/remediate
        Parameters:
        - body (List['PackagesUpdateRemediateRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.update_packages_remediate(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_packages_trade_sample_flag(self, body: Optional[List['PackagesUpdateTradeSampleFlagRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Flags or unflags one or more packages at the specified Facility as trade samples.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT /packages/v2/tradesample/flag
        Parameters:
        - body (List['PackagesUpdateTradeSampleFlagRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.update_packages_trade_sample_flag(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_packages_trade_sample_unflag(self, body: Optional[List['PackagesUpdateTradeSampleUnflagRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Removes the trade sample flag from one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Manage Packages Inventory
        PUT /packages/v2/tradesample/unflag
        Parameters:
        - body (List['PackagesUpdateTradeSampleUnflagRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.update_packages_trade_sample_unflag(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_packages_use_by_date(self, body: Optional[List['PackagesUpdateUseByDateRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates the use-by date for one or more packages at the specified Facility.
        Permissions Required:
        - View Packages
        - Create/Submit/Discontinue Packages
        PUT /packages/v2/usebydate
        Parameters:
        - body (List['PackagesUpdateUseByDateRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.packages.update_packages_use_by_date(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

