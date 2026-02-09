using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantBatchesWaste
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchId")]
        public required int PlantBatchId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchName")]
        public required string PlantBatchName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantCount")]
        public required int PlantCount { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantWasteNumber")]
        public required string PlantWasteNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteDate")]
        public required string WasteDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteMethodName")]
        public required string WasteMethodName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteReasonName")]
        public required string WasteReasonName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteUnitOfMeasureName")]
        public required string WasteUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteWeight")]
        public required double WasteWeight { get; set; }
    }
}
