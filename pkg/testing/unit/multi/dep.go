package multi

import (
	"github.com/robrobbins/guest/pkg/testing/unit"
)

type Dep struct {
  Multi *Multi
}

func Deploy(e *unit.Env) (*Dep, error) {
	 _, _, _, err := unit.DeployAssert(e.Admin.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	 _, _, multiContract, err := DeployMulti(e.Admin.Opts, e.Blockchain)

	if err != nil {
		return nil, err
	}

	e.Blockchain.Commit()

	// _, _, multiTestContract, err := DeployMultiTest(e.Admin.Opts, e.Blockchain, assertAddr, multiAddr)

	// if err != nil {
		// return nil, err
	// }

	// e.Blockchain.Commit()

	return &Dep{
		Multi:     multiContract,
	}, nil
}
