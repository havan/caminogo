// Copyright (C) 2022, Chain4Travel AG. All rights reserved.
//
// This file is a derived work, based on ava-labs code whose
// original notices appear below.
//
// It is distributed under the same license conditions as the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********************************************************

// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"errors"
	"fmt"

	"github.com/chain4travel/caminogo/chains/atomic"
	"github.com/chain4travel/caminogo/database"
	"github.com/chain4travel/caminogo/ids"
	"github.com/chain4travel/caminogo/snow"
	"github.com/chain4travel/caminogo/vms/components/avax"
	"github.com/chain4travel/caminogo/vms/components/verify"
)

var (
	errAlreadyPresentKey = errors.New("key already present, wait for target chain to consume")
	errKeyTooLarge       = errors.New("key exceeds the 50 byte limit")
	errValueTooLarge     = errors.New("value exceeds the 50 byte limit")

	AtomicWritePrefix                  = []byte("ATOMIC")
	_                 UnsignedAtomicTx = &UnsignedWriteAtomicTx{}

	maxValueBytes = 50
	maxKeyBytes   = 50
)

// UnsignedWriteAtomicTx is a debug tx that is used to arbitraily write to the shared memory of other primary chains
// It encasupates the logic to write to shared memory
type UnsignedWriteAtomicTx struct {
	avax.Metadata
	Key           []byte `serialize:"true" json:"key"`
	Value         []byte `serialize:"true" json:"value"`
	TargetChainID ids.ID `serialize:"true" json:"targetChain"`
}

// satisfy the atomicTx interface
func (tx *UnsignedWriteAtomicTx) InitCtx(ctx *snow.Context) {
}

func (tx *UnsignedWriteAtomicTx) InputIDs() ids.Set {
	return nil
}

// SyntacticVerify this transaction is well-formed
func (tx *UnsignedWriteAtomicTx) SyntacticVerify(ctx *snow.Context) error {
	if len(tx.GetKeyWithPrefix()) > maxKeyBytes {
		return errKeyTooLarge
	}

	if len(tx.Value) > maxValueBytes {
		return errValueTooLarge
	}

	return nil
}

func (tx *UnsignedWriteAtomicTx) InputUTXOs() ids.Set { return nil }

// Attempts to verify this transaction with the provided state.
func (tx *UnsignedWriteAtomicTx) SemanticVerify(vm *VM, parentState MutableState, stx *Tx) error {
	_, err := tx.AtomicExecute(vm, parentState, stx)
	return err
}

// Execute this transaction.
func (tx *UnsignedWriteAtomicTx) Execute(
	vm *VM,
	vs VersionedState,
	stx *Tx,
) (
	func() error,
	error,
) {
	switch {
	case tx == nil:
		return nil, errNilTx
	case len(stx.Creds) != 0:
		return nil, errWrongNumberOfCredentials
	}

	if err := tx.SyntacticVerify(vm.ctx); err != nil {
		return nil, err
	}

	if vm.bootstrapped.GetValue() {
		if err := verify.SameSubnet(vm.ctx, tx.TargetChainID); err != nil {
			return nil, err
		}
	}

	// TODO @jax once proposals are working we need to check if the provided proposal has concluded, and it would be best if we grab the base fee from
	// TODO the proposal instead of getting it from the TX currently i assume the propose_block function reads all active proposals and checks which of them need
	// TODO need to be executed. Still i assume anyone can just craft this TX so we should make sure the proposal is first of all valid

	// TODO @jax add these methods, we should be able across the codebase to know if a proposal is valid or is in effect
	// if err := isProposalValid(vs, tx.ProposalID); err != nil {
	// 	return nil, err
	// }

	// if err := isProposalInEffect(vs, tx.ProposalID); err != nil {
	// 	return nil, err
	// }

	chainID, requests, err := tx.AtomicOperations()
	if err != nil {
		return nil, fmt.Errorf("unable to build atomic operations: %v", err)
	}

	if err := vm.ctx.SharedMemory.Apply(map[ids.ID]*atomic.Requests{chainID: requests}); err != nil {
		//TODO @jax this shadows other errors
		return nil, errAlreadyPresentKey
	}

	return nil, nil
}

// AtomicOperations returns the shared memory requests
func (tx *UnsignedWriteAtomicTx) AtomicOperations() (ids.ID, *atomic.Requests, error) {

	elems := make([]*atomic.Element, 1)
	elems[0] = &atomic.Element{
		Key:   tx.GetKeyWithPrefix(),
		Value: tx.Value,
	}

	return tx.TargetChainID, &atomic.Requests{PutRequests: elems}, nil
}

// [AtomicExecute] to maintain consistency for the standard block.
func (tx *UnsignedWriteAtomicTx) AtomicExecute(
	vm *VM,
	parentState MutableState,
	stx *Tx,
) (VersionedState, error) {
	// Set up the state if this tx is committed
	newState := newVersionedState(
		parentState,
		parentState.CurrentStakerChainState(),
		parentState.PendingStakerChainState(),
	)
	_, err := tx.Execute(vm, newState, stx)
	return newState, err
}

// this is required to not allow arbitrary request to create utxos on other chains
func (tx *UnsignedWriteAtomicTx) GetKeyWithPrefix() []byte {
	return PrefixKeyForAtomic(tx.Key)
}

// Impl of the prefix logic
func PrefixKeyForAtomic(key []byte) []byte {
	prefixed := make([]byte, len(key)+len(AtomicWritePrefix))
	copy(prefixed, AtomicWritePrefix)
	copy(prefixed[len(AtomicWritePrefix):], prefixed)

	return prefixed
}

// Accept this transaction and write the data to the target chains shared memory
func (tx *UnsignedWriteAtomicTx) AtomicAccept(ctx *snow.Context, batch database.Batch) error {
	chainID, requests, err := tx.AtomicOperations()
	if err != nil {
		return err
	}
	return ctx.SharedMemory.Apply(map[ids.ID]*atomic.Requests{chainID: requests}, batch)
}

// Create a new transaction
func (vm *VM) newWriteAtomicTx(
	key []byte,
	value []byte,
	targetChainId ids.ID,
) (*Tx, error) {
	tx := &Tx{UnsignedTx: &UnsignedWriteAtomicTx{
		Key:           key,
		Value:         value,
		TargetChainID: targetChainId,
	}}
	return tx, tx.Sign(Codec, nil)
}
