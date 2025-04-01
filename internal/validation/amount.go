package validation

import "fmt"

func ValidateAmount(amount int) error {
	if amount < 0 {
		return fmt.Errorf("amount must be greater than or equal to 0")
	}
	return nil
}
