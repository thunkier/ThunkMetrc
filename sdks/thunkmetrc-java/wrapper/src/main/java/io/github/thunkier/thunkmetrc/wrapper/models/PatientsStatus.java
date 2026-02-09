package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PatientsStatus {
    @JsonProperty("Active")
    public Boolean active;
    @JsonProperty("ConcentrateOuncesAllowed")
    public Integer concentrateOuncesAllowed;
    @JsonProperty("ConcentrateOuncesAvailable")
    public Integer concentrateOuncesAvailable;
    @JsonProperty("ConcentrateOuncesPurchased")
    public Integer concentrateOuncesPurchased;
    @JsonProperty("FlowerOuncesAllowed")
    public Integer flowerOuncesAllowed;
    @JsonProperty("FlowerOuncesAvailable")
    public Integer flowerOuncesAvailable;
    @JsonProperty("FlowerOuncesPurchased")
    public Integer flowerOuncesPurchased;
    @JsonProperty("IdentificationMethod")
    public String identificationMethod;
    @JsonProperty("InfusedOuncesAllowed")
    public Integer infusedOuncesAllowed;
    @JsonProperty("InfusedOuncesAvailable")
    public Integer infusedOuncesAvailable;
    @JsonProperty("InfusedOuncesPurchased")
    public Integer infusedOuncesPurchased;
    @JsonProperty("PatientLicenseExpirationDate")
    public String patientLicenseExpirationDate;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("PatientRegistrationStartDate")
    public String patientRegistrationStartDate;
    @JsonProperty("PurchaseAmountDays")
    public Integer purchaseAmountDays;
    @JsonProperty("RegistrationNumber")
    public String registrationNumber;
    @JsonProperty("ThcOuncesAllowed")
    public Integer thcOuncesAllowed;
    @JsonProperty("ThcOuncesAvailable")
    public Integer thcOuncesAvailable;
    @JsonProperty("ThcOuncesPurchased")
    public Integer thcOuncesPurchased;
}
