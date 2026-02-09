package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PatientsCreatePatientsRequestItem {
    @JsonProperty("ActualDate")
    public String actualDate;
    @JsonProperty("ConcentrateOuncesAllowed")
    public Integer concentrateOuncesAllowed;
    @JsonProperty("FlowerOuncesAllowed")
    public Integer flowerOuncesAllowed;
    @JsonProperty("HasSalesLimitExemption")
    public Boolean hasSalesLimitExemption;
    @JsonProperty("InfusedOuncesAllowed")
    public Integer infusedOuncesAllowed;
    @JsonProperty("LicenseEffectiveEndDate")
    public String licenseEffectiveEndDate;
    @JsonProperty("LicenseEffectiveStartDate")
    public String licenseEffectiveStartDate;
    @JsonProperty("LicenseNumber")
    public String licenseNumber;
    @JsonProperty("MaxConcentrateThcPercentAllowed")
    public Integer maxConcentrateThcPercentAllowed;
    @JsonProperty("MaxFlowerThcPercentAllowed")
    public Integer maxFlowerThcPercentAllowed;
    @JsonProperty("RecommendedPlants")
    public Integer recommendedPlants;
    @JsonProperty("RecommendedSmokableQuantity")
    public Double recommendedSmokableQuantity;
    @JsonProperty("ThcOuncesAllowed")
    public Integer thcOuncesAllowed;
}
