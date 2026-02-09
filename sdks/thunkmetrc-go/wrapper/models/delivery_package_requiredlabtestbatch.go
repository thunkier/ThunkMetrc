package models

type DeliveryPackageRequiredlabtestbatch struct {
    LabTestBatchId int64 `json:"LabTestBatchId,omitempty"`
    LabTestBatchName string `json:"LabTestBatchName,omitempty"`
    PackageId int64 `json:"PackageId,omitempty"`
}
