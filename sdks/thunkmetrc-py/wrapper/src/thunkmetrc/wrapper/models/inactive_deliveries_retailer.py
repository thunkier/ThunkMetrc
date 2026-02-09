from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class InactiveDeliveriesRetailer(TypedDict, total=False):
    AcceptedDateTime: str
    ActualDepartureDateTime: str
    AllowFullEdit: bool
    CompletedDateTime: str
    DateTime: str
    Destinations: List[Any]
    Direction: str
    DriverEmployeeId: str
    DriverName: str
    DriversLicenseNumber: str
    EstimatedDepartureDateTime: str
    FacilityLicenseNumber: str
    FacilityName: str
    Id: int
    LastModified: str
    Leg: int
    Packages: List[Any]
    RecordedByUserName: str
    RecordedDateTime: str
    RestockDateTime: str
    RetailerDeliveryNumber: str
    RetailerDeliveryState: str
    TotalPackages: int
    TotalPrice: int
    TotalPriceSold: int
    VehicleInfo: str
    VehicleLicensePlateNumber: str
    VehicleMake: str
    VehicleModel: str
    VoidedDate: str
