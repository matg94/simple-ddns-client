package main

import (
	"log"
	"os"
	"strconv"
)

func StartGoogleClient() {
	ipProvider := &DefaultIpAddressProvider{
		"0.0.0.0",
	}

	interval, err := strconv.Atoi(os.Getenv("GOOGLE_DDNS_INTERVAL"))
	if err != nil {
		interval = 60
	}

	scheduler := &DefaultScheduler{
		Interval: interval,
	}

	user := os.Getenv("GOOGLE_DDNS_USERNAME")
	pass := os.Getenv("GOOGLE_DDNS_PASSWORD")
	domain := os.Getenv("GOOGLE_DDNS_DOMAIN_NAME")

	if user == "" || pass == "" || domain == "" {
		log.Fatalf("Could not read one or more environment variables")
	}

	ddns := &GoogleDDNSClient{
		BaseURL:  "https://domains.google.com",
		Username: user,
		Password: pass,
	}

	scheduler.Start(ipProvider, ddns, domain)
}

func main() {
	StartGoogleClient()
}
