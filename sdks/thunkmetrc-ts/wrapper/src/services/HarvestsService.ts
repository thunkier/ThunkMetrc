
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    CreateHarvestsPackagesRequestItem,
    CreateHarvestsWasteRequestItem,
    CreatePackagesTestingRequestItem,
    FinishHarvestsRequestItem,
    Harvest,
    HarvestsWaste,
    UnfinishHarvestsRequestItem,
    UpdateHarvestsLocationRequestItem,
    UpdateRenameRequestItem,
    UpdateRestoreHarvestedPlantsRequestItem,
    WasteType,
    WriteResponse,
} from '../models';

export class HarvestsService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Creates packages from harvested products for a specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Manage Harvests
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * POST CreateHarvestsPackages
   */
  public async createHarvestsPackages(body: CreateHarvestsPackagesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createHarvestsPackages(body,licenseNumber,
        )
    );
  }

  /**
   * Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
   * Permissions Required:
   * - View Harvests
   * - Manage Harvests
   * POST CreateHarvestsWaste
   */
  public async createHarvestsWaste(body: CreateHarvestsWasteRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createHarvestsWaste(body,licenseNumber,
        )
    );
  }

  /**
   * Creates packages for testing from harvested products for a specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Manage Harvests
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * POST CreatePackagesTesting
   */
  public async createHarvestsPackagesTesting(body: CreatePackagesTestingRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createHarvestsPackagesTesting(body,licenseNumber,
        )
    );
  }

  /**
   * Discontinues a specific harvest waste record by Id for the specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Discontinue Harvest Waste
   * DELETE DeleteWasteById
   */
  public async deleteHarvestsWasteById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteHarvestsWasteById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Marks one or more harvests as finished for the specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Finish/Discontinue Harvests
   * PUT FinishHarvests
   */
  public async finishHarvestsHarvests(body: FinishHarvestsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.finishHarvestsHarvests(body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of active harvests for a specified Facility.
   * Permissions Required:
   * - View Harvests
   * GET GetActiveHarvests
   */
  public async getActiveHarvests(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Harvest>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getActiveHarvests(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
   * Permissions Required:
   * - View Harvests
   * GET GetHarvestsById
   */
  public async getHarvestsById(id: string, body?: any, licenseNumber?: string): Promise<Harvest> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getHarvestsById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
   * Permissions Required:
   * - View Harvests
   * GET GetHarvestsWaste
   */
  public async getHarvestsWaste(body?: any, harvestId?: number, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<HarvestsWaste>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getHarvestsWaste(body,harvestId,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of inactive harvests for a specified Facility.
   * Permissions Required:
   * - View Harvests
   * GET GetInactiveHarvests
   */
  public async getInactiveHarvests(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Harvest>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getInactiveHarvests(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of harvests on hold for a specified Facility.
   * Permissions Required:
   * - View Harvests
   * GET GetOnHoldHarvests
   */
  public async getOnHoldHarvests(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Harvest>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getOnHoldHarvests(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of Waste types for harvests.
   * GET GetWasteTypes
   */
  public async getHarvestsWasteTypes(body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<WasteType>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getHarvestsWasteTypes(body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Reopens one or more previously finished harvests for the specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Finish/Discontinue Harvests
   * PUT UnfinishHarvests
   */
  public async unfinishHarvestsHarvests(body: UnfinishHarvestsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.unfinishHarvestsHarvests(body,licenseNumber,
        )
    );
  }

  /**
   * Updates the Location of Harvest for a specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Manage Harvests
   * PUT UpdateHarvestsLocation
   */
  public async updateHarvestsLocation(body: UpdateHarvestsLocationRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateHarvestsLocation(body,licenseNumber,
        )
    );
  }

  /**
   * Renames one or more harvests for the specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Manage Harvests
   * PUT UpdateRename
   */
  public async updateHarvestsRename(body: UpdateRenameRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateHarvestsRename(body,licenseNumber,
        )
    );
  }

  /**
   * Restores previously harvested plants to their original state for the specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Finish/Discontinue Harvests
   * PUT UpdateRestoreHarvestedPlants
   */
  public async updateHarvestsRestoreHarvestedPlants(body: UpdateRestoreHarvestedPlantsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateHarvestsRestoreHarvestedPlants(body,licenseNumber,
        )
    );
  }
}

