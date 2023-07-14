package test

import (
	"os"
	"runtime"
	"strconv"
	"testing"

	wasp_wallet_sdk "github.com/iotaledger/wasp-wallet-sdk"
	"github.com/iotaledger/wasp-wallet-sdk/types"

	"github.com/stretchr/testify/require"
)

func FromEnv(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}
	return value
}

var (
	ShimmerNetworkAPI     = FromEnv("TEST_NETWORK_API", "https://api.shimmer.network")
	UseLedgerSimulator, _ = strconv.ParseBool(FromEnv("TEST_USE_LEDGER_SIMULATOR", "false"))
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
		return wd + "/../../iota-sdk/target/release/libiota_sdk_native.so"

	case "windows":
		return wd + "/../../iota-sdk/target/debug/libiota_sdk_go.dll"

	default:
		return ""
	}
}

var sdk *wasp_wallet_sdk.IOTASDK

func GetOrInitTest(t *testing.T) *wasp_wallet_sdk.IOTASDK {
	var err error
	if sdk != nil {
		return sdk
	}

	sdk, err = wasp_wallet_sdk.NewIotaSDK(getIOTASDKLibraryPath())
	require.NoError(t, err)

	success, err := sdk.InitLogger(types.ILoggerConfig{
		LevelFilter: types.LevelFilterTrace,
	})
	require.NoError(t, err)
	require.True(t, success)

	return sdk
}
