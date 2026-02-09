from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Receive(TypedDict, total=False):
    ChildTag: str
    Eaches: List[str]
    LabelSource: str
    QrCount: int
    Ranges: List[List[int]]
    RequiresVerification: bool
    SiblingTags: List[str]
