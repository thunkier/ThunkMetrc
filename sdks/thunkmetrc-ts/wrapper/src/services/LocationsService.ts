
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    CreateLocationsRequestItem,
    LocationsLocation,
    LocationsType,
    UpdateLocationsRequestItem,
    WriteResponse,
} from '../models';

export class LocationsService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Creates new locations for a specified Facility.
   * Permissions Required:
   * - Manage Locations
   * POST CreateLocations
   */
  public async createLocations(body: CreateLocationsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createLocations(body,licenseNumber,
        )
    );
  }

  /**
   * Archives a specified Location, identified by its Id, for a Facility.
   * Permissions Required:
   * - Manage Locations
   * DELETE DeleteLocationsById
   */
  public async deleteLocationsById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteLocationsById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of active locations for a specified Facility.
   * Permissions Required:
   * - Manage Locations
   * GET GetActiveLocations
   */
  public async getActiveLocations(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<LocationsLocation>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getActiveLocations(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of inactive locations for a specified Facility.
   * Permissions Required:
   * - Manage Locations
   * GET GetInactiveLocations
   */
  public async getInactiveLocations(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<LocationsLocation>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getInactiveLocations(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a Location by its Id.
   * Permissions Required:
   * - Manage Locations
   * GET GetLocationsById
   */
  public async getLocationsById(id: string, body?: any, licenseNumber?: string): Promise<LocationsLocation> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getLocationsById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of active location types for a specified Facility.
   * Permissions Required:
   * - Manage Locations
   * GET GetLocationsTypes
   */
  public async getLocationsTypes(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<LocationsType>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getLocationsTypes(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Updates existing locations for a specified Facility.
   * Permissions Required:
   * - Manage Locations
   * PUT UpdateLocations
   */
  public async updateLocations(body: UpdateLocationsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateLocations(body,licenseNumber,
        )
    );
  }
}

