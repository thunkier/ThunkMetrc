
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    Brand,
    Category,
    CreateBrandRequestItem,
    CreateFileRequestItem,
    CreateItemsRequestItem,
    CreatePhotoRequestItem,
    File,
    Item,
    Photo,
    UpdateBrandRequestItem,
    UpdateItemsRequestItem,
    WriteResponse,
} from '../models';

export class ItemsService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Creates one or more new item brands for the specified Facility identified by the License Number.
   * Permissions Required:
   * - Manage Items
   * POST CreateBrand
   */
  public async createItemsBrand(body: CreateBrandRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createItemsBrand(body,licenseNumber,
        )
    );
  }

  /**
   * Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
   * Permissions Required:
   * - Manage Items
   * POST CreateFile
   */
  public async createItemsFile(body: CreateFileRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createItemsFile(body,licenseNumber,
        )
    );
  }

  /**
   * Permissions Required:
   * - Manage Items
   * POST CreateItems
   */
  public async createItems(body: CreateItemsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createItems(body,licenseNumber,
        )
    );
  }

  /**
   * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
   * Permissions Required:
   * - Manage Items
   * POST CreatePhoto
   */
  public async createItemsPhoto(body: CreatePhotoRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createItemsPhoto(body,licenseNumber,
        )
    );
  }

  /**
   * Archives the specified Item Brand by Id for the given Facility License Number.
   * Permissions Required:
   * - Manage Items
   * DELETE DeleteBrandById
   */
  public async deleteItemsBrandById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteItemsBrandById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Archives the specified Product by Id for the given Facility License Number.
   * Permissions Required:
   * - Manage Items
   * DELETE DeleteItemsById
   */
  public async deleteItemsById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteItemsById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Returns a list of active items for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * GET GetActiveItems
   */
  public async getActiveItems(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Item>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getActiveItems(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of active item brands for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * GET GetBrands
   */
  public async getItemsBrands(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Brand>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getItemsBrands(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of item categories.
   * GET GetCategories
   */
  public async getItemsCategories(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Category>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getItemsCategories(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a file by its Id for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * GET GetFileById
   */
  public async getItemsFileById(id: string, body?: any, licenseNumber?: string): Promise<File> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getItemsFileById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of inactive items for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * GET GetInactiveItems
   */
  public async getInactiveItems(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Item>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getInactiveItems(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves detailed information about a specific Item by Id.
   * Permissions Required:
   * - Manage Items
   * GET GetItemsById
   */
  public async getItemsById(id: string, body?: any, licenseNumber?: string): Promise<Item> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getItemsById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves an image by its Id for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * GET GetPhotoById
   */
  public async getItemsPhotoById(id: string, body?: any, licenseNumber?: string): Promise<Photo> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getItemsPhotoById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Updates one or more existing item brands for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * PUT UpdateBrand
   */
  public async updateItemsBrand(body: UpdateBrandRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateItemsBrand(body,licenseNumber,
        )
    );
  }

  /**
   * Updates one or more existing products for the specified Facility.
   * Permissions Required:
   * - Manage Items
   * PUT UpdateItems
   */
  public async updateItems(body: UpdateItemsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateItems(body,licenseNumber,
        )
    );
  }
}

