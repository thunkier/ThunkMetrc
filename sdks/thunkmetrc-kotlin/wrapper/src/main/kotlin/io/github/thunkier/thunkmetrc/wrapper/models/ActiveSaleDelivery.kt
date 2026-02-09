package io.github.thunkier.thunkmetrc.wrapper.models

import com.fasterxml.jackson.annotation.JsonProperty



data class ActiveSaleDelivery(
    @JsonProperty("AcceptedDateTime")
    val acceptedDateTime: String? = null,
    @JsonProperty("ActualArrivalDateTime")
    val actualArrivalDateTime: String? = null,
    @JsonProperty("ActualDepartureDateTime")
    val actualDepartureDateTime: String? = null,
    @JsonProperty("ActualReturnArrivalDateTime")
    val actualReturnArrivalDateTime: String? = null,
    @JsonProperty("ActualReturnDepartureDateTime")
    val actualReturnDepartureDateTime: String? = null,
    @JsonProperty("CompletedDateTime")
    val completedDateTime: String? = null,
    @JsonProperty("ConsumerId")
    val consumerId: String? = null,
    @JsonProperty("DeliveryNumber")
    val deliveryNumber: String? = null,
    @JsonProperty("Direction")
    val direction: String? = null,
    @JsonProperty("DriverEmployeeId")
    val driverEmployeeId: String? = null,
    @JsonProperty("DriverName")
    val driverName: String? = null,
    @JsonProperty("DriversLicenseNumber")
    val driversLicenseNumber: String? = null,
    @JsonProperty("EstimatedArrivalDateTime")
    val estimatedArrivalDateTime: String? = null,
    @JsonProperty("EstimatedDepartureDateTime")
    val estimatedDepartureDateTime: String? = null,
    @JsonProperty("EstimatedReturnArrivalDateTime")
    val estimatedReturnArrivalDateTime: String? = null,
    @JsonProperty("EstimatedReturnDepartureDateTime")
    val estimatedReturnDepartureDateTime: String? = null,
    @JsonProperty("FacilityLicenseNumber")
    val facilityLicenseNumber: String? = null,
    @JsonProperty("FacilityName")
    val facilityName: String? = null,
    @JsonProperty("Id")
    val id: Int? = null,
    @JsonProperty("LastModified")
    val lastModified: String? = null,
    @JsonProperty("PatientLicenseNumber")
    val patientLicenseNumber: String? = null,
    @JsonProperty("PlannedRoute")
    val plannedRoute: String? = null,
    @JsonProperty("RecipientAddressCity")
    val recipientAddressCity: String? = null,
    @JsonProperty("RecipientAddressCounty")
    val recipientAddressCounty: String? = null,
    @JsonProperty("RecipientAddressPostalCode")
    val recipientAddressPostalCode: String? = null,
    @JsonProperty("RecipientAddressState")
    val recipientAddressState: String? = null,
    @JsonProperty("RecipientAddressStreet1")
    val recipientAddressStreet1: String? = null,
    @JsonProperty("RecipientAddressStreet2")
    val recipientAddressStreet2: String? = null,
    @JsonProperty("RecipientDeliveryZoneId")
    val recipientDeliveryZoneId: Int? = null,
    @JsonProperty("RecipientDeliveryZoneName")
    val recipientDeliveryZoneName: String? = null,
    @JsonProperty("RecipientName")
    val recipientName: String? = null,
    @JsonProperty("RecipientZoneId")
    val recipientZoneId: Int? = null,
    @JsonProperty("RecordedByUserName")
    val recordedByUserName: String? = null,
    @JsonProperty("RecordedDateTime")
    val recordedDateTime: String? = null,
    @JsonProperty("SalesCustomerType")
    val salesCustomerType: String? = null,
    @JsonProperty("SalesDateTime")
    val salesDateTime: String? = null,
    @JsonProperty("SalesDeliveryState")
    val salesDeliveryState: String? = null,
    @JsonProperty("ShipperFacilityLicenseNumber")
    val shipperFacilityLicenseNumber: String? = null,
    @JsonProperty("TotalPackages")
    val totalPackages: Int? = null,
    @JsonProperty("TotalPrice")
    val totalPrice: Int? = null,
    @JsonProperty("Transactions")
    val transactions: List<Any>? = null,
    @JsonProperty("TransporterFacilityId")
    val transporterFacilityId: Int? = null,
    @JsonProperty("TransporterFacilityLicenseNumber")
    val transporterFacilityLicenseNumber: String? = null,
    @JsonProperty("TransporterFacilityName")
    val transporterFacilityName: String? = null,
    @JsonProperty("VehicleLicensePlateNumber")
    val vehicleLicensePlateNumber: String? = null,
    @JsonProperty("VehicleMake")
    val vehicleMake: String? = null,
    @JsonProperty("VehicleModel")
    val vehicleModel: String? = null,
    @JsonProperty("VoidedDate")
    val voidedDate: String? = null
)
