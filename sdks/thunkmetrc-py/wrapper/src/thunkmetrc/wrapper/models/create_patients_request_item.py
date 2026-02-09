from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreatePatientsRequestItem(TypedDict, total=False):
    ActualDate: str
    ConcentrateOuncesAllowed: int
    FlowerOuncesAllowed: int
    HasSalesLimitExemption: bool
    InfusedOuncesAllowed: int
    LicenseEffectiveEndDate: str
    LicenseEffectiveStartDate: str
    LicenseNumber: str
    MaxConcentrateThcPercentAllowed: int
    MaxFlowerThcPercentAllowed: int
    RecommendedPlants: int
    RecommendedSmokableQuantity: float
    ThcOuncesAllowed: int
