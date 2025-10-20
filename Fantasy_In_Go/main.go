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
	LeagueID string   `json:"league_id"`
	Metadata Metadata `json:"metadata"`
	OwnerID  string   `json:"owner_id"`
	Players  []string `json:"players"`
	Reserve  any      `json:"reserve"`
	RosterID int      `json:"roster_id"`
	Settings Settings `json:"settings"`
	Starters []string `json:"starters"`
	Taxi     any      `json:"taxi"`
}

type Metadata struct {
	AllowPNInactiveStarters       string `json:"allow_pn_inactive_starters"`
	AllowPNPlayerInjuryStatus     string `json:"allow_pn_player_injury_status"`
	AllowPNScoring                string `json:"allow_pn_scoring"`
	Record                        string `json:"record"`
	RestrictPNScoringStartersOnly string `json:"restrict_pn_scoring_starters_only"`
	Streak                        string `json:"streak"`
}

type Settings struct {
	Fpts               int `json:"fpts"`
	FptsAgainst        int `json:"fpts_against"`
	FptsAgainstDecimal int `json:"fpts_against_decimal"`
	FptsDecimal        int `json:"fpts_decimal"`
	Losses             int `json:"losses"`
	Ppts               int `json:"ppts"`
	PptsDecimal        int `json:"ppts_decimal"`
	Ties               int `json:"ties"`
	TotalMoves         int `json:"total_moves"`
	WaiverBudgetUsed   int `json:"waiver_budget_used"`
	WaiverPosition     int `json:"waiver_position"`
	Wins               int `json:"wins"`
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

	rawJson, err := os.ReadFile("/home/jwhowell/Code_Projects/FantasyFootball/players.json")
	if err != nil {
		log.Fatalf("Error reading json file.")
	}

	var raw map[string]Player
	if err := json.Unmarshal(rawJson, &raw); err != nil {
		log.Fatalf("unmarshal error: %v", err)
	}
	fmt.Println(raw["6462"])

	var roster map[string]Roster
	jsonResp, err := fetchHTML("https://api.sleeper.app/v1/league/1258515704311709696/rosters")
	if err != nil {
		log.Fatalf("Error fetching Json: %v", err)
	}
	if err := json.Unmarshal([]byte(jsonResp), &roster); err != nil {
		log.Fatalf("Error unmarshalling Json: % v", err)
	}

	fmt.Println(roster)

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
