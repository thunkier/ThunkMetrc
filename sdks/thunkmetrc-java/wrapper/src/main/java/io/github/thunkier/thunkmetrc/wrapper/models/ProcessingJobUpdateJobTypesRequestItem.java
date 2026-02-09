package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class ProcessingJobUpdateJobTypesRequestItem {
    @JsonProperty("Attributes")
    public List<Object> attributes;
    @JsonProperty("CategoryName")
    public String categoryName;
    @JsonProperty("Description")
    public String description;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("ProcessingSteps")
    public String processingSteps;
}
