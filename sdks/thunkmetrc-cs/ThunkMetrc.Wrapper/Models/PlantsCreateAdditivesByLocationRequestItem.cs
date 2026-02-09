using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class PlantsCreateAdditivesByLocationRequestItem
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
        [global::System.Text.Json.Serialization.JsonPropertyName("LocationName")]
        public string? LocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductSupplier")]
        public string? ProductSupplier { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductTradeName")]
        public string? ProductTradeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("SublocationName")]
        public string? SublocationName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalAmountApplied")]
        public int? TotalAmountApplied { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("TotalAmountUnitOfMeasure")]
        public string? TotalAmountUnitOfMeasure { get; set; }
        public class PlantsCreateAdditivesByLocationRequestItemActiveIngredientsItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
            public string? Name { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Percentage")]
            public int? Percentage { get; set; }
        }
    }
}
