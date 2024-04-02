package main

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func shufflePlayers2D(players [][]string) {
	rand.Seed(time.Now().UnixNano())

	// Flatten the 2D slice into a 1D slice
	flattened := make([]string, 0)
	for _, row := range players {
		flattened = append(flattened, row...)
	}

	// Shuffle the 1D slice
	rand.Shuffle(len(flattened), func(i, j int) {
		flattened[i], flattened[j] = flattened[j], flattened[i]
	})

	// Reassign the shuffled players back to the original 2D slice
	index := 0
	for i := range players {
		for j := range players[i] {
			players[i][j] = flattened[index]
			index++
		}
	}
}

// funciton to distribute players across the spots in the game
func fillArrayWithPlayers(players []string, rows, cols int) [][]string {
	totalPlayers := len(players)
	distributedPlayers := make([][]string, rows)
	playerIndex := 0
	for i := 0; i < rows; i++ {
		distributedPlayers[i] = make([]string, cols)
		for j := 0; j < cols; j++ {
			playerIndex %= totalPlayers
			distributedPlayers[i][j] = players[playerIndex]
			playerIndex++
		}
	}
	shufflePlayers2D(distributedPlayers)
	return distributedPlayers
}
func updateWinningSquare(gameID string) (homeDigit int, awayDigit int, err error) {
	//URL for the api endpoint
	url := "https://ncaa-api.henrygd.me/scoreboard/basketball-women/d1"

	// Make the api call
	response, err := http.Get(url)
	if err != nil {
		log.Println("Failed to make API call:", err)
		return 0, 0, err
	}
	defer response.Body.Close()

	// Check the status code of the response
	if response.StatusCode != http.StatusOK {
		log.Println("Failed to retrieve data. Status code:", response.StatusCode)
		return 0, 0, errors.New("failed to retrieve data")
	}

	// Decode the JSON response
	var data map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return 0, 0, err
	}

	// Extract games from the response
	games, ok := data["games"].([]interface{})
	if !ok {
		return 0, 0, errors.New("games not found in response")
	}

	// Iterate over each game
	for _, game := range games {
		gameMap := game.(map[string]interface{})
		gameData := gameMap["game"].(map[string]interface{})
		gameIDValue, ok := gameData["gameID"].(string)
		if !ok {
			continue
		}
		if gameIDValue == gameID {
			// get the scores for each team
			homeScore := gameData["home"].(map[string]interface{})["score"].(string)
			awayScore := gameData["away"].(map[string]interface{})["score"].(string)

			// get the last digit in the scores and return them with no error
			homeDigit, _ := strconv.Atoi(string(homeScore[len(homeScore)-1]))
			awayDigit, _ := strconv.Atoi(string(awayScore[len(awayScore)-1]))

			return homeDigit, awayDigit, nil
		}
	}

	fmt.Println("game", gameID)

	return 0, 0, errors.New("gameID not found")
}
