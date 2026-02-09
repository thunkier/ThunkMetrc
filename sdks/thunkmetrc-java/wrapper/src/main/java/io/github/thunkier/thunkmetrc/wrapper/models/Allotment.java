package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class Allotment {
    @JsonProperty("Allotment")
    public Integer allotment;
    @JsonProperty("FacilityId")
    public Integer facilityId;
    @JsonProperty("FacilityLicenseNumber")
    public String facilityLicenseNumber;
}
