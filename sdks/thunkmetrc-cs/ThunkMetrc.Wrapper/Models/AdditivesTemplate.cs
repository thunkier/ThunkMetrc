using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class AdditivesTemplate
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("ActiveIngredients")]
        public required List<AdditivesTemplateActiveIngredientsItem> ActiveIngredients { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdditiveType")]
        public required string AdditiveType { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("AdditiveTypeName")]
        public string? AdditiveTypeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ApplicationDevice")]
        public required string ApplicationDevice { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("EpaRegistrationNumber")]
        public required string EpaRegistrationNumber { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("FacilityId")]
        public required int FacilityId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Id")]
        public required int Id { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("IsActive")]
        public required bool IsActive { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LastModified")]
        public required string LastModified { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
        public string? Name { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("Note")]
        public string? Note { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductSupplier")]
        public required string ProductSupplier { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("ProductTradeName")]
        public required string ProductTradeName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RestrictiveEntryIntervalQuantityDescription")]
        public string? RestrictiveEntryIntervalQuantityDescription { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("RestrictiveEntryIntervalTimeDescription")]
        public string? RestrictiveEntryIntervalTimeDescription { get; set; }
        public class AdditivesTemplateActiveIngredientsItem
        {
            [global::System.Text.Json.Serialization.JsonPropertyName("Name")]
            public required string Name { get; set; }
            [global::System.Text.Json.Serialization.JsonPropertyName("Percentage")]
            public required double Percentage { get; set; }
        }
    }
}
