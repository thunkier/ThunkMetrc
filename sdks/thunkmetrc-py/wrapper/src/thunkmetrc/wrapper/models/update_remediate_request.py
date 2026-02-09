from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateRemediateRequest(TypedDict, total=False):
    PackageLabel: str
    RemediationDate: str
    RemediationMethodName: str
    RemediationSteps: str
