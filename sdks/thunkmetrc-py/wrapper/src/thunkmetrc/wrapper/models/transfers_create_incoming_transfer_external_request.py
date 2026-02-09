from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class TransfersCreateIncomingTransferExternalRequest(TypedDict, total=False):
    Destinations: List[Any]
    DriverLicenseNumber: str
    DriverName: str
    DriverOccupationalLicenseNumber: str
    PhoneNumberForQuestions: str
    ShipperAddress1: str
    ShipperAddress2: str
    ShipperAddressCity: str
    ShipperAddressPostalCode: str
    ShipperAddressState: str
    ShipperLicenseNumber: str
    ShipperMainPhoneNumber: str
    ShipperName: str
    TransporterFacilityLicenseNumber: str
    VehicleLicensePlateNumber: str
    VehicleMake: str
    VehicleModel: str
