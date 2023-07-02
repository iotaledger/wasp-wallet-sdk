package methods

import "github.com/iotaledger/wasp_wallet_sdk/types"

type BackupMethodData struct {
	// Destination corresponds to the JSON schema field "destination".
	Destination string `json:"destination" yaml:"destination" mapstructure:"destination"`

	// Password corresponds to the JSON schema field "password".
	Password string `json:"password" yaml:"password" mapstructure:"password"`
}

type ChangeStrongholdPasswordMethodData struct {
	// CurrentPassword corresponds to the JSON schema field "currentPassword".
	CurrentPassword string `json:"currentPassword" yaml:"currentPassword" mapstructure:"currentPassword"`

	// NewPassword corresponds to the JSON schema field "newPassword".
	NewPassword string `json:"newPassword" yaml:"newPassword" mapstructure:"newPassword"`
}

type ClaimOutputsMethodData struct {
	// OutputIdsToClaim corresponds to the JSON schema field "outputIdsToClaim".
	OutputIdsToClaim []string `json:"outputIdsToClaim" yaml:"outputIdsToClaim" mapstructure:"outputIdsToClaim"`
}

type ClearListenersMethodData struct {
	// Topics corresponds to the JSON schema field "topics".
	Topics []string `json:"topics" yaml:"topics" mapstructure:"topics"`
}

// Options for account creation
type CreateAccountPayloadMethodData struct {
	// Alias corresponds to the JSON schema field "alias".
	Alias string `json:"alias,omitempty" yaml:"alias,omitempty" mapstructure:"alias,omitempty"`

	// Bech32Hrp corresponds to the JSON schema field "bech32Hrp".
	Bech32Hrp string `json:"bech32Hrp,omitempty" yaml:"bech32Hrp,omitempty" mapstructure:"bech32Hrp,omitempty"`
}

type GetAccountMethodData struct {
	// AccountID corresponds to the JSON schema field "accountId".
	AccountID string `json:"accountId" yaml:"accountId" mapstructure:"accountId"`
}

type SetDefaultSyncOptionsMethodData struct {
	// Options corresponds to the JSON schema field "options".
	Options types.SyncOptions `json:"options" yaml:"options" mapstructure:"options"`
}

type SetDefaultSyncOptionsMethod struct {
	// Data corresponds to the JSON schema field "data".
	Data SetDefaultSyncOptionsMethodData `json:"data" yaml:"data" mapstructure:"data"`

	// Name corresponds to the JSON schema field "name".
	Name string `json:"name" yaml:"name" mapstructure:"name"`
}

type RecoverAccountsMethodData struct {
	// AccountGapLimit corresponds to the JSON schema field "accountGapLimit".
	AccountGapLimit float64 `json:"accountGapLimit" yaml:"accountGapLimit" mapstructure:"accountGapLimit"`

	// AccountStartIndex corresponds to the JSON schema field "accountStartIndex".
	AccountStartIndex float64 `json:"accountStartIndex" yaml:"accountStartIndex" mapstructure:"accountStartIndex"`

	// AddressGapLimit corresponds to the JSON schema field "addressGapLimit".
	AddressGapLimit float64 `json:"addressGapLimit" yaml:"addressGapLimit" mapstructure:"addressGapLimit"`

	// SyncOptions corresponds to the JSON schema field "syncOptions".
	SyncOptions *types.SyncOptions `json:"syncOptions,omitempty" yaml:"syncOptions,omitempty" mapstructure:"syncOptions,omitempty"`
}

type RestoreBackupMethodData struct {
	// IgnoreIfBech32Mismatch corresponds to the JSON schema field
	// "ignoreIfBech32Mismatch".
	IgnoreIfBech32Mismatch string `json:"ignoreIfBech32Mismatch,omitempty" yaml:"ignoreIfBech32Mismatch,omitempty" mapstructure:"ignoreIfBech32Mismatch,omitempty"`

	// IgnoreIfCoinTypeMismatch corresponds to the JSON schema field
	// "ignoreIfCoinTypeMismatch".
	IgnoreIfCoinTypeMismatch bool `json:"ignoreIfCoinTypeMismatch,omitempty" yaml:"ignoreIfCoinTypeMismatch,omitempty" mapstructure:"ignoreIfCoinTypeMismatch,omitempty"`

	// Password corresponds to the JSON schema field "password".
	Password string `json:"password" yaml:"password" mapstructure:"password"`

	// Source corresponds to the JSON schema field "source".
	Source string `json:"source" yaml:"source" mapstructure:"source"`
}

type SetClientOptionsMethodData struct {
	// ClientOptions corresponds to the JSON schema field "clientOptions".
	ClientOptions types.ClientOptions `json:"clientOptions" yaml:"clientOptions" mapstructure:"clientOptions"`
}

type SetStrongholdPasswordMethodData struct {
	// Password corresponds to the JSON schema field "password".
	Password string `json:"password" yaml:"password" mapstructure:"password"`
}

type SetStrongholdPasswordClearIntervalMethodData struct {
	// IntervalInMilliseconds corresponds to the JSON schema field
	// "intervalInMilliseconds".
	IntervalInMilliseconds uint32 `json:"intervalInMilliseconds,omitempty" yaml:"intervalInMilliseconds,omitempty" mapstructure:"intervalInMilliseconds,omitempty"`
}

type StartBackgroundSyncMethodData struct {
	// IntervalInMilliseconds corresponds to the JSON schema field
	// "intervalInMilliseconds".
	IntervalInMilliseconds float64 `json:"intervalInMilliseconds,omitempty" yaml:"intervalInMilliseconds,omitempty" mapstructure:"intervalInMilliseconds,omitempty"`

	// Options corresponds to the JSON schema field "options".
	Options *types.SyncOptions `json:"options,omitempty" yaml:"options,omitempty" mapstructure:"options,omitempty"`
}

type UpdateNodeAuthMethodData struct {
	// Auth corresponds to the JSON schema field "auth".
	Auth *types.IAuth `json:"auth,omitempty" yaml:"auth,omitempty" mapstructure:"auth,omitempty"`

	// URL corresponds to the JSON schema field "url".
	URL string `json:"url" yaml:"url" mapstructure:"url"`
}
