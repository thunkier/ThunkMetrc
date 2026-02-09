
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    Employee,
} from '../models';

export class EmployeesService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Retrieves a list of employees for a specified Facility.
   * Permissions Required:
   * - Manage Employees
   * - View Employees
   * GET GetEmployees
   */
  public async getEmployees(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Employee>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getEmployees(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
   * Permissions Required:
   * - Manage Employees
   * GET GetPermissions
   */
  public async getEmployeesPermissions(body?: any, employeeLicenseNumber?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getEmployeesPermissions(body,employeeLicenseNumber,licenseNumber,
        )
    );
  }
}

