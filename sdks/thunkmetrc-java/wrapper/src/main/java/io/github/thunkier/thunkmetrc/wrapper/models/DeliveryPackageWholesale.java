package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class DeliveryPackageWholesale {
    @JsonProperty("PackageId")
    public Integer packageId;
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("ReceiverWholesalePrice")
    public String receiverWholesalePrice;
    @JsonProperty("ShipperWholesalePrice")
    public String shipperWholesalePrice;
}
