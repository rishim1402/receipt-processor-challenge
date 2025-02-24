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
		c.IndentedJSON(http.StatusOK, gin.H{"points": point})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
}

func postReceipt(c *gin.Context) {
	var newReceipt types.Receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		// fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid receipt data: %v", err)})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pointMap[generatedID] = totalPoints
	c.JSON(http.StatusCreated, gin.H{"id": generatedID})
}

func setupRouter() *gin.Engine {
	// fmt.Println("Setting up router")
	router := gin.Default()
	router.GET("/receipt/:id/points", getReceipt)
	router.POST("/receipt/process", postReceipt)
	return router
}

func main() {
	router := setupRouter()
	router.Run("localhost:8080")
}
