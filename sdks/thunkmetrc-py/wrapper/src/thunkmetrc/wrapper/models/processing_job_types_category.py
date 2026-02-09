from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class ProcessingJobTypesCategory(TypedDict, total=False):
    ForItems: bool
    ForProcessingJobs: bool
    Id: int
    Name: str
    RequiresProcessingJobAttributes: bool
