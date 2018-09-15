package tax

import (
	validator "gopkg.in/go-playground/validator.v9"
)

const (
	tableItem = "public.items"
)

func validateNewItem(newItem item) (errorMessage []string, ok bool) {
	ok = true
	if err := validator.New().Struct(newItem); err != nil {
		ok = false

		for _, err := range err.(validator.ValidationErrors) {
			errorMessage = append(errorMessage, err.Field()+" - "+err.Tag())
		}
	}
	return
}

func calculateTax(newItem item) float64 {
	switch newItem.TaxCode {
	case 1:
		return float64(newItem.Amount) * 0.1
	case 2:
		return 10 + (float64(newItem.Amount) * 0.2)
	case 3:
		if newItem.Amount >= 0 && newItem.Amount <= 100 {
			return 0
		} else {
			return 0.1 * (float64(newItem.Amount) - 100)
		}
	default:
		return 0
	}
}

func setTaxType(taxCode int) string {
	switch taxCode {
	case 1:
		return "Food"
	case 2:
		return "Tobacco"
	case 3:
		return "Entertainment"
	default:
		return ""
	}
}

func calculateTotal(items []item) (taxAmount, totalAmount, grandTotal float64) {
	for _, item := range items {
		taxAmount += item.TaxAmount
		totalAmount += float64(item.Amount)
	}
	grandTotal = taxAmount + totalAmount
	return
}
