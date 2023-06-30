package test

import (
	"os"
	"runtime"
	"testing"

	"github.com/iotaledger/wasp_wallet_sdk"
	"github.com/iotaledger/wasp_wallet_sdk/types"

	"github.com/stretchr/testify/require"
)

const (
	ShimmerNetworkAPI  = "https://api.shimmer.network"
	UseLedgerSimulator = true
)

// Mnemonic chosen by fair dice roll.
// guaranteed to be random.
const (
	Mnemonic                  = "saddle dune lake festival gain race cancel fragile amused brush donor outer today unique actress rescue abstract curve tail find catch huge cricket crop"
	SignMessageFromEssenceHex = "0xcf30a3824d6b2d3b25ec63aa97733e4fc4dd99e6d38c97093a0abd21f5e9016c"
)

func getIOTASDKLibraryPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	switch runtime.GOOS {
	case "darwin":
		return wd + "/../../iota-sdk/target/debug/libiota_sdk_go.dylib"

	case "linux":
		return wd + "/../../iota-sdk/target/debug/libiota_sdk_go.so"

	case "windows":
		return wd + "/../../iota-sdk/target/debug/libiota_sdk_go.dll"

	default:
		return ""
	}
}

func InitTest(t *testing.T) *wasp_wallet_sdk.IOTASDK {
	sdk, err := wasp_wallet_sdk.NewIotaSDK(getIOTASDKLibraryPath())
	require.NoError(t, err)

	success, err := sdk.InitLogger(types.ILoggerConfig{
		LevelFilter: types.LevelFilterTrace,
	})
	require.True(t, success)
	require.NoError(t, err)

	return sdk
}
