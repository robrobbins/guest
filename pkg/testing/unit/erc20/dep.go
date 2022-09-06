package erc20

import (
	"github.com/robrobbins/guest/pkg/testing/unit"
)

type Dep struct {
	ERC20     *ERC20
	ERC20Test *ERC20Test
}

func Deploy(e *unit.Env) (*Dep, error) {
	assertAddr, _, _, err := unit.DeployAssert(e.Admin.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	ercAddr, _, ercContract, err := DeployERC20(e.Admin.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	_, _, ercTestContract, err := DeployERC20Test(e.Admin.Opts, e.Blockchain, assertAddr, ercAddr)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	return &Dep{
		ERC20:     ercContract,
		ERC20Test: ercTestContract,
	}, nil
}
