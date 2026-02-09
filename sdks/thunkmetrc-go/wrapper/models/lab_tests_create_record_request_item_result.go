package models

type LabTestsCreateRecordRequestItemResult struct {
    LabTestTypeName string `json:"LabTestTypeName,omitempty"`
    Notes string `json:"Notes,omitempty"`
    Passed bool `json:"Passed,omitempty"`
    Quantity float64 `json:"Quantity,omitempty"`
}
