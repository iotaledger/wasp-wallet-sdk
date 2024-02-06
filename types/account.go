package types

type BaseCallAccountMethod[T BaseCallAccountMethodWrap[any]] struct {
	AccountId uint32 `json:"accountId"`
	Method    T      `json:"method"`
}

type BaseCallAccountMethodWrap[T any] struct {
	Name string `json:"name"`
	Data T      `json:"data"`
}

type GenerateAccountEd25519Addresses struct {
	Amount  uint32                 `json:"amount"`
	Options GenerateAddressOptions `json:"options"`
}

func NewGenerateAccountEd25519Addresses(amount uint32, options GenerateAddressOptions) BaseCallAccountMethodWrap[any] {
	return BaseCallAccountMethodWrap[any]{
		Name: "generateEd25519Addresses",
		Data: GenerateAccountEd25519Addresses{
			Amount:  amount,
			Options: options,
		},
	}
}
