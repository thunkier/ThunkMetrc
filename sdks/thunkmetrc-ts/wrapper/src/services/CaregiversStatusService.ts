
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    CaregiversStatus,
} from '../models';

export class CaregiversStatusService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
   * Permissions Required:
   * - Lookup Caregivers
   * GET GetCaregiversStatusByCaregiverLicenseNumber
   */
  public async getCaregiversStatusByCaregiverLicenseNumber(caregiverLicenseNumber: string, body?: any, licenseNumber?: string): Promise<CaregiversStatus[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getCaregiversStatusByCaregiverLicenseNumber(caregiverLicenseNumber, body,licenseNumber,
        )
    );
  }
}

