package main

import (
	"fmt"
	"net/http"
	"receiptPointProcessor/helpers"
	"receiptPointProcessor/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var recpts []types.Receipt
var pointMap = make(map[string]int)

func getReceipt(c *gin.Context) {
	id := c.Param("id")
	if point, ok := pointMap[id]; ok {
		c.IndentedJSON(http.StatusOK, gin.H{"points": point})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
}

func postReceipt(c *gin.Context) {
	var newRecpt types.Receipt

	if err := c.BindJSON(&newRecpt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid receipt data"})
		return
	}

	generatedID := uuid.New().String()
	for {
		if _, ok := pointMap[generatedID]; !ok {
			break
		}
		generatedID = uuid.New().String()
	}

	totalPoints := 0

	// Check if receipt data is empty
	if newRecpt.Retailer == "" && newRecpt.Total == "" && len(newRecpt.Items) == 0 && newRecpt.PurchaseDate == "" && newRecpt.PurchaseTime == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty receipt data"})
		return
	}

	// Calculate retailer points
	if points, err := helpers.CountAlphaNumeric(newRecpt.Retailer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		totalPoints += points
	}

	// Calculate total points
	if points, err := helpers.CalculatePointsForTotal(newRecpt.Total); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		totalPoints += points
	}

	// Calculate items points
	if points, err := helpers.CalculatePointsForItems(newRecpt.Items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		totalPoints += points
	}

	// Calculate date points
	if points, err := helpers.CalculatePointsForDate(newRecpt.PurchaseDate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		totalPoints += points
	}

	// Calculate time points
	if points, err := helpers.CalculatePointsForTime(newRecpt.PurchaseTime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		totalPoints += points
	}

	pointMap[generatedID] = totalPoints
	c.JSON(http.StatusCreated, gin.H{"id": generatedID})
}

func setupRouter() *gin.Engine {
	fmt.Println("Setting up router")
	router := gin.Default()
	router.GET("/receipt/:id/points", getReceipt)
	router.POST("/receipt/process", postReceipt)
	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
	// router := gin.Default()
	// router.GET("/receipt/:id/points", getRecpt)
	// router.POST("/receipt/process", postReceipt)
	// router.Run("localhost:8080")
}
