
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    CreateStrainsRequestItem,
    Strain,
    UpdateStrainsRequestItem,
    WriteResponse,
} from '../models';

export class StrainsService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Creates new strain records for a specified Facility.
   * Permissions Required:
   * - Manage Strains
   * POST CreateStrains
   */
  public async createStrains(body: CreateStrainsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createStrains(body,licenseNumber,
        )
    );
  }

  /**
   * Archives an existing strain record for a Facility
   * Permissions Required:
   * - Manage Strains
   * DELETE DeleteStrainsById
   */
  public async deleteStrainsById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteStrainsById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
   * Permissions Required:
   * - Manage Strains
   * GET GetActiveStrains
   */
  public async getActiveStrains(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Strain>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getActiveStrains(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
   * Permissions Required:
   * - Manage Strains
   * GET GetInactiveStrains
   */
  public async getInactiveStrains(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Strain>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getInactiveStrains(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a Strain record by its Id, with an optional license number.
   * Permissions Required:
   * - Manage Strains
   * GET GetStrainsById
   */
  public async getStrainsById(id: string, body?: any, licenseNumber?: string): Promise<Strain> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getStrainsById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Updates existing strain records for a specified Facility.
   * Permissions Required:
   * - Manage Strains
   * PUT UpdateStrains
   */
  public async updateStrains(body: UpdateStrainsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateStrains(body,licenseNumber,
        )
    );
  }
}

