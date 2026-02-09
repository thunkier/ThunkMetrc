package models

type ItemsCreateFileRequestItem struct {
    EncodedImageBase64 string `json:"EncodedImageBase64,omitempty"`
    FileName string `json:"FileName,omitempty"`
}
