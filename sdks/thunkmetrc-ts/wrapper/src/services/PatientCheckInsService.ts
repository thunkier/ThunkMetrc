
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    CreatePatientCheckInsRequestItem,
    PatientCheckIn,
    PatientCheckInsLocation,
    UpdatePatientCheckInsRequestItem,
    WriteResponse,
} from '../models';

export class PatientCheckInsService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Records patient check-ins for a specified Facility.
   * Permissions Required:
   * - ManagePatientsCheckIns
   * POST CreatePatientCheckIns
   */
  public async createPatientCheckIns(body: CreatePatientCheckInsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPatientCheckIns(body,licenseNumber,
        )
    );
  }

  /**
   * Archives a Patient Check-In, identified by its Id, for a specified Facility.
   * Permissions Required:
   * - ManagePatientsCheckIns
   * DELETE DeletePatientCheckInsById
   */
  public async deletePatientCheckInsById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deletePatientCheckInsById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of Patient Check-In locations.
   * GET GetLocations
   */
  public async getPatientCheckInsLocations(body?: any): Promise<PatientCheckInsLocation[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPatientCheckInsLocations(body,
        )
    );
  }

  /**
   * Retrieves a list of patient check-ins for a specified Facility.
   * Permissions Required:
   * - ManagePatientsCheckIns
   * GET GetPatientCheckIns
   */
  public async getPatientCheckIns(body?: any, checkinDateEnd?: string, checkinDateStart?: string, licenseNumber?: string): Promise<PaginatedResponse<PatientCheckIn>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPatientCheckIns(body,checkinDateEnd,checkinDateStart,licenseNumber,
        )
    );
  }

  /**
   * Updates patient check-ins for a specified Facility.
   * Permissions Required:
   * - ManagePatientsCheckIns
   * PUT UpdatePatientCheckIns
   */
  public async updatePatientCheckIns(body: UpdatePatientCheckInsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePatientCheckIns(body,licenseNumber,
        )
    );
  }
}

