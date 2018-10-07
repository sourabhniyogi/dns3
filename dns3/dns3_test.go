package dns3

import (
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	_ "github.com/mattn/go-sqlite3"
	"github.com/miekg/dns"
)

// Core test
func TestDNSRequest(t *testing.T) {
	// Problem: Lookup the DNS value of "dev.eth.hacker" and get "35.77.66.55"
	request := "dev.eth.hacker"
	expectedResponse := "35.77.66.55"

	// Solution: Use Trustless DNS with Ethereum Contract!

	// 1. Parse the request
	domain, tld, err := ParseDomain(request)
	if err != nil {
		t.Fatalf("ParseDomain err %v", err)
	}
	fmt.Printf("DNS3 Request:\t%s\n", request)
	fmt.Printf("  tld:\t%s\n", tld) // TODO Integrate with Handshake
	fmt.Printf("  domain:\t%s\n", domain)
	domainHash := Keccak256([]byte(domain))
	var domainHash32 [32]byte
	copy(domainHash32[:], domainHash[:])
	fmt.Printf("  domainHash:\t0x%x\n", domainHash)

	// 2. Get the Zone File Hash from Ethereum, turn it into a URL
	ipfsHash, ipfsHash58, err := GetZone(domainHash32)
	if err != nil {
		t.Fatalf("GetZone Err: %v", err)
	}
	fmt.Printf("DNS3.sol Call:\tgetZone(0x%x)\n", domainHash)
	fmt.Printf("  ipfsHash:\t%x => %s", ipfsHash, ipfsHash58)
	ipfsUrl := fmt.Sprintf("https://cloudflare-ipfs.com/ipfs/%s", ipfsHash58)

	// 3. Lookup DNS in IPFS Url
	fmt.Printf("  IPFS Lookup:\t%s", ipfsUrl)
	result, err := LookupDNS(ipfsUrl, request)
	if err != nil {
		t.Fatalf("LookupDNS Err: %v", err)
	}
	fmt.Printf("... DONE\n")
	fmt.Printf("  DNS3 Result:\t%s\n", result)
	if strings.Compare(result, expectedResponse) != 0 {
		t.Fatalf("Failure to get expected response %s", expectedResponse)
	}
}

// Tests Submit Zone in Go
func TestSubmitZone(t *testing.T) {
	key, _ := crypto.HexToECDSA(PrivateKey)
	session, err := setSession(common.HexToAddress(DNS3ContractAddr), wsEndpointUrl, key)
	if err != nil {
		t.Fatalf("setSession %v", err)
	}

	// sample data
	sample := "QmXkTBPtuJ1pTYRQ1U4AsSgAy1vE7r1EaMSAJ4pKMkZj89"
	//	ipfs := base58.Decode(sample)
	hashtype, digest, err := IPFSHashToBytes(sample)
	if err != nil {
		t.Fatalf("IPFSHash %v\n", err)
	}
	var ipfsdigest [32]byte
	copy(ipfsdigest[:], digest[:])
	sz := len(digest)
	domain := "eth.hacker"
	domainHash0 := Keccak256([]byte(domain))
	var domainHash [32]byte
	copy(domainHash[0:32], domainHash0[0:32])
	domainHash0, _ = hex.DecodeString("b63f160a960a1663c5cec1d7d02e67a44d368affd1d42be3b3554c34fd2dea4b")

	// SubmitZone
	tx, err := session.SubmitZone(ipfsdigest, hashtype, uint8(sz), domainHash)
	if err != nil {
		t.Fatalf("submitZone %v", err)
	}
	fmt.Printf("tx: %x\n", tx.Hash())
}

func AnotherHelloServer(w dns.ResponseWriter, req *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(req)

	m.Extra = make([]dns.RR, 1)
	m.Extra[0] = &dns.TXT{Hdr: dns.RR_Header{Name: m.Question[0].Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 0}, Txt: []string{"104.154.155.233"}}
	w.WriteMsg(m)
}

// Demonstrates local DNS Server idea
func TestLocal(t *testing.T) {
	dns.HandleFunc("eth.hacker.", AnotherHelloServer)
	defer dns.HandleRemove("eth.hacker.")

	waitLock := sync.Mutex{}
	server := &dns.Server{Addr: ":0", Net: "udp", ReadTimeout: time.Hour, WriteTimeout: time.Hour, NotifyStartedFunc: waitLock.Unlock}
	waitLock.Lock()

	go func() {
		server.ListenAndServe()
	}()
	waitLock.Lock()

	c, m := new(dns.Client), new(dns.Msg)
	m.SetQuestion("eth.hacker.", dns.TypeTXT)
	addr := server.PacketConn.LocalAddr().String() // Get address via the PacketConn that gets set.
	r, _, err := c.Exchange(m, addr)
	if err != nil {
		t.Fatal("failed to exchange eth.hacker", err)
	}
	txt := r.Extra[0].(*dns.TXT).Txt[0]
	if txt != "104.154.155.233" {
		t.Error("unexpected result for eth.hacker", txt, "!= 104.154.155.233")
	}
	fmt.Printf("%v\n", txt)
	server.Shutdown()
}
