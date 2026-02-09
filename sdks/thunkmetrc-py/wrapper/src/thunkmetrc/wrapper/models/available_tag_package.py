from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class AvailableTagPackage(TypedDict, total=False):
    FacilityId: int
    GroupTagTypeId: str
    GroupTagTypeName: str
    Id: int
    Label: str
    MaxGroupSize: int
    TagInventoryTypeName: str
    TagTypeId: str
    TagTypeName: str
