package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class Employee {
    @JsonProperty("FullName")
    public String fullName;
    @JsonProperty("IsIndustryAdmin")
    public Boolean isIndustryAdmin;
    @JsonProperty("IsManager")
    public Boolean isManager;
    @JsonProperty("IsOwner")
    public Boolean isOwner;
    @JsonProperty("License")
    public EmployeeLicense license;
    public static class EmployeeLicense {
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
}
