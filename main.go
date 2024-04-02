package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

var submittedBettors []string // Define a global variable to store submitted bettors
var finalBettors [][]string

func main() {
	r := gin.Default()
	// Enable CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		// Make the GET request to the API endpoint

		resp, err := http.Get("https://ncaa-api.henrygd.me/scoreboard/basketball-women/d1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch data from API",
			})
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read response body",
			})
			return
		}

		// Return the response body as JSON
		c.Data(http.StatusOK, "application/json", body)
	})

	r.POST("/set-bettors", func(c *gin.Context) {
		var bettorsRequest struct {
			Bettors []string `json:"bettors"`
		}

		// Bind the JSON request body to the bettorsRequest struct
		if err := c.BindJSON(&bettorsRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			fmt.Println(err)
			return
		}

		// Access the bettors array from bettorsRequest
		bettors := bettorsRequest.Bettors

		// Store the submitted bettors in the global variable
		submittedBettors = bettors

		// Desired number of rows and columns for the 2D array
		rows := 10
		cols := 10

		distributedPlayers := fillArrayWithPlayers(submittedBettors, rows, cols)

		finalBettors = distributedPlayers

		//return the distributed players and player counts as JSON
		c.JSON(http.StatusOK, gin.H{
			"distributed_players": distributedPlayers,
		})
	})

	r.POST("/refresh", func(c *gin.Context) {
		// Get the gameID from the request body
		var requestData map[string]string
		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		gameID, ok := requestData["gameID"]
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "gameID not provided"})
			return
		}

		// Fetch the winning square based on the gameID
		homeDigit, awayDigit, err := updateWinningSquare(gameID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch winning data"})
			return
		}

		// Return the winning data as JSON response
		responseData := map[string]int{
			"home_score_digit": homeDigit,
			"away_score_digit": awayDigit,
		}
		c.JSON(http.StatusOK, responseData)
	})

	r.GET("/bettors", func(c *gin.Context) {

		//return the submitted bettors as JSON
		c.JSON(http.StatusOK, gin.H{"bettors": finalBettors})
	})

	r.Run(":8080")
}
