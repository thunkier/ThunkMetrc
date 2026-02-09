package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class ProcessingJobCreateJobTypesRequestItem {
    @JsonProperty("Attributes")
    public List<Object> attributes;
    @JsonProperty("Category")
    public String category;
    @JsonProperty("Description")
    public String description;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("ProcessingSteps")
    public String processingSteps;
}
