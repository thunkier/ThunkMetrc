using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class ProcessingJobCreateProcessingJobPackagesRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ExpirationDate")]
        public string? ExpirationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FinishDate")]
        public string? FinishDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FinishNote")]
        public string? FinishNote { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FinishProcessingJob")]
        public bool? FinishProcessingJob { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsFinishedGood")]
        public bool? IsFinishedGood { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Item")]
        public string? Item { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("JobName")]
        public string? JobName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Location")]
        public string? Location { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Note")]
        public string? Note { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageDate")]
        public string? PackageDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductionBatchNumber")]
        public string? ProductionBatchNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Quantity")]
        public double? Quantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SellByDate")]
        public string? SellByDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Sublocation")]
        public string? Sublocation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Tag")]
        public string? Tag { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfMeasure")]
        public string? UnitOfMeasure { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UseByDate")]
        public string? UseByDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteCountQuantity")]
        public double? WasteCountQuantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteCountUnitOfMeasureName")]
        public string? WasteCountUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteVolumeQuantity")]
        public double? WasteVolumeQuantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteVolumeUnitOfMeasureName")]
        public string? WasteVolumeUnitOfMeasureName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteWeightQuantity")]
        public double? WasteWeightQuantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("WasteWeightUnitOfMeasureName")]
        public string? WasteWeightUnitOfMeasureName { get; set; }
    }
}
