// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCurrentRPCChainVMCompatible(t *testing.T) {
	compatibleVersions := RPCChainVMProtocolCompatibility[RPCChainVMProtocol]
	for _, ver := range compatibleVersions {
		_ = ver.String()
	}
	require.Contains(t, compatibleVersions, Current)
}
