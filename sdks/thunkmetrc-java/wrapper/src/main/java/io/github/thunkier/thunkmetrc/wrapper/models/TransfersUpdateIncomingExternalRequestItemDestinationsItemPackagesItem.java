package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class TransfersUpdateIncomingExternalRequestItemDestinationsItemPackagesItem {
    @JsonProperty("ExpirationDate")
    public String expirationDate;
    @JsonProperty("ExternalId")
    public String externalId;
    @JsonProperty("GrossUnitOfWeightName")
    public String grossUnitOfWeightName;
    @JsonProperty("GrossWeight")
    public Double grossWeight;
    @JsonProperty("IsFinishedGood")
    public Boolean isFinishedGood;
    @JsonProperty("ItemName")
    public String itemName;
    @JsonProperty("PackagedDate")
    public String packagedDate;
    @JsonProperty("Quantity")
    public Double quantity;
    @JsonProperty("SellByDate")
    public String sellByDate;
    @JsonProperty("UnitOfMeasureName")
    public String unitOfMeasureName;
    @JsonProperty("UseByDate")
    public String useByDate;
    @JsonProperty("WholesalePrice")
    public String wholesalePrice;
}
