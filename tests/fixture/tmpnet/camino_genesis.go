// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"time"

	"github.com/ava-labs/avalanchego/genesis"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/coreth/core"
	"github.com/ava-labs/coreth/plugin/evm"
)

// Create a genesis struct valid for bootstrapping a test
// network. Note that many of the genesis fields (e.g. reward
// addresses) are randomly generated or hard-coded.
func NewCaminoTestGenesis(
	networkID uint32,
	nodes []*Node,
	keysToFund []*secp256k1.PrivateKey,
) (*genesis.UnparsedConfig, error) {
	// Validate inputs
	switch networkID {
	case constants.TestnetID, constants.MainnetID, constants.LocalID:
		return nil, errInvalidNetworkIDForGenesis
	}
	if len(nodes) == 0 {
		return nil, errMissingStakersForGenesis
	}
	if len(keysToFund) == 0 {
		return nil, errNoKeysForGenesis
	}

	now := time.Now()

	// overwrite randomly generated node IDs with hard-coded ones (used for e2e testing)
	nodeID1, err := ids.NodeIDFromString("NodeID-AK7sPBsZM9rQwse23aLhEEBPHZD5gkLrL")
	if err != nil {
		return nil, err
	}
	nodes[0].NodeID = nodeID1
	nodeID6, err := ids.NodeIDFromString("NodeID-FHseEbTVS7U3odWfjgZYyygsv5gWCqVdk")
	if err != nil {
		return nil, err
	}
	nodes[5].NodeID = nodeID6
	nodeID7, err := ids.NodeIDFromString("NodeID-AAFgkP7AVeQjmv4MSi2DaQbobg3wpZbFp")
	if err != nil {
		return nil, err
	}
	nodes[6].NodeID = nodeID7

	config := &genesis.UnparsedConfig{
		NetworkID: networkID,
		Camino: &genesis.UnparsedCamino{
			VerifyNodeSignature: true,
			LockModeBondDeposit: true,
			InitialAdmin:        "X-kopernikus1g65uqn6t77p656w64023nh8nd9updzmxh8ttv3",
			DepositOffers: []genesis.UnparsedDepositOffer{
				{
					InterestRateNominator:   110000,
					StartOffset:             0,
					EndOffset:               63075600,
					MinAmount:               1,
					MinDuration:             60,
					MaxDuration:             31536000,
					UnlockPeriodDuration:    0,
					NoRewardsPeriodDuration: 0,
					Memo:                    "depositOffer test#1 - pickable after network creation and not in genesis as duration is not choosable per allocation",
					Flags: genesis.UnparsedDepositOfferFlags{
						Locked: false,
					},
				},
				{
					InterestRateNominator:   210000,
					StartOffset:             0,
					EndOffset:               63075600,
					MinAmount:               1,
					MinDuration:             86400,
					MaxDuration:             31536000,
					UnlockPeriodDuration:    43200,
					NoRewardsPeriodDuration: 43200,
					Memo:                    "depositOffer test#2 - pickable after network creation and not in genesis as duration is not choosable per allocation!",
					Flags: genesis.UnparsedDepositOfferFlags{
						Locked: false,
					},
				},
				{
					InterestRateNominator:   80000,
					StartOffset:             0,
					EndOffset:               112795200,
					MinAmount:               1,
					MinDuration:             110376000,
					MaxDuration:             110376000,
					UnlockPeriodDuration:    31536000,
					NoRewardsPeriodDuration: 15768000,
					Memo:                    "lockedpresale3y",
					Flags: genesis.UnparsedDepositOfferFlags{
						Locked: true,
					},
				},
				{
					InterestRateNominator:   0.1 * 1_000_000 * (365 * 24 * 60 * 60),
					StartOffset:             0,
					EndOffset:               112795200,
					MinAmount:               100,
					MinDuration:             60,
					MaxDuration:             60,
					UnlockPeriodDuration:    20,
					NoRewardsPeriodDuration: 10,
					Memo:                    "presale1min",
					Flags: genesis.UnparsedDepositOfferFlags{
						Locked: false,
					},
				},
			},
			Allocations: []genesis.UnparsedCaminoAllocation{
				{
					ETHAddr:  "0x0000000000000000000000000000000000000000",
					AVAXAddr: "X-kopernikus1g65uqn6t77p656w64023nh8nd9updzmxh8ttv3",
					XAmount:  1000000000000,
					AddressStates: genesis.AddressStates{
						ConsortiumMember: true,
						KYCVerified:      true,
					},
					PlatformAllocations: []genesis.UnparsedPlatformAllocation{
						{
							Amount:            2000000000000,
							NodeID:            nodes[0].NodeID.String(),
							ValidatorDuration: 31536000,
							TimestampOffset:   0,
						},
						{
							Amount: 1000000000000,
						},
						{
							Amount: 4000000000000,
						},
					},
				},
				{
					ETHAddr:  "0x0000000000000000000000000000000000000000",
					AVAXAddr: "X-kopernikus18jma8ppw3nhx5r4ap8clazz0dps7rv5uuvjh68",
					XAmount:  1000000000000,
					AddressStates: genesis.AddressStates{
						ConsortiumMember: true,
						KYCVerified:      true,
					},
					PlatformAllocations: []genesis.UnparsedPlatformAllocation{
						{
							Amount:            2000000000000,
							NodeID:            nodes[1].NodeID.String(),
							ValidatorDuration: 31536000,
							TimestampOffset:   0,
						},
						{
							Amount: 1000000000000,
						},
					},
				},
				{
					ETHAddr:  "0x0000000000000000000000000000000000000000",
					AVAXAddr: "X-kopernikus13kyf72ftu4l77kss7xm0kshm0au29s48zjaygq",
					XAmount:  1000000000000,
					AddressStates: genesis.AddressStates{
						ConsortiumMember: true,
						KYCVerified:      true,
					},
					PlatformAllocations: []genesis.UnparsedPlatformAllocation{
						{
							Amount:            2000000000000,
							NodeID:            nodes[2].NodeID.String(),
							ValidatorDuration: 31536000,
							TimestampOffset:   0,
						},
					},
				},
				{
					ETHAddr:  "0x0000000000000000000000000000000000000000",
					AVAXAddr: "X-kopernikus1zy075lddftstzpwzvt627mvc0tep0vyk7y9v4l",
					AddressStates: genesis.AddressStates{
						ConsortiumMember: true,
						KYCVerified:      true,
					},
					PlatformAllocations: []genesis.UnparsedPlatformAllocation{
						{
							Amount:            2000000000000,
							NodeID:            nodes[3].NodeID.String(),
							ValidatorDuration: 31536000,
							TimestampOffset:   0,
						},
					},
				},
				{
					ETHAddr:  "0x0000000000000000000000000000000000000000",
					AVAXAddr: "X-kopernikus1lx58kettrnt2kyr38adyrrmxt5x57u4vg4cfky",
					XAmount:  1000000000000,
					AddressStates: genesis.AddressStates{
						ConsortiumMember: true,
						KYCVerified:      true,
					},
					PlatformAllocations: []genesis.UnparsedPlatformAllocation{
						{
							Amount:            2000000000000,
							NodeID:            nodes[4].NodeID.String(),
							ValidatorDuration: 31536000,
							TimestampOffset:   0,
						},
						{
							Amount: 1000000000000,
						},
					},
				},
				{
					ETHAddr:  "0x0000000000000000000000000000000000000000",
					AVAXAddr: "X-kopernikus1fwrv3kj5jqntuucw67lzgu9a9tkqyczxgcvpst",
					XAmount:  1000000000000,
					AddressStates: genesis.AddressStates{
						ConsortiumMember: true,
						KYCVerified:      true,
					},
					PlatformAllocations: []genesis.UnparsedPlatformAllocation{
						{
							Amount: 200000000000000000,
						},
					},
				},
				{
					ETHAddr:  "0x0000000000000000000000000000000000000000",
					AVAXAddr: "X-kopernikus1s93gzmzuvv7gz8q4l83xccrdchh8mtm3xm5s2g",
					AddressStates: genesis.AddressStates{
						ConsortiumMember: true,
						KYCVerified:      true,
					},
					PlatformAllocations: []genesis.UnparsedPlatformAllocation{
						{
							Amount: 4000000000000,
						},
					},
				},
				{
					ETHAddr:  "0x0000000000000000000000000000000000000000",
					AVAXAddr: "X-kopernikus1jla8ty5c9ud6lsj8s4re2dvzvfxpzrxdcrd8q7",
					XAmount:  1000000000000,
					AddressStates: genesis.AddressStates{
						ConsortiumMember: true,
						KYCVerified:      true,
					},
					PlatformAllocations: []genesis.UnparsedPlatformAllocation{
						{
							Amount: 200000000000000000,
						},
					},
				},
			},
			InitialMultisigAddresses: []genesis.UnparsedMultisigAlias{
				{
					Alias: "X-kopernikus1fwrv3kj5jqntuucw67lzgu9a9tkqyczxgcvpst",
					Addresses: []string{
						"X-kopernikus1jla8ty5c9ud6lsj8s4re2dvzvfxpzrxdcrd8q7",
						"X-kopernikus15hscuhlg5tkv4wwrujqgarne3tau83wrpp2d0d",
					},
					Threshold: 1,
				},
			},
		},
		StartTime: uint64(now.Unix()),
		Message:   "hello camino!",
	}

	// Ensure pre-funded keys have arbitrary large balances on both chains to support testing
	xChainBalances := make(XChainBalanceMap, len(keysToFund))
	cChainBalances := make(core.GenesisAlloc, len(keysToFund))
	for _, key := range keysToFund {
		xChainBalances[key.Address()] = defaultFundedKeyXChainAmount
		cChainBalances[evm.GetEthAddress(key)] = core.GenesisAccount{
			Balance: defaultFundedKeyCChainAmount,
		}
	}

	// TODO what about dUpgradeBlockTimestamp and cortina
	config.CChainGenesis = "{\"alloc\": {\"0100000000000000000000000000000000000000\": {\"balance\": \"0x0\",\"code\": \"0x7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80631e010439146042578063b6510bb314606e575b600080fd5b605c60048036036020811015605657600080fd5b503560b1565b60408051918252519081900360200190f35b818015607957600080fd5b5060af60048036036080811015608e57600080fd5b506001600160a01b03813516906020810135906040810135906060013560b6565b005b30cd90565b836001600160a01b031681836108fc8690811502906040516000604051808303818888878c8acf9550505050505015801560f4573d6000803e3d6000fd5b505050505056fea26469706673582212201eebce970fe3f5cb96bf8ac6ba5f5c133fc2908ae3dcd51082cfee8f583429d064736f6c634300060a0033\"  },\"1f0e5c64afdf53175f78846f7125776e76fa8f34\": {\"balance\": \"0x295BE96E64066972000000\"  },\"305cea207112c0561033133f816d7a2233699f06\": {\"balance\": \"0x295BE96E64066972000000\"  },\"8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC\": {\"balance\": \"0x295BE96E64066972000000\"  }},\"coinbase\": \"0x0000000000000000000000000000000000000000\",\"config\": {\"apricotPhase1BlockTimestamp\": 0,\"apricotPhase2BlockTimestamp\": 0,\"apricotPhase3BlockTimestamp\": 0,\"apricotPhase4BlockTimestamp\": 0,\"apricotPhase5BlockTimestamp\": 0,\"apricotPhase6BlockTimestamp\": 0,\"apricotPhasePost6BlockTimestamp\": 0,\"apricotPhasePre6BlockTimestamp\": 0,\"banffBlockTimestamp\": 0,\"byzantiumBlock\": 0,\"chainId\": 503,\"constantinopleBlock\": 0,\"daoForkBlock\": 0,\"daoForkSupport\": true,\"eip150Block\": 0,\"eip150Hash\": \"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\": 0,\"eip158Block\": 0,\"homesteadBlock\": 0,\"istanbulBlock\": 0,\"muirGlacierBlock\": 0,\"petersburgBlock\": 0,\"sunrisePhase0BlockTimestamp\": 0},\"difficulty\": \"0x0\",\"extraData\": \"0x00\",\"feeRewardExportMinAmount\": \"0x2710\",\"feeRewardExportMinTimeInterval\": \"0x3C\",\"gasLimit\": \"0x5f5e100\",\"gasUsed\": \"0x0\",\"initialAdmin\": \"0x1f0e5c64afdf53175f78846f7125776e76fa8f34\",\"mixHash\": \"0x0000000000000000000000000000000000000000000000000000000000000000\",\"nonce\": \"0x0\",\"number\": \"0x0\",\"parentHash\": \"0x0000000000000000000000000000000000000000000000000000000000000000\",\"timestamp\": \"0x0\"}"
	return config, nil
}
