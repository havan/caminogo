package commands

import (
	"fmt"
	"math/big"

	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/coreth/core/state"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	KYC_OFFSET = uint(1)
	KYB_OFFSET = uint(8)
)

type KYCUpdate struct {
	Address     common.Address `serialize:"true", json:"address"`
	KYCVerified bool           `serialize:"true", json:"kyc_verified"`
	KYBVerified bool           `serialize:"true", json:"kyb_verified"`
}

type CommandSetKYC struct {
	KYCUpdates []KYCUpdate `serialize:"true", json:"kyc_updates"`
}

func (cmd *CommandSetKYC) Verify() error {

	set := make(map[common.Address]struct{})
	for _, update := range cmd.KYCUpdates {
		set[update.Address] = struct{}{}
	}

	if len(set) != len(cmd.KYCUpdates) {
		return fmt.Errorf("no updates entries allowed")
	}

	return nil
}

func (cmd *CommandSetKYC) EVMStateTransfer(ctx *snow.Context, state *state.StateDB) error {

	for _, update := range cmd.KYCUpdates {
		kycState := big.NewInt(0)
		if update.KYCVerified {
			kycState = kycState.Or(kycState, common.Big1.Lsh(common.Big1, KYC_OFFSET))
		}

		if update.KYBVerified {
			kycState = kycState.Or(kycState, common.Big1.Lsh(common.Big1, KYB_OFFSET))
		}

		state.SetState(AdminContractAddr, KycRoleKeyFromAddr(update.Address), common.BigToHash(kycState))
	}

	return nil
}

func KycRoleKeyFromAddr(addr common.Address) common.Hash {
	return crypto.Keccak256Hash(addr.Hash().Bytes(), common.HexToHash("0x2").Bytes() /*slot 2 reference admin.sol map(address => uint)*/)
}
