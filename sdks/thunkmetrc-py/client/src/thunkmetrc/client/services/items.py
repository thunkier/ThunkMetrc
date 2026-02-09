from typing import Any, Optional, Dict, List, Union
import urllib.parse

class ItemsClient:
    def __init__(self, client):
        self.client = client
    def create_items_brand(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Creates one or more new item brands for the specified Facility identified by the License Number.
        Permissions Required:
        - Manage Items
        POST CreateBrand
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/items/v2/brand"
        return self.client._send("POST", path, body, query_params)
    def create_items_file(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
        Permissions Required:
        - Manage Items
        POST CreateFile
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/items/v2/file"
        return self.client._send("POST", path, body, query_params)
    def create_items(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Permissions Required:
        - Manage Items
        POST CreateItems
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/items/v2"
        return self.client._send("POST", path, body, query_params)
    def create_items_photo(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
        Permissions Required:
        - Manage Items
        POST CreatePhoto
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/items/v2/photo"
        return self.client._send("POST", path, body, query_params)
    def delete_items_brand_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Archives the specified Item Brand by Id for the given Facility License Number.
        Permissions Required:
        - Manage Items
        DELETE DeleteBrandById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/items/v2/brand/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def delete_items_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Archives the specified Product by Id for the given Facility License Number.
        Permissions Required:
        - Manage Items
        DELETE DeleteItemsById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/items/v2/{urllib.parse.quote(id)}"
        return self.client._send("DELETE", path, body, query_params)
    def get_active_items(self, last_modified_end: Optional[str] = None, last_modified_start: Optional[str] = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Returns a list of active items for the specified Facility.
        Permissions Required:
        - Manage Items
        GET GetActiveItems
        Parameters:
        - last_modified_end (str, optional): Filter by lastModifiedEnd
        - last_modified_start (str, optional): Filter by lastModifiedStart
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if last_modified_end is not None:
            query_params["lastModifiedEnd"] = last_modified_end
        if last_modified_start is not None:
            query_params["lastModifiedStart"] = last_modified_start
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/items/v2/active"
        return self.client._send("GET", path, body, query_params)
    def get_items_brands(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of active item brands for the specified Facility.
        Permissions Required:
        - Manage Items
        GET GetBrands
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/items/v2/brands"
        return self.client._send("GET", path, body, query_params)
    def get_items_categories(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of item categories.
        GET GetCategories
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/items/v2/categories"
        return self.client._send("GET", path, body, query_params)
    def get_items_file_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves a file by its Id for the specified Facility.
        Permissions Required:
        - Manage Items
        GET GetFileById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/items/v2/file/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def get_inactive_items(self, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> Any:
        """
        Retrieves a list of inactive items for the specified Facility.
        Permissions Required:
        - Manage Items
        GET GetInactiveItems
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        if page_number is not None:
            query_params["pageNumber"] = page_number
        if page_size is not None:
            query_params["pageSize"] = page_size
        path = f"/items/v2/inactive"
        return self.client._send("GET", path, body, query_params)
    def get_items_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves detailed information about a specific Item by Id.
        Permissions Required:
        - Manage Items
        GET GetItemsById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/items/v2/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def get_items_photo_by_id(self, id: str, license_number: Optional[str] = None) -> Any:
        """
        Retrieves an image by its Id for the specified Facility.
        Permissions Required:
        - Manage Items
        GET GetPhotoById
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/items/v2/photo/{urllib.parse.quote(id)}"
        return self.client._send("GET", path, body, query_params)
    def update_items_brand(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates one or more existing item brands for the specified Facility.
        Permissions Required:
        - Manage Items
        PUT UpdateBrand
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/items/v2/brand"
        return self.client._send("PUT", path, body, query_params)
    def update_items(self, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Updates one or more existing products for the specified Facility.
        Permissions Required:
        - Manage Items
        PUT UpdateItems
        Parameters:
        - body (Any): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        query_params = {}
        if license_number is not None:
            query_params["licenseNumber"] = license_number
        path = f"/items/v2"
        return self.client._send("PUT", path, body, query_params)

