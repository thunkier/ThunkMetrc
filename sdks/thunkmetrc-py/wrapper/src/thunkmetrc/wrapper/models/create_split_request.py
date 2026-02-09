from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateSplitRequest(TypedDict, total=False):
    ActualDate: str
    Count: int
    GroupName: str
    Location: str
    PatientLicenseNumber: str
    PlantBatch: str
    Strain: str
    Sublocation: str
