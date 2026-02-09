using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class UpdateLabtestsRequiredRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public string? Label { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiredLabTestBatches")]
        public List<object>? RequiredLabTestBatches { get; set; }
    }
}
