using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class CreateAdditivesTemplatesRequest
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActiveIngredients")]
        public List<object>? ActiveIngredients { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdditiveType")]
        public string? AdditiveType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ApplicationDevice")]
        public string? ApplicationDevice { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EpaRegistrationNumber")]
        public string? EpaRegistrationNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public string? Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Note")]
        public string? Note { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductSupplier")]
        public string? ProductSupplier { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductTradeName")]
        public string? ProductTradeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RestrictiveEntryIntervalQuantityDescription")]
        public string? RestrictiveEntryIntervalQuantityDescription { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RestrictiveEntryIntervalTimeDescription")]
        public string? RestrictiveEntryIntervalTimeDescription { get; set; }
        public class CreateAdditivesTemplatesRequestActiveIngredientsItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
            public string? Name { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Percentage")]
            public double? Percentage { get; set; }
        }
    }
}
