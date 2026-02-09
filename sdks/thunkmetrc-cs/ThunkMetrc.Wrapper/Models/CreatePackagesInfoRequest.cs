using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class CreatePackagesInfoRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("packageLabels")]
        public List<object>? packageLabels { get; set; }
    }
}
