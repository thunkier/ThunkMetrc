
import { MetrcClient as InternalMetrcClient } from '@thunkier/thunkmetrc-client';
import { MetrcRateLimiter } from '../RateLimiter';
import { iteratePages } from '../Utils';
import { 
    PaginatedResponse,
    ActiveDeliveriesRetailer,
    ActiveDelivery,
    ActiveReceipt,
    County,
    CreateDeliveriesRequestItem,
    CreateDeliveriesRetailerDepartRequestItem,
    CreateDeliveriesRetailerEndRequestItem,
    CreateDeliveriesRetailerRestockRequestItem,
    CreateReceiptsRequestItem,
    CreateSalesDeliveriesRetailerRequestItem,
    DeliveriesReturnReason,
    DeliveryRetailer,
    InactiveDeliveriesRetailer,
    InactiveDelivery,
    InactiveReceipt,
    PatientRegistrationLocation,
    ReceiptsExternalByExternalNumber,
    SalesDelivery,
    UpdateDeliveriesCompleteRequestItem,
    UpdateDeliveriesHubAcceptRequestItem,
    UpdateDeliveriesHubDepartRequestItem,
    UpdateDeliveriesHubRequestItem,
    UpdateDeliveriesHubVerifyIDRequestItem,
    UpdateDeliveriesRequestItem,
    UpdateDeliveriesRetailerRequestItem,
    UpdateReceiptsFinalizeRequestItem,
    UpdateReceiptsRequestItem,
    UpdateReceiptsUnfinalizeRequestItem,
    WriteResponse,
} from '../models';

export class SalesService {
  constructor(private client: InternalMetrcClient, private rateLimiter: MetrcRateLimiter) {}

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
  public async createSalesDeliveries(body: CreateDeliveriesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createSalesDeliveries(body,licenseNumber,
        )
    );
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
  public async createSalesDeliveriesRetailerDepart(body: CreateDeliveriesRetailerDepartRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createSalesDeliveriesRetailerDepart(body,licenseNumber,
        )
    );
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
  public async createSalesDeliveriesRetailerEnd(body: CreateDeliveriesRetailerEndRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createSalesDeliveriesRetailerEnd(body,licenseNumber,
        )
    );
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
  public async createSalesDeliveriesRetailerRestock(body: CreateDeliveriesRetailerRestockRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createSalesDeliveriesRetailerRestock(body,licenseNumber,
        )
    );
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
  public async createSalesReceipts(body: CreateReceiptsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createSalesReceipts(body,licenseNumber,
        )
    );
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
  public async createSalesDeliveriesRetailer(body: CreateSalesDeliveriesRetailerRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.createSalesDeliveriesRetailer(body,licenseNumber,
        )
    );
  }

  /**
   * Voids a sales delivery for a Facility using the provided License Number and delivery Id.
   * Permissions Required:
   * - Manage Sales Delivery
   * DELETE DeleteDeliveryById
   */
  public async deleteSalesDeliveryById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteSalesDeliveryById(id, body,licenseNumber,
        )
    );
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
  public async deleteSalesDeliveryRetailerById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteSalesDeliveryRetailerById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Archives a sales receipt for a Facility using the provided License Number and receipt Id.
   * Permissions Required:
   * - Manage Sales
   * DELETE DeleteReceiptById
   */
  public async deleteSalesReceiptById(id: string, body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,false,
        () => this.client.deleteSalesReceiptById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
   * Permissions Required:
   * - View Sales Delivery
   * - Manage Sales Delivery
   * GET GetActiveDeliveries
   */
  public async getSalesActiveDeliveries(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<ActiveDelivery>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesActiveDeliveries(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
   * Permissions Required:
   * - View Retailer Delivery
   * - Manage Retailer Delivery
   * GET GetActiveDeliveriesRetailer
   */
  public async getSalesActiveDeliveriesRetailer(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<ActiveDeliveriesRetailer>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesActiveDeliveriesRetailer(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
   * Permissions Required:
   * - View Sales
   * - Manage Sales
   * GET GetActiveReceipts
   */
  public async getSalesActiveReceipts(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<ActiveReceipt>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesActiveReceipts(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Returns a list of counties available for sales deliveries.
   * GET GetCounties
   */
  public async getSalesCounties(body?: any): Promise<PaginatedResponse<County>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesCounties(body,
        )
    );
  }

  /**
   * Returns a list of customer types.
   * GET GetCustomerTypes
   */
  public async getSalesCustomerTypes(body?: any): Promise<any> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesCustomerTypes(body,
        )
    );
  }

  /**
   * Returns a list of return reasons for sales deliveries based on the provided License Number.
   * Permissions Required:
   * - Sales Delivery
   * GET GetDeliveriesReturnReasons
   */
  public async getSalesDeliveriesReturnReasons(body?: any, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<DeliveriesReturnReason>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesDeliveriesReturnReasons(body,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Retrieves a retailer delivery record by its ID, with an optional License Number.
   * Permissions Required:
   * - View Retailer Delivery
   * - Manage Retailer Delivery
   * GET GetDeliveryRetailerById
   */
  public async getSalesDeliveryRetailerById(id: string, body?: any, licenseNumber?: string): Promise<DeliveryRetailer> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesDeliveryRetailerById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
   * Permissions Required:
   * - View Sales Delivery
   * - Manage Sales Delivery
   * GET GetInactiveDeliveries
   */
  public async getSalesInactiveDeliveries(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<InactiveDelivery>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesInactiveDeliveries(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
   * Permissions Required:
   * - View Retailer Delivery
   * - Manage Retailer Delivery
   * GET GetInactiveDeliveriesRetailer
   */
  public async getSalesInactiveDeliveriesRetailer(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<InactiveDeliveriesRetailer>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesInactiveDeliveriesRetailer(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
   * Permissions Required:
   * - View Sales
   * - Manage Sales
   * GET GetInactiveReceipts
   */
  public async getSalesInactiveReceipts(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, pageNumber?: number, pageSize?: string): Promise<PaginatedResponse<InactiveReceipt>> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesInactiveReceipts(body,lastModifiedEnd,lastModifiedStart,licenseNumber,pageNumber,pageSize,
        )
    );
  }

  /**
   * Returns a list of valid Patient registration locations for sales.
   * GET GetPatientRegistrationLocations
   */
  public async getSalesPatientRegistrationLocations(body?: any): Promise<PatientRegistrationLocation[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesPatientRegistrationLocations(body,
        )
    );
  }

  /**
   * Returns a list of available payment types for the specified License Number.
   * Permissions Required:
   * - View Sales Delivery
   * - Manage Sales Delivery
   * GET GetPaymentTypes
   */
  public async getSalesPaymentTypes(body?: any, licenseNumber?: string): Promise<any> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesPaymentTypes(body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a sales receipt by its Id, with an optional License Number.
   * Permissions Required:
   * - View Sales
   * - Manage Sales
   * GET GetReceiptById
   */
  public async getSalesReceiptById(id: string, body?: any, licenseNumber?: string): Promise<ActiveReceipt> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesReceiptById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a Sales Receipt by its external number, with an optional License Number.
   * Permissions Required:
   * - View Sales
   * - Manage Sales
   * GET GetReceiptsExternalByExternalNumber
   */
  public async getSalesReceiptsExternalByExternalNumber(externalNumber: string, body?: any, licenseNumber?: string): Promise<ReceiptsExternalByExternalNumber[]> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesReceiptsExternalByExternalNumber(externalNumber, body,licenseNumber,
        )
    );
  }

  /**
   * Retrieves a sales delivery record by its Id, with an optional License Number.
   * Permissions Required:
   * - View Sales Delivery
   * - Manage Sales Delivery
   * GET GetSalesDeliveryById
   */
  public async getSalesDeliveryById(id: string, body?: any, licenseNumber?: string): Promise<SalesDelivery> {
    return this.rateLimiter.execute(null,true,
        () => this.client.getSalesDeliveryById(id, body,licenseNumber,
        )
    );
  }

  /**
   * Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * Permissions Required:
   * - Manage Sales Delivery
   * PUT UpdateDeliveries
   */
  public async updateSalesDeliveries(body: UpdateDeliveriesRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateSalesDeliveries(body,licenseNumber,
        )
    );
  }

  /**
   * Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
   * Permissions Required:
   * - Manage Sales Delivery
   * PUT UpdateDeliveriesComplete
   */
  public async updateSalesDeliveriesComplete(body: UpdateDeliveriesCompleteRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateSalesDeliveriesComplete(body,licenseNumber,
        )
    );
  }

  /**
   * Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * Permissions Required:
   * - Manage Sales Delivery, Manage Sales Delivery Hub
   * PUT UpdateDeliveriesHub
   */
  public async updateSalesDeliveriesHub(body: UpdateDeliveriesHubRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateSalesDeliveriesHub(body,licenseNumber,
        )
    );
  }

  /**
   * Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
   * Permissions Required:
   * - Manage Sales Delivery Hub
   * PUT UpdateDeliveriesHubAccept
   */
  public async updateSalesDeliveriesHubAccept(body: UpdateDeliveriesHubAcceptRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateSalesDeliveriesHubAccept(body,licenseNumber,
        )
    );
  }

  /**
   * Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
   * Permissions Required:
   * - Manage Sales Delivery Hub
   * PUT UpdateDeliveriesHubDepart
   */
  public async updateSalesDeliveriesHubDepart(body: UpdateDeliveriesHubDepartRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateSalesDeliveriesHubDepart(body,licenseNumber,
        )
    );
  }

  /**
   * Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
   * Permissions Required:
   * - Manage Sales Delivery Hub
   * PUT UpdateDeliveriesHubVerifyID
   */
  public async updateSalesDeliveriesHubVerifyID(body: UpdateDeliveriesHubVerifyIDRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateSalesDeliveriesHubVerifyID(body,licenseNumber,
        )
    );
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
  public async updateSalesDeliveriesRetailer(body: UpdateDeliveriesRetailerRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateSalesDeliveriesRetailer(body,licenseNumber,
        )
    );
  }

  /**
   * Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * Permissions Required:
   * - Manage Sales
   * PUT UpdateReceipts
   */
  public async updateSalesReceipts(body: UpdateReceiptsRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateSalesReceipts(body,licenseNumber,
        )
    );
  }

  /**
   * Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
   * Permissions Required:
   * - Manage Sales
   * PUT UpdateReceiptsFinalize
   */
  public async updateSalesReceiptsFinalize(body: UpdateReceiptsFinalizeRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateSalesReceiptsFinalize(body,licenseNumber,
        )
    );
  }

  /**
   * Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
   * Permissions Required:
   * - Manage Sales
   * PUT UpdateReceiptsUnfinalize
   */
  public async updateSalesReceiptsUnfinalize(body: UpdateReceiptsUnfinalizeRequestItem[], licenseNumber?: string): Promise<WriteResponse> {
    return this.rateLimiter.execute(null,false,
        () => this.client.updateSalesReceiptsUnfinalize(body,licenseNumber,
        )
    );
  }
}

