using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    /// <summary>
    /// Generic wrapper for paginated API responses.
    /// </summary>
    public class PaginatedResponse<T>
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("Data")]
        public List<T> Data { get; set; } = new();

        [global::System.Text.Json.Serialization.JsonPropertyName("CurrentPage")]
        public int CurrentPage { get; set; }

        [global::System.Text.Json.Serialization.JsonPropertyName("Page")]
        public int Page { get; set; }

        [global::System.Text.Json.Serialization.JsonPropertyName("PageSize")]
        public int PageSize { get; set; }

        [global::System.Text.Json.Serialization.JsonPropertyName("RecordsOnPage")]
        public int RecordsOnPage { get; set; }

        [global::System.Text.Json.Serialization.JsonPropertyName("Total")]
        public int Total { get; set; }

        [global::System.Text.Json.Serialization.JsonPropertyName("TotalPages")]
        public int TotalPages { get; set; }

        [global::System.Text.Json.Serialization.JsonPropertyName("TotalRecords")]
        public int TotalRecords { get; set; }
    }
}
