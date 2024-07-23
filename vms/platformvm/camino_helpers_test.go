// Copyright (C) 2022-2023, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/api/keystore"
	"github.com/ava-labs/avalanchego/chains"
	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/database/memdb"
	"github.com/ava-labs/avalanchego/database/prefixdb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/uptime"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/formatting"
	"github.com/ava-labs/avalanchego/utils/formatting/address"
	"github.com/ava-labs/avalanchego/utils/json"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/nodeid"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	as "github.com/ava-labs/avalanchego/vms/platformvm/addrstate"
	"github.com/ava-labs/avalanchego/vms/platformvm/api"
	"github.com/ava-labs/avalanchego/vms/platformvm/caminoconfig"
	"github.com/ava-labs/avalanchego/vms/platformvm/config"
	"github.com/ava-labs/avalanchego/vms/platformvm/genesis"
	"github.com/ava-labs/avalanchego/vms/platformvm/locked"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

const (
	defaultCaminoValidatorWeight = 2 * units.KiloAvax
)

var (
	localStakingPath          = "../../staking/local/"
	caminoPreFundedKeys       = secp256k1.TestKeys()
	_, caminoPreFundedNodeIDs = nodeid.LoadLocalCaminoNodeKeysAndIDs(localStakingPath)
	defaultStartTime          = banffForkTime.Add(time.Second)

	testAddressID ids.ShortID
)

func init() {
	_, _, testAddressBytes, err := address.Parse(testAddress)
	if err != nil {
		panic(err)
	}
	testAddressID, err = ids.ToShortID(testAddressBytes)
	if err != nil {
		panic(err)
	}
}

func defaultCaminoService(t *testing.T, camino api.Camino, utxos []api.UTXO) *CaminoService {
	vm := newCaminoVM(t, camino, utxos, nil)

	vm.ctx.Lock.Lock()
	defer vm.ctx.Lock.Unlock()
	ks := keystore.New(logging.NoLog{}, memdb.New())
	require.NoError(t, ks.CreateUser(testUsername, testPassword))
	vm.ctx.Keystore = ks.NewBlockchainKeyStore(vm.ctx.ChainID)
	return &CaminoService{
		Service: Service{
			vm:          vm,
			addrManager: avax.NewAddressManager(vm.ctx),
		},
	}
}

func newCaminoVM(t *testing.T, genesisConfig api.Camino, genesisUTXOs []api.UTXO, startTime *time.Time) *VM {
	require := require.New(t)

	vm := &VM{Config: defaultCaminoConfig()}

	db := memdb.New()
	chainDB := prefixdb.New([]byte{0}, db)
	atomicDB := prefixdb.New([]byte{1}, db)

	if startTime == nil {
		startTime = &defaultStartTime
	}
	vm.clock.Set(*startTime)
	msgChan := make(chan common.Message, 1)
	ctx := defaultContext(t)

	m := atomic.NewMemory(atomicDB)
	msm := &mutableSharedMemory{
		SharedMemory: m.NewSharedMemory(ctx.ChainID),
	}
	ctx.SharedMemory = msm

	ctx.Lock.Lock()
	defer ctx.Lock.Unlock()
	_, genesisBytes := newCaminoGenesisWithUTXOs(t, genesisConfig, genesisUTXOs, startTime)
	// _, genesisBytes := defaultGenesis(t)
	appSender := &common.SenderTest{}
	appSender.CantSendAppGossip = true
	appSender.SendAppGossipF = func(context.Context, []byte) error {
		return nil
	}

	require.NoError(vm.Initialize(
		context.Background(),
		ctx,
		chainDB,
		genesisBytes,
		nil,
		nil,
		msgChan,
		nil,
		appSender,
	))

	require.NoError(vm.SetState(context.Background(), snow.NormalOp))

	// Create a subnet and store it in testSubnet1
	// Note: following Banff activation, block acceptance will move
	// chain time ahead
	testSubnet1, err := vm.txBuilder.NewCreateSubnetTx(
		2, // threshold; 2 sigs from control keys needed to add validator to this subnet
		[]ids.ShortID{ // control keys
			caminoPreFundedKeys[0].PublicKey().Address(),
			caminoPreFundedKeys[1].PublicKey().Address(),
			caminoPreFundedKeys[2].PublicKey().Address(),
		},
		[]*secp256k1.PrivateKey{caminoPreFundedKeys[0]},
		caminoPreFundedKeys[0].PublicKey().Address(),
	)
	require.NoError(err)
	require.NoError(vm.Builder.AddUnverifiedTx(testSubnet1))
	blk, err := vm.Builder.BuildBlock(context.Background())
	require.NoError(err)
	require.NoError(blk.Verify(context.Background()))
	require.NoError(blk.Accept(context.Background()))
	require.NoError(vm.SetPreference(context.Background(), vm.manager.LastAccepted()))

	return vm
	// return vm, baseDBManager.Current().Database, msm
}

func defaultCaminoConfig() config.Config {
	return config.Config{
		Chains:                 chains.TestManager,
		UptimeLockedCalculator: uptime.NewLockedCalculator(),
		SybilProtectionEnabled: true,
		Validators:             validators.NewManager(),
		TxFee:                  defaultTxFee,
		CreateSubnetTxFee:      100 * defaultTxFee,
		TransformSubnetTxFee:   100 * defaultTxFee,
		CreateBlockchainTxFee:  100 * defaultTxFee,
		MinValidatorStake:      defaultCaminoValidatorWeight,
		MaxValidatorStake:      defaultCaminoValidatorWeight,
		MinDelegatorStake:      1 * units.MilliAvax,
		MinStakeDuration:       defaultMinStakingDuration,
		MaxStakeDuration:       defaultMaxStakingDuration,
		RewardConfig:           defaultRewardConfig,
		ApricotPhase3Time:      defaultValidateEndTime,
		ApricotPhase5Time:      defaultValidateEndTime,
		BanffTime:              banffForkTime,
		CaminoConfig: caminoconfig.Config{
			DACProposalBondAmount: 100 * units.Avax,
		},
	}
}

// Returns:
// 1) The genesis state
// 2) The byte representation of the default genesis for tests
func newCaminoGenesisWithUTXOs(t *testing.T, caminoGenesisConfig api.Camino, genesisUTXOs []api.UTXO, starttime *time.Time) (*api.BuildGenesisArgs, []byte) {
	require := require.New(t)

	if starttime == nil {
		starttime = &defaultValidateStartTime
	}
	caminoGenesisConfig.UTXODeposits = make([]api.UTXODeposit, len(genesisUTXOs))
	caminoGenesisConfig.ValidatorDeposits = make([][]api.UTXODeposit, len(caminoPreFundedKeys))
	caminoGenesisConfig.ValidatorConsortiumMembers = make([]ids.ShortID, len(caminoPreFundedKeys))

	genesisValidators := make([]api.PermissionlessValidator, len(caminoPreFundedKeys))
	for i, key := range caminoPreFundedKeys {
		addr, err := address.FormatBech32(constants.UnitTestHRP, key.PublicKey().Address().Bytes())
		require.NoError(err)
		genesisValidators[i] = api.PermissionlessValidator{
			Staker: api.Staker{
				StartTime: json.Uint64(starttime.Unix()),
				EndTime:   json.Uint64(starttime.Add(10 * defaultMinStakingDuration).Unix()),
				NodeID:    caminoPreFundedNodeIDs[i],
			},
			RewardOwner: &api.Owner{
				Threshold: 1,
				Addresses: []string{addr},
			},
			Staked: []api.UTXO{{
				Amount:  json.Uint64(defaultCaminoValidatorWeight),
				Address: addr,
			}},
		}
		caminoGenesisConfig.ValidatorDeposits[i] = make([]api.UTXODeposit, 1)
		caminoGenesisConfig.ValidatorConsortiumMembers[i] = key.Address()
		caminoGenesisConfig.AddressStates = append(caminoGenesisConfig.AddressStates, genesis.AddressState{
			Address: key.Address(),
			State:   as.AddressStateConsortium,
		})
	}

	buildGenesisArgs := api.BuildGenesisArgs{
		Encoding:      formatting.Hex,
		NetworkID:     json.Uint32(constants.UnitTestID),
		AvaxAssetID:   avaxAssetID,
		UTXOs:         genesisUTXOs,
		Validators:    genesisValidators,
		Chains:        nil,
		Time:          json.Uint64(defaultGenesisTime.Unix()),
		InitialSupply: json.Uint64(360 * units.MegaAvax),
		Camino:        &caminoGenesisConfig,
	}

	buildGenesisResponse := api.BuildGenesisReply{}
	platformvmSS := api.StaticService{}
	require.NoError(platformvmSS.BuildGenesis(nil, &buildGenesisArgs, &buildGenesisResponse))

	genesisBytes, err := formatting.Decode(buildGenesisResponse.Encoding, buildGenesisResponse.Bytes)
	require.NoError(err)

	return &buildGenesisArgs, genesisBytes
}

func generateKeyAndOwner(t *testing.T) (*secp256k1.PrivateKey, ids.ShortID, secp256k1fx.OutputOwners) {
	t.Helper()
	key, err := secp256k1.NewPrivateKey()
	require.NoError(t, err)
	addr := key.Address()
	return key, addr, secp256k1fx.OutputOwners{
		Locktime:  0,
		Threshold: 1,
		Addrs:     []ids.ShortID{addr},
	}
}

func stopVM(t *testing.T, vm *VM, doLock bool) {
	t.Helper()
	if doLock {
		vm.ctx.Lock.Lock()
	}
	require.NoError(t, vm.Shutdown(context.TODO()))
	vm.ctx.Lock.Unlock()
}

func generateTestUTXO(txID ids.ID, assetID ids.ID, amount uint64, outputOwners secp256k1fx.OutputOwners, depositTxID, bondTxID ids.ID) *avax.UTXO {
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
	return &avax.UTXO{
		UTXOID: avax.UTXOID{TxID: txID},
		Asset:  avax.Asset{ID: assetID},
		Out:    out,
	}
}
