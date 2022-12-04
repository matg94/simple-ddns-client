package main

import (
	"io/ioutil"
	"log"
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
		log.Println("Failed to get ip address")
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read ip address", err)
		return ""
	}
	ip.lastIp = string(body)
	return string(body)
}
