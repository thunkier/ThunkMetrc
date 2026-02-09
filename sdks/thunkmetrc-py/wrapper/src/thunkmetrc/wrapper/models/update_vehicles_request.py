from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateVehiclesRequest(TypedDict, total=False):
    Id: str
    LicensePlateNumber: str
    Make: str
    Model: str
