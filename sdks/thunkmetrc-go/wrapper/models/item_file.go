package models

type ItemFile struct {
    ContentType string `json:"ContentType,omitempty"`
    FileContents string `json:"FileContents,omitempty"`
    FileDownloadName string `json:"FileDownloadName,omitempty"`
}
