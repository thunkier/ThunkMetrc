from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class LocationsType(TypedDict, total=False):
    ForHarvests: bool
    ForPackages: bool
    ForPlantBatches: bool
    ForPlants: bool
    Id: int
    Name: str
