from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateDeliveriesHubRequestItem(TypedDict, total=False):
    DriverEmployeeId: str
    DriverName: str
    DriversLicenseNumber: str
    EstimatedArrivalDateTime: str
    EstimatedDepartureDateTime: str
    Id: int
    PhoneNumberForQuestions: str
    PlannedRoute: str
    TransporterFacilityId: str
    VehicleLicensePlateNumber: str
    VehicleMake: str
    VehicleModel: str
