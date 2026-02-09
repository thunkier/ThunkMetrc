package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class SaleReceipt {
    @JsonProperty("ArchivedDate")
    public String archivedDate;
    @JsonProperty("CaregiverLicenseNumber")
    public String caregiverLicenseNumber;
    @JsonProperty("ExternalReceiptNumber")
    public String externalReceiptNumber;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("IdentificationMethod")
    public String identificationMethod;
    @JsonProperty("IsFinal")
    public Boolean isFinal;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("PatientLicenseNumber")
    public String patientLicenseNumber;
    @JsonProperty("PatientRegistrationLocationId")
    public Integer patientRegistrationLocationId;
    @JsonProperty("ReceiptNumber")
    public String receiptNumber;
    @JsonProperty("RecordedByUserName")
    public String recordedByUserName;
    @JsonProperty("RecordedDateTime")
    public String recordedDateTime;
    @JsonProperty("SalesCustomerType")
    public String salesCustomerType;
    @JsonProperty("SalesDateTime")
    public String salesDateTime;
    @JsonProperty("TotalPackages")
    public Integer totalPackages;
    @JsonProperty("TotalPrice")
    public Integer totalPrice;
    @JsonProperty("Transactions")
    public List<Object> transactions;
}
