package main

import "os"

func main() {
	scheduler := &DefaultScheduler{
		Interval: 30,
	}
	ipProvider := &DefaultIpAddressProvider{
		"0.0.0.0",
	}
	ddns := &GoogleDDNSClient{
		BaseURL:  "https://domains.google.com",
		Username: os.Getenv("GOOGLE_DDNS_USERNAME"),
		Password: os.Getenv("GOOGLE_DDNS_PASSWORD"),
	}

	scheduler.Start(ipProvider, ddns, os.Getenv("GOOGLE_DDNS_DOMAIN_NAME"))

}
