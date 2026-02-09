package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class RetailIdPackagesInfo {
    @JsonProperty("Packages")
    public List<RetailIdRetailIdPackagesInfoPackagesItem> packages;
    public static class RetailIdRetailIdPackagesInfoPackagesItemIssuancesItem {
    @JsonProperty("CreatedAt")
    public String createdAt;
    @JsonProperty("LabelQuantity")
    public Double labelQuantity;
    @JsonProperty("LabelSet")
    public Integer labelSet;
}public static class RetailIdRetailIdPackagesInfoPackagesItem {
    @JsonProperty("EstimatedBalance")
    public Integer estimatedBalance;
    @JsonProperty("HasQrs")
    public Boolean hasQrs;
    @JsonProperty("IssuanceId")
    public String issuanceId;
    @JsonProperty("Issuances")
    public List<RetailIdRetailIdPackagesInfoPackagesItemIssuancesItem> issuances;
    @JsonProperty("QrCount")
    public Integer qrCount;
    @JsonProperty("RequiresVerification")
    public Boolean requiresVerification;
    @JsonProperty("SiblingCount")
    public Integer siblingCount;
    @JsonProperty("Tag")
    public String tag;
}
}
