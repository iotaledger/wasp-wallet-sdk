package test

import (
	"testing"

	"github.com/iotaledger/wasp_wallet_sdk"
	"github.com/iotaledger/wasp_wallet_sdk/types"

	"github.com/stretchr/testify/require"
)

/**
Wallet tests are without any client configuration and are therefore offline
*/

func TestWalletMnemonic(t *testing.T) {
	sdk := InitTest(t)

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
}

func TestWalletLedger(t *testing.T) {
	sdk := InitTest(t)

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

	bip39Chain := wasp_wallet_sdk.BuildBip32Chain(types.CoinTypeSMR, 0, false, 0)
	signedEssence, err := wallet.SignTransactionEssence(types.HexEncodedString(SignMessageFromEssenceHex), bip39Chain)
	require.NoError(t, err)
	require.NotEmpty(t, signedEssence)
}

func TestWalletStronghold(t *testing.T) {
	sdk := InitTest(t)

	wallet, err := sdk.CreateWallet(types.WalletOptions{
		ClientOptions: &types.ClientOptions{},
		SecretManager: types.StrongholdSecretManager{
			Stronghold: types.StrongholdSecretManagerStronghold{
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

	res, err := wallet.CreateAccount(0, 0, "smr", nil)
	require.NoError(t, err)
	t.Log(res)

	address, err := wallet.GenerateEd25519Address(0, 0, "smr", nil)
	require.NoError(t, err)
	require.NotEmpty(t, address)

	bip39Chain := wasp_wallet_sdk.BuildBip32Chain(types.CoinTypeSMR, 0, false, 0)
	signedEssence, err := wallet.SignTransactionEssence(types.HexEncodedString(SignMessageFromEssenceHex), bip39Chain)
	require.NoError(t, err)
	require.NotEmpty(t, signedEssence)
}
