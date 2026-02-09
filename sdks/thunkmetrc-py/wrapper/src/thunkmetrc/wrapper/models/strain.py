from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Strain(TypedDict, total=False):
    CbdLevel: str
    Genetics: str
    Id: int
    IndicaPercentage: int
    IsUsed: bool
    Name: str
    SativaPercentage: int
    TestingStatus: str
    ThcLevel: str
