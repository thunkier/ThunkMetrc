
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    Additive,
    CreateAdditivesByLocationRequestItem,
    CreateAdditivesByLocationUsingTemplateRequestItem,
    CreateManicureRequestItem,
    CreatePlantBatchPackagesRequestItem,
    CreatePlantsAdditivesRequestItem,
    CreatePlantsAdditivesUsingTemplateRequestItem,
    CreatePlantsPlantingsRequestItem,
    CreatePlantsWasteRequestItem,
    DeletePlantsRequestItem,
    Mother,
    Plant,
    PlantsWaste,
    UpdateAdjustPlantsRequestItem,
    UpdateGrowthPhaseRequestItem,
    UpdateHarvestRequestItem,
    UpdateMergeRequestItem,
    UpdatePlantsLocationRequestItem,
    UpdatePlantsStrainRequestItem,
    UpdatePlantsTagRequestItem,
    UpdateSplitRequestItem,
    WasteMethod,
    WastePackage,
    WasteReason,
    WriteResponse,
} from '../models';

export class PlantsService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Records additive usage for plants based on their location within a specified Facility.
   * Permissions Required:
   * - Manage Plants
   * - Manage Plants Additives
   * POST CreateAdditivesByLocation
   */
  public async createPlantsAdditivesByLocation(body: CreateAdditivesByLocationRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantsAdditivesByLocation(body,licenseNumber,
        )
    );
  }

  /**
   * Records additive usage for plants by location using a predefined additive template at a specified Facility.
   * Permissions Required:
   * - Manage Plants Additives
   * POST CreateAdditivesByLocationUsingTemplate
   */
  public async createPlantsAdditivesByLocationUsingTemplate(body: CreateAdditivesByLocationUsingTemplateRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantsAdditivesByLocationUsingTemplate(body,licenseNumber,
        )
    );
  }

  /**
   * Creates harvest product records from plant batches at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manicure/Harvest Veg/Flower Plants
   * POST CreateManicure
   */
  public async createPlantsManicure(body: CreateManicureRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantsManicure(body,licenseNumber,
        )
    );
  }

  /**
   * Creates packages from plant batches at a specified Facility.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants Inventory
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * POST CreatePlantBatchPackages
   */
  public async createPlantsPlantBatchPackages(body: CreatePlantBatchPackagesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantsPlantBatchPackages(body,licenseNumber,
        )
    );
  }

  /**
   * Records additive usage details applied to specific plants at a Facility.
   * Permissions Required:
   * - Manage Plants Additives
   * POST CreatePlantsAdditives
   */
  public async createPlantsAdditives(body: CreatePlantsAdditivesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantsAdditives(body,licenseNumber,
        )
    );
  }

  /**
   * Records additive usage for plants using predefined additive templates at a specified Facility.
   * Permissions Required:
   * - Manage Plants Additives
   * POST CreatePlantsAdditivesUsingTemplate
   */
  public async createPlantsAdditivesUsingTemplate(body: CreatePlantsAdditivesUsingTemplateRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantsAdditivesUsingTemplate(body,licenseNumber,
        )
    );
  }

  /**
   * Creates new plant batches at a specified Facility from existing plant data.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants Inventory
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * POST CreatePlantsPlantings
   */
  public async createPlantsPlantings(body: CreatePlantsPlantingsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantsPlantings(body,licenseNumber,
        )
    );
  }

  /**
   * Records waste events for plants at a Facility, including method, reason, and location details.
   * Permissions Required:
   * - Manage Plants Waste
   * POST CreatePlantsWaste
   */
  public async createPlantsWaste(body: CreatePlantsWasteRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPlantsWaste(body,licenseNumber,
        )
    );
  }

  /**
   * Removes plants from a Facility’s inventory while recording the reason for their disposal.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Destroy Veg/Flower Plants
   * DELETE DeletePlants
   */
  public async deletePlants(body?: any, licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deletePlants(body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves additive records applied to plants at a specified Facility.
   * Permissions Required:
   * - View/Manage Plants Additives
   * GET GetAdditives
   */
  public async getPlantsAdditives(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Additive>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantsAdditives(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of all plant additive types defined within a Facility.
   * GET GetAdditivesTypes
   */
  public async getPlantsAdditivesTypes(body?: any): Promise<any> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantsAdditivesTypes(body,
        )
    );
  }

  /**
   * Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
   * Permissions Required:
   * - View Veg/Flower Plants
   * GET GetFloweringPlants
   */
  public async getFloweringPlants(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Plant>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getFloweringPlants(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves the list of growth phases supported by a specified Facility.
   * GET GetGrowthPhases
   */
  public async getPlantsGrowthPhases(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantsGrowthPhases(body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves inactive plants at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * GET GetInactivePlants
   */
  public async getInactivePlants(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Plant>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getInactivePlants(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves inactive mother-phase plants at a specified Facility.
   * Permissions Required:
   * - View Mother Plants
   * GET GetMotherInactivePlants
   */
  public async getMotherInactivePlants(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Mother>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getMotherInactivePlants(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves mother-phase plants currently marked as on hold at a specified Facility.
   * Permissions Required:
   * - View Mother Plants
   * GET GetMotherOnHoldPlants
   */
  public async getMotherOnHoldPlants(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Mother>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getMotherOnHoldPlants(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves mother-phase plants at a specified Facility.
   * Permissions Required:
   * - View Mother Plants
   * GET GetMotherPlants
   */
  public async getMotherPlants(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Mother>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getMotherPlants(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves plants that are currently on hold at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * GET GetOnHoldPlants
   */
  public async getOnHoldPlants(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Plant>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getOnHoldPlants(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a Plant by Id.
   * Permissions Required:
   * - View Veg/Flower Plants
   * GET GetPlantsById
   */
  public async getPlantsById(id: string, body?: any, licenseNumber?: string): Promise<Plant> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantsById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a Plant by label.
   * Permissions Required:
   * - View Veg/Flower Plants
   * GET GetPlantsByLabel
   */
  public async getPlantsByLabel(label: string, body?: any, licenseNumber?: string): Promise<Plant[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantsByLabel(label, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of recorded plant waste events for a specific Facility.
   * Permissions Required:
   * - View Plants Waste
   * GET GetPlantsWaste
   */
  public async getPlantsWaste(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<PlantsWaste>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantsWaste(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of all available plant waste methods for use within a Facility.
   * GET GetPlantsWasteMethods
   */
  public async getPlantsWasteMethods(body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<WasteMethod>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantsWasteMethods(body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retriveves available reasons for recording mature plant waste at a specified Facility.
   * GET GetPlantsWasteReasons
   */
  public async getPlantsWasteReasons(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<WasteReason>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantsWasteReasons(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
   * Permissions Required:
   * - View Veg/Flower Plants
   * GET GetVegetativePlants
   */
  public async getVegetativePlants(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Plant>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getVegetativePlants(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
   * Permissions Required:
   * - View Plants Waste
   * GET GetWasteById
   */
  public async getPlantsWasteById(id: string, body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<PlantsWaste>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantsWasteById(id, body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of package records linked to the specified plantWasteId for a given facility.
   * Permissions Required:
   * - View Plants Waste
   * GET GetWastePackageById
   */
  public async getPlantsWastePackageById(id: string, body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<WastePackage>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPlantsWastePackageById(id, body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Adjusts the recorded count of plants at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdateAdjustPlants
   */
  public async updateAdjustPlants(body: UpdateAdjustPlantsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateAdjustPlants(body,licenseNumber,
        )
    );
  }

  /**
   * Changes the growth phases of plants within a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdateGrowthPhase
   */
  public async updatePlantsGrowthPhase(body: UpdateGrowthPhaseRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePlantsGrowthPhase(body,licenseNumber,
        )
    );
  }

  /**
   * Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manicure/Harvest Veg/Flower Plants
   * PUT UpdateHarvest
   */
  public async updatePlantsHarvest(body: UpdateHarvestRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePlantsHarvest(body,licenseNumber,
        )
    );
  }

  /**
   * Merges multiple plant groups into a single group within a Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manicure/Harvest Veg/Flower Plants
   * PUT UpdateMerge
   */
  public async updatePlantsMerge(body: UpdateMergeRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePlantsMerge(body,licenseNumber,
        )
    );
  }

  /**
   * Moves plant batches to new locations within a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdatePlantsLocation
   */
  public async updatePlantsLocation(body: UpdatePlantsLocationRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePlantsLocation(body,licenseNumber,
        )
    );
  }

  /**
   * Updates the strain information for plants within a Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdatePlantsStrain
   */
  public async updatePlantsStrain(body: UpdatePlantsStrainRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePlantsStrain(body,licenseNumber,
        )
    );
  }

  /**
   * Replaces existing plant tags with new tags for plants within a Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdatePlantsTag
   */
  public async updatePlantsTag(body: UpdatePlantsTagRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePlantsTag(body,licenseNumber,
        )
    );
  }

  /**
   * Splits an existing plant group into multiple groups within a Facility.
   * Permissions Required:
   * - View Plant
   * PUT UpdateSplit
   */
  public async updatePlantsSplit(body: UpdateSplitRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePlantsSplit(body,licenseNumber,
        )
    );
  }
}

