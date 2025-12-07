export * from '@thunkmetrc/client';
import { MetrcClient as InternalMetrcClient } from '@thunkmetrc/client';
import { MetrcRateLimiter, RateLimiterConfig } from './RateLimiter';


export class MetrcWrapper {
  private rateLimiter: MetrcRateLimiter;

  constructor(public client: InternalMetrcClient, config?: RateLimiterConfig) {
    this.rateLimiter = new MetrcRateLimiter(config);
  }

  /**
   * Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditives V1
   */
  public async plantsCreateAdditivesV1(body: PlantsCreateAdditivesV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreateAdditivesV1(body, licenseNumber));
  }

  /**
   * Records additive usage details applied to specific plants at a Facility.
   * 
   *   Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditives V2
   */
  public async plantsCreateAdditivesV2(body: PlantsCreateAdditivesV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreateAdditivesV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Plants
   *   - Manage Plants Additives
   *
   * POST CreateAdditivesBylocation V1
   */
  public async plantsCreateAdditivesBylocationV1(body: PlantsCreateAdditivesBylocationV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreateAdditivesBylocationV1(body, licenseNumber));
  }

  /**
   * Records additive usage for plants based on their location within a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Plants
   *   - Manage Plants Additives
   *
   * POST CreateAdditivesBylocation V2
   */
  public async plantsCreateAdditivesBylocationV2(body: PlantsCreateAdditivesBylocationV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreateAdditivesBylocationV2(body, licenseNumber));
  }

  /**
   * Records additive usage for plants by location using a predefined additive template at a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditivesBylocationUsingtemplate V2
   */
  public async plantsCreateAdditivesBylocationUsingtemplateV2(body: PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreateAdditivesBylocationUsingtemplateV2(body, licenseNumber));
  }

  /**
   * Records additive usage for plants using predefined additive templates at a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditivesUsingtemplate V2
   */
  public async plantsCreateAdditivesUsingtemplateV2(body: PlantsCreateAdditivesUsingtemplateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreateAdditivesUsingtemplateV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * POST CreateChangegrowthphases V1
   */
  public async plantsCreateChangegrowthphasesV1(body: PlantsCreateChangegrowthphasesV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreateChangegrowthphasesV1(body, licenseNumber));
  }

  /**
   * NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manicure/Harvest Veg/Flower Plants
   *
   * POST CreateHarvestplants V1
   */
  public async plantsCreateHarvestplantsV1(body: PlantsCreateHarvestplantsV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreateHarvestplantsV1(body, licenseNumber));
  }

  /**
   * Creates harvest product records from plant batches at a specified Facility.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manicure/Harvest Veg/Flower Plants
   *
   * POST CreateManicure V2
   */
  public async plantsCreateManicureV2(body: PlantsCreateManicureV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreateManicureV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manicure/Harvest Veg/Flower Plants
   *
   * POST CreateManicureplants V1
   */
  public async plantsCreateManicureplantsV1(body: PlantsCreateManicureplantsV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreateManicureplantsV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * POST CreateMoveplants V1
   */
  public async plantsCreateMoveplantsV1(body: PlantsCreateMoveplantsV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreateMoveplantsV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreatePlantbatchPackage V1
   */
  public async plantsCreatePlantbatchPackageV1(body: PlantsCreatePlantbatchPackageV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreatePlantbatchPackageV1(body, licenseNumber));
  }

  /**
   * Creates packages from plant batches at a specified Facility.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreatePlantbatchPackage V2
   */
  public async plantsCreatePlantbatchPackageV2(body: PlantsCreatePlantbatchPackageV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreatePlantbatchPackageV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * POST CreatePlantings V1
   */
  public async plantsCreatePlantingsV1(body: PlantsCreatePlantingsV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreatePlantingsV1(body, licenseNumber));
  }

  /**
   * Creates new plant batches at a specified Facility from existing plant data.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * POST CreatePlantings V2
   */
  public async plantsCreatePlantingsV2(body: PlantsCreatePlantingsV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreatePlantingsV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Plants Waste
   *
   * POST CreateWaste V1
   */
  public async plantsCreateWasteV1(body: PlantsCreateWasteV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreateWasteV1(body, licenseNumber));
  }

  /**
   * Records waste events for plants at a Facility, including method, reason, and location details.
   * 
   *   Permissions Required:
   *   - Manage Plants Waste
   *
   * POST CreateWaste V2
   */
  public async plantsCreateWasteV2(body: PlantsCreateWasteV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsCreateWasteV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *   - Destroy Veg/Flower Plants
   *
   * DELETE Delete V1
   */
  public async plantsDeleteV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsDeleteV1(body, licenseNumber));
  }

  /**
   * Removes plants from a Facility’s inventory while recording the reason for their disposal.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *   - Destroy Veg/Flower Plants
   *
   * DELETE Delete V2
   */
  public async plantsDeleteV2(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsDeleteV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET Get V1
   */
  public async plantsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetV1(id, body, licenseNumber));
  }

  /**
   * Retrieves a Plant by Id.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET Get V2
   */
  public async plantsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View/Manage Plants Additives
   *
   * GET GetAdditives V1
   */
  public async plantsGetAdditivesV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetAdditivesV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves additive records applied to plants at a specified Facility.
   * 
   *   Permissions Required:
   *   - View/Manage Plants Additives
   *
   * GET GetAdditives V2
   */
  public async plantsGetAdditivesV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetAdditivesV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   -
   *
   * GET GetAdditivesTypes V1
   */
  public async plantsGetAdditivesTypesV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetAdditivesTypesV1(body, No));
  }

  /**
   * Retrieves a list of all plant additive types defined within a Facility.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetAdditivesTypes V2
   */
  public async plantsGetAdditivesTypesV2(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetAdditivesTypesV2(body, No));
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetByLabel V1
   */
  public async plantsGetByLabelV1(label: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetByLabelV1(label, body, licenseNumber));
  }

  /**
   * Retrieves a Plant by label.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetByLabel V2
   */
  public async plantsGetByLabelV2(label: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetByLabelV2(label, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetFlowering V1
   */
  public async plantsGetFloweringV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetFloweringV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetFlowering V2
   */
  public async plantsGetFloweringV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetFloweringV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetGrowthPhases V1
   */
  public async plantsGetGrowthPhasesV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetGrowthPhasesV1(body, licenseNumber));
  }

  /**
   * Retrieves the list of growth phases supported by a specified Facility.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetGrowthPhases V2
   */
  public async plantsGetGrowthPhasesV2(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetGrowthPhasesV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetInactive V1
   */
  public async plantsGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves inactive plants at a specified Facility.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetInactive V2
   */
  public async plantsGetInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves mother-phase plants at a specified Facility.
   * 
   *   Permissions Required:
   *   - View Mother Plants
   *
   * GET GetMother V2
   */
  public async plantsGetMotherV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetMotherV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves inactive mother-phase plants at a specified Facility.
   * 
   *   Permissions Required:
   *   - View Mother Plants
   *
   * GET GetMotherInactive V2
   */
  public async plantsGetMotherInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetMotherInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves mother-phase plants currently marked as on hold at a specified Facility.
   * 
   *   Permissions Required:
   *   - View Mother Plants
   *
   * GET GetMotherOnhold V2
   */
  public async plantsGetMotherOnholdV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetMotherOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetOnhold V1
   */
  public async plantsGetOnholdV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetOnholdV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves plants that are currently on hold at a specified Facility.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetOnhold V2
   */
  public async plantsGetOnholdV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetVegetative V1
   */
  public async plantsGetVegetativeV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetVegetativeV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetVegetative V2
   */
  public async plantsGetVegetativeV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetVegetativeV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of recorded plant waste events for a specific Facility.
   * 
   *   Permissions Required:
   *   - View Plants Waste
   *
   * GET GetWaste V2
   */
  public async plantsGetWasteV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetWasteV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetWasteMethodsAll V1
   */
  public async plantsGetWasteMethodsAllV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetWasteMethodsAllV1(body, No));
  }

  /**
   * Retrieves a list of all available plant waste methods for use within a Facility.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetWasteMethodsAll V2
   */
  public async plantsGetWasteMethodsAllV2(body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetWasteMethodsAllV2(body, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of package records linked to the specified plantWasteId for a given facility.
   * 
   *   Permissions Required:
   *   - View Plants Waste
   *
   * GET GetWastePackage V2
   */
  public async plantsGetWastePackageV2(id: string, body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetWastePackageV2(id, body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
   * 
   *   Permissions Required:
   *   - View Plants Waste
   *
   * GET GetWastePlant V2
   */
  public async plantsGetWastePlantV2(id: string, body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetWastePlantV2(id, body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetWasteReasons V1
   */
  public async plantsGetWasteReasonsV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetWasteReasonsV1(body, licenseNumber));
  }

  /**
   * Retriveves available reasons for recording mature plant waste at a specified Facility.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetWasteReasons V2
   */
  public async plantsGetWasteReasonsV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantsGetWasteReasonsV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Adjusts the recorded count of plants at a specified Facility.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * PUT UpdateAdjust V2
   */
  public async plantsUpdateAdjustV2(body: PlantsUpdateAdjustV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateAdjustV2(body, licenseNumber));
  }

  /**
   * Changes the growth phases of plants within a specified Facility.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * PUT UpdateGrowthphase V2
   */
  public async plantsUpdateGrowthphaseV2(body: PlantsUpdateGrowthphaseV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateGrowthphaseV2(body, licenseNumber));
  }

  /**
   * Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manicure/Harvest Veg/Flower Plants
   *
   * PUT UpdateHarvest V2
   */
  public async plantsUpdateHarvestV2(body: PlantsUpdateHarvestV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateHarvestV2(body, licenseNumber));
  }

  /**
   * Moves plant batches to new locations within a specified Facility.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * PUT UpdateLocation V2
   */
  public async plantsUpdateLocationV2(body: PlantsUpdateLocationV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateLocationV2(body, licenseNumber));
  }

  /**
   * Merges multiple plant groups into a single group within a Facility.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manicure/Harvest Veg/Flower Plants
   *
   * PUT UpdateMerge V2
   */
  public async plantsUpdateMergeV2(body: PlantsUpdateMergeV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateMergeV2(body, licenseNumber));
  }

  /**
   * Splits an existing plant group into multiple groups within a Facility.
   * 
   *   Permissions Required:
   *   - View Plant
   *
   * PUT UpdateSplit V2
   */
  public async plantsUpdateSplitV2(body: PlantsUpdateSplitV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateSplitV2(body, licenseNumber));
  }

  /**
   * Updates the strain information for plants within a Facility.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * PUT UpdateStrain V2
   */
  public async plantsUpdateStrainV2(body: PlantsUpdateStrainV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateStrainV2(body, licenseNumber));
  }

  /**
   * Replaces existing plant tags with new tags for plants within a Facility.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * PUT UpdateTag V2
   */
  public async plantsUpdateTagV2(body: PlantsUpdateTagV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantsUpdateTagV2(body, licenseNumber));
  }

  /**
   * Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
   *   - WebApi Retail ID Read Write State (All or WriteOnly)
   *   - Industry/View Packages
   *   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
   *
   * POST CreateAssociate V2
   */
  public async retailIdCreateAssociateV2(body: RetailidCreateAssociateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.retailIdCreateAssociateV2(body, licenseNumber));
  }

  /**
   * Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
   *   - WebApi Retail ID Read Write State (All or WriteOnly)
   *   - Industry/View Packages
   *   - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
   *
   * POST CreateGenerate V2
   */
  public async retailIdCreateGenerateV2(body: RetailidCreateGenerateV2Request, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.retailIdCreateGenerateV2(body, licenseNumber));
  }

  /**
   * Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
   *   - WebApi Retail ID Read Write State (All or WriteOnly)
   *   - Key Value Settings/Retail ID Merge Packages Enabled
   *
   * POST CreateMerge V2
   */
  public async retailIdCreateMergeV2(body: RetailidCreateMergeV2Request, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.retailIdCreateMergeV2(body, licenseNumber));
  }

  /**
   * Retrieves Package information for given list of Package labels.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
   *   - WebApi Retail ID Read Write State (All or WriteOnly)
   *   - Industry/View Packages
   *   - Admin/Employees/Packages Page/Product Labels(Manage)
   *
   * POST CreatePackageInfo V2
   */
  public async retailIdCreatePackageInfoV2(body: RetailidCreatePackageInfoV2Request, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.retailIdCreatePackageInfoV2(body, licenseNumber));
  }

  /**
   * Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Manage RetailId
   *   - WebApi Retail ID Read Write State (All or ReadOnly)
   *   - Industry/View Packages
   *   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
   *
   * GET GetReceiveByLabel V2
   */
  public async retailIdGetReceiveByLabelV2(label: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.retailIdGetReceiveByLabelV2(label, body, licenseNumber));
  }

  /**
   * Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Manage RetailId
   *   - WebApi Retail ID Read Write State (All or ReadOnly)
   *   - Industry/View Packages
   *   - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
   *
   * GET GetReceiveQrByShortCode V2
   */
  public async retailIdGetReceiveQrByShortCodeV2(shortCode: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.retailIdGetReceiveQrByShortCodeV2(shortCode, body, licenseNumber));
  }

  /**
   * This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
   * 
   *   Permissions Required:
   *   - None
   *
   * POST CreateIntegratorSetup V2
   */
  public async sandboxCreateIntegratorSetupV2(body?: any, userKey?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.sandboxCreateIntegratorSetupV2(body, userKey));
  }

  /**
   * This endpoint provides a list of facilities for which the authenticated user has access.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetAll V1
   */
  public async facilitiesGetAllV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.facilitiesGetAllV1(body, No));
  }

  /**
   * This endpoint provides a list of facilities for which the authenticated user has access.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetAll V2
   */
  public async facilitiesGetAllV2(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.facilitiesGetAllV2(body, No));
  }

  /**
   * Adds new patients to a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Patients
   *
   * POST Create V2
   */
  public async patientsCreateV2(body: PatientsCreateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.patientsCreateV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Patients
   *
   * POST CreateAdd V1
   */
  public async patientsCreateAddV1(body: PatientsCreateAddV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.patientsCreateAddV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Patients
   *
   * POST CreateUpdate V1
   */
  public async patientsCreateUpdateV1(body: PatientsCreateUpdateV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.patientsCreateUpdateV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Patients
   *
   * DELETE Delete V1
   */
  public async patientsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.patientsDeleteV1(id, body, licenseNumber));
  }

  /**
   * Removes a Patient, identified by an Id, from a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Patients
   *
   * DELETE Delete V2
   */
  public async patientsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.patientsDeleteV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Patients
   *
   * GET Get V1
   */
  public async patientsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.patientsGetV1(id, body, licenseNumber));
  }

  /**
   * Retrieves a Patient by Id.
   * 
   *   Permissions Required:
   *   - Manage Patients
   *
   * GET Get V2
   */
  public async patientsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.patientsGetV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Patients
   *
   * GET GetActive V1
   */
  public async patientsGetActiveV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.patientsGetActiveV1(body, licenseNumber));
  }

  /**
   * Retrieves a list of active patients for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Patients
   *
   * GET GetActive V2
   */
  public async patientsGetActiveV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.patientsGetActiveV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Updates Patient information for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Patients
   *
   * PUT Update V2
   */
  public async patientsUpdateV2(body: PatientsUpdateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.patientsUpdateV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Transfers
   *
   * POST CreateExternalIncoming V1
   */
  public async transfersCreateExternalIncomingV1(body: TransfersCreateExternalIncomingV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transfersCreateExternalIncomingV1(body, licenseNumber));
  }

  /**
   * Creates external incoming shipment plans for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *
   * POST CreateExternalIncoming V2
   */
  public async transfersCreateExternalIncomingV2(body: TransfersCreateExternalIncomingV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transfersCreateExternalIncomingV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Transfer Templates
   *
   * POST CreateTemplates V1
   */
  public async transfersCreateTemplatesV1(body: TransfersCreateTemplatesV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transfersCreateTemplatesV1(body, licenseNumber));
  }

  /**
   * Creates new transfer templates for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transfer Templates
   *
   * POST CreateTemplatesOutgoing V2
   */
  public async transfersCreateTemplatesOutgoingV2(body: TransfersCreateTemplatesOutgoingV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transfersCreateTemplatesOutgoingV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Transfers
   *
   * DELETE DeleteExternalIncoming V1
   */
  public async transfersDeleteExternalIncomingV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transfersDeleteExternalIncomingV1(id, body, licenseNumber));
  }

  /**
   * Voids an external incoming shipment plan for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *
   * DELETE DeleteExternalIncoming V2
   */
  public async transfersDeleteExternalIncomingV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transfersDeleteExternalIncomingV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Transfer Templates
   *
   * DELETE DeleteTemplates V1
   */
  public async transfersDeleteTemplatesV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transfersDeleteTemplatesV1(id, body, licenseNumber));
  }

  /**
   * Archives a transfer template for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transfer Templates
   *
   * DELETE DeleteTemplatesOutgoing V2
   */
  public async transfersDeleteTemplatesOutgoingV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transfersDeleteTemplatesOutgoingV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetDeliveriesPackagesStates V1
   */
  public async transfersGetDeliveriesPackagesStatesV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveriesPackagesStatesV1(body, No));
  }

  /**
   * Returns a list of available shipment Package states.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetDeliveriesPackagesStates V2
   */
  public async transfersGetDeliveriesPackagesStatesV2(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveriesPackagesStatesV2(body, No));
  }

  /**
   * Please note: that the {id} parameter above represents a Shipment Plan ID.
   * 
   *   Permissions Required:
   *   - Transfers
   *
   * GET GetDelivery V1
   */
  public async transfersGetDeliveryV1(id: string, body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryV1(id, body, No));
  }

  /**
   * Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *   - View Transfers
   *
   * GET GetDelivery V2
   */
  public async transfersGetDeliveryV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryV2(id, body, pageNumber, pageSize));
  }

  /**
   * Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Transfers
   *
   * GET GetDeliveryPackage V1
   */
  public async transfersGetDeliveryPackageV1(id: string, body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryPackageV1(id, body, No));
  }

  /**
   * Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *   - View Transfers
   *
   * GET GetDeliveryPackage V2
   */
  public async transfersGetDeliveryPackageV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryPackageV2(id, body, pageNumber, pageSize));
  }

  /**
   * Please note: The {id} parameter above represents a Transfer Delivery Package ID, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Transfers
   *
   * GET GetDeliveryPackageRequiredlabtestbatches V1
   */
  public async transfersGetDeliveryPackageRequiredlabtestbatchesV1(id: string, body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryPackageRequiredlabtestbatchesV1(id, body, No));
  }

  /**
   * Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *   - View Transfers
   *
   * GET GetDeliveryPackageRequiredlabtestbatches V2
   */
  public async transfersGetDeliveryPackageRequiredlabtestbatchesV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryPackageRequiredlabtestbatchesV2(id, body, pageNumber, pageSize));
  }

  /**
   * Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Transfers
   *
   * GET GetDeliveryPackageWholesale V1
   */
  public async transfersGetDeliveryPackageWholesaleV1(id: string, body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryPackageWholesaleV1(id, body, No));
  }

  /**
   * Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *   - View Transfers
   *
   * GET GetDeliveryPackageWholesale V2
   */
  public async transfersGetDeliveryPackageWholesaleV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryPackageWholesaleV2(id, body, pageNumber, pageSize));
  }

  /**
   * Please note: that the {id} parameter above represents a Shipment Delivery ID.
   * 
   *   Permissions Required:
   *   - Transfers
   *
   * GET GetDeliveryTransporters V1
   */
  public async transfersGetDeliveryTransportersV1(id: string, body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryTransportersV1(id, body, No));
  }

  /**
   * Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *   - View Transfers
   *
   * GET GetDeliveryTransporters V2
   */
  public async transfersGetDeliveryTransportersV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryTransportersV2(id, body, pageNumber, pageSize));
  }

  /**
   * Please note: The {id} parameter above represents a Shipment Delivery ID.
   * 
   *   Permissions Required:
   *   - Transfers
   *
   * GET GetDeliveryTransportersDetails V1
   */
  public async transfersGetDeliveryTransportersDetailsV1(id: string, body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryTransportersDetailsV1(id, body, No));
  }

  /**
   * Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *   - View Transfers
   *
   * GET GetDeliveryTransportersDetails V2
   */
  public async transfersGetDeliveryTransportersDetailsV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetDeliveryTransportersDetailsV2(id, body, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *   - View Transfers
   *
   * GET GetHub V2
   */
  public async transfersGetHubV2(body?: any, estimatedArrivalEnd?: string, estimatedArrivalStart?: string, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetHubV2(body, estimatedArrivalEnd, estimatedArrivalStart, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Transfers
   *
   * GET GetIncoming V1
   */
  public async transfersGetIncomingV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetIncomingV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *   - View Transfers
   *
   * GET GetIncoming V2
   */
  public async transfersGetIncomingV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetIncomingV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Transfers
   *
   * GET GetOutgoing V1
   */
  public async transfersGetOutgoingV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetOutgoingV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *   - View Transfers
   *
   * GET GetOutgoing V2
   */
  public async transfersGetOutgoingV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetOutgoingV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Transfers
   *
   * GET GetRejected V1
   */
  public async transfersGetRejectedV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetRejectedV1(body, licenseNumber));
  }

  /**
   * Retrieves a list of shipments with rejected packages for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *   - View Transfers
   *
   * GET GetRejected V2
   */
  public async transfersGetRejectedV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetRejectedV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Transfer Templates
   *
   * GET GetTemplates V1
   */
  public async transfersGetTemplatesV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Transfer Templates
   *
   * GET GetTemplatesDelivery V1
   */
  public async transfersGetTemplatesDeliveryV1(id: string, body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesDeliveryV1(id, body, No));
  }

  /**
   * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Transfers
   *
   * GET GetTemplatesDeliveryPackage V1
   */
  public async transfersGetTemplatesDeliveryPackageV1(id: string, body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesDeliveryPackageV1(id, body, No));
  }

  /**
   * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Transfer Templates
   *
   * GET GetTemplatesDeliveryTransporters V1
   */
  public async transfersGetTemplatesDeliveryTransportersV1(id: string, body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesDeliveryTransportersV1(id, body, No));
  }

  /**
   * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Transfer Templates
   *
   * GET GetTemplatesDeliveryTransportersDetails V1
   */
  public async transfersGetTemplatesDeliveryTransportersDetailsV1(id: string, body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesDeliveryTransportersDetailsV1(id, body, No));
  }

  /**
   * Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
   * 
   *   Permissions Required:
   *   - Manage Transfer Templates
   *   - View Transfer Templates
   *
   * GET GetTemplatesOutgoing V2
   */
  public async transfersGetTemplatesOutgoingV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesOutgoingV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of deliveries associated with a specific transfer template.
   * 
   *   Permissions Required:
   *   - Manage Transfer Templates
   *   - View Transfer Templates
   *
   * GET GetTemplatesOutgoingDelivery V2
   */
  public async transfersGetTemplatesOutgoingDeliveryV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesOutgoingDeliveryV2(id, body, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Manage Transfer Templates
   *   - View Transfer Templates
   *
   * GET GetTemplatesOutgoingDeliveryPackage V2
   */
  public async transfersGetTemplatesOutgoingDeliveryPackageV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesOutgoingDeliveryPackageV2(id, body, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Manage Transfer Templates
   *   - View Transfer Templates
   *
   * GET GetTemplatesOutgoingDeliveryTransporters V2
   */
  public async transfersGetTemplatesOutgoingDeliveryTransportersV2(id: string, body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesOutgoingDeliveryTransportersV2(id, body, pageNumber, pageSize));
  }

  /**
   * Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
   * 
   *   Permissions Required:
   *   - Manage Transfer Templates
   *   - View Transfer Templates
   *
   * GET GetTemplatesOutgoingDeliveryTransportersDetails V2
   */
  public async transfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id: string, body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id, body, No));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetTypes V1
   */
  public async transfersGetTypesV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetTypesV1(body, licenseNumber));
  }

  /**
   * Retrieves a list of available transfer types for a Facility based on its license number.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetTypes V2
   */
  public async transfersGetTypesV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transfersGetTypesV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Transfers
   *
   * PUT UpdateExternalIncoming V1
   */
  public async transfersUpdateExternalIncomingV1(body: TransfersUpdateExternalIncomingV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transfersUpdateExternalIncomingV1(body, licenseNumber));
  }

  /**
   * Updates external incoming shipment plans for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *
   * PUT UpdateExternalIncoming V2
   */
  public async transfersUpdateExternalIncomingV2(body: TransfersUpdateExternalIncomingV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transfersUpdateExternalIncomingV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Transfer Templates
   *
   * PUT UpdateTemplates V1
   */
  public async transfersUpdateTemplatesV1(body: TransfersUpdateTemplatesV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transfersUpdateTemplatesV1(body, licenseNumber));
  }

  /**
   * Updates existing transfer templates for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transfer Templates
   *
   * PUT UpdateTemplatesOutgoing V2
   */
  public async transfersUpdateTemplatesOutgoingV2(body: TransfersUpdateTemplatesOutgoingV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transfersUpdateTemplatesOutgoingV2(body, licenseNumber));
  }

  /**
   * Creates new driver records for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transporters
   *
   * POST CreateDriver V2
   */
  public async transportersCreateDriverV2(body: TransportersCreateDriverV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transportersCreateDriverV2(body, licenseNumber));
  }

  /**
   * Creates new vehicle records for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transporters
   *
   * POST CreateVehicle V2
   */
  public async transportersCreateVehicleV2(body: TransportersCreateVehicleV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transportersCreateVehicleV2(body, licenseNumber));
  }

  /**
   * Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
   * 
   *   Permissions Required:
   *   - Manage Transporters
   *
   * DELETE DeleteDriver V2
   */
  public async transportersDeleteDriverV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transportersDeleteDriverV2(id, body, licenseNumber));
  }

  /**
   * Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
   * 
   *   Permissions Required:
   *   - Manage Transporters
   *
   * DELETE DeleteVehicle V2
   */
  public async transportersDeleteVehicleV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transportersDeleteVehicleV2(id, body, licenseNumber));
  }

  /**
   * Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
   * 
   *   Permissions Required:
   *   - Transporters
   *
   * GET GetDriver V2
   */
  public async transportersGetDriverV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transportersGetDriverV2(id, body, licenseNumber));
  }

  /**
   * Retrieves a list of drivers for a Facility.
   * 
   *   Permissions Required:
   *   - Transporters
   *
   * GET GetDrivers V2
   */
  public async transportersGetDriversV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transportersGetDriversV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
   * 
   *   Permissions Required:
   *   - Transporters
   *
   * GET GetVehicle V2
   */
  public async transportersGetVehicleV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transportersGetVehicleV2(id, body, licenseNumber));
  }

  /**
   * Retrieves a list of vehicles for a Facility.
   * 
   *   Permissions Required:
   *   - Transporters
   *
   * GET GetVehicles V2
   */
  public async transportersGetVehiclesV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.transportersGetVehiclesV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Updates existing driver records for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transporters
   *
   * PUT UpdateDriver V2
   */
  public async transportersUpdateDriverV2(body: TransportersUpdateDriverV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transportersUpdateDriverV2(body, licenseNumber));
  }

  /**
   * Updates existing vehicle records for a facility.
   * 
   *   Permissions Required:
   *   - Manage Transporters
   *
   * PUT UpdateVehicle V2
   */
  public async transportersUpdateVehicleV2(body: TransportersUpdateVehicleV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.transportersUpdateVehicleV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetActive V1
   */
  public async unitsOfMeasureGetActiveV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.unitsOfMeasureGetActiveV1(body, No));
  }

  /**
   * Retrieves all active units of measure.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetActive V2
   */
  public async unitsOfMeasureGetActiveV2(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.unitsOfMeasureGetActiveV2(body, No));
  }

  /**
   * Retrieves all inactive units of measure.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetInactive V2
   */
  public async unitsOfMeasureGetInactiveV2(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.unitsOfMeasureGetInactiveV2(body, No));
  }

  /**
   * Retrieves all available waste methods.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetAll V2
   */
  public async wasteMethodsGetAllV2(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.wasteMethodsGetAllV2(body, No));
  }

  /**
   * Data returned by this endpoint is cached for up to one minute.
   * 
   *   Permissions Required:
   *   - Lookup Caregivers
   *
   * GET GetByCaregiverLicenseNumber V1
   */
  public async caregiversStatusGetByCaregiverLicenseNumberV1(caregiverLicenseNumber: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.caregiversStatusGetByCaregiverLicenseNumberV1(caregiverLicenseNumber, body, licenseNumber));
  }

  /**
   * Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
   * 
   *   Permissions Required:
   *   - Lookup Caregivers
   *
   * GET GetByCaregiverLicenseNumber V2
   */
  public async caregiversStatusGetByCaregiverLicenseNumberV2(caregiverLicenseNumber: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.caregiversStatusGetByCaregiverLicenseNumberV2(caregiverLicenseNumber, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Employees
   *
   * GET GetAll V1
   */
  public async employeesGetAllV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.employeesGetAllV1(body, licenseNumber));
  }

  /**
   * Retrieves a list of employees for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Employees
   *   - View Employees
   *
   * GET GetAll V2
   */
  public async employeesGetAllV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.employeesGetAllV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
   * 
   *   Permissions Required:
   *   - Manage Employees
   *
   * GET GetPermissions V2
   */
  public async employeesGetPermissionsV2(body?: any, employeeLicenseNumber?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.employeesGetPermissionsV2(body, employeeLicenseNumber, licenseNumber));
  }

  /**
   * NOTE: To include a photo with an item, first use POST /items/v1/photo to POST the photo, and then use the returned ID in the request body in this endpoint.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * POST Create V1
   */
  public async itemsCreateV1(body: ItemsCreateV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.itemsCreateV1(body, licenseNumber));
  }

  /**
   * Creates one or more new products for the specified Facility. NOTE: To include a photo with an item, first use POST /items/v2/photo to POST the photo, and then use the returned Id in the request body in this endpoint.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * POST Create V2
   */
  public async itemsCreateV2(body: ItemsCreateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.itemsCreateV2(body, licenseNumber));
  }

  /**
   * Creates one or more new item brands for the specified Facility identified by the License Number.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * POST CreateBrand V2
   */
  public async itemsCreateBrandV2(body: ItemsCreateBrandV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.itemsCreateBrandV2(body, licenseNumber));
  }

  /**
   * Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * POST CreateFile V2
   */
  public async itemsCreateFileV2(body: ItemsCreateFileV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.itemsCreateFileV2(body, licenseNumber));
  }

  /**
   * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * POST CreatePhoto V1
   */
  public async itemsCreatePhotoV1(body: ItemsCreatePhotoV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.itemsCreatePhotoV1(body, licenseNumber));
  }

  /**
   * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * POST CreatePhoto V2
   */
  public async itemsCreatePhotoV2(body: ItemsCreatePhotoV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.itemsCreatePhotoV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * POST CreateUpdate V1
   */
  public async itemsCreateUpdateV1(body: ItemsCreateUpdateV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.itemsCreateUpdateV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * DELETE Delete V1
   */
  public async itemsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.itemsDeleteV1(id, body, licenseNumber));
  }

  /**
   * Archives the specified Product by Id for the given Facility License Number.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * DELETE Delete V2
   */
  public async itemsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.itemsDeleteV2(id, body, licenseNumber));
  }

  /**
   * Archives the specified Item Brand by Id for the given Facility License Number.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * DELETE DeleteBrand V2
   */
  public async itemsDeleteBrandV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.itemsDeleteBrandV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * GET Get V1
   */
  public async itemsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.itemsGetV1(id, body, licenseNumber));
  }

  /**
   * Retrieves detailed information about a specific Item by Id.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * GET Get V2
   */
  public async itemsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.itemsGetV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * GET GetActive V1
   */
  public async itemsGetActiveV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.itemsGetActiveV1(body, licenseNumber));
  }

  /**
   * Returns a list of active items for the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * GET GetActive V2
   */
  public async itemsGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.itemsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * GET GetBrands V1
   */
  public async itemsGetBrandsV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.itemsGetBrandsV1(body, licenseNumber));
  }

  /**
   * Retrieves a list of active item brands for the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * GET GetBrands V2
   */
  public async itemsGetBrandsV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.itemsGetBrandsV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetCategories V1
   */
  public async itemsGetCategoriesV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.itemsGetCategoriesV1(body, licenseNumber));
  }

  /**
   * Retrieves a list of item categories.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetCategories V2
   */
  public async itemsGetCategoriesV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.itemsGetCategoriesV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves a file by its Id for the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * GET GetFile V2
   */
  public async itemsGetFileV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.itemsGetFileV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * GET GetInactive V1
   */
  public async itemsGetInactiveV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.itemsGetInactiveV1(body, licenseNumber));
  }

  /**
   * Retrieves a list of inactive items for the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * GET GetInactive V2
   */
  public async itemsGetInactiveV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.itemsGetInactiveV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * GET GetPhoto V1
   */
  public async itemsGetPhotoV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.itemsGetPhotoV1(id, body, licenseNumber));
  }

  /**
   * Retrieves an image by its Id for the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * GET GetPhoto V2
   */
  public async itemsGetPhotoV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.itemsGetPhotoV2(id, body, licenseNumber));
  }

  /**
   * Updates one or more existing products for the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * PUT Update V2
   */
  public async itemsUpdateV2(body: ItemsUpdateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.itemsUpdateV2(body, licenseNumber));
  }

  /**
   * Updates one or more existing item brands for the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * PUT UpdateBrand V2
   */
  public async itemsUpdateBrandV2(body: ItemsUpdateBrandV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.itemsUpdateBrandV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST Create V1
   */
  public async packagesCreateV1(body: PackagesCreateV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesCreateV1(body, licenseNumber));
  }

  /**
   * Creates new packages for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST Create V2
   */
  public async packagesCreateV2(body: PackagesCreateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesCreateV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreateAdjust V1
   */
  public async packagesCreateAdjustV1(body: PackagesCreateAdjustV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesCreateAdjustV1(body, licenseNumber));
  }

  /**
   * Records a list of adjustments for packages at a specific Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreateAdjust V2
   */
  public async packagesCreateAdjustV2(body: PackagesCreateAdjustV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesCreateAdjustV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreateChangeItem V1
   */
  public async packagesCreateChangeItemV1(body: PackagesCreateChangeItemV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesCreateChangeItemV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreateChangeLocation V1
   */
  public async packagesCreateChangeLocationV1(body: PackagesCreateChangeLocationV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesCreateChangeLocationV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreateFinish V1
   */
  public async packagesCreateFinishV1(body: PackagesCreateFinishV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesCreateFinishV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreatePlantings V1
   */
  public async packagesCreatePlantingsV1(body: PackagesCreatePlantingsV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesCreatePlantingsV1(body, licenseNumber));
  }

  /**
   * Creates new plantings from packages for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreatePlantings V2
   */
  public async packagesCreatePlantingsV2(body: PackagesCreatePlantingsV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesCreatePlantingsV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreateRemediate V1
   */
  public async packagesCreateRemediateV1(body: PackagesCreateRemediateV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesCreateRemediateV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreateTesting V1
   */
  public async packagesCreateTestingV1(body: PackagesCreateTestingV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesCreateTestingV1(body, licenseNumber));
  }

  /**
   * Creates new packages for testing for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreateTesting V2
   */
  public async packagesCreateTestingV2(body: PackagesCreateTestingV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesCreateTestingV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreateUnfinish V1
   */
  public async packagesCreateUnfinishV1(body: PackagesCreateUnfinishV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesCreateUnfinishV1(body, licenseNumber));
  }

  /**
   * Discontinues a Package at a specific Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * DELETE Delete V2
   */
  public async packagesDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesDeleteV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *
   * GET Get V1
   */
  public async packagesGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetV1(id, body, licenseNumber));
  }

  /**
   * Retrieves a Package by its Id.
   * 
   *   Permissions Required:
   *   - View Packages
   *
   * GET Get V2
   */
  public async packagesGetV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *
   * GET GetActive V1
   */
  public async packagesGetActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves a list of active packages for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *
   * GET GetActive V2
   */
  public async packagesGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetAdjustReasons V1
   */
  public async packagesGetAdjustReasonsV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetAdjustReasonsV1(body, licenseNumber));
  }

  /**
   * Retrieves a list of adjustment reasons for packages at a specified Facility.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetAdjustReasons V2
   */
  public async packagesGetAdjustReasonsV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetAdjustReasonsV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *
   * GET GetByLabel V1
   */
  public async packagesGetByLabelV1(label: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetByLabelV1(label, body, licenseNumber));
  }

  /**
   * Retrieves a Package by its label.
   * 
   *   Permissions Required:
   *   - View Packages
   *
   * GET GetByLabel V2
   */
  public async packagesGetByLabelV2(label: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetByLabelV2(label, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *
   * GET GetInactive V1
   */
  public async packagesGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves a list of inactive packages for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *
   * GET GetInactive V2
   */
  public async packagesGetInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of packages in transit for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *
   * GET GetIntransit V2
   */
  public async packagesGetIntransitV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetIntransitV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *
   * GET GetLabsamples V2
   */
  public async packagesGetLabsamplesV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetLabsamplesV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *
   * GET GetOnhold V1
   */
  public async packagesGetOnholdV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetOnholdV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves a list of packages on hold for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *
   * GET GetOnhold V2
   */
  public async packagesGetOnholdV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves the source harvests for a Package by its Id.
   * 
   *   Permissions Required:
   *   - View Package Source Harvests
   *
   * GET GetSourceHarvest V2
   */
  public async packagesGetSourceHarvestV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetSourceHarvestV2(id, body, licenseNumber));
  }

  /**
   * Retrieves a list of transferred packages for a specific Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *
   * GET GetTransferred V2
   */
  public async packagesGetTransferredV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetTransferredV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetTypes V1
   */
  public async packagesGetTypesV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetTypesV1(body, No));
  }

  /**
   * Retrieves a list of available Package types.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetTypes V2
   */
  public async packagesGetTypesV2(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.packagesGetTypesV2(body, No));
  }

  /**
   * Set the final quantity for a Package.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateAdjust V2
   */
  public async packagesUpdateAdjustV2(body: PackagesUpdateAdjustV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateAdjustV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *   - Manage Package Notes
   *
   * PUT UpdateChangeNote V1
   */
  public async packagesUpdateChangeNoteV1(body: PackagesUpdateChangeNoteV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateChangeNoteV1(body, licenseNumber));
  }

  /**
   * Updates the Product decontaminate information for a list of packages at a specific Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateDecontaminate V2
   */
  public async packagesUpdateDecontaminateV2(body: PackagesUpdateDecontaminateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateDecontaminateV2(body, licenseNumber));
  }

  /**
   * Flags one or more packages for donation at the specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateDonationFlag V2
   */
  public async packagesUpdateDonationFlagV2(body: PackagesUpdateDonationFlagV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateDonationFlagV2(body, licenseNumber));
  }

  /**
   * Removes the donation flag from one or more packages at the specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateDonationUnflag V2
   */
  public async packagesUpdateDonationUnflagV2(body: PackagesUpdateDonationUnflagV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateDonationUnflagV2(body, licenseNumber));
  }

  /**
   * Updates the external identifiers for one or more packages at the specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Package Inventory
   *   - External Id Enabled
   *
   * PUT UpdateExternalid V2
   */
  public async packagesUpdateExternalidV2(body: PackagesUpdateExternalidV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateExternalidV2(body, licenseNumber));
  }

  /**
   * Updates a list of packages as finished for a specific Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateFinish V2
   */
  public async packagesUpdateFinishV2(body: PackagesUpdateFinishV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateFinishV2(body, licenseNumber));
  }

  /**
   * Updates the associated Item for one or more packages at the specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * PUT UpdateItem V2
   */
  public async packagesUpdateItemV2(body: PackagesUpdateItemV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateItemV2(body, licenseNumber));
  }

  /**
   * Updates the list of required lab test batches for one or more packages at the specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * PUT UpdateLabTestRequired V2
   */
  public async packagesUpdateLabTestRequiredV2(body: PackagesUpdateLabTestRequiredV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateLabTestRequiredV2(body, licenseNumber));
  }

  /**
   * Updates the Location and Sublocation for one or more packages at the specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * PUT UpdateLocation V2
   */
  public async packagesUpdateLocationV2(body: PackagesUpdateLocationV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateLocationV2(body, licenseNumber));
  }

  /**
   * Updates notes associated with one or more packages for the specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *   - Manage Package Notes
   *
   * PUT UpdateNote V2
   */
  public async packagesUpdateNoteV2(body: PackagesUpdateNoteV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateNoteV2(body, licenseNumber));
  }

  /**
   * Updates a list of Product remediations for packages at a specific Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateRemediate V2
   */
  public async packagesUpdateRemediateV2(body: PackagesUpdateRemediateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateRemediateV2(body, licenseNumber));
  }

  /**
   * Flags or unflags one or more packages at the specified Facility as trade samples.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateTradesampleFlag V2
   */
  public async packagesUpdateTradesampleFlagV2(body: PackagesUpdateTradesampleFlagV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateTradesampleFlagV2(body, licenseNumber));
  }

  /**
   * Removes the trade sample flag from one or more packages at the specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateTradesampleUnflag V2
   */
  public async packagesUpdateTradesampleUnflagV2(body: PackagesUpdateTradesampleUnflagV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateTradesampleUnflagV2(body, licenseNumber));
  }

  /**
   * Updates a list of packages as unfinished for a specific Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateUnfinish V2
   */
  public async packagesUpdateUnfinishV2(body: PackagesUpdateUnfinishV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateUnfinishV2(body, licenseNumber));
  }

  /**
   * Updates the use-by date for one or more packages at the specified Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * PUT UpdateUsebydate V2
   */
  public async packagesUpdateUsebydateV2(body: PackagesUpdateUsebydateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.packagesUpdateUsebydateV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * POST Create V1
   */
  public async patientCheckInsCreateV1(body: PatientcheckinsCreateV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.patientCheckInsCreateV1(body, licenseNumber));
  }

  /**
   * Records patient check-ins for a specified Facility.
   * 
   *   Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * POST Create V2
   */
  public async patientCheckInsCreateV2(body: PatientcheckinsCreateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.patientCheckInsCreateV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * DELETE Delete V1
   */
  public async patientCheckInsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.patientCheckInsDeleteV1(id, body, licenseNumber));
  }

  /**
   * Archives a Patient Check-In, identified by its Id, for a specified Facility.
   * 
   *   Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * DELETE Delete V2
   */
  public async patientCheckInsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.patientCheckInsDeleteV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * GET GetAll V1
   */
  public async patientCheckInsGetAllV1(body?: any, checkinDateEnd?: string, checkinDateStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.patientCheckInsGetAllV1(body, checkinDateEnd, checkinDateStart, licenseNumber));
  }

  /**
   * Retrieves a list of patient check-ins for a specified Facility.
   * 
   *   Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * GET GetAll V2
   */
  public async patientCheckInsGetAllV2(body?: any, checkinDateEnd?: string, checkinDateStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.patientCheckInsGetAllV2(body, checkinDateEnd, checkinDateStart, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetLocations V1
   */
  public async patientCheckInsGetLocationsV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.patientCheckInsGetLocationsV1(body, No));
  }

  /**
   * Retrieves a list of Patient Check-In locations.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetLocations V2
   */
  public async patientCheckInsGetLocationsV2(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.patientCheckInsGetLocationsV2(body, No));
  }

  /**
   * Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * PUT Update V1
   */
  public async patientCheckInsUpdateV1(body: PatientcheckinsUpdateV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.patientCheckInsUpdateV1(body, licenseNumber));
  }

  /**
   * Updates patient check-ins for a specified Facility.
   * 
   *   Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * PUT Update V2
   */
  public async patientCheckInsUpdateV2(body: PatientcheckinsUpdateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.patientCheckInsUpdateV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditives V1
   */
  public async plantBatchesCreateAdditivesV1(body: PlantbatchesCreateAdditivesV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateAdditivesV1(body, licenseNumber));
  }

  /**
   * Records Additive usage details for plant batches at a specific Facility.
   * 
   *   Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditives V2
   */
  public async plantBatchesCreateAdditivesV2(body: PlantbatchesCreateAdditivesV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateAdditivesV2(body, licenseNumber));
  }

  /**
   * Records Additive usage for plant batches at a Facility using predefined additive templates.
   * 
   *   Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditivesUsingtemplate V2
   */
  public async plantBatchesCreateAdditivesUsingtemplateV2(body: PlantbatchesCreateAdditivesUsingtemplateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateAdditivesUsingtemplateV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *
   * POST CreateAdjust V1
   */
  public async plantBatchesCreateAdjustV1(body: PlantbatchesCreateAdjustV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateAdjustV1(body, licenseNumber));
  }

  /**
   * Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *
   * POST CreateAdjust V2
   */
  public async plantBatchesCreateAdjustV2(body: PlantbatchesCreateAdjustV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateAdjustV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * POST CreateChangegrowthphase V1
   */
  public async plantBatchesCreateChangegrowthphaseV1(body: PlantbatchesCreateChangegrowthphaseV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateChangegrowthphaseV1(body, licenseNumber));
  }

  /**
   * Updates the growth phase of plants at a specified Facility based on tracking information.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * POST CreateGrowthphase V2
   */
  public async plantBatchesCreateGrowthphaseV2(body: PlantbatchesCreateGrowthphaseV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateGrowthphaseV2(body, licenseNumber));
  }

  /**
   * Creates packages from plant batches at a Facility, with optional support for packaging from mother plants.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreatePackage V2
   */
  public async plantBatchesCreatePackageV2(body: PlantbatchesCreatePackageV2RequestItem[], isFromMotherPlant?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreatePackageV2(body, isFromMotherPlant, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreatePackageFrommotherplant V1
   */
  public async plantBatchesCreatePackageFrommotherplantV1(body: PlantbatchesCreatePackageFrommotherplantV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreatePackageFrommotherplantV1(body, licenseNumber));
  }

  /**
   * Creates packages from mother plants at the specified Facility.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreatePackageFrommotherplant V2
   */
  public async plantBatchesCreatePackageFrommotherplantV2(body: PlantbatchesCreatePackageFrommotherplantV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreatePackageFrommotherplantV2(body, licenseNumber));
  }

  /**
   * Creates new plantings for a Facility by generating plant batches based on provided planting details.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *
   * POST CreatePlantings V2
   */
  public async plantBatchesCreatePlantingsV2(body: PlantbatchesCreatePlantingsV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreatePlantingsV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *
   * POST CreateSplit V1
   */
  public async plantBatchesCreateSplitV1(body: PlantbatchesCreateSplitV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateSplitV1(body, licenseNumber));
  }

  /**
   * Splits an existing Plant Batch into multiple groups at the specified Facility.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *
   * POST CreateSplit V2
   */
  public async plantBatchesCreateSplitV2(body: PlantbatchesCreateSplitV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateSplitV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Plants Waste
   *
   * POST CreateWaste V1
   */
  public async plantBatchesCreateWasteV1(body: PlantbatchesCreateWasteV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateWasteV1(body, licenseNumber));
  }

  /**
   * Records waste information for plant batches based on the submitted data for the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Plants Waste
   *
   * POST CreateWaste V2
   */
  public async plantBatchesCreateWasteV2(body: PlantbatchesCreateWasteV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateWasteV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST Createpackages V1
   */
  public async plantBatchesCreatepackagesV1(body: PlantbatchesCreatepackagesV1RequestItem[], isFromMotherPlant?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreatepackagesV1(body, isFromMotherPlant, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *
   * POST Createplantings V1
   */
  public async plantBatchesCreateplantingsV1(body: PlantbatchesCreateplantingsV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesCreateplantingsV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Destroy Immature Plants
   *
   * DELETE Delete V1
   */
  public async plantBatchesDeleteV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesDeleteV1(body, licenseNumber));
  }

  /**
   * Completes the destruction of plant batches based on the provided input data.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *   - Destroy Immature Plants
   *
   * DELETE Delete V2
   */
  public async plantBatchesDeleteV2(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesDeleteV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *
   * GET Get V1
   */
  public async plantBatchesGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetV1(id, body, licenseNumber));
  }

  /**
   * Retrieves a Plant Batch by Id.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *
   * GET Get V2
   */
  public async plantBatchesGetV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *
   * GET GetActive V1
   */
  public async plantBatchesGetActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *
   * GET GetActive V2
   */
  public async plantBatchesGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *
   * GET GetInactive V1
   */
  public async plantBatchesGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *
   * GET GetInactive V2
   */
  public async plantBatchesGetInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetTypes V1
   */
  public async plantBatchesGetTypesV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetTypesV1(body, No));
  }

  /**
   * Retrieves a list of plant batch types.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetTypes V2
   */
  public async plantBatchesGetTypesV2(body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetTypesV2(body, pageNumber, pageSize));
  }

  /**
   * Retrieves waste details associated with plant batches at a specified Facility.
   * 
   *   Permissions Required:
   *   - View Plants Waste
   *
   * GET GetWaste V2
   */
  public async plantBatchesGetWasteV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetWasteV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetWasteReasons V1
   */
  public async plantBatchesGetWasteReasonsV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetWasteReasonsV1(body, licenseNumber));
  }

  /**
   * Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetWasteReasons V2
   */
  public async plantBatchesGetWasteReasonsV2(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.plantBatchesGetWasteReasonsV2(body, licenseNumber));
  }

  /**
   * Moves one or more plant batches to new locations with in a specified Facility.
   * 
   *   Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants
   *
   * PUT UpdateLocation V2
   */
  public async plantBatchesUpdateLocationV2(body: PlantbatchesUpdateLocationV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesUpdateLocationV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *
   * PUT UpdateMoveplantbatches V1
   */
  public async plantBatchesUpdateMoveplantbatchesV1(body: PlantbatchesUpdateMoveplantbatchesV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesUpdateMoveplantbatchesV1(body, licenseNumber));
  }

  /**
   * Renames plant batches at a specified Facility.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * PUT UpdateName V2
   */
  public async plantBatchesUpdateNameV2(body: PlantbatchesUpdateNameV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesUpdateNameV2(body, licenseNumber));
  }

  /**
   * Changes the strain of plant batches at a specified Facility.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * PUT UpdateStrain V2
   */
  public async plantBatchesUpdateStrainV2(body: PlantbatchesUpdateStrainV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesUpdateStrainV2(body, licenseNumber));
  }

  /**
   * Replaces tags for plant batches at a specified Facility.
   * 
   *   Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * PUT UpdateTag V2
   */
  public async plantBatchesUpdateTagV2(body: PlantbatchesUpdateTagV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.plantBatchesUpdateTagV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - ManageProcessingJobs
   *
   * POST CreateAdjust V1
   */
  public async processingJobsCreateAdjustV1(body: ProcessingjobsCreateAdjustV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreateAdjustV1(body, licenseNumber));
  }

  /**
   * Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * POST CreateAdjust V2
   */
  public async processingJobsCreateAdjustV2(body: ProcessingjobsCreateAdjustV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreateAdjustV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * POST CreateJobtypes V1
   */
  public async processingJobsCreateJobtypesV1(body: ProcessingjobsCreateJobtypesV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreateJobtypesV1(body, licenseNumber));
  }

  /**
   * Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * POST CreateJobtypes V2
   */
  public async processingJobsCreateJobtypesV2(body: ProcessingjobsCreateJobtypesV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreateJobtypesV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - ManageProcessingJobs
   *
   * POST CreateStart V1
   */
  public async processingJobsCreateStartV1(body: ProcessingjobsCreateStartV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreateStartV1(body, licenseNumber));
  }

  /**
   * Initiates new processing jobs at a Facility, including job details and associated packages.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * POST CreateStart V2
   */
  public async processingJobsCreateStartV2(body: ProcessingjobsCreateStartV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreateStartV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - ManageProcessingJobs
   *
   * POST Createpackages V1
   */
  public async processingJobsCreatepackagesV1(body: ProcessingjobsCreatepackagesV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreatepackagesV1(body, licenseNumber));
  }

  /**
   * Creates packages from processing jobs at a Facility, including optional location and note assignments.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * POST Createpackages V2
   */
  public async processingJobsCreatepackagesV2(body: ProcessingjobsCreatepackagesV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsCreatepackagesV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * DELETE Delete V1
   */
  public async processingJobsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsDeleteV1(id, body, licenseNumber));
  }

  /**
   * Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * DELETE Delete V2
   */
  public async processingJobsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsDeleteV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * DELETE DeleteJobtypes V1
   */
  public async processingJobsDeleteJobtypesV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsDeleteJobtypesV1(id, body, licenseNumber));
  }

  /**
   * Archives a Processing Job Type at a Facility, making it inactive for future use.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * DELETE DeleteJobtypes V2
   */
  public async processingJobsDeleteJobtypesV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsDeleteJobtypesV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET Get V1
   */
  public async processingJobsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetV1(id, body, licenseNumber));
  }

  /**
   * Retrieves a ProcessingJob by Id.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * GET Get V2
   */
  public async processingJobsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetActive V1
   */
  public async processingJobsGetActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves active processing jobs at a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetActive V2
   */
  public async processingJobsGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetInactive V1
   */
  public async processingJobsGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves inactive processing jobs at a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetInactive V2
   */
  public async processingJobsGetInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetJobtypesActive V1
   */
  public async processingJobsGetJobtypesActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves a list of all active processing job types defined within a Facility.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetJobtypesActive V2
   */
  public async processingJobsGetJobtypesActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetJobtypesAttributes V1
   */
  public async processingJobsGetJobtypesAttributesV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesAttributesV1(body, licenseNumber));
  }

  /**
   * Retrieves all processing job attributes available for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetJobtypesAttributes V2
   */
  public async processingJobsGetJobtypesAttributesV2(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesAttributesV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetJobtypesCategories V1
   */
  public async processingJobsGetJobtypesCategoriesV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesCategoriesV1(body, licenseNumber));
  }

  /**
   * Retrieves all processing job categories available for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetJobtypesCategories V2
   */
  public async processingJobsGetJobtypesCategoriesV2(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesCategoriesV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetJobtypesInactive V1
   */
  public async processingJobsGetJobtypesInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves a list of all inactive processing job types defined within a Facility.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetJobtypesInactive V2
   */
  public async processingJobsGetJobtypesInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.processingJobsGetJobtypesInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * PUT UpdateFinish V1
   */
  public async processingJobsUpdateFinishV1(body: ProcessingjobsUpdateFinishV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsUpdateFinishV1(body, licenseNumber));
  }

  /**
   * Completes processing jobs at a Facility by recording final notes and waste measurements.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * PUT UpdateFinish V2
   */
  public async processingJobsUpdateFinishV2(body: ProcessingjobsUpdateFinishV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsUpdateFinishV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * PUT UpdateJobtypes V1
   */
  public async processingJobsUpdateJobtypesV1(body: ProcessingjobsUpdateJobtypesV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsUpdateJobtypesV1(body, licenseNumber));
  }

  /**
   * Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * PUT UpdateJobtypes V2
   */
  public async processingJobsUpdateJobtypesV2(body: ProcessingjobsUpdateJobtypesV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsUpdateJobtypesV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * PUT UpdateUnfinish V1
   */
  public async processingJobsUpdateUnfinishV1(body: ProcessingjobsUpdateUnfinishV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsUpdateUnfinishV1(body, licenseNumber));
  }

  /**
   * Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * PUT UpdateUnfinish V2
   */
  public async processingJobsUpdateUnfinishV2(body: ProcessingjobsUpdateUnfinishV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.processingJobsUpdateUnfinishV2(body, licenseNumber));
  }

  /**
   * Creates new sublocation records for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * POST Create V2
   */
  public async sublocationsCreateV2(body: SublocationsCreateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.sublocationsCreateV2(body, licenseNumber));
  }

  /**
   * Archives an existing Sublocation record for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * DELETE Delete V2
   */
  public async sublocationsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.sublocationsDeleteV2(id, body, licenseNumber));
  }

  /**
   * Retrieves a Sublocation by its Id, with an optional license number.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * GET Get V2
   */
  public async sublocationsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.sublocationsGetV2(id, body, licenseNumber));
  }

  /**
   * Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * GET GetActive V2
   */
  public async sublocationsGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.sublocationsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of inactive sublocations for the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * GET GetInactive V2
   */
  public async sublocationsGetInactiveV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.sublocationsGetInactiveV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Updates existing sublocation records for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * PUT Update V2
   */
  public async sublocationsUpdateV2(body: SublocationsUpdateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.sublocationsUpdateV2(body, licenseNumber));
  }

  /**
   * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Sales Delivery
   *
   * POST CreateDelivery V1
   */
  public async salesCreateDeliveryV1(body: SalesCreateDeliveryV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryV1(body, licenseNumber));
  }

  /**
   * Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
   *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
   *
   * POST CreateDelivery V2
   */
  public async salesCreateDeliveryV2(body: SalesCreateDeliveryV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryV2(body, licenseNumber));
  }

  /**
   * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Retailer Delivery
   *
   * POST CreateDeliveryRetailer V1
   */
  public async salesCreateDeliveryRetailerV1(body: SalesCreateDeliveryRetailerV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerV1(body, licenseNumber));
  }

  /**
   * Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   *   - Industry/Facility Type/Retailer Delivery
   *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
   *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
   *   - Manage Retailer Delivery
   *
   * POST CreateDeliveryRetailer V2
   */
  public async salesCreateDeliveryRetailerV2(body: SalesCreateDeliveryRetailerV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Retailer Delivery
   *
   * POST CreateDeliveryRetailerDepart V1
   */
  public async salesCreateDeliveryRetailerDepartV1(body: SalesCreateDeliveryRetailerDepartV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerDepartV1(body, licenseNumber));
  }

  /**
   * Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   *   - Industry/Facility Type/Retailer Delivery
   *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   *   - Manage Retailer Delivery
   *
   * POST CreateDeliveryRetailerDepart V2
   */
  public async salesCreateDeliveryRetailerDepartV2(body: SalesCreateDeliveryRetailerDepartV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerDepartV2(body, licenseNumber));
  }

  /**
   * Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Retailer Delivery
   *
   * POST CreateDeliveryRetailerEnd V1
   */
  public async salesCreateDeliveryRetailerEndV1(body: SalesCreateDeliveryRetailerEndV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerEndV1(body, licenseNumber));
  }

  /**
   * Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   *   - Industry/Facility Type/Retailer Delivery
   *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   *   - Manage Retailer Delivery
   *
   * POST CreateDeliveryRetailerEnd V2
   */
  public async salesCreateDeliveryRetailerEndV2(body: SalesCreateDeliveryRetailerEndV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerEndV2(body, licenseNumber));
  }

  /**
   * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Retailer Delivery
   *
   * POST CreateDeliveryRetailerRestock V1
   */
  public async salesCreateDeliveryRetailerRestockV1(body: SalesCreateDeliveryRetailerRestockV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerRestockV1(body, licenseNumber));
  }

  /**
   * Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   *   - Industry/Facility Type/Retailer Delivery
   *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   *   - Manage Retailer Delivery
   *
   * POST CreateDeliveryRetailerRestock V2
   */
  public async salesCreateDeliveryRetailerRestockV2(body: SalesCreateDeliveryRetailerRestockV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerRestockV2(body, licenseNumber));
  }

  /**
   * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Retailer Delivery
   *
   * POST CreateDeliveryRetailerSale V1
   */
  public async salesCreateDeliveryRetailerSaleV1(body: SalesCreateDeliveryRetailerSaleV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerSaleV1(body, licenseNumber));
  }

  /**
   * Records sales deliveries originating from a retailer delivery for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   *   - Industry/Facility Type/Retailer Delivery
   *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
   *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
   *
   * POST CreateDeliveryRetailerSale V2
   */
  public async salesCreateDeliveryRetailerSaleV2(body: SalesCreateDeliveryRetailerSaleV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateDeliveryRetailerSaleV2(body, licenseNumber));
  }

  /**
   * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Sales
   *
   * POST CreateReceipt V1
   */
  public async salesCreateReceiptV1(body: SalesCreateReceiptV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateReceiptV1(body, licenseNumber));
  }

  /**
   * Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Sales (Write)
   *   - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
   *   - Industry/Facility Type/Advanced Sales
   *   - WebApi Sales Read Write State (All or WriteOnly)
   *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
   *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
   *
   * POST CreateReceipt V2
   */
  public async salesCreateReceiptV2(body: SalesCreateReceiptV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateReceiptV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * POST CreateTransactionByDate V1
   */
  public async salesCreateTransactionByDateV1(date: string, body: SalesCreateTransactionByDateV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesCreateTransactionByDateV1(date, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Sales Delivery
   *
   * DELETE DeleteDelivery V1
   */
  public async salesDeleteDeliveryV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesDeleteDeliveryV1(id, body, licenseNumber));
  }

  /**
   * Voids a sales delivery for a Facility using the provided License Number and delivery Id.
   * 
   *   Permissions Required:
   *   - Manage Sales Delivery
   *
   * DELETE DeleteDelivery V2
   */
  public async salesDeleteDeliveryV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesDeleteDeliveryV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Retailer Delivery
   *
   * DELETE DeleteDeliveryRetailer V1
   */
  public async salesDeleteDeliveryRetailerV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesDeleteDeliveryRetailerV1(id, body, licenseNumber));
  }

  /**
   * Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   *   - Industry/Facility Type/Retailer Delivery
   *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   *   - Manage Retailer Delivery
   *
   * DELETE DeleteDeliveryRetailer V2
   */
  public async salesDeleteDeliveryRetailerV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesDeleteDeliveryRetailerV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * DELETE DeleteReceipt V1
   */
  public async salesDeleteReceiptV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesDeleteReceiptV1(id, body, licenseNumber));
  }

  /**
   * Archives a sales receipt for a Facility using the provided License Number and receipt Id.
   * 
   *   Permissions Required:
   *   - Manage Sales
   *
   * DELETE DeleteReceipt V2
   */
  public async salesDeleteReceiptV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesDeleteReceiptV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetCounties V1
   */
  public async salesGetCountiesV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetCountiesV1(body, No));
  }

  /**
   * Returns a list of counties available for sales deliveries.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetCounties V2
   */
  public async salesGetCountiesV2(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetCountiesV2(body, No));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetCustomertypes V1
   */
  public async salesGetCustomertypesV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetCustomertypesV1(body, No));
  }

  /**
   * Returns a list of customer types.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetCustomertypes V2
   */
  public async salesGetCustomertypesV2(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetCustomertypesV2(body, No));
  }

  /**
   * Permissions Required:
   *   - Sales Delivery
   *
   * GET GetDeliveriesActive V1
   */
  public async salesGetDeliveriesActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart));
  }

  /**
   * Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
   * 
   *   Permissions Required:
   *   - View Sales Delivery
   *   - Manage Sales Delivery
   *
   * GET GetDeliveriesActive V2
   */
  public async salesGetDeliveriesActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart));
  }

  /**
   * Permissions Required:
   *   - Sales Delivery
   *
   * GET GetDeliveriesInactive V1
   */
  public async salesGetDeliveriesInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart));
  }

  /**
   * Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
   * 
   *   Permissions Required:
   *   - View Sales Delivery
   *   - Manage Sales Delivery
   *
   * GET GetDeliveriesInactive V2
   */
  public async salesGetDeliveriesInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart));
  }

  /**
   * Permissions Required:
   *   - Retailer Delivery
   *
   * GET GetDeliveriesRetailerActive V1
   */
  public async salesGetDeliveriesRetailerActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesRetailerActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
   * 
   *   Permissions Required:
   *   - View Retailer Delivery
   *   - Manage Retailer Delivery
   *
   * GET GetDeliveriesRetailerActive V2
   */
  public async salesGetDeliveriesRetailerActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesRetailerActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Retailer Delivery
   *
   * GET GetDeliveriesRetailerInactive V1
   */
  public async salesGetDeliveriesRetailerInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesRetailerInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
   * 
   *   Permissions Required:
   *   - View Retailer Delivery
   *   - Manage Retailer Delivery
   *
   * GET GetDeliveriesRetailerInactive V2
   */
  public async salesGetDeliveriesRetailerInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesRetailerInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   -
   *
   * GET GetDeliveriesReturnreasons V1
   */
  public async salesGetDeliveriesReturnreasonsV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesReturnreasonsV1(body, licenseNumber));
  }

  /**
   * Returns a list of return reasons for sales deliveries based on the provided License Number.
   * 
   *   Permissions Required:
   *   - Sales Delivery
   *
   * GET GetDeliveriesReturnreasons V2
   */
  public async salesGetDeliveriesReturnreasonsV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveriesReturnreasonsV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Sales Delivery
   *
   * GET GetDelivery V1
   */
  public async salesGetDeliveryV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveryV1(id, body, licenseNumber));
  }

  /**
   * Retrieves a sales delivery record by its Id, with an optional License Number.
   * 
   *   Permissions Required:
   *   - View Sales Delivery
   *   - Manage Sales Delivery
   *
   * GET GetDelivery V2
   */
  public async salesGetDeliveryV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveryV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Retailer Delivery
   *
   * GET GetDeliveryRetailer V1
   */
  public async salesGetDeliveryRetailerV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveryRetailerV1(id, body, licenseNumber));
  }

  /**
   * Retrieves a retailer delivery record by its ID, with an optional License Number.
   * 
   *   Permissions Required:
   *   - View Retailer Delivery
   *   - Manage Retailer Delivery
   *
   * GET GetDeliveryRetailer V2
   */
  public async salesGetDeliveryRetailerV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetDeliveryRetailerV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   -
   *
   * GET GetPatientRegistrationsLocations V1
   */
  public async salesGetPatientRegistrationsLocationsV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetPatientRegistrationsLocationsV1(body, No));
  }

  /**
   * Returns a list of valid Patient registration locations for sales.
   * 
   *   Permissions Required:
   *   -
   *
   * GET GetPatientRegistrationsLocations V2
   */
  public async salesGetPatientRegistrationsLocationsV2(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetPatientRegistrationsLocationsV2(body, No));
  }

  /**
   * Permissions Required:
   *   - Sales Delivery
   *
   * GET GetPaymenttypes V1
   */
  public async salesGetPaymenttypesV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetPaymenttypesV1(body, licenseNumber));
  }

  /**
   * Returns a list of available payment types for the specified License Number.
   * 
   *   Permissions Required:
   *   - View Sales Delivery
   *   - Manage Sales Delivery
   *
   * GET GetPaymenttypes V2
   */
  public async salesGetPaymenttypesV2(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetPaymenttypesV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * GET GetReceipt V1
   */
  public async salesGetReceiptV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptV1(id, body, licenseNumber));
  }

  /**
   * Retrieves a sales receipt by its Id, with an optional License Number.
   * 
   *   Permissions Required:
   *   - View Sales
   *   - Manage Sales
   *
   * GET GetReceipt V2
   */
  public async salesGetReceiptV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * GET GetReceiptsActive V1
   */
  public async salesGetReceiptsActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptsActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart));
  }

  /**
   * Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
   * 
   *   Permissions Required:
   *   - View Sales
   *   - Manage Sales
   *
   * GET GetReceiptsActive V2
   */
  public async salesGetReceiptsActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptsActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart));
  }

  /**
   * Retrieves a Sales Receipt by its external number, with an optional License Number.
   * 
   *   Permissions Required:
   *   - View Sales
   *   - Manage Sales
   *
   * GET GetReceiptsExternalByExternalNumber V2
   */
  public async salesGetReceiptsExternalByExternalNumberV2(externalNumber: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptsExternalByExternalNumberV2(externalNumber, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * GET GetReceiptsInactive V1
   */
  public async salesGetReceiptsInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptsInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber, salesDateEnd, salesDateStart));
  }

  /**
   * Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
   * 
   *   Permissions Required:
   *   - View Sales
   *   - Manage Sales
   *
   * GET GetReceiptsInactive V2
   */
  public async salesGetReceiptsInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetReceiptsInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize, salesDateEnd, salesDateStart));
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * GET GetTransactions V1
   */
  public async salesGetTransactionsV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetTransactionsV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * GET GetTransactionsBySalesDateStartAndSalesDateEnd V1
   */
  public async salesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesDateStart: string, salesDateEnd: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.salesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesDateStart, salesDateEnd, body, licenseNumber));
  }

  /**
   * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Sales Delivery
   *
   * PUT UpdateDelivery V1
   */
  public async salesUpdateDeliveryV1(body: SalesUpdateDeliveryV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryV1(body, licenseNumber));
  }

  /**
   * Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Manage Sales Delivery
   *
   * PUT UpdateDelivery V2
   */
  public async salesUpdateDeliveryV2(body: SalesUpdateDeliveryV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Sales Delivery
   *
   * PUT UpdateDeliveryComplete V1
   */
  public async salesUpdateDeliveryCompleteV1(body: SalesUpdateDeliveryCompleteV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryCompleteV1(body, licenseNumber));
  }

  /**
   * Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
   * 
   *   Permissions Required:
   *   - Manage Sales Delivery
   *
   * PUT UpdateDeliveryComplete V2
   */
  public async salesUpdateDeliveryCompleteV2(body: SalesUpdateDeliveryCompleteV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryCompleteV2(body, licenseNumber));
  }

  /**
   * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Sales Delivery
   *
   * PUT UpdateDeliveryHub V1
   */
  public async salesUpdateDeliveryHubV1(body: SalesUpdateDeliveryHubV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubV1(body, licenseNumber));
  }

  /**
   * Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Manage Sales Delivery, Manage Sales Delivery Hub
   *
   * PUT UpdateDeliveryHub V2
   */
  public async salesUpdateDeliveryHubV2(body: SalesUpdateDeliveryHubV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * PUT UpdateDeliveryHubAccept V1
   */
  public async salesUpdateDeliveryHubAcceptV1(body: SalesUpdateDeliveryHubAcceptV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubAcceptV1(body, licenseNumber));
  }

  /**
   * Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
   * 
   *   Permissions Required:
   *   - Manage Sales Delivery Hub
   *
   * PUT UpdateDeliveryHubAccept V2
   */
  public async salesUpdateDeliveryHubAcceptV2(body: SalesUpdateDeliveryHubAcceptV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubAcceptV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * PUT UpdateDeliveryHubDepart V1
   */
  public async salesUpdateDeliveryHubDepartV1(body: SalesUpdateDeliveryHubDepartV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubDepartV1(body, licenseNumber));
  }

  /**
   * Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
   * 
   *   Permissions Required:
   *   - Manage Sales Delivery Hub
   *
   * PUT UpdateDeliveryHubDepart V2
   */
  public async salesUpdateDeliveryHubDepartV2(body: SalesUpdateDeliveryHubDepartV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubDepartV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * PUT UpdateDeliveryHubVerifyID V1
   */
  public async salesUpdateDeliveryHubVerifyIdV1(body: SalesUpdateDeliveryHubVerifyIdV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubVerifyIdV1(body, licenseNumber));
  }

  /**
   * Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
   * 
   *   Permissions Required:
   *   - Manage Sales Delivery Hub
   *
   * PUT UpdateDeliveryHubVerifyID V2
   */
  public async salesUpdateDeliveryHubVerifyIdV2(body: SalesUpdateDeliveryHubVerifyIdV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryHubVerifyIdV2(body, licenseNumber));
  }

  /**
   * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Retailer Delivery
   *
   * PUT UpdateDeliveryRetailer V1
   */
  public async salesUpdateDeliveryRetailerV1(body: SalesUpdateDeliveryRetailerV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryRetailerV1(body, licenseNumber));
  }

  /**
   * Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   *   - Industry/Facility Type/Retailer Delivery
   *   - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   *   - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   *   - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
   *   - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
   *   - Manage Retailer Delivery
   *
   * PUT UpdateDeliveryRetailer V2
   */
  public async salesUpdateDeliveryRetailerV2(body: SalesUpdateDeliveryRetailerV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateDeliveryRetailerV2(body, licenseNumber));
  }

  /**
   * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Sales
   *
   * PUT UpdateReceipt V1
   */
  public async salesUpdateReceiptV1(body: SalesUpdateReceiptV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateReceiptV1(body, licenseNumber));
  }

  /**
   * Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Manage Sales
   *
   * PUT UpdateReceipt V2
   */
  public async salesUpdateReceiptV2(body: SalesUpdateReceiptV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateReceiptV2(body, licenseNumber));
  }

  /**
   * Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
   * 
   *   Permissions Required:
   *   - Manage Sales
   *
   * PUT UpdateReceiptFinalize V2
   */
  public async salesUpdateReceiptFinalizeV2(body: SalesUpdateReceiptFinalizeV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateReceiptFinalizeV2(body, licenseNumber));
  }

  /**
   * Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
   * 
   *   Permissions Required:
   *   - Manage Sales
   *
   * PUT UpdateReceiptUnfinalize V2
   */
  public async salesUpdateReceiptUnfinalizeV2(body: SalesUpdateReceiptUnfinalizeV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateReceiptUnfinalizeV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * PUT UpdateTransactionByDate V1
   */
  public async salesUpdateTransactionByDateV1(date: string, body: SalesUpdateTransactionByDateV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.salesUpdateTransactionByDateV1(date, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Strains
   *
   * POST Create V1
   */
  public async strainsCreateV1(body: StrainsCreateV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.strainsCreateV1(body, licenseNumber));
  }

  /**
   * Creates new strain records for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Strains
   *
   * POST Create V2
   */
  public async strainsCreateV2(body: StrainsCreateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.strainsCreateV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Strains
   *
   * POST CreateUpdate V1
   */
  public async strainsCreateUpdateV1(body: StrainsCreateUpdateV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.strainsCreateUpdateV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Strains
   *
   * DELETE Delete V1
   */
  public async strainsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.strainsDeleteV1(id, body, licenseNumber));
  }

  /**
   * Archives an existing strain record for a Facility
   * 
   *   Permissions Required:
   *   - Manage Strains
   *
   * DELETE Delete V2
   */
  public async strainsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.strainsDeleteV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Strains
   *
   * GET Get V1
   */
  public async strainsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.strainsGetV1(id, body, licenseNumber));
  }

  /**
   * Retrieves a Strain record by its Id, with an optional license number.
   * 
   *   Permissions Required:
   *   - Manage Strains
   *
   * GET Get V2
   */
  public async strainsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.strainsGetV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Strains
   *
   * GET GetActive V1
   */
  public async strainsGetActiveV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.strainsGetActiveV1(body, licenseNumber));
  }

  /**
   * Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
   * 
   *   Permissions Required:
   *   - Manage Strains
   *
   * GET GetActive V2
   */
  public async strainsGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.strainsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
   * 
   *   Permissions Required:
   *   - Manage Strains
   *
   * GET GetInactive V2
   */
  public async strainsGetInactiveV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.strainsGetInactiveV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Updates existing strain records for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Strains
   *
   * PUT Update V2
   */
  public async strainsUpdateV2(body: StrainsUpdateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.strainsUpdateV2(body, licenseNumber));
  }

  /**
   * Returns a list of available package tags. NOTE: This is a premium endpoint.
   * 
   *   Permissions Required:
   *   - Manage Tags
   *
   * GET GetPackageAvailable V2
   */
  public async tagsGetPackageAvailableV2(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.tagsGetPackageAvailableV2(body, licenseNumber));
  }

  /**
   * Returns a list of available plant tags. NOTE: This is a premium endpoint.
   * 
   *   Permissions Required:
   *   - Manage Tags
   *
   * GET GetPlantAvailable V2
   */
  public async tagsGetPlantAvailableV2(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.tagsGetPlantAvailableV2(body, licenseNumber));
  }

  /**
   * Returns a list of staged tags. NOTE: This is a premium endpoint.
   * 
   *   Permissions Required:
   *   - Manage Tags
   *   - RetailId.AllowPackageStaging Key Value enabled
   *
   * GET GetStaged V2
   */
  public async tagsGetStagedV2(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.tagsGetStagedV2(body, licenseNumber));
  }

  /**
   * Creates new additive templates for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Additives
   *
   * POST Create V2
   */
  public async additivesTemplatesCreateV2(body: AdditivestemplatesCreateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.additivesTemplatesCreateV2(body, licenseNumber));
  }

  /**
   * Retrieves an Additive Template by its Id.
   * 
   *   Permissions Required:
   *   - Manage Additives
   *
   * GET Get V2
   */
  public async additivesTemplatesGetV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.additivesTemplatesGetV2(id, body, licenseNumber));
  }

  /**
   * Retrieves a list of active additive templates for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Additives
   *
   * GET GetActive V2
   */
  public async additivesTemplatesGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.additivesTemplatesGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of inactive additive templates for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Additives
   *
   * GET GetInactive V2
   */
  public async additivesTemplatesGetInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.additivesTemplatesGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Updates existing additive templates for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Additives
   *
   * PUT Update V2
   */
  public async additivesTemplatesUpdateV2(body: AdditivestemplatesUpdateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.additivesTemplatesUpdateV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *   - Finish/Discontinue Harvests
   *
   * POST CreateFinish V1
   */
  public async harvestsCreateFinishV1(body: HarvestsCreateFinishV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsCreateFinishV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *   - Manage Harvests
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreatePackage V1
   */
  public async harvestsCreatePackageV1(body: HarvestsCreatePackageV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsCreatePackageV1(body, licenseNumber));
  }

  /**
   * Creates packages from harvested products for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Harvests
   *   - Manage Harvests
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreatePackage V2
   */
  public async harvestsCreatePackageV2(body: HarvestsCreatePackageV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsCreatePackageV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *   - Manage Harvests
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreatePackageTesting V1
   */
  public async harvestsCreatePackageTestingV1(body: HarvestsCreatePackageTestingV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsCreatePackageTestingV1(body, licenseNumber));
  }

  /**
   * Creates packages for testing from harvested products for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Harvests
   *   - Manage Harvests
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreatePackageTesting V2
   */
  public async harvestsCreatePackageTestingV2(body: HarvestsCreatePackageTestingV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsCreatePackageTestingV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *   - Manage Harvests
   *
   * POST CreateRemoveWaste V1
   */
  public async harvestsCreateRemoveWasteV1(body: HarvestsCreateRemoveWasteV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsCreateRemoveWasteV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *   - Finish/Discontinue Harvests
   *
   * POST CreateUnfinish V1
   */
  public async harvestsCreateUnfinishV1(body: HarvestsCreateUnfinishV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsCreateUnfinishV1(body, licenseNumber));
  }

  /**
   * Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
   * 
   *   Permissions Required:
   *   - View Harvests
   *   - Manage Harvests
   *
   * POST CreateWaste V2
   */
  public async harvestsCreateWasteV2(body: HarvestsCreateWasteV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsCreateWasteV2(body, licenseNumber));
  }

  /**
   * Discontinues a specific harvest waste record by Id for the specified Facility.
   * 
   *   Permissions Required:
   *   - View Harvests
   *   - Discontinue Harvest Waste
   *
   * DELETE DeleteWaste V2
   */
  public async harvestsDeleteWasteV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsDeleteWasteV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *
   * GET Get V1
   */
  public async harvestsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.harvestsGetV1(id, body, licenseNumber));
  }

  /**
   * Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
   * 
   *   Permissions Required:
   *   - View Harvests
   *
   * GET Get V2
   */
  public async harvestsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.harvestsGetV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *
   * GET GetActive V1
   */
  public async harvestsGetActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.harvestsGetActiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves a list of active harvests for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Harvests
   *
   * GET GetActive V2
   */
  public async harvestsGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.harvestsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *
   * GET GetInactive V1
   */
  public async harvestsGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.harvestsGetInactiveV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves a list of inactive harvests for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Harvests
   *
   * GET GetInactive V2
   */
  public async harvestsGetInactiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.harvestsGetInactiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *
   * GET GetOnhold V1
   */
  public async harvestsGetOnholdV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.harvestsGetOnholdV1(body, lastModifiedEnd, lastModifiedStart, licenseNumber));
  }

  /**
   * Retrieves a list of harvests on hold for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Harvests
   *
   * GET GetOnhold V2
   */
  public async harvestsGetOnholdV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.harvestsGetOnholdV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
   * 
   *   Permissions Required:
   *   - View Harvests
   *
   * GET GetWaste V2
   */
  public async harvestsGetWasteV2(body?: any, harvestId?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.harvestsGetWasteV2(body, harvestId, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetWasteTypes V1
   */
  public async harvestsGetWasteTypesV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.harvestsGetWasteTypesV1(body, No));
  }

  /**
   * Retrieves a list of Waste types for harvests.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetWasteTypes V2
   */
  public async harvestsGetWasteTypesV2(body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.harvestsGetWasteTypesV2(body, pageNumber, pageSize));
  }

  /**
   * Marks one or more harvests as finished for the specified Facility.
   * 
   *   Permissions Required:
   *   - View Harvests
   *   - Finish/Discontinue Harvests
   *
   * PUT UpdateFinish V2
   */
  public async harvestsUpdateFinishV2(body: HarvestsUpdateFinishV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateFinishV2(body, licenseNumber));
  }

  /**
   * Updates the Location of Harvest for a specified Facility.
   * 
   *   Permissions Required:
   *   - View Harvests
   *   - Manage Harvests
   *
   * PUT UpdateLocation V2
   */
  public async harvestsUpdateLocationV2(body: HarvestsUpdateLocationV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateLocationV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *   - Manage Harvests
   *
   * PUT UpdateMove V1
   */
  public async harvestsUpdateMoveV1(body: HarvestsUpdateMoveV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateMoveV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *   - Manage Harvests
   *
   * PUT UpdateRename V1
   */
  public async harvestsUpdateRenameV1(body: HarvestsUpdateRenameV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateRenameV1(body, licenseNumber));
  }

  /**
   * Renames one or more harvests for the specified Facility.
   * 
   *   Permissions Required:
   *   - View Harvests
   *   - Manage Harvests
   *
   * PUT UpdateRename V2
   */
  public async harvestsUpdateRenameV2(body: HarvestsUpdateRenameV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateRenameV2(body, licenseNumber));
  }

  /**
   * Restores previously harvested plants to their original state for the specified Facility.
   * 
   *   Permissions Required:
   *   - View Harvests
   *   - Finish/Discontinue Harvests
   *
   * PUT UpdateRestoreHarvestedPlants V2
   */
  public async harvestsUpdateRestoreHarvestedPlantsV2(body: HarvestsUpdateRestoreHarvestedPlantsV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateRestoreHarvestedPlantsV2(body, licenseNumber));
  }

  /**
   * Reopens one or more previously finished harvests for the specified Facility.
   * 
   *   Permissions Required:
   *   - View Harvests
   *   - Finish/Discontinue Harvests
   *
   * PUT UpdateUnfinish V2
   */
  public async harvestsUpdateUnfinishV2(body: HarvestsUpdateUnfinishV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.harvestsUpdateUnfinishV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreateRecord V1
   */
  public async labTestsCreateRecordV1(body: LabtestsCreateRecordV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.labTestsCreateRecordV1(body, licenseNumber));
  }

  /**
   * Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreateRecord V2
   */
  public async labTestsCreateRecordV2(body: LabtestsCreateRecordV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.labTestsCreateRecordV2(body, licenseNumber));
  }

  /**
   * Retrieves a list of Lab Test batches.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetBatches V2
   */
  public async labTestsGetBatchesV2(body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.labTestsGetBatchesV2(body, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * GET GetLabtestdocument V1
   */
  public async labTestsGetLabtestdocumentV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.labTestsGetLabtestdocumentV1(id, body, licenseNumber));
  }

  /**
   * Retrieves a specific Lab Test result document by its Id for a given Facility.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * GET GetLabtestdocument V2
   */
  public async labTestsGetLabtestdocumentV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.labTestsGetLabtestdocumentV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *
   * GET GetResults V1
   */
  public async labTestsGetResultsV1(body?: any, licenseNumber?: string, packageId?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.labTestsGetResultsV1(body, licenseNumber, packageId));
  }

  /**
   * Retrieves Lab Test results for a specified Package.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * GET GetResults V2
   */
  public async labTestsGetResultsV2(body?: any, licenseNumber?: string, packageId?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.labTestsGetResultsV2(body, licenseNumber, packageId, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetStates V1
   */
  public async labTestsGetStatesV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.labTestsGetStatesV1(body, No));
  }

  /**
   * Returns a list of all lab testing states.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetStates V2
   */
  public async labTestsGetStatesV2(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.labTestsGetStatesV2(body, No));
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetTypes V1
   */
  public async labTestsGetTypesV1(body?: any, No?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.labTestsGetTypesV1(body, No));
  }

  /**
   * Returns a list of Lab Test types.
   * 
   *   Permissions Required:
   *   - None
   *
   * GET GetTypes V2
   */
  public async labTestsGetTypesV2(body?: any, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.labTestsGetTypesV2(body, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateLabtestdocument V1
   */
  public async labTestsUpdateLabtestdocumentV1(body: LabtestsUpdateLabtestdocumentV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.labTestsUpdateLabtestdocumentV1(body, licenseNumber));
  }

  /**
   * Updates one or more documents for previously submitted lab tests.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateLabtestdocument V2
   */
  public async labTestsUpdateLabtestdocumentV2(body: LabtestsUpdateLabtestdocumentV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.labTestsUpdateLabtestdocumentV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateResultRelease V1
   */
  public async labTestsUpdateResultReleaseV1(body: LabtestsUpdateResultReleaseV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.labTestsUpdateResultReleaseV1(body, licenseNumber));
  }

  /**
   * Releases Lab Test results for one or more packages.
   * 
   *   Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateResultRelease V2
   */
  public async labTestsUpdateResultReleaseV2(body: LabtestsUpdateResultReleaseV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.labTestsUpdateResultReleaseV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Locations
   *
   * POST Create V1
   */
  public async locationsCreateV1(body: LocationsCreateV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.locationsCreateV1(body, licenseNumber));
  }

  /**
   * Creates new locations for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * POST Create V2
   */
  public async locationsCreateV2(body: LocationsCreateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.locationsCreateV2(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Locations
   *
   * POST CreateUpdate V1
   */
  public async locationsCreateUpdateV1(body: LocationsCreateUpdateV1RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.locationsCreateUpdateV1(body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Locations
   *
   * DELETE Delete V1
   */
  public async locationsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.locationsDeleteV1(id, body, licenseNumber));
  }

  /**
   * Archives a specified Location, identified by its Id, for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * DELETE Delete V2
   */
  public async locationsDeleteV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.locationsDeleteV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Locations
   *
   * GET Get V1
   */
  public async locationsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.locationsGetV1(id, body, licenseNumber));
  }

  /**
   * Retrieves a Location by its Id.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * GET Get V2
   */
  public async locationsGetV2(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.locationsGetV2(id, body, licenseNumber));
  }

  /**
   * Permissions Required:
   *   - Manage Locations
   *
   * GET GetActive V1
   */
  public async locationsGetActiveV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.locationsGetActiveV1(body, licenseNumber));
  }

  /**
   * Retrieves a list of active locations for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * GET GetActive V2
   */
  public async locationsGetActiveV2(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.locationsGetActiveV2(body, lastModifiedEnd, lastModifiedStart, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Retrieves a list of inactive locations for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * GET GetInactive V2
   */
  public async locationsGetInactiveV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.locationsGetInactiveV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Permissions Required:
   *   - Manage Locations
   *
   * GET GetTypes V1
   */
  public async locationsGetTypesV1(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.locationsGetTypesV1(body, licenseNumber));
  }

  /**
   * Retrieves a list of active location types for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * GET GetTypes V2
   */
  public async locationsGetTypesV2(body?: any, licenseNumber?: string, pageNumber?: string, pageSize?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.locationsGetTypesV2(body, licenseNumber, pageNumber, pageSize));
  }

  /**
   * Updates existing locations for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * PUT Update V2
   */
  public async locationsUpdateV2(body: LocationsUpdateV2RequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, false, () => this.client.locationsUpdateV2(body, licenseNumber));
  }

  /**
   * Data returned by this endpoint is cached for up to one minute.
   * 
   *   Permissions Required:
   *   - Lookup Patients
   *
   * GET GetStatusesByPatientLicenseNumber V1
   */
  public async patientsStatusGetStatusesByPatientLicenseNumberV1(patientLicenseNumber: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.patientsStatusGetStatusesByPatientLicenseNumberV1(patientLicenseNumber, body, licenseNumber));
  }

  /**
   * Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
   * 
   *   Permissions Required:
   *   - Lookup Patients
   *
   * GET GetStatusesByPatientLicenseNumber V2
   */
  public async patientsStatusGetStatusesByPatientLicenseNumberV2(patientLicenseNumber: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null, true, () => this.client.patientsStatusGetStatusesByPatientLicenseNumberV2(patientLicenseNumber, body, licenseNumber));
  }

}

// --- Request Models ---
export interface PlantsCreateAdditivesV1RequestItem {
  'ActiveIngredients'?: PlantsCreateAdditivesV1RequestItem_ActiveIngredients[];
  'ActualDate'?: string;
  'AdditiveType'?: string;
  'ApplicationDevice'?: string;
  'EpaRegistrationNumber'?: string;
  'PlantLabels'?: string[];
  'ProductSupplier'?: string;
  'ProductTradeName'?: string;
  'TotalAmountApplied'?: number;
  'TotalAmountUnitOfMeasure'?: string;
}

export interface PlantsCreateAdditivesV1RequestItem_ActiveIngredients {
  'Name'?: string;
  'Percentage'?: number;
}

export interface PlantsCreateAdditivesV2RequestItem {
  'ActiveIngredients'?: PlantsCreateAdditivesV2RequestItem_ActiveIngredients[];
  'ActualDate'?: string;
  'AdditiveType'?: string;
  'ApplicationDevice'?: string;
  'EpaRegistrationNumber'?: string;
  'PlantLabels'?: string[];
  'ProductSupplier'?: string;
  'ProductTradeName'?: string;
  'TotalAmountApplied'?: number;
  'TotalAmountUnitOfMeasure'?: string;
}

export interface PlantsCreateAdditivesV2RequestItem_ActiveIngredients {
  'Name'?: string;
  'Percentage'?: number;
}

export interface PlantsCreateAdditivesBylocationV1RequestItem {
  'ActiveIngredients'?: PlantsCreateAdditivesBylocationV1RequestItem_ActiveIngredients[];
  'ActualDate'?: string;
  'AdditiveType'?: string;
  'ApplicationDevice'?: string;
  'EpaRegistrationNumber'?: string;
  'LocationName'?: string;
  'ProductSupplier'?: string;
  'ProductTradeName'?: string;
  'SublocationName'?: string;
  'TotalAmountApplied'?: number;
  'TotalAmountUnitOfMeasure'?: string;
}

export interface PlantsCreateAdditivesBylocationV1RequestItem_ActiveIngredients {
  'Name'?: string;
  'Percentage'?: number;
}

export interface PlantsCreateAdditivesBylocationV2RequestItem {
  'ActiveIngredients'?: PlantsCreateAdditivesBylocationV2RequestItem_ActiveIngredients[];
  'ActualDate'?: string;
  'AdditiveType'?: string;
  'ApplicationDevice'?: string;
  'EpaRegistrationNumber'?: string;
  'LocationName'?: string;
  'ProductSupplier'?: string;
  'ProductTradeName'?: string;
  'SublocationName'?: string;
  'TotalAmountApplied'?: number;
  'TotalAmountUnitOfMeasure'?: string;
}

export interface PlantsCreateAdditivesBylocationV2RequestItem_ActiveIngredients {
  'Name'?: string;
  'Percentage'?: number;
}

export interface PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem {
  'ActualDate'?: string;
  'AdditivesTemplateName'?: string;
  'LocationName'?: string;
  'Rate'?: string;
  'SublocationName'?: string;
  'TotalAmountApplied'?: number;
  'TotalAmountUnitOfMeasure'?: string;
  'Volume'?: string;
}

export interface PlantsCreateAdditivesUsingtemplateV2RequestItem {
  'ActualDate'?: string;
  'AdditivesTemplateName'?: string;
  'PlantLabels'?: string[];
  'Rate'?: string;
  'TotalAmountApplied'?: number;
  'TotalAmountUnitOfMeasure'?: string;
  'Volume'?: string;
}

export interface PlantsCreateChangegrowthphasesV1RequestItem {
  'GrowthDate'?: string;
  'GrowthPhase'?: string;
  'Id'?: number;
  'Label'?: string;
  'NewLocation'?: string;
  'NewSublocation'?: string;
  'NewTag'?: string;
}

export interface PlantsCreateHarvestplantsV1RequestItem {
  'ActualDate'?: string;
  'DryingLocation'?: string;
  'DryingSublocation'?: string;
  'HarvestName'?: string;
  'PatientLicenseNumber'?: string;
  'Plant'?: string;
  'UnitOfWeight'?: string;
  'Weight'?: number;
}

export interface PlantsCreateManicureV2RequestItem {
  'ActualDate'?: string;
  'DryingLocation'?: string;
  'DryingSublocation'?: string;
  'HarvestName'?: string;
  'PatientLicenseNumber'?: string;
  'Plant'?: string;
  'PlantCount'?: number;
  'UnitOfWeight'?: string;
  'Weight'?: number;
}

export interface PlantsCreateManicureplantsV1RequestItem {
  'ActualDate'?: string;
  'DryingLocation'?: string;
  'DryingSublocation'?: string;
  'HarvestName'?: string;
  'PatientLicenseNumber'?: string;
  'Plant'?: string;
  'PlantCount'?: number;
  'UnitOfWeight'?: string;
  'Weight'?: number;
}

export interface PlantsCreateMoveplantsV1RequestItem {
  'ActualDate'?: string;
  'Id'?: number;
  'Label'?: string;
  'Location'?: string;
  'Sublocation'?: string;
}

export interface PlantsCreatePlantbatchPackageV1RequestItem {
  'ActualDate'?: string;
  'Count'?: number;
  'IsDonation'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'Location'?: string;
  'Note'?: string;
  'PackageTag'?: string;
  'PatientLicenseNumber'?: string;
  'PlantBatchType'?: string;
  'PlantLabel'?: string;
  'Sublocation'?: string;
}

export interface PlantsCreatePlantbatchPackageV2RequestItem {
  'ActualDate'?: string;
  'Count'?: number;
  'IsDonation'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'Location'?: string;
  'Note'?: string;
  'PackageTag'?: string;
  'PatientLicenseNumber'?: string;
  'PlantBatchType'?: string;
  'PlantLabel'?: string;
  'Sublocation'?: string;
}

export interface PlantsCreatePlantingsV1RequestItem {
  'ActualDate'?: string;
  'LocationName'?: string;
  'PatientLicenseNumber'?: string;
  'PlantBatchName'?: string;
  'PlantBatchType'?: string;
  'PlantCount'?: number;
  'PlantLabel'?: string;
  'StrainName'?: string;
  'SublocationName'?: string;
}

export interface PlantsCreatePlantingsV2RequestItem {
  'ActualDate'?: string;
  'LocationName'?: string;
  'PatientLicenseNumber'?: string;
  'PlantBatchName'?: string;
  'PlantBatchType'?: string;
  'PlantCount'?: number;
  'PlantLabel'?: string;
  'StrainName'?: string;
  'SublocationName'?: string;
}

export interface PlantsCreateWasteV1RequestItem {
  'LocationName'?: string;
  'MixedMaterial'?: string;
  'Note'?: string;
  'PlantLabels'?: any[];
  'ReasonName'?: string;
  'SublocationName'?: string;
  'UnitOfMeasureName'?: string;
  'WasteDate'?: string;
  'WasteMethodName'?: string;
  'WasteWeight'?: number;
}

export interface PlantsCreateWasteV2RequestItem {
  'LocationName'?: string;
  'MixedMaterial'?: string;
  'Note'?: string;
  'PlantLabels'?: any[];
  'ReasonName'?: string;
  'SublocationName'?: string;
  'UnitOfMeasureName'?: string;
  'WasteDate'?: string;
  'WasteMethodName'?: string;
  'WasteWeight'?: number;
}

export interface PlantsUpdateAdjustV2RequestItem {
  'AdjustCount'?: number;
  'AdjustReason'?: string;
  'AdjustmentDate'?: string;
  'Id'?: number;
  'Label'?: string;
  'ReasonNote'?: string;
}

export interface PlantsUpdateGrowthphaseV2RequestItem {
  'GrowthDate'?: string;
  'GrowthPhase'?: string;
  'Id'?: number;
  'Label'?: string;
  'NewLocation'?: string;
  'NewSublocation'?: string;
  'NewTag'?: string;
}

export interface PlantsUpdateHarvestV2RequestItem {
  'ActualDate'?: string;
  'DryingLocation'?: string;
  'DryingSublocation'?: string;
  'HarvestName'?: string;
  'PatientLicenseNumber'?: string;
  'Plant'?: string;
  'UnitOfWeight'?: string;
  'Weight'?: number;
}

export interface PlantsUpdateLocationV2RequestItem {
  'ActualDate'?: string;
  'Id'?: number;
  'Label'?: string;
  'Location'?: string;
  'Sublocation'?: string;
}

export interface PlantsUpdateMergeV2RequestItem {
  'MergeDate'?: string;
  'SourcePlantGroupLabel'?: string;
  'TargetPlantGroupLabel'?: string;
}

export interface PlantsUpdateSplitV2RequestItem {
  'PlantCount'?: number;
  'SourcePlantLabel'?: string;
  'SplitDate'?: string;
  'StrainLabel'?: string;
  'TagLabel'?: string;
}

export interface PlantsUpdateStrainV2RequestItem {
  'Id'?: number;
  'Label'?: string;
  'StrainId'?: number;
  'StrainName'?: string;
}

export interface PlantsUpdateTagV2RequestItem {
  'Id'?: number;
  'Label'?: string;
  'NewTag'?: string;
  'ReplaceDate'?: string;
  'TagId'?: number;
}

export interface RetailidCreateAssociateV2RequestItem {
  'PackageLabel'?: string;
  'QrUrls'?: string[];
}

export interface RetailidCreateGenerateV2Request {
  'PackageLabel'?: string;
  'Quantity'?: number;
}

export interface RetailidCreateMergeV2Request {
  'packageLabels'?: string[];
}

export interface RetailidCreatePackageInfoV2Request {
  'packageLabels'?: string[];
}

export interface PatientsCreateV2RequestItem {
  'ActualDate'?: string;
  'ConcentrateOuncesAllowed'?: number;
  'FlowerOuncesAllowed'?: number;
  'HasSalesLimitExemption'?: boolean;
  'InfusedOuncesAllowed'?: number;
  'LicenseEffectiveEndDate'?: string;
  'LicenseEffectiveStartDate'?: string;
  'LicenseNumber'?: string;
  'MaxConcentrateThcPercentAllowed'?: number;
  'MaxFlowerThcPercentAllowed'?: number;
  'RecommendedPlants'?: number;
  'RecommendedSmokableQuantity'?: number;
  'ThcOuncesAllowed'?: number;
}

export interface PatientsCreateAddV1RequestItem {
  'ActualDate'?: string;
  'ConcentrateOuncesAllowed'?: number;
  'FlowerOuncesAllowed'?: number;
  'HasSalesLimitExemption'?: boolean;
  'InfusedOuncesAllowed'?: number;
  'LicenseEffectiveEndDate'?: string;
  'LicenseEffectiveStartDate'?: string;
  'LicenseNumber'?: string;
  'MaxConcentrateThcPercentAllowed'?: number;
  'MaxFlowerThcPercentAllowed'?: number;
  'RecommendedPlants'?: number;
  'RecommendedSmokableQuantity'?: number;
  'ThcOuncesAllowed'?: number;
}

export interface PatientsCreateUpdateV1RequestItem {
  'ActualDate'?: string;
  'ConcentrateOuncesAllowed'?: number;
  'FlowerOuncesAllowed'?: number;
  'HasSalesLimitExemption'?: boolean;
  'InfusedOuncesAllowed'?: number;
  'LicenseEffectiveEndDate'?: string;
  'LicenseEffectiveStartDate'?: string;
  'LicenseNumber'?: string;
  'MaxConcentrateThcPercentAllowed'?: number;
  'MaxFlowerThcPercentAllowed'?: number;
  'NewLicenseNumber'?: string;
  'RecommendedPlants'?: number;
  'RecommendedSmokableQuantity'?: number;
  'ThcOuncesAllowed'?: number;
}

export interface PatientsUpdateV2RequestItem {
  'ActualDate'?: string;
  'ConcentrateOuncesAllowed'?: number;
  'FlowerOuncesAllowed'?: number;
  'HasSalesLimitExemption'?: boolean;
  'InfusedOuncesAllowed'?: number;
  'LicenseEffectiveEndDate'?: string;
  'LicenseEffectiveStartDate'?: string;
  'LicenseNumber'?: string;
  'MaxConcentrateThcPercentAllowed'?: number;
  'MaxFlowerThcPercentAllowed'?: number;
  'NewLicenseNumber'?: string;
  'RecommendedPlants'?: number;
  'RecommendedSmokableQuantity'?: number;
  'ThcOuncesAllowed'?: number;
}

export interface TransfersCreateExternalIncomingV1RequestItem {
  'Destinations'?: TransfersCreateExternalIncomingV1RequestItem_Destinations[];
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'PhoneNumberForQuestions'?: string;
  'ShipperAddress1'?: string;
  'ShipperAddress2'?: string;
  'ShipperAddressCity'?: string;
  'ShipperAddressPostalCode'?: string;
  'ShipperAddressState'?: string;
  'ShipperLicenseNumber'?: string;
  'ShipperMainPhoneNumber'?: string;
  'ShipperName'?: string;
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersCreateExternalIncomingV1RequestItem_Destinations {
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'GrossUnitOfWeightId'?: number;
  'GrossWeight'?: number;
  'InvoiceNumber'?: string;
  'Packages'?: TransfersCreateExternalIncomingV1RequestItem_Destinations_Packages[];
  'PlannedRoute'?: string;
  'RecipientLicenseNumber'?: string;
  'TransferTypeName'?: string;
  'Transporters'?: TransfersCreateExternalIncomingV1RequestItem_Destinations_Transporters[];
}

export interface TransfersCreateExternalIncomingV1RequestItem_Destinations_Packages {
  'ExpirationDate'?: string;
  'ExternalId'?: string;
  'GrossUnitOfWeightName'?: string;
  'GrossWeight'?: number;
  'ItemName'?: string;
  'PackagedDate'?: string;
  'Quantity'?: number;
  'SellByDate'?: string;
  'UnitOfMeasureName'?: string;
  'UseByDate'?: string;
  'WholesalePrice'?: string;
}

export interface TransfersCreateExternalIncomingV1RequestItem_Destinations_Transporters {
  'DriverLayoverLeg'?: string;
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'IsLayover'?: boolean;
  'PhoneNumberForQuestions'?: string;
  'TransporterDetails'?: string;
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersCreateExternalIncomingV2RequestItem {
  'Destinations'?: TransfersCreateExternalIncomingV2RequestItem_Destinations[];
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'PhoneNumberForQuestions'?: string;
  'ShipperAddress1'?: string;
  'ShipperAddress2'?: string;
  'ShipperAddressCity'?: string;
  'ShipperAddressPostalCode'?: string;
  'ShipperAddressState'?: string;
  'ShipperLicenseNumber'?: string;
  'ShipperMainPhoneNumber'?: string;
  'ShipperName'?: string;
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersCreateExternalIncomingV2RequestItem_Destinations {
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'GrossUnitOfWeightId'?: number;
  'GrossWeight'?: number;
  'InvoiceNumber'?: string;
  'Packages'?: TransfersCreateExternalIncomingV2RequestItem_Destinations_Packages[];
  'PlannedRoute'?: string;
  'RecipientLicenseNumber'?: string;
  'TransferTypeName'?: string;
  'Transporters'?: TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters[];
}

export interface TransfersCreateExternalIncomingV2RequestItem_Destinations_Packages {
  'ExpirationDate'?: string;
  'ExternalId'?: string;
  'GrossUnitOfWeightName'?: string;
  'GrossWeight'?: number;
  'ItemName'?: string;
  'PackagedDate'?: string;
  'Quantity'?: number;
  'SellByDate'?: string;
  'UnitOfMeasureName'?: string;
  'UseByDate'?: string;
  'WholesalePrice'?: string;
}

export interface TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters {
  'DriverLayoverLeg'?: string;
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'IsLayover'?: boolean;
  'PhoneNumberForQuestions'?: string;
  'TransporterDetails'?: TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails[];
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails {
  'DriverLayoverLeg'?: string;
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersCreateTemplatesV1RequestItem {
  'Destinations'?: TransfersCreateTemplatesV1RequestItem_Destinations[];
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'Name'?: string;
  'PhoneNumberForQuestions'?: string;
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersCreateTemplatesV1RequestItem_Destinations {
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'InvoiceNumber'?: string;
  'Packages'?: TransfersCreateTemplatesV1RequestItem_Destinations_Packages[];
  'PlannedRoute'?: string;
  'RecipientLicenseNumber'?: string;
  'TransferTypeName'?: string;
  'Transporters'?: TransfersCreateTemplatesV1RequestItem_Destinations_Transporters[];
}

export interface TransfersCreateTemplatesV1RequestItem_Destinations_Packages {
  'GrossUnitOfWeightName'?: string;
  'GrossWeight'?: number;
  'PackageLabel'?: string;
  'WholesalePrice'?: string;
}

export interface TransfersCreateTemplatesV1RequestItem_Destinations_Transporters {
  'DriverLayoverLeg'?: string;
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'IsLayover'?: boolean;
  'PhoneNumberForQuestions'?: string;
  'TransporterDetails'?: string;
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersCreateTemplatesOutgoingV2RequestItem {
  'Destinations'?: TransfersCreateTemplatesOutgoingV2RequestItem_Destinations[];
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'Name'?: string;
  'PhoneNumberForQuestions'?: string;
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersCreateTemplatesOutgoingV2RequestItem_Destinations {
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'InvoiceNumber'?: string;
  'Packages'?: TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Packages[];
  'PlannedRoute'?: string;
  'RecipientLicenseNumber'?: string;
  'TransferTypeName'?: string;
  'Transporters'?: TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters[];
}

export interface TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Packages {
  'GrossUnitOfWeightName'?: string;
  'GrossWeight'?: number;
  'PackageLabel'?: string;
  'WholesalePrice'?: string;
}

export interface TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters {
  'DriverLayoverLeg'?: string;
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'IsLayover'?: boolean;
  'PhoneNumberForQuestions'?: string;
  'TransporterDetails'?: TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails[];
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails {
  'DriverLayoverLeg'?: string;
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersUpdateExternalIncomingV1RequestItem {
  'Destinations'?: TransfersUpdateExternalIncomingV1RequestItem_Destinations[];
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'PhoneNumberForQuestions'?: string;
  'ShipperAddress1'?: string;
  'ShipperAddress2'?: string;
  'ShipperAddressCity'?: string;
  'ShipperAddressPostalCode'?: string;
  'ShipperAddressState'?: string;
  'ShipperLicenseNumber'?: string;
  'ShipperMainPhoneNumber'?: string;
  'ShipperName'?: string;
  'TransferId'?: number;
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersUpdateExternalIncomingV1RequestItem_Destinations {
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'GrossUnitOfWeightId'?: number;
  'GrossWeight'?: number;
  'InvoiceNumber'?: string;
  'Packages'?: TransfersUpdateExternalIncomingV1RequestItem_Destinations_Packages[];
  'PlannedRoute'?: string;
  'RecipientLicenseNumber'?: string;
  'TransferDestinationId'?: number;
  'TransferTypeName'?: string;
  'Transporters'?: TransfersUpdateExternalIncomingV1RequestItem_Destinations_Transporters[];
}

export interface TransfersUpdateExternalIncomingV1RequestItem_Destinations_Packages {
  'ExpirationDate'?: string;
  'ExternalId'?: string;
  'GrossUnitOfWeightName'?: string;
  'GrossWeight'?: number;
  'ItemName'?: string;
  'PackagedDate'?: string;
  'Quantity'?: number;
  'SellByDate'?: string;
  'UnitOfMeasureName'?: string;
  'UseByDate'?: string;
  'WholesalePrice'?: string;
}

export interface TransfersUpdateExternalIncomingV1RequestItem_Destinations_Transporters {
  'DriverLayoverLeg'?: string;
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'IsLayover'?: boolean;
  'PhoneNumberForQuestions'?: string;
  'TransporterDetails'?: string;
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersUpdateExternalIncomingV2RequestItem {
  'Destinations'?: TransfersUpdateExternalIncomingV2RequestItem_Destinations[];
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'PhoneNumberForQuestions'?: string;
  'ShipperAddress1'?: string;
  'ShipperAddress2'?: string;
  'ShipperAddressCity'?: string;
  'ShipperAddressPostalCode'?: string;
  'ShipperAddressState'?: string;
  'ShipperLicenseNumber'?: string;
  'ShipperMainPhoneNumber'?: string;
  'ShipperName'?: string;
  'TransferId'?: number;
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersUpdateExternalIncomingV2RequestItem_Destinations {
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'GrossUnitOfWeightId'?: number;
  'GrossWeight'?: number;
  'InvoiceNumber'?: string;
  'Packages'?: TransfersUpdateExternalIncomingV2RequestItem_Destinations_Packages[];
  'PlannedRoute'?: string;
  'RecipientLicenseNumber'?: string;
  'TransferDestinationId'?: number;
  'TransferTypeName'?: string;
  'Transporters'?: TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters[];
}

export interface TransfersUpdateExternalIncomingV2RequestItem_Destinations_Packages {
  'ExpirationDate'?: string;
  'ExternalId'?: string;
  'GrossUnitOfWeightName'?: string;
  'GrossWeight'?: number;
  'ItemName'?: string;
  'PackagedDate'?: string;
  'Quantity'?: number;
  'SellByDate'?: string;
  'UnitOfMeasureName'?: string;
  'UseByDate'?: string;
  'WholesalePrice'?: string;
}

export interface TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters {
  'DriverLayoverLeg'?: string;
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'IsLayover'?: boolean;
  'PhoneNumberForQuestions'?: string;
  'TransporterDetails'?: TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails[];
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails {
  'DriverLayoverLeg'?: string;
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersUpdateTemplatesV1RequestItem {
  'Destinations'?: TransfersUpdateTemplatesV1RequestItem_Destinations[];
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'Name'?: string;
  'PhoneNumberForQuestions'?: string;
  'TransferTemplateId'?: number;
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersUpdateTemplatesV1RequestItem_Destinations {
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'InvoiceNumber'?: string;
  'Packages'?: TransfersUpdateTemplatesV1RequestItem_Destinations_Packages[];
  'PlannedRoute'?: string;
  'RecipientLicenseNumber'?: string;
  'TransferDestinationId'?: number;
  'TransferTypeName'?: string;
  'Transporters'?: TransfersUpdateTemplatesV1RequestItem_Destinations_Transporters[];
}

export interface TransfersUpdateTemplatesV1RequestItem_Destinations_Packages {
  'GrossUnitOfWeightName'?: string;
  'GrossWeight'?: number;
  'PackageLabel'?: string;
  'WholesalePrice'?: string;
}

export interface TransfersUpdateTemplatesV1RequestItem_Destinations_Transporters {
  'DriverLayoverLeg'?: string;
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'IsLayover'?: boolean;
  'PhoneNumberForQuestions'?: string;
  'TransporterDetails'?: string;
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersUpdateTemplatesOutgoingV2RequestItem {
  'Destinations'?: TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations[];
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'Name'?: string;
  'PhoneNumberForQuestions'?: string;
  'TransferTemplateId'?: number;
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations {
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'InvoiceNumber'?: string;
  'Packages'?: TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Packages[];
  'PlannedRoute'?: string;
  'RecipientLicenseNumber'?: string;
  'TransferDestinationId'?: number;
  'TransferTypeName'?: string;
  'Transporters'?: TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters[];
}

export interface TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Packages {
  'GrossUnitOfWeightName'?: string;
  'GrossWeight'?: number;
  'PackageLabel'?: string;
  'WholesalePrice'?: string;
}

export interface TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters {
  'DriverLayoverLeg'?: string;
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'IsLayover'?: boolean;
  'PhoneNumberForQuestions'?: string;
  'TransporterDetails'?: TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails[];
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails {
  'DriverLayoverLeg'?: string;
  'DriverLicenseNumber'?: string;
  'DriverName'?: string;
  'DriverOccupationalLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface TransportersCreateDriverV2RequestItem {
  'DriversLicenseNumber'?: string;
  'EmployeeId'?: string;
  'Name'?: string;
}

export interface TransportersCreateVehicleV2RequestItem {
  'LicensePlateNumber'?: string;
  'Make'?: string;
  'Model'?: string;
}

export interface TransportersUpdateDriverV2RequestItem {
  'DriversLicenseNumber'?: string;
  'EmployeeId'?: string;
  'Id'?: string;
  'Name'?: string;
}

export interface TransportersUpdateVehicleV2RequestItem {
  'Id'?: string;
  'LicensePlateNumber'?: string;
  'Make'?: string;
  'Model'?: string;
}

export interface ItemsCreateV1RequestItem {
  'AdministrationMethod'?: string;
  'Allergens'?: string;
  'Description'?: string;
  'GlobalProductName'?: string;
  'ItemBrand'?: string;
  'ItemCategory'?: string;
  'ItemIngredients'?: string;
  'LabelImageFileSystemIds'?: string;
  'LabelPhotoDescription'?: string;
  'Name'?: string;
  'NumberOfDoses'?: string;
  'PackagingImageFileSystemIds'?: string;
  'PackagingPhotoDescription'?: string;
  'ProcessingJobCategoryName'?: string;
  'ProcessingJobTypeName'?: string;
  'ProductImageFileSystemIds'?: string;
  'ProductPDFFileSystemIds'?: string;
  'ProductPhotoDescription'?: string;
  'PublicIngredients'?: string;
  'ServingSize'?: string;
  'Strain'?: string;
  'SupplyDurationDays'?: number;
  'UnitCbdAContent'?: number;
  'UnitCbdAContentDose'?: number;
  'UnitCbdAContentDoseUnitOfMeasure'?: string;
  'UnitCbdAContentUnitOfMeasure'?: string;
  'UnitCbdAPercent'?: number;
  'UnitCbdContent'?: number;
  'UnitCbdContentDose'?: number;
  'UnitCbdContentDoseUnitOfMeasure'?: string;
  'UnitCbdContentUnitOfMeasure'?: string;
  'UnitCbdPercent'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcAContent'?: number;
  'UnitThcAContentDose'?: number;
  'UnitThcAContentDoseUnitOfMeasure'?: string;
  'UnitThcAContentUnitOfMeasure'?: string;
  'UnitThcAPercent'?: number;
  'UnitThcContent'?: number;
  'UnitThcContentDose'?: number;
  'UnitThcContentDoseUnitOfMeasure'?: string;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitVolume'?: number;
  'UnitVolumeUnitOfMeasure'?: string;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface ItemsCreateV2RequestItem {
  'AdministrationMethod'?: string;
  'Allergens'?: string;
  'Description'?: string;
  'GlobalProductName'?: string;
  'ItemBrand'?: string;
  'ItemCategory'?: string;
  'ItemIngredients'?: string;
  'LabelImageFileSystemIds'?: string;
  'LabelPhotoDescription'?: string;
  'Name'?: string;
  'NumberOfDoses'?: string;
  'PackagingImageFileSystemIds'?: string;
  'PackagingPhotoDescription'?: string;
  'ProcessingJobCategoryName'?: string;
  'ProcessingJobTypeName'?: string;
  'ProductImageFileSystemIds'?: string;
  'ProductPDFFileSystemIds'?: string;
  'ProductPhotoDescription'?: string;
  'PublicIngredients'?: string;
  'ServingSize'?: string;
  'Strain'?: string;
  'SupplyDurationDays'?: number;
  'UnitCbdAContent'?: number;
  'UnitCbdAContentDose'?: number;
  'UnitCbdAContentDoseUnitOfMeasure'?: string;
  'UnitCbdAContentUnitOfMeasure'?: string;
  'UnitCbdAPercent'?: number;
  'UnitCbdContent'?: number;
  'UnitCbdContentDose'?: number;
  'UnitCbdContentDoseUnitOfMeasure'?: string;
  'UnitCbdContentUnitOfMeasure'?: string;
  'UnitCbdPercent'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcAContent'?: number;
  'UnitThcAContentDose'?: number;
  'UnitThcAContentDoseUnitOfMeasure'?: string;
  'UnitThcAContentUnitOfMeasure'?: string;
  'UnitThcAPercent'?: number;
  'UnitThcContent'?: number;
  'UnitThcContentDose'?: number;
  'UnitThcContentDoseUnitOfMeasure'?: string;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitVolume'?: number;
  'UnitVolumeUnitOfMeasure'?: string;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface ItemsCreateBrandV2RequestItem {
  'Name'?: string;
}

export interface ItemsCreateFileV2RequestItem {
  'EncodedImageBase64'?: string;
  'FileName'?: string;
}

export interface ItemsCreatePhotoV1RequestItem {
  'EncodedImageBase64'?: string;
  'FileName'?: string;
}

export interface ItemsCreatePhotoV2RequestItem {
  'EncodedImageBase64'?: string;
  'FileName'?: string;
}

export interface ItemsCreateUpdateV1RequestItem {
  'AdministrationMethod'?: string;
  'Allergens'?: string;
  'Description'?: string;
  'GlobalProductName'?: string;
  'Id'?: number;
  'ItemBrand'?: string;
  'ItemCategory'?: string;
  'ItemIngredients'?: string;
  'LabelImageFileSystemIds'?: string;
  'LabelPhotoDescription'?: string;
  'Name'?: string;
  'NumberOfDoses'?: string;
  'PackagingImageFileSystemIds'?: string;
  'PackagingPhotoDescription'?: string;
  'ProcessingJobCategoryName'?: string;
  'ProcessingJobTypeName'?: string;
  'ProductImageFileSystemIds'?: string;
  'ProductPDFFileSystemIds'?: string;
  'ProductPhotoDescription'?: string;
  'PublicIngredients'?: string;
  'ServingSize'?: string;
  'Strain'?: string;
  'SupplyDurationDays'?: number;
  'UnitCbdAContent'?: number;
  'UnitCbdAContentDose'?: number;
  'UnitCbdAContentDoseUnitOfMeasure'?: string;
  'UnitCbdAContentUnitOfMeasure'?: string;
  'UnitCbdAPercent'?: number;
  'UnitCbdContent'?: number;
  'UnitCbdContentDose'?: number;
  'UnitCbdContentDoseUnitOfMeasure'?: string;
  'UnitCbdContentUnitOfMeasure'?: string;
  'UnitCbdPercent'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcAContent'?: number;
  'UnitThcAContentDose'?: number;
  'UnitThcAContentDoseUnitOfMeasure'?: string;
  'UnitThcAContentUnitOfMeasure'?: string;
  'UnitThcAPercent'?: number;
  'UnitThcContent'?: number;
  'UnitThcContentDose'?: number;
  'UnitThcContentDoseUnitOfMeasure'?: string;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitVolume'?: number;
  'UnitVolumeUnitOfMeasure'?: string;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface ItemsUpdateV2RequestItem {
  'AdministrationMethod'?: string;
  'Allergens'?: string;
  'Description'?: string;
  'GlobalProductName'?: string;
  'Id'?: number;
  'ItemBrand'?: string;
  'ItemCategory'?: string;
  'ItemIngredients'?: string;
  'LabelImageFileSystemIds'?: string;
  'LabelPhotoDescription'?: string;
  'Name'?: string;
  'NumberOfDoses'?: string;
  'PackagingImageFileSystemIds'?: string;
  'PackagingPhotoDescription'?: string;
  'ProcessingJobCategoryName'?: string;
  'ProcessingJobTypeName'?: string;
  'ProductImageFileSystemIds'?: string;
  'ProductPDFFileSystemIds'?: string;
  'ProductPhotoDescription'?: string;
  'PublicIngredients'?: string;
  'ServingSize'?: string;
  'Strain'?: string;
  'SupplyDurationDays'?: number;
  'UnitCbdAContent'?: number;
  'UnitCbdAContentDose'?: number;
  'UnitCbdAContentDoseUnitOfMeasure'?: string;
  'UnitCbdAContentUnitOfMeasure'?: string;
  'UnitCbdAPercent'?: number;
  'UnitCbdContent'?: number;
  'UnitCbdContentDose'?: number;
  'UnitCbdContentDoseUnitOfMeasure'?: string;
  'UnitCbdContentUnitOfMeasure'?: string;
  'UnitCbdPercent'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcAContent'?: number;
  'UnitThcAContentDose'?: number;
  'UnitThcAContentDoseUnitOfMeasure'?: string;
  'UnitThcAContentUnitOfMeasure'?: string;
  'UnitThcAPercent'?: number;
  'UnitThcContent'?: number;
  'UnitThcContentDose'?: number;
  'UnitThcContentDoseUnitOfMeasure'?: string;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitVolume'?: number;
  'UnitVolumeUnitOfMeasure'?: string;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface ItemsUpdateBrandV2RequestItem {
  'Id'?: number;
  'Name'?: string;
}

export interface PackagesCreateV1RequestItem {
  'ActualDate'?: string;
  'ExpirationDate'?: string;
  'Ingredients'?: PackagesCreateV1RequestItem_Ingredients[];
  'IsDonation'?: boolean;
  'IsProductionBatch'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'LabTestStageId'?: number;
  'Location'?: string;
  'Note'?: string;
  'PatientLicenseNumber'?: string;
  'ProcessingJobTypeId'?: number;
  'ProductRequiresRemediation'?: boolean;
  'ProductionBatchNumber'?: string;
  'Quantity'?: number;
  'RequiredLabTestBatches'?: boolean;
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UnitOfMeasure'?: string;
  'UseByDate'?: string;
  'UseSameItem'?: boolean;
}

export interface PackagesCreateV1RequestItem_Ingredients {
  'Package'?: string;
  'Quantity'?: number;
  'UnitOfMeasure'?: string;
}

export interface PackagesCreateV2RequestItem {
  'ActualDate'?: string;
  'ExpirationDate'?: string;
  'Ingredients'?: PackagesCreateV2RequestItem_Ingredients[];
  'IsDonation'?: boolean;
  'IsProductionBatch'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'LabTestStageId'?: number;
  'Location'?: string;
  'Note'?: string;
  'PatientLicenseNumber'?: string;
  'ProcessingJobTypeId'?: number;
  'ProductRequiresRemediation'?: boolean;
  'ProductionBatchNumber'?: string;
  'Quantity'?: number;
  'RequiredLabTestBatches'?: boolean;
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UnitOfMeasure'?: string;
  'UseByDate'?: string;
  'UseSameItem'?: boolean;
}

export interface PackagesCreateV2RequestItem_Ingredients {
  'Package'?: string;
  'Quantity'?: number;
  'UnitOfMeasure'?: string;
}

export interface PackagesCreateAdjustV1RequestItem {
  'AdjustmentDate'?: string;
  'AdjustmentReason'?: string;
  'Label'?: string;
  'Quantity'?: number;
  'ReasonNote'?: string;
  'UnitOfMeasure'?: string;
}

export interface PackagesCreateAdjustV2RequestItem {
  'AdjustmentDate'?: string;
  'AdjustmentReason'?: string;
  'Label'?: string;
  'Quantity'?: number;
  'ReasonNote'?: string;
  'UnitOfMeasure'?: string;
}

export interface PackagesCreateChangeItemV1RequestItem {
  'Item'?: string;
  'Label'?: string;
}

export interface PackagesCreateChangeLocationV1RequestItem {
  'Label'?: string;
  'Location'?: string;
  'MoveDate'?: string;
  'Sublocation'?: string;
}

export interface PackagesCreateFinishV1RequestItem {
  'ActualDate'?: string;
  'Label'?: string;
}

export interface PackagesCreatePlantingsV1RequestItem {
  'LocationName'?: string;
  'PackageAdjustmentAmount'?: number;
  'PackageAdjustmentUnitOfMeasureName'?: string;
  'PackageLabel'?: string;
  'PatientLicenseNumber'?: string;
  'PlantBatchName'?: string;
  'PlantBatchType'?: string;
  'PlantCount'?: number;
  'PlantedDate'?: string;
  'StrainName'?: string;
  'SublocationName'?: string;
  'UnpackagedDate'?: string;
}

export interface PackagesCreatePlantingsV2RequestItem {
  'LocationName'?: string;
  'PackageAdjustmentAmount'?: number;
  'PackageAdjustmentUnitOfMeasureName'?: string;
  'PackageLabel'?: string;
  'PatientLicenseNumber'?: string;
  'PlantBatchName'?: string;
  'PlantBatchType'?: string;
  'PlantCount'?: number;
  'PlantedDate'?: string;
  'StrainName'?: string;
  'SublocationName'?: string;
  'UnpackagedDate'?: string;
}

export interface PackagesCreateRemediateV1RequestItem {
  'PackageLabel'?: string;
  'RemediationDate'?: string;
  'RemediationMethodName'?: string;
  'RemediationSteps'?: string;
}

export interface PackagesCreateTestingV1RequestItem {
  'ActualDate'?: string;
  'ExpirationDate'?: string;
  'Ingredients'?: PackagesCreateTestingV1RequestItem_Ingredients[];
  'IsDonation'?: boolean;
  'IsProductionBatch'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'LabTestStageId'?: number;
  'Location'?: string;
  'Note'?: string;
  'PatientLicenseNumber'?: string;
  'ProcessingJobTypeId'?: number;
  'ProductRequiresRemediation'?: boolean;
  'ProductionBatchNumber'?: string;
  'Quantity'?: number;
  'RequiredLabTestBatches'?: boolean;
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UnitOfMeasure'?: string;
  'UseByDate'?: string;
  'UseSameItem'?: boolean;
}

export interface PackagesCreateTestingV1RequestItem_Ingredients {
  'Package'?: string;
  'Quantity'?: number;
  'UnitOfMeasure'?: string;
}

export interface PackagesCreateTestingV2RequestItem {
  'ActualDate'?: string;
  'ExpirationDate'?: string;
  'Ingredients'?: PackagesCreateTestingV2RequestItem_Ingredients[];
  'IsDonation'?: boolean;
  'IsProductionBatch'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'LabTestStageId'?: number;
  'Location'?: string;
  'Note'?: string;
  'PatientLicenseNumber'?: string;
  'ProcessingJobTypeId'?: number;
  'ProductRequiresRemediation'?: boolean;
  'ProductionBatchNumber'?: string;
  'Quantity'?: number;
  'RequiredLabTestBatches'?: string[];
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UnitOfMeasure'?: string;
  'UseByDate'?: string;
  'UseSameItem'?: boolean;
}

export interface PackagesCreateTestingV2RequestItem_Ingredients {
  'Package'?: string;
  'Quantity'?: number;
  'UnitOfMeasure'?: string;
}

export interface PackagesCreateUnfinishV1RequestItem {
  'Label'?: string;
}

export interface PackagesUpdateAdjustV2RequestItem {
  'AdjustmentDate'?: string;
  'AdjustmentReason'?: string;
  'Label'?: string;
  'Quantity'?: number;
  'ReasonNote'?: string;
  'UnitOfMeasure'?: string;
}

export interface PackagesUpdateChangeNoteV1RequestItem {
  'Note'?: string;
  'PackageLabel'?: string;
}

export interface PackagesUpdateDecontaminateV2RequestItem {
  'DecontaminationDate'?: string;
  'DecontaminationMethodName'?: string;
  'DecontaminationSteps'?: string;
  'PackageLabel'?: string;
}

export interface PackagesUpdateDonationFlagV2RequestItem {
  'Label'?: string;
}

export interface PackagesUpdateDonationUnflagV2RequestItem {
  'Label'?: string;
}

export interface PackagesUpdateExternalidV2RequestItem {
  'ExternalId'?: string;
  'PackageLabel'?: string;
}

export interface PackagesUpdateFinishV2RequestItem {
  'ActualDate'?: string;
  'Label'?: string;
}

export interface PackagesUpdateItemV2RequestItem {
  'Item'?: string;
  'Label'?: string;
}

export interface PackagesUpdateLabTestRequiredV2RequestItem {
  'Label'?: string;
  'RequiredLabTestBatches'?: string[];
}

export interface PackagesUpdateLocationV2RequestItem {
  'Label'?: string;
  'Location'?: string;
  'MoveDate'?: string;
  'Sublocation'?: string;
}

export interface PackagesUpdateNoteV2RequestItem {
  'Note'?: string;
  'PackageLabel'?: string;
}

export interface PackagesUpdateRemediateV2RequestItem {
  'PackageLabel'?: string;
  'RemediationDate'?: string;
  'RemediationMethodName'?: string;
  'RemediationSteps'?: string;
}

export interface PackagesUpdateTradesampleFlagV2RequestItem {
  'Label'?: string;
}

export interface PackagesUpdateTradesampleUnflagV2RequestItem {
  'Label'?: string;
}

export interface PackagesUpdateUnfinishV2RequestItem {
  'Label'?: string;
}

export interface PackagesUpdateUsebydateV2RequestItem {
  'ExpirationDate'?: string;
  'Label'?: string;
  'SellByDate'?: string;
  'UseByDate'?: string;
}

export interface PatientcheckinsCreateV1RequestItem {
  'CheckInDate'?: string;
  'CheckInLocationId'?: number;
  'PatientLicenseNumber'?: string;
  'RegistrationExpiryDate'?: string;
  'RegistrationStartDate'?: string;
}

export interface PatientcheckinsCreateV2RequestItem {
  'CheckInDate'?: string;
  'CheckInLocationId'?: number;
  'PatientLicenseNumber'?: string;
  'RegistrationExpiryDate'?: string;
  'RegistrationStartDate'?: string;
}

export interface PatientcheckinsUpdateV1RequestItem {
  'CheckInDate'?: string;
  'CheckInLocationId'?: number;
  'Id'?: number;
  'PatientLicenseNumber'?: string;
  'RegistrationExpiryDate'?: string;
  'RegistrationStartDate'?: string;
}

export interface PatientcheckinsUpdateV2RequestItem {
  'CheckInDate'?: string;
  'CheckInLocationId'?: number;
  'Id'?: number;
  'PatientLicenseNumber'?: string;
  'RegistrationExpiryDate'?: string;
  'RegistrationStartDate'?: string;
}

export interface PlantbatchesCreateAdditivesV1RequestItem {
  'ActiveIngredients'?: PlantbatchesCreateAdditivesV1RequestItem_ActiveIngredients[];
  'ActualDate'?: string;
  'AdditiveType'?: string;
  'ApplicationDevice'?: string;
  'EpaRegistrationNumber'?: string;
  'PlantBatchName'?: string;
  'ProductSupplier'?: string;
  'ProductTradeName'?: string;
  'TotalAmountApplied'?: number;
  'TotalAmountUnitOfMeasure'?: string;
}

export interface PlantbatchesCreateAdditivesV1RequestItem_ActiveIngredients {
  'Name'?: string;
  'Percentage'?: number;
}

export interface PlantbatchesCreateAdditivesV2RequestItem {
  'ActiveIngredients'?: PlantbatchesCreateAdditivesV2RequestItem_ActiveIngredients[];
  'ActualDate'?: string;
  'AdditiveType'?: string;
  'ApplicationDevice'?: string;
  'EpaRegistrationNumber'?: string;
  'PlantBatchName'?: string;
  'ProductSupplier'?: string;
  'ProductTradeName'?: string;
  'TotalAmountApplied'?: number;
  'TotalAmountUnitOfMeasure'?: string;
}

export interface PlantbatchesCreateAdditivesV2RequestItem_ActiveIngredients {
  'Name'?: string;
  'Percentage'?: number;
}

export interface PlantbatchesCreateAdditivesUsingtemplateV2RequestItem {
  'ActualDate'?: string;
  'AdditivesTemplateName'?: string;
  'PlantBatchName'?: string;
  'Rate'?: string;
  'TotalAmountApplied'?: number;
  'TotalAmountUnitOfMeasure'?: string;
  'Volume'?: string;
}

export interface PlantbatchesCreateAdjustV1RequestItem {
  'AdjustmentDate'?: string;
  'AdjustmentReason'?: string;
  'PlantBatchName'?: string;
  'Quantity'?: number;
  'ReasonNote'?: string;
}

export interface PlantbatchesCreateAdjustV2RequestItem {
  'AdjustmentDate'?: string;
  'AdjustmentReason'?: string;
  'PlantBatchName'?: string;
  'Quantity'?: number;
  'ReasonNote'?: string;
}

export interface PlantbatchesCreateChangegrowthphaseV1RequestItem {
  'Count'?: number;
  'CountPerPlant'?: string;
  'GrowthDate'?: string;
  'GrowthPhase'?: string;
  'Name'?: string;
  'NewLocation'?: string;
  'NewSublocation'?: string;
  'PatientLicenseNumber'?: string;
  'StartingTag'?: string;
}

export interface PlantbatchesCreateGrowthphaseV2RequestItem {
  'Count'?: number;
  'CountPerPlant'?: string;
  'GrowthDate'?: string;
  'GrowthPhase'?: string;
  'Name'?: string;
  'NewLocation'?: string;
  'NewSublocation'?: string;
  'PatientLicenseNumber'?: string;
  'StartingTag'?: string;
}

export interface PlantbatchesCreatePackageV2RequestItem {
  'ActualDate'?: string;
  'Count'?: number;
  'ExpirationDate'?: string;
  'Id'?: number;
  'IsDonation'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'Location'?: string;
  'Note'?: string;
  'PatientLicenseNumber'?: string;
  'PlantBatch'?: string;
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UseByDate'?: string;
}

export interface PlantbatchesCreatePackageFrommotherplantV1RequestItem {
  'ActualDate'?: string;
  'Count'?: number;
  'ExpirationDate'?: string;
  'Id'?: number;
  'IsDonation'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'Location'?: string;
  'Note'?: string;
  'PatientLicenseNumber'?: string;
  'PlantBatch'?: string;
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UseByDate'?: string;
}

export interface PlantbatchesCreatePackageFrommotherplantV2RequestItem {
  'ActualDate'?: string;
  'Count'?: number;
  'ExpirationDate'?: string;
  'Id'?: number;
  'IsDonation'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'Location'?: string;
  'Note'?: string;
  'PatientLicenseNumber'?: string;
  'PlantBatch'?: string;
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UseByDate'?: string;
}

export interface PlantbatchesCreatePlantingsV2RequestItem {
  'ActualDate'?: string;
  'Count'?: number;
  'Location'?: string;
  'Name'?: string;
  'PatientLicenseNumber'?: string;
  'SourcePlantBatches'?: string;
  'Strain'?: string;
  'Sublocation'?: string;
  'Type'?: string;
}

export interface PlantbatchesCreateSplitV1RequestItem {
  'ActualDate'?: string;
  'Count'?: number;
  'GroupName'?: string;
  'Location'?: string;
  'PatientLicenseNumber'?: string;
  'PlantBatch'?: string;
  'Strain'?: string;
  'Sublocation'?: string;
}

export interface PlantbatchesCreateSplitV2RequestItem {
  'ActualDate'?: string;
  'Count'?: number;
  'GroupName'?: string;
  'Location'?: string;
  'PatientLicenseNumber'?: string;
  'PlantBatch'?: string;
  'Strain'?: string;
  'Sublocation'?: string;
}

export interface PlantbatchesCreateWasteV1RequestItem {
  'MixedMaterial'?: string;
  'Note'?: string;
  'PlantBatchName'?: string;
  'ReasonName'?: string;
  'UnitOfMeasureName'?: string;
  'WasteDate'?: string;
  'WasteMethodName'?: string;
  'WasteWeight'?: number;
}

export interface PlantbatchesCreateWasteV2RequestItem {
  'MixedMaterial'?: string;
  'Note'?: string;
  'PlantBatchName'?: string;
  'ReasonName'?: string;
  'UnitOfMeasureName'?: string;
  'WasteDate'?: string;
  'WasteMethodName'?: string;
  'WasteWeight'?: number;
}

export interface PlantbatchesCreatepackagesV1RequestItem {
  'ActualDate'?: string;
  'Count'?: number;
  'ExpirationDate'?: string;
  'Id'?: number;
  'IsDonation'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'Location'?: string;
  'Note'?: string;
  'PatientLicenseNumber'?: string;
  'PlantBatch'?: string;
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UseByDate'?: string;
}

export interface PlantbatchesCreateplantingsV1RequestItem {
  'ActualDate'?: string;
  'Count'?: number;
  'Location'?: string;
  'Name'?: string;
  'PatientLicenseNumber'?: string;
  'SourcePlantBatches'?: string;
  'Strain'?: string;
  'Sublocation'?: string;
  'Type'?: string;
}

export interface PlantbatchesUpdateLocationV2RequestItem {
  'Location'?: string;
  'MoveDate'?: string;
  'Name'?: string;
  'Sublocation'?: string;
}

export interface PlantbatchesUpdateMoveplantbatchesV1RequestItem {
  'Location'?: string;
  'MoveDate'?: string;
  'Name'?: string;
  'Sublocation'?: string;
}

export interface PlantbatchesUpdateNameV2RequestItem {
  'Group'?: string;
  'Id'?: number;
  'NewGroup'?: string;
}

export interface PlantbatchesUpdateStrainV2RequestItem {
  'Id'?: number;
  'Name'?: string;
  'StrainId'?: number;
  'StrainName'?: string;
}

export interface PlantbatchesUpdateTagV2RequestItem {
  'Group'?: string;
  'Id'?: number;
  'NewTag'?: string;
  'ReplaceDate'?: string;
  'TagId'?: number;
}

export interface ProcessingjobsCreateAdjustV1RequestItem {
  'AdjustmentDate'?: string;
  'AdjustmentNote'?: string;
  'AdjustmentReason'?: string;
  'CountUnitOfMeasureName'?: string;
  'Id'?: number;
  'Packages'?: ProcessingjobsCreateAdjustV1RequestItem_Packages[];
  'VolumeUnitOfMeasureName'?: string;
  'WeightUnitOfMeasureName'?: string;
}

export interface ProcessingjobsCreateAdjustV1RequestItem_Packages {
  'Label'?: string;
  'Quantity'?: number;
  'UnitOfMeasure'?: string;
}

export interface ProcessingjobsCreateAdjustV2RequestItem {
  'AdjustmentDate'?: string;
  'AdjustmentNote'?: string;
  'AdjustmentReason'?: string;
  'CountUnitOfMeasureName'?: string;
  'Id'?: number;
  'Packages'?: ProcessingjobsCreateAdjustV2RequestItem_Packages[];
  'VolumeUnitOfMeasureName'?: string;
  'WeightUnitOfMeasureName'?: string;
}

export interface ProcessingjobsCreateAdjustV2RequestItem_Packages {
  'Label'?: string;
  'Quantity'?: number;
  'UnitOfMeasure'?: string;
}

export interface ProcessingjobsCreateJobtypesV1RequestItem {
  'Attributes'?: string[];
  'Category'?: string;
  'Description'?: string;
  'Name'?: string;
  'ProcessingSteps'?: string;
}

export interface ProcessingjobsCreateJobtypesV2RequestItem {
  'Attributes'?: string[];
  'Category'?: string;
  'Description'?: string;
  'Name'?: string;
  'ProcessingSteps'?: string;
}

export interface ProcessingjobsCreateStartV1RequestItem {
  'CountUnitOfMeasure'?: string;
  'JobName'?: string;
  'JobType'?: string;
  'Packages'?: ProcessingjobsCreateStartV1RequestItem_Packages[];
  'StartDate'?: string;
  'VolumeUnitOfMeasure'?: string;
  'WeightUnitOfMeasure'?: string;
}

export interface ProcessingjobsCreateStartV1RequestItem_Packages {
  'Label'?: string;
  'Quantity'?: number;
  'UnitOfMeasure'?: string;
}

export interface ProcessingjobsCreateStartV2RequestItem {
  'CountUnitOfMeasure'?: string;
  'JobName'?: string;
  'JobType'?: string;
  'Packages'?: ProcessingjobsCreateStartV2RequestItem_Packages[];
  'StartDate'?: string;
  'VolumeUnitOfMeasure'?: string;
  'WeightUnitOfMeasure'?: string;
}

export interface ProcessingjobsCreateStartV2RequestItem_Packages {
  'Label'?: string;
  'Quantity'?: number;
  'UnitOfMeasure'?: string;
}

export interface ProcessingjobsCreatepackagesV1RequestItem {
  'ExpirationDate'?: string;
  'FinishDate'?: string;
  'FinishNote'?: string;
  'FinishProcessingJob'?: boolean;
  'Item'?: string;
  'JobName'?: string;
  'Location'?: string;
  'Note'?: string;
  'PackageDate'?: string;
  'PatientLicenseNumber'?: string;
  'ProductionBatchNumber'?: string;
  'Quantity'?: number;
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UnitOfMeasure'?: string;
  'UseByDate'?: string;
  'WasteCountQuantity'?: string;
  'WasteCountUnitOfMeasureName'?: string;
  'WasteVolumeQuantity'?: string;
  'WasteVolumeUnitOfMeasureName'?: string;
  'WasteWeightQuantity'?: string;
  'WasteWeightUnitOfMeasureName'?: string;
}

export interface ProcessingjobsCreatepackagesV2RequestItem {
  'ExpirationDate'?: string;
  'FinishDate'?: string;
  'FinishNote'?: string;
  'FinishProcessingJob'?: boolean;
  'Item'?: string;
  'JobName'?: string;
  'Location'?: string;
  'Note'?: string;
  'PackageDate'?: string;
  'PatientLicenseNumber'?: string;
  'ProductionBatchNumber'?: string;
  'Quantity'?: number;
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UnitOfMeasure'?: string;
  'UseByDate'?: string;
  'WasteCountQuantity'?: string;
  'WasteCountUnitOfMeasureName'?: string;
  'WasteVolumeQuantity'?: string;
  'WasteVolumeUnitOfMeasureName'?: string;
  'WasteWeightQuantity'?: string;
  'WasteWeightUnitOfMeasureName'?: string;
}

export interface ProcessingjobsUpdateFinishV1RequestItem {
  'FinishDate'?: string;
  'FinishNote'?: string;
  'Id'?: number;
  'TotalCountWaste'?: string;
  'TotalVolumeWaste'?: string;
  'TotalWeightWaste'?: number;
  'WasteCountUnitOfMeasureName'?: string;
  'WasteVolumeUnitOfMeasureName'?: string;
  'WasteWeightUnitOfMeasureName'?: string;
}

export interface ProcessingjobsUpdateFinishV2RequestItem {
  'FinishDate'?: string;
  'FinishNote'?: string;
  'Id'?: number;
  'TotalCountWaste'?: string;
  'TotalVolumeWaste'?: string;
  'TotalWeightWaste'?: number;
  'WasteCountUnitOfMeasureName'?: string;
  'WasteVolumeUnitOfMeasureName'?: string;
  'WasteWeightUnitOfMeasureName'?: string;
}

export interface ProcessingjobsUpdateJobtypesV1RequestItem {
  'Attributes'?: string[];
  'CategoryName'?: string;
  'Description'?: string;
  'Id'?: number;
  'Name'?: string;
  'ProcessingSteps'?: string;
}

export interface ProcessingjobsUpdateJobtypesV2RequestItem {
  'Attributes'?: string[];
  'CategoryName'?: string;
  'Description'?: string;
  'Id'?: number;
  'Name'?: string;
  'ProcessingSteps'?: string;
}

export interface ProcessingjobsUpdateUnfinishV1RequestItem {
  'Id'?: number;
}

export interface ProcessingjobsUpdateUnfinishV2RequestItem {
  'Id'?: number;
}

export interface SublocationsCreateV2RequestItem {
  'Name'?: string;
}

export interface SublocationsUpdateV2RequestItem {
  'Id'?: number;
  'Name'?: string;
}

export interface SalesCreateDeliveryV1RequestItem {
  'ConsumerId'?: number;
  'DriverEmployeeId'?: string;
  'DriverName'?: string;
  'DriversLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'PatientLicenseNumber'?: string;
  'PhoneNumberForQuestions'?: string;
  'PlannedRoute'?: string;
  'RecipientAddressCity'?: string;
  'RecipientAddressCounty'?: string;
  'RecipientAddressPostalCode'?: string;
  'RecipientAddressState'?: string;
  'RecipientAddressStreet1'?: string;
  'RecipientAddressStreet2'?: string;
  'RecipientName'?: string;
  'RecipientZoneId'?: number;
  'SalesCustomerType'?: string;
  'SalesDateTime'?: string;
  'Transactions'?: SalesCreateDeliveryV1RequestItem_Transactions[];
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface SalesCreateDeliveryV1RequestItem_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesCreateDeliveryV2RequestItem {
  'ConsumerId'?: number;
  'DriverEmployeeId'?: string;
  'DriverName'?: string;
  'DriversLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'PatientLicenseNumber'?: string;
  'PhoneNumberForQuestions'?: string;
  'PlannedRoute'?: string;
  'RecipientAddressCity'?: string;
  'RecipientAddressCounty'?: string;
  'RecipientAddressPostalCode'?: string;
  'RecipientAddressState'?: string;
  'RecipientAddressStreet1'?: string;
  'RecipientAddressStreet2'?: string;
  'RecipientName'?: string;
  'RecipientZoneId'?: number;
  'SalesCustomerType'?: string;
  'SalesDateTime'?: string;
  'Transactions'?: SalesCreateDeliveryV2RequestItem_Transactions[];
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface SalesCreateDeliveryV2RequestItem_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesCreateDeliveryRetailerV1RequestItem {
  'DateTime'?: string;
  'Destinations'?: SalesCreateDeliveryRetailerV1RequestItem_Destinations[];
  'DriverEmployeeId'?: string;
  'DriverName'?: string;
  'DriversLicenseNumber'?: string;
  'EstimatedDepartureDateTime'?: string;
  'Packages'?: SalesCreateDeliveryRetailerV1RequestItem_Packages[];
  'PhoneNumberForQuestions'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface SalesCreateDeliveryRetailerV1RequestItem_Destinations {
  'ConsumerId'?: string;
  'EstimatedArrivalDateTime'?: string;
  'PatientLicenseNumber'?: string;
  'RecipientAddressCity'?: string;
  'RecipientAddressCounty'?: string;
  'RecipientAddressPostalCode'?: string;
  'RecipientAddressState'?: string;
  'RecipientAddressStreet1'?: string;
  'RecipientAddressStreet2'?: string;
  'RecipientName'?: string;
  'RecipientZoneId'?: string;
  'SalesCustomerType'?: string;
  'Transactions'?: SalesCreateDeliveryRetailerV1RequestItem_Destinations_Transactions[];
}

export interface SalesCreateDeliveryRetailerV1RequestItem_Destinations_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesCreateDeliveryRetailerV1RequestItem_Packages {
  'DateTime'?: string;
  'PackageLabel'?: string;
  'Quantity'?: number;
  'TotalPrice'?: number;
  'UnitOfMeasure'?: string;
}

export interface SalesCreateDeliveryRetailerV2RequestItem {
  'DateTime'?: string;
  'Destinations'?: SalesCreateDeliveryRetailerV2RequestItem_Destinations[];
  'DriverEmployeeId'?: string;
  'DriverName'?: string;
  'DriversLicenseNumber'?: string;
  'EstimatedDepartureDateTime'?: string;
  'Packages'?: SalesCreateDeliveryRetailerV2RequestItem_Packages[];
  'PhoneNumberForQuestions'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface SalesCreateDeliveryRetailerV2RequestItem_Destinations {
  'ConsumerId'?: string;
  'EstimatedArrivalDateTime'?: string;
  'PatientLicenseNumber'?: string;
  'RecipientAddressCity'?: string;
  'RecipientAddressCounty'?: string;
  'RecipientAddressPostalCode'?: string;
  'RecipientAddressState'?: string;
  'RecipientAddressStreet1'?: string;
  'RecipientAddressStreet2'?: string;
  'RecipientName'?: string;
  'RecipientZoneId'?: string;
  'SalesCustomerType'?: string;
  'Transactions'?: SalesCreateDeliveryRetailerV2RequestItem_Destinations_Transactions[];
}

export interface SalesCreateDeliveryRetailerV2RequestItem_Destinations_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesCreateDeliveryRetailerV2RequestItem_Packages {
  'DateTime'?: string;
  'PackageLabel'?: string;
  'Quantity'?: number;
  'TotalPrice'?: number;
  'UnitOfMeasure'?: string;
}

export interface SalesCreateDeliveryRetailerDepartV1RequestItem {
  'RetailerDeliveryId'?: number;
}

export interface SalesCreateDeliveryRetailerDepartV2RequestItem {
  'RetailerDeliveryId'?: number;
}

export interface SalesCreateDeliveryRetailerEndV1RequestItem {
  'ActualArrivalDateTime'?: string;
  'Packages'?: SalesCreateDeliveryRetailerEndV1RequestItem_Packages[];
  'RetailerDeliveryId'?: number;
}

export interface SalesCreateDeliveryRetailerEndV1RequestItem_Packages {
  'EndQuantity'?: number;
  'EndUnitOfMeasure'?: string;
  'Label'?: string;
}

export interface SalesCreateDeliveryRetailerEndV2RequestItem {
  'ActualArrivalDateTime'?: string;
  'Packages'?: SalesCreateDeliveryRetailerEndV2RequestItem_Packages[];
  'RetailerDeliveryId'?: number;
}

export interface SalesCreateDeliveryRetailerEndV2RequestItem_Packages {
  'EndQuantity'?: number;
  'EndUnitOfMeasure'?: string;
  'Label'?: string;
}

export interface SalesCreateDeliveryRetailerRestockV1RequestItem {
  'DateTime'?: string;
  'Destinations'?: string;
  'EstimatedDepartureDateTime'?: string;
  'Packages'?: SalesCreateDeliveryRetailerRestockV1RequestItem_Packages[];
  'RetailerDeliveryId'?: number;
}

export interface SalesCreateDeliveryRetailerRestockV1RequestItem_Packages {
  'PackageLabel'?: string;
  'Quantity'?: number;
  'RemoveCurrentPackage'?: boolean;
  'TotalPrice'?: number;
  'UnitOfMeasure'?: string;
}

export interface SalesCreateDeliveryRetailerRestockV2RequestItem {
  'DateTime'?: string;
  'Destinations'?: string;
  'EstimatedDepartureDateTime'?: string;
  'Packages'?: SalesCreateDeliveryRetailerRestockV2RequestItem_Packages[];
  'RetailerDeliveryId'?: number;
}

export interface SalesCreateDeliveryRetailerRestockV2RequestItem_Packages {
  'PackageLabel'?: string;
  'Quantity'?: number;
  'RemoveCurrentPackage'?: boolean;
  'TotalPrice'?: number;
  'UnitOfMeasure'?: string;
}

export interface SalesCreateDeliveryRetailerSaleV1RequestItem {
  'ConsumerId'?: number;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'PatientLicenseNumber'?: string;
  'PhoneNumberForQuestions'?: string;
  'PlannedRoute'?: string;
  'RecipientAddressCity'?: string;
  'RecipientAddressCounty'?: string;
  'RecipientAddressPostalCode'?: string;
  'RecipientAddressState'?: string;
  'RecipientAddressStreet1'?: string;
  'RecipientAddressStreet2'?: string;
  'RecipientName'?: string;
  'RecipientZoneId'?: number;
  'RetailerDeliveryId'?: number;
  'SalesCustomerType'?: string;
  'SalesDateTime'?: string;
  'Transactions'?: SalesCreateDeliveryRetailerSaleV1RequestItem_Transactions[];
}

export interface SalesCreateDeliveryRetailerSaleV1RequestItem_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesCreateDeliveryRetailerSaleV2RequestItem {
  'ConsumerId'?: number;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'PatientLicenseNumber'?: string;
  'PhoneNumberForQuestions'?: string;
  'PlannedRoute'?: string;
  'RecipientAddressCity'?: string;
  'RecipientAddressCounty'?: string;
  'RecipientAddressPostalCode'?: string;
  'RecipientAddressState'?: string;
  'RecipientAddressStreet1'?: string;
  'RecipientAddressStreet2'?: string;
  'RecipientName'?: string;
  'RecipientZoneId'?: number;
  'RetailerDeliveryId'?: number;
  'SalesCustomerType'?: string;
  'SalesDateTime'?: string;
  'Transactions'?: SalesCreateDeliveryRetailerSaleV2RequestItem_Transactions[];
}

export interface SalesCreateDeliveryRetailerSaleV2RequestItem_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesCreateReceiptV1RequestItem {
  'CaregiverLicenseNumber'?: string;
  'ExternalReceiptNumber'?: string;
  'IdentificationMethod'?: string;
  'PatientLicenseNumber'?: string;
  'PatientRegistrationLocationId'?: number;
  'SalesCustomerType'?: string;
  'SalesDateTime'?: string;
  'Transactions'?: SalesCreateReceiptV1RequestItem_Transactions[];
}

export interface SalesCreateReceiptV1RequestItem_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesCreateReceiptV2RequestItem {
  'CaregiverLicenseNumber'?: string;
  'ExternalReceiptNumber'?: string;
  'IdentificationMethod'?: string;
  'PatientLicenseNumber'?: string;
  'PatientRegistrationLocationId'?: number;
  'SalesCustomerType'?: string;
  'SalesDateTime'?: string;
  'Transactions'?: SalesCreateReceiptV2RequestItem_Transactions[];
}

export interface SalesCreateReceiptV2RequestItem_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesCreateTransactionByDateV1RequestItem {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesUpdateDeliveryV1RequestItem {
  'ConsumerId'?: number;
  'DriverEmployeeId'?: string;
  'DriverName'?: string;
  'DriversLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'Id'?: number;
  'PatientLicenseNumber'?: string;
  'PhoneNumberForQuestions'?: string;
  'PlannedRoute'?: string;
  'RecipientAddressCity'?: string;
  'RecipientAddressCounty'?: string;
  'RecipientAddressPostalCode'?: string;
  'RecipientAddressState'?: string;
  'RecipientAddressStreet1'?: string;
  'RecipientAddressStreet2'?: string;
  'RecipientName'?: string;
  'RecipientZoneId'?: string;
  'SalesCustomerType'?: string;
  'SalesDateTime'?: string;
  'Transactions'?: SalesUpdateDeliveryV1RequestItem_Transactions[];
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface SalesUpdateDeliveryV1RequestItem_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesUpdateDeliveryV2RequestItem {
  'ConsumerId'?: number;
  'DriverEmployeeId'?: string;
  'DriverName'?: string;
  'DriversLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'Id'?: number;
  'PatientLicenseNumber'?: string;
  'PhoneNumberForQuestions'?: string;
  'PlannedRoute'?: string;
  'RecipientAddressCity'?: string;
  'RecipientAddressCounty'?: string;
  'RecipientAddressPostalCode'?: string;
  'RecipientAddressState'?: string;
  'RecipientAddressStreet1'?: string;
  'RecipientAddressStreet2'?: string;
  'RecipientName'?: string;
  'RecipientZoneId'?: string;
  'SalesCustomerType'?: string;
  'SalesDateTime'?: string;
  'Transactions'?: SalesUpdateDeliveryV2RequestItem_Transactions[];
  'TransporterFacilityLicenseNumber'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface SalesUpdateDeliveryV2RequestItem_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesUpdateDeliveryCompleteV1RequestItem {
  'AcceptedPackages'?: string[];
  'ActualArrivalDateTime'?: string;
  'Id'?: number;
  'PaymentType'?: string;
  'ReturnedPackages'?: SalesUpdateDeliveryCompleteV1RequestItem_ReturnedPackages[];
}

export interface SalesUpdateDeliveryCompleteV1RequestItem_ReturnedPackages {
  'Label'?: string;
  'ReturnQuantityVerified'?: number;
  'ReturnReason'?: string;
  'ReturnReasonNote'?: string;
  'ReturnUnitOfMeasure'?: string;
}

export interface SalesUpdateDeliveryCompleteV2RequestItem {
  'AcceptedPackages'?: string[];
  'ActualArrivalDateTime'?: string;
  'Id'?: number;
  'PaymentType'?: string;
  'ReturnedPackages'?: SalesUpdateDeliveryCompleteV2RequestItem_ReturnedPackages[];
}

export interface SalesUpdateDeliveryCompleteV2RequestItem_ReturnedPackages {
  'Label'?: string;
  'ReturnQuantityVerified'?: number;
  'ReturnReason'?: string;
  'ReturnReasonNote'?: string;
  'ReturnUnitOfMeasure'?: string;
}

export interface SalesUpdateDeliveryHubV1RequestItem {
  'DriverEmployeeId'?: string;
  'DriverName'?: string;
  'DriversLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'Id'?: number;
  'PhoneNumberForQuestions'?: string;
  'PlannedRoute'?: string;
  'TransporterFacilityId'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface SalesUpdateDeliveryHubV2RequestItem {
  'DriverEmployeeId'?: string;
  'DriverName'?: string;
  'DriversLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'Id'?: number;
  'PhoneNumberForQuestions'?: string;
  'PlannedRoute'?: string;
  'TransporterFacilityId'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface SalesUpdateDeliveryHubAcceptV1RequestItem {
  'Id'?: number;
}

export interface SalesUpdateDeliveryHubAcceptV2RequestItem {
  'Id'?: number;
}

export interface SalesUpdateDeliveryHubDepartV1RequestItem {
  'Id'?: number;
}

export interface SalesUpdateDeliveryHubDepartV2RequestItem {
  'Id'?: number;
}

export interface SalesUpdateDeliveryHubVerifyIdV1RequestItem {
  'Id'?: number;
  'PaymentType'?: string;
}

export interface SalesUpdateDeliveryHubVerifyIdV2RequestItem {
  'Id'?: number;
  'PaymentType'?: string;
}

export interface SalesUpdateDeliveryRetailerV1RequestItem {
  'DateTime'?: string;
  'Destinations'?: SalesUpdateDeliveryRetailerV1RequestItem_Destinations[];
  'DriverEmployeeId'?: string;
  'DriverName'?: string;
  'DriversLicenseNumber'?: string;
  'EstimatedDepartureDateTime'?: string;
  'Id'?: number;
  'Packages'?: SalesUpdateDeliveryRetailerV1RequestItem_Packages[];
  'PhoneNumberForQuestions'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface SalesUpdateDeliveryRetailerV1RequestItem_Destinations {
  'ConsumerId'?: string;
  'DriverEmployeeId'?: number;
  'DriverName'?: string;
  'DriversLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'Id'?: number;
  'PatientLicenseNumber'?: string;
  'PhoneNumberForQuestions'?: string;
  'PlannedRoute'?: string;
  'RecipientAddressCity'?: string;
  'RecipientAddressCounty'?: string;
  'RecipientAddressPostalCode'?: string;
  'RecipientAddressState'?: string;
  'RecipientAddressStreet1'?: string;
  'RecipientAddressStreet2'?: string;
  'RecipientName'?: string;
  'RecipientZoneId'?: string;
  'SalesCustomerType'?: string;
  'SalesDateTime'?: string;
  'Transactions'?: SalesUpdateDeliveryRetailerV1RequestItem_Destinations_Transactions[];
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface SalesUpdateDeliveryRetailerV1RequestItem_Destinations_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesUpdateDeliveryRetailerV1RequestItem_Packages {
  'DateTime'?: string;
  'PackageLabel'?: string;
  'Quantity'?: number;
  'TotalPrice'?: number;
  'UnitOfMeasure'?: string;
}

export interface SalesUpdateDeliveryRetailerV2RequestItem {
  'DateTime'?: string;
  'Destinations'?: SalesUpdateDeliveryRetailerV2RequestItem_Destinations[];
  'DriverEmployeeId'?: string;
  'DriverName'?: string;
  'DriversLicenseNumber'?: string;
  'EstimatedDepartureDateTime'?: string;
  'Id'?: number;
  'Packages'?: SalesUpdateDeliveryRetailerV2RequestItem_Packages[];
  'PhoneNumberForQuestions'?: string;
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface SalesUpdateDeliveryRetailerV2RequestItem_Destinations {
  'ConsumerId'?: string;
  'DriverEmployeeId'?: number;
  'DriverName'?: string;
  'DriversLicenseNumber'?: string;
  'EstimatedArrivalDateTime'?: string;
  'EstimatedDepartureDateTime'?: string;
  'Id'?: number;
  'PatientLicenseNumber'?: string;
  'PhoneNumberForQuestions'?: string;
  'PlannedRoute'?: string;
  'RecipientAddressCity'?: string;
  'RecipientAddressCounty'?: string;
  'RecipientAddressPostalCode'?: string;
  'RecipientAddressState'?: string;
  'RecipientAddressStreet1'?: string;
  'RecipientAddressStreet2'?: string;
  'RecipientName'?: string;
  'RecipientZoneId'?: string;
  'SalesCustomerType'?: string;
  'SalesDateTime'?: string;
  'Transactions'?: SalesUpdateDeliveryRetailerV2RequestItem_Destinations_Transactions[];
  'VehicleLicensePlateNumber'?: string;
  'VehicleMake'?: string;
  'VehicleModel'?: string;
}

export interface SalesUpdateDeliveryRetailerV2RequestItem_Destinations_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesUpdateDeliveryRetailerV2RequestItem_Packages {
  'DateTime'?: string;
  'PackageLabel'?: string;
  'Quantity'?: number;
  'TotalPrice'?: number;
  'UnitOfMeasure'?: string;
}

export interface SalesUpdateReceiptV1RequestItem {
  'CaregiverLicenseNumber'?: string;
  'ExternalReceiptNumber'?: string;
  'Id'?: number;
  'IdentificationMethod'?: string;
  'PatientLicenseNumber'?: string;
  'PatientRegistrationLocationId'?: number;
  'SalesCustomerType'?: string;
  'SalesDateTime'?: string;
  'Transactions'?: SalesUpdateReceiptV1RequestItem_Transactions[];
}

export interface SalesUpdateReceiptV1RequestItem_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesUpdateReceiptV2RequestItem {
  'CaregiverLicenseNumber'?: string;
  'ExternalReceiptNumber'?: string;
  'Id'?: number;
  'IdentificationMethod'?: string;
  'PatientLicenseNumber'?: string;
  'PatientRegistrationLocationId'?: number;
  'SalesCustomerType'?: string;
  'SalesDateTime'?: string;
  'Transactions'?: SalesUpdateReceiptV2RequestItem_Transactions[];
}

export interface SalesUpdateReceiptV2RequestItem_Transactions {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface SalesUpdateReceiptFinalizeV2RequestItem {
  'Id'?: number;
}

export interface SalesUpdateReceiptUnfinalizeV2RequestItem {
  'Id'?: number;
}

export interface SalesUpdateTransactionByDateV1RequestItem {
  'CityTax'?: string;
  'CountyTax'?: string;
  'DiscountAmount'?: string;
  'ExciseTax'?: string;
  'InvoiceNumber'?: string;
  'MunicipalTax'?: string;
  'PackageLabel'?: string;
  'Price'?: string;
  'QrCodes'?: string;
  'Quantity'?: number;
  'SalesTax'?: string;
  'SubTotal'?: string;
  'TotalAmount'?: number;
  'UnitOfMeasure'?: string;
  'UnitThcContent'?: number;
  'UnitThcContentUnitOfMeasure'?: string;
  'UnitThcPercent'?: number;
  'UnitWeight'?: number;
  'UnitWeightUnitOfMeasure'?: string;
}

export interface StrainsCreateV1RequestItem {
  'CbdLevel'?: number;
  'IndicaPercentage'?: number;
  'Name'?: string;
  'SativaPercentage'?: number;
  'TestingStatus'?: string;
  'ThcLevel'?: number;
}

export interface StrainsCreateV2RequestItem {
  'CbdLevel'?: number;
  'IndicaPercentage'?: number;
  'Name'?: string;
  'SativaPercentage'?: number;
  'TestingStatus'?: string;
  'ThcLevel'?: number;
}

export interface StrainsCreateUpdateV1RequestItem {
  'CbdLevel'?: number;
  'Id'?: number;
  'IndicaPercentage'?: number;
  'Name'?: string;
  'SativaPercentage'?: number;
  'TestingStatus'?: string;
  'ThcLevel'?: number;
}

export interface StrainsUpdateV2RequestItem {
  'CbdLevel'?: number;
  'Id'?: number;
  'IndicaPercentage'?: number;
  'Name'?: string;
  'SativaPercentage'?: number;
  'TestingStatus'?: string;
  'ThcLevel'?: number;
}

export interface AdditivestemplatesCreateV2RequestItem {
  'ActiveIngredients'?: AdditivestemplatesCreateV2RequestItem_ActiveIngredients[];
  'AdditiveType'?: string;
  'ApplicationDevice'?: string;
  'EpaRegistrationNumber'?: string;
  'Name'?: string;
  'Note'?: string;
  'ProductSupplier'?: string;
  'ProductTradeName'?: string;
  'RestrictiveEntryIntervalQuantityDescription'?: string;
  'RestrictiveEntryIntervalTimeDescription'?: string;
}

export interface AdditivestemplatesCreateV2RequestItem_ActiveIngredients {
  'Name'?: string;
  'Percentage'?: number;
}

export interface AdditivestemplatesUpdateV2RequestItem {
  'ActiveIngredients'?: AdditivestemplatesUpdateV2RequestItem_ActiveIngredients[];
  'AdditiveType'?: string;
  'ApplicationDevice'?: string;
  'EpaRegistrationNumber'?: string;
  'Id'?: number;
  'Name'?: string;
  'Note'?: string;
  'ProductSupplier'?: string;
  'ProductTradeName'?: string;
  'RestrictiveEntryIntervalQuantityDescription'?: string;
  'RestrictiveEntryIntervalTimeDescription'?: string;
}

export interface AdditivestemplatesUpdateV2RequestItem_ActiveIngredients {
  'Name'?: string;
  'Percentage'?: number;
}

export interface HarvestsCreateFinishV1RequestItem {
  'ActualDate'?: string;
  'Id'?: number;
}

export interface HarvestsCreatePackageV1RequestItem {
  'ActualDate'?: string;
  'DecontaminateProduct'?: boolean;
  'DecontaminationDate'?: string;
  'DecontaminationSteps'?: string;
  'ExpirationDate'?: string;
  'Ingredients'?: HarvestsCreatePackageV1RequestItem_Ingredients[];
  'IsDonation'?: boolean;
  'IsProductionBatch'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'LabTestStageId'?: number;
  'Location'?: string;
  'Note'?: string;
  'PatientLicenseNumber'?: string;
  'ProcessingJobTypeId'?: number;
  'ProductRequiresDecontamination'?: boolean;
  'ProductRequiresRemediation'?: boolean;
  'ProductionBatchNumber'?: string;
  'RemediateProduct'?: boolean;
  'RemediationDate'?: string;
  'RemediationMethodId'?: number;
  'RemediationSteps'?: string;
  'RequiredLabTestBatches'?: any[];
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UnitOfWeight'?: string;
  'UseByDate'?: string;
}

export interface HarvestsCreatePackageV1RequestItem_Ingredients {
  'HarvestId'?: number;
  'HarvestName'?: string;
  'UnitOfWeight'?: string;
  'Weight'?: number;
}

export interface HarvestsCreatePackageV2RequestItem {
  'ActualDate'?: string;
  'DecontaminateProduct'?: boolean;
  'DecontaminationDate'?: string;
  'DecontaminationSteps'?: string;
  'ExpirationDate'?: string;
  'Ingredients'?: HarvestsCreatePackageV2RequestItem_Ingredients[];
  'IsDonation'?: boolean;
  'IsProductionBatch'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'LabTestStageId'?: number;
  'Location'?: string;
  'Note'?: string;
  'PatientLicenseNumber'?: string;
  'ProcessingJobTypeId'?: number;
  'ProductRequiresDecontamination'?: boolean;
  'ProductRequiresRemediation'?: boolean;
  'ProductionBatchNumber'?: string;
  'RemediateProduct'?: boolean;
  'RemediationDate'?: string;
  'RemediationMethodId'?: number;
  'RemediationSteps'?: string;
  'RequiredLabTestBatches'?: any[];
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UnitOfWeight'?: string;
  'UseByDate'?: string;
}

export interface HarvestsCreatePackageV2RequestItem_Ingredients {
  'HarvestId'?: number;
  'HarvestName'?: string;
  'UnitOfWeight'?: string;
  'Weight'?: number;
}

export interface HarvestsCreatePackageTestingV1RequestItem {
  'ActualDate'?: string;
  'DecontaminateProduct'?: boolean;
  'DecontaminationDate'?: string;
  'DecontaminationSteps'?: string;
  'ExpirationDate'?: string;
  'Ingredients'?: HarvestsCreatePackageTestingV1RequestItem_Ingredients[];
  'IsDonation'?: boolean;
  'IsProductionBatch'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'LabTestStageId'?: number;
  'Location'?: string;
  'Note'?: string;
  'PatientLicenseNumber'?: string;
  'ProcessingJobTypeId'?: number;
  'ProductRequiresDecontamination'?: boolean;
  'ProductRequiresRemediation'?: boolean;
  'ProductionBatchNumber'?: string;
  'RemediateProduct'?: boolean;
  'RemediationDate'?: string;
  'RemediationMethodId'?: number;
  'RemediationSteps'?: string;
  'RequiredLabTestBatches'?: any[];
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UnitOfWeight'?: string;
  'UseByDate'?: string;
}

export interface HarvestsCreatePackageTestingV1RequestItem_Ingredients {
  'HarvestId'?: number;
  'HarvestName'?: string;
  'UnitOfWeight'?: string;
  'Weight'?: number;
}

export interface HarvestsCreatePackageTestingV2RequestItem {
  'ActualDate'?: string;
  'DecontaminateProduct'?: boolean;
  'DecontaminationDate'?: string;
  'DecontaminationSteps'?: string;
  'ExpirationDate'?: string;
  'Ingredients'?: HarvestsCreatePackageTestingV2RequestItem_Ingredients[];
  'IsDonation'?: boolean;
  'IsProductionBatch'?: boolean;
  'IsTradeSample'?: boolean;
  'Item'?: string;
  'LabTestStageId'?: number;
  'Location'?: string;
  'Note'?: string;
  'PatientLicenseNumber'?: string;
  'ProcessingJobTypeId'?: number;
  'ProductRequiresDecontamination'?: boolean;
  'ProductRequiresRemediation'?: boolean;
  'ProductionBatchNumber'?: string;
  'RemediateProduct'?: boolean;
  'RemediationDate'?: string;
  'RemediationMethodId'?: number;
  'RemediationSteps'?: string;
  'RequiredLabTestBatches'?: any[];
  'SellByDate'?: string;
  'Sublocation'?: string;
  'Tag'?: string;
  'UnitOfWeight'?: string;
  'UseByDate'?: string;
}

export interface HarvestsCreatePackageTestingV2RequestItem_Ingredients {
  'HarvestId'?: number;
  'HarvestName'?: string;
  'UnitOfWeight'?: string;
  'Weight'?: number;
}

export interface HarvestsCreateRemoveWasteV1RequestItem {
  'ActualDate'?: string;
  'Id'?: number;
  'UnitOfWeight'?: string;
  'WasteType'?: string;
  'WasteWeight'?: number;
}

export interface HarvestsCreateUnfinishV1RequestItem {
  'Id'?: number;
}

export interface HarvestsCreateWasteV2RequestItem {
  'ActualDate'?: string;
  'Id'?: number;
  'UnitOfWeight'?: string;
  'WasteType'?: string;
  'WasteWeight'?: number;
}

export interface HarvestsUpdateFinishV2RequestItem {
  'ActualDate'?: string;
  'Id'?: number;
}

export interface HarvestsUpdateLocationV2RequestItem {
  'ActualDate'?: string;
  'DryingLocation'?: string;
  'DryingSublocation'?: string;
  'HarvestName'?: string;
  'Id'?: number;
}

export interface HarvestsUpdateMoveV1RequestItem {
  'ActualDate'?: string;
  'DryingLocation'?: string;
  'DryingSublocation'?: string;
  'HarvestName'?: string;
  'Id'?: number;
}

export interface HarvestsUpdateRenameV1RequestItem {
  'Id'?: number;
  'NewName'?: string;
  'OldName'?: string;
}

export interface HarvestsUpdateRenameV2RequestItem {
  'Id'?: number;
  'NewName'?: string;
  'OldName'?: string;
}

export interface HarvestsUpdateRestoreHarvestedPlantsV2RequestItem {
  'HarvestId'?: number;
  'PlantTags'?: string[];
}

export interface HarvestsUpdateUnfinishV2RequestItem {
  'Id'?: number;
}

export interface LabtestsCreateRecordV1RequestItem {
  'DocumentFileBase64'?: string;
  'DocumentFileName'?: string;
  'Label'?: string;
  'ResultDate'?: string;
  'Results'?: LabtestsCreateRecordV1RequestItem_Results[];
}

export interface LabtestsCreateRecordV1RequestItem_Results {
  'LabTestTypeName'?: string;
  'Notes'?: string;
  'Passed'?: boolean;
  'Quantity'?: number;
}

export interface LabtestsCreateRecordV2RequestItem {
  'DocumentFileBase64'?: string;
  'DocumentFileName'?: string;
  'Label'?: string;
  'ResultDate'?: string;
  'Results'?: LabtestsCreateRecordV2RequestItem_Results[];
}

export interface LabtestsCreateRecordV2RequestItem_Results {
  'LabTestTypeName'?: string;
  'Notes'?: string;
  'Passed'?: boolean;
  'Quantity'?: number;
}

export interface LabtestsUpdateLabtestdocumentV1RequestItem {
  'DocumentFileBase64'?: string;
  'DocumentFileName'?: string;
  'LabTestResultId'?: number;
}

export interface LabtestsUpdateLabtestdocumentV2RequestItem {
  'DocumentFileBase64'?: string;
  'DocumentFileName'?: string;
  'LabTestResultId'?: number;
}

export interface LabtestsUpdateResultReleaseV1RequestItem {
  'PackageLabel'?: string;
}

export interface LabtestsUpdateResultReleaseV2RequestItem {
  'PackageLabel'?: string;
}

export interface LocationsCreateV1RequestItem {
  'LocationTypeName'?: string;
  'Name'?: string;
}

export interface LocationsCreateV2RequestItem {
  'LocationTypeName'?: string;
  'Name'?: string;
}

export interface LocationsCreateUpdateV1RequestItem {
  'Id'?: number;
  'LocationTypeName'?: string;
  'Name'?: string;
}

export interface LocationsUpdateV2RequestItem {
  'Id'?: number;
  'LocationTypeName'?: string;
  'Name'?: string;
}

