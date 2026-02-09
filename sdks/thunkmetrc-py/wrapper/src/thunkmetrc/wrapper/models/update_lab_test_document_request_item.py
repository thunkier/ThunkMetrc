from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateLabTestDocumentRequestItem(TypedDict, total=False):
    DocumentFileBase64: str
    DocumentFileName: str
    LabTestResultId: int
