//go:build IOTA_SDK_WITH_WALLET

package wasp_wallet_sdk

import (
	"github.com/iotaledger/wasp-wallet-sdk/methods"
	"github.com/iotaledger/wasp-wallet-sdk/types"
)

type Wallet struct {
	sdk              *IOTASDK
	walletPtr        IotaWalletPtr
	clientPtr        IotaClientPtr
	secretManagerPtr IotaSecretManagerPtr
}

func (i *IOTASDK) CreateWallet(walletOptions types.WalletOptions) (wallet *Wallet, err error) {
	msg, free, err := SerializeGuarded(walletOptions)
	defer free()
	if err != nil {
		return nil, err
	}

	var walletPtr IotaWalletPtr
	if walletPtr = i.libCreateWallet(msg); walletPtr == 0 {
		return nil, i.GetLastError()
	}

	clientPtr, err := i.GetClientFromWallet(walletPtr)
	if err != nil {
		return nil, err
	}

	secretManagerPtr, err := i.GetSecretManagerFromWallet(walletPtr)
	if err != nil {
		return nil, err
	}

	return NewWallet(i, walletPtr, clientPtr, secretManagerPtr), nil
}

func NewWallet(sdk *IOTASDK, walletPtr IotaWalletPtr, clientPtr IotaClientPtr, secretManagerPtr IotaSecretManagerPtr) *Wallet {
	return &Wallet{
		sdk:              sdk,
		walletPtr:        walletPtr,
		clientPtr:        clientPtr,
		secretManagerPtr: secretManagerPtr,
	}
}

func (s *Wallet) Destroy() {
	_ = s.sdk.DestroyWallet(s.walletPtr)
}

func (s *Wallet) GetLedgerStatus() (*types.LedgerNanoStatus, error) {
	ledgerNanoStatus, err := s.sdk.CallWalletMethod(s.walletPtr, methods.GetLedgerNanoStatusMethod())
	if err != nil {
		return nil, err
	}

	status, err := methods.ParseResponse[types.LedgerNanoStatus](ledgerNanoStatus, err)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (s *Wallet) CreateAccount(alias string, bech32Hrp string, options *types.GenerateAddressOptions) (any, error) {
	ledgerNanoStatus, err := s.sdk.CallWalletMethod(s.walletPtr, methods.CreateAccountMethod(methods.CreateAccountPayloadMethodData{
		Bech32Hrp: bech32Hrp,
		Alias:     alias,
	}))
	if err != nil {
		return "", err
	}

	address, err := methods.ParseResponse[any](ledgerNanoStatus, err)
	if err != nil || address == nil {
		return "", err
	}

	return *address, nil
}

func (s *Wallet) GenerateEd25519Address(addressIndex uint32, accountIndex uint32, bech32Hrp string, options *types.GenerateAddressOptions) (string, error) {
	ledgerNanoStatus, err := s.sdk.CallWalletMethod(s.walletPtr, methods.GenerateEd25519AddressMethod(methods.GenerateEd25519AddressMethodData{
		AddressIndex: addressIndex,
		AccountIndex: accountIndex,
		Bech32Hrp:    bech32Hrp,
		Options:      options,
	}))
	if err != nil {
		return "", err
	}

	address, err := methods.ParseResponse[string](ledgerNanoStatus, err)
	if err != nil || address == nil {
		return "", err
	}

	return *address, nil
}

func (s *Wallet) StoreMnemonic(mnemonic string) (bool, error) {
	success, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.StoreMnemonicMethod(methods.StoreMnemonicMethodData{
		Mnemonic: mnemonic,
	}))
	if err != nil {
		return false, err
	}

	return methods.ParseResponseStatus(success, err)
}

func (s *Wallet) SignTransactionEssence(txEssence types.HexEncodedString, bip44Chain types.Bip44Chain) (*types.Ed25519Signature, error) {
	signedMessageStr, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.SignEd25519Method(methods.SignEd25519MethodData{
		Message: txEssence,
		Chain:   bip44Chain,
	}))
	if err != nil {
		return nil, err
	}

	return methods.ParseResponse[types.Ed25519Signature](signedMessageStr, err)
}
