from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class SalesUpdateSaleDeliveriesRetailerRequest(TypedDict, total=False):
    DateTime: str
    Destinations: List[Any]
    DriverEmployeeId: str
    DriverName: str
    DriversLicenseNumber: str
    EstimatedDepartureDateTime: str
    Id: int
    Packages: List[Any]
    PhoneNumberForQuestions: str
    VehicleLicensePlateNumber: str
    VehicleMake: str
    VehicleModel: str
