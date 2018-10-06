package dns3

import (
	// "database/sql"

	"fmt"
	"log"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	// "encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestDNS3(t *testing.T) {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("https://rinkeby.infura.io/metamask")

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	//  C0    a5718e79ae2fe43431820cba7315f48ac0a79e5305da6988c9f3358003784d85    0x12233992092D7B405355D771940E5115c17f959F
	var key = `{"address":"90fb0de606507e989247797c6a30952cae4d5cbe","crypto":{"cipher":"aes-128-ctr","ciphertext":"54396d6ed0335e4b4874cd4440d24eabeca895fcbafb15d310c25c6b1e4bb306","cipherparams":{"iv":"e3a2457cf8420d3072e5adf118d31df8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"d25987f2f2429e53f51d87eb6474e3f12a67c63603fd860b558657cee19a6ea9"},"mac":"023fc8a29a6e323db43e0c7795d2d59d0c1f295a62cbb9bc625951fca9c385dd"},"id":"dc849ada-c6be-4f12-bfa2-5200ec560c2e","version":3}`
	auth, err := bind.NewTransactor(strings.NewReader(key), "mdotm")
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Instantiate the contract and display its name
	dns3, err := NewDNS3(common.HexToAddress("0x8116a77cf44457a455ffc24001c521ddeebc9606"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Simplestens contract: %v", err)
	}

	sample := "QmXkTBPtuJ1pTYRQ1U4AsSgAy1vE7r1EaMSAJ4pKMkZj89"
	hashtype, ipfsHashByte, err := IPFSHashToBytes(sample)
	if err != nil {
		t.Fatalf("IPFSHashToBytes %v", err)
	}
	fmt.Printf("TestIPFS: %d %x (%d bytes)\n", hashtype, ipfsHashByte, len(ipfsHashByte))

	domainHash0 := Keccak256([]byte("eth.hacker"))
	var domainHash [32]byte
	copy(domainHash[0:32], domainHash0[0:32])
	tx, err := dns3.SubmitZone(auth, ipfsHashByte, domainHash)
	if err != nil {
		t.Fatalf("submitZone %v", err)
	}
	fmt.Printf("tx: %v\n", tx)
}
