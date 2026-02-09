
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    AdjustReason,
    Adjustment,
    CreateAdjustPackagesRequestItem,
    CreatePackagesPackagesRequestItem,
    CreatePackagesPlantingsRequestItem,
    CreateTestingRequestItem,
    FinishPackagesRequestItem,
    FinishedgoodFlagRequestItem,
    FinishedgoodUnflagRequestItem,
    InTransit,
    PackagesPackage,
    SourceHarvest,
    UnfinishPackagesRequestItem,
    UpdateAdjustPackagesRequestItem,
    UpdateDecontaminateRequestItem,
    UpdateDonationFlagRequestItem,
    UpdateDonationUnflagRequestItem,
    UpdateExternalidRequestItem,
    UpdateItemRequestItem,
    UpdateLabtestsRequiredRequestItem,
    UpdateNoteRequestItem,
    UpdatePackagesLocationRequestItem,
    UpdateRemediateRequestItem,
    UpdateTradeSampleFlagRequestItem,
    UpdateTradeSampleUnflagRequestItem,
    UpdateUseByDateRequestItem,
    WriteResponse,
} from '../models';

export class PackagesService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Records a list of adjustments for packages at a specific Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * POST CreateAdjustPackages
   */
  public async createAdjustPackages(body: CreateAdjustPackagesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createAdjustPackages(body,licenseNumber,
        )
    );
  }

  /**
   * Creates new packages for a specified Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * POST CreatePackagesPackages
   */
  public async createPackagesPackages(body: CreatePackagesPackagesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPackagesPackages(body,licenseNumber,
        )
    );
  }

  /**
   * Creates new plantings from packages for a specified Facility.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants
   * - View Packages
   * - Manage Packages Inventory
   * POST CreatePackagesPlantings
   */
  public async createPackagesPlantings(body: CreatePackagesPlantingsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPackagesPlantings(body,licenseNumber,
        )
    );
  }

  /**
   * Creates new packages for testing for a specified Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * POST CreateTesting
   */
  public async createPackagesTesting(body: CreateTestingRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createPackagesTesting(body,licenseNumber,
        )
    );
  }

  /**
   * Discontinues a Package at a specific Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * DELETE DeletePackagesById
   */
  public async deletePackagesById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deletePackagesById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Updates a list of packages as finished for a specific Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT FinishPackages
   */
  public async finishPackagesPackages(body: FinishPackagesRequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.finishPackagesPackages(body,licenseNumber,
        )
    );
  }

  /**
   * Flags one or more Packages at the specified Facility as Finished Goods.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT FinishedgoodFlag
   */
  public async finishedgoodFlagPackages(body: FinishedgoodFlagRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.finishedgoodFlagPackages(body,licenseNumber,
        )
    );
  }

  /**
   * Removes the Finished Good flag one or more Packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT FinishedgoodUnflag
   */
  public async finishedgoodUnflagPackages(body: FinishedgoodUnflagRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.finishedgoodUnflagPackages(body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of active packages for a specified Facility.
   * Permissions Required:
   * - View Packages
   * GET GetActivePackages
   */
  public async getActivePackages(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<PackagesPackage>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getActivePackages(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of adjustment reasons for packages at a specified Facility.
   * GET GetAdjustReasons
   */
  public async getPackagesAdjustReasons(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<AdjustReason>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPackagesAdjustReasons(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves the Package Adjustments for a Facility
   * Permissions Required:
   * - View Packages
   * GET GetAdjustments
   */
  public async getPackagesAdjustments(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Adjustment>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPackagesAdjustments(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of packages in transit for a specified Facility.
   * Permissions Required:
   * - View Packages
   * GET GetInTransitPackages
   */
  public async getInTransitPackages(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<InTransit>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getInTransitPackages(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of inactive packages for a specified Facility.
   * Permissions Required:
   * - View Packages
   * GET GetInactivePackages
   */
  public async getInactivePackages(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<PackagesPackage>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getInactivePackages(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
   * Permissions Required:
   * - View Packages
   * GET GetLabSamples
   */
  public async getPackagesLabSamples(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<PackagesPackage>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPackagesLabSamples(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of packages on hold for a specified Facility.
   * Permissions Required:
   * - View Packages
   * GET GetOnHoldPackages
   */
  public async getOnHoldPackages(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<PackagesPackage>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getOnHoldPackages(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a Package by its Id.
   * Permissions Required:
   * - View Packages
   * GET GetPackagesById
   */
  public async getPackagesById(id: string, body?: any, licenseNumber?: string): Promise<PackagesPackage> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPackagesById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a Package by its label.
   * Permissions Required:
   * - View Packages
   * GET GetPackagesByLabel
   */
  public async getPackagesByLabel(label: string, body?: any, licenseNumber?: string): Promise<PackagesPackage[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPackagesByLabel(label, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of available Package types.
   * GET GetPackagesTypes
   */
  public async getPackagesTypes(body?: any): Promise<any> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPackagesTypes(body,
        )
    );
  }

  /**
   * Retrieves the source harvests for a Package by its Id.
   * Permissions Required:
   * - View Package Source Harvests
   * GET GetSourceHarvestById
   */
  public async getPackagesSourceHarvestById(id: string, body?: any, licenseNumber?: string): Promise<SourceHarvest> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPackagesSourceHarvestById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of transferred packages for a specific Facility.
   * Permissions Required:
   * - View Packages
   * GET GetTransferred
   */
  public async getPackagesTransferred(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<PackagesPackage>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getPackagesTransferred(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Updates a list of packages as unfinished for a specific Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UnfinishPackages
   */
  public async unfinishPackagesPackages(body: UnfinishPackagesRequestItem[], licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.unfinishPackagesPackages(body,licenseNumber,
        )
    );
  }

  /**
   * Set the final quantity for a Package.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateAdjustPackages
   */
  public async updateAdjustPackages(body: UpdateAdjustPackagesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateAdjustPackages(body,licenseNumber,
        )
    );
  }

  /**
   * Updates the Product decontaminate information for a list of packages at a specific Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateDecontaminate
   */
  public async updatePackagesDecontaminate(body: UpdateDecontaminateRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePackagesDecontaminate(body,licenseNumber,
        )
    );
  }

  /**
   * Flags one or more packages for donation at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateDonationFlag
   */
  public async updatePackagesDonationFlag(body: UpdateDonationFlagRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePackagesDonationFlag(body,licenseNumber,
        )
    );
  }

  /**
   * Removes the donation flag from one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateDonationUnflag
   */
  public async updatePackagesDonationUnflag(body: UpdateDonationUnflagRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePackagesDonationUnflag(body,licenseNumber,
        )
    );
  }

  /**
   * Updates the external identifiers for one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Package Inventory
   * - External Id Enabled
   * PUT UpdateExternalid
   */
  public async updatePackagesExternalid(body: UpdateExternalidRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePackagesExternalid(body,licenseNumber,
        )
    );
  }

  /**
   * Updates the associated Item for one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * PUT UpdateItem
   */
  public async updatePackagesItem(body: UpdateItemRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePackagesItem(body,licenseNumber,
        )
    );
  }

  /**
   * Updates the list of required lab test batches for one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * PUT UpdateLabtestsRequired
   */
  public async updatePackagesLabtestsRequired(body: UpdateLabtestsRequiredRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePackagesLabtestsRequired(body,licenseNumber,
        )
    );
  }

  /**
   * Updates notes associated with one or more packages for the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * - Manage Package Notes
   * PUT UpdateNote
   */
  public async updatePackagesNote(body: UpdateNoteRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePackagesNote(body,licenseNumber,
        )
    );
  }

  /**
   * Updates the Location and Sublocation for one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * PUT UpdatePackagesLocation
   */
  public async updatePackagesLocation(body: UpdatePackagesLocationRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePackagesLocation(body,licenseNumber,
        )
    );
  }

  /**
   * Updates a list of Product remediations for packages at a specific Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateRemediate
   */
  public async updatePackagesRemediate(body: UpdateRemediateRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePackagesRemediate(body,licenseNumber,
        )
    );
  }

  /**
   * Flags or unflags one or more packages at the specified Facility as trade samples.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateTradeSampleFlag
   */
  public async updatePackagesTradeSampleFlag(body: UpdateTradeSampleFlagRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePackagesTradeSampleFlag(body,licenseNumber,
        )
    );
  }

  /**
   * Removes the trade sample flag from one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateTradeSampleUnflag
   */
  public async updatePackagesTradeSampleUnflag(body: UpdateTradeSampleUnflagRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePackagesTradeSampleUnflag(body,licenseNumber,
        )
    );
  }

  /**
   * Updates the use-by date for one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * PUT UpdateUseByDate
   */
  public async updatePackagesUseByDate(body: UpdateUseByDateRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updatePackagesUseByDate(body,licenseNumber,
        )
    );
  }
}

