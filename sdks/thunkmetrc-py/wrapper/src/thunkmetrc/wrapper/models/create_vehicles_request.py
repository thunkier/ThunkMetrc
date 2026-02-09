from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateVehiclesRequest(TypedDict, total=False):
    LicensePlateNumber: str
    Make: str
    Model: str
