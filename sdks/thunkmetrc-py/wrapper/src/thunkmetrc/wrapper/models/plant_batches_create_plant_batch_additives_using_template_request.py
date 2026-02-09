from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class PlantBatchesCreatePlantBatchAdditivesUsingTemplateRequest(TypedDict, total=False):
    ActualDate: str
    AdditivesTemplateName: str
    PlantBatchName: str
    Rate: str
    TotalAmountApplied: int
    TotalAmountUnitOfMeasure: str
    Volume: str
