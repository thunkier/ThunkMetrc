from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class ActiveJobType(TypedDict, total=False):
    Attributes: List['ProcessingJobActiveJobTypeAttributesItem']
    CategoryName: str
    Description: str
    ForItems: bool
    ForProcessingJobs: bool
    Id: int
    Name: str
    ProcessingSteps: str
    RequiresProcessingJobAttributes: bool
