package sap_api_input_reader

import (
	"sap-api-integrations-material-document-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.MaterialDocument
	return &requests.Header{
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
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataMaterialDocument := sdc.MaterialDocument
	data := sdc.MaterialDocument.MaterialDocumentItem
	return &requests.Item{
		MaterialDocumentYear:           dataMaterialDocument.MaterialDocumentYear,
		MaterialDocument:               dataMaterialDocument.MaterialDocument,
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
}
