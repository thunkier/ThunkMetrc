from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class LabTestBatche(TypedDict, total=False):
    Id: int
    ItemCategories: List['LabTestsLabTestBatcheItemCategoriesItem']
    ItemCategoryCount: int
    LabTestTypeCount: int
    LabTestTypes: List['LabTestsLabTestBatcheLabTestTypesItem']
    Name: str
    RequiresAllFromLabTestBatch: bool
