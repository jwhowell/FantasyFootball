package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

type Players struct {
	PlayerID int `json:"player_id"`
	FullName string `json:"full_name"`
	Active bool `json:"active"`
	Status string `json:"status"`
	FantasyPositions []string `json:"fantasy_positions"`
	Position string `json:"position"`

}

func main() {
	var rosters []Roster

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
