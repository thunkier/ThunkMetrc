from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreatePlantBatchesAdditivesRequest(TypedDict, total=False):
    ActiveIngredients: List[Any]
    ActualDate: str
    AdditiveType: str
    ApplicationDevice: str
    EpaRegistrationNumber: str
    PlantBatchName: str
    ProductSupplier: str
    ProductTradeName: str
    TotalAmountApplied: int
    TotalAmountUnitOfMeasure: str
