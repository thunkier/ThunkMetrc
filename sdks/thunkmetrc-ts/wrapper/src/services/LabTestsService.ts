
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    Batch,
    CreateRecordRequestItem,
    LabTestsType,
    Result,
    UpdateLabTestDocumentRequestItem,
    UpdateResultsReleaseRequestItem,
    WriteResponse,
} from '../models';

export class LabTestsService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * POST CreateRecord
   */
  public async createLabTestsRecord(body: CreateRecordRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createLabTestsRecord(body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of Lab Test batches.
   * GET GetBatches
   */
  public async getLabTestsBatches(body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Batch>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getLabTestsBatches(body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a specific Lab Test result document by its Id for a given Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * GET GetLabTestDocumentById
   */
  public async getLabTestsLabTestDocumentById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getLabTestsLabTestDocumentById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Returns a list of Lab Test types.
   * GET GetLabTestsTypes
   */
  public async getLabTestsTypes(body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<LabTestsType>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getLabTestsTypes(body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves Lab Test results for a specified Package.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * GET GetResults
   */
  public async getLabTestsResults(body?: any, licenseNumber?: string, packageId?: number, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Result>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getLabTestsResults(body,licenseNumber,packageId,pageNumber,pageSize,
        )
    );
  }

  /**
   * Returns a list of all lab testing states.
   * GET GetStates
   */
  public async getLabTestsStates(body?: any): Promise<any> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getLabTestsStates(body,
        )
    );
  }

  /**
   * Updates one or more documents for previously submitted lab tests.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateLabTestDocument
   */
  public async updateLabTestsLabTestDocument(body: UpdateLabTestDocumentRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateLabTestsLabTestDocument(body,licenseNumber,
        )
    );
  }

  /**
   * Releases Lab Test results for one or more packages.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateResultsRelease
   */
  public async updateLabTestsResultsRelease(body: UpdateResultsReleaseRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateLabTestsResultsRelease(body,licenseNumber,
        )
    );
  }
}

