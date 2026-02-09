from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class TransferTemplateDeliveryTransporterDetail(TypedDict, total=False):
    ActualDriverStartDateTime: str
    DriverLayoverLeg: str
    DriverName: str
    DriverOccupationalLicenseNumber: str
    DriverVehicleLicenseNumber: str
    VehicleLicensePlateNumber: str
    VehicleMake: str
    VehicleModel: str
