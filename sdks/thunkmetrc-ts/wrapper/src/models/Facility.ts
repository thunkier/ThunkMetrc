
import type { FacilityFacilityType } from './FacilityFacilityType';
import type { FacilityLicense } from './FacilityLicense';
export interface Facility {
    Alias?: string;
    AllowTransferOfOnHoldPackages?: boolean;
    AllowTransferOfOnRecallPackages?: boolean;
    CredentialedDate?: string;
    DisplayName?: string;
    Email?: string;
    FacilityId?: number;
    FacilityType?: FacilityFacilityType;
    HireDate?: string;
    IsFinancialContact?: boolean;
    IsManager?: boolean;
    IsOwner?: boolean;
    License?: FacilityLicense;
    Name?: string;
    Occupations?: any[];
    SupportActivationDate?: string;
    SupportExpirationDate?: string;
}
