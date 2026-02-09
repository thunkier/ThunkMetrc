
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    PatientsStatus,
} from '../models';

export class PatientsStatusService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
   * Permissions Required:
   * - Lookup Patients
   * GET GetPatientsStatusesByPatientLicenseNumber
   */
  public async getPatientsStatusesByPatientLicenseNumber(patientLicenseNumber: string, body?: any, licenseNumber?: string): Promise<PatientsStatus[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPatientsStatusesByPatientLicenseNumber(patientLicenseNumber, body,licenseNumber,
        )
    );
  }
}

