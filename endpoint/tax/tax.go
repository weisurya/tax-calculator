package tax

import (
	"encoding/json"
	"net/http"

	common "tax-calculator/common"
	provider "tax-calculator/provider"
)

type item struct {
	Name        string  `json:"name"validate:"required"`
	TaxCode     int     `json:"taxcode"validate:"required"`
	Amount      int     `json:"amount"validate:"required"`
	Type        string  `json:"type"`
	TaxAmount   float64 `json:"taxamount"`
	TotalAmount float64 `json:"totalamount"`
}

func GetTotal(w http.ResponseWriter, r *http.Request) {
	http := common.HTTP{
		Write: w,
		Read:  r,
	}

	// Retrieve all records from database
	items, err := retrieveAllItems(provider.Provider.Database)
	if err != nil {
		common.CreateResponse(http, common.ErrorUnprocessableEntityCode, err.Error())
		return
	}
	if len(items) == 0 {
		common.CreateResponse(http, common.ErrorNotFound, nil)
		return
	}

	// Calculate grand total
	taxAmount, totalAmount, grandTotal := calculateTotal(items)

	output := make(map[string]interface{})
	output["items"] = items
	output["total_amount"] = totalAmount
	output["tax_amount"] = taxAmount
	output["grand_total"] = grandTotal

	// Create success response
	common.CreateResponse(http, common.NoErrorCode, output)
	return
}

func PostItem(w http.ResponseWriter, r *http.Request) {
	http := common.HTTP{
		Write: w,
		Read:  r,
	}

	body, err := common.ReadBody(r)
	if err != nil {
		common.CreateResponse(http, common.ErrorUnprocessableEntityCode, err.Error())
		return
	}

	newItem := item{}
	if err := json.Unmarshal(body, &newItem); err != nil {
		common.CreateResponse(http, common.ErrorUnprocessableEntityCode, err.Error())
		return
	}

	errorMessage, passed := validateNewItem(newItem)
	if passed != true {
		common.CreateResponse(http, common.ErrorBadRequestCode, errorMessage)
		return
	}

	newItem.TaxAmount = calculateTax(newItem)
	newItem.Type = setTaxType(newItem.TaxCode)
	newItem.TotalAmount = float64(newItem.Amount) + newItem.TaxAmount

	if err := insertNewItem(provider.Provider.Database, newItem); err != nil {
		common.CreateResponse(http, common.ErrorUnprocessableEntityCode, err.Error())
		return
	}

	common.CreateResponse(http, common.NoErrorCode, "OK")
	return
}
