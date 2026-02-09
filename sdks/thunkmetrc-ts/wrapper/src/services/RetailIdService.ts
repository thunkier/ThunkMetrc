
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    Allotment,
    CreateAssociateRequestItem,
    CreateGenerateRequest,
    CreateMergeRequest,
    CreatePackagesInfoRequest,
    Receive,
    ReceiveQrByShortCode,
    WriteResponse,
} from '../models';

export class RetailIdService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
   * - WebApi Retail ID Read Write State (All or WriteOnly)
   * - Industry/View Packages
   * - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
   * POST CreateAssociate
   */
  public async createRetailIdAssociate(body: CreateAssociateRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createRetailIdAssociate(body,licenseNumber,
        )
    );
  }

  /**
   * Allows you to generate a specific quantity of QR codes. Id value returned (issuance ID) could be used for printing.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
   * - WebApi Retail ID Read Write State (All or WriteOnly)
   * - Industry/View Packages
   * - One of the following: Industry/Facility Type/Can Download Product Label, Licensee/Download Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
   * POST CreateGenerate
   */
  public async createRetailIdGenerate(body: CreateGenerateRequest, licenseNumber?: string): Promise<WriteResponse[]> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createRetailIdGenerate(body,licenseNumber,
        )
    );
  }

  /**
   * Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
   * - WebApi Retail ID Read Write State (All or WriteOnly)
   * - Key Value Settings/Retail ID Merge Packages Enabled
   * POST CreateMerge
   */
  public async createRetailIdMerge(body: CreateMergeRequest, licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createRetailIdMerge(body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves Package information for given list of Package labels.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
   * - WebApi Retail ID Read Write State (All or WriteOnly)
   * - Industry/View Packages
   * - Admin/Employees/Packages Page/Product Labels(Manage)
   * POST CreatePackagesInfo
   */
  public async createRetailIdPackagesInfo(body: CreatePackagesInfoRequest, licenseNumber?: string): Promise<WriteResponse[]> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createRetailIdPackagesInfo(body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves the available Retail Item ID quota for a facility.
   * Permissions Required:
   * - Download Product Labels
   * - Manage Product Labels
   * - Manage Tag Orders
   * GET GetAllotment
   */
  public async getRetailIdAllotment(body?: any, licenseNumber?: string): Promise<Allotment[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getRetailIdAllotment(body,licenseNumber,
        )
    );
  }

  /**
   * Get a list of eaches (Retail ID QR code URL) and sibling tags based on given Package label.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Manage RetailId
   * - WebApi Retail ID Read Write State (All or ReadOnly)
   * - Industry/View Packages
   * - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
   * GET GetReceiveByLabel
   */
  public async getRetailIdReceiveByLabel(label: string, body?: any, licenseNumber?: string): Promise<Receive[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getRetailIdReceiveByLabel(label, body,licenseNumber,
        )
    );
  }

  /**
   * Get a list of eaches (Retail ID QR code URL) and sibling tags based on given short code value (first segment in Retail ID QR code URL).
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Manage RetailId
   * - WebApi Retail ID Read Write State (All or ReadOnly)
   * - Industry/View Packages
   * - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(View or Manage)
   * GET GetReceiveQrByShortCode
   */
  public async getRetailIdReceiveQrByShortCode(shortCode: string, body?: any, licenseNumber?: string): Promise<ReceiveQrByShortCode[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getRetailIdReceiveQrByShortCode(shortCode, body,licenseNumber,
        )
    );
  }
}

