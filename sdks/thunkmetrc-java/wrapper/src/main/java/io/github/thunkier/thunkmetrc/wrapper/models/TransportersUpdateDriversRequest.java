package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class TransportersUpdateDriversRequest {
    @JsonProperty("DriversLicenseNumber")
    public String driversLicenseNumber;
    @JsonProperty("EmployeeId")
    public String employeeId;
    @JsonProperty("Id")
    public String id;
    @JsonProperty("Name")
    public String name;
}
