package main

import (
	"fmt"
	"log"
	"net/http"
)

type DDNSClient interface {
	UpdateIP(string, string) error
}

type GoogleDDNSClient struct {
	BaseURL  string
	Username string
	Password string
}

func (client *GoogleDDNSClient) UpdateIP(domainName, ipAddress string) error {
	url := fmt.Sprintf("%s/nic/update?hostname=%s&ip=%s", client.BaseURL, domainName, ipAddress)

	httpClient := &http.Client{}

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		return err
	}

	encodedAuth := ""
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodedAuth))

	resp, err := httpClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	log.Printf("Updated IP Address for %s to %s\n", domainName, ipAddress)
	return nil
}
