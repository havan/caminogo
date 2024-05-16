// Copyright (C) 2022-2023, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"context"
	"fmt"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/formatting"
	"github.com/ava-labs/avalanchego/utils/rpc"
)

type CaminoClient interface {
	// GetConfiguration returns genesis information of the primary network
	GetConfiguration(ctx context.Context, options ...rpc.Option) (*GetConfigurationReply, error)

	// GetMultisigAlias returns the alias definition of the given multisig address
	GetMultisigAlias(ctx context.Context, multisigAddress string, options ...rpc.Option) (*GetMultisigAliasReply, error)

	GetRegisteredShortIDLink(ctx context.Context, addrStr ids.ShortID, options ...rpc.Option) (string, error)
	GetLastAcceptedBlock(ctx context.Context, options ...rpc.Option) ([]byte, error)
}

func (c *client) GetConfiguration(ctx context.Context, options ...rpc.Option) (*GetConfigurationReply, error) {
	res := &GetConfigurationReply{}
	err := c.requester.SendRequest(ctx, "platform.getConfiguration", struct{}{}, res, options...)
	return res, err
}

func (c *client) GetMultisigAlias(ctx context.Context, multisigAddress string, options ...rpc.Option) (*GetMultisigAliasReply, error) {
	res := &GetMultisigAliasReply{}
	err := c.requester.SendRequest(ctx, "platform.getMultisigAlias", &api.JSONAddress{
		Address: multisigAddress,
	}, res, options...)
	return res, err
}

func (c *client) GetRegisteredShortIDLink(ctx context.Context, addrStr ids.ShortID, options ...rpc.Option) (string, error) {
	res := &api.JSONAddress{}
	err := c.requester.SendRequest(ctx, "platform.getRegisteredShortIDLink", &api.JSONAddress{
		Address: addrStr.String(),
	}, res, options...)
	return res.Address, err
}

func (c *client) GetLastAcceptedBlock(ctx context.Context, options ...rpc.Option) ([]byte, error) {
	res := &api.GetBlockResponse{}
	err := c.requester.SendRequest(ctx, "platform.getLastAcceptedBlock", struct{}{}, res, options...)
	if err != nil {
		return nil, err
	}
	blkBytesStr, ok := res.Block.(string)
	if !ok {
		return nil, fmt.Errorf("platform.getLastAcceptedBlock.reply.Block expected []byte, got %T", res.Block)
	}
	return formatting.Decode(formatting.Hex, blkBytesStr)
}
