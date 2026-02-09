from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreatePlantBatchesWasteRequestItem(TypedDict, total=False):
    MixedMaterial: str
    Note: str
    PlantBatchName: str
    ReasonName: str
    UnitOfMeasureName: str
    WasteDate: str
    WasteMethodName: str
    WasteWeight: float
