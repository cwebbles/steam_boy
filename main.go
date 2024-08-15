package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Game struct {
	Cmd   string `json:"cmd"`
	Name  string `json:"name"`
	AppId string `json:"id"`
}

// Add support for adding a new game to the list
// Add support for removing a game from the list
// Add support for listing all games in the list
// Add the steam app id search program to cli
func main() {
	// Open config file
	file, err := os.Open("games.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Decode JSON file into array of Game structs
	var game []Game
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&game)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Please provide a name to search for.")
	}
	searchCmd := os.Args[1]

	// Find provided game cmd in the list
	g, err := findGame(game, searchCmd)

	if err != nil {
		log.Fatalf("Closing program: %v", err)
	} else {
		// Run the game
		runGame(g)
	}
}

func findGame(game []Game, id string) (Game, error) {
	for _, g := range game {
		if g.Cmd == id {
			return g, nil
		}
	}
	return Game{}, fmt.Errorf("Game not found: %s", id)
}

// Try to use steams appid stuff to run game
func runGame(g Game) error {
	steamURI := fmt.Sprintf("steam://rungameid/%s", g.AppId)
	cmd := exec.Command("cmd", "/C", "start", steamURI)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run game: %v", err)
	}
	fmt.Printf("Running %s\n", g.Name)
	return nil
}
