using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class File
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ContentType")]
        public required string ContentType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FileContents")]
        public required string FileContents { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FileDownloadName")]
        public required string FileDownloadName { get; set; }
    }
}
