package main

import (
	"log"
	"time"
)

type Scheduler interface {
	Start(IpAddressProvider, DDNSClient, string)
	Execute(IpAddressProvider, DDNSClient, string)
}

type DefaultScheduler struct {
	Interval int
}

func (s *DefaultScheduler) Start(provider IpAddressProvider, ddns DDNSClient, domainName string) {
	for {
		go s.Execute(provider, ddns, domainName)
		time.Sleep(time.Duration(s.Interval) * time.Second)
	}
}

func (s *DefaultScheduler) Execute(provider IpAddressProvider, ddns DDNSClient, domainName string) {
	currentIp := provider.GetLastIP()
	latestIp := provider.GetLatestIP()
	if latestIp == "" {
		return
	}
	if currentIp == latestIp {
		return
	}
	err := ddns.UpdateIP(domainName, latestIp)
	if err != nil {
		log.Println("Failed to update ip", err)
		return
	}
}
