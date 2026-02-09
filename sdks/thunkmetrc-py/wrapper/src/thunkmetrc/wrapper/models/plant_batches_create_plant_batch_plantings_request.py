from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class PlantBatchesCreatePlantBatchPlantingsRequest(TypedDict, total=False):
    ActualDate: str
    Count: int
    Location: str
    Name: str
    PatientLicenseNumber: str
    SourcePlantBatches: str
    Strain: str
    Sublocation: str
    Type: str
