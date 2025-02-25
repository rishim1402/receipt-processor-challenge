package helpers

import (
	"fmt"
	"math"
	"receiptPointProcessor/constants"
	"receiptPointProcessor/types"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// CountAlphaNumeric counts and returns the number of alphanumeric characters in a given string.
// It takes a retailer name as input.
func CountAlphaNumeric(retailer_name string) (int, error) {
	count := 0
	for _, r := range retailer_name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			count += constants.Retailer_name_point
		}
	}
	return count, nil
}

// CalculatePointsForTotal calculates points based on the total value string.
// It returns points if the total is divisible by 0.25 and if it is a whole number.
func CalculatePointsForTotal(total string) (int, error) {
	totalPoints := 0
	if total == "" {
		return 0, fmt.Errorf("empty total value")
	}
	total_float, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting total to float: %v", err)
	}
	if total_float < 0 {
		return 0, fmt.Errorf("Total cannot be negative")
	}
	if math.Mod(total_float, 0.25) == 0 {
		totalPoints += constants.Total_divisible_by_0_25
	}
	if (total_float - math.Floor(total_float)) == 0.0 {
		totalPoints += constants.Total_whole_number
	}
	return totalPoints, nil
}

// CalculatePointsForItems calculates points for a list of items.
// Points are assigned based on the item price and description length conditions.
func CalculatePointsForItems(items []types.Item) (int, error) {
	if len(items) == 0 {
		return 0, fmt.Errorf("Items List cannot be empty")
	}

	itemPoints := 0
	itemLen := len(items)
	itemPoints += (constants.Item_length_point * (itemLen / 2))

	for _, item := range items {
		priceFloat, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			return 0, fmt.Errorf("Error converting price to float: %v", err)
		}
		if priceFloat < 0 {
			return 0, fmt.Errorf("Price cannot be negative")
		}
		trimmedItem := strings.TrimSpace(item.ShortDescription)
		trimmedItemLen := len(trimmedItem)
		if trimmedItemLen == 0 {
			return 0, fmt.Errorf("Item Description cannot be empty")
		}
		if trimmedItemLen%constants.Item_description_lenght_multiple == 0 {
			if item.Price == "" {
				return 0, fmt.Errorf("Price cannot be empty")
			}
			itemPoints += int(math.Ceil(priceFloat * constants.Item_description_lenght_based_points))
		}
	}
	return itemPoints, nil
}

// CalculatePointsForDate calculates points based on the purchase date.
// Returns 6 points if the day of the month is odd.
func CalculatePointsForDate(date string) (int, error) {
	if date == "" {
		return 0, fmt.Errorf("Date cannot be empty")
	}

	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, fmt.Errorf("Error parsing date: %v", err)
	}

	if parsedDate.Day()%2 != 0 {
		return constants.Odd_date_point, nil
	}
	return constants.Even_date_point, nil
}

// CalculatePointsForTime calculates points based on the purchase time.
// It awards 10 points if the time is between 14:00 and 16:00.
func CalculatePointsForTime(timeStr string) (int, error) {
	if timeStr == "" {
		return 0, fmt.Errorf("Time cannot be empty")
	}

	parseTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return 0, fmt.Errorf("Error parsing time: %v", err)
	}

	purchaseHour := parseTime.Hour()
	if purchaseHour >= constants.Start_time_deadline && purchaseHour < constants.End_time_deadline {
		return constants.Peak_hour_point, nil
	}
	return 0, nil
}
