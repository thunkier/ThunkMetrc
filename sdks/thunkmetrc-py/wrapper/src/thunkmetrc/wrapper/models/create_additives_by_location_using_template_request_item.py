from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class CreateAdditivesByLocationUsingTemplateRequestItem(TypedDict, total=False):
    ActualDate: str
    AdditivesTemplateName: str
    LocationName: str
    Rate: str
    SublocationName: str
    TotalAmountApplied: int
    TotalAmountUnitOfMeasure: str
    Volume: str
