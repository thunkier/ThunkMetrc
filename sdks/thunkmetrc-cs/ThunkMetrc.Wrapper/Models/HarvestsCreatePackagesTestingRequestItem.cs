using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class HarvestsCreatePackagesTestingRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DecontaminateProduct")]
        public bool? DecontaminateProduct { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DecontaminationDate")]
        public string? DecontaminationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("DecontaminationSteps")]
        public string? DecontaminationSteps { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ExpirationDate")]
        public string? ExpirationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Ingredients")]
        public List<object>? Ingredients { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsDonation")]
        public bool? IsDonation { get; set; }
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
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductRequiresDecontamination")]
        public bool? ProductRequiresDecontamination { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductRequiresRemediation")]
        public bool? ProductRequiresRemediation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductionBatchNumber")]
        public string? ProductionBatchNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RemediateProduct")]
        public bool? RemediateProduct { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RemediationDate")]
        public string? RemediationDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RemediationMethodId")]
        public int? RemediationMethodId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RemediationSteps")]
        public string? RemediationSteps { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RequiredLabTestBatches")]
        public List<object>? RequiredLabTestBatches { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SellByDate")]
        public string? SellByDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Sublocation")]
        public string? Sublocation { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Tag")]
        public string? Tag { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfWeight")]
        public string? UnitOfWeight { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("UseByDate")]
        public string? UseByDate { get; set; }
        public class HarvestsCreatePackagesTestingRequestItemIngredientsItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("HarvestId")]
            public int? HarvestId { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("HarvestName")]
            public string? HarvestName { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("UnitOfWeight")]
            public string? UnitOfWeight { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Weight")]
            public double? Weight { get; set; }
        }
    }
}
