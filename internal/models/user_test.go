package models

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "birthday passed this year",
			dob:      time.Date(1990, 5, 10, 0, 0, 0, 0, time.UTC),
			expected: time.Now().Year() - 1990,
		},
		{
			name:     "birthday not yet passed this year",
			dob:      time.Date(1990, 12, 31, 0, 0, 0, 0, time.UTC),
			expected: time.Now().Year() - 1990 - 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			age := CalculateAge(tt.dob)
			if age != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, age)
			}
		})
	}
}
