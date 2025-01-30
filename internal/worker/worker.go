package worker

import (
	"encoding/json"
	"fmt"
	"github.com/200sh/200sh-dashboard/models"
	"log"
	"net/http"
	"time"
)

func RunWorker(token, dashboardURL string) error {
	// 1) Fetch monitors from the dashboard
	monitors, err := fetchMonitors(dashboardURL, token)
	if err != nil {
		return fmt.Errorf("failed to fetch monitors: %w", err)
	}

	// 2) For each monitor, check uptime
	for _, m := range monitors {
		go checkMonitor(m) // or do it sequentially, depending on your design
	}

	// 3) Possibly sleep or run continuously
	time.Sleep(time.Minute) // example: wait 1 minute
	return nil              // or loop
}

func fetchMonitors(baseURL, token string) ([]models.Monitor, error) {
	req, err := http.NewRequest("GET", baseURL+"/api/monitors", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	var monitors []models.Monitor
	if err := json.NewDecoder(resp.Body).Decode(&monitors); err != nil {
		return nil, err
	}
	return monitors, nil
}

func checkMonitor(m models.Monitor) {
	// Perform a simple HTTP GET/HEAD
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(m.Url)
	if err != nil {
		log.Printf("Monitor %d (%s) is DOWN: %v", m.Id, m.Url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("Monitor %d (%s) is UP", m.Id, m.Url)
	} else {
		log.Printf("Monitor %d (%s) is DOWN (status %d)", m.Id, m.Url, resp.StatusCode)
	}
}
