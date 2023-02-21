// Copyright (C) 2022, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/wrappers"

	"github.com/ava-labs/avalanchego/vms/platformvm/stakeable"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/locked"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

const (
	defaultCaminoValidatorWeight = 2 * units.KiloAvax
	defaultMinStakingDuration    = 24 * time.Hour
	defaultTxFee                 = uint64(100)
	testNetworkID                = 10
)

var (
	xChainID    = ids.Empty.Prefix(0)
	cChainID    = ids.Empty.Prefix(1)
	avaxAssetID = ids.ID{'y', 'e', 'e', 't'}
)

var (
	caminoPreFundedKeys      = crypto.BuildTestKeys()
	defaultGenesisTime       = time.Date(1997, 1, 1, 0, 0, 0, 0, time.UTC)
	defaultValidateStartTime = defaultGenesisTime
	defaultValidateEndTime   = defaultValidateStartTime.Add(10 * defaultMinStakingDuration)
)

type snLookup struct {
	chainsToSubnet map[ids.ID]ids.ID
}

func (sn *snLookup) SubnetID(chainID ids.ID) (ids.ID, error) {
	subnetID, ok := sn.chainsToSubnet[chainID]
	if !ok {
		return ids.ID{}, errors.New("missing subnet associated with requested chainID")
	}
	return subnetID, nil
}

func defaultContext() *snow.Context {
	ctx := snow.DefaultContextTest()
	ctx.NetworkID = testNetworkID
	ctx.XChainID = xChainID
	ctx.CChainID = cChainID
	ctx.AVAXAssetID = avaxAssetID
	aliaser := ids.NewAliaser()

	errs := wrappers.Errs{}
	errs.Add(
		aliaser.Alias(constants.PlatformChainID, "P"),
		aliaser.Alias(constants.PlatformChainID, constants.PlatformChainID.String()),
		aliaser.Alias(xChainID, "X"),
		aliaser.Alias(xChainID, xChainID.String()),
		aliaser.Alias(cChainID, "C"),
		aliaser.Alias(cChainID, cChainID.String()),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
	ctx.BCLookup = aliaser

	ctx.SNLookup = &snLookup{
		chainsToSubnet: map[ids.ID]ids.ID{
			constants.PlatformChainID: constants.PrimaryNetworkID,
			xChainID:                  constants.PrimaryNetworkID,
			cChainID:                  constants.PrimaryNetworkID,
		},
	}
	return ctx
}

func generateTestOut(assetID ids.ID, amount uint64, outputOwners secp256k1fx.OutputOwners, depositTxID, bondTxID ids.ID) *avax.TransferableOutput {
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
	return &avax.TransferableOutput{
		Asset: avax.Asset{ID: assetID},
		Out:   out,
	}
}

func generateTestStakeableOut(assetID ids.ID, amount, locktime uint64, outputOwners secp256k1fx.OutputOwners) *avax.TransferableOutput {
	return &avax.TransferableOutput{
		Asset: avax.Asset{ID: assetID},
		Out: &stakeable.LockOut{
			Locktime: locktime,
			TransferableOut: &secp256k1fx.TransferOutput{
				Amt:          amount,
				OutputOwners: outputOwners,
			},
		},
	}
}

func generateTestIn(assetID ids.ID, amount uint64, depositTxID, bondTxID ids.ID, sigIndices []uint32) *avax.TransferableInput {
	var in avax.TransferableIn = &secp256k1fx.TransferInput{
		Amt: amount,
		Input: secp256k1fx.Input{
			SigIndices: sigIndices,
		},
	}
	if depositTxID != ids.Empty || bondTxID != ids.Empty {
		in = &locked.In{
			IDs: locked.IDs{
				DepositTxID: depositTxID,
				BondTxID:    bondTxID,
			},
			TransferableIn: in,
		}
	}
	return &avax.TransferableInput{
		UTXOID: avax.UTXOID{TxID: ids.GenerateTestID(), OutputIndex: 0},
		Asset:  avax.Asset{ID: assetID},
		In:     in,
	}
}
