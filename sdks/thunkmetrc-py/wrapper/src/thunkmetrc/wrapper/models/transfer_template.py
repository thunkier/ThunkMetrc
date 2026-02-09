from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class TransferTemplate(TypedDict, total=False):
    ActualArrivalDateTime: str
    ActualDepartureDateTime: str
    ActualReturnArrivalDateTime: str
    ActualReturnDepartureDateTime: str
    ContainsDonation: bool
    ContainsPlantPackage: bool
    ContainsProductPackage: bool
    ContainsProductRequiresRemediation: bool
    ContainsRemediatedProductPackage: bool
    ContainsTestingSample: bool
    ContainsTradeSample: bool
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
    InvoiceNumber: str
    IsVoided: bool
    LastModified: str
    ManifestNumber: str
    Name: str
    OriginatingTemplateId: int
    PackageCount: int
    ReceivedDateTime: str
    ReceivedDeliveryCount: int
    ReceivedPackageCount: int
    RecipientFacilityLicenseNumber: str
    RecipientFacilityName: str
    ShipmentLicenseType: int
    ShipmentTransactionType: str
    ShipmentTypeName: str
    ShipperFacilityLicenseNumber: str
    ShipperFacilityName: str
    TransporterFacilityLicenseNumber: str
    TransporterFacilityName: str
    VehicleLicensePlateNumber: str
    VehicleMake: str
    VehicleModel: str
