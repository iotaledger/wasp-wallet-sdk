//go:build IOTA_SDK_WITH_WALLET

package test

import (
	"testing"

	wasp_wallet_sdk "github.com/iotaledger/wasp-wallet-sdk"
	"github.com/iotaledger/wasp-wallet-sdk/types"

	"github.com/stretchr/testify/require"
)

/**
Wallet tests are without any client configuration and are therefore _offline_
*/

func TestWalletMnemonic(t *testing.T) {
	sdk := GetOrInitTest(t)

	wallet, err := sdk.CreateWallet(types.WalletOptions{
		SecretManager: types.MnemonicSecretManager{
			Mnemonic: Mnemonic,
		},
		ClientOptions: &types.ClientOptions{},
		StoragePath:   "./testdb/mnemonic",
		CoinType:      types.CoinTypeSMR,
	})

	defer wallet.Destroy()
	require.NoError(t, err)
	require.NotNil(t, wallet)

	t.Log(wallet)

	bip32Chain := wasp_wallet_sdk.BuildBip32Chain(types.CoinTypeSMR, 0, false, 0)
	result, err := wallet.SignTransactionEssence(SignMessageFromEssenceHex, bip32Chain)

	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestWalletLedger(t *testing.T) {
	sdk := GetOrInitTest(t)

	wallet, err := sdk.CreateWallet(types.WalletOptions{
		ClientOptions: &types.ClientOptions{},
		SecretManager: types.LedgerNanoSecretManager{
			LedgerNano: UseLedgerSimulator,
		},
		StoragePath: "./testdb/ledger",
		CoinType:    types.CoinTypeSMR,
	})
	defer wallet.Destroy()

	require.NoError(t, err)
	require.NotNil(t, wallet)

	status, err := wallet.GetLedgerStatus()
	require.NoError(t, err)
	require.NotNil(t, status)

	address, err := wallet.GenerateEd25519Address(0, 0, "smr", nil)
	require.NoError(t, err)
	require.NotEmpty(t, address)

	bip32Chain := wasp_wallet_sdk.BuildBip32Chain(types.CoinTypeSMR, 0, false, 0)
	signedEssence, err := wallet.SignTransactionEssence(types.HexEncodedString(SignMessageFromEssenceHex), bip32Chain)
	require.NoError(t, err)
	require.NotEmpty(t, signedEssence)
}

func TestWalletStronghold(t *testing.T) {
	sdk := GetOrInitTest(t)

	wallet, err := sdk.CreateWallet(types.WalletOptions{
		ClientOptions: &types.ClientOptions{},
		SecretManager: types.StrongholdSecretManager{
			Stronghold: types.StrongholdSecretManagerOptions{
				Password:     "1074",
				SnapshotPath: "./testdb/stronghold_snapshots",
			},
		},
		StoragePath: "./testdb/stronghold",
		CoinType:    types.CoinTypeSMR,
	})
	defer wallet.Destroy()

	require.NoError(t, err)
	require.NotNil(t, wallet)

	res, err := wallet.StoreMnemonic(Mnemonic)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	t.Log(res)

	address, err := wallet.GenerateEd25519Address(0, 0, "smr", nil)
	require.NoError(t, err)
	require.NotEmpty(t, address)

	bip32Chain := wasp_wallet_sdk.BuildBip32Chain(types.CoinTypeSMR, 0, false, 0)
	signedEssence, err := wallet.SignTransactionEssence(types.HexEncodedString(SignMessageFromEssenceHex), bip32Chain)
	require.NoError(t, err)
	require.NotEmpty(t, signedEssence)
}
