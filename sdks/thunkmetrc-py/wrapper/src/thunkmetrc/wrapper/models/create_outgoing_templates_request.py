from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateOutgoingTemplatesRequest(TypedDict, total=False):
    Destinations: List[Any]
    DriverLicenseNumber: str
    DriverName: str
    DriverOccupationalLicenseNumber: str
    Name: str
    PhoneNumberForQuestions: str
    TransporterFacilityLicenseNumber: str
    VehicleLicensePlateNumber: str
    VehicleMake: str
    VehicleModel: str
