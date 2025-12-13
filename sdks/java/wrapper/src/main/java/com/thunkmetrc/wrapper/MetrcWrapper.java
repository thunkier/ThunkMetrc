package com.thunkmetrc.wrapper;

import com.thunkmetrc.client.MetrcClient;
import java.io.IOException;
import java.util.List;

/**
 * Metrc Wrapper with rate limiting.
 */
public class MetrcWrapper {
    private final MetrcClient client;
    private final RateLimiter rateLimiter;

    /**
     * Constructor.
     * @param client MetrcClient instance
     */
    public MetrcWrapper(MetrcClient client) {
        this(client, new RateLimiter.Config());
    }

    /**
     * Constructor.
     * @param client MetrcClient instance
     * @param config RateLimiter configuration
     */
    public MetrcWrapper(MetrcClient client, RateLimiter.Config config) {
        this.client = client;
        this.rateLimiter = new RateLimiter(config);
    }

    public static class SublocationsCreateV2RequestItem {
        public String Name;
    }

    /**
     * Creates new sublocation records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * POST Create V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object sublocationsCreateV2(String licensenumber, List<SublocationsCreateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.sublocationsCreateV2(licensenumber, body));
    }

    /**
     * Archives an existing Sublocation record for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * DELETE Delete V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object sublocationsDeleteV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.sublocationsDeleteV2(id, licensenumber));
    }

    /**
     * Retrieves a Sublocation by its Id, with an optional license number.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET Get V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object sublocationsGetV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.sublocationsGetV2(id, licensenumber));
    }

    /**
     * Retrieves a list of active sublocations for the current Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object sublocationsGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.sublocationsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves a list of inactive sublocations for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object sublocationsGetInactiveV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.sublocationsGetInactiveV2(licensenumber, pagenumber, pagesize));
    }

    public static class SublocationsUpdateV2RequestItem {
        public Integer Id;
        public String Name;
    }

    /**
     * Updates existing sublocation records for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * PUT Update V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object sublocationsUpdateV2(String licensenumber, List<SublocationsUpdateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.sublocationsUpdateV2(licensenumber, body));
    }

    public static class TransfersCreateExternalIncomingV1RequestItem {
        public List<TransfersCreateExternalIncomingV1RequestItem_Destinations> Destinations;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String PhoneNumberForQuestions;
        public String ShipperAddress1;
        public String ShipperAddress2;
        public String ShipperAddressCity;
        public String ShipperAddressPostalCode;
        public String ShipperAddressState;
        public String ShipperLicenseNumber;
        public String ShipperMainPhoneNumber;
        public String ShipperName;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class TransfersCreateExternalIncomingV1RequestItem_Destinations {
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Integer GrossUnitOfWeightId;
        public Double GrossWeight;
        public String InvoiceNumber;
        public List<TransfersCreateExternalIncomingV1RequestItem_Destinations_Packages> Packages;
        public String PlannedRoute;
        public String RecipientLicenseNumber;
        public String TransferTypeName;
        public List<TransfersCreateExternalIncomingV1RequestItem_Destinations_Transporters> Transporters;
    }

    public static class TransfersCreateExternalIncomingV1RequestItem_Destinations_Packages {
        public String ExpirationDate;
        public String ExternalId;
        public String GrossUnitOfWeightName;
        public Double GrossWeight;
        public String ItemName;
        public String PackagedDate;
        public Integer Quantity;
        public String SellByDate;
        public String UnitOfMeasureName;
        public String UseByDate;
        public String WholesalePrice;
    }

    public static class TransfersCreateExternalIncomingV1RequestItem_Destinations_Transporters {
        public String DriverLayoverLeg;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Boolean IsLayover;
        public String PhoneNumberForQuestions;
        public String TransporterDetails;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * POST CreateExternalIncoming V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersCreateExternalIncomingV1(String licensenumber, List<TransfersCreateExternalIncomingV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transfersCreateExternalIncomingV1(licensenumber, body));
    }

    public static class TransfersCreateExternalIncomingV2RequestItem {
        public List<TransfersCreateExternalIncomingV2RequestItem_Destinations> Destinations;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String PhoneNumberForQuestions;
        public String ShipperAddress1;
        public String ShipperAddress2;
        public String ShipperAddressCity;
        public String ShipperAddressPostalCode;
        public String ShipperAddressState;
        public String ShipperLicenseNumber;
        public String ShipperMainPhoneNumber;
        public String ShipperName;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class TransfersCreateExternalIncomingV2RequestItem_Destinations {
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Integer GrossUnitOfWeightId;
        public Double GrossWeight;
        public String InvoiceNumber;
        public List<TransfersCreateExternalIncomingV2RequestItem_Destinations_Packages> Packages;
        public String PlannedRoute;
        public String RecipientLicenseNumber;
        public String TransferTypeName;
        public List<TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters> Transporters;
    }

    public static class TransfersCreateExternalIncomingV2RequestItem_Destinations_Packages {
        public String ExpirationDate;
        public String ExternalId;
        public String GrossUnitOfWeightName;
        public Double GrossWeight;
        public String ItemName;
        public String PackagedDate;
        public Integer Quantity;
        public String SellByDate;
        public String UnitOfMeasureName;
        public String UseByDate;
        public String WholesalePrice;
    }

    public static class TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters {
        public String DriverLayoverLeg;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Boolean IsLayover;
        public String PhoneNumberForQuestions;
        public List<TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails> TransporterDetails;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class TransfersCreateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails {
        public String DriverLayoverLeg;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    /**
     * Creates external incoming shipment plans for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *
     * POST CreateExternalIncoming V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersCreateExternalIncomingV2(String licensenumber, List<TransfersCreateExternalIncomingV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transfersCreateExternalIncomingV2(licensenumber, body));
    }

    public static class TransfersCreateTemplatesV1RequestItem {
        public List<TransfersCreateTemplatesV1RequestItem_Destinations> Destinations;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String Name;
        public String PhoneNumberForQuestions;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class TransfersCreateTemplatesV1RequestItem_Destinations {
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public String InvoiceNumber;
        public List<TransfersCreateTemplatesV1RequestItem_Destinations_Packages> Packages;
        public String PlannedRoute;
        public String RecipientLicenseNumber;
        public String TransferTypeName;
        public List<TransfersCreateTemplatesV1RequestItem_Destinations_Transporters> Transporters;
    }

    public static class TransfersCreateTemplatesV1RequestItem_Destinations_Packages {
        public String GrossUnitOfWeightName;
        public Double GrossWeight;
        public String PackageLabel;
        public String WholesalePrice;
    }

    public static class TransfersCreateTemplatesV1RequestItem_Destinations_Transporters {
        public String DriverLayoverLeg;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Boolean IsLayover;
        public String PhoneNumberForQuestions;
        public String TransporterDetails;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * POST CreateTemplates V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersCreateTemplatesV1(String licensenumber, List<TransfersCreateTemplatesV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transfersCreateTemplatesV1(licensenumber, body));
    }

    public static class TransfersCreateTemplatesOutgoingV2RequestItem {
        public List<TransfersCreateTemplatesOutgoingV2RequestItem_Destinations> Destinations;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String Name;
        public String PhoneNumberForQuestions;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class TransfersCreateTemplatesOutgoingV2RequestItem_Destinations {
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public String InvoiceNumber;
        public List<TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Packages> Packages;
        public String PlannedRoute;
        public String RecipientLicenseNumber;
        public String TransferTypeName;
        public List<TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters> Transporters;
    }

    public static class TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Packages {
        public String GrossUnitOfWeightName;
        public Double GrossWeight;
        public String PackageLabel;
        public String WholesalePrice;
    }

    public static class TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters {
        public String DriverLayoverLeg;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Boolean IsLayover;
        public String PhoneNumberForQuestions;
        public List<TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails> TransporterDetails;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class TransfersCreateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails {
        public String DriverLayoverLeg;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    /**
     * Creates new transfer templates for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *
     * POST CreateTemplatesOutgoing V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersCreateTemplatesOutgoingV2(String licensenumber, List<TransfersCreateTemplatesOutgoingV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transfersCreateTemplatesOutgoingV2(licensenumber, body));
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * DELETE DeleteExternalIncoming V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersDeleteExternalIncomingV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transfersDeleteExternalIncomingV1(id, licensenumber));
    }

    /**
     * Voids an external incoming shipment plan for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *
     * DELETE DeleteExternalIncoming V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersDeleteExternalIncomingV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transfersDeleteExternalIncomingV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * DELETE DeleteTemplates V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersDeleteTemplatesV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transfersDeleteTemplatesV1(id, licensenumber));
    }

    /**
     * Archives a transfer template for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *
     * DELETE DeleteTemplatesOutgoing V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersDeleteTemplatesOutgoingV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transfersDeleteTemplatesOutgoingV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetDeliveriesPackagesStates V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveriesPackagesStatesV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveriesPackagesStatesV1(no));
    }

    /**
     * Returns a list of available shipment Package states.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetDeliveriesPackagesStates V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveriesPackagesStatesV2(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveriesPackagesStatesV2(no));
    }

    /**
     * Please note: that the {id} parameter above represents a Shipment Plan ID.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDelivery V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryV1(String id, String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveryV1(id, no));
    }

    /**
     * Retrieves a list of shipment deliveries for a given Transfer Id. Please note: The {id} parameter above represents a Transfer Id.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDelivery V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryV2(String id, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveryV2(id, pagenumber, pagesize));
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryPackage V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryPackageV1(String id, String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveryPackageV1(id, no));
    }

    /**
     * Retrieves a list of packages associated with a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryPackage V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryPackageV2(String id, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveryPackageV2(id, pagenumber, pagesize));
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Delivery Package ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryPackageRequiredlabtestbatches V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryPackageRequiredlabtestbatchesV1(String id, String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveryPackageRequiredlabtestbatchesV1(id, no));
    }

    /**
     * Retrieves a list of required lab test batches for a given Transfer Delivery Package Id. Please note: The {id} parameter above represents a Transfer Delivery Package Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryPackageRequiredlabtestbatches V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryPackageRequiredlabtestbatchesV2(String id, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveryPackageRequiredlabtestbatchesV2(id, pagenumber, pagesize));
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryPackageWholesale V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryPackageWholesaleV1(String id, String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveryPackageWholesaleV1(id, no));
    }

    /**
     * Retrieves a list of wholesale shipment packages for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryPackageWholesale V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryPackageWholesaleV2(String id, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveryPackageWholesaleV2(id, pagenumber, pagesize));
    }

    /**
     * Please note: that the {id} parameter above represents a Shipment Delivery ID.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryTransporters V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryTransportersV1(String id, String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveryTransportersV1(id, no));
    }

    /**
     * Retrieves a list of transporters for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryTransporters V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryTransportersV2(String id, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveryTransportersV2(id, pagenumber, pagesize));
    }

    /**
     * Please note: The {id} parameter above represents a Shipment Delivery ID.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetDeliveryTransportersDetails V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryTransportersDetailsV1(String id, String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveryTransportersDetailsV1(id, no));
    }

    /**
     * Retrieves a list of transporter details for a given Transfer Delivery Id. Please note: The {id} parameter above represents a Transfer Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetDeliveryTransportersDetails V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetDeliveryTransportersDetailsV2(String id, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetDeliveryTransportersDetailsV2(id, pagenumber, pagesize));
    }

    /**
     * Retrieves a list of transfer hub shipments for a Facility, filtered by either last modified or estimated arrival date range.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetHub V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetHubV2(String estimatedarrivalend, String estimatedarrivalstart, String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetHubV2(estimatedarrivalend, estimatedarrivalstart, lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * GET GetIncoming V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetIncomingV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetIncomingV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves a list of incoming shipments for a Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetIncoming V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetIncomingV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetIncomingV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * GET GetOutgoing V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetOutgoingV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetOutgoingV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves a list of outgoing shipments for a Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetOutgoing V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetOutgoingV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetOutgoingV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * GET GetRejected V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetRejectedV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetRejectedV1(licensenumber));
    }

    /**
     * Retrieves a list of shipments with rejected packages for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *   - View Transfers
     *
     * GET GetRejected V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetRejectedV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetRejectedV2(licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplates V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetTemplatesV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplatesDelivery V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesDeliveryV1(String id, String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetTemplatesDeliveryV1(id, no));
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfers
     *
     * GET GetTemplatesDeliveryPackage V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesDeliveryPackageV1(String id, String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetTemplatesDeliveryPackageV1(id, no));
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplatesDeliveryTransporters V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesDeliveryTransportersV1(String id, String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetTemplatesDeliveryTransportersV1(id, no));
    }

    /**
     * Please note: The {id} parameter above represents a Transfer Template Delivery ID, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Transfer Templates
     *
     * GET GetTemplatesDeliveryTransportersDetails V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesDeliveryTransportersDetailsV1(String id, String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetTemplatesDeliveryTransportersDetailsV1(id, no));
    }

    /**
     * Retrieves a list of transfer templates for a Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoing V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesOutgoingV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetTemplatesOutgoingV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves a list of deliveries associated with a specific transfer template.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDelivery V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesOutgoingDeliveryV2(String id, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetTemplatesOutgoingDeliveryV2(id, pagenumber, pagesize));
    }

    /**
     * Retrieves a list of delivery package templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDeliveryPackage V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesOutgoingDeliveryPackageV2(String id, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetTemplatesOutgoingDeliveryPackageV2(id, pagenumber, pagesize));
    }

    /**
     * Retrieves a list of transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDeliveryTransporters V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesOutgoingDeliveryTransportersV2(String id, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetTemplatesOutgoingDeliveryTransportersV2(id, pagenumber, pagesize));
    }

    /**
     * Retrieves detailed transporter templates for a given Transfer Template Delivery Id. Please note: The {id} parameter above represents a Transfer Template Delivery Id, not a Manifest Number.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *   - View Transfer Templates
     *
     * GET GetTemplatesOutgoingDeliveryTransportersDetails V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(String id, String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetTemplatesOutgoingDeliveryTransportersDetailsV2(id, no));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetTypesV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetTypesV1(licensenumber));
    }

    /**
     * Retrieves a list of available transfer types for a Facility based on its license number.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersGetTypesV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transfersGetTypesV2(licensenumber, pagenumber, pagesize));
    }

    public static class TransfersUpdateExternalIncomingV1RequestItem {
        public List<TransfersUpdateExternalIncomingV1RequestItem_Destinations> Destinations;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String PhoneNumberForQuestions;
        public String ShipperAddress1;
        public String ShipperAddress2;
        public String ShipperAddressCity;
        public String ShipperAddressPostalCode;
        public String ShipperAddressState;
        public String ShipperLicenseNumber;
        public String ShipperMainPhoneNumber;
        public String ShipperName;
        public Integer TransferId;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class TransfersUpdateExternalIncomingV1RequestItem_Destinations {
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Integer GrossUnitOfWeightId;
        public Double GrossWeight;
        public String InvoiceNumber;
        public List<TransfersUpdateExternalIncomingV1RequestItem_Destinations_Packages> Packages;
        public String PlannedRoute;
        public String RecipientLicenseNumber;
        public Integer TransferDestinationId;
        public String TransferTypeName;
        public List<TransfersUpdateExternalIncomingV1RequestItem_Destinations_Transporters> Transporters;
    }

    public static class TransfersUpdateExternalIncomingV1RequestItem_Destinations_Packages {
        public String ExpirationDate;
        public String ExternalId;
        public String GrossUnitOfWeightName;
        public Double GrossWeight;
        public String ItemName;
        public String PackagedDate;
        public Integer Quantity;
        public String SellByDate;
        public String UnitOfMeasureName;
        public String UseByDate;
        public String WholesalePrice;
    }

    public static class TransfersUpdateExternalIncomingV1RequestItem_Destinations_Transporters {
        public String DriverLayoverLeg;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Boolean IsLayover;
        public String PhoneNumberForQuestions;
        public String TransporterDetails;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    /**
     * Permissions Required:
     *   - Transfers
     *
     * PUT UpdateExternalIncoming V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersUpdateExternalIncomingV1(String licensenumber, List<TransfersUpdateExternalIncomingV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transfersUpdateExternalIncomingV1(licensenumber, body));
    }

    public static class TransfersUpdateExternalIncomingV2RequestItem {
        public List<TransfersUpdateExternalIncomingV2RequestItem_Destinations> Destinations;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String PhoneNumberForQuestions;
        public String ShipperAddress1;
        public String ShipperAddress2;
        public String ShipperAddressCity;
        public String ShipperAddressPostalCode;
        public String ShipperAddressState;
        public String ShipperLicenseNumber;
        public String ShipperMainPhoneNumber;
        public String ShipperName;
        public Integer TransferId;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class TransfersUpdateExternalIncomingV2RequestItem_Destinations {
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Integer GrossUnitOfWeightId;
        public Double GrossWeight;
        public String InvoiceNumber;
        public List<TransfersUpdateExternalIncomingV2RequestItem_Destinations_Packages> Packages;
        public String PlannedRoute;
        public String RecipientLicenseNumber;
        public Integer TransferDestinationId;
        public String TransferTypeName;
        public List<TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters> Transporters;
    }

    public static class TransfersUpdateExternalIncomingV2RequestItem_Destinations_Packages {
        public String ExpirationDate;
        public String ExternalId;
        public String GrossUnitOfWeightName;
        public Double GrossWeight;
        public String ItemName;
        public String PackagedDate;
        public Integer Quantity;
        public String SellByDate;
        public String UnitOfMeasureName;
        public String UseByDate;
        public String WholesalePrice;
    }

    public static class TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters {
        public String DriverLayoverLeg;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Boolean IsLayover;
        public String PhoneNumberForQuestions;
        public List<TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails> TransporterDetails;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class TransfersUpdateExternalIncomingV2RequestItem_Destinations_Transporters_TransporterDetails {
        public String DriverLayoverLeg;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    /**
     * Updates external incoming shipment plans for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfers
     *
     * PUT UpdateExternalIncoming V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersUpdateExternalIncomingV2(String licensenumber, List<TransfersUpdateExternalIncomingV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transfersUpdateExternalIncomingV2(licensenumber, body));
    }

    public static class TransfersUpdateTemplatesV1RequestItem {
        public List<TransfersUpdateTemplatesV1RequestItem_Destinations> Destinations;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String Name;
        public String PhoneNumberForQuestions;
        public Integer TransferTemplateId;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class TransfersUpdateTemplatesV1RequestItem_Destinations {
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public String InvoiceNumber;
        public List<TransfersUpdateTemplatesV1RequestItem_Destinations_Packages> Packages;
        public String PlannedRoute;
        public String RecipientLicenseNumber;
        public Integer TransferDestinationId;
        public String TransferTypeName;
        public List<TransfersUpdateTemplatesV1RequestItem_Destinations_Transporters> Transporters;
    }

    public static class TransfersUpdateTemplatesV1RequestItem_Destinations_Packages {
        public String GrossUnitOfWeightName;
        public Double GrossWeight;
        public String PackageLabel;
        public String WholesalePrice;
    }

    public static class TransfersUpdateTemplatesV1RequestItem_Destinations_Transporters {
        public String DriverLayoverLeg;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Boolean IsLayover;
        public String PhoneNumberForQuestions;
        public String TransporterDetails;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    /**
     * Permissions Required:
     *   - Transfer Templates
     *
     * PUT UpdateTemplates V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersUpdateTemplatesV1(String licensenumber, List<TransfersUpdateTemplatesV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transfersUpdateTemplatesV1(licensenumber, body));
    }

    public static class TransfersUpdateTemplatesOutgoingV2RequestItem {
        public List<TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations> Destinations;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String Name;
        public String PhoneNumberForQuestions;
        public Integer TransferTemplateId;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations {
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public String InvoiceNumber;
        public List<TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Packages> Packages;
        public String PlannedRoute;
        public String RecipientLicenseNumber;
        public Integer TransferDestinationId;
        public String TransferTypeName;
        public List<TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters> Transporters;
    }

    public static class TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Packages {
        public String GrossUnitOfWeightName;
        public Double GrossWeight;
        public String PackageLabel;
        public String WholesalePrice;
    }

    public static class TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters {
        public String DriverLayoverLeg;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Boolean IsLayover;
        public String PhoneNumberForQuestions;
        public List<TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails> TransporterDetails;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class TransfersUpdateTemplatesOutgoingV2RequestItem_Destinations_Transporters_TransporterDetails {
        public String DriverLayoverLeg;
        public String DriverLicenseNumber;
        public String DriverName;
        public String DriverOccupationalLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    /**
     * Updates existing transfer templates for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transfer Templates
     *
     * PUT UpdateTemplatesOutgoing V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transfersUpdateTemplatesOutgoingV2(String licensenumber, List<TransfersUpdateTemplatesOutgoingV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transfersUpdateTemplatesOutgoingV2(licensenumber, body));
    }

    public static class PlantsCreateAdditivesV1RequestItem {
        public List<PlantsCreateAdditivesV1RequestItem_ActiveIngredients> ActiveIngredients;
        public String ActualDate;
        public String AdditiveType;
        public String ApplicationDevice;
        public String EpaRegistrationNumber;
        public List<String> PlantLabels;
        public String ProductSupplier;
        public String ProductTradeName;
        public Integer TotalAmountApplied;
        public String TotalAmountUnitOfMeasure;
    }

    public static class PlantsCreateAdditivesV1RequestItem_ActiveIngredients {
        public String Name;
        public Double Percentage;
    }

    /**
     * Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreateAdditivesV1(String licensenumber, List<PlantsCreateAdditivesV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreateAdditivesV1(licensenumber, body));
    }

    public static class PlantsCreateAdditivesV2RequestItem {
        public List<PlantsCreateAdditivesV2RequestItem_ActiveIngredients> ActiveIngredients;
        public String ActualDate;
        public String AdditiveType;
        public String ApplicationDevice;
        public String EpaRegistrationNumber;
        public List<String> PlantLabels;
        public String ProductSupplier;
        public String ProductTradeName;
        public Integer TotalAmountApplied;
        public String TotalAmountUnitOfMeasure;
    }

    public static class PlantsCreateAdditivesV2RequestItem_ActiveIngredients {
        public String Name;
        public Double Percentage;
    }

    /**
     * Records additive usage details applied to specific plants at a Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreateAdditivesV2(String licensenumber, List<PlantsCreateAdditivesV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreateAdditivesV2(licensenumber, body));
    }

    public static class PlantsCreateAdditivesBylocationV1RequestItem {
        public List<PlantsCreateAdditivesBylocationV1RequestItem_ActiveIngredients> ActiveIngredients;
        public String ActualDate;
        public String AdditiveType;
        public String ApplicationDevice;
        public String EpaRegistrationNumber;
        public String LocationName;
        public String ProductSupplier;
        public String ProductTradeName;
        public String SublocationName;
        public Integer TotalAmountApplied;
        public String TotalAmountUnitOfMeasure;
    }

    public static class PlantsCreateAdditivesBylocationV1RequestItem_ActiveIngredients {
        public String Name;
        public Double Percentage;
    }

    /**
     * Permissions Required:
     *   - Manage Plants
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesBylocation V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreateAdditivesBylocationV1(String licensenumber, List<PlantsCreateAdditivesBylocationV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreateAdditivesBylocationV1(licensenumber, body));
    }

    public static class PlantsCreateAdditivesBylocationV2RequestItem {
        public List<PlantsCreateAdditivesBylocationV2RequestItem_ActiveIngredients> ActiveIngredients;
        public String ActualDate;
        public String AdditiveType;
        public String ApplicationDevice;
        public String EpaRegistrationNumber;
        public String LocationName;
        public String ProductSupplier;
        public String ProductTradeName;
        public String SublocationName;
        public Integer TotalAmountApplied;
        public String TotalAmountUnitOfMeasure;
    }

    public static class PlantsCreateAdditivesBylocationV2RequestItem_ActiveIngredients {
        public String Name;
        public Double Percentage;
    }

    /**
     * Records additive usage for plants based on their location within a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesBylocation V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreateAdditivesBylocationV2(String licensenumber, List<PlantsCreateAdditivesBylocationV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreateAdditivesBylocationV2(licensenumber, body));
    }

    public static class PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem {
        public String ActualDate;
        public String AdditivesTemplateName;
        public String LocationName;
        public String Rate;
        public String SublocationName;
        public Integer TotalAmountApplied;
        public String TotalAmountUnitOfMeasure;
        public String Volume;
    }

    /**
     * Records additive usage for plants by location using a predefined additive template at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesBylocationUsingtemplate V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreateAdditivesBylocationUsingtemplateV2(String licensenumber, List<PlantsCreateAdditivesBylocationUsingtemplateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreateAdditivesBylocationUsingtemplateV2(licensenumber, body));
    }

    public static class PlantsCreateAdditivesUsingtemplateV2RequestItem {
        public String ActualDate;
        public String AdditivesTemplateName;
        public List<String> PlantLabels;
        public String Rate;
        public Integer TotalAmountApplied;
        public String TotalAmountUnitOfMeasure;
        public String Volume;
    }

    /**
     * Records additive usage for plants using predefined additive templates at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesUsingtemplate V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreateAdditivesUsingtemplateV2(String licensenumber, List<PlantsCreateAdditivesUsingtemplateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreateAdditivesUsingtemplateV2(licensenumber, body));
    }

    public static class PlantsCreateChangegrowthphasesV1RequestItem {
        public String GrowthDate;
        public String GrowthPhase;
        public Integer Id;
        public String Label;
        public String NewLocation;
        public String NewSublocation;
        public String NewTag;
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateChangegrowthphases V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreateChangegrowthphasesV1(String licensenumber, List<PlantsCreateChangegrowthphasesV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreateChangegrowthphasesV1(licensenumber, body));
    }

    public static class PlantsCreateHarvestplantsV1RequestItem {
        public String ActualDate;
        public String DryingLocation;
        public String DryingSublocation;
        public String HarvestName;
        public String PatientLicenseNumber;
        public String Plant;
        public String UnitOfWeight;
        public Double Weight;
    }

    /**
     * NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * POST CreateHarvestplants V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreateHarvestplantsV1(String licensenumber, List<PlantsCreateHarvestplantsV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreateHarvestplantsV1(licensenumber, body));
    }

    public static class PlantsCreateManicureV2RequestItem {
        public String ActualDate;
        public String DryingLocation;
        public String DryingSublocation;
        public String HarvestName;
        public String PatientLicenseNumber;
        public String Plant;
        public Integer PlantCount;
        public String UnitOfWeight;
        public Double Weight;
    }

    /**
     * Creates harvest product records from plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * POST CreateManicure V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreateManicureV2(String licensenumber, List<PlantsCreateManicureV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreateManicureV2(licensenumber, body));
    }

    public static class PlantsCreateManicureplantsV1RequestItem {
        public String ActualDate;
        public String DryingLocation;
        public String DryingSublocation;
        public String HarvestName;
        public String PatientLicenseNumber;
        public String Plant;
        public Integer PlantCount;
        public String UnitOfWeight;
        public Double Weight;
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * POST CreateManicureplants V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreateManicureplantsV1(String licensenumber, List<PlantsCreateManicureplantsV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreateManicureplantsV1(licensenumber, body));
    }

    public static class PlantsCreateMoveplantsV1RequestItem {
        public String ActualDate;
        public Integer Id;
        public String Label;
        public String Location;
        public String Sublocation;
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateMoveplants V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreateMoveplantsV1(String licensenumber, List<PlantsCreateMoveplantsV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreateMoveplantsV1(licensenumber, body));
    }

    public static class PlantsCreatePlantbatchPackageV1RequestItem {
        public String ActualDate;
        public Integer Count;
        public Boolean IsDonation;
        public Boolean IsTradeSample;
        public String Item;
        public String Location;
        public String Note;
        public String PackageTag;
        public String PatientLicenseNumber;
        public String PlantBatchType;
        public String PlantLabel;
        public String Sublocation;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreatePlantbatchPackageV1(String licensenumber, List<PlantsCreatePlantbatchPackageV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreatePlantbatchPackageV1(licensenumber, body));
    }

    public static class PlantsCreatePlantbatchPackageV2RequestItem {
        public String ActualDate;
        public Integer Count;
        public Boolean IsDonation;
        public Boolean IsTradeSample;
        public String Item;
        public String Location;
        public String Note;
        public String PackageTag;
        public String PatientLicenseNumber;
        public String PlantBatchType;
        public String PlantLabel;
        public String Sublocation;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreatePlantbatchPackageV2(String licensenumber, List<PlantsCreatePlantbatchPackageV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreatePlantbatchPackageV2(licensenumber, body));
    }

    public static class PlantsCreatePlantingsV1RequestItem {
        public String ActualDate;
        public String LocationName;
        public String PatientLicenseNumber;
        public String PlantBatchName;
        public String PlantBatchType;
        public Integer PlantCount;
        public String PlantLabel;
        public String StrainName;
        public String SublocationName;
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreatePlantings V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreatePlantingsV1(String licensenumber, List<PlantsCreatePlantingsV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreatePlantingsV1(licensenumber, body));
    }

    public static class PlantsCreatePlantingsV2RequestItem {
        public String ActualDate;
        public String LocationName;
        public String PatientLicenseNumber;
        public String PlantBatchName;
        public String PlantBatchType;
        public Integer PlantCount;
        public String PlantLabel;
        public String StrainName;
        public String SublocationName;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreatePlantingsV2(String licensenumber, List<PlantsCreatePlantingsV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreatePlantingsV2(licensenumber, body));
    }

    public static class PlantsCreateWasteV1RequestItem {
        public String LocationName;
        public String MixedMaterial;
        public String Note;
        public List<Object> PlantLabels;
        public String ReasonName;
        public String SublocationName;
        public String UnitOfMeasureName;
        public String WasteDate;
        public String WasteMethodName;
        public Double WasteWeight;
    }

    /**
     * Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreateWasteV1(String licensenumber, List<PlantsCreateWasteV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreateWasteV1(licensenumber, body));
    }

    public static class PlantsCreateWasteV2RequestItem {
        public String LocationName;
        public String MixedMaterial;
        public String Note;
        public List<Object> PlantLabels;
        public String ReasonName;
        public String SublocationName;
        public String UnitOfMeasureName;
        public String WasteDate;
        public String WasteMethodName;
        public Double WasteWeight;
    }

    /**
     * Records waste events for plants at a Facility, including method, reason, and location details.
     * 
     *   Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsCreateWasteV2(String licensenumber, List<PlantsCreateWasteV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsCreateWasteV2(licensenumber, body));
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *   - Destroy Veg/Flower Plants
     *
     * DELETE Delete V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsDeleteV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsDeleteV1(licensenumber));
    }

    /**
     * Removes plants from a Facility’s inventory while recording the reason for their disposal.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Destroy Veg/Flower Plants
     *
     * DELETE Delete V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsDeleteV2(String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsDeleteV2(licensenumber));
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET Get V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetV1(id, licensenumber));
    }

    /**
     * Retrieves a Plant by Id.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET Get V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - View/Manage Plants Additives
     *
     * GET GetAdditives V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetAdditivesV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetAdditivesV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves additive records applied to plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View/Manage Plants Additives
     *
     * GET GetAdditives V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetAdditivesV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetAdditivesV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   -
     *
     * GET GetAdditivesTypes V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetAdditivesTypesV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetAdditivesTypesV1(no));
    }

    /**
     * Retrieves a list of all plant additive types defined within a Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAdditivesTypes V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetAdditivesTypesV2(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetAdditivesTypesV2(no));
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetByLabel V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetByLabelV1(String label, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetByLabelV1(label, licensenumber));
    }

    /**
     * Retrieves a Plant by label.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetByLabel V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetByLabelV2(String label, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetByLabelV2(label, licensenumber));
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetFlowering V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetFloweringV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetFloweringV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves flowering-phase plants at a specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetFlowering V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetFloweringV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetFloweringV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetGrowthPhases V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetGrowthPhasesV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetGrowthPhasesV1(licensenumber));
    }

    /**
     * Retrieves the list of growth phases supported by a specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetGrowthPhases V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetGrowthPhasesV2(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetGrowthPhasesV2(licensenumber));
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetInactive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves inactive plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves mother-phase plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Mother Plants
     *
     * GET GetMother V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetMotherV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetMotherV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves inactive mother-phase plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Mother Plants
     *
     * GET GetMotherInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetMotherInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetMotherInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves mother-phase plants currently marked as on hold at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Mother Plants
     *
     * GET GetMotherOnhold V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetMotherOnholdV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetMotherOnholdV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetOnhold V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetOnholdV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetOnholdV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves plants that are currently on hold at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetOnhold V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetOnholdV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetOnholdV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetVegetative V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetVegetativeV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetVegetativeV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves vegetative-phase plants at a specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *
     * GET GetVegetative V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetVegetativeV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetVegetativeV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves a list of recorded plant waste events for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWaste V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetWasteV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetWasteV2(licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteMethodsAll V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetWasteMethodsAllV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetWasteMethodsAllV1(no));
    }

    /**
     * Retrieves a list of all available plant waste methods for use within a Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteMethodsAll V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetWasteMethodsAllV2(String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetWasteMethodsAllV2(pagenumber, pagesize));
    }

    /**
     * Retrieves a list of package records linked to the specified plantWasteId for a given facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWastePackage V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetWastePackageV2(String id, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetWastePackageV2(id, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves a list of plants records linked to the specified plantWasteId for a given facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWastePlant V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetWastePlantV2(String id, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetWastePlantV2(id, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetWasteReasonsV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetWasteReasonsV1(licensenumber));
    }

    /**
     * Retriveves available reasons for recording mature plant waste at a specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsGetWasteReasonsV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantsGetWasteReasonsV2(licensenumber, pagenumber, pagesize));
    }

    public static class PlantsUpdateAdjustV2RequestItem {
        public Integer AdjustCount;
        public String AdjustReason;
        public String AdjustmentDate;
        public Integer Id;
        public String Label;
        public String ReasonNote;
    }

    /**
     * Adjusts the recorded count of plants at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateAdjust V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsUpdateAdjustV2(String licensenumber, List<PlantsUpdateAdjustV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsUpdateAdjustV2(licensenumber, body));
    }

    public static class PlantsUpdateGrowthphaseV2RequestItem {
        public String GrowthDate;
        public String GrowthPhase;
        public Integer Id;
        public String Label;
        public String NewLocation;
        public String NewSublocation;
        public String NewTag;
    }

    /**
     * Changes the growth phases of plants within a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateGrowthphase V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsUpdateGrowthphaseV2(String licensenumber, List<PlantsUpdateGrowthphaseV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsUpdateGrowthphaseV2(licensenumber, body));
    }

    public static class PlantsUpdateHarvestV2RequestItem {
        public String ActualDate;
        public String DryingLocation;
        public String DryingSublocation;
        public String HarvestName;
        public String PatientLicenseNumber;
        public String Plant;
        public String UnitOfWeight;
        public Double Weight;
    }

    /**
     * Processes whole plant Harvest data for a specific Facility. NOTE: If HarvestName is excluded from the request body, or if it is passed in as null, the harvest name is auto-generated.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * PUT UpdateHarvest V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsUpdateHarvestV2(String licensenumber, List<PlantsUpdateHarvestV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsUpdateHarvestV2(licensenumber, body));
    }

    public static class PlantsUpdateLocationV2RequestItem {
        public String ActualDate;
        public Integer Id;
        public String Label;
        public String Location;
        public String Sublocation;
    }

    /**
     * Moves plant batches to new locations within a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateLocation V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsUpdateLocationV2(String licensenumber, List<PlantsUpdateLocationV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsUpdateLocationV2(licensenumber, body));
    }

    public static class PlantsUpdateMergeV2RequestItem {
        public String MergeDate;
        public String SourcePlantGroupLabel;
        public String TargetPlantGroupLabel;
    }

    /**
     * Merges multiple plant groups into a single group within a Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manicure/Harvest Veg/Flower Plants
     *
     * PUT UpdateMerge V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsUpdateMergeV2(String licensenumber, List<PlantsUpdateMergeV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsUpdateMergeV2(licensenumber, body));
    }

    public static class PlantsUpdateSplitV2RequestItem {
        public Integer PlantCount;
        public String SourcePlantLabel;
        public String SplitDate;
        public String StrainLabel;
        public String TagLabel;
    }

    /**
     * Splits an existing plant group into multiple groups within a Facility.
     * 
     *   Permissions Required:
     *   - View Plant
     *
     * PUT UpdateSplit V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsUpdateSplitV2(String licensenumber, List<PlantsUpdateSplitV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsUpdateSplitV2(licensenumber, body));
    }

    public static class PlantsUpdateStrainV2RequestItem {
        public Integer Id;
        public String Label;
        public Integer StrainId;
        public String StrainName;
    }

    /**
     * Updates the strain information for plants within a Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateStrain V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsUpdateStrainV2(String licensenumber, List<PlantsUpdateStrainV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsUpdateStrainV2(licensenumber, body));
    }

    public static class PlantsUpdateTagV2RequestItem {
        public Integer Id;
        public String Label;
        public String NewTag;
        public String ReplaceDate;
        public Integer TagId;
    }

    /**
     * Replaces existing plant tags with new tags for plants within a Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateTag V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantsUpdateTagV2(String licensenumber, List<PlantsUpdateTagV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantsUpdateTagV2(licensenumber, body));
    }

    public static class AdditivestemplatesCreateV2RequestItem {
        public List<AdditivestemplatesCreateV2RequestItem_ActiveIngredients> ActiveIngredients;
        public String AdditiveType;
        public String ApplicationDevice;
        public String EpaRegistrationNumber;
        public String Name;
        public String Note;
        public String ProductSupplier;
        public String ProductTradeName;
        public String RestrictiveEntryIntervalQuantityDescription;
        public String RestrictiveEntryIntervalTimeDescription;
    }

    public static class AdditivestemplatesCreateV2RequestItem_ActiveIngredients {
        public String Name;
        public Double Percentage;
    }

    /**
     * Creates new additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * POST Create V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object additivesTemplatesCreateV2(String licensenumber, List<AdditivestemplatesCreateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.additivesTemplatesCreateV2(licensenumber, body));
    }

    /**
     * Retrieves an Additive Template by its Id.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * GET Get V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object additivesTemplatesGetV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.additivesTemplatesGetV2(id, licensenumber));
    }

    /**
     * Retrieves a list of active additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * GET GetActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object additivesTemplatesGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.additivesTemplatesGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves a list of inactive additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * GET GetInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object additivesTemplatesGetInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.additivesTemplatesGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    public static class AdditivestemplatesUpdateV2RequestItem {
        public List<AdditivestemplatesUpdateV2RequestItem_ActiveIngredients> ActiveIngredients;
        public String AdditiveType;
        public String ApplicationDevice;
        public String EpaRegistrationNumber;
        public Integer Id;
        public String Name;
        public String Note;
        public String ProductSupplier;
        public String ProductTradeName;
        public String RestrictiveEntryIntervalQuantityDescription;
        public String RestrictiveEntryIntervalTimeDescription;
    }

    public static class AdditivestemplatesUpdateV2RequestItem_ActiveIngredients {
        public String Name;
        public Double Percentage;
    }

    /**
     * Updates existing additive templates for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Additives
     *
     * PUT Update V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object additivesTemplatesUpdateV2(String licensenumber, List<AdditivestemplatesUpdateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.additivesTemplatesUpdateV2(licensenumber, body));
    }

    /**
     * Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Caregivers
     *
     * GET GetByCaregiverLicenseNumber V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object caregiversStatusGetByCaregiverLicenseNumberV1(String caregiverlicensenumber, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.caregiversStatusGetByCaregiverLicenseNumberV1(caregiverlicensenumber, licensenumber));
    }

    /**
     * Retrieves the status of a Caregiver by their License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Caregivers
     *
     * GET GetByCaregiverLicenseNumber V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object caregiversStatusGetByCaregiverLicenseNumberV2(String caregiverlicensenumber, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.caregiversStatusGetByCaregiverLicenseNumberV2(caregiverlicensenumber, licensenumber));
    }

    public static class HarvestsCreateFinishV1RequestItem {
        public String ActualDate;
        public Integer Id;
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * POST CreateFinish V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsCreateFinishV1(String licensenumber, List<HarvestsCreateFinishV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsCreateFinishV1(licensenumber, body));
    }

    public static class HarvestsCreatePackageV1RequestItem {
        public String ActualDate;
        public Boolean DecontaminateProduct;
        public String DecontaminationDate;
        public String DecontaminationSteps;
        public String ExpirationDate;
        public List<HarvestsCreatePackageV1RequestItem_Ingredients> Ingredients;
        public Boolean IsDonation;
        public Boolean IsProductionBatch;
        public Boolean IsTradeSample;
        public String Item;
        public Integer LabTestStageId;
        public String Location;
        public String Note;
        public String PatientLicenseNumber;
        public Integer ProcessingJobTypeId;
        public Boolean ProductRequiresDecontamination;
        public Boolean ProductRequiresRemediation;
        public String ProductionBatchNumber;
        public Boolean RemediateProduct;
        public String RemediationDate;
        public Integer RemediationMethodId;
        public String RemediationSteps;
        public List<Object> RequiredLabTestBatches;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UnitOfWeight;
        public String UseByDate;
    }

    public static class HarvestsCreatePackageV1RequestItem_Ingredients {
        public Integer HarvestId;
        public String HarvestName;
        public String UnitOfWeight;
        public Double Weight;
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackage V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsCreatePackageV1(String licensenumber, List<HarvestsCreatePackageV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsCreatePackageV1(licensenumber, body));
    }

    public static class HarvestsCreatePackageV2RequestItem {
        public String ActualDate;
        public Boolean DecontaminateProduct;
        public String DecontaminationDate;
        public String DecontaminationSteps;
        public String ExpirationDate;
        public List<HarvestsCreatePackageV2RequestItem_Ingredients> Ingredients;
        public Boolean IsDonation;
        public Boolean IsProductionBatch;
        public Boolean IsTradeSample;
        public String Item;
        public Integer LabTestStageId;
        public String Location;
        public String Note;
        public String PatientLicenseNumber;
        public Integer ProcessingJobTypeId;
        public Boolean ProductRequiresDecontamination;
        public Boolean ProductRequiresRemediation;
        public String ProductionBatchNumber;
        public Boolean RemediateProduct;
        public String RemediationDate;
        public Integer RemediationMethodId;
        public String RemediationSteps;
        public List<Object> RequiredLabTestBatches;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UnitOfWeight;
        public String UseByDate;
    }

    public static class HarvestsCreatePackageV2RequestItem_Ingredients {
        public Integer HarvestId;
        public String HarvestName;
        public String UnitOfWeight;
        public Double Weight;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsCreatePackageV2(String licensenumber, List<HarvestsCreatePackageV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsCreatePackageV2(licensenumber, body));
    }

    public static class HarvestsCreatePackageTestingV1RequestItem {
        public String ActualDate;
        public Boolean DecontaminateProduct;
        public String DecontaminationDate;
        public String DecontaminationSteps;
        public String ExpirationDate;
        public List<HarvestsCreatePackageTestingV1RequestItem_Ingredients> Ingredients;
        public Boolean IsDonation;
        public Boolean IsProductionBatch;
        public Boolean IsTradeSample;
        public String Item;
        public Integer LabTestStageId;
        public String Location;
        public String Note;
        public String PatientLicenseNumber;
        public Integer ProcessingJobTypeId;
        public Boolean ProductRequiresDecontamination;
        public Boolean ProductRequiresRemediation;
        public String ProductionBatchNumber;
        public Boolean RemediateProduct;
        public String RemediationDate;
        public Integer RemediationMethodId;
        public String RemediationSteps;
        public List<Object> RequiredLabTestBatches;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UnitOfWeight;
        public String UseByDate;
    }

    public static class HarvestsCreatePackageTestingV1RequestItem_Ingredients {
        public Integer HarvestId;
        public String HarvestName;
        public String UnitOfWeight;
        public Double Weight;
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackageTesting V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsCreatePackageTestingV1(String licensenumber, List<HarvestsCreatePackageTestingV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsCreatePackageTestingV1(licensenumber, body));
    }

    public static class HarvestsCreatePackageTestingV2RequestItem {
        public String ActualDate;
        public Boolean DecontaminateProduct;
        public String DecontaminationDate;
        public String DecontaminationSteps;
        public String ExpirationDate;
        public List<HarvestsCreatePackageTestingV2RequestItem_Ingredients> Ingredients;
        public Boolean IsDonation;
        public Boolean IsProductionBatch;
        public Boolean IsTradeSample;
        public String Item;
        public Integer LabTestStageId;
        public String Location;
        public String Note;
        public String PatientLicenseNumber;
        public Integer ProcessingJobTypeId;
        public Boolean ProductRequiresDecontamination;
        public Boolean ProductRequiresRemediation;
        public String ProductionBatchNumber;
        public Boolean RemediateProduct;
        public String RemediationDate;
        public Integer RemediationMethodId;
        public String RemediationSteps;
        public List<Object> RequiredLabTestBatches;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UnitOfWeight;
        public String UseByDate;
    }

    public static class HarvestsCreatePackageTestingV2RequestItem_Ingredients {
        public Integer HarvestId;
        public String HarvestName;
        public String UnitOfWeight;
        public Double Weight;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsCreatePackageTestingV2(String licensenumber, List<HarvestsCreatePackageTestingV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsCreatePackageTestingV2(licensenumber, body));
    }

    public static class HarvestsCreateRemoveWasteV1RequestItem {
        public String ActualDate;
        public Integer Id;
        public String UnitOfWeight;
        public String WasteType;
        public Double WasteWeight;
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * POST CreateRemoveWaste V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsCreateRemoveWasteV1(String licensenumber, List<HarvestsCreateRemoveWasteV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsCreateRemoveWasteV1(licensenumber, body));
    }

    public static class HarvestsCreateUnfinishV1RequestItem {
        public Integer Id;
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * POST CreateUnfinish V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsCreateUnfinishV1(String licensenumber, List<HarvestsCreateUnfinishV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsCreateUnfinishV1(licensenumber, body));
    }

    public static class HarvestsCreateWasteV2RequestItem {
        public String ActualDate;
        public Integer Id;
        public String UnitOfWeight;
        public String WasteType;
        public Double WasteWeight;
    }

    /**
     * Records Waste from harvests for a specified Facility. NOTE: The IDs passed in the request body are the harvest IDs for which you are documenting waste.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * POST CreateWaste V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsCreateWasteV2(String licensenumber, List<HarvestsCreateWasteV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsCreateWasteV2(licensenumber, body));
    }

    /**
     * Discontinues a specific harvest waste record by Id for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Discontinue Harvest Waste
     *
     * DELETE DeleteWaste V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsDeleteWasteV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsDeleteWasteV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET Get V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsGetV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.harvestsGetV1(id, licensenumber));
    }

    /**
     * Retrieves a Harvest by its Id, optionally validated against a specified Facility License Number.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET Get V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsGetV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.harvestsGetV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET GetActive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsGetActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.harvestsGetActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves a list of active harvests for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.harvestsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET GetInactive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsGetInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.harvestsGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves a list of inactive harvests for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsGetInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.harvestsGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *
     * GET GetOnhold V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsGetOnholdV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.harvestsGetOnholdV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves a list of harvests on hold for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetOnhold V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsGetOnholdV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.harvestsGetOnholdV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves a list of Waste records for a specified Harvest, identified by its Harvest Id, within a Facility identified by its License Number.
     * 
     *   Permissions Required:
     *   - View Harvests
     *
     * GET GetWaste V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsGetWasteV2(String harvestid, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.harvestsGetWasteV2(harvestid, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteTypes V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsGetWasteTypesV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.harvestsGetWasteTypesV1(no));
    }

    /**
     * Retrieves a list of Waste types for harvests.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteTypes V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsGetWasteTypesV2(String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.harvestsGetWasteTypesV2(pagenumber, pagesize));
    }

    public static class HarvestsUpdateFinishV2RequestItem {
        public String ActualDate;
        public Integer Id;
    }

    /**
     * Marks one or more harvests as finished for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * PUT UpdateFinish V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsUpdateFinishV2(String licensenumber, List<HarvestsUpdateFinishV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsUpdateFinishV2(licensenumber, body));
    }

    public static class HarvestsUpdateLocationV2RequestItem {
        public String ActualDate;
        public String DryingLocation;
        public String DryingSublocation;
        public String HarvestName;
        public Integer Id;
    }

    /**
     * Updates the Location of Harvest for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateLocation V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsUpdateLocationV2(String licensenumber, List<HarvestsUpdateLocationV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsUpdateLocationV2(licensenumber, body));
    }

    public static class HarvestsUpdateMoveV1RequestItem {
        public String ActualDate;
        public String DryingLocation;
        public String DryingSublocation;
        public String HarvestName;
        public Integer Id;
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateMove V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsUpdateMoveV1(String licensenumber, List<HarvestsUpdateMoveV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsUpdateMoveV1(licensenumber, body));
    }

    public static class HarvestsUpdateRenameV1RequestItem {
        public Integer Id;
        public String NewName;
        public String OldName;
    }

    /**
     * Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateRename V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsUpdateRenameV1(String licensenumber, List<HarvestsUpdateRenameV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsUpdateRenameV1(licensenumber, body));
    }

    public static class HarvestsUpdateRenameV2RequestItem {
        public Integer Id;
        public String NewName;
        public String OldName;
    }

    /**
     * Renames one or more harvests for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Manage Harvests
     *
     * PUT UpdateRename V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsUpdateRenameV2(String licensenumber, List<HarvestsUpdateRenameV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsUpdateRenameV2(licensenumber, body));
    }

    public static class HarvestsUpdateRestoreHarvestedPlantsV2RequestItem {
        public Integer HarvestId;
        public List<String> PlantTags;
    }

    /**
     * Restores previously harvested plants to their original state for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * PUT UpdateRestoreHarvestedPlants V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsUpdateRestoreHarvestedPlantsV2(String licensenumber, List<HarvestsUpdateRestoreHarvestedPlantsV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsUpdateRestoreHarvestedPlantsV2(licensenumber, body));
    }

    public static class HarvestsUpdateUnfinishV2RequestItem {
        public Integer Id;
    }

    /**
     * Reopens one or more previously finished harvests for the specified Facility.
     * 
     *   Permissions Required:
     *   - View Harvests
     *   - Finish/Discontinue Harvests
     *
     * PUT UpdateUnfinish V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object harvestsUpdateUnfinishV2(String licensenumber, List<HarvestsUpdateUnfinishV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.harvestsUpdateUnfinishV2(licensenumber, body));
    }

    public static class LabtestsCreateRecordV1RequestItem {
        public String DocumentFileBase64;
        public String DocumentFileName;
        public String Label;
        public String ResultDate;
        public List<LabtestsCreateRecordV1RequestItem_Results> Results;
    }

    public static class LabtestsCreateRecordV1RequestItem_Results {
        public String LabTestTypeName;
        public String Notes;
        public Boolean Passed;
        public Double Quantity;
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateRecord V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsCreateRecordV1(String licensenumber, List<LabtestsCreateRecordV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.labTestsCreateRecordV1(licensenumber, body));
    }

    public static class LabtestsCreateRecordV2RequestItem {
        public String DocumentFileBase64;
        public String DocumentFileName;
        public String Label;
        public String ResultDate;
        public List<LabtestsCreateRecordV2RequestItem_Results> Results;
    }

    public static class LabtestsCreateRecordV2RequestItem_Results {
        public String LabTestTypeName;
        public String Notes;
        public Boolean Passed;
        public Double Quantity;
    }

    /**
     * Submits Lab Test results for one or more packages. NOTE: This endpoint allows only PDF files, and uploaded files can be no more than 5 MB in size. The Label element in the request is a Package Label.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateRecord V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsCreateRecordV2(String licensenumber, List<LabtestsCreateRecordV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.labTestsCreateRecordV2(licensenumber, body));
    }

    /**
     * Retrieves a list of Lab Test batches.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetBatches V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsGetBatchesV2(String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.labTestsGetBatchesV2(pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * GET GetLabtestdocument V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsGetLabtestdocumentV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.labTestsGetLabtestdocumentV1(id, licensenumber));
    }

    /**
     * Retrieves a specific Lab Test result document by its Id for a given Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * GET GetLabtestdocument V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsGetLabtestdocumentV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.labTestsGetLabtestdocumentV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetResults V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsGetResultsV1(String licensenumber, String packageid) throws Exception {
        return rateLimiter.execute(null, true, () -> client.labTestsGetResultsV1(licensenumber, packageid));
    }

    /**
     * Retrieves Lab Test results for a specified Package.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * GET GetResults V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsGetResultsV2(String licensenumber, String packageid, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.labTestsGetResultsV2(licensenumber, packageid, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetStates V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsGetStatesV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.labTestsGetStatesV1(no));
    }

    /**
     * Returns a list of all lab testing states.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetStates V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsGetStatesV2(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.labTestsGetStatesV2(no));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsGetTypesV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.labTestsGetTypesV1(no));
    }

    /**
     * Returns a list of Lab Test types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsGetTypesV2(String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.labTestsGetTypesV2(pagenumber, pagesize));
    }

    public static class LabtestsUpdateLabtestdocumentV1RequestItem {
        public String DocumentFileBase64;
        public String DocumentFileName;
        public Integer LabTestResultId;
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateLabtestdocument V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsUpdateLabtestdocumentV1(String licensenumber, List<LabtestsUpdateLabtestdocumentV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.labTestsUpdateLabtestdocumentV1(licensenumber, body));
    }

    public static class LabtestsUpdateLabtestdocumentV2RequestItem {
        public String DocumentFileBase64;
        public String DocumentFileName;
        public Integer LabTestResultId;
    }

    /**
     * Updates one or more documents for previously submitted lab tests.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateLabtestdocument V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsUpdateLabtestdocumentV2(String licensenumber, List<LabtestsUpdateLabtestdocumentV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.labTestsUpdateLabtestdocumentV2(licensenumber, body));
    }

    public static class LabtestsUpdateResultReleaseV1RequestItem {
        public String PackageLabel;
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateResultRelease V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsUpdateResultReleaseV1(String licensenumber, List<LabtestsUpdateResultReleaseV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.labTestsUpdateResultReleaseV1(licensenumber, body));
    }

    public static class LabtestsUpdateResultReleaseV2RequestItem {
        public String PackageLabel;
    }

    /**
     * Releases Lab Test results for one or more packages.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateResultRelease V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object labTestsUpdateResultReleaseV2(String licensenumber, List<LabtestsUpdateResultReleaseV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.labTestsUpdateResultReleaseV2(licensenumber, body));
    }

    public static class SalesCreateDeliveryV1RequestItem {
        public Integer ConsumerId;
        public String DriverEmployeeId;
        public String DriverName;
        public String DriversLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public String PatientLicenseNumber;
        public String PhoneNumberForQuestions;
        public String PlannedRoute;
        public String RecipientAddressCity;
        public String RecipientAddressCounty;
        public String RecipientAddressPostalCode;
        public String RecipientAddressState;
        public String RecipientAddressStreet1;
        public String RecipientAddressStreet2;
        public String RecipientName;
        public Integer RecipientZoneId;
        public String SalesCustomerType;
        public String SalesDateTime;
        public List<SalesCreateDeliveryV1RequestItem_Transactions> Transactions;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class SalesCreateDeliveryV1RequestItem_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * POST CreateDelivery V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryV1(String licensenumber, List<SalesCreateDeliveryV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateDeliveryV1(licensenumber, body));
    }

    public static class SalesCreateDeliveryV2RequestItem {
        public Integer ConsumerId;
        public String DriverEmployeeId;
        public String DriverName;
        public String DriversLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public String PatientLicenseNumber;
        public String PhoneNumberForQuestions;
        public String PlannedRoute;
        public String RecipientAddressCity;
        public String RecipientAddressCounty;
        public String RecipientAddressPostalCode;
        public String RecipientAddressState;
        public String RecipientAddressStreet1;
        public String RecipientAddressStreet2;
        public String RecipientName;
        public Integer RecipientZoneId;
        public String SalesCustomerType;
        public String SalesDateTime;
        public List<SalesCreateDeliveryV2RequestItem_Transactions> Transactions;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class SalesCreateDeliveryV2RequestItem_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryV2(String licensenumber, List<SalesCreateDeliveryV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateDeliveryV2(licensenumber, body));
    }

    public static class SalesCreateDeliveryRetailerV1RequestItem {
        public String DateTime;
        public List<SalesCreateDeliveryRetailerV1RequestItem_Destinations> Destinations;
        public String DriverEmployeeId;
        public String DriverName;
        public String DriversLicenseNumber;
        public String EstimatedDepartureDateTime;
        public List<SalesCreateDeliveryRetailerV1RequestItem_Packages> Packages;
        public String PhoneNumberForQuestions;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class SalesCreateDeliveryRetailerV1RequestItem_Destinations {
        public String ConsumerId;
        public String EstimatedArrivalDateTime;
        public String PatientLicenseNumber;
        public String RecipientAddressCity;
        public String RecipientAddressCounty;
        public String RecipientAddressPostalCode;
        public String RecipientAddressState;
        public String RecipientAddressStreet1;
        public String RecipientAddressStreet2;
        public String RecipientName;
        public String RecipientZoneId;
        public String SalesCustomerType;
        public List<SalesCreateDeliveryRetailerV1RequestItem_Destinations_Transactions> Transactions;
    }

    public static class SalesCreateDeliveryRetailerV1RequestItem_Destinations_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    public static class SalesCreateDeliveryRetailerV1RequestItem_Packages {
        public String DateTime;
        public String PackageLabel;
        public Integer Quantity;
        public Double TotalPrice;
        public String UnitOfMeasure;
    }

    /**
     * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailer V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerV1(String licensenumber, List<SalesCreateDeliveryRetailerV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateDeliveryRetailerV1(licensenumber, body));
    }

    public static class SalesCreateDeliveryRetailerV2RequestItem {
        public String DateTime;
        public List<SalesCreateDeliveryRetailerV2RequestItem_Destinations> Destinations;
        public String DriverEmployeeId;
        public String DriverName;
        public String DriversLicenseNumber;
        public String EstimatedDepartureDateTime;
        public List<SalesCreateDeliveryRetailerV2RequestItem_Packages> Packages;
        public String PhoneNumberForQuestions;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class SalesCreateDeliveryRetailerV2RequestItem_Destinations {
        public String ConsumerId;
        public String EstimatedArrivalDateTime;
        public String PatientLicenseNumber;
        public String RecipientAddressCity;
        public String RecipientAddressCounty;
        public String RecipientAddressPostalCode;
        public String RecipientAddressState;
        public String RecipientAddressStreet1;
        public String RecipientAddressStreet2;
        public String RecipientName;
        public String RecipientZoneId;
        public String SalesCustomerType;
        public List<SalesCreateDeliveryRetailerV2RequestItem_Destinations_Transactions> Transactions;
    }

    public static class SalesCreateDeliveryRetailerV2RequestItem_Destinations_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    public static class SalesCreateDeliveryRetailerV2RequestItem_Packages {
        public String DateTime;
        public String PackageLabel;
        public Integer Quantity;
        public Double TotalPrice;
        public String UnitOfMeasure;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerV2(String licensenumber, List<SalesCreateDeliveryRetailerV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateDeliveryRetailerV2(licensenumber, body));
    }

    public static class SalesCreateDeliveryRetailerDepartV1RequestItem {
        public Integer RetailerDeliveryId;
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerDepart V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerDepartV1(String licensenumber, List<SalesCreateDeliveryRetailerDepartV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateDeliveryRetailerDepartV1(licensenumber, body));
    }

    public static class SalesCreateDeliveryRetailerDepartV2RequestItem {
        public Integer RetailerDeliveryId;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerDepartV2(String licensenumber, List<SalesCreateDeliveryRetailerDepartV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateDeliveryRetailerDepartV2(licensenumber, body));
    }

    public static class SalesCreateDeliveryRetailerEndV1RequestItem {
        public String ActualArrivalDateTime;
        public List<SalesCreateDeliveryRetailerEndV1RequestItem_Packages> Packages;
        public Integer RetailerDeliveryId;
    }

    public static class SalesCreateDeliveryRetailerEndV1RequestItem_Packages {
        public Integer EndQuantity;
        public String EndUnitOfMeasure;
        public String Label;
    }

    /**
     * Please note: The ActualArrivalDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerEnd V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerEndV1(String licensenumber, List<SalesCreateDeliveryRetailerEndV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateDeliveryRetailerEndV1(licensenumber, body));
    }

    public static class SalesCreateDeliveryRetailerEndV2RequestItem {
        public String ActualArrivalDateTime;
        public List<SalesCreateDeliveryRetailerEndV2RequestItem_Packages> Packages;
        public Integer RetailerDeliveryId;
    }

    public static class SalesCreateDeliveryRetailerEndV2RequestItem_Packages {
        public Integer EndQuantity;
        public String EndUnitOfMeasure;
        public String Label;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerEndV2(String licensenumber, List<SalesCreateDeliveryRetailerEndV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateDeliveryRetailerEndV2(licensenumber, body));
    }

    public static class SalesCreateDeliveryRetailerRestockV1RequestItem {
        public String DateTime;
        public String Destinations;
        public String EstimatedDepartureDateTime;
        public List<SalesCreateDeliveryRetailerRestockV1RequestItem_Packages> Packages;
        public Integer RetailerDeliveryId;
    }

    public static class SalesCreateDeliveryRetailerRestockV1RequestItem_Packages {
        public String PackageLabel;
        public Integer Quantity;
        public Boolean RemoveCurrentPackage;
        public Double TotalPrice;
        public String UnitOfMeasure;
    }

    /**
     * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerRestock V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerRestockV1(String licensenumber, List<SalesCreateDeliveryRetailerRestockV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateDeliveryRetailerRestockV1(licensenumber, body));
    }

    public static class SalesCreateDeliveryRetailerRestockV2RequestItem {
        public String DateTime;
        public String Destinations;
        public String EstimatedDepartureDateTime;
        public List<SalesCreateDeliveryRetailerRestockV2RequestItem_Packages> Packages;
        public Integer RetailerDeliveryId;
    }

    public static class SalesCreateDeliveryRetailerRestockV2RequestItem_Packages {
        public String PackageLabel;
        public Integer Quantity;
        public Boolean RemoveCurrentPackage;
        public Double TotalPrice;
        public String UnitOfMeasure;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerRestockV2(String licensenumber, List<SalesCreateDeliveryRetailerRestockV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateDeliveryRetailerRestockV2(licensenumber, body));
    }

    public static class SalesCreateDeliveryRetailerSaleV1RequestItem {
        public Integer ConsumerId;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public String PatientLicenseNumber;
        public String PhoneNumberForQuestions;
        public String PlannedRoute;
        public String RecipientAddressCity;
        public String RecipientAddressCounty;
        public String RecipientAddressPostalCode;
        public String RecipientAddressState;
        public String RecipientAddressStreet1;
        public String RecipientAddressStreet2;
        public String RecipientName;
        public Integer RecipientZoneId;
        public Integer RetailerDeliveryId;
        public String SalesCustomerType;
        public String SalesDateTime;
        public List<SalesCreateDeliveryRetailerSaleV1RequestItem_Transactions> Transactions;
    }

    public static class SalesCreateDeliveryRetailerSaleV1RequestItem_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * POST CreateDeliveryRetailerSale V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerSaleV1(String licensenumber, List<SalesCreateDeliveryRetailerSaleV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateDeliveryRetailerSaleV1(licensenumber, body));
    }

    public static class SalesCreateDeliveryRetailerSaleV2RequestItem {
        public Integer ConsumerId;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public String PatientLicenseNumber;
        public String PhoneNumberForQuestions;
        public String PlannedRoute;
        public String RecipientAddressCity;
        public String RecipientAddressCounty;
        public String RecipientAddressPostalCode;
        public String RecipientAddressState;
        public String RecipientAddressStreet1;
        public String RecipientAddressStreet2;
        public String RecipientName;
        public Integer RecipientZoneId;
        public Integer RetailerDeliveryId;
        public String SalesCustomerType;
        public String SalesDateTime;
        public List<SalesCreateDeliveryRetailerSaleV2RequestItem_Transactions> Transactions;
    }

    public static class SalesCreateDeliveryRetailerSaleV2RequestItem_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateDeliveryRetailerSaleV2(String licensenumber, List<SalesCreateDeliveryRetailerSaleV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateDeliveryRetailerSaleV2(licensenumber, body));
    }

    public static class SalesCreateReceiptV1RequestItem {
        public String CaregiverLicenseNumber;
        public String ExternalReceiptNumber;
        public String IdentificationMethod;
        public String PatientLicenseNumber;
        public Integer PatientRegistrationLocationId;
        public String SalesCustomerType;
        public String SalesDateTime;
        public List<SalesCreateReceiptV1RequestItem_Transactions> Transactions;
    }

    public static class SalesCreateReceiptV1RequestItem_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales
     *
     * POST CreateReceipt V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateReceiptV1(String licensenumber, List<SalesCreateReceiptV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateReceiptV1(licensenumber, body));
    }

    public static class SalesCreateReceiptV2RequestItem {
        public String CaregiverLicenseNumber;
        public String ExternalReceiptNumber;
        public String IdentificationMethod;
        public String PatientLicenseNumber;
        public Integer PatientRegistrationLocationId;
        public String SalesCustomerType;
        public String SalesDateTime;
        public List<SalesCreateReceiptV2RequestItem_Transactions> Transactions;
    }

    public static class SalesCreateReceiptV2RequestItem_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateReceiptV2(String licensenumber, List<SalesCreateReceiptV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateReceiptV2(licensenumber, body));
    }

    public static class SalesCreateTransactionByDateV1RequestItem {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * POST CreateTransactionByDate V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesCreateTransactionByDateV1(String date, String licensenumber, List<SalesCreateTransactionByDateV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesCreateTransactionByDateV1(date, licensenumber, body));
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * DELETE DeleteDelivery V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesDeleteDeliveryV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesDeleteDeliveryV1(id, licensenumber));
    }

    /**
     * Voids a sales delivery for a Facility using the provided License Number and delivery Id.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery
     *
     * DELETE DeleteDelivery V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesDeleteDeliveryV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesDeleteDeliveryV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * DELETE DeleteDeliveryRetailer V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesDeleteDeliveryRetailerV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesDeleteDeliveryRetailerV1(id, licensenumber));
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesDeleteDeliveryRetailerV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesDeleteDeliveryRetailerV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * DELETE DeleteReceipt V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesDeleteReceiptV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesDeleteReceiptV1(id, licensenumber));
    }

    /**
     * Archives a sales receipt for a Facility using the provided License Number and receipt Id.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * DELETE DeleteReceipt V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesDeleteReceiptV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesDeleteReceiptV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetCounties V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetCountiesV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetCountiesV1(no));
    }

    /**
     * Returns a list of counties available for sales deliveries.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetCounties V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetCountiesV2(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetCountiesV2(no));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetCustomertypes V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetCustomertypesV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetCustomertypesV1(no));
    }

    /**
     * Returns a list of customer types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetCustomertypes V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetCustomertypesV2(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetCustomertypesV2(no));
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDeliveriesActive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String salesdateend, String salesdatestart) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveriesActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber, salesdateend, salesdatestart));
    }

    /**
     * Returns a list of active sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetDeliveriesActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize, String salesdateend, String salesdatestart) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveriesActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize, salesdateend, salesdatestart));
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDeliveriesInactive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String salesdateend, String salesdatestart) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveriesInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber, salesdateend, salesdatestart));
    }

    /**
     * Returns a list of inactive sales deliveries for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetDeliveriesInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize, String salesdateend, String salesdatestart) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveriesInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize, salesdateend, salesdatestart));
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * GET GetDeliveriesRetailerActive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesRetailerActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveriesRetailerActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Returns a list of active retailer deliveries for a Facility, optionally filtered by last modified date range
     * 
     *   Permissions Required:
     *   - View Retailer Delivery
     *   - Manage Retailer Delivery
     *
     * GET GetDeliveriesRetailerActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesRetailerActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveriesRetailerActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * GET GetDeliveriesRetailerInactive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesRetailerInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveriesRetailerInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Returns a list of inactive retailer deliveries for a Facility, optionally filtered by last modified date range
     * 
     *   Permissions Required:
     *   - View Retailer Delivery
     *   - Manage Retailer Delivery
     *
     * GET GetDeliveriesRetailerInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesRetailerInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveriesRetailerInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   -
     *
     * GET GetDeliveriesReturnreasons V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesReturnreasonsV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveriesReturnreasonsV1(licensenumber));
    }

    /**
     * Returns a list of return reasons for sales deliveries based on the provided License Number.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDeliveriesReturnreasons V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveriesReturnreasonsV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveriesReturnreasonsV2(licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetDelivery V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveryV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveryV1(id, licensenumber));
    }

    /**
     * Retrieves a sales delivery record by its Id, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetDelivery V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveryV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveryV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Retailer Delivery
     *
     * GET GetDeliveryRetailer V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveryRetailerV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveryRetailerV1(id, licensenumber));
    }

    /**
     * Retrieves a retailer delivery record by its ID, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Retailer Delivery
     *   - Manage Retailer Delivery
     *
     * GET GetDeliveryRetailer V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetDeliveryRetailerV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetDeliveryRetailerV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   -
     *
     * GET GetPatientRegistrationsLocations V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetPatientRegistrationsLocationsV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetPatientRegistrationsLocationsV1(no));
    }

    /**
     * Returns a list of valid Patient registration locations for sales.
     * 
     *   Permissions Required:
     *   -
     *
     * GET GetPatientRegistrationsLocations V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetPatientRegistrationsLocationsV2(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetPatientRegistrationsLocationsV2(no));
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * GET GetPaymenttypes V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetPaymenttypesV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetPaymenttypesV1(licensenumber));
    }

    /**
     * Returns a list of available payment types for the specified License Number.
     * 
     *   Permissions Required:
     *   - View Sales Delivery
     *   - Manage Sales Delivery
     *
     * GET GetPaymenttypes V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetPaymenttypesV2(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetPaymenttypesV2(licensenumber));
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetReceipt V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetReceiptV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetReceiptV1(id, licensenumber));
    }

    /**
     * Retrieves a sales receipt by its Id, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceipt V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetReceiptV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetReceiptV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetReceiptsActive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetReceiptsActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String salesdateend, String salesdatestart) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetReceiptsActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber, salesdateend, salesdatestart));
    }

    /**
     * Returns a list of active sales receipts for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceiptsActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetReceiptsActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize, String salesdateend, String salesdatestart) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetReceiptsActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize, salesdateend, salesdatestart));
    }

    /**
     * Retrieves a Sales Receipt by its external number, with an optional License Number.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceiptsExternalByExternalNumber V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetReceiptsExternalByExternalNumberV2(String externalnumber, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetReceiptsExternalByExternalNumberV2(externalnumber, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetReceiptsInactive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetReceiptsInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String salesdateend, String salesdatestart) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetReceiptsInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber, salesdateend, salesdatestart));
    }

    /**
     * Returns a list of inactive sales receipts for a Facility, filtered by optional sales or last modified date ranges.
     * 
     *   Permissions Required:
     *   - View Sales
     *   - Manage Sales
     *
     * GET GetReceiptsInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetReceiptsInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize, String salesdateend, String salesdatestart) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetReceiptsInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize, salesdateend, salesdatestart));
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetTransactions V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetTransactionsV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetTransactionsV1(licensenumber));
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * GET GetTransactionsBySalesDateStartAndSalesDateEnd V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesGetTransactionsBySalesDateStartAndSalesDateEndV1(String salesdatestart, String salesdateend, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.salesGetTransactionsBySalesDateStartAndSalesDateEndV1(salesdatestart, salesdateend, licensenumber));
    }

    public static class SalesUpdateDeliveryV1RequestItem {
        public Integer ConsumerId;
        public String DriverEmployeeId;
        public String DriverName;
        public String DriversLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Integer Id;
        public String PatientLicenseNumber;
        public String PhoneNumberForQuestions;
        public String PlannedRoute;
        public String RecipientAddressCity;
        public String RecipientAddressCounty;
        public String RecipientAddressPostalCode;
        public String RecipientAddressState;
        public String RecipientAddressStreet1;
        public String RecipientAddressStreet2;
        public String RecipientName;
        public String RecipientZoneId;
        public String SalesCustomerType;
        public String SalesDateTime;
        public List<SalesUpdateDeliveryV1RequestItem_Transactions> Transactions;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class SalesUpdateDeliveryV1RequestItem_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * PUT UpdateDelivery V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryV1(String licensenumber, List<SalesUpdateDeliveryV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryV1(licensenumber, body));
    }

    public static class SalesUpdateDeliveryV2RequestItem {
        public Integer ConsumerId;
        public String DriverEmployeeId;
        public String DriverName;
        public String DriversLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Integer Id;
        public String PatientLicenseNumber;
        public String PhoneNumberForQuestions;
        public String PlannedRoute;
        public String RecipientAddressCity;
        public String RecipientAddressCounty;
        public String RecipientAddressPostalCode;
        public String RecipientAddressState;
        public String RecipientAddressStreet1;
        public String RecipientAddressStreet2;
        public String RecipientName;
        public String RecipientZoneId;
        public String SalesCustomerType;
        public String SalesDateTime;
        public List<SalesUpdateDeliveryV2RequestItem_Transactions> Transactions;
        public String TransporterFacilityLicenseNumber;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class SalesUpdateDeliveryV2RequestItem_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    /**
     * Updates sales delivery records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery
     *
     * PUT UpdateDelivery V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryV2(String licensenumber, List<SalesUpdateDeliveryV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryV2(licensenumber, body));
    }

    public static class SalesUpdateDeliveryCompleteV1RequestItem {
        public List<String> AcceptedPackages;
        public String ActualArrivalDateTime;
        public Integer Id;
        public String PaymentType;
        public List<SalesUpdateDeliveryCompleteV1RequestItem_ReturnedPackages> ReturnedPackages;
    }

    public static class SalesUpdateDeliveryCompleteV1RequestItem_ReturnedPackages {
        public String Label;
        public Integer ReturnQuantityVerified;
        public String ReturnReason;
        public String ReturnReasonNote;
        public String ReturnUnitOfMeasure;
    }

    /**
     * Permissions Required:
     *   - Sales Delivery
     *
     * PUT UpdateDeliveryComplete V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryCompleteV1(String licensenumber, List<SalesUpdateDeliveryCompleteV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryCompleteV1(licensenumber, body));
    }

    public static class SalesUpdateDeliveryCompleteV2RequestItem {
        public List<String> AcceptedPackages;
        public String ActualArrivalDateTime;
        public Integer Id;
        public String PaymentType;
        public List<SalesUpdateDeliveryCompleteV2RequestItem_ReturnedPackages> ReturnedPackages;
    }

    public static class SalesUpdateDeliveryCompleteV2RequestItem_ReturnedPackages {
        public String Label;
        public Integer ReturnQuantityVerified;
        public String ReturnReason;
        public String ReturnReasonNote;
        public String ReturnUnitOfMeasure;
    }

    /**
     * Completes a list of sales deliveries for a Facility using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery
     *
     * PUT UpdateDeliveryComplete V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryCompleteV2(String licensenumber, List<SalesUpdateDeliveryCompleteV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryCompleteV2(licensenumber, body));
    }

    public static class SalesUpdateDeliveryHubV1RequestItem {
        public String DriverEmployeeId;
        public String DriverName;
        public String DriversLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Integer Id;
        public String PhoneNumberForQuestions;
        public String PlannedRoute;
        public String TransporterFacilityId;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales Delivery
     *
     * PUT UpdateDeliveryHub V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubV1(String licensenumber, List<SalesUpdateDeliveryHubV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryHubV1(licensenumber, body));
    }

    public static class SalesUpdateDeliveryHubV2RequestItem {
        public String DriverEmployeeId;
        public String DriverName;
        public String DriversLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Integer Id;
        public String PhoneNumberForQuestions;
        public String PlannedRoute;
        public String TransporterFacilityId;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    /**
     * Updates hub transporter details for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery, Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHub V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubV2(String licensenumber, List<SalesUpdateDeliveryHubV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryHubV2(licensenumber, body));
    }

    public static class SalesUpdateDeliveryHubAcceptV1RequestItem {
        public Integer Id;
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateDeliveryHubAccept V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubAcceptV1(String licensenumber, List<SalesUpdateDeliveryHubAcceptV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryHubAcceptV1(licensenumber, body));
    }

    public static class SalesUpdateDeliveryHubAcceptV2RequestItem {
        public Integer Id;
    }

    /**
     * Accepts a list of hub sales deliveries for a Facility based on the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHubAccept V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubAcceptV2(String licensenumber, List<SalesUpdateDeliveryHubAcceptV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryHubAcceptV2(licensenumber, body));
    }

    public static class SalesUpdateDeliveryHubDepartV1RequestItem {
        public Integer Id;
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateDeliveryHubDepart V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubDepartV1(String licensenumber, List<SalesUpdateDeliveryHubDepartV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryHubDepartV1(licensenumber, body));
    }

    public static class SalesUpdateDeliveryHubDepartV2RequestItem {
        public Integer Id;
    }

    /**
     * Processes the departure of hub sales deliveries for a Facility using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHubDepart V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubDepartV2(String licensenumber, List<SalesUpdateDeliveryHubDepartV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryHubDepartV2(licensenumber, body));
    }

    public static class SalesUpdateDeliveryHubVerifyIdV1RequestItem {
        public Integer Id;
        public String PaymentType;
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateDeliveryHubVerifyID V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubVerifyIdV1(String licensenumber, List<SalesUpdateDeliveryHubVerifyIdV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryHubVerifyIdV1(licensenumber, body));
    }

    public static class SalesUpdateDeliveryHubVerifyIdV2RequestItem {
        public Integer Id;
        public String PaymentType;
    }

    /**
     * Verifies identification for a list of hub sales deliveries using the provided License Number and delivery data.
     * 
     *   Permissions Required:
     *   - Manage Sales Delivery Hub
     *
     * PUT UpdateDeliveryHubVerifyID V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryHubVerifyIdV2(String licensenumber, List<SalesUpdateDeliveryHubVerifyIdV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryHubVerifyIdV2(licensenumber, body));
    }

    public static class SalesUpdateDeliveryRetailerV1RequestItem {
        public String DateTime;
        public List<SalesUpdateDeliveryRetailerV1RequestItem_Destinations> Destinations;
        public String DriverEmployeeId;
        public String DriverName;
        public String DriversLicenseNumber;
        public String EstimatedDepartureDateTime;
        public Integer Id;
        public List<SalesUpdateDeliveryRetailerV1RequestItem_Packages> Packages;
        public String PhoneNumberForQuestions;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class SalesUpdateDeliveryRetailerV1RequestItem_Destinations {
        public String ConsumerId;
        public Integer DriverEmployeeId;
        public String DriverName;
        public String DriversLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Integer Id;
        public String PatientLicenseNumber;
        public String PhoneNumberForQuestions;
        public String PlannedRoute;
        public String RecipientAddressCity;
        public String RecipientAddressCounty;
        public String RecipientAddressPostalCode;
        public String RecipientAddressState;
        public String RecipientAddressStreet1;
        public String RecipientAddressStreet2;
        public String RecipientName;
        public String RecipientZoneId;
        public String SalesCustomerType;
        public String SalesDateTime;
        public List<SalesUpdateDeliveryRetailerV1RequestItem_Destinations_Transactions> Transactions;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class SalesUpdateDeliveryRetailerV1RequestItem_Destinations_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    public static class SalesUpdateDeliveryRetailerV1RequestItem_Packages {
        public String DateTime;
        public String PackageLabel;
        public Integer Quantity;
        public Double TotalPrice;
        public String UnitOfMeasure;
    }

    /**
     * Please note: The DateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Retailer Delivery
     *
     * PUT UpdateDeliveryRetailer V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryRetailerV1(String licensenumber, List<SalesUpdateDeliveryRetailerV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryRetailerV1(licensenumber, body));
    }

    public static class SalesUpdateDeliveryRetailerV2RequestItem {
        public String DateTime;
        public List<SalesUpdateDeliveryRetailerV2RequestItem_Destinations> Destinations;
        public String DriverEmployeeId;
        public String DriverName;
        public String DriversLicenseNumber;
        public String EstimatedDepartureDateTime;
        public Integer Id;
        public List<SalesUpdateDeliveryRetailerV2RequestItem_Packages> Packages;
        public String PhoneNumberForQuestions;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class SalesUpdateDeliveryRetailerV2RequestItem_Destinations {
        public String ConsumerId;
        public Integer DriverEmployeeId;
        public String DriverName;
        public String DriversLicenseNumber;
        public String EstimatedArrivalDateTime;
        public String EstimatedDepartureDateTime;
        public Integer Id;
        public String PatientLicenseNumber;
        public String PhoneNumberForQuestions;
        public String PlannedRoute;
        public String RecipientAddressCity;
        public String RecipientAddressCounty;
        public String RecipientAddressPostalCode;
        public String RecipientAddressState;
        public String RecipientAddressStreet1;
        public String RecipientAddressStreet2;
        public String RecipientName;
        public String RecipientZoneId;
        public String SalesCustomerType;
        public String SalesDateTime;
        public List<SalesUpdateDeliveryRetailerV2RequestItem_Destinations_Transactions> Transactions;
        public String VehicleLicensePlateNumber;
        public String VehicleMake;
        public String VehicleModel;
    }

    public static class SalesUpdateDeliveryRetailerV2RequestItem_Destinations_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    public static class SalesUpdateDeliveryRetailerV2RequestItem_Packages {
        public String DateTime;
        public String PackageLabel;
        public Integer Quantity;
        public Double TotalPrice;
        public String UnitOfMeasure;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateDeliveryRetailerV2(String licensenumber, List<SalesUpdateDeliveryRetailerV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateDeliveryRetailerV2(licensenumber, body));
    }

    public static class SalesUpdateReceiptV1RequestItem {
        public String CaregiverLicenseNumber;
        public String ExternalReceiptNumber;
        public Integer Id;
        public String IdentificationMethod;
        public String PatientLicenseNumber;
        public Integer PatientRegistrationLocationId;
        public String SalesCustomerType;
        public String SalesDateTime;
        public List<SalesUpdateReceiptV1RequestItem_Transactions> Transactions;
    }

    public static class SalesUpdateReceiptV1RequestItem_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    /**
     * Please note: The SalesDateTime field must be the actual date and time of the transaction without time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Sales
     *
     * PUT UpdateReceipt V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateReceiptV1(String licensenumber, List<SalesUpdateReceiptV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateReceiptV1(licensenumber, body));
    }

    public static class SalesUpdateReceiptV2RequestItem {
        public String CaregiverLicenseNumber;
        public String ExternalReceiptNumber;
        public Integer Id;
        public String IdentificationMethod;
        public String PatientLicenseNumber;
        public Integer PatientRegistrationLocationId;
        public String SalesCustomerType;
        public String SalesDateTime;
        public List<SalesUpdateReceiptV2RequestItem_Transactions> Transactions;
    }

    public static class SalesUpdateReceiptV2RequestItem_Transactions {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    /**
     * Updates sales receipt records for a given License Number. Please note: The SalesDateTime field must be the actual date and time of the transaction without the time zone. This date/time must already be in the same time zone as the Facility recording the sales. For example, if the Facility is in Pacific Time, then this time must be in Pacific Standard (or Daylight Savings) Time and not in UTC.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * PUT UpdateReceipt V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateReceiptV2(String licensenumber, List<SalesUpdateReceiptV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateReceiptV2(licensenumber, body));
    }

    public static class SalesUpdateReceiptFinalizeV2RequestItem {
        public Integer Id;
    }

    /**
     * Finalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * PUT UpdateReceiptFinalize V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateReceiptFinalizeV2(String licensenumber, List<SalesUpdateReceiptFinalizeV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateReceiptFinalizeV2(licensenumber, body));
    }

    public static class SalesUpdateReceiptUnfinalizeV2RequestItem {
        public Integer Id;
    }

    /**
     * Unfinalizes a list of sales receipts for a Facility using the provided License Number and receipt data.
     * 
     *   Permissions Required:
     *   - Manage Sales
     *
     * PUT UpdateReceiptUnfinalize V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateReceiptUnfinalizeV2(String licensenumber, List<SalesUpdateReceiptUnfinalizeV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateReceiptUnfinalizeV2(licensenumber, body));
    }

    public static class SalesUpdateTransactionByDateV1RequestItem {
        public String CityTax;
        public String CountyTax;
        public String DiscountAmount;
        public String ExciseTax;
        public String InvoiceNumber;
        public String MunicipalTax;
        public String PackageLabel;
        public String Price;
        public String QrCodes;
        public Integer Quantity;
        public String SalesTax;
        public String SubTotal;
        public Double TotalAmount;
        public String UnitOfMeasure;
        public Double UnitThcContent;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    /**
     * Permissions Required:
     *   - Sales
     *
     * PUT UpdateTransactionByDate V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object salesUpdateTransactionByDateV1(String date, String licensenumber, List<SalesUpdateTransactionByDateV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.salesUpdateTransactionByDateV1(date, licensenumber, body));
    }

    public static class TransportersCreateDriverV2RequestItem {
        public String DriversLicenseNumber;
        public String EmployeeId;
        public String Name;
    }

    /**
     * Creates new driver records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * POST CreateDriver V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transportersCreateDriverV2(String licensenumber, List<TransportersCreateDriverV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transportersCreateDriverV2(licensenumber, body));
    }

    public static class TransportersCreateVehicleV2RequestItem {
        public String LicensePlateNumber;
        public String Make;
        public String Model;
    }

    /**
     * Creates new vehicle records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * POST CreateVehicle V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transportersCreateVehicleV2(String licensenumber, List<TransportersCreateVehicleV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transportersCreateVehicleV2(licensenumber, body));
    }

    /**
     * Archives a Driver record for a Facility.  Please note: The {id} parameter above represents a Driver Id.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * DELETE DeleteDriver V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transportersDeleteDriverV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transportersDeleteDriverV2(id, licensenumber));
    }

    /**
     * Archives a Vehicle for a facility.  Please note: The {id} parameter above represents a Vehicle Id.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * DELETE DeleteVehicle V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transportersDeleteVehicleV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transportersDeleteVehicleV2(id, licensenumber));
    }

    /**
     * Retrieves a Driver by its Id, with an optional license number. Please note: The {id} parameter above represents a Driver Id.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetDriver V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transportersGetDriverV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transportersGetDriverV2(id, licensenumber));
    }

    /**
     * Retrieves a list of drivers for a Facility.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetDrivers V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transportersGetDriversV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transportersGetDriversV2(licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves a Vehicle by its Id, with an optional license number. Please note: The {id} parameter above represents a Vehicle Id.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetVehicle V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transportersGetVehicleV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transportersGetVehicleV2(id, licensenumber));
    }

    /**
     * Retrieves a list of vehicles for a Facility.
     * 
     *   Permissions Required:
     *   - Transporters
     *
     * GET GetVehicles V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transportersGetVehiclesV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.transportersGetVehiclesV2(licensenumber, pagenumber, pagesize));
    }

    public static class TransportersUpdateDriverV2RequestItem {
        public String DriversLicenseNumber;
        public String EmployeeId;
        public String Id;
        public String Name;
    }

    /**
     * Updates existing driver records for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * PUT UpdateDriver V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transportersUpdateDriverV2(String licensenumber, List<TransportersUpdateDriverV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transportersUpdateDriverV2(licensenumber, body));
    }

    public static class TransportersUpdateVehicleV2RequestItem {
        public String Id;
        public String LicensePlateNumber;
        public String Make;
        public String Model;
    }

    /**
     * Updates existing vehicle records for a facility.
     * 
     *   Permissions Required:
     *   - Manage Transporters
     *
     * PUT UpdateVehicle V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object transportersUpdateVehicleV2(String licensenumber, List<TransportersUpdateVehicleV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.transportersUpdateVehicleV2(licensenumber, body));
    }

    /**
     * This endpoint provides a list of facilities for which the authenticated user has access.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAll V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object facilitiesGetAllV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.facilitiesGetAllV1(no));
    }

    /**
     * This endpoint provides a list of facilities for which the authenticated user has access.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAll V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object facilitiesGetAllV2(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.facilitiesGetAllV2(no));
    }

    public static class PackagesCreateV1RequestItem {
        public String ActualDate;
        public String ExpirationDate;
        public List<PackagesCreateV1RequestItem_Ingredients> Ingredients;
        public Boolean IsDonation;
        public Boolean IsProductionBatch;
        public Boolean IsTradeSample;
        public String Item;
        public Integer LabTestStageId;
        public String Location;
        public String Note;
        public String PatientLicenseNumber;
        public Integer ProcessingJobTypeId;
        public Boolean ProductRequiresRemediation;
        public String ProductionBatchNumber;
        public Integer Quantity;
        public Boolean RequiredLabTestBatches;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UnitOfMeasure;
        public String UseByDate;
        public Boolean UseSameItem;
    }

    public static class PackagesCreateV1RequestItem_Ingredients {
        public String Package;
        public Integer Quantity;
        public String UnitOfMeasure;
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST Create V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesCreateV1(String licensenumber, List<PackagesCreateV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesCreateV1(licensenumber, body));
    }

    public static class PackagesCreateV2RequestItem {
        public String ActualDate;
        public String ExpirationDate;
        public List<PackagesCreateV2RequestItem_Ingredients> Ingredients;
        public Boolean IsDonation;
        public Boolean IsProductionBatch;
        public Boolean IsTradeSample;
        public String Item;
        public Integer LabTestStageId;
        public String Location;
        public String Note;
        public String PatientLicenseNumber;
        public Integer ProcessingJobTypeId;
        public Boolean ProductRequiresRemediation;
        public String ProductionBatchNumber;
        public Integer Quantity;
        public Boolean RequiredLabTestBatches;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UnitOfMeasure;
        public String UseByDate;
        public Boolean UseSameItem;
    }

    public static class PackagesCreateV2RequestItem_Ingredients {
        public String Package;
        public Integer Quantity;
        public String UnitOfMeasure;
    }

    /**
     * Creates new packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST Create V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesCreateV2(String licensenumber, List<PackagesCreateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesCreateV2(licensenumber, body));
    }

    public static class PackagesCreateAdjustV1RequestItem {
        public String AdjustmentDate;
        public String AdjustmentReason;
        public String Label;
        public Integer Quantity;
        public String ReasonNote;
        public String UnitOfMeasure;
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateAdjust V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesCreateAdjustV1(String licensenumber, List<PackagesCreateAdjustV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesCreateAdjustV1(licensenumber, body));
    }

    public static class PackagesCreateAdjustV2RequestItem {
        public String AdjustmentDate;
        public String AdjustmentReason;
        public String Label;
        public Integer Quantity;
        public String ReasonNote;
        public String UnitOfMeasure;
    }

    /**
     * Records a list of adjustments for packages at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateAdjust V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesCreateAdjustV2(String licensenumber, List<PackagesCreateAdjustV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesCreateAdjustV2(licensenumber, body));
    }

    public static class PackagesCreateChangeItemV1RequestItem {
        public String Item;
        public String Label;
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateChangeItem V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesCreateChangeItemV1(String licensenumber, List<PackagesCreateChangeItemV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesCreateChangeItemV1(licensenumber, body));
    }

    public static class PackagesCreateChangeLocationV1RequestItem {
        public String Label;
        public String Location;
        public String MoveDate;
        public String Sublocation;
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateChangeLocation V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesCreateChangeLocationV1(String licensenumber, List<PackagesCreateChangeLocationV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesCreateChangeLocationV1(licensenumber, body));
    }

    public static class PackagesCreateFinishV1RequestItem {
        public String ActualDate;
        public String Label;
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateFinish V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesCreateFinishV1(String licensenumber, List<PackagesCreateFinishV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesCreateFinishV1(licensenumber, body));
    }

    public static class PackagesCreatePlantingsV1RequestItem {
        public String LocationName;
        public Integer PackageAdjustmentAmount;
        public String PackageAdjustmentUnitOfMeasureName;
        public String PackageLabel;
        public String PatientLicenseNumber;
        public String PlantBatchName;
        public String PlantBatchType;
        public Integer PlantCount;
        public String PlantedDate;
        public String StrainName;
        public String SublocationName;
        public String UnpackagedDate;
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreatePlantings V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesCreatePlantingsV1(String licensenumber, List<PackagesCreatePlantingsV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesCreatePlantingsV1(licensenumber, body));
    }

    public static class PackagesCreatePlantingsV2RequestItem {
        public String LocationName;
        public Integer PackageAdjustmentAmount;
        public String PackageAdjustmentUnitOfMeasureName;
        public String PackageLabel;
        public String PatientLicenseNumber;
        public String PlantBatchName;
        public String PlantBatchType;
        public Integer PlantCount;
        public String PlantedDate;
        public String StrainName;
        public String SublocationName;
        public String UnpackagedDate;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesCreatePlantingsV2(String licensenumber, List<PackagesCreatePlantingsV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesCreatePlantingsV2(licensenumber, body));
    }

    public static class PackagesCreateRemediateV1RequestItem {
        public String PackageLabel;
        public String RemediationDate;
        public String RemediationMethodName;
        public String RemediationSteps;
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateRemediate V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesCreateRemediateV1(String licensenumber, List<PackagesCreateRemediateV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesCreateRemediateV1(licensenumber, body));
    }

    public static class PackagesCreateTestingV1RequestItem {
        public String ActualDate;
        public String ExpirationDate;
        public List<PackagesCreateTestingV1RequestItem_Ingredients> Ingredients;
        public Boolean IsDonation;
        public Boolean IsProductionBatch;
        public Boolean IsTradeSample;
        public String Item;
        public Integer LabTestStageId;
        public String Location;
        public String Note;
        public String PatientLicenseNumber;
        public Integer ProcessingJobTypeId;
        public Boolean ProductRequiresRemediation;
        public String ProductionBatchNumber;
        public Integer Quantity;
        public Boolean RequiredLabTestBatches;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UnitOfMeasure;
        public String UseByDate;
        public Boolean UseSameItem;
    }

    public static class PackagesCreateTestingV1RequestItem_Ingredients {
        public String Package;
        public Integer Quantity;
        public String UnitOfMeasure;
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateTesting V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesCreateTestingV1(String licensenumber, List<PackagesCreateTestingV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesCreateTestingV1(licensenumber, body));
    }

    public static class PackagesCreateTestingV2RequestItem {
        public String ActualDate;
        public String ExpirationDate;
        public List<PackagesCreateTestingV2RequestItem_Ingredients> Ingredients;
        public Boolean IsDonation;
        public Boolean IsProductionBatch;
        public Boolean IsTradeSample;
        public String Item;
        public Integer LabTestStageId;
        public String Location;
        public String Note;
        public String PatientLicenseNumber;
        public Integer ProcessingJobTypeId;
        public Boolean ProductRequiresRemediation;
        public String ProductionBatchNumber;
        public Integer Quantity;
        public List<String> RequiredLabTestBatches;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UnitOfMeasure;
        public String UseByDate;
        public Boolean UseSameItem;
    }

    public static class PackagesCreateTestingV2RequestItem_Ingredients {
        public String Package;
        public Integer Quantity;
        public String UnitOfMeasure;
    }

    /**
     * Creates new packages for testing for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreateTesting V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesCreateTestingV2(String licensenumber, List<PackagesCreateTestingV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesCreateTestingV2(licensenumber, body));
    }

    public static class PackagesCreateUnfinishV1RequestItem {
        public String Label;
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * POST CreateUnfinish V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesCreateUnfinishV1(String licensenumber, List<PackagesCreateUnfinishV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesCreateUnfinishV1(licensenumber, body));
    }

    /**
     * Discontinues a Package at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * DELETE Delete V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesDeleteV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesDeleteV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET Get V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetV1(id, licensenumber));
    }

    /**
     * Retrieves a Package by its Id.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET Get V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetActive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves a list of active packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetAdjustReasons V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetAdjustReasonsV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetAdjustReasonsV1(licensenumber));
    }

    /**
     * Retrieves a list of adjustment reasons for packages at a specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAdjustReasons V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetAdjustReasonsV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetAdjustReasonsV2(licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetByLabel V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetByLabelV1(String label, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetByLabelV1(label, licensenumber));
    }

    /**
     * Retrieves a Package by its label.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetByLabel V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetByLabelV2(String label, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetByLabelV2(label, licensenumber));
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetInactive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves a list of inactive packages for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves a list of packages in transit for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetIntransit V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetIntransitV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetIntransitV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves a list of lab sample packages created or sent for testing for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetLabsamples V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetLabsamplesV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetLabsamplesV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - View Packages
     *
     * GET GetOnhold V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetOnholdV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetOnholdV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves a list of packages on hold for a specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetOnhold V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetOnholdV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetOnholdV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves the source harvests for a Package by its Id.
     * 
     *   Permissions Required:
     *   - View Package Source Harvests
     *
     * GET GetSourceHarvest V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetSourceHarvestV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetSourceHarvestV2(id, licensenumber));
    }

    /**
     * Retrieves a list of transferred packages for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *
     * GET GetTransferred V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetTransferredV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetTransferredV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetTypesV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetTypesV1(no));
    }

    /**
     * Retrieves a list of available Package types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesGetTypesV2(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.packagesGetTypesV2(no));
    }

    public static class PackagesUpdateAdjustV2RequestItem {
        public String AdjustmentDate;
        public String AdjustmentReason;
        public String Label;
        public Integer Quantity;
        public String ReasonNote;
        public String UnitOfMeasure;
    }

    /**
     * Set the final quantity for a Package.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateAdjust V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateAdjustV2(String licensenumber, List<PackagesUpdateAdjustV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateAdjustV2(licensenumber, body));
    }

    public static class PackagesUpdateChangeNoteV1RequestItem {
        public String Note;
        public String PackageLabel;
    }

    /**
     * Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *   - Manage Package Notes
     *
     * PUT UpdateChangeNote V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateChangeNoteV1(String licensenumber, List<PackagesUpdateChangeNoteV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateChangeNoteV1(licensenumber, body));
    }

    public static class PackagesUpdateDecontaminateV2RequestItem {
        public String DecontaminationDate;
        public String DecontaminationMethodName;
        public String DecontaminationSteps;
        public String PackageLabel;
    }

    /**
     * Updates the Product decontaminate information for a list of packages at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateDecontaminate V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateDecontaminateV2(String licensenumber, List<PackagesUpdateDecontaminateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateDecontaminateV2(licensenumber, body));
    }

    public static class PackagesUpdateDonationFlagV2RequestItem {
        public String Label;
    }

    /**
     * Flags one or more packages for donation at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateDonationFlag V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateDonationFlagV2(String licensenumber, List<PackagesUpdateDonationFlagV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateDonationFlagV2(licensenumber, body));
    }

    public static class PackagesUpdateDonationUnflagV2RequestItem {
        public String Label;
    }

    /**
     * Removes the donation flag from one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateDonationUnflag V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateDonationUnflagV2(String licensenumber, List<PackagesUpdateDonationUnflagV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateDonationUnflagV2(licensenumber, body));
    }

    public static class PackagesUpdateExternalidV2RequestItem {
        public String ExternalId;
        public String PackageLabel;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateExternalidV2(String licensenumber, List<PackagesUpdateExternalidV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateExternalidV2(licensenumber, body));
    }

    public static class PackagesUpdateFinishV2RequestItem {
        public String ActualDate;
        public String Label;
    }

    /**
     * Updates a list of packages as finished for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateFinish V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateFinishV2(String licensenumber, List<PackagesUpdateFinishV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateFinishV2(licensenumber, body));
    }

    public static class PackagesUpdateItemV2RequestItem {
        public String Item;
        public String Label;
    }

    /**
     * Updates the associated Item for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateItem V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateItemV2(String licensenumber, List<PackagesUpdateItemV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateItemV2(licensenumber, body));
    }

    public static class PackagesUpdateLabTestRequiredV2RequestItem {
        public String Label;
        public List<String> RequiredLabTestBatches;
    }

    /**
     * Updates the list of required lab test batches for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateLabTestRequired V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateLabTestRequiredV2(String licensenumber, List<PackagesUpdateLabTestRequiredV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateLabTestRequiredV2(licensenumber, body));
    }

    public static class PackagesUpdateLocationV2RequestItem {
        public String Label;
        public String Location;
        public String MoveDate;
        public String Sublocation;
    }

    /**
     * Updates the Location and Sublocation for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateLocation V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateLocationV2(String licensenumber, List<PackagesUpdateLocationV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateLocationV2(licensenumber, body));
    }

    public static class PackagesUpdateNoteV2RequestItem {
        public String Note;
        public String PackageLabel;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateNoteV2(String licensenumber, List<PackagesUpdateNoteV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateNoteV2(licensenumber, body));
    }

    public static class PackagesUpdateRemediateV2RequestItem {
        public String PackageLabel;
        public String RemediationDate;
        public String RemediationMethodName;
        public String RemediationSteps;
    }

    /**
     * Updates a list of Product remediations for packages at a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateRemediate V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateRemediateV2(String licensenumber, List<PackagesUpdateRemediateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateRemediateV2(licensenumber, body));
    }

    public static class PackagesUpdateTradesampleFlagV2RequestItem {
        public String Label;
    }

    /**
     * Flags or unflags one or more packages at the specified Facility as trade samples.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateTradesampleFlag V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateTradesampleFlagV2(String licensenumber, List<PackagesUpdateTradesampleFlagV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateTradesampleFlagV2(licensenumber, body));
    }

    public static class PackagesUpdateTradesampleUnflagV2RequestItem {
        public String Label;
    }

    /**
     * Removes the trade sample flag from one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateTradesampleUnflag V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateTradesampleUnflagV2(String licensenumber, List<PackagesUpdateTradesampleUnflagV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateTradesampleUnflagV2(licensenumber, body));
    }

    public static class PackagesUpdateUnfinishV2RequestItem {
        public String Label;
    }

    /**
     * Updates a list of packages as unfinished for a specific Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Manage Packages Inventory
     *
     * PUT UpdateUnfinish V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateUnfinishV2(String licensenumber, List<PackagesUpdateUnfinishV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateUnfinishV2(licensenumber, body));
    }

    public static class PackagesUpdateUsebydateV2RequestItem {
        public String ExpirationDate;
        public String Label;
        public String SellByDate;
        public String UseByDate;
    }

    /**
     * Updates the use-by date for one or more packages at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * PUT UpdateUsebydate V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object packagesUpdateUsebydateV2(String licensenumber, List<PackagesUpdateUsebydateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.packagesUpdateUsebydateV2(licensenumber, body));
    }

    public static class PatientcheckinsCreateV1RequestItem {
        public String CheckInDate;
        public Integer CheckInLocationId;
        public String PatientLicenseNumber;
        public String RegistrationExpiryDate;
        public String RegistrationStartDate;
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * POST Create V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientCheckInsCreateV1(String licensenumber, List<PatientcheckinsCreateV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.patientCheckInsCreateV1(licensenumber, body));
    }

    public static class PatientcheckinsCreateV2RequestItem {
        public String CheckInDate;
        public Integer CheckInLocationId;
        public String PatientLicenseNumber;
        public String RegistrationExpiryDate;
        public String RegistrationStartDate;
    }

    /**
     * Records patient check-ins for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * POST Create V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientCheckInsCreateV2(String licensenumber, List<PatientcheckinsCreateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.patientCheckInsCreateV2(licensenumber, body));
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * DELETE Delete V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientCheckInsDeleteV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.patientCheckInsDeleteV1(id, licensenumber));
    }

    /**
     * Archives a Patient Check-In, identified by its Id, for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * DELETE Delete V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientCheckInsDeleteV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.patientCheckInsDeleteV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * GET GetAll V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientCheckInsGetAllV1(String checkindateend, String checkindatestart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.patientCheckInsGetAllV1(checkindateend, checkindatestart, licensenumber));
    }

    /**
     * Retrieves a list of patient check-ins for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * GET GetAll V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientCheckInsGetAllV2(String checkindateend, String checkindatestart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.patientCheckInsGetAllV2(checkindateend, checkindatestart, licensenumber));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetLocations V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientCheckInsGetLocationsV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.patientCheckInsGetLocationsV1(no));
    }

    /**
     * Retrieves a list of Patient Check-In locations.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetLocations V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientCheckInsGetLocationsV2(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.patientCheckInsGetLocationsV2(no));
    }

    public static class PatientcheckinsUpdateV1RequestItem {
        public String CheckInDate;
        public Integer CheckInLocationId;
        public Integer Id;
        public String PatientLicenseNumber;
        public String RegistrationExpiryDate;
        public String RegistrationStartDate;
    }

    /**
     * Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * PUT Update V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientCheckInsUpdateV1(String licensenumber, List<PatientcheckinsUpdateV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.patientCheckInsUpdateV1(licensenumber, body));
    }

    public static class PatientcheckinsUpdateV2RequestItem {
        public String CheckInDate;
        public Integer CheckInLocationId;
        public Integer Id;
        public String PatientLicenseNumber;
        public String RegistrationExpiryDate;
        public String RegistrationStartDate;
    }

    /**
     * Updates patient check-ins for a specified Facility.
     * 
     *   Permissions Required:
     *   - ManagePatientsCheckIns
     *
     * PUT Update V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientCheckInsUpdateV2(String licensenumber, List<PatientcheckinsUpdateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.patientCheckInsUpdateV2(licensenumber, body));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetActive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object unitsOfMeasureGetActiveV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.unitsOfMeasureGetActiveV1(no));
    }

    /**
     * Retrieves all active units of measure.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object unitsOfMeasureGetActiveV2(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.unitsOfMeasureGetActiveV2(no));
    }

    /**
     * Retrieves all inactive units of measure.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object unitsOfMeasureGetInactiveV2(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.unitsOfMeasureGetInactiveV2(no));
    }

    /**
     * Retrieves all available waste methods.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetAll V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object wasteMethodsGetAllV2(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.wasteMethodsGetAllV2(no));
    }

    /**
     * Permissions Required:
     *   - Manage Employees
     *
     * GET GetAll V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object employeesGetAllV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.employeesGetAllV1(licensenumber));
    }

    /**
     * Retrieves a list of employees for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Employees
     *   - View Employees
     *
     * GET GetAll V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object employeesGetAllV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.employeesGetAllV2(licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves the permissions of a specified Employee, identified by their Employee License Number, for a given Facility.
     * 
     *   Permissions Required:
     *   - Manage Employees
     *
     * GET GetPermissions V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object employeesGetPermissionsV2(String employeelicensenumber, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.employeesGetPermissionsV2(employeelicensenumber, licensenumber));
    }

    public static class ItemsCreateV1RequestItem {
        public String AdministrationMethod;
        public String Allergens;
        public String Description;
        public String GlobalProductName;
        public String ItemBrand;
        public String ItemCategory;
        public String ItemIngredients;
        public String LabelImageFileSystemIds;
        public String LabelPhotoDescription;
        public String Name;
        public String NumberOfDoses;
        public String PackagingImageFileSystemIds;
        public String PackagingPhotoDescription;
        public String ProcessingJobCategoryName;
        public String ProcessingJobTypeName;
        public String ProductImageFileSystemIds;
        public String ProductPDFFileSystemIds;
        public String ProductPhotoDescription;
        public String PublicIngredients;
        public String ServingSize;
        public String Strain;
        public Integer SupplyDurationDays;
        public Double UnitCbdAContent;
        public Double UnitCbdAContentDose;
        public String UnitCbdAContentDoseUnitOfMeasure;
        public String UnitCbdAContentUnitOfMeasure;
        public Double UnitCbdAPercent;
        public Double UnitCbdContent;
        public Double UnitCbdContentDose;
        public String UnitCbdContentDoseUnitOfMeasure;
        public String UnitCbdContentUnitOfMeasure;
        public Double UnitCbdPercent;
        public String UnitOfMeasure;
        public Double UnitThcAContent;
        public Double UnitThcAContentDose;
        public String UnitThcAContentDoseUnitOfMeasure;
        public String UnitThcAContentUnitOfMeasure;
        public Double UnitThcAPercent;
        public Double UnitThcContent;
        public Double UnitThcContentDose;
        public String UnitThcContentDoseUnitOfMeasure;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitVolume;
        public String UnitVolumeUnitOfMeasure;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    /**
     * NOTE: To include a photo with an item, first use POST /items/v1/photo to POST the photo, and then use the returned ID in the request body in this endpoint.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST Create V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsCreateV1(String licensenumber, List<ItemsCreateV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.itemsCreateV1(licensenumber, body));
    }

    public static class ItemsCreateV2RequestItem {
        public String AdministrationMethod;
        public String Allergens;
        public String Description;
        public String GlobalProductName;
        public String ItemBrand;
        public String ItemCategory;
        public String ItemIngredients;
        public String LabelImageFileSystemIds;
        public String LabelPhotoDescription;
        public String Name;
        public String NumberOfDoses;
        public String PackagingImageFileSystemIds;
        public String PackagingPhotoDescription;
        public String ProcessingJobCategoryName;
        public String ProcessingJobTypeName;
        public String ProductImageFileSystemIds;
        public String ProductPDFFileSystemIds;
        public String ProductPhotoDescription;
        public String PublicIngredients;
        public String ServingSize;
        public String Strain;
        public Integer SupplyDurationDays;
        public Double UnitCbdAContent;
        public Double UnitCbdAContentDose;
        public String UnitCbdAContentDoseUnitOfMeasure;
        public String UnitCbdAContentUnitOfMeasure;
        public Double UnitCbdAPercent;
        public Double UnitCbdContent;
        public Double UnitCbdContentDose;
        public String UnitCbdContentDoseUnitOfMeasure;
        public String UnitCbdContentUnitOfMeasure;
        public Double UnitCbdPercent;
        public String UnitOfMeasure;
        public Double UnitThcAContent;
        public Double UnitThcAContentDose;
        public String UnitThcAContentDoseUnitOfMeasure;
        public String UnitThcAContentUnitOfMeasure;
        public Double UnitThcAPercent;
        public Double UnitThcContent;
        public Double UnitThcContentDose;
        public String UnitThcContentDoseUnitOfMeasure;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitVolume;
        public String UnitVolumeUnitOfMeasure;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    /**
     * Creates one or more new products for the specified Facility. NOTE: To include a photo with an item, first use POST /items/v2/photo to POST the photo, and then use the returned Id in the request body in this endpoint.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST Create V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsCreateV2(String licensenumber, List<ItemsCreateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.itemsCreateV2(licensenumber, body));
    }

    public static class ItemsCreateBrandV2RequestItem {
        public String Name;
    }

    /**
     * Creates one or more new item brands for the specified Facility identified by the License Number.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreateBrand V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsCreateBrandV2(String licensenumber, List<ItemsCreateBrandV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.itemsCreateBrandV2(licensenumber, body));
    }

    public static class ItemsCreateFileV2RequestItem {
        public String EncodedImageBase64;
        public String FileName;
    }

    /**
     * Uploads one or more image or PDF files for products, labels, packaging, or documents at the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreateFile V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsCreateFileV2(String licensenumber, List<ItemsCreateFileV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.itemsCreateFileV2(licensenumber, body));
    }

    public static class ItemsCreatePhotoV1RequestItem {
        public String EncodedImageBase64;
        public String FileName;
    }

    /**
     * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreatePhoto V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsCreatePhotoV1(String licensenumber, List<ItemsCreatePhotoV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.itemsCreatePhotoV1(licensenumber, body));
    }

    public static class ItemsCreatePhotoV2RequestItem {
        public String EncodedImageBase64;
        public String FileName;
    }

    /**
     * This endpoint allows only BMP, GIF, JPG, and PNG files and uploaded files can be no more than 5 MB in size.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * POST CreatePhoto V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsCreatePhotoV2(String licensenumber, List<ItemsCreatePhotoV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.itemsCreatePhotoV2(licensenumber, body));
    }

    public static class ItemsCreateUpdateV1RequestItem {
        public String AdministrationMethod;
        public String Allergens;
        public String Description;
        public String GlobalProductName;
        public Integer Id;
        public String ItemBrand;
        public String ItemCategory;
        public String ItemIngredients;
        public String LabelImageFileSystemIds;
        public String LabelPhotoDescription;
        public String Name;
        public String NumberOfDoses;
        public String PackagingImageFileSystemIds;
        public String PackagingPhotoDescription;
        public String ProcessingJobCategoryName;
        public String ProcessingJobTypeName;
        public String ProductImageFileSystemIds;
        public String ProductPDFFileSystemIds;
        public String ProductPhotoDescription;
        public String PublicIngredients;
        public String ServingSize;
        public String Strain;
        public Integer SupplyDurationDays;
        public Double UnitCbdAContent;
        public Double UnitCbdAContentDose;
        public String UnitCbdAContentDoseUnitOfMeasure;
        public String UnitCbdAContentUnitOfMeasure;
        public Double UnitCbdAPercent;
        public Double UnitCbdContent;
        public Double UnitCbdContentDose;
        public String UnitCbdContentDoseUnitOfMeasure;
        public String UnitCbdContentUnitOfMeasure;
        public Double UnitCbdPercent;
        public String UnitOfMeasure;
        public Double UnitThcAContent;
        public Double UnitThcAContentDose;
        public String UnitThcAContentDoseUnitOfMeasure;
        public String UnitThcAContentUnitOfMeasure;
        public Double UnitThcAPercent;
        public Double UnitThcContent;
        public Double UnitThcContentDose;
        public String UnitThcContentDoseUnitOfMeasure;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitVolume;
        public String UnitVolumeUnitOfMeasure;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * POST CreateUpdate V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsCreateUpdateV1(String licensenumber, List<ItemsCreateUpdateV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.itemsCreateUpdateV1(licensenumber, body));
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * DELETE Delete V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsDeleteV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.itemsDeleteV1(id, licensenumber));
    }

    /**
     * Archives the specified Product by Id for the given Facility License Number.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * DELETE Delete V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsDeleteV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.itemsDeleteV2(id, licensenumber));
    }

    /**
     * Archives the specified Item Brand by Id for the given Facility License Number.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * DELETE DeleteBrand V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsDeleteBrandV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.itemsDeleteBrandV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET Get V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsGetV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.itemsGetV1(id, licensenumber));
    }

    /**
     * Retrieves detailed information about a specific Item by Id.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET Get V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsGetV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.itemsGetV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetActive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsGetActiveV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.itemsGetActiveV1(licensenumber));
    }

    /**
     * Returns a list of active items for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.itemsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetBrands V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsGetBrandsV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.itemsGetBrandsV1(licensenumber));
    }

    /**
     * Retrieves a list of active item brands for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetBrands V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsGetBrandsV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.itemsGetBrandsV2(licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetCategories V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsGetCategoriesV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.itemsGetCategoriesV1(licensenumber));
    }

    /**
     * Retrieves a list of item categories.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetCategories V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsGetCategoriesV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.itemsGetCategoriesV2(licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves a file by its Id for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetFile V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsGetFileV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.itemsGetFileV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetInactive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsGetInactiveV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.itemsGetInactiveV1(licensenumber));
    }

    /**
     * Retrieves a list of inactive items for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsGetInactiveV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.itemsGetInactiveV2(licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - Manage Items
     *
     * GET GetPhoto V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsGetPhotoV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.itemsGetPhotoV1(id, licensenumber));
    }

    /**
     * Retrieves an image by its Id for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * GET GetPhoto V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsGetPhotoV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.itemsGetPhotoV2(id, licensenumber));
    }

    public static class ItemsUpdateV2RequestItem {
        public String AdministrationMethod;
        public String Allergens;
        public String Description;
        public String GlobalProductName;
        public Integer Id;
        public String ItemBrand;
        public String ItemCategory;
        public String ItemIngredients;
        public String LabelImageFileSystemIds;
        public String LabelPhotoDescription;
        public String Name;
        public String NumberOfDoses;
        public String PackagingImageFileSystemIds;
        public String PackagingPhotoDescription;
        public String ProcessingJobCategoryName;
        public String ProcessingJobTypeName;
        public String ProductImageFileSystemIds;
        public String ProductPDFFileSystemIds;
        public String ProductPhotoDescription;
        public String PublicIngredients;
        public String ServingSize;
        public String Strain;
        public Integer SupplyDurationDays;
        public Double UnitCbdAContent;
        public Double UnitCbdAContentDose;
        public String UnitCbdAContentDoseUnitOfMeasure;
        public String UnitCbdAContentUnitOfMeasure;
        public Double UnitCbdAPercent;
        public Double UnitCbdContent;
        public Double UnitCbdContentDose;
        public String UnitCbdContentDoseUnitOfMeasure;
        public String UnitCbdContentUnitOfMeasure;
        public Double UnitCbdPercent;
        public String UnitOfMeasure;
        public Double UnitThcAContent;
        public Double UnitThcAContentDose;
        public String UnitThcAContentDoseUnitOfMeasure;
        public String UnitThcAContentUnitOfMeasure;
        public Double UnitThcAPercent;
        public Double UnitThcContent;
        public Double UnitThcContentDose;
        public String UnitThcContentDoseUnitOfMeasure;
        public String UnitThcContentUnitOfMeasure;
        public Double UnitThcPercent;
        public Double UnitVolume;
        public String UnitVolumeUnitOfMeasure;
        public Double UnitWeight;
        public String UnitWeightUnitOfMeasure;
    }

    /**
     * Updates one or more existing products for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * PUT Update V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsUpdateV2(String licensenumber, List<ItemsUpdateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.itemsUpdateV2(licensenumber, body));
    }

    public static class ItemsUpdateBrandV2RequestItem {
        public Integer Id;
        public String Name;
    }

    /**
     * Updates one or more existing item brands for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Items
     *
     * PUT UpdateBrand V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object itemsUpdateBrandV2(String licensenumber, List<ItemsUpdateBrandV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.itemsUpdateBrandV2(licensenumber, body));
    }

    public static class LocationsCreateV1RequestItem {
        public String LocationTypeName;
        public String Name;
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * POST Create V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object locationsCreateV1(String licensenumber, List<LocationsCreateV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.locationsCreateV1(licensenumber, body));
    }

    public static class LocationsCreateV2RequestItem {
        public String LocationTypeName;
        public String Name;
    }

    /**
     * Creates new locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * POST Create V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object locationsCreateV2(String licensenumber, List<LocationsCreateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.locationsCreateV2(licensenumber, body));
    }

    public static class LocationsCreateUpdateV1RequestItem {
        public Integer Id;
        public String LocationTypeName;
        public String Name;
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * POST CreateUpdate V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object locationsCreateUpdateV1(String licensenumber, List<LocationsCreateUpdateV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.locationsCreateUpdateV1(licensenumber, body));
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * DELETE Delete V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object locationsDeleteV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.locationsDeleteV1(id, licensenumber));
    }

    /**
     * Archives a specified Location, identified by its Id, for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * DELETE Delete V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object locationsDeleteV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.locationsDeleteV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * GET Get V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object locationsGetV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.locationsGetV1(id, licensenumber));
    }

    /**
     * Retrieves a Location by its Id.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET Get V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object locationsGetV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.locationsGetV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * GET GetActive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object locationsGetActiveV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.locationsGetActiveV1(licensenumber));
    }

    /**
     * Retrieves a list of active locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object locationsGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.locationsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves a list of inactive locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object locationsGetInactiveV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.locationsGetInactiveV2(licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - Manage Locations
     *
     * GET GetTypes V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object locationsGetTypesV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.locationsGetTypesV1(licensenumber));
    }

    /**
     * Retrieves a list of active location types for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * GET GetTypes V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object locationsGetTypesV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.locationsGetTypesV2(licensenumber, pagenumber, pagesize));
    }

    public static class LocationsUpdateV2RequestItem {
        public Integer Id;
        public String LocationTypeName;
        public String Name;
    }

    /**
     * Updates existing locations for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Locations
     *
     * PUT Update V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object locationsUpdateV2(String licensenumber, List<LocationsUpdateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.locationsUpdateV2(licensenumber, body));
    }

    public static class PatientsCreateV2RequestItem {
        public String ActualDate;
        public Integer ConcentrateOuncesAllowed;
        public Integer FlowerOuncesAllowed;
        public Boolean HasSalesLimitExemption;
        public Integer InfusedOuncesAllowed;
        public String LicenseEffectiveEndDate;
        public String LicenseEffectiveStartDate;
        public String LicenseNumber;
        public Integer MaxConcentrateThcPercentAllowed;
        public Integer MaxFlowerThcPercentAllowed;
        public Integer RecommendedPlants;
        public Integer RecommendedSmokableQuantity;
        public Integer ThcOuncesAllowed;
    }

    /**
     * Adds new patients to a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * POST Create V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientsCreateV2(String licensenumber, List<PatientsCreateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.patientsCreateV2(licensenumber, body));
    }

    public static class PatientsCreateAddV1RequestItem {
        public String ActualDate;
        public Integer ConcentrateOuncesAllowed;
        public Integer FlowerOuncesAllowed;
        public Boolean HasSalesLimitExemption;
        public Integer InfusedOuncesAllowed;
        public String LicenseEffectiveEndDate;
        public String LicenseEffectiveStartDate;
        public String LicenseNumber;
        public Integer MaxConcentrateThcPercentAllowed;
        public Integer MaxFlowerThcPercentAllowed;
        public Integer RecommendedPlants;
        public Integer RecommendedSmokableQuantity;
        public Integer ThcOuncesAllowed;
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * POST CreateAdd V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientsCreateAddV1(String licensenumber, List<PatientsCreateAddV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.patientsCreateAddV1(licensenumber, body));
    }

    public static class PatientsCreateUpdateV1RequestItem {
        public String ActualDate;
        public Integer ConcentrateOuncesAllowed;
        public Integer FlowerOuncesAllowed;
        public Boolean HasSalesLimitExemption;
        public Integer InfusedOuncesAllowed;
        public String LicenseEffectiveEndDate;
        public String LicenseEffectiveStartDate;
        public String LicenseNumber;
        public Integer MaxConcentrateThcPercentAllowed;
        public Integer MaxFlowerThcPercentAllowed;
        public String NewLicenseNumber;
        public Integer RecommendedPlants;
        public Integer RecommendedSmokableQuantity;
        public Integer ThcOuncesAllowed;
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * POST CreateUpdate V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientsCreateUpdateV1(String licensenumber, List<PatientsCreateUpdateV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.patientsCreateUpdateV1(licensenumber, body));
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * DELETE Delete V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientsDeleteV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.patientsDeleteV1(id, licensenumber));
    }

    /**
     * Removes a Patient, identified by an Id, from a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * DELETE Delete V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientsDeleteV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.patientsDeleteV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * GET Get V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientsGetV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.patientsGetV1(id, licensenumber));
    }

    /**
     * Retrieves a Patient by Id.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * GET Get V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientsGetV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.patientsGetV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Patients
     *
     * GET GetActive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientsGetActiveV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.patientsGetActiveV1(licensenumber));
    }

    /**
     * Retrieves a list of active patients for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * GET GetActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientsGetActiveV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.patientsGetActiveV2(licensenumber, pagenumber, pagesize));
    }

    public static class PatientsUpdateV2RequestItem {
        public String ActualDate;
        public Integer ConcentrateOuncesAllowed;
        public Integer FlowerOuncesAllowed;
        public Boolean HasSalesLimitExemption;
        public Integer InfusedOuncesAllowed;
        public String LicenseEffectiveEndDate;
        public String LicenseEffectiveStartDate;
        public String LicenseNumber;
        public Integer MaxConcentrateThcPercentAllowed;
        public Integer MaxFlowerThcPercentAllowed;
        public String NewLicenseNumber;
        public Integer RecommendedPlants;
        public Integer RecommendedSmokableQuantity;
        public Integer ThcOuncesAllowed;
    }

    /**
     * Updates Patient information for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Patients
     *
     * PUT Update V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientsUpdateV2(String licensenumber, List<PatientsUpdateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.patientsUpdateV2(licensenumber, body));
    }

    public static class RetailidCreateAssociateV2RequestItem {
        public String PackageLabel;
        public List<String> QrUrls;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object retailIdCreateAssociateV2(String licensenumber, List<RetailidCreateAssociateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.retailIdCreateAssociateV2(licensenumber, body));
    }

    public static class RetailidCreateGenerateV2Request {
        public String PackageLabel;
        public Integer Quantity;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object retailIdCreateGenerateV2(String licensenumber, RetailidCreateGenerateV2Request body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.retailIdCreateGenerateV2(licensenumber, body));
    }

    public static class RetailidCreateMergeV2Request {
        public List<String> packageLabels;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object retailIdCreateMergeV2(String licensenumber, RetailidCreateMergeV2Request body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.retailIdCreateMergeV2(licensenumber, body));
    }

    public static class RetailidCreatePackageInfoV2Request {
        public List<String> packageLabels;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object retailIdCreatePackageInfoV2(String licensenumber, RetailidCreatePackageInfoV2Request body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.retailIdCreatePackageInfoV2(licensenumber, body));
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object retailIdGetReceiveByLabelV2(String label, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.retailIdGetReceiveByLabelV2(label, licensenumber));
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object retailIdGetReceiveQrByShortCodeV2(String shortcode, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.retailIdGetReceiveQrByShortCodeV2(shortcode, licensenumber));
    }

    /**
     * This endpoint is used to handle the setup of an external integrator for sandbox environments. It processes a request to create a new sandbox user for integration based on an external source's API key. It checks whether the API key is valid, manages the user creation process, and returns an appropriate status based on the current state of the request.
     * 
     *   Permissions Required:
     *   - None
     *
     * POST CreateIntegratorSetup V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object sandboxCreateIntegratorSetupV2(String userkey, Object body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.sandboxCreateIntegratorSetupV2(userkey, body));
    }

    public static class StrainsCreateV1RequestItem {
        public Double CbdLevel;
        public Double IndicaPercentage;
        public String Name;
        public Double SativaPercentage;
        public String TestingStatus;
        public Double ThcLevel;
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * POST Create V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object strainsCreateV1(String licensenumber, List<StrainsCreateV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.strainsCreateV1(licensenumber, body));
    }

    public static class StrainsCreateV2RequestItem {
        public Double CbdLevel;
        public Double IndicaPercentage;
        public String Name;
        public Double SativaPercentage;
        public String TestingStatus;
        public Double ThcLevel;
    }

    /**
     * Creates new strain records for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * POST Create V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object strainsCreateV2(String licensenumber, List<StrainsCreateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.strainsCreateV2(licensenumber, body));
    }

    public static class StrainsCreateUpdateV1RequestItem {
        public Double CbdLevel;
        public Integer Id;
        public Double IndicaPercentage;
        public String Name;
        public Double SativaPercentage;
        public String TestingStatus;
        public Double ThcLevel;
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * POST CreateUpdate V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object strainsCreateUpdateV1(String licensenumber, List<StrainsCreateUpdateV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.strainsCreateUpdateV1(licensenumber, body));
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * DELETE Delete V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object strainsDeleteV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.strainsDeleteV1(id, licensenumber));
    }

    /**
     * Archives an existing strain record for a Facility
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * DELETE Delete V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object strainsDeleteV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.strainsDeleteV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * GET Get V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object strainsGetV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.strainsGetV1(id, licensenumber));
    }

    /**
     * Retrieves a Strain record by its Id, with an optional license number.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * GET Get V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object strainsGetV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.strainsGetV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Strains
     *
     * GET GetActive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object strainsGetActiveV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.strainsGetActiveV1(licensenumber));
    }

    /**
     * Retrieves a list of active strains for the current Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * GET GetActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object strainsGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.strainsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Retrieves a list of inactive strains for the current Facility, optionally filtered by last modified date range.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * GET GetInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object strainsGetInactiveV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.strainsGetInactiveV2(licensenumber, pagenumber, pagesize));
    }

    public static class StrainsUpdateV2RequestItem {
        public Double CbdLevel;
        public Integer Id;
        public Double IndicaPercentage;
        public String Name;
        public Double SativaPercentage;
        public String TestingStatus;
        public Double ThcLevel;
    }

    /**
     * Updates existing strain records for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Strains
     *
     * PUT Update V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object strainsUpdateV2(String licensenumber, List<StrainsUpdateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.strainsUpdateV2(licensenumber, body));
    }

    /**
     * Returns a list of available package tags. NOTE: This is a premium endpoint.
     * 
     *   Permissions Required:
     *   - Manage Tags
     *
     * GET GetPackageAvailable V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object tagsGetPackageAvailableV2(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.tagsGetPackageAvailableV2(licensenumber));
    }

    /**
     * Returns a list of available plant tags. NOTE: This is a premium endpoint.
     * 
     *   Permissions Required:
     *   - Manage Tags
     *
     * GET GetPlantAvailable V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object tagsGetPlantAvailableV2(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.tagsGetPlantAvailableV2(licensenumber));
    }

    /**
     * Returns a list of staged tags. NOTE: This is a premium endpoint.
     * 
     *   Permissions Required:
     *   - Manage Tags
     *   - RetailId.AllowPackageStaging Key Value enabled
     *
     * GET GetStaged V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object tagsGetStagedV2(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.tagsGetStagedV2(licensenumber));
    }

    /**
     * Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Patients
     *
     * GET GetStatusesByPatientLicenseNumber V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientsStatusGetStatusesByPatientLicenseNumberV1(String patientlicensenumber, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.patientsStatusGetStatusesByPatientLicenseNumberV1(patientlicensenumber, licensenumber));
    }

    /**
     * Retrieves a list of statuses for a Patient License Number for a specified Facility. Data returned by this endpoint is cached for up to one minute.
     * 
     *   Permissions Required:
     *   - Lookup Patients
     *
     * GET GetStatusesByPatientLicenseNumber V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object patientsStatusGetStatusesByPatientLicenseNumberV2(String patientlicensenumber, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.patientsStatusGetStatusesByPatientLicenseNumberV2(patientlicensenumber, licensenumber));
    }

    public static class PlantbatchesCreateAdditivesV1RequestItem {
        public List<PlantbatchesCreateAdditivesV1RequestItem_ActiveIngredients> ActiveIngredients;
        public String ActualDate;
        public String AdditiveType;
        public String ApplicationDevice;
        public String EpaRegistrationNumber;
        public String PlantBatchName;
        public String ProductSupplier;
        public String ProductTradeName;
        public Integer TotalAmountApplied;
        public String TotalAmountUnitOfMeasure;
    }

    public static class PlantbatchesCreateAdditivesV1RequestItem_ActiveIngredients {
        public String Name;
        public Double Percentage;
    }

    /**
     * Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreateAdditivesV1(String licensenumber, List<PlantbatchesCreateAdditivesV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreateAdditivesV1(licensenumber, body));
    }

    public static class PlantbatchesCreateAdditivesV2RequestItem {
        public List<PlantbatchesCreateAdditivesV2RequestItem_ActiveIngredients> ActiveIngredients;
        public String ActualDate;
        public String AdditiveType;
        public String ApplicationDevice;
        public String EpaRegistrationNumber;
        public String PlantBatchName;
        public String ProductSupplier;
        public String ProductTradeName;
        public Integer TotalAmountApplied;
        public String TotalAmountUnitOfMeasure;
    }

    public static class PlantbatchesCreateAdditivesV2RequestItem_ActiveIngredients {
        public String Name;
        public Double Percentage;
    }

    /**
     * Records Additive usage details for plant batches at a specific Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditives V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreateAdditivesV2(String licensenumber, List<PlantbatchesCreateAdditivesV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreateAdditivesV2(licensenumber, body));
    }

    public static class PlantbatchesCreateAdditivesUsingtemplateV2RequestItem {
        public String ActualDate;
        public String AdditivesTemplateName;
        public String PlantBatchName;
        public String Rate;
        public Integer TotalAmountApplied;
        public String TotalAmountUnitOfMeasure;
        public String Volume;
    }

    /**
     * Records Additive usage for plant batches at a Facility using predefined additive templates.
     * 
     *   Permissions Required:
     *   - Manage Plants Additives
     *
     * POST CreateAdditivesUsingtemplate V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreateAdditivesUsingtemplateV2(String licensenumber, List<PlantbatchesCreateAdditivesUsingtemplateV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreateAdditivesUsingtemplateV2(licensenumber, body));
    }

    public static class PlantbatchesCreateAdjustV1RequestItem {
        public String AdjustmentDate;
        public String AdjustmentReason;
        public String PlantBatchName;
        public Integer Quantity;
        public String ReasonNote;
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateAdjust V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreateAdjustV1(String licensenumber, List<PlantbatchesCreateAdjustV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreateAdjustV1(licensenumber, body));
    }

    public static class PlantbatchesCreateAdjustV2RequestItem {
        public String AdjustmentDate;
        public String AdjustmentReason;
        public String PlantBatchName;
        public Integer Quantity;
        public String ReasonNote;
    }

    /**
     * Applies Facility specific adjustments to plant batches based on submitted reasons and input data.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateAdjust V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreateAdjustV2(String licensenumber, List<PlantbatchesCreateAdjustV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreateAdjustV2(licensenumber, body));
    }

    public static class PlantbatchesCreateChangegrowthphaseV1RequestItem {
        public Integer Count;
        public String CountPerPlant;
        public String GrowthDate;
        public String GrowthPhase;
        public String Name;
        public String NewLocation;
        public String NewSublocation;
        public String PatientLicenseNumber;
        public String StartingTag;
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * POST CreateChangegrowthphase V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreateChangegrowthphaseV1(String licensenumber, List<PlantbatchesCreateChangegrowthphaseV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreateChangegrowthphaseV1(licensenumber, body));
    }

    public static class PlantbatchesCreateGrowthphaseV2RequestItem {
        public Integer Count;
        public String CountPerPlant;
        public String GrowthDate;
        public String GrowthPhase;
        public String Name;
        public String NewLocation;
        public String NewSublocation;
        public String PatientLicenseNumber;
        public String StartingTag;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreateGrowthphaseV2(String licensenumber, List<PlantbatchesCreateGrowthphaseV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreateGrowthphaseV2(licensenumber, body));
    }

    public static class PlantbatchesCreatePackageV2RequestItem {
        public String ActualDate;
        public Integer Count;
        public String ExpirationDate;
        public Integer Id;
        public Boolean IsDonation;
        public Boolean IsTradeSample;
        public String Item;
        public String Location;
        public String Note;
        public String PatientLicenseNumber;
        public String PlantBatch;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UseByDate;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreatePackageV2(String isfrommotherplant, String licensenumber, List<PlantbatchesCreatePackageV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreatePackageV2(isfrommotherplant, licensenumber, body));
    }

    public static class PlantbatchesCreatePackageFrommotherplantV1RequestItem {
        public String ActualDate;
        public Integer Count;
        public String ExpirationDate;
        public Integer Id;
        public Boolean IsDonation;
        public Boolean IsTradeSample;
        public String Item;
        public String Location;
        public String Note;
        public String PatientLicenseNumber;
        public String PlantBatch;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UseByDate;
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST CreatePackageFrommotherplant V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreatePackageFrommotherplantV1(String licensenumber, List<PlantbatchesCreatePackageFrommotherplantV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreatePackageFrommotherplantV1(licensenumber, body));
    }

    public static class PlantbatchesCreatePackageFrommotherplantV2RequestItem {
        public String ActualDate;
        public Integer Count;
        public String ExpirationDate;
        public Integer Id;
        public Boolean IsDonation;
        public Boolean IsTradeSample;
        public String Item;
        public String Location;
        public String Note;
        public String PatientLicenseNumber;
        public String PlantBatch;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UseByDate;
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
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreatePackageFrommotherplantV2(String licensenumber, List<PlantbatchesCreatePackageFrommotherplantV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreatePackageFrommotherplantV2(licensenumber, body));
    }

    public static class PlantbatchesCreatePlantingsV2RequestItem {
        public String ActualDate;
        public Integer Count;
        public String Location;
        public String Name;
        public String PatientLicenseNumber;
        public String SourcePlantBatches;
        public String Strain;
        public String Sublocation;
        public String Type;
    }

    /**
     * Creates new plantings for a Facility by generating plant batches based on provided planting details.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreatePlantings V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreatePlantingsV2(String licensenumber, List<PlantbatchesCreatePlantingsV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreatePlantingsV2(licensenumber, body));
    }

    public static class PlantbatchesCreateSplitV1RequestItem {
        public String ActualDate;
        public Integer Count;
        public String GroupName;
        public String Location;
        public String PatientLicenseNumber;
        public String PlantBatch;
        public String Strain;
        public String Sublocation;
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateSplit V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreateSplitV1(String licensenumber, List<PlantbatchesCreateSplitV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreateSplitV1(licensenumber, body));
    }

    public static class PlantbatchesCreateSplitV2RequestItem {
        public String ActualDate;
        public Integer Count;
        public String GroupName;
        public String Location;
        public String PatientLicenseNumber;
        public String PlantBatch;
        public String Strain;
        public String Sublocation;
    }

    /**
     * Splits an existing Plant Batch into multiple groups at the specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST CreateSplit V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreateSplitV2(String licensenumber, List<PlantbatchesCreateSplitV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreateSplitV2(licensenumber, body));
    }

    public static class PlantbatchesCreateWasteV1RequestItem {
        public String MixedMaterial;
        public String Note;
        public String PlantBatchName;
        public String ReasonName;
        public String UnitOfMeasureName;
        public String WasteDate;
        public String WasteMethodName;
        public Double WasteWeight;
    }

    /**
     * Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreateWasteV1(String licensenumber, List<PlantbatchesCreateWasteV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreateWasteV1(licensenumber, body));
    }

    public static class PlantbatchesCreateWasteV2RequestItem {
        public String MixedMaterial;
        public String Note;
        public String PlantBatchName;
        public String ReasonName;
        public String UnitOfMeasureName;
        public String WasteDate;
        public String WasteMethodName;
        public Double WasteWeight;
    }

    /**
     * Records waste information for plant batches based on the submitted data for the specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Plants Waste
     *
     * POST CreateWaste V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreateWasteV2(String licensenumber, List<PlantbatchesCreateWasteV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreateWasteV2(licensenumber, body));
    }

    public static class PlantbatchesCreatepackagesV1RequestItem {
        public String ActualDate;
        public Integer Count;
        public String ExpirationDate;
        public Integer Id;
        public Boolean IsDonation;
        public Boolean IsTradeSample;
        public String Item;
        public String Location;
        public String Note;
        public String PatientLicenseNumber;
        public String PlantBatch;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UseByDate;
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *   - View Packages
     *   - Create/Submit/Discontinue Packages
     *
     * POST Createpackages V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreatepackagesV1(String isfrommotherplant, String licensenumber, List<PlantbatchesCreatepackagesV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreatepackagesV1(isfrommotherplant, licensenumber, body));
    }

    public static class PlantbatchesCreateplantingsV1RequestItem {
        public String ActualDate;
        public Integer Count;
        public String Location;
        public String Name;
        public String PatientLicenseNumber;
        public String SourcePlantBatches;
        public String Strain;
        public String Sublocation;
        public String Type;
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants Inventory
     *
     * POST Createplantings V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesCreateplantingsV1(String licensenumber, List<PlantbatchesCreateplantingsV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesCreateplantingsV1(licensenumber, body));
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *   - Destroy Immature Plants
     *
     * DELETE Delete V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesDeleteV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesDeleteV1(licensenumber));
    }

    /**
     * Completes the destruction of plant batches based on the provided input data.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Destroy Immature Plants
     *
     * DELETE Delete V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesDeleteV2(String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesDeleteV2(licensenumber));
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * GET Get V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesGetV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantBatchesGetV1(id, licensenumber));
    }

    /**
     * Retrieves a Plant Batch by Id.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *
     * GET Get V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesGetV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantBatchesGetV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * GET GetActive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesGetActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantBatchesGetActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves a list of active plant batches for the specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *
     * GET GetActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantBatchesGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * GET GetInactive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesGetInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantBatchesGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves a list of inactive plant batches for the specified Facility, optionally filtered by last modified date.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *
     * GET GetInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesGetInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantBatchesGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetTypes V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesGetTypesV1(String no) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantBatchesGetTypesV1(no));
    }

    /**
     * Retrieves a list of plant batch types.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetTypes V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesGetTypesV2(String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantBatchesGetTypesV2(pagenumber, pagesize));
    }

    /**
     * Retrieves waste details associated with plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Plants Waste
     *
     * GET GetWaste V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesGetWasteV2(String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantBatchesGetWasteV2(licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesGetWasteReasonsV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantBatchesGetWasteReasonsV1(licensenumber));
    }

    /**
     * Retrieves a list of valid waste reasons associated with immature plant batches for the specified Facility.
     * 
     *   Permissions Required:
     *   - None
     *
     * GET GetWasteReasons V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesGetWasteReasonsV2(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.plantBatchesGetWasteReasonsV2(licensenumber));
    }

    public static class PlantbatchesUpdateLocationV2RequestItem {
        public String Location;
        public String MoveDate;
        public String Name;
        public String Sublocation;
    }

    /**
     * Moves one or more plant batches to new locations with in a specified Facility.
     * 
     *   Permissions Required:
     *   - View Immature Plants
     *   - Manage Immature Plants
     *
     * PUT UpdateLocation V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesUpdateLocationV2(String licensenumber, List<PlantbatchesUpdateLocationV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesUpdateLocationV2(licensenumber, body));
    }

    public static class PlantbatchesUpdateMoveplantbatchesV1RequestItem {
        public String Location;
        public String MoveDate;
        public String Name;
        public String Sublocation;
    }

    /**
     * Permissions Required:
     *   - View Immature Plants
     *
     * PUT UpdateMoveplantbatches V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesUpdateMoveplantbatchesV1(String licensenumber, List<PlantbatchesUpdateMoveplantbatchesV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesUpdateMoveplantbatchesV1(licensenumber, body));
    }

    public static class PlantbatchesUpdateNameV2RequestItem {
        public String Group;
        public Integer Id;
        public String NewGroup;
    }

    /**
     * Renames plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateName V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesUpdateNameV2(String licensenumber, List<PlantbatchesUpdateNameV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesUpdateNameV2(licensenumber, body));
    }

    public static class PlantbatchesUpdateStrainV2RequestItem {
        public Integer Id;
        public String Name;
        public Integer StrainId;
        public String StrainName;
    }

    /**
     * Changes the strain of plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateStrain V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesUpdateStrainV2(String licensenumber, List<PlantbatchesUpdateStrainV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesUpdateStrainV2(licensenumber, body));
    }

    public static class PlantbatchesUpdateTagV2RequestItem {
        public String Group;
        public Integer Id;
        public String NewTag;
        public String ReplaceDate;
        public Integer TagId;
    }

    /**
     * Replaces tags for plant batches at a specified Facility.
     * 
     *   Permissions Required:
     *   - View Veg/Flower Plants
     *   - Manage Veg/Flower Plants Inventory
     *
     * PUT UpdateTag V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object plantBatchesUpdateTagV2(String licensenumber, List<PlantbatchesUpdateTagV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.plantBatchesUpdateTagV2(licensenumber, body));
    }

    public static class ProcessingjobsCreateAdjustV1RequestItem {
        public String AdjustmentDate;
        public String AdjustmentNote;
        public String AdjustmentReason;
        public String CountUnitOfMeasureName;
        public Integer Id;
        public List<ProcessingjobsCreateAdjustV1RequestItem_Packages> Packages;
        public String VolumeUnitOfMeasureName;
        public String WeightUnitOfMeasureName;
    }

    public static class ProcessingjobsCreateAdjustV1RequestItem_Packages {
        public String Label;
        public Integer Quantity;
        public String UnitOfMeasure;
    }

    /**
     * Permissions Required:
     *   - ManageProcessingJobs
     *
     * POST CreateAdjust V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsCreateAdjustV1(String licensenumber, List<ProcessingjobsCreateAdjustV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsCreateAdjustV1(licensenumber, body));
    }

    public static class ProcessingjobsCreateAdjustV2RequestItem {
        public String AdjustmentDate;
        public String AdjustmentNote;
        public String AdjustmentReason;
        public String CountUnitOfMeasureName;
        public Integer Id;
        public List<ProcessingjobsCreateAdjustV2RequestItem_Packages> Packages;
        public String VolumeUnitOfMeasureName;
        public String WeightUnitOfMeasureName;
    }

    public static class ProcessingjobsCreateAdjustV2RequestItem_Packages {
        public String Label;
        public Integer Quantity;
        public String UnitOfMeasure;
    }

    /**
     * Adjusts the details of existing processing jobs at a Facility, including units of measure and associated packages.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateAdjust V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsCreateAdjustV2(String licensenumber, List<ProcessingjobsCreateAdjustV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsCreateAdjustV2(licensenumber, body));
    }

    public static class ProcessingjobsCreateJobtypesV1RequestItem {
        public List<String> Attributes;
        public String Category;
        public String Description;
        public String Name;
        public String ProcessingSteps;
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateJobtypes V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsCreateJobtypesV1(String licensenumber, List<ProcessingjobsCreateJobtypesV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsCreateJobtypesV1(licensenumber, body));
    }

    public static class ProcessingjobsCreateJobtypesV2RequestItem {
        public List<String> Attributes;
        public String Category;
        public String Description;
        public String Name;
        public String ProcessingSteps;
    }

    /**
     * Creates new processing job types for a Facility, including name, category, description, steps, and attributes.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateJobtypes V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsCreateJobtypesV2(String licensenumber, List<ProcessingjobsCreateJobtypesV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsCreateJobtypesV2(licensenumber, body));
    }

    public static class ProcessingjobsCreateStartV1RequestItem {
        public String CountUnitOfMeasure;
        public String JobName;
        public String JobType;
        public List<ProcessingjobsCreateStartV1RequestItem_Packages> Packages;
        public String StartDate;
        public String VolumeUnitOfMeasure;
        public String WeightUnitOfMeasure;
    }

    public static class ProcessingjobsCreateStartV1RequestItem_Packages {
        public String Label;
        public Integer Quantity;
        public String UnitOfMeasure;
    }

    /**
     * Permissions Required:
     *   - ManageProcessingJobs
     *
     * POST CreateStart V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsCreateStartV1(String licensenumber, List<ProcessingjobsCreateStartV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsCreateStartV1(licensenumber, body));
    }

    public static class ProcessingjobsCreateStartV2RequestItem {
        public String CountUnitOfMeasure;
        public String JobName;
        public String JobType;
        public List<ProcessingjobsCreateStartV2RequestItem_Packages> Packages;
        public String StartDate;
        public String VolumeUnitOfMeasure;
        public String WeightUnitOfMeasure;
    }

    public static class ProcessingjobsCreateStartV2RequestItem_Packages {
        public String Label;
        public Integer Quantity;
        public String UnitOfMeasure;
    }

    /**
     * Initiates new processing jobs at a Facility, including job details and associated packages.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST CreateStart V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsCreateStartV2(String licensenumber, List<ProcessingjobsCreateStartV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsCreateStartV2(licensenumber, body));
    }

    public static class ProcessingjobsCreatepackagesV1RequestItem {
        public String ExpirationDate;
        public String FinishDate;
        public String FinishNote;
        public Boolean FinishProcessingJob;
        public String Item;
        public String JobName;
        public String Location;
        public String Note;
        public String PackageDate;
        public String PatientLicenseNumber;
        public String ProductionBatchNumber;
        public Integer Quantity;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UnitOfMeasure;
        public String UseByDate;
        public String WasteCountQuantity;
        public String WasteCountUnitOfMeasureName;
        public String WasteVolumeQuantity;
        public String WasteVolumeUnitOfMeasureName;
        public String WasteWeightQuantity;
        public String WasteWeightUnitOfMeasureName;
    }

    /**
     * Permissions Required:
     *   - ManageProcessingJobs
     *
     * POST Createpackages V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsCreatepackagesV1(String licensenumber, List<ProcessingjobsCreatepackagesV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsCreatepackagesV1(licensenumber, body));
    }

    public static class ProcessingjobsCreatepackagesV2RequestItem {
        public String ExpirationDate;
        public String FinishDate;
        public String FinishNote;
        public Boolean FinishProcessingJob;
        public String Item;
        public String JobName;
        public String Location;
        public String Note;
        public String PackageDate;
        public String PatientLicenseNumber;
        public String ProductionBatchNumber;
        public Integer Quantity;
        public String SellByDate;
        public String Sublocation;
        public String Tag;
        public String UnitOfMeasure;
        public String UseByDate;
        public String WasteCountQuantity;
        public String WasteCountUnitOfMeasureName;
        public String WasteVolumeQuantity;
        public String WasteVolumeUnitOfMeasureName;
        public String WasteWeightQuantity;
        public String WasteWeightUnitOfMeasureName;
    }

    /**
     * Creates packages from processing jobs at a Facility, including optional location and note assignments.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * POST Createpackages V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsCreatepackagesV2(String licensenumber, List<ProcessingjobsCreatepackagesV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsCreatepackagesV2(licensenumber, body));
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE Delete V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsDeleteV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsDeleteV1(id, licensenumber));
    }

    /**
     * Archives a Processing Job at a Facility by marking it as inactive and removing it from active use.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE Delete V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsDeleteV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsDeleteV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE DeleteJobtypes V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsDeleteJobtypesV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsDeleteJobtypesV1(id, licensenumber));
    }

    /**
     * Archives a Processing Job Type at a Facility, making it inactive for future use.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * DELETE DeleteJobtypes V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsDeleteJobtypesV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsDeleteJobtypesV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET Get V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetV1(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetV1(id, licensenumber));
    }

    /**
     * Retrieves a ProcessingJob by Id.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET Get V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetV2(String id, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetV2(id, licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetActive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves active processing jobs at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetInactive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves inactive processing jobs at a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesActive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesActiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetJobtypesActiveV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves a list of all active processing job types defined within a Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesActive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesActiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetJobtypesActiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesAttributes V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesAttributesV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetJobtypesAttributesV1(licensenumber));
    }

    /**
     * Retrieves all processing job attributes available for a Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesAttributes V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesAttributesV2(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetJobtypesAttributesV2(licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesCategories V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesCategoriesV1(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetJobtypesCategoriesV1(licensenumber));
    }

    /**
     * Retrieves all processing job categories available for a specified Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesCategories V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesCategoriesV2(String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetJobtypesCategoriesV2(licensenumber));
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesInactive V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesInactiveV1(String lastmodifiedend, String lastmodifiedstart, String licensenumber) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetJobtypesInactiveV1(lastmodifiedend, lastmodifiedstart, licensenumber));
    }

    /**
     * Retrieves a list of all inactive processing job types defined within a Facility.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * GET GetJobtypesInactive V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsGetJobtypesInactiveV2(String lastmodifiedend, String lastmodifiedstart, String licensenumber, String pagenumber, String pagesize) throws Exception {
        return rateLimiter.execute(null, true, () -> client.processingJobsGetJobtypesInactiveV2(lastmodifiedend, lastmodifiedstart, licensenumber, pagenumber, pagesize));
    }

    public static class ProcessingjobsUpdateFinishV1RequestItem {
        public String FinishDate;
        public String FinishNote;
        public Integer Id;
        public String TotalCountWaste;
        public String TotalVolumeWaste;
        public Integer TotalWeightWaste;
        public String WasteCountUnitOfMeasureName;
        public String WasteVolumeUnitOfMeasureName;
        public String WasteWeightUnitOfMeasureName;
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateFinish V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsUpdateFinishV1(String licensenumber, List<ProcessingjobsUpdateFinishV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsUpdateFinishV1(licensenumber, body));
    }

    public static class ProcessingjobsUpdateFinishV2RequestItem {
        public String FinishDate;
        public String FinishNote;
        public Integer Id;
        public String TotalCountWaste;
        public String TotalVolumeWaste;
        public Integer TotalWeightWaste;
        public String WasteCountUnitOfMeasureName;
        public String WasteVolumeUnitOfMeasureName;
        public String WasteWeightUnitOfMeasureName;
    }

    /**
     * Completes processing jobs at a Facility by recording final notes and waste measurements.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateFinish V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsUpdateFinishV2(String licensenumber, List<ProcessingjobsUpdateFinishV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsUpdateFinishV2(licensenumber, body));
    }

    public static class ProcessingjobsUpdateJobtypesV1RequestItem {
        public List<String> Attributes;
        public String CategoryName;
        public String Description;
        public Integer Id;
        public String Name;
        public String ProcessingSteps;
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateJobtypes V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsUpdateJobtypesV1(String licensenumber, List<ProcessingjobsUpdateJobtypesV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsUpdateJobtypesV1(licensenumber, body));
    }

    public static class ProcessingjobsUpdateJobtypesV2RequestItem {
        public List<String> Attributes;
        public String CategoryName;
        public String Description;
        public Integer Id;
        public String Name;
        public String ProcessingSteps;
    }

    /**
     * Updates existing processing job types at a Facility, including their name, category, description, steps, and attributes.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateJobtypes V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsUpdateJobtypesV2(String licensenumber, List<ProcessingjobsUpdateJobtypesV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsUpdateJobtypesV2(licensenumber, body));
    }

    public static class ProcessingjobsUpdateUnfinishV1RequestItem {
        public Integer Id;
    }

    /**
     * Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateUnfinish V1
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsUpdateUnfinishV1(String licensenumber, List<ProcessingjobsUpdateUnfinishV1RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsUpdateUnfinishV1(licensenumber, body));
    }

    public static class ProcessingjobsUpdateUnfinishV2RequestItem {
        public Integer Id;
    }

    /**
     * Reopens previously completed processing jobs at a Facility to allow further updates or corrections.
     * 
     *   Permissions Required:
     *   - Manage Processing Job
     *
     * PUT UpdateUnfinish V2
     * @throws Exception If request fails
     * @return Response object
     */
    public Object processingJobsUpdateUnfinishV2(String licensenumber, List<ProcessingjobsUpdateUnfinishV2RequestItem> body) throws Exception {
        return rateLimiter.execute(null, false, () -> client.processingJobsUpdateUnfinishV2(licensenumber, body));
    }

}
