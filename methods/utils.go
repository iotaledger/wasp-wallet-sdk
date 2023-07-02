package methods

func GenerateMnemonicMethod() BaseRequest[NoType] {
	method := "generateMnemonic"

	return NewBaseRequestNoData(method)
}
