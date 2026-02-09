from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreatePlantsWasteRequestItem(TypedDict, total=False):
    LocationName: str
    MixedMaterial: str
    Note: str
    PlantLabels: List[Any]
    ReasonName: str
    SublocationName: str
    UnitOfMeasureName: str
    WasteDate: str
    WasteMethodName: str
    WasteWeight: float
