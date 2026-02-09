package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class Vehicle {
    @JsonProperty("FacilityId")
    public Integer facilityId;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("IsArchived")
    public Boolean isArchived;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("LicensePlateNumber")
    public String licensePlateNumber;
    @JsonProperty("Make")
    public String make;
    @JsonProperty("Model")
    public String model;
}
