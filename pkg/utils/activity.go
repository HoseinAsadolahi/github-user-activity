package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func fetchData(username string, page int) (map[string]any, error) {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events?page=%d", username, page))
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		if response.StatusCode == 404 {
			return nil, fmt.Errorf("username not found")
		} else {
			return nil, fmt.Errorf("error fetching data. status code: %d", response.StatusCode)
		}
	}
	var result map[string]any
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return result, nil
}
