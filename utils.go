package wasp_wallet_sdk

import (
	"github.com/awnumar/memguard"

	"github.com/iotaledger/wasp-wallet-sdk/methods"
)

type Utils struct {
	sdk *IOTASDK
}

func (u *Utils) GenerateMnemonic() (*memguard.Enclave, error) {
	mnemonic, free, err := u.sdk.CallUtilsMethod(methods.GenerateMnemonicMethod())
	defer free()
	if err != nil {
		return nil, err
	}

	response, err := methods.ParseResponseProtectedString(mnemonic, err)
	if err != nil {
		return nil, err
	}

	return response, nil
}
