// Copyright (C) 2023, Chain4Travel AG. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import "github.com/ava-labs/avalanchego/ids"

type CaminoRewardMessage struct {
	Timestamp uint64 `serialize:"true" json:"timestamp"`
}

type CaminoCommandMessage struct {
	CommandTxID ids.ID `serialize:"true" json:"commandTxID"`
}
