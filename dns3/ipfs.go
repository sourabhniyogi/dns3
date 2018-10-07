package dns3

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

func Keccak256(data ...[]byte) []byte {
	hasher := sha3.NewKeccak256()
	for _, b := range data {
		hasher.Write(b)
	}
	return hasher.Sum(nil)
}

func IPFSHashToBytes(h string) (hashtype uint8, digest []byte, err error) {
	ipfs := base58.Decode(h)

	hashtype = ipfs[0]
	sz := int(ipfs[1])
	digest = ipfs[2:]
	if len(digest) != sz {
		return hashtype, digest, fmt.Errorf("incorrect size")
	}
	return hashtype, digest, nil
}

func BuildIPFSHash(hashtype uint8, digest []byte) (ipfshash58 string) {
	prefix := make([]byte, 2)
	prefix[0] = hashtype
	prefix[1] = byte(len(digest))
	b := append(prefix, digest...)
	return base58.Encode(b)
}

func submitZone(fn string) {
	cmd := exec.Command("ipfs", "add", fn)
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}

func getZone(fn string) {
	cmd := exec.Command("ipfs", "cat", fn)
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}
