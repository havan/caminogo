package commands

import (
	"fmt"
	"math/big"

	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/coreth/core/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type CommandSetBaseFee struct {
	BaseFee uint64 `serialize:"true"`
}

func (cmd *CommandSetBaseFee) Verify() error {

	switch {
	case cmd.BaseFee > 0:
		return fmt.Errorf("baseFee has to be greater that 0")
	}

	return nil
}

func (cmd *CommandSetBaseFee) EVMStateTransfer(ctx *snow.Context, state *state.StateDB) error {
	state.SetState(AdminContractAddr, crypto.Keccak256Hash(common.HexToHash("0x1").Bytes()), common.BigToHash(big.NewInt(int64(cmd.BaseFee))))
	return nil
}
