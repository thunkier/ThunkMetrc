from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateRenameRequest(TypedDict, total=False):
    Id: int
    NewName: str
    OldName: str
