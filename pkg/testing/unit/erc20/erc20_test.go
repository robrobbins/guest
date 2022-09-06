package erc20

import (
	"math/big"
	"os"
	"testing"

	"github.com/robrobbins/guest/pkg/testing/unit"
)

var dep *Dep
var env *unit.Env
var depErr error

func TestERC20(t *testing.T) {
	// &bind.TransactOpts{From: env.Admin.Opts.From, Signer: env.Admin.Opts.Signer}

	// assure contracts are deployed
	if depErr != nil {
		t.Fatal(depErr)
	}

	if err := dep.ERC20Test.Run(nil); err != nil {
		t.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	env := unit.NewEnv(big.NewInt(1000000000000000000))
	dep, depErr = Deploy(env)

	os.Exit(m.Run())
}
