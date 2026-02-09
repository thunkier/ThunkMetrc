from typing import TypedDict, List, Optional, Any, Generic, TypeVar, TYPE_CHECKING
from typing_extensions import TypedDict

if TYPE_CHECKING:
    from . import *

class UpdateWebhooksRequest(TypedDict, total=False):
    errorResponseJsonTemplate: str
    facilityLicenseNumbers: str
    objectType: str
    serverPublicKeyFingerprint: str
    status: str
    template: str
    tpiApiKey: str
    url: str
    userApiKey: str
    verb: str
