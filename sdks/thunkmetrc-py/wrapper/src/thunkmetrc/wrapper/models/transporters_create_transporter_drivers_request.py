from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class TransportersCreateTransporterDriversRequest(TypedDict, total=False):
    DriversLicenseNumber: str
    EmployeeId: str
    Name: str
