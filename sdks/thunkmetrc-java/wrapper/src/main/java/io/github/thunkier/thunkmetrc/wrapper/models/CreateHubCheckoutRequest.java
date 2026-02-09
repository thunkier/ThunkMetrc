package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreateHubCheckoutRequest {
    @JsonProperty("ShipmentDeliveryId")
    public Integer shipmentDeliveryId;
    @JsonProperty("TransporterDirection")
    public String transporterDirection;
}
