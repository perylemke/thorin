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
		panic(err)
	}

	now := time.Now()
	dateTime := now.Format("02012006_150406")
	snapshotName := dropletName + "_" + dateTime

	apiCreateSnapshot := "https://api.digitalocean.com/v2/droplets/" + string(dropletID) + "/actions"

	data := payload{
		Type: "snapshot",
		Name: snapshotName,
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", apiCreateSnapshot, body)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Error - Status returned: %v\n", resp.Status)
	} else {
		fmt.Printf("Success - Status returned: %v\n", resp.Status)
	}
}
