package requests

type Header struct {
	MaterialDocumentYear       string  `json:"MaterialDocumentYear"`
	MaterialDocument           string  `json:"MaterialDocument"`
	InventoryTransactionType   *string `json:"InventoryTransactionType"`
	DocumentDate               *string `json:"DocumentDate"`
	PostingDate                *string `json:"PostingDate"`
	CreationDate               *string `json:"CreationDate"`
	CreationTime               *string `json:"CreationTime"`
	MaterialDocumentHeaderText *string `json:"MaterialDocumentHeaderText"`
	ReferenceDocument          *string `json:"ReferenceDocument"`
	GoodsMovementCode          *string `json:"GoodsMovementCode"`
//	ToItem                     *struct {
//		Deferred struct {
//			URI string `json:"uri"`
//		} `json:"__deferred"`
//	} `json:"to_MaterialDocumentItem"`
}
