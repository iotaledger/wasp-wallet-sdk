package wasp_wallet_sdk

import (
	"github.com/iotaledger/wasp-wallet-sdk/methods"
)

type Utils struct {
	sdk *IOTASDK
}

func (u *Utils) GenerateMnemonic() (*string, error) {
	signedMessageStr, err := u.sdk.CallUtilsMethod(methods.GenerateMnemonicMethod())
	if err != nil {
		return nil, err
	}

	return methods.ParseResponse[string](signedMessageStr, err)
}
