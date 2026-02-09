from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateAdjustPackagesRequest(TypedDict, total=False):
    AdjustmentDate: str
    AdjustmentReason: str
    Label: str
    Quantity: int
    ReasonNote: str
    UnitOfMeasure: str
