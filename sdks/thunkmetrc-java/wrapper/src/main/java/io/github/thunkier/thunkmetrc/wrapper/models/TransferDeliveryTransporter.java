package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class TransferDeliveryTransporter {
    @JsonProperty("TransporterDirection")
    public String transporterDirection;
    @JsonProperty("TransporterFacilityLicenseNumber")
    public String transporterFacilityLicenseNumber;
    @JsonProperty("TransporterFacilityName")
    public String transporterFacilityName;
}
