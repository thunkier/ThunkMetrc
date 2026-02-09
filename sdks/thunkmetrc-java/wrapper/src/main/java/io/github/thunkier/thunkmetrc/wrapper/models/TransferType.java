package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class TransferType {
    @JsonProperty("BypassApproval")
    public Boolean bypassApproval;
    @JsonProperty("ExternalIncomingCanRecordExternalIdentifier")
    public Boolean externalIncomingCanRecordExternalIdentifier;
    @JsonProperty("ExternalIncomingExternalIdentifierRequired")
    public Boolean externalIncomingExternalIdentifierRequired;
    @JsonProperty("ExternalOutgoingCanRecordExternalIdentifier")
    public Boolean externalOutgoingCanRecordExternalIdentifier;
    @JsonProperty("ExternalOutgoingExternalIdentifierRequired")
    public Boolean externalOutgoingExternalIdentifierRequired;
    @JsonProperty("ForExternalIncomingShipments")
    public Boolean forExternalIncomingShipments;
    @JsonProperty("ForExternalOutgoingShipments")
    public Boolean forExternalOutgoingShipments;
    @JsonProperty("ForLicensedShipments")
    public Boolean forLicensedShipments;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("RequiresDestinationGrossWeight")
    public Boolean requiresDestinationGrossWeight;
    @JsonProperty("RequiresInvoiceNumber")
    public Boolean requiresInvoiceNumber;
    @JsonProperty("RequiresPDFDocument")
    public Boolean requiresPDFDocument;
    @JsonProperty("RequiresPackagesGrossWeight")
    public Boolean requiresPackagesGrossWeight;
    @JsonProperty("TransactionType")
    public String transactionType;
}
