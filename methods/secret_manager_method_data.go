package methods

import "github.com/iotaledger/wasp-wallet-sdk/types"

type GenerateEd25519AddressMethodData struct {
	// AccountIndex corresponds to the JSON schema field "accountIndex".
	AccountIndex uint32 `json:"accountIndex" yaml:"accountIndex" mapstructure:"accountIndex"`

	// AddressIndex corresponds to the JSON schema field "addressIndex".
	AddressIndex uint32 `json:"addressIndex" yaml:"addressIndex" mapstructure:"addressIndex"`

	// Bech32Hrp corresponds to the JSON schema field "bech32Hrp".
	Bech32Hrp string `json:"bech32Hrp,omitempty" yaml:"bech32Hrp,omitempty" mapstructure:"bech32Hrp,omitempty"`

	// Options corresponds to the JSON schema field "options".
	Options *types.GenerateAddressOptions `json:"options,omitempty" yaml:"options,omitempty" mapstructure:"options,omitempty"`
}

type SignSecp256K1EcdsaMethodData struct {
	// Chain corresponds to the JSON schema field "chain".
	Chain types.IBip32Chain `json:"chain" yaml:"chain" mapstructure:"chain"`

	// Message corresponds to the JSON schema field "message".
	Message types.HexEncodedString `json:"message" yaml:"message" mapstructure:"message"`
}

type StoreMnemonicMethodData struct {
	// Mnemonic corresponds to the JSON schema field "mnemonic".
	Mnemonic string `json:"mnemonic" yaml:"mnemonic" mapstructure:"mnemonic"`
}

type SignatureUnlockMethodData struct {
	// Chain corresponds to the JSON schema field "chain".
	Chain types.IBip32Chain `json:"chain" yaml:"chain" mapstructure:"chain"`

	// SecretManager corresponds to the JSON schema field "secretManager".
	SecretManager types.SignatureUnlockMethodDataSecretManager `json:"secretManager" yaml:"secretManager" mapstructure:"secretManager"`

	// TransactionEssenceHash corresponds to the JSON schema field
	// "transactionEssenceHash".
	TransactionEssenceHash types.HexEncodedString `json:"transactionEssenceHash" yaml:"transactionEssenceHash" mapstructure:"transactionEssenceHash"`
}

type SignTransactionMethodData struct {
	// PreparedTransactionData corresponds to the JSON schema field
	// "preparedTransactionData".
	PreparedTransactionData types.PreparedTransactionData `json:"preparedTransactionData" yaml:"preparedTransactionData" mapstructure:"preparedTransactionData"`

	// SecretManager corresponds to the JSON schema field "secretManager".
	SecretManager SignTransactionMethodDataSecretManager `json:"secretManager" yaml:"secretManager" mapstructure:"secretManager"`
}

type SignTransactionMethodDataSecretManager interface{}

type SignEd25519MethodData struct {
	// Chain corresponds to the JSON schema field "chain".
	Chain types.IBip32Chain `json:"chain" yaml:"chain" mapstructure:"chain"`

	// Message corresponds to the JSON schema field "message".
	Message types.HexEncodedString `json:"message" yaml:"message" mapstructure:"message"`
}

type GenerateEvmAddressesMethodData struct {
	// Options corresponds to the JSON schema field "options".
	Options types.IGenerateAddressesOptions `json:"options" yaml:"options" mapstructure:"options"`
}

type GenerateEd25519AddressesMethodData struct {
	// Options corresponds to the JSON schema field "options".
	Options types.IGenerateAddressesOptions `json:"options" yaml:"options" mapstructure:"options"`
}
