package main

import (
	"fmt"
	"net/http"
	"receiptPointProcessor/processing"
	"receiptPointProcessor/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receipts []types.Receipt
var pointMap = make(map[string]int)

func getReceipt(c *gin.Context) {
	id := c.Param("id")
	if point, ok := pointMap[id]; ok {
		c.IndentedJSON(http.StatusOK, gin.H{"Points": point})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No receipt found for that ID."})
}

func postReceipt(c *gin.Context) {
	var newReceipt types.Receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("The receipt is invalid.")})
		return
	}

	generatedID := uuid.New().String()
	for {
		if _, ok := pointMap[generatedID]; !ok {
			break
		}
		generatedID = uuid.New().String()
	}

	totalPoints, err := processing.ProcessReceipt(newReceipt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("The receipt is invalid.")})
		return
	}
	pointMap[generatedID] = totalPoints
	c.JSON(http.StatusCreated, gin.H{"id": generatedID})
}

func setupRouter() *gin.Engine {
	// fmt.Println("Setting up router")
	router := gin.Default()
	router.GET("/receipts/:id/points", getReceipt)
	router.POST("/receipts/process", postReceipt)
	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}
