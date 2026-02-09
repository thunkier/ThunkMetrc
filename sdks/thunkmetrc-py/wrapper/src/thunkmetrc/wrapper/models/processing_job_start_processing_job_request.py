from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class ProcessingJobStartProcessingJobRequest(TypedDict, total=False):
    CountUnitOfMeasure: str
    JobName: str
    JobType: str
    Packages: List[Any]
    StartDate: str
    VolumeUnitOfMeasure: str
    WeightUnitOfMeasure: str
