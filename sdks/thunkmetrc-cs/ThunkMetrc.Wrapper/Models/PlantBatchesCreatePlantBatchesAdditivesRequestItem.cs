using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantBatchesCreatePlantBatchesAdditivesRequestItem
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActiveIngredients")]
        public List<object>? ActiveIngredients { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ActualDate")]
        public string? ActualDate { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdditiveType")]
        public string? AdditiveType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ApplicationDevice")]
        public string? ApplicationDevice { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EpaRegistrationNumber")]
        public string? EpaRegistrationNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PlantBatchName")]
        public string? PlantBatchName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductSupplier")]
        public string? ProductSupplier { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductTradeName")]
        public string? ProductTradeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalAmountApplied")]
        public int? TotalAmountApplied { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalAmountUnitOfMeasure")]
        public string? TotalAmountUnitOfMeasure { get; set; }
        public class PlantBatchesCreatePlantBatchesAdditivesRequestItemActiveIngredientsItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
            public string? Name { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Percentage")]
            public int? Percentage { get; set; }
        }
    }
}
