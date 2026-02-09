package models

type Receive struct {
    ChildTag string `json:"ChildTag,omitempty"`
    Eaches []string `json:"Eaches,omitempty"`
    LabelSource string `json:"LabelSource,omitempty"`
    QrCount int64 `json:"QrCount,omitempty"`
    Ranges [][]int64 `json:"Ranges,omitempty"`
    RequiresVerification bool `json:"RequiresVerification,omitempty"`
    SiblingTags []string `json:"SiblingTags,omitempty"`
}
