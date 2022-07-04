package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// record represents data about a record.
type record struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Artist string    `json:"artist"`
	Price  float64   `json:"price"`
}

// make some dummy data to work with
var records = []record{
	{ID: uuid.New(), Title: "Volume", Artist: "Jauz", Price: 19.99},
	{ID: uuid.New(), Title: "DAMN", Artist: "Joyryde", Price: 17.99},
	{ID: uuid.New(), Title: "Indian Summer", Artist: "Jai Wolf", Price: 16.49},
}

func main() {
	fmt.Println("\n\nStarting the application. Initial record IDs:")
	for _, record := range records {
		fmt.Println("    ", record.ID)
	}

	router := gin.Default()
	router.GET("/ping", getPing)
	router.GET("/records", getRecords)
	router.GET("/records/:id", getRecordByID)
	router.POST("/records", postRecords)

	router.Run("0.0.0.0:8080")
}

// getPing sends "pong," so quirky! hehe
func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// getRecords responds with the list of all records as JSON.
func getRecords(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, records)
}

// getRecordByID returns a record with the given ID, if there is one
func getRecordByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range records {
		if a.ID.String() == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "record not found"})
}

// postRecords adds a record to the slice of records
func postRecords(c *gin.Context) {
	var newRecord record

	// unmarshal the JSON from the request body
	if err := c.BindJSON(&newRecord); err != nil {
		return
	}
	records = append(records, newRecord)
	c.IndentedJSON(http.StatusCreated, newRecord)
}
