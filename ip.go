package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type IpAddressProvider interface {
	GetLastIP() string
	GetLatestIP() string
}

type DefaultIpAddressProvider struct {
	lastIp string
}

func (ip *DefaultIpAddressProvider) GetLastIP() string {
	return ip.lastIp
}

func (ip *DefaultIpAddressProvider) GetLatestIP() string {
	resp, err := http.Get("https://ifconfig.me")
	if err != nil {
		fmt.Println("Failed to get ip address")
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read ip address", err)
		return ""
	}

	return string(body)
}
