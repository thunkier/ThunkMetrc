package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class InactiveJobType {
    @JsonProperty("Attributes")
    public List<ProcessingJobInactiveJobTypeAttributesItem> attributes;
    @JsonProperty("CategoryName")
    public String categoryName;
    @JsonProperty("Description")
    public String description;
    @JsonProperty("ForItems")
    public Boolean forItems;
    @JsonProperty("ForProcessingJobs")
    public Boolean forProcessingJobs;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("ProcessingSteps")
    public String processingSteps;
    @JsonProperty("RequiresProcessingJobAttributes")
    public Boolean requiresProcessingJobAttributes;
    public static class ProcessingJobInactiveJobTypeAttributesItem {
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("LastModified")
    public String lastModified;
    @JsonProperty("Name")
    public String name;
}
}
