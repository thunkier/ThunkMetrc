from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Hub(TypedDict, total=False):
    ActualArrivalDateTime: str
    ActualDepartureDateTime: str
    ActualReturnArrivalDateTime: str
    ActualReturnDepartureDateTime: str
    CreatedByUserName: str
    CreatedDateTime: str
    DeliveryCount: int
    DeliveryId: int
    DeliveryPackageCount: int
    DeliveryReceivedPackageCount: int
    DriverName: str
    DriverOccupationalLicenseNumber: str
    DriverVehicleLicenseNumber: str
    EstimatedArrivalDateTime: str
    EstimatedDepartureDateTime: str
    EstimatedReturnArrivalDateTime: str
    EstimatedReturnDepartureDateTime: str
    Id: int
    IsLayover: bool
    IsVoided: bool
    LastModified: str
    ManifestNumber: str
    PackageCount: int
    ReceivedDateTime: str
    ReceivedDeliveryCount: int
    ReceivedPackageCount: int
    RecipientFacilityLicenseNumber: str
    RecipientFacilityName: str
    RejectedPackagesReturned: bool
    ShipmentTransactionType: int
    ShipmentTransporterDetails: List['TransfersHubShipmentTransporterDetailsItem']
    ShipmentTypeName: str
    ShipperFacilityLicenseNumber: str
    ShipperFacilityName: str
    TransporterAcceptedDateTime: str
    TransporterActualArrivalDateTime: str
    TransporterActualDepartureDateTime: str
    TransporterEstimatedArrivalDateTime: str
    TransporterEstimatedDepartureDateTime: str
    TransporterFacilityLicenseNumber: str
    TransporterFacilityName: str
    VehicleLicensePlateNumber: str
    VehicleMake: str
    VehicleModel: str
