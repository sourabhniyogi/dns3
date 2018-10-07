package dns3

import (
	"crypto/ecdsa"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	UserAddress      = "0x12233992092D7B405355D771940E5115c17f959F"
	PrivateKey       = "a5718e79ae2fe43431820cba7315f48ac0a79e5305da6988c9f3358003784d85"
	DNS3ContractAddr = "0x8116a77cf44457a455ffc24001c521ddeebc9606"
	wsEndpointUrl    = "wss://rinkeby.infura.io/ws"
)

func setConnection(endpointUrl string) (conn *ethclient.Client, err error) {
	conn, err = ethclient.Dial(endpointUrl)
	if err != nil {
		return conn, err
	} else {
		fmt.Printf("Successfully connected to: %v\n", endpointUrl)
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
			From:   authT.From,
			Signer: authT.Signer,
		},
	}
	return session, nil
}

func DomainHash(domain string) (domainHash [32]byte) {
	domainHash0 := Keccak256([]byte(domain))
	copy(domainHash[0:32], domainHash0[0:32])
	return domainHash
}

func ParseDomain(request string) (domain string, tld string, err error) {
	// [Sourabh] TODO: parse "dev.eth.hacker" into "eth.hacker"
	pieces := strings.Split(request, ".")
	if len(pieces) < 2 {
		return domain, tld, fmt.Errorf("Invalid domain")
	}
	tld = pieces[len(pieces)-1]
	domainarr := pieces[len(pieces)-2 : len(pieces)-1]
	domain = strings.Join(domainarr, ".")
	return domain, tld, nil
}

func GetZone(domainHash [32]byte) (ipfsHash [32]byte, ipfsHash58 string, err error) {
	// [Michael] TODO: set up session and do get Call

	key, _ := crypto.HexToECDSA(PrivateKey)
	session, err := setSession(common.HexToAddress(DNS3ContractAddr), wsEndpointUrl, key)
	if err != nil {
		return ipfsHash, ipfsHash58, fmt.Errorf("setSession %v", err)
	}

	res, err := session.GetZone(domainHash)
	if err != nil {
		return ipfsHash, ipfsHash58, err
	}
	return res.Digest, ipfsHash58, nil
}

func LookupDNS(ipfsUrl string, request string) (result string, err error) {
	// [Sourabh] TODO: HTTP Get of ipfsurl, parse each line and get the request
	return result, nil
}
