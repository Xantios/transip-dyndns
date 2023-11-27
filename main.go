package main

import (
	"fmt"
	externalip "github.com/glendc/go-external-ip"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/domain"
	"net"
	"os"
	"time"
)

func getWanIP() (net.IP, error) {
	cons := externalip.DefaultConsensus(nil, nil)
	cons.UseIPProtocol(4)
	return cons.ExternalIP()
}

func setDynDns(wanIp net.IP) {

	accountName := os.Getenv("USERNAME")
	if accountName == "" {
		fmt.Printf("Error: Missing account name. Read readme.md \n")
		os.Exit(1)
	}

	domainName := os.Getenv("DOMAIN")
	if domainName == "" {
		fmt.Printf("Error: Missing domain name. Read readme.md\n")
		os.Exit(2)
	}

	subDomain := os.Getenv("SUBDOMAIN")
	if subDomain == "" {
		fmt.Printf("Error: Missing subdomain. Read readme.md\n")
		os.Exit(3)
	}

	client, err := gotransip.NewClient(gotransip.ClientConfiguration{
		AccountName:    accountName,
		PrivateKeyPath: "./key.private",
	})

	if err != nil {
		fmt.Printf("Cant create client. %s\n", err)
		return
	}

	domains := domain.Repository{Client: client}
	err = domains.UpdateDNSEntry(domainName, domain.DNSEntry{
		Name:    subDomain,
		Expire:  60,
		Type:    "A",
		Content: wanIp.String(),
	})

	if err != nil {
		fmt.Printf("Error while updating DNS entry :: %s\n", err)
	}
}

func main() {

	fmt.Printf("TransIP DynDNS client - Started\n")

	for {
		wan, _ := getWanIP()
		setDynDns(wan)

		time.Sleep(120 * time.Second)
	}
}
