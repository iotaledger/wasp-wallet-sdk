package types

type CoinType uint32

const (
	CoinTypeEther CoinType = 60
	CoinTypeSMR   CoinType = 4219
	CoinTypeIOTA  CoinType = 4218
)

type WalletType uint32

const HDWalletType WalletType = 44

type LedgerDeviceType string

const (
	LedgerDeviceTypeLedgerNanoS     LedgerDeviceType = "ledgerNanoS"
	LedgerDeviceTypeLedgerNanoSPlus LedgerDeviceType = "ledgerNanoSPlus"
	LedgerDeviceTypeLedgerNanoX     LedgerDeviceType = "ledgerNanoX"
)

type LedgerApp struct {
	// Name corresponds to the JSON schema field "name".
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	// Version corresponds to the JSON schema field "version".
	Version string `json:"version" yaml:"version" mapstructure:"version"`
}

type LedgerNanoStatus struct {
	// App corresponds to the JSON schema field "app".
	App *LedgerApp `json:"app,omitempty" yaml:"app,omitempty" mapstructure:"app,omitempty"`

	// BlindSigningEnabled corresponds to the JSON schema field "blindSigningEnabled".
	BlindSigningEnabled bool `json:"blindSigningEnabled,omitempty" yaml:"blindSigningEnabled,omitempty" mapstructure:"blindSigningEnabled,omitempty"`

	// BufferSize corresponds to the JSON schema field "bufferSize".
	BufferSize uint32 `json:"bufferSize,omitempty" yaml:"bufferSize,omitempty" mapstructure:"bufferSize,omitempty"`

	// Connected corresponds to the JSON schema field "connected".
	Connected bool `json:"connected" yaml:"connected" mapstructure:"connected"`

	// Device corresponds to the JSON schema field "device".
	Device *LedgerDeviceType `json:"device,omitempty" yaml:"device,omitempty" mapstructure:"device,omitempty"`

	// Locked corresponds to the JSON schema field "locked".
	Locked bool `json:"locked,omitempty" yaml:"locked,omitempty" mapstructure:"locked,omitempty"`
}

// GenerateAddressOptions Options for address generation, useful with a Ledger Nano SecretManager
type GenerateAddressOptions struct {
	// Internal corresponds to the JSON schema field "internal".
	Internal bool `json:"internal" yaml:"internal" mapstructure:"internal"`

	// LedgerNanoPrompt corresponds to the JSON schema field "ledgerNanoPrompt".
	LedgerNanoPrompt bool `json:"ledgerNanoPrompt" yaml:"ledgerNanoPrompt" mapstructure:"ledgerNanoPrompt"`
}

// IGenerateAddressOptions Options provided to Generate Address
type IGenerateAddressOptions struct {
	// Display the address on ledger devices.
	LedgerNanoPrompt bool `json:"ledgerNanoPrompt" yaml:"ledgerNanoPrompt" mapstructure:"ledgerNanoPrompt"`
}

type Ed25519Signature struct {
	// Message corresponds to the JSON schema field "message".
	Message string `json:"message" yaml:"message" mapstructure:"message"`

	PublicKey string `json:"publicKey" yaml:"message" mapstructure:"message"`

	// Ed25519Signature signature.
	Signature string `json:"signature" yaml:"signature" mapstructure:"signature"`
}

// Secret manager that uses a mnemonic.
type MnemonicSecretManager struct {
	// Mnemonic corresponds to the JSON schema field "mnemonic".
	Mnemonic string `json:"mnemonic" yaml:"mnemonic" mapstructure:"mnemonic"`
}

// Secret manager that uses a Ledger Nano hardware wallet or Speculos simulator.
type LedgerNanoSecretManager struct {
	// boolean indicates whether it's a simulator or not.
	LedgerNano bool `json:"ledgerNano" yaml:"ledgerNano" mapstructure:"ledgerNano"`
}

// Secret manager that uses Stronghold.
type StrongholdSecretManager struct {
	// Stronghold corresponds to the JSON schema field "stronghold".
	Stronghold StrongholdSecretManagerStronghold `json:"stronghold" yaml:"stronghold" mapstructure:"stronghold"`
}

type StrongholdSecretManagerStronghold struct {
	// Password corresponds to the JSON schema field "password".
	Password string `json:"password,omitempty" yaml:"password,omitempty" mapstructure:"password,omitempty"`

	// SnapshotPath corresponds to the JSON schema field "snapshotPath".
	SnapshotPath string `json:"snapshotPath,omitempty" yaml:"snapshotPath,omitempty" mapstructure:"snapshotPath,omitempty"`
}

// Options for the Wallet builder
type WalletOptions struct {
	// ClientOptions corresponds to the JSON schema field "clientOptions".
	ClientOptions *ClientOptions `json:"clientOptions,omitempty" yaml:"clientOptions,omitempty" mapstructure:"clientOptions,omitempty"`

	// CoinType corresponds to the JSON schema field "coinType".
	CoinType CoinType `json:"coinType,omitempty" yaml:"coinType,omitempty" mapstructure:"coinType,omitempty"`

	// SecretManager corresponds to the JSON schema field "secretManager".
	SecretManager WalletOptionsSecretManager `json:"secretManager,omitempty" yaml:"secretManager,omitempty" mapstructure:"secretManager,omitempty"`

	// StoragePath corresponds to the JSON schema field "storagePath".
	StoragePath string `json:"storagePath,omitempty" yaml:"storagePath,omitempty" mapstructure:"storagePath,omitempty"`
}

// Either LedgerNanoSecretManager | MnemonicSecretManager | StrongholdSecretManager
type WalletOptionsSecretManager interface{}

type InputSigningDataOutput map[string]interface{}

type IOutputMetadataResponse struct {
	// The block id the output was contained in.
	BlockID HexEncodedString `json:"blockId" yaml:"blockId" mapstructure:"blockId"`

	// Is the output spent.
	IsSpent bool `json:"isSpent" yaml:"isSpent" mapstructure:"isSpent"`

	// The ledger index at which these output was available at.
	LedgerIndex float64 `json:"ledgerIndex" yaml:"ledgerIndex" mapstructure:"ledgerIndex"`

	// The milestone index at which this output was booked into the ledger.
	MilestoneIndexBooked float64 `json:"milestoneIndexBooked" yaml:"milestoneIndexBooked" mapstructure:"milestoneIndexBooked"`

	// The milestone index at which this output was spent.
	MilestoneIndexSpent float64 `json:"milestoneIndexSpent,omitempty" yaml:"milestoneIndexSpent,omitempty" mapstructure:"milestoneIndexSpent,omitempty"`

	// The milestone timestamp this output was booked in the ledger.
	MilestoneTimestampBooked float64 `json:"milestoneTimestampBooked" yaml:"milestoneTimestampBooked" mapstructure:"milestoneTimestampBooked"`

	// The milestone timestamp this output was spent.
	MilestoneTimestampSpent float64 `json:"milestoneTimestampSpent,omitempty" yaml:"milestoneTimestampSpent,omitempty" mapstructure:"milestoneTimestampSpent,omitempty"`

	// The index for the output.
	OutputIndex float64 `json:"outputIndex" yaml:"outputIndex" mapstructure:"outputIndex"`

	// The transaction id for the output.
	TransactionID HexEncodedString `json:"transactionId" yaml:"transactionId" mapstructure:"transactionId"`

	// The transaction this output was spent with.
	TransactionIDSpent *HexEncodedString `json:"transactionIdSpent,omitempty" yaml:"transactionIdSpent,omitempty" mapstructure:"transactionIdSpent,omitempty"`
}

// Data for transaction inputs for signing and ordering of unlock blocks
type InputSigningData struct {
	// The chain derived from seed, only for ed25519 addresses
	Chain IBip32Chain `json:"chain,omitempty" yaml:"chain,omitempty" mapstructure:"chain,omitempty"`

	// The output
	Output InputSigningDataOutput `json:"output" yaml:"output" mapstructure:"output"`

	// The output metadata
	OutputMetadata IOutputMetadataResponse `json:"outputMetadata" yaml:"outputMetadata" mapstructure:"outputMetadata"`
}

// The remainder address
type RemainderAddress map[string]interface{}

// The remainder output
type RemainderOutput map[string]interface{}

// The RemainderValueStrategy
type RemainderValueStrategy interface{}

type Remainder struct {
	// The remainder address
	Address RemainderAddress `json:"address" yaml:"address" mapstructure:"address"`

	// The chain derived from seed, for the remainder addresses
	Chain IBip32Chain `json:"chain,omitempty" yaml:"chain,omitempty" mapstructure:"chain,omitempty"`

	// The remainder output
	Output RemainderOutput `json:"output" yaml:"output" mapstructure:"output"`
}

type RegularTransactionEssence struct {
	Type uint32 `json:"type" yaml:"type" mapstructure:"type"`

	PublicKey          string `json:"publicKey" yaml:"publicKey" mapstructure:"publicKey"`
	TransactionEssence string `json:"transactionEssence" yaml:"transactionEssence" mapstructure:"transactionEssence"`
}

type PreparedTransactionData struct {
	// Transaction essence
	Essence RegularTransactionEssence `json:"essence" yaml:"essence" mapstructure:"essence"`

	// Required address information for signing
	InputsData []InputSigningData `json:"inputsData" yaml:"inputsData" mapstructure:"inputsData"`

	// Optional remainder output information
	Remainder *Remainder `json:"remainder,omitempty" yaml:"remainder,omitempty" mapstructure:"remainder,omitempty"`
}

// A range with start and end values.
type Range struct {
	// End corresponds to the JSON schema field "end".
	End uint32 `json:"end" yaml:"end" mapstructure:"end"`

	// Start corresponds to the JSON schema field "start".
	Start uint32 `json:"start" yaml:"start" mapstructure:"start"`
}

func NewRange(start uint32, end uint32) Range {
	return Range{Start: start, End: end}
}

// Input options for GenerateAddresses
type IGenerateAddressesOptions struct {
	// AccountIndex corresponds to the JSON schema field "accountIndex".
	AccountIndex uint32 `json:"accountIndex,omitempty" yaml:"accountIndex,omitempty" mapstructure:"accountIndex,omitempty"`

	// Bech32 human readable part
	Bech32Hrp string `json:"bech32Hrp,omitempty" yaml:"bech32Hrp,omitempty" mapstructure:"bech32Hrp,omitempty"`

	// CoinType corresponds to the JSON schema field "coinType".
	CoinType CoinType `json:"coinType,omitempty" yaml:"coinType,omitempty" mapstructure:"coinType,omitempty"`

	// Internal addresses
	Internal bool `json:"internal,omitempty" yaml:"internal,omitempty" mapstructure:"internal,omitempty"`

	// Options corresponds to the JSON schema field "options".
	Options *IGenerateAddressOptions `json:"options,omitempty" yaml:"options,omitempty" mapstructure:"options,omitempty"`

	// Range corresponds to the JSON schema field "range".
	Range Range `json:"range,omitempty" yaml:"range,omitempty" mapstructure:"range,omitempty"`
}
