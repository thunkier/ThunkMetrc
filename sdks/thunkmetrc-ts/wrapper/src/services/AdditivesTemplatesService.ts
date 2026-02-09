
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    AdditivesTemplate,
    CreateAdditivesTemplatesRequestItem,
    UpdateAdditivesTemplatesRequestItem,
    WriteResponse,
} from '../models';

export class AdditivesTemplatesService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Creates new additive templates for a specified Facility.
   * Permissions Required:
   * - Manage Additives
   * POST CreateAdditivesTemplates
   */
  public async createAdditivesTemplates(body: CreateAdditivesTemplatesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createAdditivesTemplates(body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of active additive templates for a specified Facility.
   * Permissions Required:
   * - Manage Additives
   * GET GetActiveAdditivesTemplates
   */
  public async getActiveAdditivesTemplates(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<AdditivesTemplate>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getActiveAdditivesTemplates(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves an Additive Template by its Id.
   * Permissions Required:
   * - Manage Additives
   * GET GetAdditivesTemplatesById
   */
  public async getAdditivesTemplatesById(id: string, body?: any, licenseNumber?: string): Promise<AdditivesTemplate> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getAdditivesTemplatesById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of inactive additive templates for a specified Facility.
   * Permissions Required:
   * - Manage Additives
   * GET GetInactiveAdditivesTemplates
   */
  public async getInactiveAdditivesTemplates(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<AdditivesTemplate>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getInactiveAdditivesTemplates(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Updates existing additive templates for a specified Facility.
   * Permissions Required:
   * - Manage Additives
   * PUT UpdateAdditivesTemplates
   */
  public async updateAdditivesTemplates(body: UpdateAdditivesTemplatesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateAdditivesTemplates(body,licenseNumber,
        )
    );
  }
}

