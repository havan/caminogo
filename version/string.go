// Copyright (C) 2022-2023, Chain4Travel AG. All rights reserved.
//
// This file is a derived work, based on ava-labs code whose
// original notices appear below.
//
// It is distributed under the same license conditions as the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********************************************************
// Copyright (C) 2019-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"fmt"
	"runtime"
	"strings"
)

// String is displayed when CLI arg --version is used
var String string

func init() {
	goVersion := runtime.Version()
	goVersionNumber := strings.TrimPrefix(goVersion, "go")
	String = fmt.Sprintf("%s [git %s, %s; database: %s, rpcchainvm %d, go: %s]\n",
		CurrentApp,
		GitVersion,
		GitCommit,
		CurrentDatabase,
		RPCChainVMProtocol,
		goVersionNumber,
	)
}
