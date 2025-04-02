package validation_test

import (
	"testing"
	"time"

	"upsider-coding-test/internal/validation"
)

func TestValidateDateFormat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		dateStr string
		wantErr bool
	}{
		{"valid date", "20230101", false},
		{"invalid length", "202301", true},
		{"invalid date", "20230132", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validation.ValidateDateFormat(tt.dateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateDateFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateDateRange(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		startAtStr string
		endAtStr   string
		wantErr    bool
	}{
		{"valid range", "20230101", "20230102", false},
		{"invalid start date", "20230132", "20230102", true},
		{"invalid end date", "20230101", "20230132", true},
		{"start after end", "20230102", "20230101", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validation.ValidateDateRange(tt.startAtStr, tt.endAtStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateDateRange() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePastDate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		dateStr string
		wantErr bool
	}{
		{"valid future date", time.Now().AddDate(0, 0, 1).Format("20060102"), false},
		{"invalid date", "20230132", true},
		{"past date", time.Now().AddDate(0, 0, -1).Format("20060102"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := validation.ValidatePastDate(tt.dateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePastDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
