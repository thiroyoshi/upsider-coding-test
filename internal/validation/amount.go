package validation

import (
	"errors"
)

func ValidateAmount(amount int) error {
	if amount < 0 {
		return errors.New("amount must be greater than or equal to 0")
	}
	return nil
}
