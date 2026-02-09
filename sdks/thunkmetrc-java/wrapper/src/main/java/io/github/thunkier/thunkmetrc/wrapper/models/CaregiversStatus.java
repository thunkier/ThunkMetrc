package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class CaregiversStatus {
    @JsonProperty("Active")
    public Boolean active;
    @JsonProperty("CaregiverLicenseNumber")
    public String caregiverLicenseNumber;
    @JsonProperty("Patients")
    public List<String> patients;
}
