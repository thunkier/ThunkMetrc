package models

type LabTestsCreateRecordRequestItem struct {
    DocumentFileBase64 string `json:"DocumentFileBase64,omitempty"`
    DocumentFileName string `json:"DocumentFileName,omitempty"`
    Label string `json:"Label,omitempty"`
    ResultDate string `json:"ResultDate,omitempty"`
    Results []*LabTestsCreateRecordRequestItemResult `json:"Results,omitempty"`
}
