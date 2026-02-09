from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class SalesCreateSaleDeliveriesRequest(TypedDict, total=False):
    ConsumerId: int
    DriverEmployeeId: str
    DriverName: str
    DriversLicenseNumber: str
    EstimatedArrivalDateTime: str
    EstimatedDepartureDateTime: str
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
    RecipientZoneId: int
    SalesCustomerType: str
    SalesDateTime: str
    Transactions: List[Any]
    TransporterFacilityLicenseNumber: str
    VehicleLicensePlateNumber: str
    VehicleMake: str
    VehicleModel: str
