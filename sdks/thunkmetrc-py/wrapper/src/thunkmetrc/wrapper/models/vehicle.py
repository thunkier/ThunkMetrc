from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Vehicle(TypedDict, total=False):
    FacilityId: int
    Id: int
    IsArchived: bool
    LastModified: str
    LicensePlateNumber: str
    Make: str
    Model: str
