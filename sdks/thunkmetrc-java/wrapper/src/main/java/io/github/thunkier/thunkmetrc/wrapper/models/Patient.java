package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class Patient {
    @JsonProperty("HasSalesLimitExemption")
    public Boolean hasSalesLimitExemption;
    @JsonProperty("LicenseEffectiveEndDate")
    public String licenseEffectiveEndDate;
    @JsonProperty("LicenseEffectiveStartDate")
    public String licenseEffectiveStartDate;
    @JsonProperty("LicenseNumber")
    public String licenseNumber;
    @JsonProperty("OtherFacilitiesCount")
    public Integer otherFacilitiesCount;
    @JsonProperty("PatientId")
    public Integer patientId;
    @JsonProperty("RecommendedPlants")
    public Integer recommendedPlants;
    @JsonProperty("RecommendedSmokableQuantity")
    public Double recommendedSmokableQuantity;
    @JsonProperty("RegistrationDate")
    public String registrationDate;
}
