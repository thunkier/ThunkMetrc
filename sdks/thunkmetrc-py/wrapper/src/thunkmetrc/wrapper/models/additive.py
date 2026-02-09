from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Additive(TypedDict, total=False):
    AdditiveTypeName: str
    AmountUnitOfMeasure: str
    ApplicationDevice: str
    EpaRegistrationNumber: str
    Note: str
    PlantBatchId: int
    PlantBatchName: str
    PlantCount: int
    ProductSupplier: str
    ProductTradeName: str
    Rate: str
    RestrictiveEntryIntervalQuantityDescription: str
    RestrictiveEntryIntervalTimeDescription: str
    TotalAmountApplied: int
    Volume: float
