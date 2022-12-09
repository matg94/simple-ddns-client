package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func CreateClient(key, clientType string) (DDNSClient, error) {
	username := os.Getenv(fmt.Sprintf("%s_DDNS_USERNAME_%s", clientType, key))
	password := os.Getenv(fmt.Sprintf("%s_DDNS_PASSWORD_%s", clientType, key))
	domain := os.Getenv(fmt.Sprintf("%s_DDNS_DOMAIN_NAME_%s", clientType, key))
	if username == "" || password == "" || domain == "" {
		log.Printf("Could not load client of type %s and key %s because of missing env variables", clientType, key)
		return &GoogleDDNSClient{}, errors.New("failed to load env variables")
	}

	interval, err := strconv.Atoi(os.Getenv(fmt.Sprintf("%s_DDNS_INTERVAL_%s", clientType, key)))
	if err != nil {
		interval = 60
	}

	if clientType == "GOOGLE" {
		return &GoogleDDNSClient{
			BaseURL:    "https://domains.google.com",
			Username:   username,
			Password:   password,
			DomainName: domain,
			Interval:   interval,
		}, nil
	} else {
		return &GoogleDDNSClient{}, errors.New("client type not yet supported")
	}
}

func StartClient(scheduler Scheduler, ddnsClient DDNSClient, ipProvider IpAddressProvider) {
	scheduler.Start(ipProvider, ddnsClient)
}

func main() {
	ipProvider := &DefaultIpAddressProvider{
		"0.0.0.0",
	}
	scheduler := &DefaultScheduler{}
	valheimDDNS, err := CreateClient("VALHEIM", "GOOGLE")
	if err != nil {
		log.Fatalf("Error creating Valheim client %s", err)
	}
	openVPNDDNS, err := CreateClient("OPENVPN", "GOOGLE")
	if err != nil {
		log.Fatalf("Error creating OpenVPN client %s", err)
	}
	go StartClient(scheduler, valheimDDNS, ipProvider)
	StartClient(scheduler, openVPNDDNS, ipProvider)
}
