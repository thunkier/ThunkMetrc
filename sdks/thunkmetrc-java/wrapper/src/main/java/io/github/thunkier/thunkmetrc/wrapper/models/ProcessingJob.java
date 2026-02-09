package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class ProcessingJob {
    @JsonProperty("CountUnitOfMeasureAbbreviation")
    public String countUnitOfMeasureAbbreviation;
    @JsonProperty("CountUnitOfMeasureId")
    public Integer countUnitOfMeasureId;
    @JsonProperty("CountUnitOfMeasureName")
    public String countUnitOfMeasureName;
    @JsonProperty("FinishNote")
    public String finishNote;
    @JsonProperty("FinishedDate")
    public String finishedDate;
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("IsFinished")
    public Boolean isFinished;
    @JsonProperty("JobTypeId")
    public Integer jobTypeId;
    @JsonProperty("JobTypeName")
    public String jobTypeName;
    @JsonProperty("Name")
    public String name;
    @JsonProperty("Number")
    public String number;
    @JsonProperty("Packages")
    public List<Object> packages;
    @JsonProperty("StartDate")
    public String startDate;
    @JsonProperty("TotalCount")
    public Integer totalCount;
    @JsonProperty("TotalCountWaste")
    public String totalCountWaste;
    @JsonProperty("TotalQuantity")
    public Double totalQuantity;
    @JsonProperty("TotalUnitOfMeasureId")
    public Integer totalUnitOfMeasureId;
    @JsonProperty("TotalVolume")
    public Double totalVolume;
    @JsonProperty("TotalVolumeWaste")
    public String totalVolumeWaste;
    @JsonProperty("TotalWeight")
    public Double totalWeight;
    @JsonProperty("TotalWeightWaste")
    public String totalWeightWaste;
    @JsonProperty("VolumeUnitOfMeasureAbbreviation")
    public String volumeUnitOfMeasureAbbreviation;
    @JsonProperty("VolumeUnitOfMeasureId")
    public Integer volumeUnitOfMeasureId;
    @JsonProperty("VolumeUnitOfMeasureName")
    public String volumeUnitOfMeasureName;
    @JsonProperty("WasteCountUnitOfMeasureAbbreviation")
    public String wasteCountUnitOfMeasureAbbreviation;
    @JsonProperty("WasteCountUnitOfMeasureId")
    public Integer wasteCountUnitOfMeasureId;
    @JsonProperty("WasteCountUnitOfMeasureName")
    public String wasteCountUnitOfMeasureName;
    @JsonProperty("WasteVolumeUnitOfMeasureAbbreviation")
    public String wasteVolumeUnitOfMeasureAbbreviation;
    @JsonProperty("WasteVolumeUnitOfMeasureId")
    public Integer wasteVolumeUnitOfMeasureId;
    @JsonProperty("WasteVolumeUnitOfMeasureName")
    public String wasteVolumeUnitOfMeasureName;
    @JsonProperty("WasteWeightUnitOfMeasureAbbreviation")
    public String wasteWeightUnitOfMeasureAbbreviation;
    @JsonProperty("WasteWeightUnitOfMeasureId")
    public Integer wasteWeightUnitOfMeasureId;
    @JsonProperty("WasteWeightUnitOfMeasureName")
    public String wasteWeightUnitOfMeasureName;
    @JsonProperty("WeightUnitOfMeasureAbbreviation")
    public String weightUnitOfMeasureAbbreviation;
    @JsonProperty("WeightUnitOfMeasureId")
    public Integer weightUnitOfMeasureId;
    @JsonProperty("WeightUnitOfMeasureName")
    public String weightUnitOfMeasureName;
}
