package models

type LabTestsUpdateLabTestDocumentRequestItem struct {
    DocumentFileBase64 string `json:"DocumentFileBase64,omitempty"`
    DocumentFileName string `json:"DocumentFileName,omitempty"`
    LabTestResultId int `json:"LabTestResultId,omitempty"`
}
