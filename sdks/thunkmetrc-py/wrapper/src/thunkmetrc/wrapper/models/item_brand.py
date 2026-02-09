from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class ItemBrand(TypedDict, total=False):
    Id: int
    Name: str
    Status: str
