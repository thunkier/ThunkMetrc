package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PackageSourceHarvest {
    @JsonProperty("HarvestStartDate")
    public String harvestStartDate;
    @JsonProperty("HarvestedByFacilityLicenseNumber")
    public String harvestedByFacilityLicenseNumber;
    @JsonProperty("HarvestedByFacilityName")
    public String harvestedByFacilityName;
    @JsonProperty("Name")
    public String name;
}
