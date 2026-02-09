from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Driver(TypedDict, total=False):
    DriversLicenseNumber: str
    EmployeeId: str
    FacilityId: int
    Id: int
    IsArchived: bool
    LastModified: str
    Name: str
