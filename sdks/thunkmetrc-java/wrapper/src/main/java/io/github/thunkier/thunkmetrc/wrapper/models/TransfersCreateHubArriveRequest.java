package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class TransfersCreateHubArriveRequest {
    @JsonProperty("ShipmentDeliveryId")
    public Integer shipmentDeliveryId;
    @JsonProperty("TransporterDirection")
    public String transporterDirection;
}
