from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class InactiveDelivery(TypedDict, total=False):
    AcceptedDateTime: str
    ActualArrivalDateTime: str
    ActualDepartureDateTime: str
    ActualReturnArrivalDateTime: str
    ActualReturnDepartureDateTime: str
    CompletedDateTime: str
    ConsumerId: str
    DeliveryNumber: str
    Direction: str
    DriverEmployeeId: str
    DriverName: str
    DriversLicenseNumber: str
    EstimatedArrivalDateTime: str
    EstimatedDepartureDateTime: str
    EstimatedReturnArrivalDateTime: str
    EstimatedReturnDepartureDateTime: str
    FacilityLicenseNumber: str
    FacilityName: str
    Id: int
    LastModified: str
    PatientLicenseNumber: str
    PlannedRoute: str
    RecipientAddressCity: str
    RecipientAddressCounty: str
    RecipientAddressPostalCode: str
    RecipientAddressState: str
    RecipientAddressStreet1: str
    RecipientAddressStreet2: str
    RecipientDeliveryZoneId: int
    RecipientDeliveryZoneName: str
    RecipientName: str
    RecipientZoneId: int
    RecordedByUserName: str
    RecordedDateTime: str
    SalesCustomerType: str
    SalesDateTime: str
    SalesDeliveryState: str
    ShipperFacilityLicenseNumber: str
    TotalPackages: int
    TotalPrice: int
    Transactions: List[Any]
    TransporterFacilityId: int
    TransporterFacilityLicenseNumber: str
    TransporterFacilityName: str
    VehicleLicensePlateNumber: str
    VehicleMake: str
    VehicleModel: str
    VoidedDate: str
