from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateAdjustPlantBatchesRequest(TypedDict, total=False):
    AdjustmentDate: str
    AdjustmentReason: str
    PlantBatchName: str
    Quantity: int
    ReasonNote: str
