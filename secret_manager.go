package wasp_wallet_sdk

import (
	"errors"

	"github.com/iotaledger/wasp_wallet_sdk/methods"
	"github.com/iotaledger/wasp_wallet_sdk/types"
)

type SecretManager struct {
	sdk              *IOTASDK
	secretManagerPtr IotaSecretManagerPtr
}

func (s *SecretManager) Destroy() {
	_ = s.sdk.DestroySecretManager(s.secretManagerPtr)
}

func (s *SecretManager) GetLedgerStatus() (*types.LedgerNanoStatus, error) {
	ledgerNanoStatus, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.GetLedgerNanoStatusMethod())
	if err != nil {
		return nil, err
	}

	status, err := methods.ParseResponse[types.LedgerNanoStatus](ledgerNanoStatus, err)
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (s *SecretManager) CreateAccount(bech32Hrp string, alias string) (any, error) {
	ledgerNanoStatus, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.CreateAccountMethod(methods.CreateAccountPayloadMethodData{
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

func (s *SecretManager) GenerateEvmAddresses(addressRange types.Range, accountIndex uint32, bech32Hrp string, options *types.IGenerateAddressOptions) ([]string, error) {
	evmAddresses, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.GenerateEVMAddressMethod(methods.GenerateEvmAddressesMethodData{
		Options: types.IGenerateAddressesOptions{
			AccountIndex: accountIndex,
			Bech32Hrp:    bech32Hrp,
			CoinType:     types.CoinTypeEther,
			Internal:     false,
			Options:      options,
			Range:        addressRange,
		},
	}))
	if err != nil {
		return []string{}, err
	}

	address, err := methods.ParseResponse[[]string](evmAddresses, err)
	if err != nil || address == nil {
		return []string{}, err
	}

	return *address, nil
}

func (s *SecretManager) GenerateEd25519Addresses(addressRange types.Range, accountIndex uint32, bech32Hrp string, coinType types.CoinType, options *types.IGenerateAddressOptions) ([]string, error) {
	ledgerNanoStatus, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.GenerateEd25519AddressesMethod(methods.GenerateEd25519AddressesMethodData{
		Options: types.IGenerateAddressesOptions{
			AccountIndex: accountIndex,
			Bech32Hrp:    bech32Hrp,
			CoinType:     coinType,
			Internal:     false,
			Options:      options,
			Range:        addressRange,
		},
	}))
	if err != nil {
		return []string{}, err
	}

	address, err := methods.ParseResponse[[]string](ledgerNanoStatus, err)
	if err != nil || address == nil {
		return []string{}, err
	}

	return *address, nil
}

func (s *SecretManager) GenerateEd25519Address(addressIndex uint32, accountIndex uint32, bech32Hrp string, coinType types.CoinType, options *types.IGenerateAddressOptions) (string, error) {
	addresses, err := s.GenerateEd25519Addresses(types.NewRange(addressIndex, addressIndex+1), accountIndex, bech32Hrp, coinType, options)
	if err != nil {
		return "", err
	}
	if len(addresses) == 0 {
		return "", errors.New("failed to get address")
	}

	return addresses[0], nil
}

func (s *SecretManager) StoreMnemonic(mnemonic string) (bool, error) {
	success, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.StoreMnemonicMethod(methods.StoreMnemonicMethodData{
		Mnemonic: mnemonic,
	}))
	if err != nil {
		return false, err
	}

	return methods.ParseResponseStatus(success, err)
}

func (s *SecretManager) SignTransactionEssence(txEssence types.HexEncodedString, bip32Chain types.IBip32Chain) (*types.Ed25519Signature, error) {
	signedMessageStr, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.SignEd25519Method(methods.SignEd25519MethodData{
		Message: txEssence,
		Chain:   bip32Chain,
	}))
	if err != nil {
		return nil, err
	}

	return methods.ParseResponse[types.Ed25519Signature](signedMessageStr, err)
}

func NewMnemonicSecretManager(sdk *IOTASDK, mnemonicOptions types.MnemonicSecretManager) (*SecretManager, error) {
	secretManagerPtr, err := sdk.CreateSecretManager(mnemonicOptions)
	if err != nil {
		return nil, err
	}

	return &SecretManager{
		sdk:              sdk,
		secretManagerPtr: secretManagerPtr,
	}, nil
}

func NewStrongholdSecretManager(sdk *IOTASDK, strongholdOptions types.StrongholdSecretManagerStronghold) (*SecretManager, error) {
	secretManagerPtr, err := sdk.CreateSecretManager(types.StrongholdSecretManager{Stronghold: strongholdOptions})
	if err != nil {
		return nil, err
	}

	return &SecretManager{
		sdk:              sdk,
		secretManagerPtr: secretManagerPtr,
	}, nil
}

func NewLedgerSecretManager(sdk *IOTASDK, ledgerOptions types.LedgerNanoSecretManager) (*SecretManager, error) {
	secretManagerPtr, err := sdk.CreateSecretManager(ledgerOptions)
	if err != nil {
		return nil, err
	}

	return &SecretManager{
		sdk:              sdk,
		secretManagerPtr: secretManagerPtr,
	}, nil
}
