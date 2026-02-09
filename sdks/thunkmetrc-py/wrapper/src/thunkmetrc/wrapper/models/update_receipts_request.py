from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateReceiptsRequest(TypedDict, total=False):
    CaregiverLicenseNumber: str
    ExternalReceiptNumber: str
    Id: int
    IdentificationMethod: str
    PatientLicenseNumber: str
    PatientRegistrationLocationId: int
    SalesCustomerType: str
    SalesDateTime: str
    Transactions: List[Any]
