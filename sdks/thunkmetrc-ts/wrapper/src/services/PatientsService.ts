
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    CreatePatientsRequestItem,
    Patient,
    UpdatePatientsRequestItem,
    WriteResponse,
} from '../models';

export class PatientsService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Adds new patients to a specified Facility.
   * Permissions Required:
   * - Manage Patients
   * POST CreatePatients
   */
  public async createPatients(body: CreatePatientsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPatients(body,licenseNumber,
        )
    );
  }

  /**
   * Removes a Patient, identified by an Id, from a specified Facility.
   * Permissions Required:
   * - Manage Patients
   * DELETE DeletePatientsById
   */
  public async deletePatientsById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deletePatientsById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of active patients for a specified Facility.
   * Permissions Required:
   * - Manage Patients
   * GET GetActivePatients
   */
  public async getActivePatients(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Patient>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getActivePatients(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a Patient by Id.
   * Permissions Required:
   * - Manage Patients
   * GET GetPatientsById
   */
  public async getPatientsById(id: string, body?: any, licenseNumber?: string): Promise<Patient> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPatientsById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Updates Patient information for a specified Facility.
   * Permissions Required:
   * - Manage Patients
   * PUT UpdatePatients
   */
  public async updatePatients(body: UpdatePatientsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePatients(body,licenseNumber,
        )
    );
  }
}

