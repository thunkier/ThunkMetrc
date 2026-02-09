using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantsDeletePlantsRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Count")]
        public int? Count { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public int? Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Label")]
        public string? Label { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ReasonNote")]
        public string? ReasonNote { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteMaterialMixed")]
        public string? WasteMaterialMixed { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteMethodName")]
        public string? WasteMethodName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteReasonName")]
        public string? WasteReasonName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteUnitOfMeasureName")]
        public string? WasteUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteWeight")]
        public double? WasteWeight { get; set; }
    }
}
