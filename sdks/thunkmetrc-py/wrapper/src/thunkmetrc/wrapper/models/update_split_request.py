from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateSplitRequest(TypedDict, total=False):
    PlantCount: int
    SourcePlantLabel: str
    SplitDate: str
    StrainLabel: str
    TagLabel: str
