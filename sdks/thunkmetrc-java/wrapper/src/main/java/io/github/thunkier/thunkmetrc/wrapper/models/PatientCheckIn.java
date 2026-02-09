package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PatientCheckIn {
    @JsonProperty("CheckInDate")
    public String checkInDate;
    @JsonProperty("CheckInLocationId")
    public Integer checkInLocationId;
    @JsonProperty("CheckInLocationName")
    public String checkInLocationName;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("RegistrationExpiryDate")
    public String registrationExpiryDate;
    @JsonProperty("RegistrationStartDate")
    public String registrationStartDate;
}
