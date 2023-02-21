// Copyright (C) 2023, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"testing"

	"github.com/ava-labs/avalanchego/utils/nodeid"

	"github.com/ava-labs/avalanchego/vms/platformvm/locked"
	"github.com/ava-labs/avalanchego/vms/platformvm/validator"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/treasury"
	"github.com/stretchr/testify/require"
)

func TestCaminoSerializeTx(t *testing.T) {
	ctx := defaultContext()

	var signers [][]*crypto.PrivateKeySECP256K1R
	signers = append(signers, []*crypto.PrivateKeySECP256K1R{caminoPreFundedKeys[0]})
	outputOwners := secp256k1fx.OutputOwners{
		Locktime:  0,
		Threshold: 1,
		Addrs:     []ids.ShortID{caminoPreFundedKeys[0].PublicKey().Address()},
	}
	_, nodeID := nodeid.GenerateCaminoNodeKeyAndID()

	tests := map[string]struct {
		utx         UnsignedTx
		expectedErr error
	}{
		"AddValidatorTx": {
			utx: &CaminoAddValidatorTx{
				AddValidatorTx: AddValidatorTx{
					BaseTx: BaseTx{BaseTx: avax.BaseTx{
						NetworkID:    ctx.NetworkID,
						BlockchainID: ctx.ChainID,
						Ins: []*avax.TransferableInput{
							generateTestIn(ctx.AVAXAssetID, defaultCaminoValidatorWeight*2, ids.Empty, ids.Empty, []uint32{0}),
						},
						Outs: []*avax.TransferableOutput{
							generateTestOut(ctx.AVAXAssetID, defaultCaminoValidatorWeight-defaultTxFee, outputOwners, ids.Empty, ids.Empty),
							generateTestOut(ctx.AVAXAssetID, defaultCaminoValidatorWeight, outputOwners, ids.Empty, locked.ThisTxID),
						},
					}},
					Validator: validator.Validator{
						NodeID: nodeID,
						Start:  uint64(defaultValidateStartTime.Unix()) + 1,
						End:    uint64(defaultValidateEndTime.Unix()),
						Wght:   defaultCaminoValidatorWeight,
					},
					RewardsOwner: &secp256k1fx.OutputOwners{
						Locktime:  0,
						Threshold: 1,
						Addrs:     []ids.ShortID{ids.ShortEmpty},
					},
				},
			},
		},
		"RewardsImportTx": {
			utx: &RewardsImportTx{BaseTx: BaseTx{BaseTx: avax.BaseTx{
				Ins: []*avax.TransferableInput{
					generateTestIn(ctx.AVAXAssetID, 1, ids.Empty, ids.Empty, []uint32{}),
				},
				Outs: []*avax.TransferableOutput{
					generateTestOut(ctx.AVAXAssetID, 1, *treasury.Owner, ids.Empty, ids.Empty),
				},
			}}},
		},
		"AddressStateTx": {
			utx: &AddressStateTx{
				BaseTx: BaseTx{BaseTx: avax.BaseTx{
					NetworkID:    ctx.NetworkID,
					BlockchainID: ctx.ChainID,
					Ins: []*avax.TransferableInput{
						generateTestIn(ctx.AVAXAssetID, 1, ids.Empty, ids.Empty, []uint32{0}),
					},
					Outs: []*avax.TransferableOutput{
						generateTestOut(ctx.AVAXAssetID, 1, outputOwners, ids.Empty, ids.Empty),
					},
					Memo: []byte{1, 2, 3, 4, 5, 6, 7, 8},
				}},
				Address: preFundedKeys[0].PublicKey().Address(),
				State:   AddressStateRoleAdmin,
				Remove:  false,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			tt.utx.InitCtx(ctx)
			stx, err := NewSigned(tt.utx, Codec, signers)
			require.NoError(t, stx.SyntacticVerify(ctx))
			require.NoError(t, err)

			txBytes, err := Codec.Marshal(Version, stx)
			require.NoError(t, err)

			parsedTx, err := Parse(Codec, txBytes)
			require.NoError(t, err)
			require.Equal(t, tt.utx.Bytes(), parsedTx.Unsigned.Bytes())
		})
	}
}
