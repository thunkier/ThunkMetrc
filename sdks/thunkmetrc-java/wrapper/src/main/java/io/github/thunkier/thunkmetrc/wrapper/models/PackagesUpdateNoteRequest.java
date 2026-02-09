package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class PackagesUpdateNoteRequest {
    @JsonProperty("Note")
    public String note;
    @JsonProperty("PackageLabel")
    public String packageLabel;
}
