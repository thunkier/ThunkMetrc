from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Employee(TypedDict, total=False):
    FullName: str
    IsIndustryAdmin: bool
    IsManager: bool
    IsOwner: bool
    License: 'EmployeeLicense'
