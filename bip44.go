package wasp_wallet_sdk

import "github.com/iotaledger/wasp-wallet-sdk/types"

func BuildBip44Chain(coinType types.CoinType, accountIndex uint32, addressIndex uint32) types.Bip44Chain {
	return types.Bip44Chain{
		CoinType:     uint32(coinType),
		Account:      accountIndex,
		Change:       uint32(0),
		AddressIndex: addressIndex,
	}
}
