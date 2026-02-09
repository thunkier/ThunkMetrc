from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateAdjustPlantsRequest(TypedDict, total=False):
    AdjustCount: int
    AdjustReason: str
    AdjustmentDate: str
    Id: int
    Label: str
    ReasonNote: str
