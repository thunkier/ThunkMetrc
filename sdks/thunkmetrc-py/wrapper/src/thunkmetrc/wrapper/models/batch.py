from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Batch(TypedDict, total=False):
    Id: int
    ItemCategories: List['LabTestsBatchItemCategoriesItem']
    ItemCategoryCount: int
    LabTestTypeCount: int
    LabTestTypes: List['LabTestsBatchLabTestTypesItem']
    Name: str
    RequiresAllFromLabTestBatch: bool
