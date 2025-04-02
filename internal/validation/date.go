package validation

import (
	"errors"
	"fmt"
	"time"
)

const EIGHT = 8

func ValidateDateFormat(dateStr string) error {
	if len(dateStr) != EIGHT {
		return errors.New("invalid date format: must be YYYYMMDD")
	}

	year := dateStr[0:4]
	month := dateStr[4:6]
	day := dateStr[6:8]

	_, err := time.Parse("20060102", fmt.Sprintf("%s%s%s", year, month, day))
	if err != nil {
		return fmt.Errorf("invalid date: %s", dateStr)
	}

	return nil
}

func ValidateDateRange(startAtStr, endAtStr string) error {
	start, err := time.Parse("20060102", startAtStr)
	if err != nil {
		return fmt.Errorf("invalid start date: %s", startAtStr)
	}

	end, err := time.Parse("20060102", endAtStr)
	if err != nil {
		return fmt.Errorf("invalid end date: %s", endAtStr)
	}

	if start.After(end) {
		return errors.New("startAt must be before or equal to endAt")
	}
	return nil
}

func ValidatePastDate(dateStr string) error {
	date, err := time.Parse("20060102", dateStr)
	if err != nil {
		return fmt.Errorf("invalid date: %s", dateStr)
	}

	if date.Before(time.Now()) {
		return errors.New("date must be in the future")
	}
	return nil
}
