// //go:build IOTA_SDK_WITH_WALLET

package wasp_wallet_sdk

import (
	"fmt"

	"github.com/ebitengine/purego"

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
	ledgerNanoStatus, free, err := s.sdk.CallWalletMethod(s.walletPtr, methods.GetLedgerNanoStatusMethod())
	defer free()

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
	ledgerNanoStatus, free, err := s.sdk.CallWalletMethod(s.walletPtr, methods.CreateAccountMethod(methods.CreateAccountPayloadMethodData{
		Bech32Hrp: bech32Hrp,
		Alias:     alias,
	}))
	defer free()
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
	ledgerNanoStatus, free, err := s.sdk.CallWalletMethod(s.walletPtr, methods.GenerateEd25519AddressMethod(methods.GenerateEd25519AddressMethodData{
		AddressIndex: addressIndex,
		AccountIndex: accountIndex,
		Bech32Hrp:    bech32Hrp,
		Options:      options,
	}))
	defer free()
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
	success, free, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.StoreMnemonicMethod(methods.StoreMnemonicMethodData{
		Mnemonic: mnemonic,
	}))
	defer free()
	if err != nil {
		return false, err
	}

	return methods.ParseResponseStatus(success, err)
}

func (s *Wallet) CallAccountMethod(accountId uint32, method types.BaseCallAccountMethodWrap[any]) (any, error) {
	call := types.BaseCallAccountMethod[types.BaseCallAccountMethodWrap[any]]{
		AccountId: accountId,
		Method:    method,
	}

	result, free, err := s.sdk.CallWalletMethod(s.walletPtr, methods.CallAccountMethod(call))
	defer free()
	if err != nil {
		return false, err
	}

	return methods.ParseResponseStatus(result, err)
}

func (s *Wallet) SignTransactionEssence(txEssence types.HexEncodedString, bip44Chain types.Bip44Chain) (*types.Ed25519Signature, error) {
	signedMessageStr, free, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.SignEd25519Method(methods.SignEd25519MethodData{
		Message: txEssence,
		Chain:   bip44Chain,
	}))
	defer free()
	if err != nil {
		return nil, err
	}

	return methods.ParseResponse[types.Ed25519Signature](signedMessageStr, err)
}

func (s *Wallet) ListenToUpdates() error {
	events := "[0,1,2,3,4,5]"
	eventsPtr, _ := CStringGo([]byte(events))

	cb := func(b uintptr) {
		fmt.Printf("HAI\n")
		str, free, err := s.sdk.CopyAndDestroyOriginalStringPtr(b)
		defer free()

		// TODO
		fmt.Println(err)
		fmt.Println("LEDGER EVENT FIRED")
		fmt.Printf("%s\n", string(str))
	}

	cbptr := purego.NewCallback(cb)

	res := s.sdk.libListenWallet(s.walletPtr, eventsPtr, cbptr)
	if !res {
		fmt.Printf("ERR \n")
		return s.sdk.GetLastError()
	}

	return nil
}
