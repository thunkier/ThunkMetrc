package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class ProcessingJobStartProcessingJobRequestItem {
    @JsonProperty("CountUnitOfMeasure")
    public String countUnitOfMeasure;
    @JsonProperty("JobName")
    public String jobName;
    @JsonProperty("JobType")
    public String jobType;
    @JsonProperty("Packages")
    public List<Object> packages;
    @JsonProperty("StartDate")
    public String startDate;
    @JsonProperty("VolumeUnitOfMeasure")
    public String volumeUnitOfMeasure;
    @JsonProperty("WeightUnitOfMeasure")
    public String weightUnitOfMeasure;
    public static class ProcessingJobStartProcessingJobRequestItemPackagesItem {
    @JsonProperty("Label")
    public String label;
    @JsonProperty("Quantity")
    public Double quantity;
    @JsonProperty("UnitOfMeasure")
    public String unitOfMeasure;
}
}
