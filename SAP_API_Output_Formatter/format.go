package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-material-document-creates/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	header := &Header{
		MaterialDocumentYear:       data.MaterialDocumentYear,
		MaterialDocument:           data.MaterialDocument,
		InventoryTransactionType:   data.InventoryTransactionType,
		DocumentDate:               data.DocumentDate,
		PostingDate:                data.PostingDate,
		CreationDate:               data.CreationDate,
		CreationTime:               data.CreationTime,
		MaterialDocumentHeaderText: data.MaterialDocumentHeaderText,
		ReferenceDocument:          data.ReferenceDocument,
		GoodsMovementCode:          data.GoodsMovementCode,
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D
	item := &Item{
	MaterialDocumentYear:           data.MaterialDocumentYear,
	MaterialDocument:               data.MaterialDocument,
	MaterialDocumentItem:           data.MaterialDocumentItem,
	Material:                       data.Material,
	Plant:                          data.Plant,
	StorageLocation:                data.StorageLocation,
	Batch:                          data.Batch,
	GoodsMovementType:              data.GoodsMovementType,
	InventoryStockType:             data.InventoryStockType,
	InventoryValuationType:         data.InventoryValuationType,
	InventorySpecialStockType:      data.InventorySpecialStockType,
	Supplier:                       data.Supplier,
	Customer:                       data.Customer,
	SalesOrder:                     data.SalesOrder,
	SalesOrderItem:                 data.SalesOrderItem,
	SalesOrderScheduleLine:         data.SalesOrderScheduleLine,
	PurchaseOrder:                  data.PurchaseOrder,
	PurchaseOrderItem:              data.PurchaseOrderItem,
	WBSElement:                     data.WBSElement,
	ManufacturingOrder:             data.ManufacturingOrder,
	ManufacturingOrderItem:         data.ManufacturingOrderItem,
	GoodsMovementRefDocType:        data.GoodsMovementRefDocType,
	GoodsMovementReasonCode:        data.GoodsMovementReasonCode,
	Delivery:                       data.Delivery,
	DeliveryItem:                   data.DeliveryItem,
	AccountAssignmentCategory:      data.AccountAssignmentCategory,
	CostCenter:                     data.CostCenter,
	ControllingArea:                data.ControllingArea,
	CostObject:                     data.CostObject,
	ProfitabilitySegment:           data.ProfitabilitySegment,
	ProfitCenter:                   data.ProfitCenter,
	GLAccount:                      data.GLAccount,
	FunctionalArea:                 data.FunctionalArea,
	MaterialBaseUnit:               data.MaterialBaseUnit,
	QuantityInBaseUnit:             data.QuantityInBaseUnit,
	EntryUnit:                      data.EntryUnit,
	QuantityInEntryUnit:            data.QuantityInEntryUnit,
	CompanyCodeCurrency:            data.CompanyCodeCurrency,
	FiscalYear:                     data.FiscalYear,
	FiscalYearPeriod:               data.FiscalYearPeriod,
	IssgOrRcvgMaterial:             data.IssgOrRcvgMaterial,
	IssgOrRcvgBatch:                data.IssgOrRcvgBatch,
	IssuingOrReceivingPlant:        data.IssuingOrReceivingPlant,
	IssuingOrReceivingStorageLoc:   data.IssuingOrReceivingStorageLoc,
	IssuingOrReceivingStockType:    data.IssuingOrReceivingStockType,
	IssgOrRcvgSpclStockInd:         data.IssgOrRcvgSpclStockInd,
	IssuingOrReceivingValType:      data.IssuingOrReceivingValType,
	IsCompletelyDelivered:          data.IsCompletelyDelivered,
	MaterialDocumentItemText:       data.MaterialDocumentItemText,
	UnloadingPointName:             data.UnloadingPointName,
	ShelfLifeExpirationDate:        data.ShelfLifeExpirationDate,
	ManufactureDate:                data.ManufactureDate,
	SerialNumbersAreCreatedAutomly: data.SerialNumbersAreCreatedAutomly,
	Reservation:                    data.Reservation,
	ReservationItem:                data.ReservationItem,
	ReservationIsFinallyIssued:     data.ReservationIsFinallyIssued,
	IsAutomaticallyCreated:         data.IsAutomaticallyCreated,
	GoodsMovementIsCancelled:       data.GoodsMovementIsCancelled,
	ReversedMaterialDocumentYear:   data.ReversedMaterialDocumentYear,
	ReversedMaterialDocument:       data.ReversedMaterialDocument,
	ReversedMaterialDocumentItem:   data.ReversedMaterialDocumentItem,
	ReferenceDocumentFiscalYear:    data.ReferenceDocumentFiscalYear,
	InvtryMgmtRefDocumentItem:      data.InvtryMgmtRefDocumentItem,
	InvtryMgmtReferenceDocument:    data.InvtryMgmtReferenceDocument,
	MaterialDocumentPostingType:    data.MaterialDocumentPostingType,
	}

	return item, nil
}
