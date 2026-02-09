from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdatePlantsStrainRequest(TypedDict, total=False):
    Id: int
    Label: str
    StrainId: int
    StrainName: str
