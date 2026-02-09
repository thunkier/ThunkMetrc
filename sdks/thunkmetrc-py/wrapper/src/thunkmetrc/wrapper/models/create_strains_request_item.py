from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateStrainsRequestItem(TypedDict, total=False):
    CbdLevel: float
    IndicaPercentage: int
    Name: str
    SativaPercentage: int
    TestingStatus: str
    ThcLevel: float
