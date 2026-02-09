from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class PlantsCreatePlantAdditivesUsingTemplateRequest(TypedDict, total=False):
    ActualDate: str
    AdditivesTemplateName: str
    PlantLabels: List[Any]
    Rate: str
    TotalAmountApplied: int
    TotalAmountUnitOfMeasure: str
    Volume: str
