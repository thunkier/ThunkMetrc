package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class TransfersCreateOutgoingTransferTemplatesRequestItemDestinationsPackages {
    @JsonProperty("GrossUnitOfWeightName")
    public String grossUnitOfWeightName;
    @JsonProperty("GrossWeight")
    public Double grossWeight;
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("WholesalePrice")
    public String wholesalePrice;
}
