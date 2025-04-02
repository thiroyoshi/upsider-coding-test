package validation_test

import (
	"testing"

	"upsider-coding-test/internal/validation"
)

func TestValidateAmount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		amount  int
		wantErr bool
	}{
		{
			name:    "valid amount",
			amount:  10,
			wantErr: false,
		},
		{
			name:    "invalid amount",
			amount:  -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validation.ValidateAmount(tt.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
