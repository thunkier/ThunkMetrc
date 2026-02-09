from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class RetailIdService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_retail_id_associate(self, body: Optional[List['RetailIdCreateAssociateRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        - WebApi Retail ID Read Write State (All or WriteOnly)
        - Industry/View Packages
        - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
        POST /retailid/v2/associate
        Parameters:
        - body (List['RetailIdCreateAssociateRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.retail_id.create_retail_id_associate(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_retail_id_generate(self, body: Optional['RetailIdCreateGenerateRequest'] = None, license_number: Optional[str] = None) -> List['WriteResponse']:
        """
        Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        - WebApi Retail ID Read Write State (All or WriteOnly)
        - Industry/View Packages
        - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
        POST /retailid/v2/generate
        Parameters:
        - body ('RetailIdCreateGenerateRequest'): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.retail_id.create_retail_id_generate(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_retail_id_merge(self, body: Optional['RetailIdCreateMergeRequest'] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        - WebApi Retail ID Read Write State (All or WriteOnly)
        - Key Value Settings/Retail ID Merge Packages Enabled
        POST /retailid/v2/merge
        Parameters:
        - body ('RetailIdCreateMergeRequest'): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.retail_id.create_retail_id_merge(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_retail_id_packages_info(self, body: Optional['RetailIdCreatePackagesInfoRequest'] = None, license_number: Optional[str] = None) -> List['WriteResponse']:
        """
        Retrieves Package information for given list of Package labels.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
        - WebApi Retail ID Read Write State (All or WriteOnly)
        - Industry/View Packages
        - Admin/Employees/Packages Page/Product Labels(Manage)
        POST /retailid/v2/packages/info
        Parameters:
        - body ('RetailIdCreatePackagesInfoRequest'): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.retail_id.create_retail_id_packages_info(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_retail_id_allotment(self, body: Any = None, license_number: Optional[str] = None) -> List['Allotment']:
        """
        Retrieves the available Retail Item ID quota for a facility.
        Permissions Required:
        - Download Product Labels
        - Manage Product Labels
        - Manage Tag Orders
        GET /retailid/v2/allotment
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.retail_id.get_retail_id_allotment(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_retail_id_receive_by_label(self, label: str, body: Any = None, license_number: Optional[str] = None) -> List['Receive']:
        """
        Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Manage RetailId
        - WebApi Retail ID Read Write State (All or ReadOnly)
        - Industry/View Packages
        - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
        GET /retailid/v2/receive/{label}
        Parameters:
        - label (str): Path parameter label
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.retail_id.get_retail_id_receive_by_label(label, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_retail_id_receive_qr_by_short_code(self, short_code: str, body: Any = None, license_number: Optional[str] = None) -> List['ReceiveQrByShortCode']:
        """
        Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
        Permissions Required:
        - External Sources(ThirdPartyVendorV2)/Manage RetailId
        - WebApi Retail ID Read Write State (All or ReadOnly)
        - Industry/View Packages
        - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
        GET /retailid/v2/receive/qr/{shortCode}
        Parameters:
        - short_code (str): Path parameter shortCode
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.retail_id.get_retail_id_receive_qr_by_short_code(short_code, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

