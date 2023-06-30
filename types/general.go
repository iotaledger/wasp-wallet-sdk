package types

type IBip32Chain [5]uint32
type HexEncodedString string

type ILoggerConfigLevelFilter string

const (
	LevelFilterOff   ILoggerConfigLevelFilter = "off"
	LevelFilterError ILoggerConfigLevelFilter = "error"
	LevelFilterWarn  ILoggerConfigLevelFilter = "warn"
	LevelFilterInfo  ILoggerConfigLevelFilter = "info"
	LevelFilterDebug ILoggerConfigLevelFilter = "debug"
	LevelFilterTrace ILoggerConfigLevelFilter = "trace"
)

// Logger output configuration.
type ILoggerConfig struct {
	// Color flag of an output.
	ColorEnabled bool `json:"colorEnabled,omitempty" yaml:"colorEnabled,omitempty" mapstructure:"colorEnabled,omitempty"`

	// Log level filter of an output.
	LevelFilter ILoggerConfigLevelFilter `json:"levelFilter,omitempty" yaml:"levelFilter,omitempty" mapstructure:"levelFilter,omitempty"`

	// Name of an output file, or `stdout` for standard output.
	Name string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	// Log target exclusions of an output.
	TargetExclusions []string `json:"targetExclusions,omitempty" yaml:"targetExclusions,omitempty" mapstructure:"targetExclusions,omitempty"`

	// Log target filters of an output.
	TargetFilter []string `json:"targetFilter,omitempty" yaml:"targetFilter,omitempty" mapstructure:"targetFilter,omitempty"`
}

// Time duration
type IDuration struct {
	// Nanos corresponds to the JSON schema field "nanos".
	Nanos float64 `json:"nanos" yaml:"nanos" mapstructure:"nanos"`

	// Secs corresponds to the JSON schema field "secs".
	Secs float64 `json:"secs" yaml:"secs" mapstructure:"secs"`
}

// Struct containing network and PoW related information
type INetworkInfo struct {
	// Fallback to local proof of work if the node doesn't support remote Pow
	FallbackToLocalPow bool `json:"fallbackToLocalPow" yaml:"fallbackToLocalPow" mapstructure:"fallbackToLocalPow"`

	// Local proof of work
	LocalPow bool `json:"localPow" yaml:"localPow" mapstructure:"localPow"`

	// Minimum proof of work score
	MinPowScore float64 `json:"minPowScore" yaml:"minPowScore" mapstructure:"minPowScore"`

	// Protocol parameters
	ProtocolParameters INodeInfoProtocol `json:"protocolParameters" yaml:"protocolParameters" mapstructure:"protocolParameters"`

	// Tips request interval during PoW in seconds
	TipsInterval float64 `json:"tipsInterval" yaml:"tipsInterval" mapstructure:"tipsInterval"`
}

// The Protocol Info.
type INodeInfoProtocol struct {
	// The human readable part of bech32 addresses.
	Bech32Hrp string `json:"bech32Hrp" yaml:"bech32Hrp" mapstructure:"bech32Hrp"`

	// The minimum score required for PoW.
	MinPowScore float64 `json:"minPowScore" yaml:"minPowScore" mapstructure:"minPowScore"`

	// The human friendly name of the network on which the node operates on.
	NetworkName string `json:"networkName" yaml:"networkName" mapstructure:"networkName"`

	// The rent structure used by given node/network.
	RentStructure IRent `json:"rentStructure" yaml:"rentStructure" mapstructure:"rentStructure"`

	// The token supply.
	TokenSupply string `json:"tokenSupply" yaml:"tokenSupply" mapstructure:"tokenSupply"`

	// The protocol version.
	Version float64 `json:"version" yaml:"version" mapstructure:"version"`
}

// Defines the parameters of rent cost calculations on objects which take node
// resources.
type IRent struct {
	// Defines the rent of a single virtual byte denoted in IOTA token.
	VByteCost uint32 `json:"vByteCost" yaml:"vByteCost" mapstructure:"vByteCost"`

	// The factor to be used for data only fields.
	VByteFactorData byte `json:"vByteFactorData" yaml:"vByteFactorData" mapstructure:"vByteFactorData"`

	// The factor to be used for key/lookup generating fields.
	VByteFactorKey byte `json:"vByteFactorKey" yaml:"vByteFactorKey" mapstructure:"vByteFactorKey"`
}

// Options for the MQTT broker.
type IMqttBrokerOptions struct {
	// AutomaticDisconnect corresponds to the JSON schema field "automaticDisconnect".
	AutomaticDisconnect bool `json:"automaticDisconnect,omitempty" yaml:"automaticDisconnect,omitempty" mapstructure:"automaticDisconnect,omitempty"`

	// MaxReconnectionAttempts corresponds to the JSON schema field
	// "maxReconnectionAttempts".
	MaxReconnectionAttempts uint32 `json:"maxReconnectionAttempts,omitempty" yaml:"maxReconnectionAttempts,omitempty" mapstructure:"maxReconnectionAttempts,omitempty"`

	// Port corresponds to the JSON schema field "port".
	Port uint32 `json:"port,omitempty" yaml:"port,omitempty" mapstructure:"port,omitempty"`

	// timeout in seconds
	Timeout uint32 `json:"timeout,omitempty" yaml:"timeout,omitempty" mapstructure:"timeout,omitempty"`

	// UseWs corresponds to the JSON schema field "useWs".
	UseWs bool `json:"useWs,omitempty" yaml:"useWs,omitempty" mapstructure:"useWs,omitempty"`
}

type AccountSyncOptions struct {
	// AliasOutputs corresponds to the JSON schema field "aliasOutputs".
	AliasOutputs bool `json:"aliasOutputs,omitempty" yaml:"aliasOutputs,omitempty" mapstructure:"aliasOutputs,omitempty"`

	// BasicOutputs corresponds to the JSON schema field "basicOutputs".
	BasicOutputs bool `json:"basicOutputs,omitempty" yaml:"basicOutputs,omitempty" mapstructure:"basicOutputs,omitempty"`

	// NftOutputs corresponds to the JSON schema field "nftOutputs".
	NftOutputs bool `json:"nftOutputs,omitempty" yaml:"nftOutputs,omitempty" mapstructure:"nftOutputs,omitempty"`
}

type AliasSyncOptions struct {
	// AliasOutputs corresponds to the JSON schema field "aliasOutputs".
	AliasOutputs bool `json:"aliasOutputs,omitempty" yaml:"aliasOutputs,omitempty" mapstructure:"aliasOutputs,omitempty"`

	// BasicOutputs corresponds to the JSON schema field "basicOutputs".
	BasicOutputs bool `json:"basicOutputs,omitempty" yaml:"basicOutputs,omitempty" mapstructure:"basicOutputs,omitempty"`

	// FoundryOutputs corresponds to the JSON schema field "foundryOutputs".
	FoundryOutputs bool `json:"foundryOutputs,omitempty" yaml:"foundryOutputs,omitempty" mapstructure:"foundryOutputs,omitempty"`

	// NftOutputs corresponds to the JSON schema field "nftOutputs".
	NftOutputs bool `json:"nftOutputs,omitempty" yaml:"nftOutputs,omitempty" mapstructure:"nftOutputs,omitempty"`
}

type NftSyncOptions struct {
	// AliasOutputs corresponds to the JSON schema field "aliasOutputs".
	AliasOutputs bool `json:"aliasOutputs,omitempty" yaml:"aliasOutputs,omitempty" mapstructure:"aliasOutputs,omitempty"`

	// BasicOutputs corresponds to the JSON schema field "basicOutputs".
	BasicOutputs bool `json:"basicOutputs,omitempty" yaml:"basicOutputs,omitempty" mapstructure:"basicOutputs,omitempty"`

	// NftOutputs corresponds to the JSON schema field "nftOutputs".
	NftOutputs bool `json:"nftOutputs,omitempty" yaml:"nftOutputs,omitempty" mapstructure:"nftOutputs,omitempty"`
}

// Sync options for an account
type SyncOptions struct {
	// Specifies what outputs should be synced for the ed25519 addresses from the
	// account.
	Account *AccountSyncOptions `json:"account,omitempty" yaml:"account,omitempty" mapstructure:"account,omitempty"`

	// Address index from which to start syncing addresses. 0 by default, using a
	// higher index will be faster because addresses with a lower index will be
	// skipped, but could result in a wrong balance for that reason
	AddressStartIndex float64 `json:"addressStartIndex,omitempty" yaml:"addressStartIndex,omitempty" mapstructure:"addressStartIndex,omitempty"`

	// Address index from which to start syncing internal addresses. 0 by default,
	// using a higher index will be faster because addresses with a lower index will
	// be skipped, but could result in a wrong balance for that reason
	AddressStartIndexInternal float64 `json:"addressStartIndexInternal,omitempty" yaml:"addressStartIndexInternal,omitempty" mapstructure:"addressStartIndexInternal,omitempty"`

	// Specific Bech32 encoded addresses of the account to sync, if addresses are
	// provided, then `address_start_index` will be ignored
	Addresses []string `json:"addresses,omitempty" yaml:"addresses,omitempty" mapstructure:"addresses,omitempty"`

	// Specifies what outputs should be synced for the address of an alias output.
	Alias *AliasSyncOptions `json:"alias,omitempty" yaml:"alias,omitempty" mapstructure:"alias,omitempty"`

	// Usually syncing is skipped if it's called in between 200ms, because there can
	// only be new changes every milestone and calling it twice "at the same time"
	// will not return new data When this to true, we will sync anyways, even if it's
	// called 0ms after the las sync finished. Default: false.
	ForceSyncing bool `json:"forceSyncing,omitempty" yaml:"forceSyncing,omitempty" mapstructure:"forceSyncing,omitempty"`

	// Specifies what outputs should be synced for the address of an nft output.
	Nft *NftSyncOptions `json:"nft,omitempty" yaml:"nft,omitempty" mapstructure:"nft,omitempty"`

	// SyncIncomingTransactions corresponds to the JSON schema field
	// "syncIncomingTransactions".
	SyncIncomingTransactions bool `json:"syncIncomingTransactions,omitempty" yaml:"syncIncomingTransactions,omitempty" mapstructure:"syncIncomingTransactions,omitempty"`

	// Sync native token foundries, so their metadata can be returned in the balance.
	// Default: false.
	SyncNativeTokenFoundries bool `json:"syncNativeTokenFoundries,omitempty" yaml:"syncNativeTokenFoundries,omitempty" mapstructure:"syncNativeTokenFoundries,omitempty"`

	// Specifies if only basic outputs with an AddressUnlockCondition alone should be
	// synced, will overwrite `account`, `alias` and `nft` options. Default: false.
	SyncOnlyMostBasicOutputs bool `json:"syncOnlyMostBasicOutputs,omitempty" yaml:"syncOnlyMostBasicOutputs,omitempty" mapstructure:"syncOnlyMostBasicOutputs,omitempty"`

	// Checks pending transactions and promotes/reattaches them if necessary. Default:
	// true.
	SyncPendingTransactions bool `json:"syncPendingTransactions,omitempty" yaml:"syncPendingTransactions,omitempty" mapstructure:"syncPendingTransactions,omitempty"`
}

type IAuth struct {
	// BasicAuthNamePwd corresponds to the JSON schema field "basicAuthNamePwd".
	BasicAuthNamePwd []string `json:"basicAuthNamePwd,omitempty" yaml:"basicAuthNamePwd,omitempty" mapstructure:"basicAuthNamePwd,omitempty"`

	// Jwt corresponds to the JSON schema field "jwt".
	Jwt string `json:"jwt,omitempty" yaml:"jwt,omitempty" mapstructure:"jwt,omitempty"`
}
