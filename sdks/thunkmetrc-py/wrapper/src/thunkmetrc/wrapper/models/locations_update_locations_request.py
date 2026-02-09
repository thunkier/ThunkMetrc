from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class LocationsUpdateLocationsRequest(TypedDict, total=False):
    Id: int
    LocationTypeName: str
    Name: str
