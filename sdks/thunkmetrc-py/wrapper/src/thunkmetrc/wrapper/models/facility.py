from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Facility(TypedDict, total=False):
    Alias: str
    AllowTransferOfOnHoldPackages: bool
    AllowTransferOfOnRecallPackages: bool
    CredentialedDate: str
    DisplayName: str
    Email: str
    FacilityId: int
    FacilityType: 'FacilityFacilityType'
    HireDate: str
    IsFinancialContact: bool
    IsManager: bool
    IsOwner: bool
    License: 'FacilityLicense'
    Name: str
    Occupations: List[Any]
    SupportActivationDate: str
    SupportExpirationDate: str
