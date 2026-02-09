from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateStrainsRequestItem(TypedDict, total=False):
    CbdLevel: float
    Id: int
    IndicaPercentage: int
    Name: str
    SativaPercentage: int
    TestingStatus: str
    ThcLevel: float
