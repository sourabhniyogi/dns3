package dns3

import (
	// "database/sql"

	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/btcsuite/btcutil/base58"
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

	//  C0        0x12233992092D7B405355D771940E5115c17f959F
	//	var key = `{"address":"90fb0de606507e989247797c6a30952cae4d5cbe","crypto":{"cipher":"aes-128-ctr","ciphertext":"54396d6ed0335e4b4874cd4440d24eabeca895fcbafb15d310c25c6b1e4bb306","cipherparams":{"iv":"e3a2457cf8420d3072e5adf118d31df8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"d25987f2f2429e53f51d87eb6474e3f12a67c63603fd860b558657cee19a6ea9"},"mac":"023fc8a29a6e323db43e0c7795d2d59d0c1f295a62cbb9bc625951fca9c385dd"},"id":"dc849ada-c6be-4f12-bfa2-5200ec560c2e","version":3}`
	//	auth, err := bind.NewTransactor(strings.NewReader(key), "mdotm")
	//	if err != nil {
	//		log.Fatalf("Failed to create authorized transactor: %v", err)
	//	}

	PrivateKey := "a5718e79ae2fe43431820cba7315f48ac0a79e5305da6988c9f3358003784d85"
	key, _ := crypto.HexToECDSA(PrivateKey)

	contractAddr := "0x8116a77cf44457a455ffc24001c521ddeebc9606"
	rpc_endpointUrl := "wss://rinkeby.infura.io/ws" // "https://rinkeby.infura.io/metamask"
	//conn, err := ethclient.Dial("https://rinkeby.infura.io/metamask")
	//if err != nil {
	//	log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	//}

	//	ws_endpointUrl := "wss://rinkeby.infura.io/ws"
	// Create an IPC based RPC connection to a remote node

	//if conn, err := setConnection(ws_endpointUrl); err != nil {
	//	t.Fatalf("setConn %v")
	//}

	session, err := setSession(common.HexToAddress(contractAddr), rpc_endpointUrl, key)
	if err != nil {
		t.Fatalf("setSession %v", err)
	}

	//
	domain := "eth.hacker"
	domainHash0 := Keccak256([]byte(domain))
	var domainHash [32]byte
	copy(domainHash[0:32], domainHash0[0:32])

	// RegisterDomain
	/*
		tx, err := dns3.RegisterDomain(auth, domain)
		if err != nil {
			t.Fatalf("RegisterDomain %v", err)
		}
		fmt.Printf("tx: %v\n", tx)
	*/

	sample := "QmXkTBPtuJ1pTYRQ1U4AsSgAy1vE7r1EaMSAJ4pKMkZj89"
	ipfs := base58.Decode(sample)
	fmt.Printf("Hex: 0x%x", ipfs)
	// SubmitZone
	domainHash0, _ = hex.DecodeString("b63f160a960a1663c5cec1d7d02e67a44d368affd1d42be3b3554c34fd2dea4b")
	//var domainHash [32]byte
	copy(domainHash[:], domainHash0[:])

	tx, err := session.SubmitZone(ipfs, domainHash)
	if err != nil {
		t.Fatalf("submitZone %v", err)
	}
	fmt.Printf("tx: %v\n", tx)
	fmt.Printf("%x\n%s\n", tx.Hash(), tx.Hash())
}
