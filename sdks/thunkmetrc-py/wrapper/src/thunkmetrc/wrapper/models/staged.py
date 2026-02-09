from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Staged(TypedDict, total=False):
    CommissionedDateTime: str
    DetachedDateTime: str
    FacilityId: int
    Id: int
    IsArchived: bool
    IsStaged: bool
    IsUsed: bool
    Label: str
    LastModified: str
    MaxGroupSize: int
    ProductLabel: str
    QrCount: int
    StatusName: str
    TagInventoryTypeName: str
    TagTypeId: int
    TagTypeName: str
    UsedDateTime: str
