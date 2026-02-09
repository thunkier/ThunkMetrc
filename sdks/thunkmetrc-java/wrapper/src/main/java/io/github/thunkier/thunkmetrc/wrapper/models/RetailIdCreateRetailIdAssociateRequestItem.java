package io.github.thunkier.thunkmetrc.wrapper.models;import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;public class RetailIdCreateRetailIdAssociateRequestItem {
    @JsonProperty("PackageLabel")
    public String packageLabel;
    @JsonProperty("QrUrls")
    public List<String> qrUrls;
}
