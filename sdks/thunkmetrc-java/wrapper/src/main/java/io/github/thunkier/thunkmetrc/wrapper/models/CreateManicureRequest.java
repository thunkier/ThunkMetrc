package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreateManicureRequest {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("DryingLocation")
    public String dryingLocation;
    @JsonProperty("DryingSublocation")
    public String dryingSublocation;
    @JsonProperty("HarvestName")
    public String harvestName;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("Plant")
    public String plant;
    @JsonProperty("PlantCount")
    public Integer plantCount;
    @JsonProperty("UnitOfWeight")
    public String unitOfWeight;
    @JsonProperty("Weight")
    public Double weight;
}
