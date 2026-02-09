package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class CreateAssociateRequest {
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("QrUrls")
    public List<Object> qrUrls;
}
