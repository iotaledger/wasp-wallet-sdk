package methods

func SignEd25519Method[T SignEd25519MethodData](data T) BaseRequest[T] {
	method := "signEd25519"

	return NewBaseRequest(method, data)
}

func SignSecp256K1EcdsaMethod[T SignSecp256K1EcdsaMethodData](data T) BaseRequest[T] {
	method := "signSecp256k1Ecdsa"

	return NewBaseRequest(method, data)
}

func StoreMnemonicMethod[T StoreMnemonicMethodData](data T) BaseRequest[T] {
	method := "storeMnemonic"

	return NewBaseRequest(method, data)
}

func SignatureUnlockMethod[T SignatureUnlockMethodData](data T) BaseRequest[T] {
	method := "signatureUnlock"

	return NewBaseRequest(method, data)
}

func SignTransactionMethod[T SignTransactionMethodData](data T) BaseRequest[T] {
	method := "signTransaction"

	return NewBaseRequest(method, data)
}

func GetLedgerNanoStatusMethod() BaseRequest[NoType] {
	method := "getLedgerNanoStatus"

	return NewBaseRequestNoData(method)
}

func GenerateEd25519AddressMethod[T GenerateEd25519AddressMethodData](data T) BaseRequest[T] {
	method := "generateEd25519Address"

	return NewBaseRequest(method, data)
}
