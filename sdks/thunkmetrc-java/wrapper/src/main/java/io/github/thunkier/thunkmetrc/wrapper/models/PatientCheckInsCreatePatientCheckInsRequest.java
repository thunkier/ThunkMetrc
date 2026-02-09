package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PatientCheckInsCreatePatientCheckInsRequest {
    @JsonProperty("CheckInDate")
    public String checkInDate;
    @JsonProperty("CheckInLocationId")
    public Integer checkInLocationId;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("RegistrationExpiryDate")
    public String registrationExpiryDate;
    @JsonProperty("RegistrationStartDate")
    public String registrationStartDate;
}
