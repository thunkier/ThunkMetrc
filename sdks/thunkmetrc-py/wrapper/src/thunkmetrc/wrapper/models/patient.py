from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Patient(TypedDict, total=False):
    HasSalesLimitExemption: bool
    LicenseEffectiveEndDate: str
    LicenseEffectiveStartDate: str
    LicenseNumber: str
    OtherFacilitiesCount: int
    PatientId: int
    RecommendedPlants: int
    RecommendedSmokableQuantity: float
    RegistrationDate: str
