package main

import (
	"github.com/200sh/200sh-dashboard/internal/worker"
	"log"
	"os"
)

func main() {
	// 1) Read environment variables
	token := os.Getenv("UPTIME_API_TOKEN")
	dashboardURL := os.Getenv("DASHBOARD_URL")

	// 2) Start worker loop or single run
	if err := worker.RunWorker(token, dashboardURL); err != nil {
		log.Fatal(err)
	}
}
