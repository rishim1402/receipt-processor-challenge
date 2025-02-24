package tests

import (
	"receiptPointProcessor/helpers"
	"receiptPointProcessor/types"
	"testing"
)

func TestCountAlphaNumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Basic alphanumeric", "Target", 6},
		{"With spaces", "M&M Corner Market", 14},
		{"With special chars", "7-11!@#$%", 3},
		{"Empty string", "", 0},
		{"Only special chars", "@#$%^&", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count, err := helpers.CountAlphaNumeric(tt.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if count != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, count)
			}
		})
	}
}

func TestCalculatePointsForTotal(t *testing.T) {
	tests := []struct {
		name     string
		total    string
		expected int
		hasError bool
	}{
		{"Whole number", "100.00", 75, false}, // 50 for whole number + 25 for divisible by 0.25
		{"Quarter divisible", "35.25", 25, false},
		{"Neither whole nor quarter", "35.78", 0, false},
		{"Empty string", "", 0, true},
		{"Invalid number", "abc", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points, err := helpers.CalculatePointsForTotal(tt.total)
			if tt.hasError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.hasError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}

func TestCalculatePointsForItems(t *testing.T) {
	tests := []struct {
		name     string
		items    []types.Item
		expected int
		hasError bool
	}{
		{
			"Valid items",
			[]types.Item{
				{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
				{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			},
			8,
			false, // 5 points for 2 items
		},
		{
			"Item with length multiple of 3",
			[]types.Item{
				{ShortDescription: "ABC", Price: "10.00"},
			},
			2,
			false, // Ceil(10.00 * 0.2)
		},
		{
			"Empty items list",
			[]types.Item{},
			0,
			true,
		},
		{
			"Empty description",
			[]types.Item{{ShortDescription: "", Price: "10.00"}},
			0,
			true,
		},
		{
			"Empty price",
			[]types.Item{{ShortDescription: "ABC", Price: ""}},
			0,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points, err := helpers.CalculatePointsForItems(tt.items)
			if tt.hasError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.hasError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}

func TestCalculatePointsForDate(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		expected int
		hasError bool
	}{
		{"Odd day", "2023-09-15", 6, false},
		{"Even day", "2023-09-16", 0, false},
		{"Empty date", "", 0, true},
		{"Invalid date", "2023-13-45", 0, true},
		{"Wrong format", "15-09-2023", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points, err := helpers.CalculatePointsForDate(tt.date)
			if tt.hasError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.hasError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}

func TestCalculatePointsForTime(t *testing.T) {
	tests := []struct {
		name     string
		time     string
		expected int
		hasError bool
	}{
		{"Within range (2:30 PM)", "14:30", 10, false},
		{"Within range (3:30 PM)", "15:30", 10, false},
		{"Outside range (1:30 PM)", "13:30", 0, false},
		{"Outside range (4:30 PM)", "16:30", 0, false},
		{"Empty time", "", 0, true},
		{"Invalid time", "25:70", 0, true},
		{"Wrong format", "2:30 PM", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			points, err := helpers.CalculatePointsForTime(tt.time)
			if tt.hasError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.hasError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if points != tt.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}
