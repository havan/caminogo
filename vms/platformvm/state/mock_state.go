// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/vms/platformvm/state (interfaces: State)

// Package state is a generated GoMock package.
package state

import (
	reflect "reflect"
	time "time"

	database "github.com/ava-labs/avalanchego/database"
	ids "github.com/ava-labs/avalanchego/ids"
	choices "github.com/ava-labs/avalanchego/snow/choices"
	validators "github.com/ava-labs/avalanchego/snow/validators"
	avax "github.com/ava-labs/avalanchego/vms/components/avax"
	blocks "github.com/ava-labs/avalanchego/vms/platformvm/blocks"
	genesis "github.com/ava-labs/avalanchego/vms/platformvm/genesis"
	locked "github.com/ava-labs/avalanchego/vms/platformvm/locked"
	status "github.com/ava-labs/avalanchego/vms/platformvm/status"
	txs "github.com/ava-labs/avalanchego/vms/platformvm/txs"
	gomock "github.com/golang/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// Abort mocks base method.
func (m *MockState) Abort() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Abort")
}

// Abort indicates an expected call of Abort.
func (mr *MockStateMockRecorder) Abort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abort", reflect.TypeOf((*MockState)(nil).Abort))
}

// AddChain mocks base method.
func (m *MockState) AddChain(arg0 *txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddChain", arg0)
}

// AddChain indicates an expected call of AddChain.
func (mr *MockStateMockRecorder) AddChain(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddChain", reflect.TypeOf((*MockState)(nil).AddChain), arg0)
}

// AddRewardUTXO mocks base method.
func (m *MockState) AddRewardUTXO(arg0 ids.ID, arg1 *avax.UTXO) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddRewardUTXO", arg0, arg1)
}

// AddRewardUTXO indicates an expected call of AddRewardUTXO.
func (mr *MockStateMockRecorder) AddRewardUTXO(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRewardUTXO", reflect.TypeOf((*MockState)(nil).AddRewardUTXO), arg0, arg1)
}

// AddStatelessBlock mocks base method.
func (m *MockState) AddStatelessBlock(arg0 blocks.Block, arg1 choices.Status) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddStatelessBlock", arg0, arg1)
}

// AddStatelessBlock indicates an expected call of AddStatelessBlock.
func (mr *MockStateMockRecorder) AddStatelessBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStatelessBlock", reflect.TypeOf((*MockState)(nil).AddStatelessBlock), arg0, arg1)
}

// AddSubnet mocks base method.
func (m *MockState) AddSubnet(arg0 *txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSubnet", arg0)
}

// AddSubnet indicates an expected call of AddSubnet.
func (mr *MockStateMockRecorder) AddSubnet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubnet", reflect.TypeOf((*MockState)(nil).AddSubnet), arg0)
}

// AddSubnetTransformation mocks base method.
func (m *MockState) AddSubnetTransformation(arg0 *txs.Tx) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddSubnetTransformation", arg0)
}

// AddSubnetTransformation indicates an expected call of AddSubnetTransformation.
func (mr *MockStateMockRecorder) AddSubnetTransformation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubnetTransformation", reflect.TypeOf((*MockState)(nil).AddSubnetTransformation), arg0)
}

// AddTx mocks base method.
func (m *MockState) AddTx(arg0 *txs.Tx, arg1 status.Status) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddTx", arg0, arg1)
}

// AddTx indicates an expected call of AddTx.
func (mr *MockStateMockRecorder) AddTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTx", reflect.TypeOf((*MockState)(nil).AddTx), arg0, arg1)
}

// AddUTXO mocks base method.
func (m *MockState) AddUTXO(arg0 *avax.UTXO) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddUTXO", arg0)
}

// AddUTXO indicates an expected call of AddUTXO.
func (mr *MockStateMockRecorder) AddUTXO(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUTXO", reflect.TypeOf((*MockState)(nil).AddUTXO), arg0)
}

// CaminoGenesisState mocks base method.
func (m *MockState) CaminoGenesisState() (*genesis.Camino, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CaminoGenesisState")
	ret0, _ := ret[0].(*genesis.Camino)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CaminoGenesisState indicates an expected call of CaminoGenesisState.
func (mr *MockStateMockRecorder) CaminoGenesisState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CaminoGenesisState", reflect.TypeOf((*MockState)(nil).CaminoGenesisState))
}

// Close mocks base method.
func (m *MockState) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStateMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockState)(nil).Close))
}

// Commit mocks base method.
func (m *MockState) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockStateMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockState)(nil).Commit))
}

// CommitBatch mocks base method.
func (m *MockState) CommitBatch() (database.Batch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitBatch")
	ret0, _ := ret[0].(database.Batch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommitBatch indicates an expected call of CommitBatch.
func (mr *MockStateMockRecorder) CommitBatch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitBatch", reflect.TypeOf((*MockState)(nil).CommitBatch))
}

// DeleteCurrentDelegator mocks base method.
func (m *MockState) DeleteCurrentDelegator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteCurrentDelegator", arg0)
}

// DeleteCurrentDelegator indicates an expected call of DeleteCurrentDelegator.
func (mr *MockStateMockRecorder) DeleteCurrentDelegator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCurrentDelegator", reflect.TypeOf((*MockState)(nil).DeleteCurrentDelegator), arg0)
}

// DeleteCurrentValidator mocks base method.
func (m *MockState) DeleteCurrentValidator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteCurrentValidator", arg0)
}

// DeleteCurrentValidator indicates an expected call of DeleteCurrentValidator.
func (mr *MockStateMockRecorder) DeleteCurrentValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCurrentValidator", reflect.TypeOf((*MockState)(nil).DeleteCurrentValidator), arg0)
}

// DeletePendingDelegator mocks base method.
func (m *MockState) DeletePendingDelegator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeletePendingDelegator", arg0)
}

// DeletePendingDelegator indicates an expected call of DeletePendingDelegator.
func (mr *MockStateMockRecorder) DeletePendingDelegator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePendingDelegator", reflect.TypeOf((*MockState)(nil).DeletePendingDelegator), arg0)
}

// DeletePendingValidator mocks base method.
func (m *MockState) DeletePendingValidator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeletePendingValidator", arg0)
}

// DeletePendingValidator indicates an expected call of DeletePendingValidator.
func (mr *MockStateMockRecorder) DeletePendingValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePendingValidator", reflect.TypeOf((*MockState)(nil).DeletePendingValidator), arg0)
}

// DeleteUTXO mocks base method.
func (m *MockState) DeleteUTXO(arg0 ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteUTXO", arg0)
}

// DeleteUTXO indicates an expected call of DeleteUTXO.
func (mr *MockStateMockRecorder) DeleteUTXO(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUTXO", reflect.TypeOf((*MockState)(nil).DeleteUTXO), arg0)
}

// GetAddressStates mocks base method.
func (m *MockState) GetAddressStates(arg0 ids.ShortID) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddressStates", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddressStates indicates an expected call of GetAddressStates.
func (mr *MockStateMockRecorder) GetAddressStates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressStates", reflect.TypeOf((*MockState)(nil).GetAddressStates), arg0)
}

// GetChains mocks base method.
func (m *MockState) GetChains(arg0 ids.ID) ([]*txs.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChains", arg0)
	ret0, _ := ret[0].([]*txs.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChains indicates an expected call of GetChains.
func (mr *MockStateMockRecorder) GetChains(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChains", reflect.TypeOf((*MockState)(nil).GetChains), arg0)
}

// GetCurrentDelegatorIterator mocks base method.
func (m *MockState) GetCurrentDelegatorIterator(arg0 ids.ID, arg1 ids.NodeID) (StakerIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentDelegatorIterator", arg0, arg1)
	ret0, _ := ret[0].(StakerIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentDelegatorIterator indicates an expected call of GetCurrentDelegatorIterator.
func (mr *MockStateMockRecorder) GetCurrentDelegatorIterator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentDelegatorIterator", reflect.TypeOf((*MockState)(nil).GetCurrentDelegatorIterator), arg0, arg1)
}

// GetCurrentStakerIterator mocks base method.
func (m *MockState) GetCurrentStakerIterator() (StakerIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentStakerIterator")
	ret0, _ := ret[0].(StakerIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentStakerIterator indicates an expected call of GetCurrentStakerIterator.
func (mr *MockStateMockRecorder) GetCurrentStakerIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentStakerIterator", reflect.TypeOf((*MockState)(nil).GetCurrentStakerIterator))
}

// GetCurrentSupply mocks base method.
func (m *MockState) GetCurrentSupply(arg0 ids.ID) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentSupply", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentSupply indicates an expected call of GetCurrentSupply.
func (mr *MockStateMockRecorder) GetCurrentSupply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentSupply", reflect.TypeOf((*MockState)(nil).GetCurrentSupply), arg0)
}

// GetCurrentValidator mocks base method.
func (m *MockState) GetCurrentValidator(arg0 ids.ID, arg1 ids.NodeID) (*Staker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentValidator", arg0, arg1)
	ret0, _ := ret[0].(*Staker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentValidator indicates an expected call of GetCurrentValidator.
func (mr *MockStateMockRecorder) GetCurrentValidator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentValidator", reflect.TypeOf((*MockState)(nil).GetCurrentValidator), arg0, arg1)
}

// GetLastAccepted mocks base method.
func (m *MockState) GetLastAccepted() ids.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastAccepted")
	ret0, _ := ret[0].(ids.ID)
	return ret0
}

// GetLastAccepted indicates an expected call of GetLastAccepted.
func (mr *MockStateMockRecorder) GetLastAccepted() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastAccepted", reflect.TypeOf((*MockState)(nil).GetLastAccepted))
}

// GetPendingDelegatorIterator mocks base method.
func (m *MockState) GetPendingDelegatorIterator(arg0 ids.ID, arg1 ids.NodeID) (StakerIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingDelegatorIterator", arg0, arg1)
	ret0, _ := ret[0].(StakerIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingDelegatorIterator indicates an expected call of GetPendingDelegatorIterator.
func (mr *MockStateMockRecorder) GetPendingDelegatorIterator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingDelegatorIterator", reflect.TypeOf((*MockState)(nil).GetPendingDelegatorIterator), arg0, arg1)
}

// GetPendingStakerIterator mocks base method.
func (m *MockState) GetPendingStakerIterator() (StakerIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingStakerIterator")
	ret0, _ := ret[0].(StakerIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingStakerIterator indicates an expected call of GetPendingStakerIterator.
func (mr *MockStateMockRecorder) GetPendingStakerIterator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingStakerIterator", reflect.TypeOf((*MockState)(nil).GetPendingStakerIterator))
}

// GetPendingValidator mocks base method.
func (m *MockState) GetPendingValidator(arg0 ids.ID, arg1 ids.NodeID) (*Staker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPendingValidator", arg0, arg1)
	ret0, _ := ret[0].(*Staker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPendingValidator indicates an expected call of GetPendingValidator.
func (mr *MockStateMockRecorder) GetPendingValidator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPendingValidator", reflect.TypeOf((*MockState)(nil).GetPendingValidator), arg0, arg1)
}

// GetRewardUTXOs mocks base method.
func (m *MockState) GetRewardUTXOs(arg0 ids.ID) ([]*avax.UTXO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRewardUTXOs", arg0)
	ret0, _ := ret[0].([]*avax.UTXO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRewardUTXOs indicates an expected call of GetRewardUTXOs.
func (mr *MockStateMockRecorder) GetRewardUTXOs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRewardUTXOs", reflect.TypeOf((*MockState)(nil).GetRewardUTXOs), arg0)
}

// GetStartTime mocks base method.
func (m *MockState) GetStartTime(arg0 ids.NodeID) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStartTime", arg0)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStartTime indicates an expected call of GetStartTime.
func (mr *MockStateMockRecorder) GetStartTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStartTime", reflect.TypeOf((*MockState)(nil).GetStartTime), arg0)
}

// GetStatelessBlock mocks base method.
func (m *MockState) GetStatelessBlock(arg0 ids.ID) (blocks.Block, choices.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatelessBlock", arg0)
	ret0, _ := ret[0].(blocks.Block)
	ret1, _ := ret[1].(choices.Status)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStatelessBlock indicates an expected call of GetStatelessBlock.
func (mr *MockStateMockRecorder) GetStatelessBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatelessBlock", reflect.TypeOf((*MockState)(nil).GetStatelessBlock), arg0)
}

// GetSubnetTransformation mocks base method.
func (m *MockState) GetSubnetTransformation(arg0 ids.ID) (*txs.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnetTransformation", arg0)
	ret0, _ := ret[0].(*txs.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnetTransformation indicates an expected call of GetSubnetTransformation.
func (mr *MockStateMockRecorder) GetSubnetTransformation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnetTransformation", reflect.TypeOf((*MockState)(nil).GetSubnetTransformation), arg0)
}

// GetSubnets mocks base method.
func (m *MockState) GetSubnets() ([]*txs.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubnets")
	ret0, _ := ret[0].([]*txs.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubnets indicates an expected call of GetSubnets.
func (mr *MockStateMockRecorder) GetSubnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubnets", reflect.TypeOf((*MockState)(nil).GetSubnets))
}

// GetTimestamp mocks base method.
func (m *MockState) GetTimestamp() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimestamp")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimestamp indicates an expected call of GetTimestamp.
func (mr *MockStateMockRecorder) GetTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimestamp", reflect.TypeOf((*MockState)(nil).GetTimestamp))
}

// GetTx mocks base method.
func (m *MockState) GetTx(arg0 ids.ID) (*txs.Tx, status.Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTx", arg0)
	ret0, _ := ret[0].(*txs.Tx)
	ret1, _ := ret[1].(status.Status)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTx indicates an expected call of GetTx.
func (mr *MockStateMockRecorder) GetTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTx", reflect.TypeOf((*MockState)(nil).GetTx), arg0)
}

// GetUTXO mocks base method.
func (m *MockState) GetUTXO(arg0 ids.ID) (*avax.UTXO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUTXO", arg0)
	ret0, _ := ret[0].(*avax.UTXO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUTXO indicates an expected call of GetUTXO.
func (mr *MockStateMockRecorder) GetUTXO(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUTXO", reflect.TypeOf((*MockState)(nil).GetUTXO), arg0)
}

// GetUptime mocks base method.
func (m *MockState) GetUptime(arg0 ids.NodeID) (time.Duration, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUptime", arg0)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(time.Time)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUptime indicates an expected call of GetUptime.
func (mr *MockStateMockRecorder) GetUptime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUptime", reflect.TypeOf((*MockState)(nil).GetUptime), arg0)
}

// GetValidatorWeightDiffs mocks base method.
func (m *MockState) GetValidatorWeightDiffs(arg0 uint64, arg1 ids.ID) (map[ids.NodeID]*ValidatorWeightDiff, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorWeightDiffs", arg0, arg1)
	ret0, _ := ret[0].(map[ids.NodeID]*ValidatorWeightDiff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorWeightDiffs indicates an expected call of GetValidatorWeightDiffs.
func (mr *MockStateMockRecorder) GetValidatorWeightDiffs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorWeightDiffs", reflect.TypeOf((*MockState)(nil).GetValidatorWeightDiffs), arg0, arg1)
}

// LockedUTXOs mocks base method.
func (m *MockState) LockedUTXOs(arg0 ids.Set, arg1 ids.ShortSet, arg2 locked.State) ([]*avax.UTXO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockedUTXOs", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*avax.UTXO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LockedUTXOs indicates an expected call of LockedUTXOs.
func (mr *MockStateMockRecorder) LockedUTXOs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockedUTXOs", reflect.TypeOf((*MockState)(nil).LockedUTXOs), arg0, arg1, arg2)
}

// PutCurrentDelegator mocks base method.
func (m *MockState) PutCurrentDelegator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutCurrentDelegator", arg0)
}

// PutCurrentDelegator indicates an expected call of PutCurrentDelegator.
func (mr *MockStateMockRecorder) PutCurrentDelegator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCurrentDelegator", reflect.TypeOf((*MockState)(nil).PutCurrentDelegator), arg0)
}

// PutCurrentValidator mocks base method.
func (m *MockState) PutCurrentValidator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutCurrentValidator", arg0)
}

// PutCurrentValidator indicates an expected call of PutCurrentValidator.
func (mr *MockStateMockRecorder) PutCurrentValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutCurrentValidator", reflect.TypeOf((*MockState)(nil).PutCurrentValidator), arg0)
}

// PutPendingDelegator mocks base method.
func (m *MockState) PutPendingDelegator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutPendingDelegator", arg0)
}

// PutPendingDelegator indicates an expected call of PutPendingDelegator.
func (mr *MockStateMockRecorder) PutPendingDelegator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPendingDelegator", reflect.TypeOf((*MockState)(nil).PutPendingDelegator), arg0)
}

// PutPendingValidator mocks base method.
func (m *MockState) PutPendingValidator(arg0 *Staker) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PutPendingValidator", arg0)
}

// PutPendingValidator indicates an expected call of PutPendingValidator.
func (mr *MockStateMockRecorder) PutPendingValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPendingValidator", reflect.TypeOf((*MockState)(nil).PutPendingValidator), arg0)
}

// SetAddressStates mocks base method.
func (m *MockState) SetAddressStates(arg0 ids.ShortID, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAddressStates", arg0, arg1)
}

// SetAddressStates indicates an expected call of SetAddressStates.
func (mr *MockStateMockRecorder) SetAddressStates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAddressStates", reflect.TypeOf((*MockState)(nil).SetAddressStates), arg0, arg1)
}

// SetCurrentSupply mocks base method.
func (m *MockState) SetCurrentSupply(arg0 ids.ID, arg1 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCurrentSupply", arg0, arg1)
}

// SetCurrentSupply indicates an expected call of SetCurrentSupply.
func (mr *MockStateMockRecorder) SetCurrentSupply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCurrentSupply", reflect.TypeOf((*MockState)(nil).SetCurrentSupply), arg0, arg1)
}

// SetHeight mocks base method.
func (m *MockState) SetHeight(arg0 uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHeight", arg0)
}

// SetHeight indicates an expected call of SetHeight.
func (mr *MockStateMockRecorder) SetHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeight", reflect.TypeOf((*MockState)(nil).SetHeight), arg0)
}

// SetLastAccepted mocks base method.
func (m *MockState) SetLastAccepted(arg0 ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLastAccepted", arg0)
}

// SetLastAccepted indicates an expected call of SetLastAccepted.
func (mr *MockStateMockRecorder) SetLastAccepted(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastAccepted", reflect.TypeOf((*MockState)(nil).SetLastAccepted), arg0)
}

// SetTimestamp mocks base method.
func (m *MockState) SetTimestamp(arg0 time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTimestamp", arg0)
}

// SetTimestamp indicates an expected call of SetTimestamp.
func (mr *MockStateMockRecorder) SetTimestamp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimestamp", reflect.TypeOf((*MockState)(nil).SetTimestamp), arg0)
}

// SetUptime mocks base method.
func (m *MockState) SetUptime(arg0 ids.NodeID, arg1 time.Duration, arg2 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUptime", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUptime indicates an expected call of SetUptime.
func (mr *MockStateMockRecorder) SetUptime(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUptime", reflect.TypeOf((*MockState)(nil).SetUptime), arg0, arg1, arg2)
}

// UTXOIDs mocks base method.
func (m *MockState) UTXOIDs(arg0 []byte, arg1 ids.ID, arg2 int) ([]ids.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UTXOIDs", arg0, arg1, arg2)
	ret0, _ := ret[0].([]ids.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UTXOIDs indicates an expected call of UTXOIDs.
func (mr *MockStateMockRecorder) UTXOIDs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UTXOIDs", reflect.TypeOf((*MockState)(nil).UTXOIDs), arg0, arg1, arg2)
}

// ValidatorSet mocks base method.
func (m *MockState) ValidatorSet(arg0 ids.ID) (validators.Set, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorSet", arg0)
	ret0, _ := ret[0].(validators.Set)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorSet indicates an expected call of ValidatorSet.
func (mr *MockStateMockRecorder) ValidatorSet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorSet", reflect.TypeOf((*MockState)(nil).ValidatorSet), arg0)
}
