from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class PlantBatchesCreatePlantBatchGrowthPhaseRequest(TypedDict, total=False):
    Count: int
    GrowthDate: str
    GrowthPhase: str
    Name: str
    NewLocation: str
    NewSublocation: str
    PatientLicenseNumber: str
    StartingTag: str
