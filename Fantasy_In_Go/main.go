package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Roster struct {
	Points         float64            `json:"points"`
	Players        []string           `json:"players"`
	RosterID       int                `json:"roster_id"`
	CustomPoints   *float64           `json:"custom_points"` // nullable
	MatchupID      int                `json:"matchup_id"`
	Starters       []string           `json:"starters"`
	StartersPoints []float64          `json:"starters_points"`
	PlayersPoints  map[string]float64 `json:"players_points"`
}

type Player struct {
	PlayerID         string   `json:"player_id"` // accept "6462" as string
	FullName         string   `json:"full_name"`
	Active           bool     `json:"active"` // pointer because it may be null
	Status           string   `json:"status"`
	FantasyPositions []string `json:"fantasy_positions"`
	Position         string   `json:"position"`
	BirthDate        string   `json:"birth_date"`
	Number           int      `json:"number"`
	Age              int      `json:"age"`
}

func fetchHTML(url string) (string, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body), err

}

func main() {

	//var rosters []Roster
	rawJson, err := os.ReadFile("/home/jwhowell/Code_Projects/FantasyFootball/players.json")
	if err != nil {
		log.Fatalf("Error reading json file.")
	}

	var raw map[string]Player
	if err := json.Unmarshal(rawJson, &raw); err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}
	fmt.Println(raw)
	/*
	   data, err := fetchHTML("https://api.sleeper.app/v1/league/1258515704311709696/matchups/1")

	   	if err != nil {
	   		log.Printf("Error at fetchHTML")
	   	}

	   err = json.Unmarshal([]byte(data), &rosters)

	   	if err != nil {
	   		log.Print(err)
	   	}

	   fmt.Println(rosters[0])

	   db, err := sql.Open("sqlite", "football.db?_foreign_keys=1")

	   	if err != nil {
	   		log.Fatal(err)
	   	}

	   defer db.Close()

	   ctx := context.Background()

	   _, err = db.ExecContext(ctx, `

	   	CREATE TABLE IF NOT EXISTS football (
	   		id INTEGER PRIMARY KEY AUTOINCREMENT,
	   		title TEXT NOT NULL,
	   		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	   	);

	   `)

	   	if err != nil {
	   		log.Fatal(err)
	   	}
	*/
}
