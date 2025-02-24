package helpers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"takehome/types"
	"time"
	"unicode"
)

func CountAlphaNumeric(retailer_name string) (int, error) {
	count := 0
	for _, r := range retailer_name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			count++
		}
	}
	return count, nil
}

func CalculatePointsForTotal(total string) (int, error) {
	totalPoints := 0
	if total == "" {
		return 0, fmt.Errorf("empty total value")
	}
	total_float, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting total to float: %v", err)
	}
	// Check if total is divisible by 0.25
	if math.Mod(total_float, 0.25) == 0 {
		totalPoints += 25
	}
	// Check if total is a whole number
	if (total_float - math.Floor(total_float)) == 0.0 {
		totalPoints += 50
	}
	return totalPoints, nil
}

func CalculatePointsForItems(items []types.Item) (int, error) {
	if len(items) == 0 {
		return 0, fmt.Errorf("Items List cannot be empty")
	}

	itemPoints := 0
	itemLen := len(items)
	itemPoints += (5 * (itemLen / 2))

	for _, item := range items {
		trimmedItem := strings.TrimSpace(item.ShortDescription)
		trimmedItemLen := len(trimmedItem)
		if trimmedItemLen == 0 {
			return 0, fmt.Errorf("Item Description cannot be empty")
		}
		if trimmedItemLen%3 == 0 {
			if item.Price == "" {
				return 0, fmt.Errorf("Price cannot be empty")
			}
			priceFloat, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, fmt.Errorf("Error converting price to float: %v", err)
			}
			itemPoints += int(math.Ceil(priceFloat * 0.2))
		}
	}
	return itemPoints, nil
}

func CalculatePointsForDate(date string) (int, error) {
	if date == "" {
		return 0, fmt.Errorf("Date cannot be empty")
	}

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, fmt.Errorf("Error parsing date: %v", err)
	}

	if parsedDate.Day()%2 != 0 {
		return 6, nil
	}
	return 0, nil
}

func CalculatePointsForTime(timeStr string) (int, error) {
	if timeStr == "" {
		return 0, fmt.Errorf("Time cannot be empty")
	}

	parseTime, err := time.Parse("05:04", timeStr)
	if err != nil {
		return 0, fmt.Errorf("Error parsing time: %v", err)
	}

	purchaseHour := parseTime.Hour()
	if purchaseHour >= 14 && purchaseHour < 16 {
		return 10, nil
	}
	return 0, nil
}
