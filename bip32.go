package wasp_wallet_sdk

import "github.com/iotaledger/wasp-wallet-sdk/types"

func BuildBip32Chain(coinType types.CoinType, accountIndex uint32, internalAddress bool, addressIndex uint32) types.IBip32Chain {
	var internalAddressInt uint32

	if internalAddress {
		internalAddressInt = 1
	}

	return types.IBip32Chain{
		uint32(types.HDWalletType),
		uint32(coinType),
		accountIndex,
		internalAddressInt,
		addressIndex,
	}
}
