from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class AdditivesTemplate(TypedDict, total=False):
    ActiveIngredients: List['AdditivesTemplateActiveIngredientsItem']
    AdditiveType: str
    AdditiveTypeName: str
    ApplicationDevice: str
    EpaRegistrationNumber: str
    FacilityId: int
    Id: int
    IsActive: bool
    LastModified: str
    Name: str
    Note: str
    ProductSupplier: str
    ProductTradeName: str
    RestrictiveEntryIntervalQuantityDescription: str
    RestrictiveEntryIntervalTimeDescription: str
