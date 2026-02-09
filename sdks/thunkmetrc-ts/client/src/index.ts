import axios, { AxiosInstance, AxiosResponse, AxiosError } from 'axios';

export interface Logger {
  debug(message: string): void;
  info(message: string): void;
  warn(message: string): void;
  error(message: string): void;
}

export class MetrcError extends Error {
  constructor(
    public statusCode: number,
    public message: string,
    public validationErrors: string[] = []
  ) {
    super(`API Error ${statusCode}: ${message}`);
    this.name = 'MetrcError';
  }
}

export interface MetrcClientConfig {
  baseUrl: string;
  vendorKey: string;
  userKey: string;
  timeout?: number;
  userAgent?: string;
  logger?: Logger;
}

export class MetrcClient {
  private client: AxiosInstance;
  private logger?: Logger;

  constructor(config: MetrcClientConfig) {
    this.logger = config.logger;
    this.client = axios.create({
      baseURL: config.baseUrl,
      timeout: config.timeout || 100000,
      auth: {
        username: config.vendorKey,
        password: config.userKey
      },
      headers: {
        'User-Agent': config.userAgent || 'ThunkMetrc/0.1 TypeScript'
      }
    });

    // Add logging interceptor
    this.client.interceptors.request.use(req => {
      this.logger?.debug(`Sending Request: ${req.method?.toUpperCase()} ${req.url}`);
      return req;
    });

    // Add error handling interceptor
    this.client.interceptors.response.use(
      res => res,
      (error: AxiosError<any>) => {
        const status = error.response?.status || 500;
        this.logger?.warn(`API Error Response: ${status}`);

        const data = error.response?.data;
        if (data && (data.Message || data.ValidationErrors)) {
            throw new MetrcError(status, data.Message || 'Unknown Error', data.ValidationErrors || []);
        }
        throw error;
      }
    );
  }
  /**
   * Creates new additive templates for a specified Facility.
   * Permissions Required:
   * - Manage Additives
   * POST CreateAdditivesTemplates
   */
  public async createAdditivesTemplates<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/additivestemplates/v2?${queryStr}` : `/additivestemplates/v2`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Retrieves a list of active additive templates for a specified Facility.
   * Permissions Required:
   * - Manage Additives
   * GET GetActiveAdditivesTemplates
   */
  public async getActiveAdditivesTemplates<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/additivestemplates/v2/active?${queryStr}` : `/additivestemplates/v2/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves an Additive Template by its Id.
   * Permissions Required:
   * - Manage Additives
   * GET GetAdditivesTemplatesById
   */
  public async getAdditivesTemplatesById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/additivestemplates/v2/${encodeURIComponent(id)}?${queryStr}` : `/additivestemplates/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of inactive additive templates for a specified Facility.
   * Permissions Required:
   * - Manage Additives
   * GET GetInactiveAdditivesTemplates
   */
  public async getInactiveAdditivesTemplates<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/additivestemplates/v2/inactive?${queryStr}` : `/additivestemplates/v2/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Updates existing additive templates for a specified Facility.
   * Permissions Required:
   * - Manage Additives
   * PUT UpdateAdditivesTemplates
   */
  public async updateAdditivesTemplates<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/additivestemplates/v2?${queryStr}` : `/additivestemplates/v2`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
   * Permissions Required:
   * - Lookup Caregivers
   * GET GetCaregiversStatusByCaregiverLicenseNumber
   */
  public async getCaregiversStatusByCaregiverLicenseNumber<TResponse = any, TBody = any>(caregiverLicenseNumber: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/caregivers/v2/status/${encodeURIComponent(caregiverLicenseNumber)}?${queryStr}` : `/caregivers/v2/status/${encodeURIComponent(caregiverLicenseNumber)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of employees for a specified Facility.
   * Permissions Required:
   * - Manage Employees
   * - View Employees
   * GET GetEmployees
   */
  public async getEmployees<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/employees/v2?${queryStr}` : `/employees/v2`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
   * Permissions Required:
   * - Manage Employees
   * GET GetPermissions
   */
  public async getEmployeesPermissions<TResponse = any, TBody = any>(body?: TBody, employeeLicenseNumber?: string, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (employeeLicenseNumber !== undefined) query.append('employeeLicenseNumber', String(employeeLicenseNumber));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/employees/v2/permissions?${queryStr}` : `/employees/v2/permissions`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * This endpoint provides a list of facilities for which the authenticated user has access.
   * GET GetFacilities
   */
  public async getFacilities<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/facilities/v2?${queryStr}` : `/facilities/v2`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Creates packages from harvested products for a specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Manage Harvests
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * POST CreateHarvestsPackages
   */
  public async createHarvestsPackages<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/packages?${queryStr}` : `/harvests/v2/packages`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
   * Permissions Required:
   * - View Harvests
   * - Manage Harvests
   * POST CreateHarvestsWaste
   */
  public async createHarvestsWaste<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/waste?${queryStr}` : `/harvests/v2/waste`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
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
  public async createHarvestsPackagesTesting<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/packages/testing?${queryStr}` : `/harvests/v2/packages/testing`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Discontinues a specific harvest waste record by Id for the specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Discontinue Harvest Waste
   * DELETE DeleteWasteById
   */
  public async deleteHarvestsWasteById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/waste/${encodeURIComponent(id)}?${queryStr}` : `/harvests/v2/waste/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Marks one or more harvests as finished for the specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Finish/Discontinue Harvests
   * PUT FinishHarvests
   */
  public async finishHarvestsHarvests<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/finish?${queryStr}` : `/harvests/v2/finish`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Retrieves a list of active harvests for a specified Facility.
   * Permissions Required:
   * - View Harvests
   * GET GetActiveHarvests
   */
  public async getActiveHarvests<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/active?${queryStr}` : `/harvests/v2/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
   * Permissions Required:
   * - View Harvests
   * GET GetHarvestsById
   */
  public async getHarvestsById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/${encodeURIComponent(id)}?${queryStr}` : `/harvests/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
   * Permissions Required:
   * - View Harvests
   * GET GetHarvestsWaste
   */
  public async getHarvestsWaste<TResponse = any, TBody = any>(body?: TBody, harvestId?: number, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (harvestId !== undefined) query.append('harvestId', String(harvestId));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/waste?${queryStr}` : `/harvests/v2/waste`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of inactive harvests for a specified Facility.
   * Permissions Required:
   * - View Harvests
   * GET GetInactiveHarvests
   */
  public async getInactiveHarvests<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/inactive?${queryStr}` : `/harvests/v2/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of harvests on hold for a specified Facility.
   * Permissions Required:
   * - View Harvests
   * GET GetOnHoldHarvests
   */
  public async getOnHoldHarvests<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/onhold?${queryStr}` : `/harvests/v2/onhold`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of Waste types for harvests.
   * GET GetWasteTypes
   */
  public async getHarvestsWasteTypes<TResponse = any, TBody = any>(body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/waste/types?${queryStr}` : `/harvests/v2/waste/types`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Reopens one or more previously finished harvests for the specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Finish/Discontinue Harvests
   * PUT UnfinishHarvests
   */
  public async unfinishHarvestsHarvests<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/unfinish?${queryStr}` : `/harvests/v2/unfinish`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates the Location of Harvest for a specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Manage Harvests
   * PUT UpdateHarvestsLocation
   */
  public async updateHarvestsLocation<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/location?${queryStr}` : `/harvests/v2/location`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Renames one or more harvests for the specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Manage Harvests
   * PUT UpdateRename
   */
  public async updateHarvestsRename<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/rename?${queryStr}` : `/harvests/v2/rename`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Restores previously harvested plants to their original state for the specified Facility.
   * Permissions Required:
   * - View Harvests
   * - Finish/Discontinue Harvests
   * PUT UpdateRestoreHarvestedPlants
   */
  public async updateHarvestsRestoreHarvestedPlants<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/restore/harvestedplants?${queryStr}` : `/harvests/v2/restore/harvestedplants`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates one or more new item brands for the specified Facility identified by the License Number.
   * Permissions Required:
   * - Manage Items
   * POST CreateBrand
   */
  public async createItemsBrand<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/brand?${queryStr}` : `/items/v2/brand`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
   * Permissions Required:
   * - Manage Items
   * POST CreateFile
   */
  public async createItemsFile<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/file?${queryStr}` : `/items/v2/file`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Permissions Required:
   * - Manage Items
   * POST CreateItems
   */
  public async createItems<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2?${queryStr}` : `/items/v2`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
   * Permissions Required:
   * - Manage Items
   * POST CreatePhoto
   */
  public async createItemsPhoto<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/photo?${queryStr}` : `/items/v2/photo`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Archives the specified Item Brand by Id for the given Facility License Number.
   * Permissions Required:
   * - Manage Items
   * DELETE DeleteBrandById
   */
  public async deleteItemsBrandById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/brand/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/brand/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Archives the specified Product by Id for the given Facility License Number.
   * Permissions Required:
   * - Manage Items
   * DELETE DeleteItemsById
   */
  public async deleteItemsById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of active items for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * GET GetActiveItems
   */
  public async getActiveItems<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/active?${queryStr}` : `/items/v2/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of active item brands for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * GET GetBrands
   */
  public async getItemsBrands<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/brands?${queryStr}` : `/items/v2/brands`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of item categories.
   * GET GetCategories
   */
  public async getItemsCategories<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/categories?${queryStr}` : `/items/v2/categories`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a file by its Id for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * GET GetFileById
   */
  public async getItemsFileById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/file/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/file/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of inactive items for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * GET GetInactiveItems
   */
  public async getInactiveItems<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/inactive?${queryStr}` : `/items/v2/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves detailed information about a specific Item by Id.
   * Permissions Required:
   * - Manage Items
   * GET GetItemsById
   */
  public async getItemsById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves an image by its Id for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * GET GetPhotoById
   */
  public async getItemsPhotoById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/photo/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/photo/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Updates one or more existing item brands for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * PUT UpdateBrand
   */
  public async updateItemsBrand<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/brand?${queryStr}` : `/items/v2/brand`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates one or more existing products for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * PUT UpdateItems
   */
  public async updateItems<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2?${queryStr}` : `/items/v2`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * POST CreateRecord
   */
  public async createLabTestsRecord<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/record?${queryStr}` : `/labtests/v2/record`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Retrieves a list of Lab Test batches.
   * GET GetBatches
   */
  public async getLabTestsBatches<TResponse = any, TBody = any>(body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/batches?${queryStr}` : `/labtests/v2/batches`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a specific Lab Test result document by its Id for a given Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * GET GetLabTestDocumentById
   */
  public async getLabTestsLabTestDocumentById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/labtestdocument/${encodeURIComponent(id)}?${queryStr}` : `/labtests/v2/labtestdocument/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of Lab Test types.
   * GET GetLabTestsTypes
   */
  public async getLabTestsTypes<TResponse = any, TBody = any>(body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/types?${queryStr}` : `/labtests/v2/types`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves Lab Test results for a specified Package.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * GET GetResults
   */
  public async getLabTestsResults<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, packageId?: number, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (packageId !== undefined) query.append('packageId', String(packageId));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/results?${queryStr}` : `/labtests/v2/results`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of all lab testing states.
   * GET GetStates
   */
  public async getLabTestsStates<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/states?${queryStr}` : `/labtests/v2/states`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Updates one or more documents for previously submitted lab tests.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateLabTestDocument
   */
  public async updateLabTestsLabTestDocument<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/labtestdocument?${queryStr}` : `/labtests/v2/labtestdocument`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Releases Lab Test results for one or more packages.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateResultsRelease
   */
  public async updateLabTestsResultsRelease<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/results/release?${queryStr}` : `/labtests/v2/results/release`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates new locations for a specified Facility.
   * Permissions Required:
   * - Manage Locations
   * POST CreateLocations
   */
  public async createLocations<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2?${queryStr}` : `/locations/v2`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Archives a specified Location, identified by its Id, for a Facility.
   * Permissions Required:
   * - Manage Locations
   * DELETE DeleteLocationsById
   */
  public async deleteLocationsById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2/${encodeURIComponent(id)}?${queryStr}` : `/locations/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of active locations for a specified Facility.
   * Permissions Required:
   * - Manage Locations
   * GET GetActiveLocations
   */
  public async getActiveLocations<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2/active?${queryStr}` : `/locations/v2/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of inactive locations for a specified Facility.
   * Permissions Required:
   * - Manage Locations
   * GET GetInactiveLocations
   */
  public async getInactiveLocations<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2/inactive?${queryStr}` : `/locations/v2/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a Location by its Id.
   * Permissions Required:
   * - Manage Locations
   * GET GetLocationsById
   */
  public async getLocationsById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2/${encodeURIComponent(id)}?${queryStr}` : `/locations/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of active location types for a specified Facility.
   * Permissions Required:
   * - Manage Locations
   * GET GetLocationsTypes
   */
  public async getLocationsTypes<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2/types?${queryStr}` : `/locations/v2/types`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Updates existing locations for a specified Facility.
   * Permissions Required:
   * - Manage Locations
   * PUT UpdateLocations
   */
  public async updateLocations<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2?${queryStr}` : `/locations/v2`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records a list of adjustments for packages at a specific Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * POST CreateAdjustPackages
   */
  public async createAdjustPackages<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/adjust?${queryStr}` : `/packages/v2/adjust`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates new packages for a specified Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * POST CreatePackagesPackages
   */
  public async createPackagesPackages<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2?${queryStr}` : `/packages/v2`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
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
  public async createPackagesPlantings<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/plantings?${queryStr}` : `/packages/v2/plantings`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates new packages for testing for a specified Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * POST CreateTesting
   */
  public async createPackagesTesting<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/testing?${queryStr}` : `/packages/v2/testing`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Discontinues a Package at a specific Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * DELETE DeletePackagesById
   */
  public async deletePackagesById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/${encodeURIComponent(id)}?${queryStr}` : `/packages/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Updates a list of packages as finished for a specific Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT FinishPackages
   */
  public async finishPackagesPackages<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/finish?${queryStr}` : `/packages/v2/finish`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Flags one or more Packages at the specified Facility as Finished Goods.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT FinishedgoodFlag
   */
  public async finishedgoodFlagPackages<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/finishedgood/flag?${queryStr}` : `/packages/v2/finishedgood/flag`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Removes the Finished Good flag one or more Packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT FinishedgoodUnflag
   */
  public async finishedgoodUnflagPackages<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/finishedgood/unflag?${queryStr}` : `/packages/v2/finishedgood/unflag`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Retrieves a list of active packages for a specified Facility.
   * Permissions Required:
   * - View Packages
   * GET GetActivePackages
   */
  public async getActivePackages<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/active?${queryStr}` : `/packages/v2/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of adjustment reasons for packages at a specified Facility.
   * GET GetAdjustReasons
   */
  public async getPackagesAdjustReasons<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/adjust/reasons?${queryStr}` : `/packages/v2/adjust/reasons`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves the Package Adjustments for a Facility
   * Permissions Required:
   * - View Packages
   * GET GetAdjustments
   */
  public async getPackagesAdjustments<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/adjustments?${queryStr}` : `/packages/v2/adjustments`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of packages in transit for a specified Facility.
   * Permissions Required:
   * - View Packages
   * GET GetInTransitPackages
   */
  public async getInTransitPackages<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/intransit?${queryStr}` : `/packages/v2/intransit`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of inactive packages for a specified Facility.
   * Permissions Required:
   * - View Packages
   * GET GetInactivePackages
   */
  public async getInactivePackages<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/inactive?${queryStr}` : `/packages/v2/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
   * Permissions Required:
   * - View Packages
   * GET GetLabSamples
   */
  public async getPackagesLabSamples<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/labsamples?${queryStr}` : `/packages/v2/labsamples`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of packages on hold for a specified Facility.
   * Permissions Required:
   * - View Packages
   * GET GetOnHoldPackages
   */
  public async getOnHoldPackages<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/onhold?${queryStr}` : `/packages/v2/onhold`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a Package by its Id.
   * Permissions Required:
   * - View Packages
   * GET GetPackagesById
   */
  public async getPackagesById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/${encodeURIComponent(id)}?${queryStr}` : `/packages/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a Package by its label.
   * Permissions Required:
   * - View Packages
   * GET GetPackagesByLabel
   */
  public async getPackagesByLabel<TResponse = any, TBody = any>(label: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/${encodeURIComponent(label)}?${queryStr}` : `/packages/v2/${encodeURIComponent(label)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of available Package types.
   * GET GetPackagesTypes
   */
  public async getPackagesTypes<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/types?${queryStr}` : `/packages/v2/types`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves the source harvests for a Package by its Id.
   * Permissions Required:
   * - View Package Source Harvests
   * GET GetSourceHarvestById
   */
  public async getPackagesSourceHarvestById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/${encodeURIComponent(id)}/source/harvests?${queryStr}` : `/packages/v2/${encodeURIComponent(id)}/source/harvests`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of transferred packages for a specific Facility.
   * Permissions Required:
   * - View Packages
   * GET GetTransferred
   */
  public async getPackagesTransferred<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/transferred?${queryStr}` : `/packages/v2/transferred`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Updates a list of packages as unfinished for a specific Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UnfinishPackages
   */
  public async unfinishPackagesPackages<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/unfinish?${queryStr}` : `/packages/v2/unfinish`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Set the final quantity for a Package.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateAdjustPackages
   */
  public async updateAdjustPackages<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/adjust?${queryStr}` : `/packages/v2/adjust`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates the Product decontaminate information for a list of packages at a specific Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateDecontaminate
   */
  public async updatePackagesDecontaminate<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/decontaminate?${queryStr}` : `/packages/v2/decontaminate`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Flags one or more packages for donation at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateDonationFlag
   */
  public async updatePackagesDonationFlag<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/donation/flag?${queryStr}` : `/packages/v2/donation/flag`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Removes the donation flag from one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateDonationUnflag
   */
  public async updatePackagesDonationUnflag<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/donation/unflag?${queryStr}` : `/packages/v2/donation/unflag`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates the external identifiers for one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Package Inventory
   * - External Id Enabled
   * PUT UpdateExternalid
   */
  public async updatePackagesExternalid<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/externalid?${queryStr}` : `/packages/v2/externalid`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates the associated Item for one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * PUT UpdateItem
   */
  public async updatePackagesItem<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/item?${queryStr}` : `/packages/v2/item`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates the list of required lab test batches for one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * PUT UpdateLabtestsRequired
   */
  public async updatePackagesLabtestsRequired<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/labtests/required?${queryStr}` : `/packages/v2/labtests/required`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates notes associated with one or more packages for the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * - Manage Package Notes
   * PUT UpdateNote
   */
  public async updatePackagesNote<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/note?${queryStr}` : `/packages/v2/note`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates the Location and Sublocation for one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * PUT UpdatePackagesLocation
   */
  public async updatePackagesLocation<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/location?${queryStr}` : `/packages/v2/location`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates a list of Product remediations for packages at a specific Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateRemediate
   */
  public async updatePackagesRemediate<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/remediate?${queryStr}` : `/packages/v2/remediate`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Flags or unflags one or more packages at the specified Facility as trade samples.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateTradeSampleFlag
   */
  public async updatePackagesTradeSampleFlag<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/tradesample/flag?${queryStr}` : `/packages/v2/tradesample/flag`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Removes the trade sample flag from one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Manage Packages Inventory
   * PUT UpdateTradeSampleUnflag
   */
  public async updatePackagesTradeSampleUnflag<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/tradesample/unflag?${queryStr}` : `/packages/v2/tradesample/unflag`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates the use-by date for one or more packages at the specified Facility.
   * Permissions Required:
   * - View Packages
   * - Create/Submit/Discontinue Packages
   * PUT UpdateUseByDate
   */
  public async updatePackagesUseByDate<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/usebydate?${queryStr}` : `/packages/v2/usebydate`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records patient check-ins for a specified Facility.
   * Permissions Required:
   * - ManagePatientsCheckIns
   * POST CreatePatientCheckIns
   */
  public async createPatientCheckIns<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v2?${queryStr}` : `/patient-checkins/v2`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Archives a Patient Check-In, identified by its Id, for a specified Facility.
   * Permissions Required:
   * - ManagePatientsCheckIns
   * DELETE DeletePatientCheckInsById
   */
  public async deletePatientCheckInsById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v2/${encodeURIComponent(id)}?${queryStr}` : `/patient-checkins/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of Patient Check-In locations.
   * GET GetLocations
   */
  public async getPatientCheckInsLocations<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v2/locations?${queryStr}` : `/patient-checkins/v2/locations`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of patient check-ins for a specified Facility.
   * Permissions Required:
   * - ManagePatientsCheckIns
   * GET GetPatientCheckIns
   */
  public async getPatientCheckIns<TResponse = any, TBody = any>(body?: TBody, checkinDateEnd?: string, checkinDateStart?: string, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (checkinDateEnd !== undefined) query.append('checkinDateEnd', String(checkinDateEnd));
    if (checkinDateStart !== undefined) query.append('checkinDateStart', String(checkinDateStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v2?${queryStr}` : `/patient-checkins/v2`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Updates patient check-ins for a specified Facility.
   * Permissions Required:
   * - ManagePatientsCheckIns
   * PUT UpdatePatientCheckIns
   */
  public async updatePatientCheckIns<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v2?${queryStr}` : `/patient-checkins/v2`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Adds new patients to a specified Facility.
   * Permissions Required:
   * - Manage Patients
   * POST CreatePatients
   */
  public async createPatients<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v2?${queryStr}` : `/patients/v2`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Removes a Patient, identified by an Id, from a specified Facility.
   * Permissions Required:
   * - Manage Patients
   * DELETE DeletePatientsById
   */
  public async deletePatientsById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v2/${encodeURIComponent(id)}?${queryStr}` : `/patients/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of active patients for a specified Facility.
   * Permissions Required:
   * - Manage Patients
   * GET GetActivePatients
   */
  public async getActivePatients<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v2/active?${queryStr}` : `/patients/v2/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a Patient by Id.
   * Permissions Required:
   * - Manage Patients
   * GET GetPatientsById
   */
  public async getPatientsById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v2/${encodeURIComponent(id)}?${queryStr}` : `/patients/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Updates Patient information for a specified Facility.
   * Permissions Required:
   * - Manage Patients
   * PUT UpdatePatients
   */
  public async updatePatients<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v2?${queryStr}` : `/patients/v2`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
   * Permissions Required:
   * - Lookup Patients
   * GET GetPatientsStatusesByPatientLicenseNumber
   */
  public async getPatientsStatusesByPatientLicenseNumber<TResponse = any, TBody = any>(patientLicenseNumber: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v2/statuses/${encodeURIComponent(patientLicenseNumber)}?${queryStr}` : `/patients/v2/statuses/${encodeURIComponent(patientLicenseNumber)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants Inventory
   * POST CreateAdjustPlantBatches
   */
  public async createAdjustPlantBatches<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/adjust?${queryStr}` : `/plantbatches/v2/adjust`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
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
  public async createPlantBatchesGrowthPhase<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/growthphase?${queryStr}` : `/plantbatches/v2/growthphase`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
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
  public async createPlantBatchesPackagesFromMotherPlant<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/packages/frommotherplant?${queryStr}` : `/plantbatches/v2/packages/frommotherplant`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records Additive usage details for plant batches at a specific Facility.
   * Permissions Required:
   * - Manage Plants Additives
   * POST CreatePlantBatchesAdditives
   */
  public async createPlantBatchesAdditives<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/additives?${queryStr}` : `/plantbatches/v2/additives`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records Additive usage for plant batches at a Facility using predefined additive templates.
   * Permissions Required:
   * - Manage Plants Additives
   * POST CreatePlantBatchesAdditivesUsingTemplate
   */
  public async createPlantBatchesAdditivesUsingTemplate<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/additives/usingtemplate?${queryStr}` : `/plantbatches/v2/additives/usingtemplate`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
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
  public async createPlantBatchesPackages<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/packages?${queryStr}` : `/plantbatches/v2/packages`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates new plantings for a Facility by generating plant batches based on provided planting details.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants Inventory
   * POST CreatePlantBatchesPlantings
   */
  public async createPlantBatchesPlantings<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/plantings?${queryStr}` : `/plantbatches/v2/plantings`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records waste information for plant batches based on the submitted data for the specified Facility.
   * Permissions Required:
   * - Manage Plants Waste
   * POST CreatePlantBatchesWaste
   */
  public async createPlantBatchesWaste<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/waste?${queryStr}` : `/plantbatches/v2/waste`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Splits an existing Plant Batch into multiple groups at the specified Facility.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants Inventory
   * POST CreateSplit
   */
  public async createPlantBatchesSplit<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/split?${queryStr}` : `/plantbatches/v2/split`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Completes the destruction of plant batches based on the provided input data.
   * Permissions Required:
   * - View Immature Plants
   * - Destroy Immature Plants
   * DELETE DeletePlantBatches
   */
  public async deletePlantBatches<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2?${queryStr}` : `/plantbatches/v2`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
   * Permissions Required:
   * - View Immature Plants
   * GET GetActivePlantBatches
   */
  public async getActivePlantBatches<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/active?${queryStr}` : `/plantbatches/v2/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
   * Permissions Required:
   * - View Immature Plants
   * GET GetInactivePlantBatches
   */
  public async getInactivePlantBatches<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/inactive?${queryStr}` : `/plantbatches/v2/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a Plant Batch by Id.
   * Permissions Required:
   * - View Immature Plants
   * GET GetPlantBatchesById
   */
  public async getPlantBatchesById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/${encodeURIComponent(id)}?${queryStr}` : `/plantbatches/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of plant batch types.
   * GET GetPlantBatchesTypes
   */
  public async getPlantBatchesTypes<TResponse = any, TBody = any>(body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/types?${queryStr}` : `/plantbatches/v2/types`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves waste details associated with plant batches at a specified Facility.
   * Permissions Required:
   * - View Plants Waste
   * GET GetPlantBatchesWaste
   */
  public async getPlantBatchesWaste<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/waste?${queryStr}` : `/plantbatches/v2/waste`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
   * GET GetPlantBatchesWasteReasons
   */
  public async getPlantBatchesWasteReasons<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/waste/reasons?${queryStr}` : `/plantbatches/v2/waste/reasons`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Renames plant batches at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdateName
   */
  public async updatePlantBatchesName<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/name?${queryStr}` : `/plantbatches/v2/name`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Moves one or more plant batches to new locations with in a specified Facility.
   * Permissions Required:
   * - View Immature Plants
   * - Manage Immature Plants
   * PUT UpdatePlantBatchesLocation
   */
  public async updatePlantBatchesLocation<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/location?${queryStr}` : `/plantbatches/v2/location`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Changes the strain of plant batches at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdatePlantBatchesStrain
   */
  public async updatePlantBatchesStrain<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/strain?${queryStr}` : `/plantbatches/v2/strain`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Replaces tags for plant batches at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdatePlantBatchesTag
   */
  public async updatePlantBatchesTag<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/tag?${queryStr}` : `/plantbatches/v2/tag`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records additive usage for plants based on their location within a specified Facility.
   * Permissions Required:
   * - Manage Plants
   * - Manage Plants Additives
   * POST CreateAdditivesByLocation
   */
  public async createPlantsAdditivesByLocation<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/additives/bylocation?${queryStr}` : `/plants/v2/additives/bylocation`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records additive usage for plants by location using a predefined additive template at a specified Facility.
   * Permissions Required:
   * - Manage Plants Additives
   * POST CreateAdditivesByLocationUsingTemplate
   */
  public async createPlantsAdditivesByLocationUsingTemplate<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/additives/bylocation/usingtemplate?${queryStr}` : `/plants/v2/additives/bylocation/usingtemplate`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates harvest product records from plant batches at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manicure/Harvest Veg/Flower Plants
   * POST CreateManicure
   */
  public async createPlantsManicure<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/manicure?${queryStr}` : `/plants/v2/manicure`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
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
  public async createPlantsPlantBatchPackages<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/plantbatch/packages?${queryStr}` : `/plants/v2/plantbatch/packages`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records additive usage details applied to specific plants at a Facility.
   * Permissions Required:
   * - Manage Plants Additives
   * POST CreatePlantsAdditives
   */
  public async createPlantsAdditives<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/additives?${queryStr}` : `/plants/v2/additives`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records additive usage for plants using predefined additive templates at a specified Facility.
   * Permissions Required:
   * - Manage Plants Additives
   * POST CreatePlantsAdditivesUsingTemplate
   */
  public async createPlantsAdditivesUsingTemplate<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/additives/usingtemplate?${queryStr}` : `/plants/v2/additives/usingtemplate`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
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
  public async createPlantsPlantings<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/plantings?${queryStr}` : `/plants/v2/plantings`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records waste events for plants at a Facility, including method, reason, and location details.
   * Permissions Required:
   * - Manage Plants Waste
   * POST CreatePlantsWaste
   */
  public async createPlantsWaste<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/waste?${queryStr}` : `/plants/v2/waste`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Removes plants from a Facility’s inventory while recording the reason for their disposal.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Destroy Veg/Flower Plants
   * DELETE DeletePlants
   */
  public async deletePlants<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2?${queryStr}` : `/plants/v2`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves additive records applied to plants at a specified Facility.
   * Permissions Required:
   * - View/Manage Plants Additives
   * GET GetAdditives
   */
  public async getPlantsAdditives<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/additives?${queryStr}` : `/plants/v2/additives`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of all plant additive types defined within a Facility.
   * GET GetAdditivesTypes
   */
  public async getPlantsAdditivesTypes<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/additives/types?${queryStr}` : `/plants/v2/additives/types`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
   * Permissions Required:
   * - View Veg/Flower Plants
   * GET GetFloweringPlants
   */
  public async getFloweringPlants<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/flowering?${queryStr}` : `/plants/v2/flowering`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves the list of growth phases supported by a specified Facility.
   * GET GetGrowthPhases
   */
  public async getPlantsGrowthPhases<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/growthphases?${queryStr}` : `/plants/v2/growthphases`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves inactive plants at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * GET GetInactivePlants
   */
  public async getInactivePlants<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/inactive?${queryStr}` : `/plants/v2/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves inactive mother-phase plants at a specified Facility.
   * Permissions Required:
   * - View Mother Plants
   * GET GetMotherInactivePlants
   */
  public async getMotherInactivePlants<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/mother/inactive?${queryStr}` : `/plants/v2/mother/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves mother-phase plants currently marked as on hold at a specified Facility.
   * Permissions Required:
   * - View Mother Plants
   * GET GetMotherOnHoldPlants
   */
  public async getMotherOnHoldPlants<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/mother/onhold?${queryStr}` : `/plants/v2/mother/onhold`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves mother-phase plants at a specified Facility.
   * Permissions Required:
   * - View Mother Plants
   * GET GetMotherPlants
   */
  public async getMotherPlants<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/mother?${queryStr}` : `/plants/v2/mother`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves plants that are currently on hold at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * GET GetOnHoldPlants
   */
  public async getOnHoldPlants<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/onhold?${queryStr}` : `/plants/v2/onhold`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a Plant by Id.
   * Permissions Required:
   * - View Veg/Flower Plants
   * GET GetPlantsById
   */
  public async getPlantsById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/${encodeURIComponent(id)}?${queryStr}` : `/plants/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a Plant by label.
   * Permissions Required:
   * - View Veg/Flower Plants
   * GET GetPlantsByLabel
   */
  public async getPlantsByLabel<TResponse = any, TBody = any>(label: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/${encodeURIComponent(label)}?${queryStr}` : `/plants/v2/${encodeURIComponent(label)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of recorded plant waste events for a specific Facility.
   * Permissions Required:
   * - View Plants Waste
   * GET GetPlantsWaste
   */
  public async getPlantsWaste<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/waste?${queryStr}` : `/plants/v2/waste`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of all available plant waste methods for use within a Facility.
   * GET GetPlantsWasteMethods
   */
  public async getPlantsWasteMethods<TResponse = any, TBody = any>(body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/waste/methods/all?${queryStr}` : `/plants/v2/waste/methods/all`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retriveves available reasons for recording mature plant waste at a specified Facility.
   * GET GetPlantsWasteReasons
   */
  public async getPlantsWasteReasons<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/waste/reasons?${queryStr}` : `/plants/v2/waste/reasons`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
   * Permissions Required:
   * - View Veg/Flower Plants
   * GET GetVegetativePlants
   */
  public async getVegetativePlants<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/vegetative?${queryStr}` : `/plants/v2/vegetative`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
   * Permissions Required:
   * - View Plants Waste
   * GET GetWasteById
   */
  public async getPlantsWasteById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/waste/${encodeURIComponent(id)}/plant?${queryStr}` : `/plants/v2/waste/${encodeURIComponent(id)}/plant`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of package records linked to the specified plantWasteId for a given facility.
   * Permissions Required:
   * - View Plants Waste
   * GET GetWastePackageById
   */
  public async getPlantsWastePackageById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/waste/${encodeURIComponent(id)}/package?${queryStr}` : `/plants/v2/waste/${encodeURIComponent(id)}/package`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Adjusts the recorded count of plants at a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdateAdjustPlants
   */
  public async updateAdjustPlants<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/adjust?${queryStr}` : `/plants/v2/adjust`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Changes the growth phases of plants within a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdateGrowthPhase
   */
  public async updatePlantsGrowthPhase<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/growthphase?${queryStr}` : `/plants/v2/growthphase`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manicure/Harvest Veg/Flower Plants
   * PUT UpdateHarvest
   */
  public async updatePlantsHarvest<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/harvest?${queryStr}` : `/plants/v2/harvest`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Merges multiple plant groups into a single group within a Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manicure/Harvest Veg/Flower Plants
   * PUT UpdateMerge
   */
  public async updatePlantsMerge<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/merge?${queryStr}` : `/plants/v2/merge`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Moves plant batches to new locations within a specified Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdatePlantsLocation
   */
  public async updatePlantsLocation<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/location?${queryStr}` : `/plants/v2/location`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates the strain information for plants within a Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdatePlantsStrain
   */
  public async updatePlantsStrain<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/strain?${queryStr}` : `/plants/v2/strain`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Replaces existing plant tags with new tags for plants within a Facility.
   * Permissions Required:
   * - View Veg/Flower Plants
   * - Manage Veg/Flower Plants Inventory
   * PUT UpdatePlantsTag
   */
  public async updatePlantsTag<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/tag?${queryStr}` : `/plants/v2/tag`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Splits an existing plant group into multiple groups within a Facility.
   * Permissions Required:
   * - View Plant
   * PUT UpdateSplit
   */
  public async updatePlantsSplit<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/split?${queryStr}` : `/plants/v2/split`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
   * Permissions Required:
   * - Manage Processing Job
   * POST CreateAdjustProcessingJob
   */
  public async createAdjustProcessingJob<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/adjust?${queryStr}` : `/processing/v2/adjust`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
   * Permissions Required:
   * - Manage Processing Job
   * POST CreateJobTypes
   */
  public async createProcessingJobJobTypes<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes?${queryStr}` : `/processing/v2/jobtypes`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates packages from processing jobs at a Facility, including optional location and note assignments.
   * Permissions Required:
   * - Manage Processing Job
   * POST CreateProcessingJobPackages
   */
  public async createProcessingJobPackages<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/createpackages?${queryStr}` : `/processing/v2/createpackages`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Archives a Processing Job Type at a Facility, making it inactive for future use.
   * Permissions Required:
   * - Manage Processing Job
   * DELETE DeleteJobTypeById
   */
  public async deleteProcessingJobJobTypeById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes/${encodeURIComponent(id)}?${queryStr}` : `/processing/v2/jobtypes/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
   * Permissions Required:
   * - Manage Processing Job
   * DELETE DeleteProcessingJobById
   */
  public async deleteProcessingJobById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/${encodeURIComponent(id)}?${queryStr}` : `/processing/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Completes processing jobs at a Facility by recording final notes and waste measurements.
   * Permissions Required:
   * - Manage Processing Job
   * PUT FinishProcessingJob
   */
  public async finishProcessingJobProcessingJob<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/finish?${queryStr}` : `/processing/v2/finish`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Retrieves a list of all active processing job types defined within a Facility.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetActiveJobTypes
   */
  public async getProcessingJobActiveJobTypes<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes/active?${queryStr}` : `/processing/v2/jobtypes/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves active processing jobs at a specified Facility.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetActiveProcessingJob
   */
  public async getActiveProcessingJob<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/active?${queryStr}` : `/processing/v2/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of all inactive processing job types defined within a Facility.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetInactiveJobTypes
   */
  public async getProcessingJobInactiveJobTypes<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes/inactive?${queryStr}` : `/processing/v2/jobtypes/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves inactive processing jobs at a specified Facility.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetInactiveProcessingJob
   */
  public async getInactiveProcessingJob<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/inactive?${queryStr}` : `/processing/v2/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves all processing job attributes available for a Facility.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetJobTypesAttributes
   */
  public async getProcessingJobJobTypesAttributes<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes/attributes?${queryStr}` : `/processing/v2/jobtypes/attributes`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves all processing job categories available for a specified Facility.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetJobTypesCategories
   */
  public async getProcessingJobJobTypesCategories<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes/categories?${queryStr}` : `/processing/v2/jobtypes/categories`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a ProcessingJob by Id.
   * Permissions Required:
   * - Manage Processing Job
   * GET GetProcessingJobById
   */
  public async getProcessingJobById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/${encodeURIComponent(id)}?${queryStr}` : `/processing/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Initiates new processing jobs at a Facility, including job details and associated packages.
   * Permissions Required:
   * - Manage Processing Job
   * POST StartProcessingJob
   */
  public async startProcessingJobProcessingJob<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/start?${queryStr}` : `/processing/v2/start`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
   * Permissions Required:
   * - Manage Processing Job
   * PUT UnfinishProcessingJob
   */
  public async unfinishProcessingJobProcessingJob<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/unfinish?${queryStr}` : `/processing/v2/unfinish`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
   * Permissions Required:
   * - Manage Processing Job
   * PUT UpdateJobTypes
   */
  public async updateProcessingJobJobTypes<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes?${queryStr}` : `/processing/v2/jobtypes`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Facilitate association of QR codes and Package labels. This will return the count of packages and QR codes associated that were added or replaced.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
   * - WebApi Retail ID Read Write State (All or WriteOnly)
   * - Industry/View Packages
   * - One of the following: Industry/Facility Type/Can Receive Associate Product Label, Licensee/Receive Associate Product Label or Admin/Employees/Packages Page/Product Labels(Manage)
   * POST CreateAssociate
   */
  public async createRetailIdAssociate<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/retailid/v2/associate?${queryStr}` : `/retailid/v2/associate`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
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
  public async createRetailIdGenerate<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/retailid/v2/generate?${queryStr}` : `/retailid/v2/generate`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Merge and adjust one source to one target Package. First Package detected will be processed as target Package. This requires an action reason with name containing the 'Merge' word and setup with 'Package adjustment' area.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Retail ID(Write)
   * - WebApi Retail ID Read Write State (All or WriteOnly)
   * - Key Value Settings/Retail ID Merge Packages Enabled
   * POST CreateMerge
   */
  public async createRetailIdMerge<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/retailid/v2/merge?${queryStr}` : `/retailid/v2/merge`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
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
  public async createRetailIdPackagesInfo<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/retailid/v2/packages/info?${queryStr}` : `/retailid/v2/packages/info`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Retrieves the available Retail Item ID quota for a facility.
   * Permissions Required:
   * - Download Product Labels
   * - Manage Product Labels
   * - Manage Tag Orders
   * GET GetAllotment
   */
  public async getRetailIdAllotment<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/retailid/v2/allotment?${queryStr}` : `/retailid/v2/allotment`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
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
  public async getRetailIdReceiveByLabel<TResponse = any, TBody = any>(label: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/retailid/v2/receive/${encodeURIComponent(label)}?${queryStr}` : `/retailid/v2/receive/${encodeURIComponent(label)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
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
  public async getRetailIdReceiveQrByShortCode<TResponse = any, TBody = any>(shortCode: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/retailid/v2/receive/qr/${encodeURIComponent(shortCode)}?${queryStr}` : `/retailid/v2/receive/qr/${encodeURIComponent(shortCode)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Records new sales delivery entries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   * - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
   * - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
   * POST CreateDeliveries
   */
  public async createSalesDeliveries<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries?${queryStr}` : `/sales/v2/deliveries`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Processes the departure of retailer deliveries for a Facility using the provided License Number and delivery data.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   * - Industry/Facility Type/Retailer Delivery
   * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   * - Manage Retailer Delivery
   * POST CreateDeliveriesRetailerDepart
   */
  public async createSalesDeliveriesRetailerDepart<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/depart?${queryStr}` : `/sales/v2/deliveries/retailer/depart`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Ends retailer delivery records for a given License Number. Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   * - Industry/Facility Type/Retailer Delivery
   * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   * - Manage Retailer Delivery
   * POST CreateDeliveriesRetailerEnd
   */
  public async createSalesDeliveriesRetailerEnd<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/end?${queryStr}` : `/sales/v2/deliveries/retailer/end`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records restock deliveries for retailer facilities using the provided License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   * - Industry/Facility Type/Retailer Delivery
   * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   * - Manage Retailer Delivery
   * POST CreateDeliveriesRetailerRestock
   */
  public async createSalesDeliveriesRetailerRestock<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/restock?${queryStr}` : `/sales/v2/deliveries/retailer/restock`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records a list of sales deliveries for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Sales (Write)
   * - Industry/Facility Type/Consumer Sales or Industry/Facility Type/Patient Sales or Industry/Facility Type/External Patient Sales or Industry/Facility Type/Caregiver Sales
   * - Industry/Facility Type/Advanced Sales
   * - WebApi Sales Read Write State (All or WriteOnly)
   * - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
   * - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
   * POST CreateReceipts
   */
  public async createSalesReceipts<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts?${queryStr}` : `/sales/v2/receipts`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Records retailer delivery data for a given License Number, including delivery destinations. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   * - Industry/Facility Type/Retailer Delivery
   * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   * - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
   * - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
   * - Manage Retailer Delivery
   * POST CreateSalesDeliveriesRetailer
   */
  public async createSalesDeliveriesRetailer<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer?${queryStr}` : `/sales/v2/deliveries/retailer`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Voids a sales delivery for a Facility using the provided License Number and delivery Id.
   * Permissions Required:
   * - Manage Sales Delivery
   * DELETE DeleteDeliveryById
   */
  public async deleteSalesDeliveryById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/deliveries/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Voids a retailer delivery for a Facility using the provided License Number and delivery Id.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   * - Industry/Facility Type/Retailer Delivery
   * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   * - Manage Retailer Delivery
   * DELETE DeleteDeliveryRetailerById
   */
  public async deleteSalesDeliveryRetailerById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/deliveries/retailer/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Archives a sales receipt for a Facility using the provided License Number and receipt Id.
   * Permissions Required:
   * - Manage Sales
   * DELETE DeleteReceiptById
   */
  public async deleteSalesReceiptById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/receipts/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
   * Permissions Required:
   * - View Sales Delivery
   * - Manage Sales Delivery
   * GET GetActiveDeliveries
   */
  public async getSalesActiveDeliveries<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/active?${queryStr}` : `/sales/v2/deliveries/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
   * Permissions Required:
   * - View Retailer Delivery
   * - Manage Retailer Delivery
   * GET GetActiveDeliveriesRetailer
   */
  public async getSalesActiveDeliveriesRetailer<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/active?${queryStr}` : `/sales/v2/deliveries/retailer/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
   * Permissions Required:
   * - View Sales
   * - Manage Sales
   * GET GetActiveReceipts
   */
  public async getSalesActiveReceipts<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/active?${queryStr}` : `/sales/v2/receipts/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of counties available for sales deliveries.
   * GET GetCounties
   */
  public async getSalesCounties<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/counties?${queryStr}` : `/sales/v2/counties`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of customer types.
   * GET GetCustomerTypes
   */
  public async getSalesCustomerTypes<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/customertypes?${queryStr}` : `/sales/v2/customertypes`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of return reasons for sales deliveries based on the provided License Number.
   * Permissions Required:
   * - Sales Delivery
   * GET GetDeliveriesReturnReasons
   */
  public async getSalesDeliveriesReturnReasons<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/returnreasons?${queryStr}` : `/sales/v2/deliveries/returnreasons`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a retailer delivery record by its ID, with an optional License Number.
   * Permissions Required:
   * - View Retailer Delivery
   * - Manage Retailer Delivery
   * GET GetDeliveryRetailerById
   */
  public async getSalesDeliveryRetailerById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/deliveries/retailer/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
   * Permissions Required:
   * - View Sales Delivery
   * - Manage Sales Delivery
   * GET GetInactiveDeliveries
   */
  public async getSalesInactiveDeliveries<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/inactive?${queryStr}` : `/sales/v2/deliveries/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
   * Permissions Required:
   * - View Retailer Delivery
   * - Manage Retailer Delivery
   * GET GetInactiveDeliveriesRetailer
   */
  public async getSalesInactiveDeliveriesRetailer<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/inactive?${queryStr}` : `/sales/v2/deliveries/retailer/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
   * Permissions Required:
   * - View Sales
   * - Manage Sales
   * GET GetInactiveReceipts
   */
  public async getSalesInactiveReceipts<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/inactive?${queryStr}` : `/sales/v2/receipts/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of valid Patient registration locations for sales.
   * GET GetPatientRegistrationLocations
   */
  public async getSalesPatientRegistrationLocations<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/patientregistration/locations?${queryStr}` : `/sales/v2/patientregistration/locations`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of available payment types for the specified License Number.
   * Permissions Required:
   * - View Sales Delivery
   * - Manage Sales Delivery
   * GET GetPaymentTypes
   */
  public async getSalesPaymentTypes<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/paymenttypes?${queryStr}` : `/sales/v2/paymenttypes`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a sales receipt by its Id, with an optional License Number.
   * Permissions Required:
   * - View Sales
   * - Manage Sales
   * GET GetReceiptById
   */
  public async getSalesReceiptById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/receipts/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a Sales Receipt by its external number, with an optional License Number.
   * Permissions Required:
   * - View Sales
   * - Manage Sales
   * GET GetReceiptsExternalByExternalNumber
   */
  public async getSalesReceiptsExternalByExternalNumber<TResponse = any, TBody = any>(externalNumber: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/external/${encodeURIComponent(externalNumber)}?${queryStr}` : `/sales/v2/receipts/external/${encodeURIComponent(externalNumber)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a sales delivery record by its Id, with an optional License Number.
   * Permissions Required:
   * - View Sales Delivery
   * - Manage Sales Delivery
   * GET GetSalesDeliveryById
   */
  public async getSalesDeliveryById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/deliveries/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * Permissions Required:
   * - Manage Sales Delivery
   * PUT UpdateDeliveries
   */
  public async updateSalesDeliveries<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries?${queryStr}` : `/sales/v2/deliveries`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
   * Permissions Required:
   * - Manage Sales Delivery
   * PUT UpdateDeliveriesComplete
   */
  public async updateSalesDeliveriesComplete<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/complete?${queryStr}` : `/sales/v2/deliveries/complete`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * Permissions Required:
   * - Manage Sales Delivery, Manage Sales Delivery Hub
   * PUT UpdateDeliveriesHub
   */
  public async updateSalesDeliveriesHub<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/hub?${queryStr}` : `/sales/v2/deliveries/hub`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
   * Permissions Required:
   * - Manage Sales Delivery Hub
   * PUT UpdateDeliveriesHubAccept
   */
  public async updateSalesDeliveriesHubAccept<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/hub/accept?${queryStr}` : `/sales/v2/deliveries/hub/accept`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
   * Permissions Required:
   * - Manage Sales Delivery Hub
   * PUT UpdateDeliveriesHubDepart
   */
  public async updateSalesDeliveriesHubDepart<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/hub/depart?${queryStr}` : `/sales/v2/deliveries/hub/depart`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
   * Permissions Required:
   * - Manage Sales Delivery Hub
   * PUT UpdateDeliveriesHubVerifyID
   */
  public async updateSalesDeliveriesHubVerifyID<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/hub/verifyID?${queryStr}` : `/sales/v2/deliveries/hub/verifyID`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates retailer delivery records for a given License Number. Please note: The DateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * Permissions Required:
   * - External Sources(ThirdPartyVendorV2)/Sales Deliveries(Write)
   * - Industry/Facility Type/Retailer Delivery
   * - Industry/Facility Type/Consumer Sales Delivery or Industry/Facility Type/Patient Sales Delivery
   * - WebApi Sales Deliveries Read Write State (All or WriteOnly)
   * - WebApi Retail ID Read Write State (All or WriteOnly) - Required for RID only.
   * - External Sources(ThirdPartyVendorV2)/Retail ID(Write) - Required for RID only.
   * - Manage Retailer Delivery
   * PUT UpdateDeliveriesRetailer
   */
  public async updateSalesDeliveriesRetailer<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer?${queryStr}` : `/sales/v2/deliveries/retailer`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * Permissions Required:
   * - Manage Sales
   * PUT UpdateReceipts
   */
  public async updateSalesReceipts<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts?${queryStr}` : `/sales/v2/receipts`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
   * Permissions Required:
   * - Manage Sales
   * PUT UpdateReceiptsFinalize
   */
  public async updateSalesReceiptsFinalize<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/finalize?${queryStr}` : `/sales/v2/receipts/finalize`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
   * Permissions Required:
   * - Manage Sales
   * PUT UpdateReceiptsUnfinalize
   */
  public async updateSalesReceiptsUnfinalize<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/unfinalize?${queryStr}` : `/sales/v2/receipts/unfinalize`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
   * POST CreateIntegratorSetup
   */
  public async createSandboxIntegratorSetup<TResponse = any, TBody = any>(body?: TBody, userKey?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (userKey !== undefined) query.append('userKey', String(userKey));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sandbox/v2/integrator/setup?${queryStr}` : `/sandbox/v2/integrator/setup`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates new strain records for a specified Facility.
   * Permissions Required:
   * - Manage Strains
   * POST CreateStrains
   */
  public async createStrains<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v2?${queryStr}` : `/strains/v2`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Archives an existing strain record for a Facility
   * Permissions Required:
   * - Manage Strains
   * DELETE DeleteStrainsById
   */
  public async deleteStrainsById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v2/${encodeURIComponent(id)}?${queryStr}` : `/strains/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
   * Permissions Required:
   * - Manage Strains
   * GET GetActiveStrains
   */
  public async getActiveStrains<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v2/active?${queryStr}` : `/strains/v2/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
   * Permissions Required:
   * - Manage Strains
   * GET GetInactiveStrains
   */
  public async getInactiveStrains<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v2/inactive?${queryStr}` : `/strains/v2/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a Strain record by its Id, with an optional license number.
   * Permissions Required:
   * - Manage Strains
   * GET GetStrainsById
   */
  public async getStrainsById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v2/${encodeURIComponent(id)}?${queryStr}` : `/strains/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Updates existing strain records for a specified Facility.
   * Permissions Required:
   * - Manage Strains
   * PUT UpdateStrains
   */
  public async updateStrains<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v2?${queryStr}` : `/strains/v2`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates new sublocation records for a Facility.
   * Permissions Required:
   * - Manage Locations
   * POST CreateSublocations
   */
  public async createSublocations<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sublocations/v2?${queryStr}` : `/sublocations/v2`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Archives an existing Sublocation record for a Facility.
   * Permissions Required:
   * - Manage Locations
   * DELETE DeleteSublocationsById
   */
  public async deleteSublocationsById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sublocations/v2/${encodeURIComponent(id)}?${queryStr}` : `/sublocations/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
   * Permissions Required:
   * - Manage Locations
   * GET GetActiveSublocations
   */
  public async getActiveSublocations<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sublocations/v2/active?${queryStr}` : `/sublocations/v2/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of inactive sublocations for the specified Facility.
   * Permissions Required:
   * - Manage Locations
   * GET GetInactiveSublocations
   */
  public async getInactiveSublocations<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sublocations/v2/inactive?${queryStr}` : `/sublocations/v2/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a Sublocation by its Id, with an optional license number.
   * Permissions Required:
   * - Manage Locations
   * GET GetSublocationsById
   */
  public async getSublocationsById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sublocations/v2/${encodeURIComponent(id)}?${queryStr}` : `/sublocations/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Updates existing sublocation records for a specified Facility.
   * Permissions Required:
   * - Manage Locations
   * PUT UpdateSublocations
   */
  public async updateSublocations<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sublocations/v2?${queryStr}` : `/sublocations/v2`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Returns a list of available package tags. NOTE: This is a premium endpoint.
   * Permissions Required:
   * - Manage Tags
   * GET GetAvailablePackage
   */
  public async getTagsAvailablePackage<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/tags/v2/package/available?${queryStr}` : `/tags/v2/package/available`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of available plant tags. NOTE: This is a premium endpoint.
   * Permissions Required:
   * - Manage Tags
   * GET GetAvailablePlant
   */
  public async getTagsAvailablePlant<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/tags/v2/plant/available?${queryStr}` : `/tags/v2/plant/available`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of staged tags. NOTE: This is a premium endpoint.
   * Permissions Required:
   * - Manage Tags
   * - RetailId.AllowPackageStaging Key Value enabled
   * GET GetStagedTags
   */
  public async getStagedTags<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/tags/v2/staged?${queryStr}` : `/tags/v2/staged`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Arrive a transfer for a Facility.
   * Permissions Required:
   * - Manage Transfer Hub
   * POST CreateHubArrive
   */
  public async createTransfersHubArrive<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/hub/arrive?${queryStr}` : `/transfers/v2/hub/arrive`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * CheckIn a transfer for a Facility.
   * Permissions Required:
   * - Manage Transfer Hub
   * POST CreateHubCheckin
   */
  public async createTransfersHubCheckin<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/hub/checkin?${queryStr}` : `/transfers/v2/hub/checkin`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * CheckOut a transfer for a Facility.
   * Permissions Required:
   * - Manage Transfer Hub
   * POST CreateHubCheckout
   */
  public async createTransfersHubCheckout<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/hub/checkout?${queryStr}` : `/transfers/v2/hub/checkout`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Depart a transfer for a Facility.
   * Permissions Required:
   * - Manage Transfer Hub
   * POST CreateHubDepart
   */
  public async createTransfersHubDepart<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/hub/depart?${queryStr}` : `/transfers/v2/hub/depart`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates external incoming shipment plans for a Facility.
   * Permissions Required:
   * - Manage Transfers
   * POST CreateIncomingExternal
   */
  public async createTransfersIncomingExternal<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/external/incoming?${queryStr}` : `/transfers/v2/external/incoming`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates new transfer templates for a Facility.
   * Permissions Required:
   * - Manage Transfer Templates
   * POST CreateOutgoingTemplates
   */
  public async createTransfersOutgoingTemplates<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing?${queryStr}` : `/transfers/v2/templates/outgoing`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Voids an external incoming shipment plan for a Facility.
   * Permissions Required:
   * - Manage Transfers
   * DELETE DeleteIncomingExternalById
   */
  public async deleteTransfersIncomingExternalById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/external/incoming/${encodeURIComponent(id)}?${queryStr}` : `/transfers/v2/external/incoming/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Archives a transfer template for a Facility.
   * Permissions Required:
   * - Manage Transfer Templates
   * DELETE DeleteOutgoingTemplateById
   */
  public async deleteTransfersOutgoingTemplateById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/${encodeURIComponent(id)}?${queryStr}` : `/transfers/v2/templates/outgoing/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Returns a list of available shipment Package states.
   * GET GetDeliveriesPackagesStates
   */
  public async getTransfersDeliveriesPackagesStates<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/deliveries/packages/states?${queryStr}` : `/transfers/v2/deliveries/packages/states`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetDeliveryPackageById
   */
  public async getTransfersDeliveryPackageById<TResponse = any, TBody = any>(id: string, body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/deliveries/${encodeURIComponent(id)}/packages?${queryStr}` : `/transfers/v2/deliveries/${encodeURIComponent(id)}/packages`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetDeliveryPackageRequiredlabtestbatchById
   */
  public async getTransfersDeliveryPackageRequiredlabtestbatchById<TResponse = any, TBody = any>(id: string, body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/deliveries/package/${encodeURIComponent(id)}/requiredlabtestbatches?${queryStr}` : `/transfers/v2/deliveries/package/${encodeURIComponent(id)}/requiredlabtestbatches`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetDeliveryPackageWholesaleById
   */
  public async getTransfersDeliveryPackageWholesaleById<TResponse = any, TBody = any>(id: string, body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/deliveries/${encodeURIComponent(id)}/packages/wholesale?${queryStr}` : `/transfers/v2/deliveries/${encodeURIComponent(id)}/packages/wholesale`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetDeliveryTransporterById
   */
  public async getTransfersDeliveryTransporterById<TResponse = any, TBody = any>(id: string, body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/deliveries/${encodeURIComponent(id)}/transporters?${queryStr}` : `/transfers/v2/deliveries/${encodeURIComponent(id)}/transporters`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetDeliveryTransporterDetailById
   */
  public async getTransfersDeliveryTransporterDetailById<TResponse = any, TBody = any>(id: string, body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/deliveries/${encodeURIComponent(id)}/transporters/details?${queryStr}` : `/transfers/v2/deliveries/${encodeURIComponent(id)}/transporters/details`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetHub
   */
  public async getTransfersHub<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/hub?${queryStr}` : `/transfers/v2/hub`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetIncomingTransfers
   */
  public async getIncomingTransfers<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/incoming?${queryStr}` : `/transfers/v2/incoming`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Get Transfer Manifest HTML for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetManifestHtmlById
   */
  public async getTransfersManifestHtmlById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/manifest/${encodeURIComponent(id)}/html?${queryStr}` : `/transfers/v2/manifest/${encodeURIComponent(id)}/html`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Get Transfer Manifest PDF for a given Transfer Id
   * Permissions Required:
   * - Manage Transfer Templates
   * - View Transfer Templates
   * GET GetManifestPdfById
   */
  public async getTransfersManifestPdfById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/manifest/${encodeURIComponent(id)}/pdf?${queryStr}` : `/transfers/v2/manifest/${encodeURIComponent(id)}/pdf`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of deliveries associated with a specific transfer template.
   * Permissions Required:
   * - Manage Transfer Templates
   * - View Transfer Templates
   * GET GetOutgoingTemplateDeliveryById
   */
  public async getTransfersOutgoingTemplateDeliveryById<TResponse = any, TBody = any>(id: string, body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/${encodeURIComponent(id)}/deliveries?${queryStr}` : `/transfers/v2/templates/outgoing/${encodeURIComponent(id)}/deliveries`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfer Templates
   * - View Transfer Templates
   * GET GetOutgoingTemplateDeliveryPackageById
   */
  public async getTransfersOutgoingTemplateDeliveryPackageById<TResponse = any, TBody = any>(id: string, body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/packages?${queryStr}` : `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/packages`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfer Templates
   * - View Transfer Templates
   * GET GetOutgoingTemplateDeliveryTransporterById
   */
  public async getTransfersOutgoingTemplateDeliveryTransporterById<TResponse = any, TBody = any>(id: string, body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/transporters?${queryStr}` : `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/transporters`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfer Templates
   * - View Transfer Templates
   * GET GetOutgoingTemplateDeliveryTransporterDetailById
   */
  public async getTransfersOutgoingTemplateDeliveryTransporterDetailById<TResponse = any, TBody = any>(id: string, body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/transporters/details?${queryStr}` : `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/transporters/details`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
   * Permissions Required:
   * - Manage Transfer Templates
   * - View Transfer Templates
   * GET GetOutgoingTemplates
   */
  public async getTransfersOutgoingTemplates<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing?${queryStr}` : `/transfers/v2/templates/outgoing`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetOutgoingTransfers
   */
  public async getOutgoingTransfers<TResponse = any, TBody = any>(body?: TBody, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (lastModifiedEnd !== undefined) query.append('lastModifiedEnd', String(lastModifiedEnd));
    if (lastModifiedStart !== undefined) query.append('lastModifiedStart', String(lastModifiedStart));
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/outgoing?${queryStr}` : `/transfers/v2/outgoing`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of shipments with rejected packages for a Facility.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetRejectedTransfers
   */
  public async getRejectedTransfers<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/rejected?${queryStr}` : `/transfers/v2/rejected`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetTransfersDeliveryById
   */
  public async getTransfersDeliveryById<TResponse = any, TBody = any>(id: string, body?: TBody, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/${encodeURIComponent(id)}/deliveries?${queryStr}` : `/transfers/v2/${encodeURIComponent(id)}/deliveries`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of available transfer types for a Facility based on its license number.
   * GET GetTransfersTypes
   */
  public async getTransfersTypes<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/types?${queryStr}` : `/transfers/v2/types`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Updates external incoming shipment plans for a Facility.
   * Permissions Required:
   * - Manage Transfers
   * PUT UpdateIncomingExternal
   */
  public async updateTransfersIncomingExternal<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/external/incoming?${queryStr}` : `/transfers/v2/external/incoming`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates existing transfer templates for a Facility.
   * Permissions Required:
   * - Manage Transfer Templates
   * PUT UpdateOutgoingTemplates
   */
  public async updateTransfersOutgoingTemplates<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing?${queryStr}` : `/transfers/v2/templates/outgoing`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates new driver records for a Facility.
   * Permissions Required:
   * - Manage Transporters
   * POST CreateDrivers
   */
  public async createTransportersDrivers<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/drivers?${queryStr}` : `/transporters/v2/drivers`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Creates new vehicle records for a Facility.
   * Permissions Required:
   * - Manage Transporters
   * POST CreateVehicles
   */
  public async createTransportersVehicles<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/vehicles?${queryStr}` : `/transporters/v2/vehicles`;
    const { data } = await this.client.post<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
   * Permissions Required:
   * - Manage Transporters
   * DELETE DeleteDriverById
   */
  public async deleteTransportersDriverById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/drivers/${encodeURIComponent(id)}?${queryStr}` : `/transporters/v2/drivers/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
   * Permissions Required:
   * - Manage Transporters
   * DELETE DeleteVehicleById
   */
  public async deleteTransportersVehicleById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/vehicles/${encodeURIComponent(id)}?${queryStr}` : `/transporters/v2/vehicles/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
   * Permissions Required:
   * - Transporters
   * GET GetDriverById
   */
  public async getTransportersDriverById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/drivers/${encodeURIComponent(id)}?${queryStr}` : `/transporters/v2/drivers/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of drivers for a Facility.
   * Permissions Required:
   * - Transporters
   * GET GetDrivers
   */
  public async getTransportersDrivers<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/drivers?${queryStr}` : `/transporters/v2/drivers`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
   * Permissions Required:
   * - Transporters
   * GET GetVehicleById
   */
  public async getTransportersVehicleById<TResponse = any, TBody = any>(id: string, body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/vehicles/${encodeURIComponent(id)}?${queryStr}` : `/transporters/v2/vehicles/${encodeURIComponent(id)}`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves a list of vehicles for a Facility.
   * Permissions Required:
   * - Transporters
   * GET GetVehicles
   */
  public async getTransportersVehicles<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    if (pageNumber !== undefined) query.append('pageNumber', String(pageNumber));
    if (pageSize !== undefined) query.append('pageSize', String(pageSize));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/vehicles?${queryStr}` : `/transporters/v2/vehicles`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Updates existing driver records for a Facility.
   * Permissions Required:
   * - Manage Transporters
   * PUT UpdateDrivers
   */
  public async updateTransportersDrivers<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/drivers?${queryStr}` : `/transporters/v2/drivers`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Updates existing vehicle records for a facility.
   * Permissions Required:
   * - Manage Transporters
   * PUT UpdateVehicles
   */
  public async updateTransportersVehicles<TResponse = any, TBody = any>(body?: TBody, licenseNumber?: string): Promise<TResponse> {
    const query = new URLSearchParams();
    if (licenseNumber !== undefined) query.append('licenseNumber', String(licenseNumber));
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/vehicles?${queryStr}` : `/transporters/v2/vehicles`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * Retrieves all active units of measure.
   * GET GetActiveUnitsOfMeasure
   */
  public async getActiveUnitsOfMeasure<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/unitsofmeasure/v2/active?${queryStr}` : `/unitsofmeasure/v2/active`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves all inactive units of measure.
   * GET GetInactiveUnitsOfMeasure
   */
  public async getInactiveUnitsOfMeasure<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/unitsofmeasure/v2/inactive?${queryStr}` : `/unitsofmeasure/v2/inactive`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * Retrieves all available waste methods.
   * GET GetWasteMethodsWasteMethods
   */
  public async getWasteMethodsWasteMethods<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/wastemethods/v2?${queryStr}` : `/wastemethods/v2`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * 
   * DELETE DeleteWebhooksBySubscriptionId
   */
  public async deleteWebhooksBySubscriptionId<TResponse = any, TBody = any>(subscriptionId: string, body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/webhooks/v2/${encodeURIComponent(subscriptionId)}?${queryStr}` : `/webhooks/v2/${encodeURIComponent(subscriptionId)}`;
    const { data } = await this.client.delete<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * 
   * GET GetWebhooks
   */
  public async getWebhooks<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/webhooks/v2?${queryStr}` : `/webhooks/v2`;
    const { data } = await this.client.get<TResponse>(fullUrl, { data: body });
    return data;
  }
  /**
   * 
   * PUT UpdateDisableBySubscriptionId
   */
  public async updateWebhooksDisableBySubscriptionId<TResponse = any, TBody = any>(subscriptionId: string, body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/webhooks/v2/disable/${encodeURIComponent(subscriptionId)}?${queryStr}` : `/webhooks/v2/disable/${encodeURIComponent(subscriptionId)}`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * 
   * PUT UpdateEnableBySubscriptionId
   */
  public async updateWebhooksEnableBySubscriptionId<TResponse = any, TBody = any>(subscriptionId: string, body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/webhooks/v2/enable/${encodeURIComponent(subscriptionId)}?${queryStr}` : `/webhooks/v2/enable/${encodeURIComponent(subscriptionId)}`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
  /**
   * 
   * PUT UpdateWebhooks
   */
  public async updateWebhooks<TResponse = any, TBody = any>(body?: TBody): Promise<TResponse> {
    const query = new URLSearchParams();
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/webhooks/v2?${queryStr}` : `/webhooks/v2`;
    const { data } = await this.client.put<TResponse>(fullUrl, body);
    return data;
  }
}

