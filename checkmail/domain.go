package main

import (
	"fmt"
	"net"
)

type DomainInfo struct {
	HasMXRecord bool      // whether has 1 or more MX record
	Records     []*net.MX // represent DNS MX records
}

// prüft Domäne (MX, spf und dmarc)
func CheckDomain(domain string) (*DomainInfo, error) {
	// TODO: verwende net.LookupMX und net.LookupTXT
	mx, err := net.LookupMX(domain)
	if err != nil {
		return nil, err
	}
	//fmt.Println(mx)
	txt, err := net.LookupTXT(domain)
	if err != nil {
		return nil, err
	}
	fmt.Println(txt)
	return &DomainInfo{
		HasMXRecord: len(mx) > 0,
		Records:     mx,
	}, nil
}

func (m *DomainInfo) String() string {
	res := fmt.Sprintf("MXRecord: %v\n", m.HasMXRecord)
	for _, v := range m.Records {
		res += fmt.Sprintf("%v\n", *v)
	}
	return res
}
