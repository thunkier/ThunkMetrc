package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class SalesUpdateSaleDeliveriesCompleteRequestItem {
    @JsonProperty("AcceptedPackages")
    public List<String> acceptedPackages;
    @JsonProperty("ActualArrivalDateTime")
    public String actualArrivalDateTime;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("PaymentType")
    public String paymentType;
    @JsonProperty("ReturnedPackages")
    public List<SalesUpdateSaleDeliveriesCompleteRequestItemReturnedPackages> returnedPackages;
    public static class SalesUpdateSaleDeliveriesCompleteRequestItemReturnedPackages {
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
}
