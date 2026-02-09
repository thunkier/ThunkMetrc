package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class RetailIdCreateAssociateRequestItem {
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("QrUrls")
    public List<Object> qrUrls;
}
