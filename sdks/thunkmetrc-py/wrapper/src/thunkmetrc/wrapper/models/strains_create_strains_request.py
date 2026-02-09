from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class StrainsCreateStrainsRequest(TypedDict, total=False):
    CbdLevel: float
    IndicaPercentage: float
    Name: str
    SativaPercentage: float
    TestingStatus: str
    ThcLevel: float
