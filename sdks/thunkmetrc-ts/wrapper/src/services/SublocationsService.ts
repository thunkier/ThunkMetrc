
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    CreateSublocationsRequestItem,
    Sublocation,
    UpdateSublocationsRequestItem,
    WriteResponse,
} from '../models';

export class SublocationsService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Creates new sublocation records for a Facility.
   * Permissions Required:
   * - Manage Locations
   * POST CreateSublocations
   */
  public async createSublocations(body: CreateSublocationsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createSublocations(body,licenseNumber,
        )
    );
  }

  /**
   * Archives an existing Sublocation record for a Facility.
   * Permissions Required:
   * - Manage Locations
   * DELETE DeleteSublocationsById
   */
  public async deleteSublocationsById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteSublocationsById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
   * Permissions Required:
   * - Manage Locations
   * GET GetActiveSublocations
   */
  public async getActiveSublocations(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Sublocation>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getActiveSublocations(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of inactive sublocations for the specified Facility.
   * Permissions Required:
   * - Manage Locations
   * GET GetInactiveSublocations
   */
  public async getInactiveSublocations(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Sublocation>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getInactiveSublocations(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a Sublocation by its Id, with an optional license number.
   * Permissions Required:
   * - Manage Locations
   * GET GetSublocationsById
   */
  public async getSublocationsById(id: string, body?: any, licenseNumber?: string): Promise<Sublocation> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSublocationsById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Updates existing sublocation records for a specified Facility.
   * Permissions Required:
   * - Manage Locations
   * PUT UpdateSublocations
   */
  public async updateSublocations(body: UpdateSublocationsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateSublocations(body,licenseNumber,
        )
    );
  }
}

