package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wolkdb/dns3/dns3"
	"strings"
)

func main() {
	// Problem: Lookup the DNS value of "dev.eth.hacker" and get "35.77.66.55"
	request := "dev.eth.hacker"
	expectedResponse := "35.77.66.55"

	// Solution: Use Trustless DNS with Ethereum Contract!

	// 1. Parse the request
	domain, tld, err := dns3.ParseDomain(request)
	if err != nil {
		fmt.Printf("ParseDomain err %v", err)
	}
	fmt.Printf("DNS3 Request:\t%s\n", request)
	fmt.Printf("  tld:\t%s\n", tld) // TODO Integrate with Handshake
	fmt.Printf("  domain:\t%s\n", domain)
	domainHash := dns3.DomainHash(domain)
	fmt.Printf("  domainHash:\t0x%x\n", domainHash)

	// 2. Get the Zone File Hash from Ethereum, turn it into a URL
	ipfsHash, ipfsHash58, err := dns3.GetZone(domainHash)
	if err != nil {
		fmt.Printf("GetZone Err: %v", err)
	}
	fmt.Printf("DNS3.sol Call:\tgetZone(0x%x)\n", domainHash)
	fmt.Printf("  ipfsHash:\t%x => %s", ipfsHash, ipfsHash58)
	ipfsUrl := fmt.Sprintf("https://cloudflare-ipfs.com/ipfs/%s", ipfsHash58)

	// 3. Lookup DNS in IPFS Url
	fmt.Printf("  IPFS Lookup:\t%s", ipfsUrl)
	result, found, err := dns3.LookupDNS(ipfsUrl, request)
	if err != nil {
		fmt.Printf("LookupDNS Err: %v", err)
	}
	if found {
		fmt.Printf("... FOUND\n")
		fmt.Printf("  DNS3 Result:\t%s\n", result)
		if strings.Compare(result, expectedResponse) != 0 {
			fmt.Printf("Failure to get expected response %s", expectedResponse)
		}
	} else {
		fmt.Printf("... NOT FOUND\n")
	}
}
