package models

type ManifestPdf struct {
    ContentType string `json:"ContentType,omitempty"`
    FileContents string `json:"FileContents,omitempty"`
    FileDownloadName string `json:"FileDownloadName,omitempty"`
    HttpStatusCode string `json:"HttpStatusCode,omitempty"`
}
