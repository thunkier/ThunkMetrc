from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateSalesDeliveriesRetailerRequest(TypedDict, total=False):
    ConsumerId: int
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
    RetailerDeliveryId: int
    SalesCustomerType: str
    SalesDateTime: str
    Transactions: List[Any]
