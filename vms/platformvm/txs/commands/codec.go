package commands

import (
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/codec/linearcodec"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/utils/wrappers"
)

const (
	Version        = uint16(0)
	maxMessageSize = 1 * units.MiB
)

var (
	Codec           codec.Manager
	CrossChainCodec codec.Manager
)

func init() {
	Codec = codec.NewManager(maxMessageSize)
	c := linearcodec.NewDefault()

	errs := wrappers.Errs{}
	errs.Add(
		// commands
		c.RegisterType(CommandSetBaseFee{}),
		c.RegisterType(CommandSetKYC{}),
	)
	c.SkipRegistrations(64)

	errs.Add(
		Codec.RegisterCodec(Version, c),
	)

	if errs.Errored() {
		panic(errs.Err)
	}

	// CrossChainCodec = codec.NewManager(maxMessageSize)
	// ccc := linearcodec.NewDefault()

	// errs = wrappers.Errs{}
	// errs.Add(
	// 	// CrossChainRequest Types
	// 	ccc.RegisterType(EthCallRequest{}),
	// 	ccc.RegisterType(EthCallResponse{}),

	// 	CrossChainCodec.RegisterCodec(Version, ccc),
	// )

	// if errs.Errored() {
	// 	panic(errs.Err)
	// }
}
