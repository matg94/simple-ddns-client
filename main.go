package main

import (
	"log"
	"os"
)

func StartGoogleClient() {
	scheduler := &DefaultScheduler{
		Interval: 30,
	}
	ipProvider := &DefaultIpAddressProvider{
		"0.0.0.0",
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
