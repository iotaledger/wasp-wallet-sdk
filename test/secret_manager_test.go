package test

import (
	"testing"

	"github.com/awnumar/memguard"
	"github.com/stretchr/testify/require"

	wasp_wallet_sdk "github.com/iotaledger/wasp-wallet-sdk"
	"github.com/iotaledger/wasp-wallet-sdk/types"
)

/**
Bare-bones secret manager test without Wallet or Client functionality
*/

func TestSecretManagerMnemonic(t *testing.T) {
	sdk := GetOrInitTest(t)

	secretManager, err := wasp_wallet_sdk.NewMnemonicSecretManager(sdk, memguard.NewEnclave([]byte(Mnemonic)))
	require.NoError(t, err)
	require.NotNil(t, secretManager)
	defer secretManager.Destroy()

	bip44Chain := wasp_wallet_sdk.BuildBip44Chain(types.CoinTypeSMR, 0, 0)
	result, err := secretManager.SignTransactionEssence(SignMessageFromEssenceHex, bip44Chain)

	require.NoError(t, err)
	require.NotNil(t, result)

	evmAddresses, err := secretManager.GenerateEvmAddresses(types.NewRange(0, 10), 0, "smr", nil)
	require.NoError(t, err)
	require.NotNil(t, evmAddresses)
	require.Len(t, evmAddresses, 10)

	ed25519Address, err := secretManager.GenerateEd25519Address(0, 0, "smr", types.CoinTypeSMR, nil)
	require.NoError(t, err)
	require.NotNil(t, ed25519Address)

	ed25519Addresses, err := secretManager.GenerateEd25519Addresses(types.NewRange(0, 10), 0, "smr", types.CoinTypeSMR, nil)
	require.NoError(t, err)
	require.NotNil(t, ed25519Address)
	require.Equal(t, ed25519Addresses[0], ed25519Address)
}

func TestSecretManagerLedger(t *testing.T) {
	sdk := GetOrInitTest(t)

	secretManager, err := wasp_wallet_sdk.NewLedgerSecretManager(sdk, UseLedgerSimulator)
	require.NoError(t, err)
	require.NotNil(t, secretManager)
	defer secretManager.Destroy()

	status, err := secretManager.GetLedgerStatus()
	require.NoError(t, err)
	require.NotNil(t, status)

	bip44Chain := wasp_wallet_sdk.BuildBip44Chain(types.CoinTypeSMR, 0, 0)
	signedEssence, err := secretManager.SignTransactionEssence(types.HexEncodedString(SignMessageFromEssenceHex), bip44Chain)
	require.NoError(t, err)
	require.NotEmpty(t, signedEssence)

	/* Ledger (Shimmer app) does not support generating EVM Addresses
	evmAddresses, err := secretManager.GenerateEvmAddresses(types.NewRange(0, 10), 0, "smr", nil)
	require.NoError(t, err)
	require.NotNil(t, evmAddresses)
	require.Len(t, evmAddresses, 10)
	*/

	ed25519Address, err := secretManager.GenerateEd25519Address(0, 0, "smr", types.CoinTypeSMR, nil)
	require.NoError(t, err)
	require.NotNil(t, ed25519Address)

	ed25519Addresses, err := secretManager.GenerateEd25519Addresses(types.NewRange(0, 10), 0, "smr", types.CoinTypeSMR, nil)
	require.NoError(t, err)
	require.NotNil(t, ed25519Address)

	require.Equal(t, ed25519Addresses[0], ed25519Address)
}

func TestSecretManagerStronghold(t *testing.T) {
	sdk := GetOrInitTest(t)

	secretManager, err := wasp_wallet_sdk.NewStrongholdSecretManager(sdk, memguard.NewEnclave([]byte("4389t!!$$hbg02pgn")), "./testdb/client.stronghold")
	require.NoError(t, err)
	require.NotNil(t, secretManager)
	defer secretManager.Destroy()

	res, err := secretManager.StoreMnemonic(memguard.NewEnclave([]byte(Mnemonic)))
	require.NoError(t, err)
	require.NotEmpty(t, res)
	t.Log(res)

	address, err := secretManager.GenerateEd25519Address(0, 0, "smr", types.CoinTypeSMR, nil)
	require.NoError(t, err)
	require.NotEmpty(t, address)

	bip44Chain := wasp_wallet_sdk.BuildBip44Chain(types.CoinTypeSMR, 0, 0)
	signedEssence, err := secretManager.SignTransactionEssence(types.HexEncodedString(SignMessageFromEssenceHex), bip44Chain)
	require.NoError(t, err)
	require.NotEmpty(t, signedEssence)

	evmAddresses, err := secretManager.GenerateEvmAddresses(types.NewRange(0, 10), 0, "smr", nil)
	require.NoError(t, err)
	require.NotNil(t, evmAddresses)
	require.Len(t, evmAddresses, 10)

	ed25519Address, err := secretManager.GenerateEd25519Address(0, 0, "smr", types.CoinTypeSMR, nil)
	require.NoError(t, err)
	require.NotNil(t, ed25519Address)

	ed25519Addresses, err := secretManager.GenerateEd25519Addresses(types.NewRange(0, 10), 0, "smr", types.CoinTypeSMR, nil)
	require.NoError(t, err)
	require.NotNil(t, ed25519Address)

	require.Equal(t, ed25519Addresses[0], ed25519Address)
}
