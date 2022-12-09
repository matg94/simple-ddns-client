package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DDNSClient interface {
	UpdateIP(string) error
	GetInterval() int
}

type GoogleDDNSClient struct {
	BaseURL    string
	Username   string
	Password   string
	DomainName string
	Interval   int
}

func (client *GoogleDDNSClient) GetInterval() int {
	return client.Interval
}

func (client *GoogleDDNSClient) UpdateIP(ipAddress string) error {
	url := fmt.Sprintf("%s/nic/update?hostname=%s&ip=%s", client.BaseURL, client.DomainName, ipAddress)

	httpClient := &http.Client{}

	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		return err
	}

	encodedAuth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", client.Username, client.Password)))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodedAuth))

	resp, err := httpClient.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if string(body) == "badauth" {
		return errors.New("bad auth error")
	}

	log.Printf("Status: %d | Updated IP Address for %s to %s\n", resp.StatusCode, client.DomainName, ipAddress)
	return nil
}
