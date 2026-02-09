from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateAdditivesTemplatesRequestItem(TypedDict, total=False):
    ActiveIngredients: List[Any]
    AdditiveType: str
    ApplicationDevice: str
    EpaRegistrationNumber: str
    Id: int
    Name: str
    Note: str
    ProductSupplier: str
    ProductTradeName: str
    RestrictiveEntryIntervalQuantityDescription: str
    RestrictiveEntryIntervalTimeDescription: str
