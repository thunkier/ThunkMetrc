from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateDecontaminateRequest(TypedDict, total=False):
    DecontaminationDate: str
    DecontaminationMethodName: str
    DecontaminationSteps: str
    PackageLabel: str
