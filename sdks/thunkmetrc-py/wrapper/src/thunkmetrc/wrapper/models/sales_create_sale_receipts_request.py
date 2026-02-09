from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class SalesCreateSaleReceiptsRequest(TypedDict, total=False):
    CaregiverLicenseNumber: str
    ExternalReceiptNumber: str
    IdentificationMethod: str
    PatientLicenseNumber: str
    PatientRegistrationLocationId: int
    SalesCustomerType: str
    SalesDateTime: str
    Transactions: List[Any]
