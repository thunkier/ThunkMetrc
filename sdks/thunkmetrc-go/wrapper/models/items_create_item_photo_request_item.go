package models

type ItemsCreateItemPhotoRequestItem struct {
    EncodedImageBase64 string `json:"EncodedImageBase64,omitempty"`
    FileName string `json:"FileName,omitempty"`
}
