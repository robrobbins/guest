package unit

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

func NewAuth(c *big.Int) (*ecdsa.PrivateKey, *bind.TransactOpts) {
	pk, _ := crypto.GenerateKey()
	opts, _ := bind.NewKeyedTransactorWithChainID(pk, c)
	return pk, opts
}

// GenBytes32 is a convenience function for some tests where a "GetHash" method is not available.
// Returns a type compatible with bytes32, given a string 32 chars or less.
func GenBytes32(s string) [32]byte {
	bytes := [32]byte{}
	copy(bytes[:], []byte(s))
	return bytes
}
