from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateGrowthPhaseRequestItem(TypedDict, total=False):
    GrowthDate: str
    GrowthPhase: str
    Id: int
    Label: str
    NewLocation: str
    NewSublocation: str
    NewTag: str
