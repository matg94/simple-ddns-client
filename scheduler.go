package main

import (
	"log"
	"time"
)

type Scheduler interface {
	Start(IpAddressProvider, DDNSClient)
	Execute(IpAddressProvider, DDNSClient)
}

type DefaultScheduler struct {
}

func (s *DefaultScheduler) Start(provider IpAddressProvider, ddns DDNSClient) {
	for {
		go s.Execute(provider, ddns)
		time.Sleep(time.Duration(ddns.GetInterval()) * time.Second)
	}
}

func (s *DefaultScheduler) Execute(provider IpAddressProvider, ddns DDNSClient) {
	currentIp := provider.GetLastIP()
	latestIp := provider.GetLatestIP()
	if latestIp == "" {
		return
	}
	if currentIp == latestIp {
		return
	}
	err := ddns.UpdateIP(latestIp)
	if err != nil {
		log.Println("Failed to update ip", err)
		return
	}
}
