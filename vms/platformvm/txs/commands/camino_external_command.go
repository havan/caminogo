package commands

import (
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/components/verify"
	"github.com/ava-labs/coreth/core/state"
)

type ExternalCommand interface {
	verify.Verifiable
	EVMStateTransfer(ctx *snow.Context, state *state.StateDB) error

	// Maybe this needs equality? (to filter/block out duplicate commands, etc)
}
