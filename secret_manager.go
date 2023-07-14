package wasp_wallet_sdk

import (
	"errors"

	"github.com/awnumar/memguard"

	"github.com/iotaledger/wasp-wallet-sdk/methods"
	"github.com/iotaledger/wasp-wallet-sdk/types"
)

type SecretManager struct {
	sdk              *IOTASDK
	secretManagerPtr IotaSecretManagerPtr
}

// NewMnemonicSecretManager creates or opens an in-memory Mnemonic based secret storage
func NewMnemonicSecretManager(sdk *IOTASDK, mnemonic *memguard.Enclave) (*SecretManager, error) {
	buffer, err := mnemonic.Open()
	if err != nil {
		return nil, err
	}
	defer buffer.Destroy()

	secretManagerPtr, err := sdk.CreateSecretManager(types.MnemonicSecretManager{
		Mnemonic: buffer.String(),
	})
	if err != nil {
		return nil, err
	}

	return &SecretManager{
		sdk:              sdk,
		secretManagerPtr: secretManagerPtr,
	}, nil
}

func NewStrongholdSecretManager(sdk *IOTASDK, password *memguard.Enclave, snapshotPath string) (*SecretManager, error) {
	buffer, err := password.Open()
	if err != nil {
		return nil, err
	}
	defer buffer.Destroy()

	secretManagerPtr, err := sdk.CreateSecretManager(types.StrongholdSecretManager{
		Stronghold: types.StrongholdSecretManagerOptions{
			Password:     buffer.String(),
			SnapshotPath: snapshotPath,
		},
	})
	if err != nil {
		return nil, err
	}

	return &SecretManager{
		sdk:              sdk,
		secretManagerPtr: secretManagerPtr,
	}, nil
}

// NewLedgerSecretManager creates or opens a Ledger based secret storage
func NewLedgerSecretManager(sdk *IOTASDK, isEmulator bool) (*SecretManager, error) {
	secretManagerPtr, err := sdk.CreateSecretManager(&types.LedgerNanoSecretManager{
		LedgerNano: isEmulator,
	})
	if err != nil {
		return nil, err
	}

	return &SecretManager{
		sdk:              sdk,
		secretManagerPtr: secretManagerPtr,
	}, nil
}

func (s *SecretManager) Destroy() {
	_ = s.sdk.DestroySecretManager(s.secretManagerPtr)
}

func (s *SecretManager) GetLedgerStatus() (*types.LedgerNanoStatus, error) {
	ledgerNanoStatus, free, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.GetLedgerNanoStatusMethod())
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

func (s *SecretManager) CreateAccount(bech32Hrp string, alias string) (any, error) {
	ledgerNanoStatus, free, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.CreateAccountMethod(methods.CreateAccountPayloadMethodData{
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

func (s *SecretManager) GenerateEvmAddresses(addressRange types.Range, accountIndex uint32, bech32Hrp string, options *types.IGenerateAddressOptions) ([]string, error) {
	evmAddresses, free, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.GenerateEVMAddressMethod(methods.GenerateEvmAddressesMethodData{
		Options: types.IGenerateAddressesOptions{
			AccountIndex: accountIndex,
			Bech32Hrp:    bech32Hrp,
			CoinType:     types.CoinTypeEther,
			Internal:     false,
			Options:      options,
			Range:        addressRange,
		},
	}))
	defer free()
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
	ledgerNanoStatus, free, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.GenerateEd25519AddressesMethod(methods.GenerateEd25519AddressesMethodData{
		Options: types.IGenerateAddressesOptions{
			AccountIndex: accountIndex,
			Bech32Hrp:    bech32Hrp,
			CoinType:     coinType,
			Internal:     false,
			Options:      options,
			Range:        addressRange,
		},
	}))
	defer free()
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

func (s *SecretManager) StoreMnemonic(mnemonic *memguard.Enclave) (bool, error) {
	buffer, err := mnemonic.Open()
	if err != nil {
		return false, err
	}
	defer buffer.Destroy()

	success, free, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.StoreMnemonicMethod(methods.StoreMnemonicMethodData{
		Mnemonic: buffer.String(),
	}))
	defer free()
	if err != nil {
		return false, err
	}

	return methods.ParseResponseStatus(success, err)
}

func (s *SecretManager) SignTransactionEssence(txEssence types.HexEncodedString, bip32Chain types.IBip32Chain) (*types.Ed25519Signature, error) {
	signedMessageStr, free, err := s.sdk.CallSecretManagerMethod(s.secretManagerPtr, methods.SignEd25519Method(methods.SignEd25519MethodData{
		Message: txEssence,
		Chain:   bip32Chain,
	}))
	defer free()
	if err != nil {
		return nil, err
	}

	return methods.ParseResponse[types.Ed25519Signature](signedMessageStr, err)
}
