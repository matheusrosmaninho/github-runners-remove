package helpers

import (
	"testing"
	"time"
)

func TestSubtractDate(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2022-10-31")
	if err != nil {
		t.Errorf("Error parsing initial date: %+v", err)
	}

	finalDate, err := time.Parse("2006-01-02", "2022-10-30")
	if err != nil {
		t.Errorf("Error parsing final date: %+v", err)
	}
	result := SubtractDaysInDate(date, 1)
	if result != finalDate {
		t.Errorf("Error subtracting days in date: %+v", result)
	}
}
