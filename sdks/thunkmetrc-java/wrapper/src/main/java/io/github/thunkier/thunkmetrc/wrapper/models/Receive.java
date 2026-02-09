package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class Receive {
    @JsonProperty("ChildTag")
    public String childTag;
    @JsonProperty("Eaches")
    public List<String> eaches;
    @JsonProperty("LabelSource")
    public String labelSource;
    @JsonProperty("QrCount")
    public Integer qrCount;
    @JsonProperty("Ranges")
    public List<List<Integer>> ranges;
    @JsonProperty("RequiresVerification")
    public Boolean requiresVerification;
    @JsonProperty("SiblingTags")
    public List<String> siblingTags;
}
