package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PackagesUpdatePackageUseByDateRequestItem {
    @JsonProperty("ExpirationDate")
    public String expirationDate;
    @JsonProperty("Label")
    public String label;
    @JsonProperty("SellByDate")
    public String sellByDate;
    @JsonProperty("UseByDate")
    public String useByDate;
}
