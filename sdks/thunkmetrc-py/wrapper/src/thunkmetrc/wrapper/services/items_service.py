from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class ItemsService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_items_brand(self, body: Optional[List['ItemsCreateBrandRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates one or more new item brands for the specified Facility identified by the License Number.
        Permissions Required:
        - Manage Items
        POST /items/v2/brand
        Parameters:
        - body (List['ItemsCreateBrandRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.items.create_items_brand(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_items_file(self, body: Optional[List['ItemsCreateFileRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
        Permissions Required:
        - Manage Items
        POST /items/v2/file
        Parameters:
        - body (List['ItemsCreateFileRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.items.create_items_file(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_items(self, body: Optional[List['ItemsCreateItemsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Permissions Required:
        - Manage Items
        POST /items/v2
        Parameters:
        - body (List['ItemsCreateItemsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.items.create_items(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_items_photo(self, body: Optional[List['ItemsCreatePhotoRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
        Permissions Required:
        - Manage Items
        POST /items/v2/photo
        Parameters:
        - body (List['ItemsCreatePhotoRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.items.create_items_photo(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_items_brand_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Archives the specified Item Brand by Id for the given Facility License Number.
        Permissions Required:
        - Manage Items
        DELETE /items/v2/brand/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.items.delete_items_brand_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_items_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Archives the specified Product by Id for the given Facility License Number.
        Permissions Required:
        - Manage Items
        DELETE /items/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.items.delete_items_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_active_items(self, body: Any = None, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Item']:
        """
        Returns a list of active items for the specified Facility.
        Permissions Required:
        - Manage Items
        GET /items/v2/active
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.items.get_active_items(body=body,
                last_modified_end=last_modified_end,
                last_modified_start=last_modified_start,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_items_brands(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Brand']:
        """
        Retrieves a list of active item brands for the specified Facility.
        Permissions Required:
        - Manage Items
        GET /items/v2/brands
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.items.get_items_brands(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_items_categories(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Category']:
        """
        Retrieves a list of item categories.
        GET /items/v2/categories
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.items.get_items_categories(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_items_file_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'File':
        """
        Retrieves a file by its Id for the specified Facility.
        Permissions Required:
        - Manage Items
        GET /items/v2/file/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.items.get_items_file_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_inactive_items(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['Item']:
        """
        Retrieves a list of inactive items for the specified Facility.
        Permissions Required:
        - Manage Items
        GET /items/v2/inactive
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.items.get_inactive_items(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_items_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'Item':
        """
        Retrieves detailed information about a specific Item by Id.
        Permissions Required:
        - Manage Items
        GET /items/v2/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.items.get_items_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_items_photo_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'Photo':
        """
        Retrieves an image by its Id for the specified Facility.
        Permissions Required:
        - Manage Items
        GET /items/v2/photo/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.items.get_items_photo_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def update_items_brand(self, body: Optional[List['ItemsUpdateBrandRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates one or more existing item brands for the specified Facility.
        Permissions Required:
        - Manage Items
        PUT /items/v2/brand
        Parameters:
        - body (List['ItemsUpdateBrandRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.items.update_items_brand(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_items(self, body: Optional[List['ItemsUpdateItemsRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates one or more existing products for the specified Facility.
        Permissions Required:
        - Manage Items
        PUT /items/v2
        Parameters:
        - body (List['ItemsUpdateItemsRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.items.update_items(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

