
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    CreateAdjustPlantBatchesRequestItem,
    CreateGrowthPhaseRequestItem,
    CreatePackagesFromMotherPlantRequestItem,
    CreatePlantBatchesAdditivesRequestItem,
    CreatePlantBatchesAdditivesUsingTemplateRequestItem,
    CreatePlantBatchesPackagesRequestItem,
    CreatePlantBatchesPlantingsRequestItem,
    CreatePlantBatchesWasteRequestItem,
    CreateSplitRequestItem,
    DeletePlantBatchesRequestItem,
    PlantBatch,
    PlantBatchesType,
    PlantBatchesWaste,
    PlantBatchesWasteReason,
    UpdateNameRequestItem,
    UpdatePlantBatchesLocationRequestItem,
    UpdatePlantBatchesStrainRequestItem,
    UpdatePlantBatchesTagRequestItem,
    WriteResponse,
} from '../models';

export class PlantBatchesService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants Inventory
   * POST CreateAdjustPlantBatches
   */
  public async createAdjustPlantBatches(body: CreateAdjustPlantBatchesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createAdjustPlantBatches(body,licenseNumber,
        )
    );
  }

  /**
   * Updates the growth phase of plants at a specified Facility based on tracking information.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants Inventory
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * POST CreateGrowthPhase
   */
  public async createPlantBatchesGrowthPhase(body: CreateGrowthPhaseRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantBatchesGrowthPhase(body,licenseNumber,
        )
    );
  }

  /**
   * Creates packages from mother plants at the specified Facility.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants Inventory
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * POST CreatePackagesFromMotherPlant
   */
  public async createPlantBatchesPackagesFromMotherPlant(body: CreatePackagesFromMotherPlantRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantBatchesPackagesFromMotherPlant(body,licenseNumber,
        )
    );
  }

  /**
   * Records Additive usage details for plant batches at a specific Facility.
   * Permissions Required:
   * - Manage Plants Additives
   * POST CreatePlantBatchesAdditives
   */
  public async createPlantBatchesAdditives(body: CreatePlantBatchesAdditivesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantBatchesAdditives(body,licenseNumber,
        )
    );
  }

  /**
   * Records Additive usage for plant batches at a Facility using predefined additive templates.
   * Permissions Required:
   * - Manage Plants Additives
   * POST CreatePlantBatchesAdditivesUsingTemplate
   */
  public async createPlantBatchesAdditivesUsingTemplate(body: CreatePlantBatchesAdditivesUsingTemplateRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantBatchesAdditivesUsingTemplate(body,licenseNumber,
        )
    );
  }

  /**
   * Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants Inventory
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * POST CreatePlantBatchesPackages
   */
  public async createPlantBatchesPackages(body: CreatePlantBatchesPackagesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantBatchesPackages(body,licenseNumber,
        )
    );
  }

  /**
   * Creates new plantings for a Facility by generating plant batches based on provided planting details.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants Inventory
   * POST CreatePlantBatchesPlantings
   */
  public async createPlantBatchesPlantings(body: CreatePlantBatchesPlantingsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantBatchesPlantings(body,licenseNumber,
        )
    );
  }

  /**
   * Records waste information for plant batches based on the submitted data for the specified Facility.
   * Permissions Required:
   * - Manage Plants Waste
   * POST CreatePlantBatchesWaste
   */
  public async createPlantBatchesWaste(body: CreatePlantBatchesWasteRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantBatchesWaste(body,licenseNumber,
        )
    );
  }

  /**
   * Splits an existing Plant Batch into multiple groups at the specified Facility.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants Inventory
   * POST CreateSplit
   */
  public async createPlantBatchesSplit(body: CreateSplitRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantBatchesSplit(body,licenseNumber,
        )
    );
  }

  /**
   * Completes the destruction of plant batches based on the provided input data.
   * Permissions Required:
   * - View Immature Plants
   * - Destroy Immature Plants
   * DELETE DeletePlantBatches
   */
  public async deletePlantBatches(body?: any, licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deletePlantBatches(body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
   * Permissions Required:
   * - View Immature Plants
   * GET GetActivePlantBatches
   */
  public async getActivePlantBatches(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<PlantBatch>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getActivePlantBatches(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
   * Permissions Required:
   * - View Immature Plants
   * GET GetInactivePlantBatches
   */
  public async getInactivePlantBatches(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<PlantBatch>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getInactivePlantBatches(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a Plant Batch by Id.
   * Permissions Required:
   * - View Immature Plants
   * GET GetPlantBatchesById
   */
  public async getPlantBatchesById(id: string, body?: any, licenseNumber?: string): Promise<PlantBatch> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantBatchesById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of plant batch types.
   * GET GetPlantBatchesTypes
   */
  public async getPlantBatchesTypes(body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<PlantBatchesType>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantBatchesTypes(body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves waste details associated with plant batches at a specified Facility.
   * Permissions Required:
   * - View Plants Waste
   * GET GetPlantBatchesWaste
   */
  public async getPlantBatchesWaste(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<PlantBatchesWaste>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantBatchesWaste(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
   * GET GetPlantBatchesWasteReasons
   */
  public async getPlantBatchesWasteReasons(body?: any, licenseNumber?: string): Promise<PlantBatchesWasteReason[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantBatchesWasteReasons(body,licenseNumber,
        )
    );
  }

  /**
   * Renames plant batches at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdateName
   */
  public async updatePlantBatchesName(body: UpdateNameRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePlantBatchesName(body,licenseNumber,
        )
    );
  }

  /**
   * Moves one or more plant batches to new locations with in a specified Facility.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants
   * PUT UpdatePlantBatchesLocation
   */
  public async updatePlantBatchesLocation(body: UpdatePlantBatchesLocationRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePlantBatchesLocation(body,licenseNumber,
        )
    );
  }

  /**
   * Changes the strain of plant batches at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdatePlantBatchesStrain
   */
  public async updatePlantBatchesStrain(body: UpdatePlantBatchesStrainRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePlantBatchesStrain(body,licenseNumber,
        )
    );
  }

  /**
   * Replaces tags for plant batches at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdatePlantBatchesTag
   */
  public async updatePlantBatchesTag(body: UpdatePlantBatchesTagRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePlantBatchesTag(body,licenseNumber,
        )
    );
  }
}

