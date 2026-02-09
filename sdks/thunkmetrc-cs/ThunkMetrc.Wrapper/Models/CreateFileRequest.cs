using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class CreateFileRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("EncodedImageBase64")]
        public string? EncodedImageBase64 { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FileName")]
        public string? FileName { get; set; }
    }
}
