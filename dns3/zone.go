package dns3

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	UserAddress      = "0x12233992092D7B405355D771940E5115c17f959F"
	PrivateKey       = "a5718e79ae2fe43431820cba7315f48ac0a79e5305da6988c9f3358003784d85"
	DNS3ContractAddr = "0x8c36f7e95f53b5ee7a35ec2dad854308877a0a94"
	wsEndpointUrl    = "wss://rinkeby.infura.io/ws"
)

// For HTTP calls, can be tuned based on how IPFS nodes work
var DefaultTransport http.RoundTripper = &http.Transport{
	Dial: (&net.Dialer{
		// limits the time spent establishing a TCP connection (if a new one is needed)
		Timeout:   5 * time.Second,
		KeepAlive: 3 * time.Second, // 60 * time.Second,
	}).Dial,

	MaxIdleConnsPerHost: 100,

	// limits the time spent reading the headers of the response.
	ResponseHeaderTimeout: 5 * time.Second,
	IdleConnTimeout:       4 * time.Second, // 90 * time.Second,

	// limits the time the client will wait between sending the request headers when including an Expect: 100-continue and receiving the go-ahead to send the body.
	ExpectContinueTimeout: 1 * time.Second,

	// limits the time spent performing the TLS handshake.
	TLSHandshakeTimeout: 5 * time.Second,
}

func setConnection(endpointUrl string) (conn *ethclient.Client, err error) {
	conn, err = ethclient.Dial(endpointUrl)
	if err != nil {
		return conn, err
	} else {
		//fmt.Printf("Successfully connected to: %v\n", endpointUrl)
		return conn, err
	}
}

func setSession(contractAddr common.Address, endpointUrl string, key *ecdsa.PrivateKey) (session *DNS3Session, err error) {
	// Instantiate the contract {caller, transactor, filterer}, whitout
	// setting Pre-Configured CallOpts and TransactOpts
	conn, err := setConnection(endpointUrl)
	if err != nil {
		return session, err
	}
	contract, err := NewDNS3(contractAddr, conn)
	if err != nil {
		fmt.Printf("Failed to instantiate session to contract: %v", err)
		return session, err
	}

	authT := bind.NewKeyedTransactor(key)
	session = &DNS3Session{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     authT.From,
			Signer:   authT.Signer,
			GasLimit: 500000,
		},
	}
	return session, nil
}

func DomainHash(domain string) (domainHash [32]byte) {
	domainHash0 := Keccak256([]byte(domain))
	copy(domainHash[0:32], domainHash0[0:32])
	return domainHash
}

// Parse a full domain (e.g. www.example.com) into the core domain (example.com) and tld (hacker)
func ParseDomain(request string) (domain string, tld string, err error) {
	// [Sourabh] TODO: parse "dev.eth.hacker" into "eth.hacker"
	pieces := strings.Split(request, ".")
	if len(pieces) < 2 {
		return domain, tld, fmt.Errorf("Invalid domain")
	}
	tld = pieces[len(pieces)-1]
	domainarr := pieces[len(pieces)-2 : len(pieces)]
	domain = strings.Join(domainarr, ".")
	return domain, tld, nil
}

// Look up in Ethereum the zone file hash based on the domain hash
func GetZone(domainHash [32]byte) (ipfsHash [32]byte, ipfsHash58 string, err error) {

	key, _ := crypto.HexToECDSA(PrivateKey)
	session, err := setSession(common.HexToAddress(DNS3ContractAddr), wsEndpointUrl, key)
	if err != nil {
		return ipfsHash, ipfsHash58, fmt.Errorf("setSession %v", err)
	}

	res, err := session.GetZone(domainHash)
	if err != nil {
		return ipfsHash, ipfsHash58, err
	}
	// 0xb63f160a960a1663c5cec1d7d02e67a44d368affd1d42be3b3554c34fd2dea4b
	ipfsHash58 = BuildIPFSHash(uint8(18), res.Digest[:])
	return res.Digest, ipfsHash58, nil
}

// Given an IPFS URL (assumed to be a valid Zone File)
// and a request, does an HTTP GET and parses out an "A" record entry
func LookupDNS(url string, request string) (result string, found bool, err error) {
	// HTTP Get of ipfsurl holding a
	httpclient := &http.Client{Timeout: time.Second * 5, Transport: DefaultTransport}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, false, err
	}
	resp, do_err := httpclient.Do(req)
	if do_err != nil {
		return result, false, do_err
	}
	body, readall_err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if readall_err != nil {
		return result, false, do_err
	}

	// Parse each line and match against the input request ("www.eth.hacker")
	/* Sample Data: https://cloudflare-ipfs.com/ipfs/QmXThgG1gUnfywM4e9QpEYDkBZNJwSbpPogJjXtewVgYmi
	   www.eth.hacker.     3600   IN  A   104.154.155.233
	   ganache.eth.hacker. 1      IN  A   35.224.4.165
	   remix.eth.hacker.   1      IN  A   52.4.38.70
	   dev.eth.hacker.     3600   IN  A   35.77.66.55 	*/
	records := strings.Split(string(body), "\n")
	for _, rec := range records {
		pieces := strings.Split(strings.Trim(rec, " \n"), " ")
		if strings.Contains(pieces[0], request) { // should check for "IN", "A" record status
			// FOUND! validate if its a real IP, or just return net.IP
			result := pieces[len(pieces)-1]
			return result, true, nil
		}
	}
	return result, false, nil
}
