package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class ItemFile {
    @JsonProperty("ContentType")
    public String contentType;
    @JsonProperty("FileContents")
    public String fileContents;
    @JsonProperty("FileDownloadName")
    public String fileDownloadName;
}
