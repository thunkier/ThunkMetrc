from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateHarvestsLocationRequestItem(TypedDict, total=False):
    ActualDate: str
    DryingLocation: str
    DryingSublocation: str
    HarvestName: str
    Id: int
