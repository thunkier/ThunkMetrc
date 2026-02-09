package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class HarvestsUpdateHarvestRenameRequestItem {
    @JsonProperty("Id")
    public Integer id;
    @JsonProperty("NewName")
    public String newName;
    @JsonProperty("OldName")
    public String oldName;
}
