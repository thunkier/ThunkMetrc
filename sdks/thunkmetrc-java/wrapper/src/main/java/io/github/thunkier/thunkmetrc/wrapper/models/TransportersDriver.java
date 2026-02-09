package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class TransportersDriver {
    @JsonProperty("DriversLicenseNumber")
    public String driversLicenseNumber;
    @JsonProperty("EmployeeId")
    public String employeeId;
    @JsonProperty("FacilityId")
    public Integer facilityId;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("IsArchived")
    public Boolean isArchived;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("Name")
    public String name;
}
