from typing import Any, Optional, List, AsyncGenerator
from thunkmetrc.client import MetrcClient
from ..ratelimiter import MetrcRateLimiter
from ..models import *
from ..utils import iterate_pages, deserialize_page

class TransportersService:
    def __init__(self, client: MetrcClient, limiter: MetrcRateLimiter):
        self.client = client
        self.limiter = limiter

    async def create_transporters_drivers(self, body: Optional[List['TransportersCreateDriversRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates new driver records for a Facility.
        Permissions Required:
        - Manage Transporters
        POST /transporters/v2/drivers
        Parameters:
        - body (List['TransportersCreateDriversRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transporters.create_transporters_drivers(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def create_transporters_vehicles(self, body: Optional[List['TransportersCreateVehiclesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Creates new vehicle records for a Facility.
        Permissions Required:
        - Manage Transporters
        POST /transporters/v2/vehicles
        Parameters:
        - body (List['TransportersCreateVehiclesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transporters.create_transporters_vehicles(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_transporters_driver_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
        Permissions Required:
        - Manage Transporters
        DELETE /transporters/v2/drivers/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transporters.delete_transporters_driver_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def delete_transporters_vehicle_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> Any:
        """
        Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
        Permissions Required:
        - Manage Transporters
        DELETE /transporters/v2/vehicles/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transporters.delete_transporters_vehicle_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def get_transporters_driver_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'TransportersDriver':
        """
        Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
        Permissions Required:
        - Transporters
        GET /transporters/v2/drivers/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transporters.get_transporters_driver_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transporters_drivers(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['TransportersDriver']:
        """
        Retrieves a list of drivers for a Facility.
        Permissions Required:
        - Transporters
        GET /transporters/v2/drivers
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transporters.get_transporters_drivers(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transporters_vehicle_by_id(self, id: str, body: Any = None, license_number: Optional[str] = None) -> 'TransportersVehicle':
        """
        Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
        Permissions Required:
        - Transporters
        GET /transporters/v2/vehicles/{id}
        Parameters:
        - id (str): Path parameter id
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transporters.get_transporters_vehicle_by_id(id, body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, True, op)

    async def get_transporters_vehicles(self, body: Any = None, license_number: Optional[str] = None, page_number: Optional[str] = None, page_size: Optional[str] = None) -> PaginatedResponse['TransportersVehicle']:
        """
        Retrieves a list of vehicles for a Facility.
        Permissions Required:
        - Transporters
        GET /transporters/v2/vehicles
        Parameters:
        - license_number (str, optional): Filter by licenseNumber
        - page_number (str, optional): Filter by pageNumber
        - page_size (str, optional): Filter by pageSize
        """
        async def op():
            return await self.client.transporters.get_transporters_vehicles(body=body,
                license_number=license_number,
                page_number=page_number,
                page_size=page_size,
            )
        return await self.limiter.execute(None, True, op)

    async def update_transporters_drivers(self, body: Optional[List['TransportersUpdateDriversRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates existing driver records for a Facility.
        Permissions Required:
        - Manage Transporters
        PUT /transporters/v2/drivers
        Parameters:
        - body (List['TransportersUpdateDriversRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transporters.update_transporters_drivers(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

    async def update_transporters_vehicles(self, body: Optional[List['TransportersUpdateVehiclesRequest']] = None, license_number: Optional[str] = None) -> 'WriteResponse':
        """
        Updates existing vehicle records for a facility.
        Permissions Required:
        - Manage Transporters
        PUT /transporters/v2/vehicles
        Parameters:
        - body (List['TransportersUpdateVehiclesRequest']): Request body
        - license_number (str, optional): Filter by licenseNumber
        """
        async def op():
            return await self.client.transporters.update_transporters_vehicles(body=body,
                license_number=license_number,
            )
        return await self.limiter.execute(None, False, op)

