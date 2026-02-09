package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreateDriversRequest {
    @JsonProperty("DriversLicenseNumber")
    public String driversLicenseNumber;
    @JsonProperty("EmployeeId")
    public String employeeId;
    @JsonProperty("Name")
    public String name;
}
