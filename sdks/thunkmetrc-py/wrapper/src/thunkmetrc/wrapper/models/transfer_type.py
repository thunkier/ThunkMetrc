from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class TransferType(TypedDict, total=False):
    BypassApproval: bool
    ExternalIncomingCanRecordExternalIdentifier: bool
    ExternalIncomingExternalIdentifierRequired: bool
    ExternalOutgoingCanRecordExternalIdentifier: bool
    ExternalOutgoingExternalIdentifierRequired: bool
    ForExternalIncomingShipments: bool
    ForExternalOutgoingShipments: bool
    ForLicensedShipments: bool
    Name: str
    RequiresDestinationGrossWeight: bool
    RequiresInvoiceNumber: bool
    RequiresPDFDocument: bool
    RequiresPackagesGrossWeight: bool
    TransactionType: str
