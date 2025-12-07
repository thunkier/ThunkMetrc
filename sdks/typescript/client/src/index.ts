import axios, { AxiosInstance, AxiosResponse } from 'axios';

export class MetrcClient {
  private client: AxiosInstance;

  constructor(baseUrl: string, vendorKey: string, userKey: string) {
    this.client = axios.create({
      baseURL: baseUrl,
      auth: {
        username: vendorKey,
        password: userKey
      }
    });
  }

  /**
   * Permissions Required:
   *   - Manage Locations
   *
   * POST Create V1
   */
  public async locationsCreateV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v1/create?${queryStr}` : `/locations/v1/create`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Creates new locations for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * POST Create V2
   */
  public async locationsCreateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2?${queryStr}` : `/locations/v2`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Locations
   *
   * POST CreateUpdate V1
   */
  public async locationsCreateUpdateV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v1/update?${queryStr}` : `/locations/v1/update`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Locations
   *
   * DELETE Delete V1
   */
  public async locationsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v1/${encodeURIComponent(id)}?${queryStr}` : `/locations/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2/${encodeURIComponent(id)}?${queryStr}` : `/locations/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Locations
   *
   * GET Get V1
   */
  public async locationsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v1/${encodeURIComponent(id)}?${queryStr}` : `/locations/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2/${encodeURIComponent(id)}?${queryStr}` : `/locations/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Locations
   *
   * GET GetActive V1
   */
  public async locationsGetActiveV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v1/active?${queryStr}` : `/locations/v1/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2/active?${queryStr}` : `/locations/v2/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2/inactive?${queryStr}` : `/locations/v2/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Locations
   *
   * GET GetTypes V1
   */
  public async locationsGetTypesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v1/types?${queryStr}` : `/locations/v1/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2/types?${queryStr}` : `/locations/v2/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Updates existing locations for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * PUT Update V2
   */
  public async locationsUpdateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/locations/v2?${queryStr}` : `/locations/v2`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v1/statuses/${encodeURIComponent(patientLicenseNumber)}?${queryStr}` : `/patients/v1/statuses/${encodeURIComponent(patientLicenseNumber)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v2/statuses/${encodeURIComponent(patientLicenseNumber)}?${queryStr}` : `/patients/v2/statuses/${encodeURIComponent(patientLicenseNumber)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditives V1
   */
  public async plantsCreateAdditivesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/additives?${queryStr}` : `/plants/v1/additives`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Records additive usage details applied to specific plants at a Facility.
   * 
   *   Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditives V2
   */
  public async plantsCreateAdditivesV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/additives?${queryStr}` : `/plants/v2/additives`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Plants
   *   - Manage Plants Additives
   *
   * POST CreateAdditivesBylocation V1
   */
  public async plantsCreateAdditivesBylocationV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/additives/bylocation?${queryStr}` : `/plants/v1/additives/bylocation`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantsCreateAdditivesBylocationV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/additives/bylocation?${queryStr}` : `/plants/v2/additives/bylocation`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Records additive usage for plants by location using a predefined additive template at a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditivesBylocationUsingtemplate V2
   */
  public async plantsCreateAdditivesBylocationUsingtemplateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/additives/bylocation/usingtemplate?${queryStr}` : `/plants/v2/additives/bylocation/usingtemplate`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Records additive usage for plants using predefined additive templates at a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditivesUsingtemplate V2
   */
  public async plantsCreateAdditivesUsingtemplateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/additives/usingtemplate?${queryStr}` : `/plants/v2/additives/usingtemplate`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * POST CreateChangegrowthphases V1
   */
  public async plantsCreateChangegrowthphasesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/changegrowthphases?${queryStr}` : `/plants/v1/changegrowthphases`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantsCreateHarvestplantsV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/harvestplants?${queryStr}` : `/plants/v1/harvestplants`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantsCreateManicureV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/manicure?${queryStr}` : `/plants/v2/manicure`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manicure/Harvest Veg/Flower Plants
   *
   * POST CreateManicureplants V1
   */
  public async plantsCreateManicureplantsV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/manicureplants?${queryStr}` : `/plants/v1/manicureplants`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *   - Manage Veg/Flower Plants Inventory
   *
   * POST CreateMoveplants V1
   */
  public async plantsCreateMoveplantsV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/moveplants?${queryStr}` : `/plants/v1/moveplants`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantsCreatePlantbatchPackageV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/create/plantbatch/packages?${queryStr}` : `/plants/v1/create/plantbatch/packages`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantsCreatePlantbatchPackageV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/plantbatch/packages?${queryStr}` : `/plants/v2/plantbatch/packages`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantsCreatePlantingsV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/create/plantings?${queryStr}` : `/plants/v1/create/plantings`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantsCreatePlantingsV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/plantings?${queryStr}` : `/plants/v2/plantings`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Plants Waste
   *
   * POST CreateWaste V1
   */
  public async plantsCreateWasteV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/waste?${queryStr}` : `/plants/v1/waste`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Records waste events for plants at a Facility, including method, reason, and location details.
   * 
   *   Permissions Required:
   *   - Manage Plants Waste
   *
   * POST CreateWaste V2
   */
  public async plantsCreateWasteV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/waste?${queryStr}` : `/plants/v2/waste`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *   - Destroy Veg/Flower Plants
   *
   * DELETE Delete V1
   */
  public async plantsDeleteV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1?${queryStr}` : `/plants/v1`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2?${queryStr}` : `/plants/v2`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET Get V1
   */
  public async plantsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/${encodeURIComponent(id)}?${queryStr}` : `/plants/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/${encodeURIComponent(id)}?${queryStr}` : `/plants/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View/Manage Plants Additives
   *
   * GET GetAdditives V1
   */
  public async plantsGetAdditivesV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/additives?${queryStr}` : `/plants/v1/additives`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/additives?${queryStr}` : `/plants/v2/additives`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   -
   *
   * GET GetAdditivesTypes V1
   */
  public async plantsGetAdditivesTypesV1(body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/additives/types?${queryStr}` : `/plants/v1/additives/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/additives/types?${queryStr}` : `/plants/v2/additives/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetByLabel V1
   */
  public async plantsGetByLabelV1(label: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/${encodeURIComponent(label)}?${queryStr}` : `/plants/v1/${encodeURIComponent(label)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/${encodeURIComponent(label)}?${queryStr}` : `/plants/v2/${encodeURIComponent(label)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetFlowering V1
   */
  public async plantsGetFloweringV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/flowering?${queryStr}` : `/plants/v1/flowering`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/flowering?${queryStr}` : `/plants/v2/flowering`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetGrowthPhases V1
   */
  public async plantsGetGrowthPhasesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/growthphases?${queryStr}` : `/plants/v1/growthphases`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/growthphases?${queryStr}` : `/plants/v2/growthphases`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetInactive V1
   */
  public async plantsGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/inactive?${queryStr}` : `/plants/v1/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/inactive?${queryStr}` : `/plants/v2/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/mother?${queryStr}` : `/plants/v2/mother`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/mother/inactive?${queryStr}` : `/plants/v2/mother/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/mother/onhold?${queryStr}` : `/plants/v2/mother/onhold`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetOnhold V1
   */
  public async plantsGetOnholdV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/onhold?${queryStr}` : `/plants/v1/onhold`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/onhold?${queryStr}` : `/plants/v2/onhold`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Veg/Flower Plants
   *
   * GET GetVegetative V1
   */
  public async plantsGetVegetativeV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/vegetative?${queryStr}` : `/plants/v1/vegetative`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/vegetative?${queryStr}` : `/plants/v2/vegetative`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/waste?${queryStr}` : `/plants/v2/waste`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetWasteMethodsAll V1
   */
  public async plantsGetWasteMethodsAllV1(body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/waste/methods/all?${queryStr}` : `/plants/v1/waste/methods/all`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/waste/methods/all?${queryStr}` : `/plants/v2/waste/methods/all`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/waste/${encodeURIComponent(id)}/package?${queryStr}` : `/plants/v2/waste/${encodeURIComponent(id)}/package`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/waste/${encodeURIComponent(id)}/plant?${queryStr}` : `/plants/v2/waste/${encodeURIComponent(id)}/plant`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetWasteReasons V1
   */
  public async plantsGetWasteReasonsV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v1/waste/reasons?${queryStr}` : `/plants/v1/waste/reasons`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/waste/reasons?${queryStr}` : `/plants/v2/waste/reasons`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
  public async plantsUpdateAdjustV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/adjust?${queryStr}` : `/plants/v2/adjust`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async plantsUpdateGrowthphaseV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/growthphase?${queryStr}` : `/plants/v2/growthphase`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async plantsUpdateHarvestV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/harvest?${queryStr}` : `/plants/v2/harvest`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async plantsUpdateLocationV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/location?${queryStr}` : `/plants/v2/location`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async plantsUpdateMergeV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/merge?${queryStr}` : `/plants/v2/merge`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Splits an existing plant group into multiple groups within a Facility.
   * 
   *   Permissions Required:
   *   - View Plant
   *
   * PUT UpdateSplit V2
   */
  public async plantsUpdateSplitV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/split?${queryStr}` : `/plants/v2/split`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async plantsUpdateStrainV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/strain?${queryStr}` : `/plants/v2/strain`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async plantsUpdateTagV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plants/v2/tag?${queryStr}` : `/plants/v2/tag`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async retailIdCreateAssociateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/retailid/v2/associate?${queryStr}` : `/retailid/v2/associate`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async retailIdCreateGenerateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/retailid/v2/generate?${queryStr}` : `/retailid/v2/generate`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async retailIdCreateMergeV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/retailid/v2/merge?${queryStr}` : `/retailid/v2/merge`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async retailIdCreatePackageInfoV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/retailid/v2/packages/info?${queryStr}` : `/retailid/v2/packages/info`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/retailid/v2/receive/${encodeURIComponent(label)}?${queryStr}` : `/retailid/v2/receive/${encodeURIComponent(label)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/retailid/v2/receive/qr/${encodeURIComponent(shortCode)}?${queryStr}` : `/retailid/v2/receive/qr/${encodeURIComponent(shortCode)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (userKey) query.append('userKey', userKey);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sandbox/v2/integrator/setup?${queryStr}` : `/sandbox/v2/integrator/setup`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/facilities/v1?${queryStr}` : `/facilities/v1`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/facilities/v2?${queryStr}` : `/facilities/v2`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Adds new patients to a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Patients
   *
   * POST Create V2
   */
  public async patientsCreateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v2?${queryStr}` : `/patients/v2`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Patients
   *
   * POST CreateAdd V1
   */
  public async patientsCreateAddV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v1/add?${queryStr}` : `/patients/v1/add`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Patients
   *
   * POST CreateUpdate V1
   */
  public async patientsCreateUpdateV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v1/update?${queryStr}` : `/patients/v1/update`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Patients
   *
   * DELETE Delete V1
   */
  public async patientsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v1/${encodeURIComponent(id)}?${queryStr}` : `/patients/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v2/${encodeURIComponent(id)}?${queryStr}` : `/patients/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Patients
   *
   * GET Get V1
   */
  public async patientsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v1/${encodeURIComponent(id)}?${queryStr}` : `/patients/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v2/${encodeURIComponent(id)}?${queryStr}` : `/patients/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Patients
   *
   * GET GetActive V1
   */
  public async patientsGetActiveV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v1/active?${queryStr}` : `/patients/v1/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v2/active?${queryStr}` : `/patients/v2/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Updates Patient information for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Patients
   *
   * PUT Update V2
   */
  public async patientsUpdateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patients/v2?${queryStr}` : `/patients/v2`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Transfers
   *
   * POST CreateExternalIncoming V1
   */
  public async transfersCreateExternalIncomingV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/external/incoming?${queryStr}` : `/transfers/v1/external/incoming`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Creates external incoming shipment plans for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *
   * POST CreateExternalIncoming V2
   */
  public async transfersCreateExternalIncomingV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/external/incoming?${queryStr}` : `/transfers/v2/external/incoming`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Transfer Templates
   *
   * POST CreateTemplates V1
   */
  public async transfersCreateTemplatesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/templates?${queryStr}` : `/transfers/v1/templates`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Creates new transfer templates for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transfer Templates
   *
   * POST CreateTemplatesOutgoing V2
   */
  public async transfersCreateTemplatesOutgoingV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing?${queryStr}` : `/transfers/v2/templates/outgoing`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Transfers
   *
   * DELETE DeleteExternalIncoming V1
   */
  public async transfersDeleteExternalIncomingV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/external/incoming/${encodeURIComponent(id)}?${queryStr}` : `/transfers/v1/external/incoming/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/external/incoming/${encodeURIComponent(id)}?${queryStr}` : `/transfers/v2/external/incoming/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Transfer Templates
   *
   * DELETE DeleteTemplates V1
   */
  public async transfersDeleteTemplatesV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/templates/${encodeURIComponent(id)}?${queryStr}` : `/transfers/v1/templates/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/${encodeURIComponent(id)}?${queryStr}` : `/transfers/v2/templates/outgoing/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetDeliveriesPackagesStates V1
   */
  public async transfersGetDeliveriesPackagesStatesV1(body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/deliveries/packages/states?${queryStr}` : `/transfers/v1/deliveries/packages/states`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/deliveries/packages/states?${queryStr}` : `/transfers/v2/deliveries/packages/states`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/${encodeURIComponent(id)}/deliveries?${queryStr}` : `/transfers/v1/${encodeURIComponent(id)}/deliveries`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/${encodeURIComponent(id)}/deliveries?${queryStr}` : `/transfers/v2/${encodeURIComponent(id)}/deliveries`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/deliveries/${encodeURIComponent(id)}/packages?${queryStr}` : `/transfers/v1/deliveries/${encodeURIComponent(id)}/packages`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/deliveries/${encodeURIComponent(id)}/packages?${queryStr}` : `/transfers/v2/deliveries/${encodeURIComponent(id)}/packages`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/deliveries/package/${encodeURIComponent(id)}/requiredlabtestbatches?${queryStr}` : `/transfers/v1/deliveries/package/${encodeURIComponent(id)}/requiredlabtestbatches`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/deliveries/package/${encodeURIComponent(id)}/requiredlabtestbatches?${queryStr}` : `/transfers/v2/deliveries/package/${encodeURIComponent(id)}/requiredlabtestbatches`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/deliveries/${encodeURIComponent(id)}/packages/wholesale?${queryStr}` : `/transfers/v1/deliveries/${encodeURIComponent(id)}/packages/wholesale`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/deliveries/${encodeURIComponent(id)}/packages/wholesale?${queryStr}` : `/transfers/v2/deliveries/${encodeURIComponent(id)}/packages/wholesale`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/deliveries/${encodeURIComponent(id)}/transporters?${queryStr}` : `/transfers/v1/deliveries/${encodeURIComponent(id)}/transporters`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/deliveries/${encodeURIComponent(id)}/transporters?${queryStr}` : `/transfers/v2/deliveries/${encodeURIComponent(id)}/transporters`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/deliveries/${encodeURIComponent(id)}/transporters/details?${queryStr}` : `/transfers/v1/deliveries/${encodeURIComponent(id)}/transporters/details`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/deliveries/${encodeURIComponent(id)}/transporters/details?${queryStr}` : `/transfers/v2/deliveries/${encodeURIComponent(id)}/transporters/details`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (estimatedArrivalEnd) query.append('estimatedArrivalEnd', estimatedArrivalEnd);
    if (estimatedArrivalStart) query.append('estimatedArrivalStart', estimatedArrivalStart);
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/hub?${queryStr}` : `/transfers/v2/hub`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Transfers
   *
   * GET GetIncoming V1
   */
  public async transfersGetIncomingV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/incoming?${queryStr}` : `/transfers/v1/incoming`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/incoming?${queryStr}` : `/transfers/v2/incoming`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Transfers
   *
   * GET GetOutgoing V1
   */
  public async transfersGetOutgoingV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/outgoing?${queryStr}` : `/transfers/v1/outgoing`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/outgoing?${queryStr}` : `/transfers/v2/outgoing`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Transfers
   *
   * GET GetRejected V1
   */
  public async transfersGetRejectedV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/rejected?${queryStr}` : `/transfers/v1/rejected`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/rejected?${queryStr}` : `/transfers/v2/rejected`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Transfer Templates
   *
   * GET GetTemplates V1
   */
  public async transfersGetTemplatesV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/templates?${queryStr}` : `/transfers/v1/templates`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Transfer Templates
   *
   * GET GetTemplatesDelivery V1
   */
  public async transfersGetTemplatesDeliveryV1(id: string, body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/templates/${encodeURIComponent(id)}/deliveries?${queryStr}` : `/transfers/v1/templates/${encodeURIComponent(id)}/deliveries`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/templates/deliveries/${encodeURIComponent(id)}/packages?${queryStr}` : `/transfers/v1/templates/deliveries/${encodeURIComponent(id)}/packages`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/templates/deliveries/${encodeURIComponent(id)}/transporters?${queryStr}` : `/transfers/v1/templates/deliveries/${encodeURIComponent(id)}/transporters`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/templates/deliveries/${encodeURIComponent(id)}/transporters/details?${queryStr}` : `/transfers/v1/templates/deliveries/${encodeURIComponent(id)}/transporters/details`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing?${queryStr}` : `/transfers/v2/templates/outgoing`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/${encodeURIComponent(id)}/deliveries?${queryStr}` : `/transfers/v2/templates/outgoing/${encodeURIComponent(id)}/deliveries`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/packages?${queryStr}` : `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/packages`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/transporters?${queryStr}` : `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/transporters`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/transporters/details?${queryStr}` : `/transfers/v2/templates/outgoing/deliveries/${encodeURIComponent(id)}/transporters/details`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetTypes V1
   */
  public async transfersGetTypesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/types?${queryStr}` : `/transfers/v1/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/types?${queryStr}` : `/transfers/v2/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Transfers
   *
   * PUT UpdateExternalIncoming V1
   */
  public async transfersUpdateExternalIncomingV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/external/incoming?${queryStr}` : `/transfers/v1/external/incoming`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Updates external incoming shipment plans for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transfers
   *
   * PUT UpdateExternalIncoming V2
   */
  public async transfersUpdateExternalIncomingV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/external/incoming?${queryStr}` : `/transfers/v2/external/incoming`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Transfer Templates
   *
   * PUT UpdateTemplates V1
   */
  public async transfersUpdateTemplatesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v1/templates?${queryStr}` : `/transfers/v1/templates`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Updates existing transfer templates for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transfer Templates
   *
   * PUT UpdateTemplatesOutgoing V2
   */
  public async transfersUpdateTemplatesOutgoingV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transfers/v2/templates/outgoing?${queryStr}` : `/transfers/v2/templates/outgoing`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Creates new driver records for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transporters
   *
   * POST CreateDriver V2
   */
  public async transportersCreateDriverV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/drivers?${queryStr}` : `/transporters/v2/drivers`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Creates new vehicle records for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transporters
   *
   * POST CreateVehicle V2
   */
  public async transportersCreateVehicleV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/vehicles?${queryStr}` : `/transporters/v2/vehicles`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/drivers/${encodeURIComponent(id)}?${queryStr}` : `/transporters/v2/drivers/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/vehicles/${encodeURIComponent(id)}?${queryStr}` : `/transporters/v2/vehicles/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/drivers/${encodeURIComponent(id)}?${queryStr}` : `/transporters/v2/drivers/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/drivers?${queryStr}` : `/transporters/v2/drivers`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/vehicles/${encodeURIComponent(id)}?${queryStr}` : `/transporters/v2/vehicles/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/vehicles?${queryStr}` : `/transporters/v2/vehicles`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Updates existing driver records for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Transporters
   *
   * PUT UpdateDriver V2
   */
  public async transportersUpdateDriverV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/drivers?${queryStr}` : `/transporters/v2/drivers`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Updates existing vehicle records for a facility.
   * 
   *   Permissions Required:
   *   - Manage Transporters
   *
   * PUT UpdateVehicle V2
   */
  public async transportersUpdateVehicleV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/transporters/v2/vehicles?${queryStr}` : `/transporters/v2/vehicles`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetActive V1
   */
  public async unitsOfMeasureGetActiveV1(body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/unitsofmeasure/v1/active?${queryStr}` : `/unitsofmeasure/v1/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/unitsofmeasure/v2/active?${queryStr}` : `/unitsofmeasure/v2/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/unitsofmeasure/v2/inactive?${queryStr}` : `/unitsofmeasure/v2/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/wastemethods/v2?${queryStr}` : `/wastemethods/v2`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/caregivers/v1/status/${encodeURIComponent(caregiverLicenseNumber)}?${queryStr}` : `/caregivers/v1/status/${encodeURIComponent(caregiverLicenseNumber)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/caregivers/v2/status/${encodeURIComponent(caregiverLicenseNumber)}?${queryStr}` : `/caregivers/v2/status/${encodeURIComponent(caregiverLicenseNumber)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Employees
   *
   * GET GetAll V1
   */
  public async employeesGetAllV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/employees/v1?${queryStr}` : `/employees/v1`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/employees/v2?${queryStr}` : `/employees/v2`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (employeeLicenseNumber) query.append('employeeLicenseNumber', employeeLicenseNumber);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/employees/v2/permissions?${queryStr}` : `/employees/v2/permissions`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * NOTE: To include a photo with an item, first use POST /items/v1/photo to POST the photo, and then use the returned ID in the request body in this endpoint.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * POST Create V1
   */
  public async itemsCreateV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v1/create?${queryStr}` : `/items/v1/create`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Creates one or more new products for the specified Facility. NOTE: To include a photo with an item, first use POST /items/v2/photo to POST the photo, and then use the returned Id in the request body in this endpoint.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * POST Create V2
   */
  public async itemsCreateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2?${queryStr}` : `/items/v2`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Creates one or more new item brands for the specified Facility identified by the License Number.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * POST CreateBrand V2
   */
  public async itemsCreateBrandV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/brand?${queryStr}` : `/items/v2/brand`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * POST CreateFile V2
   */
  public async itemsCreateFileV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/file?${queryStr}` : `/items/v2/file`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * POST CreatePhoto V1
   */
  public async itemsCreatePhotoV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v1/photo?${queryStr}` : `/items/v1/photo`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * POST CreatePhoto V2
   */
  public async itemsCreatePhotoV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/photo?${queryStr}` : `/items/v2/photo`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * POST CreateUpdate V1
   */
  public async itemsCreateUpdateV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v1/update?${queryStr}` : `/items/v1/update`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * DELETE Delete V1
   */
  public async itemsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v1/${encodeURIComponent(id)}?${queryStr}` : `/items/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/brand/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/brand/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * GET Get V1
   */
  public async itemsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v1/${encodeURIComponent(id)}?${queryStr}` : `/items/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * GET GetActive V1
   */
  public async itemsGetActiveV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v1/active?${queryStr}` : `/items/v1/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/active?${queryStr}` : `/items/v2/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * GET GetBrands V1
   */
  public async itemsGetBrandsV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v1/brands?${queryStr}` : `/items/v1/brands`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/brands?${queryStr}` : `/items/v2/brands`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetCategories V1
   */
  public async itemsGetCategoriesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v1/categories?${queryStr}` : `/items/v1/categories`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/categories?${queryStr}` : `/items/v2/categories`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/file/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/file/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * GET GetInactive V1
   */
  public async itemsGetInactiveV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v1/inactive?${queryStr}` : `/items/v1/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/inactive?${queryStr}` : `/items/v2/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Items
   *
   * GET GetPhoto V1
   */
  public async itemsGetPhotoV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v1/photo/${encodeURIComponent(id)}?${queryStr}` : `/items/v1/photo/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/photo/${encodeURIComponent(id)}?${queryStr}` : `/items/v2/photo/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Updates one or more existing products for the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * PUT Update V2
   */
  public async itemsUpdateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2?${queryStr}` : `/items/v2`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Updates one or more existing item brands for the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Items
   *
   * PUT UpdateBrand V2
   */
  public async itemsUpdateBrandV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/items/v2/brand?${queryStr}` : `/items/v2/brand`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST Create V1
   */
  public async packagesCreateV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/create?${queryStr}` : `/packages/v1/create`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async packagesCreateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2?${queryStr}` : `/packages/v2`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreateAdjust V1
   */
  public async packagesCreateAdjustV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/adjust?${queryStr}` : `/packages/v1/adjust`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async packagesCreateAdjustV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/adjust?${queryStr}` : `/packages/v2/adjust`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreateChangeItem V1
   */
  public async packagesCreateChangeItemV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/change/item?${queryStr}` : `/packages/v1/change/item`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreateChangeLocation V1
   */
  public async packagesCreateChangeLocationV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/change/locations?${queryStr}` : `/packages/v1/change/locations`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreateFinish V1
   */
  public async packagesCreateFinishV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/finish?${queryStr}` : `/packages/v1/finish`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async packagesCreatePlantingsV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/create/plantings?${queryStr}` : `/packages/v1/create/plantings`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async packagesCreatePlantingsV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/plantings?${queryStr}` : `/packages/v2/plantings`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreateRemediate V1
   */
  public async packagesCreateRemediateV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/remediate?${queryStr}` : `/packages/v1/remediate`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Create/Submit/Discontinue Packages
   *
   * POST CreateTesting V1
   */
  public async packagesCreateTestingV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/create/testing?${queryStr}` : `/packages/v1/create/testing`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async packagesCreateTestingV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/testing?${queryStr}` : `/packages/v2/testing`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreateUnfinish V1
   */
  public async packagesCreateUnfinishV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/unfinish?${queryStr}` : `/packages/v1/unfinish`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/${encodeURIComponent(id)}?${queryStr}` : `/packages/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *
   * GET Get V1
   */
  public async packagesGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/${encodeURIComponent(id)}?${queryStr}` : `/packages/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/${encodeURIComponent(id)}?${queryStr}` : `/packages/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *
   * GET GetActive V1
   */
  public async packagesGetActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/active?${queryStr}` : `/packages/v1/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/active?${queryStr}` : `/packages/v2/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetAdjustReasons V1
   */
  public async packagesGetAdjustReasonsV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/adjust/reasons?${queryStr}` : `/packages/v1/adjust/reasons`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/adjust/reasons?${queryStr}` : `/packages/v2/adjust/reasons`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *
   * GET GetByLabel V1
   */
  public async packagesGetByLabelV1(label: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/${encodeURIComponent(label)}?${queryStr}` : `/packages/v1/${encodeURIComponent(label)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/${encodeURIComponent(label)}?${queryStr}` : `/packages/v2/${encodeURIComponent(label)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *
   * GET GetInactive V1
   */
  public async packagesGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/inactive?${queryStr}` : `/packages/v1/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/inactive?${queryStr}` : `/packages/v2/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/intransit?${queryStr}` : `/packages/v2/intransit`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/labsamples?${queryStr}` : `/packages/v2/labsamples`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *
   * GET GetOnhold V1
   */
  public async packagesGetOnholdV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/onhold?${queryStr}` : `/packages/v1/onhold`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/onhold?${queryStr}` : `/packages/v2/onhold`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/${encodeURIComponent(id)}/source/harvests?${queryStr}` : `/packages/v2/${encodeURIComponent(id)}/source/harvests`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/transferred?${queryStr}` : `/packages/v2/transferred`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetTypes V1
   */
  public async packagesGetTypesV1(body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/types?${queryStr}` : `/packages/v1/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/types?${queryStr}` : `/packages/v2/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
  public async packagesUpdateAdjustV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/adjust?${queryStr}` : `/packages/v2/adjust`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *   - Manage Package Notes
   *
   * PUT UpdateChangeNote V1
   */
  public async packagesUpdateChangeNoteV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v1/change/note?${queryStr}` : `/packages/v1/change/note`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateDecontaminateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/decontaminate?${queryStr}` : `/packages/v2/decontaminate`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateDonationFlagV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/donation/flag?${queryStr}` : `/packages/v2/donation/flag`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateDonationUnflagV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/donation/unflag?${queryStr}` : `/packages/v2/donation/unflag`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateExternalidV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/externalid?${queryStr}` : `/packages/v2/externalid`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateFinishV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/finish?${queryStr}` : `/packages/v2/finish`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateItemV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/item?${queryStr}` : `/packages/v2/item`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateLabTestRequiredV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/labtests/required?${queryStr}` : `/packages/v2/labtests/required`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateLocationV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/location?${queryStr}` : `/packages/v2/location`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateNoteV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/note?${queryStr}` : `/packages/v2/note`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateRemediateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/remediate?${queryStr}` : `/packages/v2/remediate`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateTradesampleFlagV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/tradesample/flag?${queryStr}` : `/packages/v2/tradesample/flag`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateTradesampleUnflagV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/tradesample/unflag?${queryStr}` : `/packages/v2/tradesample/unflag`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateUnfinishV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/unfinish?${queryStr}` : `/packages/v2/unfinish`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async packagesUpdateUsebydateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/packages/v2/usebydate?${queryStr}` : `/packages/v2/usebydate`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * POST Create V1
   */
  public async patientCheckInsCreateV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v1?${queryStr}` : `/patient-checkins/v1`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Records patient check-ins for a specified Facility.
   * 
   *   Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * POST Create V2
   */
  public async patientCheckInsCreateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v2?${queryStr}` : `/patient-checkins/v2`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * DELETE Delete V1
   */
  public async patientCheckInsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v1/${encodeURIComponent(id)}?${queryStr}` : `/patient-checkins/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v2/${encodeURIComponent(id)}?${queryStr}` : `/patient-checkins/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * GET GetAll V1
   */
  public async patientCheckInsGetAllV1(body?: any, checkinDateEnd?: string, checkinDateStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (checkinDateEnd) query.append('checkinDateEnd', checkinDateEnd);
    if (checkinDateStart) query.append('checkinDateStart', checkinDateStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v1?${queryStr}` : `/patient-checkins/v1`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (checkinDateEnd) query.append('checkinDateEnd', checkinDateEnd);
    if (checkinDateStart) query.append('checkinDateStart', checkinDateStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v2?${queryStr}` : `/patient-checkins/v2`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetLocations V1
   */
  public async patientCheckInsGetLocationsV1(body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v1/locations?${queryStr}` : `/patient-checkins/v1/locations`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v2/locations?${queryStr}` : `/patient-checkins/v2/locations`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * PUT Update V1
   */
  public async patientCheckInsUpdateV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v1?${queryStr}` : `/patient-checkins/v1`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Updates patient check-ins for a specified Facility.
   * 
   *   Permissions Required:
   *   - ManagePatientsCheckIns
   *
   * PUT Update V2
   */
  public async patientCheckInsUpdateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/patient-checkins/v2?${queryStr}` : `/patient-checkins/v2`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditives V1
   */
  public async plantBatchesCreateAdditivesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/additives?${queryStr}` : `/plantbatches/v1/additives`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Records Additive usage details for plant batches at a specific Facility.
   * 
   *   Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditives V2
   */
  public async plantBatchesCreateAdditivesV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/additives?${queryStr}` : `/plantbatches/v2/additives`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Records Additive usage for plant batches at a Facility using predefined additive templates.
   * 
   *   Permissions Required:
   *   - Manage Plants Additives
   *
   * POST CreateAdditivesUsingtemplate V2
   */
  public async plantBatchesCreateAdditivesUsingtemplateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/additives/usingtemplate?${queryStr}` : `/plantbatches/v2/additives/usingtemplate`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *
   * POST CreateAdjust V1
   */
  public async plantBatchesCreateAdjustV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/adjust?${queryStr}` : `/plantbatches/v1/adjust`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantBatchesCreateAdjustV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/adjust?${queryStr}` : `/plantbatches/v2/adjust`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantBatchesCreateChangegrowthphaseV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/changegrowthphase?${queryStr}` : `/plantbatches/v1/changegrowthphase`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantBatchesCreateGrowthphaseV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/growthphase?${queryStr}` : `/plantbatches/v2/growthphase`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantBatchesCreatePackageV2(body?: any, isFromMotherPlant?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (isFromMotherPlant) query.append('isFromMotherPlant', isFromMotherPlant);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/packages?${queryStr}` : `/plantbatches/v2/packages`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantBatchesCreatePackageFrommotherplantV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/create/packages/frommotherplant?${queryStr}` : `/plantbatches/v1/create/packages/frommotherplant`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantBatchesCreatePackageFrommotherplantV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/packages/frommotherplant?${queryStr}` : `/plantbatches/v2/packages/frommotherplant`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantBatchesCreatePlantingsV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/plantings?${queryStr}` : `/plantbatches/v2/plantings`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *
   * POST CreateSplit V1
   */
  public async plantBatchesCreateSplitV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/split?${queryStr}` : `/plantbatches/v1/split`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantBatchesCreateSplitV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/split?${queryStr}` : `/plantbatches/v2/split`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Plants Waste
   *
   * POST CreateWaste V1
   */
  public async plantBatchesCreateWasteV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/waste?${queryStr}` : `/plantbatches/v1/waste`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Records waste information for plant batches based on the submitted data for the specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Plants Waste
   *
   * POST CreateWaste V2
   */
  public async plantBatchesCreateWasteV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/waste?${queryStr}` : `/plantbatches/v2/waste`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async plantBatchesCreatepackagesV1(body?: any, isFromMotherPlant?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (isFromMotherPlant) query.append('isFromMotherPlant', isFromMotherPlant);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/createpackages?${queryStr}` : `/plantbatches/v1/createpackages`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Manage Immature Plants Inventory
   *
   * POST Createplantings V1
   */
  public async plantBatchesCreateplantingsV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/createplantings?${queryStr}` : `/plantbatches/v1/createplantings`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *   - Destroy Immature Plants
   *
   * DELETE Delete V1
   */
  public async plantBatchesDeleteV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1?${queryStr}` : `/plantbatches/v1`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2?${queryStr}` : `/plantbatches/v2`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *
   * GET Get V1
   */
  public async plantBatchesGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/${encodeURIComponent(id)}?${queryStr}` : `/plantbatches/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/${encodeURIComponent(id)}?${queryStr}` : `/plantbatches/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *
   * GET GetActive V1
   */
  public async plantBatchesGetActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/active?${queryStr}` : `/plantbatches/v1/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/active?${queryStr}` : `/plantbatches/v2/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *
   * GET GetInactive V1
   */
  public async plantBatchesGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/inactive?${queryStr}` : `/plantbatches/v1/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/inactive?${queryStr}` : `/plantbatches/v2/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetTypes V1
   */
  public async plantBatchesGetTypesV1(body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/types?${queryStr}` : `/plantbatches/v1/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/types?${queryStr}` : `/plantbatches/v2/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/waste?${queryStr}` : `/plantbatches/v2/waste`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetWasteReasons V1
   */
  public async plantBatchesGetWasteReasonsV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/waste/reasons?${queryStr}` : `/plantbatches/v1/waste/reasons`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/waste/reasons?${queryStr}` : `/plantbatches/v2/waste/reasons`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
  public async plantBatchesUpdateLocationV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/location?${queryStr}` : `/plantbatches/v2/location`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Immature Plants
   *
   * PUT UpdateMoveplantbatches V1
   */
  public async plantBatchesUpdateMoveplantbatchesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v1/moveplantbatches?${queryStr}` : `/plantbatches/v1/moveplantbatches`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async plantBatchesUpdateNameV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/name?${queryStr}` : `/plantbatches/v2/name`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async plantBatchesUpdateStrainV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/strain?${queryStr}` : `/plantbatches/v2/strain`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async plantBatchesUpdateTagV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/plantbatches/v2/tag?${queryStr}` : `/plantbatches/v2/tag`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - ManageProcessingJobs
   *
   * POST CreateAdjust V1
   */
  public async processingJobsCreateAdjustV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/adjust?${queryStr}` : `/processing/v1/adjust`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * POST CreateAdjust V2
   */
  public async processingJobsCreateAdjustV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/adjust?${queryStr}` : `/processing/v2/adjust`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * POST CreateJobtypes V1
   */
  public async processingJobsCreateJobtypesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/jobtypes?${queryStr}` : `/processing/v1/jobtypes`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * POST CreateJobtypes V2
   */
  public async processingJobsCreateJobtypesV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes?${queryStr}` : `/processing/v2/jobtypes`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - ManageProcessingJobs
   *
   * POST CreateStart V1
   */
  public async processingJobsCreateStartV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/start?${queryStr}` : `/processing/v1/start`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Initiates new processing jobs at a Facility, including job details and associated packages.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * POST CreateStart V2
   */
  public async processingJobsCreateStartV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/start?${queryStr}` : `/processing/v2/start`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - ManageProcessingJobs
   *
   * POST Createpackages V1
   */
  public async processingJobsCreatepackagesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/createpackages?${queryStr}` : `/processing/v1/createpackages`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Creates packages from processing jobs at a Facility, including optional location and note assignments.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * POST Createpackages V2
   */
  public async processingJobsCreatepackagesV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/createpackages?${queryStr}` : `/processing/v2/createpackages`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * DELETE Delete V1
   */
  public async processingJobsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/${encodeURIComponent(id)}?${queryStr}` : `/processing/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/${encodeURIComponent(id)}?${queryStr}` : `/processing/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * DELETE DeleteJobtypes V1
   */
  public async processingJobsDeleteJobtypesV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/jobtypes/${encodeURIComponent(id)}?${queryStr}` : `/processing/v1/jobtypes/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes/${encodeURIComponent(id)}?${queryStr}` : `/processing/v2/jobtypes/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET Get V1
   */
  public async processingJobsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/${encodeURIComponent(id)}?${queryStr}` : `/processing/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/${encodeURIComponent(id)}?${queryStr}` : `/processing/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetActive V1
   */
  public async processingJobsGetActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/active?${queryStr}` : `/processing/v1/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/active?${queryStr}` : `/processing/v2/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetInactive V1
   */
  public async processingJobsGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/inactive?${queryStr}` : `/processing/v1/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/inactive?${queryStr}` : `/processing/v2/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetJobtypesActive V1
   */
  public async processingJobsGetJobtypesActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/jobtypes/active?${queryStr}` : `/processing/v1/jobtypes/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes/active?${queryStr}` : `/processing/v2/jobtypes/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetJobtypesAttributes V1
   */
  public async processingJobsGetJobtypesAttributesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/jobtypes/attributes?${queryStr}` : `/processing/v1/jobtypes/attributes`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes/attributes?${queryStr}` : `/processing/v2/jobtypes/attributes`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetJobtypesCategories V1
   */
  public async processingJobsGetJobtypesCategoriesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/jobtypes/categories?${queryStr}` : `/processing/v1/jobtypes/categories`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes/categories?${queryStr}` : `/processing/v2/jobtypes/categories`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * GET GetJobtypesInactive V1
   */
  public async processingJobsGetJobtypesInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/jobtypes/inactive?${queryStr}` : `/processing/v1/jobtypes/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes/inactive?${queryStr}` : `/processing/v2/jobtypes/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * PUT UpdateFinish V1
   */
  public async processingJobsUpdateFinishV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/finish?${queryStr}` : `/processing/v1/finish`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Completes processing jobs at a Facility by recording final notes and waste measurements.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * PUT UpdateFinish V2
   */
  public async processingJobsUpdateFinishV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/finish?${queryStr}` : `/processing/v2/finish`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * PUT UpdateJobtypes V1
   */
  public async processingJobsUpdateJobtypesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/jobtypes?${queryStr}` : `/processing/v1/jobtypes`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * PUT UpdateJobtypes V2
   */
  public async processingJobsUpdateJobtypesV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/jobtypes?${queryStr}` : `/processing/v2/jobtypes`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Processing Job
   *
   * PUT UpdateUnfinish V1
   */
  public async processingJobsUpdateUnfinishV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v1/unfinish?${queryStr}` : `/processing/v1/unfinish`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
   * 
   *   Permissions Required:
   *   - Manage Processing Job
   *
   * PUT UpdateUnfinish V2
   */
  public async processingJobsUpdateUnfinishV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/processing/v2/unfinish?${queryStr}` : `/processing/v2/unfinish`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Creates new sublocation records for a Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * POST Create V2
   */
  public async sublocationsCreateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sublocations/v2?${queryStr}` : `/sublocations/v2`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sublocations/v2/${encodeURIComponent(id)}?${queryStr}` : `/sublocations/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sublocations/v2/${encodeURIComponent(id)}?${queryStr}` : `/sublocations/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sublocations/v2/active?${queryStr}` : `/sublocations/v2/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sublocations/v2/inactive?${queryStr}` : `/sublocations/v2/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Updates existing sublocation records for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Locations
   *
   * PUT Update V2
   */
  public async sublocationsUpdateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sublocations/v2?${queryStr}` : `/sublocations/v2`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Sales Delivery
   *
   * POST CreateDelivery V1
   */
  public async salesCreateDeliveryV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries?${queryStr}` : `/sales/v1/deliveries`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async salesCreateDeliveryV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries?${queryStr}` : `/sales/v2/deliveries`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Retailer Delivery
   *
   * POST CreateDeliveryRetailer V1
   */
  public async salesCreateDeliveryRetailerV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/retailer?${queryStr}` : `/sales/v1/deliveries/retailer`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async salesCreateDeliveryRetailerV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer?${queryStr}` : `/sales/v2/deliveries/retailer`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Retailer Delivery
   *
   * POST CreateDeliveryRetailerDepart V1
   */
  public async salesCreateDeliveryRetailerDepartV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/depart?${queryStr}` : `/sales/v1/deliveries/retailer/depart`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async salesCreateDeliveryRetailerDepartV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/depart?${queryStr}` : `/sales/v2/deliveries/retailer/depart`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Retailer Delivery
   *
   * POST CreateDeliveryRetailerEnd V1
   */
  public async salesCreateDeliveryRetailerEndV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/end?${queryStr}` : `/sales/v1/deliveries/retailer/end`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async salesCreateDeliveryRetailerEndV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/end?${queryStr}` : `/sales/v2/deliveries/retailer/end`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Retailer Delivery
   *
   * POST CreateDeliveryRetailerRestock V1
   */
  public async salesCreateDeliveryRetailerRestockV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/restock?${queryStr}` : `/sales/v1/deliveries/retailer/restock`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async salesCreateDeliveryRetailerRestockV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/restock?${queryStr}` : `/sales/v2/deliveries/retailer/restock`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Retailer Delivery
   *
   * POST CreateDeliveryRetailerSale V1
   */
  public async salesCreateDeliveryRetailerSaleV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/sale?${queryStr}` : `/sales/v1/deliveries/retailer/sale`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async salesCreateDeliveryRetailerSaleV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/sale?${queryStr}` : `/sales/v2/deliveries/retailer/sale`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Sales
   *
   * POST CreateReceipt V1
   */
  public async salesCreateReceiptV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/receipts?${queryStr}` : `/sales/v1/receipts`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async salesCreateReceiptV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts?${queryStr}` : `/sales/v2/receipts`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * POST CreateTransactionByDate V1
   */
  public async salesCreateTransactionByDateV1(date: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/transactions/${encodeURIComponent(date)}?${queryStr}` : `/sales/v1/transactions/${encodeURIComponent(date)}`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales Delivery
   *
   * DELETE DeleteDelivery V1
   */
  public async salesDeleteDeliveryV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/${encodeURIComponent(id)}?${queryStr}` : `/sales/v1/deliveries/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/deliveries/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Retailer Delivery
   *
   * DELETE DeleteDeliveryRetailer V1
   */
  public async salesDeleteDeliveryRetailerV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/${encodeURIComponent(id)}?${queryStr}` : `/sales/v1/deliveries/retailer/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/deliveries/retailer/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * DELETE DeleteReceipt V1
   */
  public async salesDeleteReceiptV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/receipts/${encodeURIComponent(id)}?${queryStr}` : `/sales/v1/receipts/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/receipts/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetCounties V1
   */
  public async salesGetCountiesV1(body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/counties?${queryStr}` : `/sales/v1/counties`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/counties?${queryStr}` : `/sales/v2/counties`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetCustomertypes V1
   */
  public async salesGetCustomertypesV1(body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/customertypes?${queryStr}` : `/sales/v1/customertypes`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/customertypes?${queryStr}` : `/sales/v2/customertypes`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales Delivery
   *
   * GET GetDeliveriesActive V1
   */
  public async salesGetDeliveriesActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (salesDateEnd) query.append('salesDateEnd', salesDateEnd);
    if (salesDateStart) query.append('salesDateStart', salesDateStart);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/active?${queryStr}` : `/sales/v1/deliveries/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    if (salesDateEnd) query.append('salesDateEnd', salesDateEnd);
    if (salesDateStart) query.append('salesDateStart', salesDateStart);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/active?${queryStr}` : `/sales/v2/deliveries/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales Delivery
   *
   * GET GetDeliveriesInactive V1
   */
  public async salesGetDeliveriesInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (salesDateEnd) query.append('salesDateEnd', salesDateEnd);
    if (salesDateStart) query.append('salesDateStart', salesDateStart);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/inactive?${queryStr}` : `/sales/v1/deliveries/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    if (salesDateEnd) query.append('salesDateEnd', salesDateEnd);
    if (salesDateStart) query.append('salesDateStart', salesDateStart);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/inactive?${queryStr}` : `/sales/v2/deliveries/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Retailer Delivery
   *
   * GET GetDeliveriesRetailerActive V1
   */
  public async salesGetDeliveriesRetailerActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/active?${queryStr}` : `/sales/v1/deliveries/retailer/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/active?${queryStr}` : `/sales/v2/deliveries/retailer/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Retailer Delivery
   *
   * GET GetDeliveriesRetailerInactive V1
   */
  public async salesGetDeliveriesRetailerInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/inactive?${queryStr}` : `/sales/v1/deliveries/retailer/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/inactive?${queryStr}` : `/sales/v2/deliveries/retailer/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   -
   *
   * GET GetDeliveriesReturnreasons V1
   */
  public async salesGetDeliveriesReturnreasonsV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/returnreasons?${queryStr}` : `/sales/v1/deliveries/returnreasons`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/returnreasons?${queryStr}` : `/sales/v2/deliveries/returnreasons`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales Delivery
   *
   * GET GetDelivery V1
   */
  public async salesGetDeliveryV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/${encodeURIComponent(id)}?${queryStr}` : `/sales/v1/deliveries/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/deliveries/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Retailer Delivery
   *
   * GET GetDeliveryRetailer V1
   */
  public async salesGetDeliveryRetailerV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/retailer/${encodeURIComponent(id)}?${queryStr}` : `/sales/v1/deliveries/retailer/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/deliveries/retailer/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   -
   *
   * GET GetPatientRegistrationsLocations V1
   */
  public async salesGetPatientRegistrationsLocationsV1(body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/patientregistration/locations?${queryStr}` : `/sales/v1/patientregistration/locations`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/patientregistration/locations?${queryStr}` : `/sales/v2/patientregistration/locations`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales Delivery
   *
   * GET GetPaymenttypes V1
   */
  public async salesGetPaymenttypesV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/paymenttypes?${queryStr}` : `/sales/v1/paymenttypes`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/paymenttypes?${queryStr}` : `/sales/v2/paymenttypes`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * GET GetReceipt V1
   */
  public async salesGetReceiptV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/receipts/${encodeURIComponent(id)}?${queryStr}` : `/sales/v1/receipts/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/${encodeURIComponent(id)}?${queryStr}` : `/sales/v2/receipts/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * GET GetReceiptsActive V1
   */
  public async salesGetReceiptsActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (salesDateEnd) query.append('salesDateEnd', salesDateEnd);
    if (salesDateStart) query.append('salesDateStart', salesDateStart);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/receipts/active?${queryStr}` : `/sales/v1/receipts/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    if (salesDateEnd) query.append('salesDateEnd', salesDateEnd);
    if (salesDateStart) query.append('salesDateStart', salesDateStart);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/active?${queryStr}` : `/sales/v2/receipts/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/external/${encodeURIComponent(externalNumber)}?${queryStr}` : `/sales/v2/receipts/external/${encodeURIComponent(externalNumber)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * GET GetReceiptsInactive V1
   */
  public async salesGetReceiptsInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string, salesDateEnd?: string, salesDateStart?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (salesDateEnd) query.append('salesDateEnd', salesDateEnd);
    if (salesDateStart) query.append('salesDateStart', salesDateStart);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/receipts/inactive?${queryStr}` : `/sales/v1/receipts/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    if (salesDateEnd) query.append('salesDateEnd', salesDateEnd);
    if (salesDateStart) query.append('salesDateStart', salesDateStart);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/inactive?${queryStr}` : `/sales/v2/receipts/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * GET GetTransactions V1
   */
  public async salesGetTransactionsV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/transactions?${queryStr}` : `/sales/v1/transactions`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * GET GetTransactionsBySalesDateStartAndSalesDateEnd V1
   */
  public async salesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesDateStart: string, salesDateEnd: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/transactions/${encodeURIComponent(salesDateStart)}/${encodeURIComponent(salesDateEnd)}?${queryStr}` : `/sales/v1/transactions/${encodeURIComponent(salesDateStart)}/${encodeURIComponent(salesDateEnd)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Sales Delivery
   *
   * PUT UpdateDelivery V1
   */
  public async salesUpdateDeliveryV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries?${queryStr}` : `/sales/v1/deliveries`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Manage Sales Delivery
   *
   * PUT UpdateDelivery V2
   */
  public async salesUpdateDeliveryV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries?${queryStr}` : `/sales/v2/deliveries`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales Delivery
   *
   * PUT UpdateDeliveryComplete V1
   */
  public async salesUpdateDeliveryCompleteV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/complete?${queryStr}` : `/sales/v1/deliveries/complete`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
   * 
   *   Permissions Required:
   *   - Manage Sales Delivery
   *
   * PUT UpdateDeliveryComplete V2
   */
  public async salesUpdateDeliveryCompleteV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/complete?${queryStr}` : `/sales/v2/deliveries/complete`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Sales Delivery
   *
   * PUT UpdateDeliveryHub V1
   */
  public async salesUpdateDeliveryHubV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/hub?${queryStr}` : `/sales/v1/deliveries/hub`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Manage Sales Delivery, Manage Sales Delivery Hub
   *
   * PUT UpdateDeliveryHub V2
   */
  public async salesUpdateDeliveryHubV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/hub?${queryStr}` : `/sales/v2/deliveries/hub`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * PUT UpdateDeliveryHubAccept V1
   */
  public async salesUpdateDeliveryHubAcceptV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/hub/accept?${queryStr}` : `/sales/v1/deliveries/hub/accept`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
   * 
   *   Permissions Required:
   *   - Manage Sales Delivery Hub
   *
   * PUT UpdateDeliveryHubAccept V2
   */
  public async salesUpdateDeliveryHubAcceptV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/hub/accept?${queryStr}` : `/sales/v2/deliveries/hub/accept`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * PUT UpdateDeliveryHubDepart V1
   */
  public async salesUpdateDeliveryHubDepartV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/hub/depart?${queryStr}` : `/sales/v1/deliveries/hub/depart`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
   * 
   *   Permissions Required:
   *   - Manage Sales Delivery Hub
   *
   * PUT UpdateDeliveryHubDepart V2
   */
  public async salesUpdateDeliveryHubDepartV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/hub/depart?${queryStr}` : `/sales/v2/deliveries/hub/depart`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * PUT UpdateDeliveryHubVerifyID V1
   */
  public async salesUpdateDeliveryHubVerifyIdV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/hub/verifyID?${queryStr}` : `/sales/v1/deliveries/hub/verifyID`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
   * 
   *   Permissions Required:
   *   - Manage Sales Delivery Hub
   *
   * PUT UpdateDeliveryHubVerifyID V2
   */
  public async salesUpdateDeliveryHubVerifyIdV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/hub/verifyID?${queryStr}` : `/sales/v2/deliveries/hub/verifyID`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Retailer Delivery
   *
   * PUT UpdateDeliveryRetailer V1
   */
  public async salesUpdateDeliveryRetailerV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/deliveries/retailer?${queryStr}` : `/sales/v1/deliveries/retailer`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async salesUpdateDeliveryRetailerV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/deliveries/retailer?${queryStr}` : `/sales/v2/deliveries/retailer`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Sales
   *
   * PUT UpdateReceipt V1
   */
  public async salesUpdateReceiptV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/receipts?${queryStr}` : `/sales/v1/receipts`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
   * 
   *   Permissions Required:
   *   - Manage Sales
   *
   * PUT UpdateReceipt V2
   */
  public async salesUpdateReceiptV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts?${queryStr}` : `/sales/v2/receipts`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
   * 
   *   Permissions Required:
   *   - Manage Sales
   *
   * PUT UpdateReceiptFinalize V2
   */
  public async salesUpdateReceiptFinalizeV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/finalize?${queryStr}` : `/sales/v2/receipts/finalize`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
   * 
   *   Permissions Required:
   *   - Manage Sales
   *
   * PUT UpdateReceiptUnfinalize V2
   */
  public async salesUpdateReceiptUnfinalizeV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v2/receipts/unfinalize?${queryStr}` : `/sales/v2/receipts/unfinalize`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Sales
   *
   * PUT UpdateTransactionByDate V1
   */
  public async salesUpdateTransactionByDateV1(date: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/sales/v1/transactions/${encodeURIComponent(date)}?${queryStr}` : `/sales/v1/transactions/${encodeURIComponent(date)}`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Strains
   *
   * POST Create V1
   */
  public async strainsCreateV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v1/create?${queryStr}` : `/strains/v1/create`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Creates new strain records for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Strains
   *
   * POST Create V2
   */
  public async strainsCreateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v2?${queryStr}` : `/strains/v2`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Strains
   *
   * POST CreateUpdate V1
   */
  public async strainsCreateUpdateV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v1/update?${queryStr}` : `/strains/v1/update`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Strains
   *
   * DELETE Delete V1
   */
  public async strainsDeleteV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v1/${encodeURIComponent(id)}?${queryStr}` : `/strains/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v2/${encodeURIComponent(id)}?${queryStr}` : `/strains/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Strains
   *
   * GET Get V1
   */
  public async strainsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v1/${encodeURIComponent(id)}?${queryStr}` : `/strains/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v2/${encodeURIComponent(id)}?${queryStr}` : `/strains/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - Manage Strains
   *
   * GET GetActive V1
   */
  public async strainsGetActiveV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v1/active?${queryStr}` : `/strains/v1/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v2/active?${queryStr}` : `/strains/v2/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v2/inactive?${queryStr}` : `/strains/v2/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Updates existing strain records for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Strains
   *
   * PUT Update V2
   */
  public async strainsUpdateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/strains/v2?${queryStr}` : `/strains/v2`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/tags/v2/package/available?${queryStr}` : `/tags/v2/package/available`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/tags/v2/plant/available?${queryStr}` : `/tags/v2/plant/available`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/tags/v2/staged?${queryStr}` : `/tags/v2/staged`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Creates new additive templates for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Additives
   *
   * POST Create V2
   */
  public async additivesTemplatesCreateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/additivestemplates/v2?${queryStr}` : `/additivestemplates/v2`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/additivestemplates/v2/${encodeURIComponent(id)}?${queryStr}` : `/additivestemplates/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/additivestemplates/v2/active?${queryStr}` : `/additivestemplates/v2/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/additivestemplates/v2/inactive?${queryStr}` : `/additivestemplates/v2/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Updates existing additive templates for a specified Facility.
   * 
   *   Permissions Required:
   *   - Manage Additives
   *
   * PUT Update V2
   */
  public async additivesTemplatesUpdateV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/additivestemplates/v2?${queryStr}` : `/additivestemplates/v2`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *   - Finish/Discontinue Harvests
   *
   * POST CreateFinish V1
   */
  public async harvestsCreateFinishV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v1/finish?${queryStr}` : `/harvests/v1/finish`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async harvestsCreatePackageV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v1/create/packages?${queryStr}` : `/harvests/v1/create/packages`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async harvestsCreatePackageV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/packages?${queryStr}` : `/harvests/v2/packages`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async harvestsCreatePackageTestingV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v1/create/packages/testing?${queryStr}` : `/harvests/v1/create/packages/testing`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async harvestsCreatePackageTestingV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/packages/testing?${queryStr}` : `/harvests/v2/packages/testing`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *   - Manage Harvests
   *
   * POST CreateRemoveWaste V1
   */
  public async harvestsCreateRemoveWasteV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v1/removewaste?${queryStr}` : `/harvests/v1/removewaste`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *   - Finish/Discontinue Harvests
   *
   * POST CreateUnfinish V1
   */
  public async harvestsCreateUnfinishV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v1/unfinish?${queryStr}` : `/harvests/v1/unfinish`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async harvestsCreateWasteV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/waste?${queryStr}` : `/harvests/v2/waste`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/waste/${encodeURIComponent(id)}?${queryStr}` : `/harvests/v2/waste/${encodeURIComponent(id)}`;
    const { data } = await this.client.delete(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *
   * GET Get V1
   */
  public async harvestsGetV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v1/${encodeURIComponent(id)}?${queryStr}` : `/harvests/v1/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/${encodeURIComponent(id)}?${queryStr}` : `/harvests/v2/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *
   * GET GetActive V1
   */
  public async harvestsGetActiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v1/active?${queryStr}` : `/harvests/v1/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/active?${queryStr}` : `/harvests/v2/active`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *
   * GET GetInactive V1
   */
  public async harvestsGetInactiveV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v1/inactive?${queryStr}` : `/harvests/v1/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/inactive?${queryStr}` : `/harvests/v2/inactive`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *
   * GET GetOnhold V1
   */
  public async harvestsGetOnholdV1(body?: any, lastModifiedEnd?: string, lastModifiedStart?: string, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v1/onhold?${queryStr}` : `/harvests/v1/onhold`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (lastModifiedEnd) query.append('lastModifiedEnd', lastModifiedEnd);
    if (lastModifiedStart) query.append('lastModifiedStart', lastModifiedStart);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/onhold?${queryStr}` : `/harvests/v2/onhold`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (harvestId) query.append('harvestId', harvestId);
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/waste?${queryStr}` : `/harvests/v2/waste`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetWasteTypes V1
   */
  public async harvestsGetWasteTypesV1(body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v1/waste/types?${queryStr}` : `/harvests/v1/waste/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/waste/types?${queryStr}` : `/harvests/v2/waste/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
  public async harvestsUpdateFinishV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/finish?${queryStr}` : `/harvests/v2/finish`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async harvestsUpdateLocationV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/location?${queryStr}` : `/harvests/v2/location`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *   - Manage Harvests
   *
   * PUT UpdateMove V1
   */
  public async harvestsUpdateMoveV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v1/move?${queryStr}` : `/harvests/v1/move`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Harvests
   *   - Manage Harvests
   *
   * PUT UpdateRename V1
   */
  public async harvestsUpdateRenameV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v1/rename?${queryStr}` : `/harvests/v1/rename`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async harvestsUpdateRenameV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/rename?${queryStr}` : `/harvests/v2/rename`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async harvestsUpdateRestoreHarvestedPlantsV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/restore/harvestedplants?${queryStr}` : `/harvests/v2/restore/harvestedplants`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async harvestsUpdateUnfinishV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/harvests/v2/unfinish?${queryStr}` : `/harvests/v2/unfinish`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * POST CreateRecord V1
   */
  public async labTestsCreateRecordV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v1/record?${queryStr}` : `/labtests/v1/record`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
  public async labTestsCreateRecordV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/record?${queryStr}` : `/labtests/v2/record`;
    const { data } = await this.client.post(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/batches?${queryStr}` : `/labtests/v2/batches`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * GET GetLabtestdocument V1
   */
  public async labTestsGetLabtestdocumentV1(id: string, body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v1/labtestdocument/${encodeURIComponent(id)}?${queryStr}` : `/labtests/v1/labtestdocument/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/labtestdocument/${encodeURIComponent(id)}?${queryStr}` : `/labtests/v2/labtestdocument/${encodeURIComponent(id)}`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *
   * GET GetResults V1
   */
  public async labTestsGetResultsV1(body?: any, licenseNumber?: string, packageId?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (packageId) query.append('packageId', packageId);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v1/results?${queryStr}` : `/labtests/v1/results`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    if (packageId) query.append('packageId', packageId);
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/results?${queryStr}` : `/labtests/v2/results`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetStates V1
   */
  public async labTestsGetStatesV1(body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v1/states?${queryStr}` : `/labtests/v1/states`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/states?${queryStr}` : `/labtests/v2/states`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - None
   *
   * GET GetTypes V1
   */
  public async labTestsGetTypesV1(body?: any, No?: string): Promise<any> {
    const query = new URLSearchParams();
    if (No) query.append('No', No);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v1/types?${queryStr}` : `/labtests/v1/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
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
    const query = new URLSearchParams();
    if (pageNumber) query.append('pageNumber', pageNumber);
    if (pageSize) query.append('pageSize', pageSize);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/types?${queryStr}` : `/labtests/v2/types`;
    const { data } = await this.client.get(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateLabtestdocument V1
   */
  public async labTestsUpdateLabtestdocumentV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v1/labtestdocument?${queryStr}` : `/labtests/v1/labtestdocument`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async labTestsUpdateLabtestdocumentV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/labtestdocument?${queryStr}` : `/labtests/v2/labtestdocument`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

  /**
   * Permissions Required:
   *   - View Packages
   *   - Manage Packages Inventory
   *
   * PUT UpdateResultRelease V1
   */
  public async labTestsUpdateResultReleaseV1(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v1/results/release?${queryStr}` : `/labtests/v1/results/release`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
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
  public async labTestsUpdateResultReleaseV2(body?: any, licenseNumber?: string): Promise<any> {
    const query = new URLSearchParams();
    if (licenseNumber) query.append('licenseNumber', licenseNumber);
    const queryStr = query.toString();
    const fullUrl = queryStr ? `/labtests/v2/results/release?${queryStr}` : `/labtests/v2/results/release`;
    const { data } = await this.client.put(fullUrl, body);
    return data;
  }

}

// --- Request Models ---
