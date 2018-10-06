package dns3

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/btcsuite/btcutil/base58"
)

func IPFSHashToBytes(h string) (hashtype uint8, digest []byte, err error) {
	ipfs := base58.Decode(h)
	// 12 20 8bd215d69eef287f6cc093501672d6b1f8c908bf5adedc47993d5676e39577b6
	hashtype = ipfs[0]
	sz := int(ipfs[1])
	digest = ipfs[2:]
	if len(digest) != sz {
		return hashtype, digest, fmt.Errorf("incorrect size")
	}
	return hashtype, digest, nil
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
