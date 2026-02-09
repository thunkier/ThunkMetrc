from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class ProcessingJobCreateJobTypesRequest(TypedDict, total=False):
    Attributes: List[Any]
    Category: str
    Description: str
    Name: str
    ProcessingSteps: str
