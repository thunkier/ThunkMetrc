from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class Result(TypedDict, total=False):
    ExpirationDateTime: str
    LabFacilityLicenseNumber: str
    LabFacilityName: str
    LabTestDetailRevokedDate: str
    LabTestResultDocumentFileId: int
    LabTestResultId: int
    OverallPassed: bool
    PackageId: int
    ProductCategoryName: str
    ProductName: str
    ResultReleaseDateTime: str
    ResultReleased: bool
    RevokedDate: str
    SourcePackageLabel: str
    TestComment: str
    TestInformationalOnly: bool
    TestPassed: bool
    TestPerformedDate: str
    TestResultLevel: float
    TestTypeName: str
