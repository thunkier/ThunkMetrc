from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class PlantWaste(TypedDict, total=False):
    GrowthPhase: int
    Label: str
    LocationId: int
    LocationName: str
    PlantId: int
    PlantWasteId: int
    StateName: str
    SublocationId: int
    SublocationName: str
    WasteAmount: int
    WasteUnitOfMeasureAbbreviation: str
