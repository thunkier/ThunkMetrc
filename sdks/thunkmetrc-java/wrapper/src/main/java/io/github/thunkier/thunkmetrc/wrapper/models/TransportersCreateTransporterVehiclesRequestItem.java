package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class TransportersCreateTransporterVehiclesRequestItem {
    @JsonProperty("LicensePlateNumber")
    public String licensePlateNumber;
    @JsonProperty("Make")
    public String make;
    @JsonProperty("Model")
    public String model;
}
