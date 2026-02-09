using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PackagesCreatePackagesPackagesRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExpirationDate")]
        public string? ExpirationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Ingredients")]
        public List<object>? Ingredients { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsDonation")]
        public bool? IsDonation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsFinishedGood")]
        public bool? IsFinishedGood { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsProductionBatch")]
        public bool? IsProductionBatch { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsTradeSample")]
        public bool? IsTradeSample { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Item")]
        public string? Item { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestStageId")]
        public int? LabTestStageId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Location")]
        public string? Location { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Note")]
        public string? Note { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PatientLicenseNumber")]
        public string? PatientLicenseNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProcessingJobTypeId")]
        public int? ProcessingJobTypeId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductRequiresRemediation")]
        public bool? ProductRequiresRemediation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductionBatchNumber")]
        public string? ProductionBatchNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Quantity")]
        public double? Quantity { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiredLabTestBatches")]
        public bool? RequiredLabTestBatches { get; set; }
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
        [global::System.Text.Json.Serialization.JsonPropertyName("UseSameItem")]
        public bool? UseSameItem { get; set; }
        public class PackagesCreatePackagesPackagesRequestItemIngredientsItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("Package")]
            public string? Package { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Quantity")]
            public double? Quantity { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfMeasure")]
            public string? UnitOfMeasure { get; set; }
        }
    }
}
