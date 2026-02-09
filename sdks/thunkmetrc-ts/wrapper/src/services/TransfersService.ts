
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    CreateHubArriveRequestItem,
    CreateHubCheckinRequestItem,
    CreateHubCheckoutRequestItem,
    CreateHubDepartRequestItem,
    CreateIncomingExternalRequestItem,
    CreateOutgoingTemplatesRequestItem,
    DeliveryPackage,
    DeliveryPackageRequiredlabtestbatch,
    DeliveryPackageWholesale,
    DeliveryTransporter,
    DeliveryTransporterDetail,
    Hub,
    ManifestPdf,
    Template,
    TemplateDelivery,
    TemplateDeliveryPackage,
    TemplateDeliveryTransporter,
    TemplateDeliveryTransporterDetail,
    Transfer,
    TransfersDelivery,
    TransfersType,
    UpdateIncomingExternalRequestItem,
    UpdateOutgoingTemplatesRequestItem,
    WriteResponse,
} from '../models';

export class TransfersService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

  /**
   * Arrive a transfer for a Facility.
   * Permissions Required:
   * - Manage Transfer Hub
   * POST CreateHubArrive
   */
  public async createTransfersHubArrive(body: CreateHubArriveRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createTransfersHubArrive(body,licenseNumber,
        )
    );
  }

  /**
   * CheckIn a transfer for a Facility.
   * Permissions Required:
   * - Manage Transfer Hub
   * POST CreateHubCheckin
   */
  public async createTransfersHubCheckin(body: CreateHubCheckinRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createTransfersHubCheckin(body,licenseNumber,
        )
    );
  }

  /**
   * CheckOut a transfer for a Facility.
   * Permissions Required:
   * - Manage Transfer Hub
   * POST CreateHubCheckout
   */
  public async createTransfersHubCheckout(body: CreateHubCheckoutRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createTransfersHubCheckout(body,licenseNumber,
        )
    );
  }

  /**
   * Depart a transfer for a Facility.
   * Permissions Required:
   * - Manage Transfer Hub
   * POST CreateHubDepart
   */
  public async createTransfersHubDepart(body: CreateHubDepartRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createTransfersHubDepart(body,licenseNumber,
        )
    );
  }

  /**
   * Creates external incoming shipment plans for a Facility.
   * Permissions Required:
   * - Manage Transfers
   * POST CreateIncomingExternal
   */
  public async createTransfersIncomingExternal(body: CreateIncomingExternalRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createTransfersIncomingExternal(body,licenseNumber,
        )
    );
  }

  /**
   * Creates new transfer templates for a Facility.
   * Permissions Required:
   * - Manage Transfer Templates
   * POST CreateOutgoingTemplates
   */
  public async createTransfersOutgoingTemplates(body: CreateOutgoingTemplatesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createTransfersOutgoingTemplates(body,licenseNumber,
        )
    );
  }

  /**
   * Voids an external incoming shipment plan for a Facility.
   * Permissions Required:
   * - Manage Transfers
   * DELETE DeleteIncomingExternalById
   */
  public async deleteTransfersIncomingExternalById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteTransfersIncomingExternalById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Archives a transfer template for a Facility.
   * Permissions Required:
   * - Manage Transfer Templates
   * DELETE DeleteOutgoingTemplateById
   */
  public async deleteTransfersOutgoingTemplateById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteTransfersOutgoingTemplateById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Returns a list of available shipment Package states.
   * GET GetDeliveriesPackagesStates
   */
  public async getTransfersDeliveriesPackagesStates(body?: any): Promise<any> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersDeliveriesPackagesStates(body,
        )
    );
  }

  /**
   * Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetDeliveryPackageById
   */
  public async getTransfersDeliveryPackageById(id: string, body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<DeliveryPackage>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersDeliveryPackageById(id, body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetDeliveryPackageRequiredlabtestbatchById
   */
  public async getTransfersDeliveryPackageRequiredlabtestbatchById(id: string, body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<DeliveryPackageRequiredlabtestbatch>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersDeliveryPackageRequiredlabtestbatchById(id, body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetDeliveryPackageWholesaleById
   */
  public async getTransfersDeliveryPackageWholesaleById(id: string, body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<DeliveryPackageWholesale>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersDeliveryPackageWholesaleById(id, body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetDeliveryTransporterById
   */
  public async getTransfersDeliveryTransporterById(id: string, body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<DeliveryTransporter>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersDeliveryTransporterById(id, body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetDeliveryTransporterDetailById
   */
  public async getTransfersDeliveryTransporterDetailById(id: string, body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<DeliveryTransporterDetail>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersDeliveryTransporterDetailById(id, body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetHub
   */
  public async getTransfersHub(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Hub>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersHub(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetIncomingTransfers
   */
  public async getIncomingTransfers(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Transfer>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getIncomingTransfers(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Get Transfer Manifest HTML for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetManifestHtmlById
   */
  public async getTransfersManifestHtmlById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersManifestHtmlById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Get Transfer Manifest PDF for a given Transfer Id
   * Permissions Required:
   * - Manage Transfer Templates
   * - View Transfer Templates
   * GET GetManifestPdfById
   */
  public async getTransfersManifestPdfById(id: string, body?: any, licenseNumber?: string): Promise<ManifestPdf> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersManifestPdfById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a list of deliveries associated with a specific transfer template.
   * Permissions Required:
   * - Manage Transfer Templates
   * - View Transfer Templates
   * GET GetOutgoingTemplateDeliveryById
   */
  public async getTransfersOutgoingTemplateDeliveryById(id: string, body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<TemplateDelivery>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersOutgoingTemplateDeliveryById(id, body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfer Templates
   * - View Transfer Templates
   * GET GetOutgoingTemplateDeliveryPackageById
   */
  public async getTransfersOutgoingTemplateDeliveryPackageById(id: string, body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<TemplateDeliveryPackage>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersOutgoingTemplateDeliveryPackageById(id, body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfer Templates
   * - View Transfer Templates
   * GET GetOutgoingTemplateDeliveryTransporterById
   */
  public async getTransfersOutgoingTemplateDeliveryTransporterById(id: string, body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<TemplateDeliveryTransporter>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersOutgoingTemplateDeliveryTransporterById(id, body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
   * Permissions Required:
   * - Manage Transfer Templates
   * - View Transfer Templates
   * GET GetOutgoingTemplateDeliveryTransporterDetailById
   */
  public async getTransfersOutgoingTemplateDeliveryTransporterDetailById(id: string, body?: any): Promise<TemplateDeliveryTransporterDetail> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersOutgoingTemplateDeliveryTransporterDetailById(id, body,
        )
    );
  }

  /**
   * Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
   * Permissions Required:
   * - Manage Transfer Templates
   * - View Transfer Templates
   * GET GetOutgoingTemplates
   */
  public async getTransfersOutgoingTemplates(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Template>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersOutgoingTemplates(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetOutgoingTransfers
   */
  public async getOutgoingTransfers(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Transfer>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getOutgoingTransfers(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of shipments with rejected packages for a Facility.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetRejectedTransfers
   */
  public async getRejectedTransfers(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<Transfer>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getRejectedTransfers(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
   * Permissions Required:
   * - Manage Transfers
   * - View Transfers
   * GET GetTransfersDeliveryById
   */
  public async getTransfersDeliveryById(id: string, body?: any, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<TransfersDelivery>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersDeliveryById(id, body,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a list of available transfer types for a Facility based on its license number.
   * GET GetTransfersTypes
   */
  public async getTransfersTypes(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<TransfersType>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getTransfersTypes(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Updates external incoming shipment plans for a Facility.
   * Permissions Required:
   * - Manage Transfers
   * PUT UpdateIncomingExternal
   */
  public async updateTransfersIncomingExternal(body: UpdateIncomingExternalRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateTransfersIncomingExternal(body,licenseNumber,
        )
    );
  }

  /**
   * Updates existing transfer templates for a Facility.
   * Permissions Required:
   * - Manage Transfer Templates
   * PUT UpdateOutgoingTemplates
   */
  public async updateTransfersOutgoingTemplates(body: UpdateOutgoingTemplatesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateTransfersOutgoingTemplates(body,licenseNumber,
        )
    );
  }
}

