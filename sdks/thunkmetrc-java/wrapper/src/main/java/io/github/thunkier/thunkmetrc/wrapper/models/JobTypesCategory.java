package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class JobTypesCategory {
    @JsonProperty("ForItems")
    public Boolean forItems;
    @JsonProperty("ForProcessingJobs")
    public Boolean forProcessingJobs;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("RequiresProcessingJobAttributes")
    public Boolean requiresProcessingJobAttributes;
}
