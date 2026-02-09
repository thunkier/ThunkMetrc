from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateNameRequestItem(TypedDict, total=False):
    Group: str
    Id: int
    NewGroup: str
