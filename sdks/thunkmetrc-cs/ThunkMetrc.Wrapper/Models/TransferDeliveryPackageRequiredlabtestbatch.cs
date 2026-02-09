using System;
using System.Collections.Generic;

namespace ThunkMetrc.Wrapper.Models
{
    public class TransferDeliveryPackageRequiredlabtestbatch
    {
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestBatchId")]
        public required int LabTestBatchId { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("LabTestBatchName")]
        public required string LabTestBatchName { get; set; }
        [global::System.Text.Json.Serialization.JsonPropertyName("PackageId")]
        public required int PackageId { get; set; }
    }
}
