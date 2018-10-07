package dns3

import (
	// "database/sql"

	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/mattn/go-sqlite3"
	// "encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func TestDNS3(t *testing.T) {
	//  User:   0x12233992092D7B405355D771940E5115c17f959F
	PrivateKey := "a5718e79ae2fe43431820cba7315f48ac0a79e5305da6988c9f3358003784d85"
	key, _ := crypto.HexToECDSA(PrivateKey)

	contractAddr := "0x8116a77cf44457a455ffc24001c521ddeebc9606"
	ws_endpointUrl := "wss://rinkeby.infura.io/ws"
	session, err := setSession(common.HexToAddress(contractAddr), ws_endpointUrl, key)
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

	// RegisterDomain
	/*
		tx, err := dns3.RegisterDomain(auth, domain)
		if err != nil {
			t.Fatalf("RegisterDomain %v", err)
		}
		fmt.Printf("tx: %v\n", tx)
	*/

	// SubmitZone
	tx, err := session.SubmitZone(ipfsdigest, hashtype, uint8(sz), domainHash)
	if err != nil {
		t.Fatalf("submitZone %v", err)
	}
	fmt.Printf("tx: %x\n", tx.Hash())
}
