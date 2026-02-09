package io.github.thunkier.thunkmetrc.wrapper.models;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;

public class PaginatedResponse<T> {
    @JsonProperty("Total")
    public Integer Total;
    @JsonProperty("PageSize")
    public Integer PageSize;
    @JsonProperty("Current")
    public Integer Current;
    @JsonProperty("Next")
    public Integer Next;
    @JsonProperty("Previous")
    public Integer Previous;
    @JsonProperty("Data")
    public List<T> Data;
}
