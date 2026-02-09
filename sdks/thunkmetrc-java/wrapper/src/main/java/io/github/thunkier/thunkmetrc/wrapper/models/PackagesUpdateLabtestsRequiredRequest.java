package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class PackagesUpdateLabtestsRequiredRequest {
    @JsonProperty("Label")
    public String label;
    @JsonProperty("RequiredLabTestBatches")
    public List<Object> requiredLabTestBatches;
}
