from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class InactiveReceipt(TypedDict, total=False):
    ArchivedDate: str
    CaregiverLicenseNumber: str
    ExternalReceiptNumber: str
    Id: int
    IdentificationMethod: str
    IsFinal: bool
    LastModified: str
    PatientLicenseNumber: str
    PatientRegistrationLocationId: int
    ReceiptNumber: str
    RecordedByUserName: str
    RecordedDateTime: str
    SalesCustomerType: str
    SalesDateTime: str
    TotalPackages: int
    TotalPrice: int
    Transactions: List[Any]
