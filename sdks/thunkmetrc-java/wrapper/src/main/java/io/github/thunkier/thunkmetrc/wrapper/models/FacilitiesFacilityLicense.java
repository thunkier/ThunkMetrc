package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class FacilitiesFacilityLicense {
    @JsonProperty("EmployeeLicenseNumber")
    public String employeeLicenseNumber;
    @JsonProperty("EndDate")
    public String endDate;
    @JsonProperty("LicenseType")
    public String licenseType;
    @JsonProperty("Number")
    public String number;
    @JsonProperty("StartDate")
    public String startDate;
}
