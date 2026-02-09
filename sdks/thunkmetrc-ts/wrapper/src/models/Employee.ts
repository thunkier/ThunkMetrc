
import type { EmployeeLicense } from './EmployeeLicense';
export interface Employee {
    FullName?: string;
    IsIndustryAdmin?: boolean;
    IsManager?: boolean;
    IsOwner?: boolean;
    License?: EmployeeLicense;
}
