from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateAdjustProcessingJobRequestItem(TypedDict, total=False):
    AdjustmentDate: str
    AdjustmentNote: str
    AdjustmentReason: str
    CountUnitOfMeasureName: str
    Id: int
    Packages: List[Any]
    VolumeUnitOfMeasureName: str
    WeightUnitOfMeasureName: str
