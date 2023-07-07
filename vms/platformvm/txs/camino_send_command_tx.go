// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"github.com/pkg/errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/coreth/commands"
)

var (
	_ UnsignedTx = (*SendCommandTx)(nil)

	errNopForbidden      = errors.New("no NOP tx allowed")
	errOnlyCChainAllwoed = errors.New("only CChain commands are allowed")
)

// SendCommandTx is an unsigned SendCommandTx
type SendCommandTx struct {
	BaseTx `serialize:"true"`

	// futureproofing to send commands to other chains
	DestinationChain ids.ID `serialize:"true" json:"destinationChain"`

	// what to execute on the other chain
	Command commands.ExternalCommand `serialize:"true", json:"command"`
}

// InitCtx sets the FxID fields in the inputs and outputs of this
// [UnsignedSendCommandTx]. Also sets the [ctx] to the given [vm.ctx] so that
// the addresses can be json marshalled into human readable format
func (tx *SendCommandTx) InitCtx(ctx *snow.Context) {
	tx.BaseTx.InitCtx(ctx)
}

// SyntacticVerify this transaction is well-formed
func (tx *SendCommandTx) SyntacticVerify(ctx *snow.Context) error {
	switch {
	case tx == nil:
		return ErrNilTx
	case tx.Command == nil:
		return errNopForbidden
	case tx.DestinationChain != ctx.CChainID:
		return errOnlyCChainAllwoed
	case tx.SyntacticallyVerified: // already passed syntactic verification
		return nil
	}

	if err := tx.BaseTx.SyntacticVerify(ctx); err != nil {
		return err
	}

	if err := tx.Command.Verify(); err != nil {
		return err
	}

	tx.SyntacticallyVerified = true
	return nil
}

func (tx *SendCommandTx) Visit(visitor Visitor) error {
	return visitor.SendCommandTx(tx)
}
