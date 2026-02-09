package io.github.thunkier.thunkmetrc.wrapper.models;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreateFileRequest {
    @JsonProperty("EncodedImageBase64")
    public String encodedImageBase64;
    @JsonProperty("FileName")
    public String fileName;
}
