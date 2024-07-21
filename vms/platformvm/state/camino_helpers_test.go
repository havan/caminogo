// Copyright (C) 2022-2023, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/ava-labs/avalanchego/database/memdb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/config"
	"github.com/ava-labs/avalanchego/vms/platformvm/locked"
	"github.com/ava-labs/avalanchego/vms/platformvm/metrics"
	"github.com/ava-labs/avalanchego/vms/platformvm/reward"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

func generateBaseTx(assetID ids.ID, amount uint64, outputOwners secp256k1fx.OutputOwners, depositTxID, bondTxID ids.ID) *txs.BaseTx {
	var out avax.TransferableOut = &secp256k1fx.TransferOutput{
		Amt:          amount,
		OutputOwners: outputOwners,
	}
	if depositTxID != ids.Empty || bondTxID != ids.Empty {
		out = &locked.Out{
			IDs: locked.IDs{
				DepositTxID: depositTxID,
				BondTxID:    bondTxID,
			},
			TransferableOut: out,
		}
	}

	return &txs.BaseTx{
		BaseTx: avax.BaseTx{
			Outs: []*avax.TransferableOutput{
				{
					Asset: avax.Asset{ID: assetID},
					Out:   out,
				},
			},
		},
	}
}

func newEmptyState(t *testing.T) *state {
	execCfg, _ := config.GetExecutionConfig(nil)
	newState, err := newState(
		memdb.New(),
		metrics.Noop,
		&config.Config{
			Validators: validators.NewManager(),
		},
		execCfg,
		&snow.Context{},
		prometheus.NewRegistry(),
		reward.NewCalculator(reward.Config{
			MaxConsumptionRate: .12 * reward.PercentDenominator,
			MinConsumptionRate: .1 * reward.PercentDenominator,
			MintingPeriod:      365 * 24 * time.Hour,
			SupplyCap:          720 * units.MegaAvax,
		}),
		&utils.Atomic[bool]{},
	)
	require.NoError(t, err)
	require.NotNil(t, newState)
	return newState
}

func newMockStateVersions(c *gomock.Controller, parentStateID ids.ID, parentState Chain) *MockVersions {
	stateVersions := NewMockVersions(c)
	stateVersions.EXPECT().GetState(parentStateID).Return(parentState, true)
	return stateVersions
}
