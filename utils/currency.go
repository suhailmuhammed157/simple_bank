package utils

const (
	USD = "USD"
	EUR = "EUR"
)

func IsValidCurrency(currency string) bool {
	switch currency {
	case USD, EUR:
		return true
	}
	return false
}
