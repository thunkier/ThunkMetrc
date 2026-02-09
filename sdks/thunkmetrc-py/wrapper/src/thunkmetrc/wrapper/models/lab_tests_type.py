from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class LabTestsType(TypedDict, total=False):
    AlwaysPasses: bool
    DependencyMode: str
    Id: int
    InformationalOnly: bool
    LabTestResultExpirationDays: int
    LabTestResultMaximum: int
    LabTestResultMinimum: int
    LabTestResultMode: str
    MaxAllowedFailureCount: int
    Name: str
    RequiresTestResult: bool
    ResearchAndDevelopment: bool
