package models

type SalesUpdateSaleDeliveriesCompleteRequestItemReturnedPackage struct {
    Label string `json:"Label,omitempty"`
    ReturnQuantityVerified int `json:"ReturnQuantityVerified,omitempty"`
    ReturnReason string `json:"ReturnReason,omitempty"`
    ReturnReasonNote string `json:"ReturnReasonNote,omitempty"`
    ReturnUnitOfMeasure string `json:"ReturnUnitOfMeasure,omitempty"`
}
