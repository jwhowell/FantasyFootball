package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	html, err := fetchHTML("https://api.sleeper.app/v1/league/1258515704311709696/matchups/1")
	if err != nil {
		log.Printf("Error at fetchHTML")
	}
	fmt.Println(html)

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
