from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Adjustment(TypedDict, total=False):
    AdjustedByUserName: str
    AdjustmentDate: str
    AdjustmentNote: str
    AdjustmentQuantity: float
    AdjustmentReasonName: str
    AdjustmentUnitOfMeasureAbbreviation: str
    AdjustmentUnitOfMeasureName: str
    ItemCategoryName: str
    ItemName: str
    PackageId: int
    PackageLabTestResultExpirationDateTime: str
    PackageLabel: str
    PackagedDate: str
