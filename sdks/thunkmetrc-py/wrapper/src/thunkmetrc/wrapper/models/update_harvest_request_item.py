from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateHarvestRequestItem(TypedDict, total=False):
    ActualDate: str
    DryingLocation: str
    DryingSublocation: str
    HarvestName: str
    PatientLicenseNumber: str
    Plant: str
    UnitOfWeight: str
    Weight: float
