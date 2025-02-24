package processing

import (
	"fmt"
	"receiptPointProcessor/helpers"
	"receiptPointProcessor/types"
)

func ProcessReceipt(newRecpt types.Receipt) (int, error) {
	totalPoints := 0

	// Check if receipt data is empty
	if newRecpt.Retailer == "" && newRecpt.Total == "" && len(newRecpt.Items) == 0 && newRecpt.PurchaseDate == "" && newRecpt.PurchaseTime == "" {
		return 0, fmt.Errorf("Empty receipt data")
	}

	// Calculate retailer points
	if points, err := helpers.CountAlphaNumeric(newRecpt.Retailer); err != nil {
		return 0, err
	} else {
		totalPoints += points
	}

	// Calculate total points
	if points, err := helpers.CalculatePointsForTotal(newRecpt.Total); err != nil {
		return 0, err
	} else {
		totalPoints += points
	}

	// Calculate items points
	if points, err := helpers.CalculatePointsForItems(newRecpt.Items); err != nil {
		return 0, err
	} else {
		totalPoints += points
	}

	// Calculate date points
	if points, err := helpers.CalculatePointsForDate(newRecpt.PurchaseDate); err != nil {
		return 0, err
	} else {
		totalPoints += points
	}

	// Calculate time points
	if points, err := helpers.CalculatePointsForTime(newRecpt.PurchaseTime); err != nil {
		return 0, err
	} else {
		totalPoints += points
	}

	return totalPoints, nil
}
