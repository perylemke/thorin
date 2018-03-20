package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type payload struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

func main() {
	// Variables
	apiKey := os.Getenv("DO_API_KEY")
	dropletID := os.Getenv("DO_DROPLET_ID")

	dropletName, err := os.Hostname()
	if err != nil {
		fmt.Printf("error trying to get hostname: %v", err)
		os.Exit(1)
	}

	dateTime := time.Now().Format("02012006_150406")
	snapshotName := dropletName + "_" + dateTime

	apiCreateSnapshot := "https://api.digitalocean.com/v2/droplets/" + string(dropletID) + "/actions"

	data := payload{
		Type: "snapshot",
		Name: snapshotName,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error while marshal payload's json: %v", err)
		os.Exit(1)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", apiCreateSnapshot, body)
	if err != nil {
		fmt.Printf("error while preparing the request: %v", err)
		os.Exit(1)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("service response contains error %v", err)
		os.Exit(1)
	}

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("returned status is not %d: %d", http.StatusCreated, resp.StatusCode)
		os.Exit(1)
	}

	fmt.Printf("Success! - Status returned: %v\n", resp.Status)

}
