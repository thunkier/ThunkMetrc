package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class SalesUpdateDeliveriesCompleteRequestItemReturnedPackagesItem {
    @JsonProperty("Label")
    public String label;
    @JsonProperty("ReturnQuantityVerified")
    public Integer returnQuantityVerified;
    @JsonProperty("ReturnReason")
    public String returnReason;
    @JsonProperty("ReturnReasonNote")
    public String returnReasonNote;
    @JsonProperty("ReturnUnitOfMeasure")
    public String returnUnitOfMeasure;
}
