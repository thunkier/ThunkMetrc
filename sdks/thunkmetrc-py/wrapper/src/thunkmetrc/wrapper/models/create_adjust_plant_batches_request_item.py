from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateAdjustPlantBatchesRequestItem(TypedDict, total=False):
    AdjustmentDate: str
    AdjustmentReason: str
    PlantBatchName: str
    Quantity: float
    ReasonNote: str
