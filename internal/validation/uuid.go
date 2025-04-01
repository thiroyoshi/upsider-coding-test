package validation

import (
	"fmt"

	"github.com/google/uuid"
)

func ValidateUUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		return fmt.Errorf("invalid UUID format: %s", id)
	}
	return nil
}
