
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    CreateDriversRequestItem,
    CreateVehiclesRequestItem,
    TransportersDriver,
    TransportersVehicle,
    UpdateDriversRequestItem,
    UpdateVehiclesRequestItem,
    WriteResponse,
} from '../models';

export class TransportersService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Creates new driver records for a Facility.
   * Permissions Required:
   * - Manage Transporters
   * POST CreateDrivers
   */
  public async createTransportersDrivers(body: CreateDriversRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createTransportersDrivers(body,licenseNumber,
        )
    );
  }

  /**
   * Creates new vehicle records for a Facility.
   * Permissions Required:
   * - Manage Transporters
   * POST CreateVehicles
   */
  public async createTransportersVehicles(body: CreateVehiclesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createTransportersVehicles(body,licenseNumber,
        )
    );
  }

  /**
   * Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
   * Permissions Required:
   * - Manage Transporters
   * DELETE DeleteDriverById
   */
  public async deleteTransportersDriverById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteTransportersDriverById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
   * Permissions Required:
   * - Manage Transporters
   * DELETE DeleteVehicleById
   */
  public async deleteTransportersVehicleById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteTransportersVehicleById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
   * Permissions Required:
   * - Transporters
   * GET GetDriverById
   */
  public async getTransportersDriverById(id: string, body?: any, licenseNumber?: string): Promise<TransportersDriver> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransportersDriverById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of drivers for a Facility.
   * Permissions Required:
   * - Transporters
   * GET GetDrivers
   */
  public async getTransportersDrivers(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<TransportersDriver>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransportersDrivers(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
   * Permissions Required:
   * - Transporters
   * GET GetVehicleById
   */
  public async getTransportersVehicleById(id: string, body?: any, licenseNumber?: string): Promise<TransportersVehicle> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransportersVehicleById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of vehicles for a Facility.
   * Permissions Required:
   * - Transporters
   * GET GetVehicles
   */
  public async getTransportersVehicles(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<TransportersVehicle>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransportersVehicles(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Updates existing driver records for a Facility.
   * Permissions Required:
   * - Manage Transporters
   * PUT UpdateDrivers
   */
  public async updateTransportersDrivers(body: UpdateDriversRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateTransportersDrivers(body,licenseNumber,
        )
    );
  }

  /**
   * Updates existing vehicle records for a facility.
   * Permissions Required:
   * - Manage Transporters
   * PUT UpdateVehicles
   */
  public async updateTransportersVehicles(body: UpdateVehiclesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateTransportersVehicles(body,licenseNumber,
        )
    );
  }
}

