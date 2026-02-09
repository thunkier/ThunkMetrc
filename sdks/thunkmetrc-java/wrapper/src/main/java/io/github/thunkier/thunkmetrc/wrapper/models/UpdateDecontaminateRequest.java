package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class UpdateDecontaminateRequest {
    @JsonProperty("DecontaminationDate")
    public String decontaminationDate;
    @JsonProperty("DecontaminationMethodName")
    public String decontaminationMethodName;
    @JsonProperty("DecontaminationSteps")
    public String decontaminationSteps;
    @JsonProperty("PackageLabel")
    public String packageLabel;
}
