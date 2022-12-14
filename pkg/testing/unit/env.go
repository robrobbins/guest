package unit

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
)

var chainId = big.NewInt(1337)

// Auth is a custom type which allows us easy access to the dynamically generated
// ecdsa.PrivateKey along with the bind.TransactOpts
type Auth struct {
	PK   *ecdsa.PrivateKey
	Opts *bind.TransactOpts
}

// Env, holds Auth objects capable of signing transactions.
// Also holds the Geth simulated backend.
type Env struct {
	Alloc      core.GenesisAlloc
	Admin      *Auth
	User1      *Auth
	User2      *Auth
	Blockchain *backends.SimulatedBackend
}

// NewEnv returns a hydrated Env struct, ready for use.
// Given a balance argument, it assigns this as the wallet balance for
// each authorization object in the Ctx
func NewEnv(b *big.Int) *Env {
	pk, admin := NewAuth(chainId)
	pk1, u1 := NewAuth(chainId)
	pk2, u2 := NewAuth(chainId)
	alloc := make(core.GenesisAlloc)
	alloc[admin.From] = core.GenesisAccount{Balance: b}
	alloc[u1.From] = core.GenesisAccount{Balance: b}
	alloc[u2.From] = core.GenesisAccount{Balance: b}
	// 2nd arg is a gas limit, a uint64. we'll use 4.7 million
	bc := backends.NewSimulatedBackend(alloc, 4700000)

	return &Env{
		Alloc:      alloc,
		Admin:      &Auth{PK: pk, Opts: admin},
		User1:      &Auth{PK: pk1, Opts: u1},
		User2:      &Auth{PK: pk2, Opts: u2},
		Blockchain: bc,
	}
}
