from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateDeliveriesRequest(TypedDict, total=False):
    ConsumerId: int
    DriverEmployeeId: str
    DriverName: str
    DriversLicenseNumber: str
    EstimatedArrivalDateTime: str
    EstimatedDepartureDateTime: str
    Id: int
    PatientLicenseNumber: str
    PhoneNumberForQuestions: str
    PlannedRoute: str
    RecipientAddressCity: str
    RecipientAddressCounty: str
    RecipientAddressPostalCode: str
    RecipientAddressState: str
    RecipientAddressStreet1: str
    RecipientAddressStreet2: str
    RecipientName: str
    RecipientZoneId: str
    SalesCustomerType: str
    SalesDateTime: str
    Transactions: List[Any]
    TransporterFacilityLicenseNumber: str
    VehicleLicensePlateNumber: str
    VehicleMake: str
    VehicleModel: str
