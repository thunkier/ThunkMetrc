package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class PackagesUpdatePackageLabtestsRequiredRequestItem {
    @JsonProperty("Label")
    public String label;
    @JsonProperty("RequiredLabTestBatches")
    public List<String> requiredLabTestBatches;
}
