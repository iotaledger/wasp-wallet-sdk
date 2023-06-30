package methods

func BackupMethod[T BackupMethodData](data T) BaseRequest[T] {
	method := "backup"

	return NewBaseRequest(method, data)
}

func ChangeStrongholdPasswordMethod[T ChangeStrongholdPasswordMethodData](data T) BaseRequest[T] {
	method := "changeStrongholdPassword"

	return NewBaseRequest(method, data)
}

func ClearStrongholdPasswordMethod() BaseRequest[NoType] {
	method := "clearStrongholdPassword"

	return NewBaseRequestNoData(method)
}

func ClearListenersMethod[T ClearListenersMethodData](data T) BaseRequest[T] {
	method := "clearListeners"

	return NewBaseRequest(method, data)
}

func CreateAccountMethod[T CreateAccountPayloadMethodData](data T) BaseRequest[T] {
	method := "createAccount"

	return NewBaseRequest(method, data)
}

func GenerateMnemonicMethod() BaseRequest[NoType] {
	method := "generateMnemonic"

	return NewBaseRequestNoData(method)
}

func GetAccountIndexesMethod() BaseRequest[NoType] {
	method := "getAccountIndexes"

	return NewBaseRequestNoData(method)
}

func GetAccountsMethod() BaseRequest[NoType] {
	method := "getAccounts"

	return NewBaseRequestNoData(method)
}

func GetAccountMethod[T GetAccountMethodData](data T) BaseRequest[T] {
	method := "getAccount"

	return NewBaseRequest(method, data)
}

func IsStrongholdPasswordAvailableMethod() BaseRequest[NoType] {
	method := "isStrongholdPasswordAvailable"

	return NewBaseRequestNoData(method)
}

func RecoverAccountsMethod[T RecoverAccountsMethodData](data T) BaseRequest[T] {
	method := "recoverAccounts"

	return NewBaseRequest(method, data)
}

func RemoveLatestAccountMethod() BaseRequest[NoType] {
	method := "removeLatestAccount"

	return NewBaseRequestNoData(method)
}

func RestoreBackupMethod[T RestoreBackupMethodData](data T) BaseRequest[T] {
	method := "restoreBackup"

	return NewBaseRequest(method, data)
}

func SetClientOptionsMethod[T SetClientOptionsMethodData](data T) BaseRequest[T] {
	method := "setClientOptions"

	return NewBaseRequest(method, data)
}

func SetStrongholdPasswordMethod[T SetStrongholdPasswordMethodData](data T) BaseRequest[T] {
	method := "setStrongholdPassword"

	return NewBaseRequest(method, data)
}

func SetStrongholdPasswordClearIntervalMethod[T SetStrongholdPasswordClearIntervalMethodData](data T) BaseRequest[T] {
	method := "setStrongholdPasswordClearInterval"

	return NewBaseRequest(method, data)
}

func StartBackgroundSyncMethod[T StartBackgroundSyncMethodData](data T) BaseRequest[T] {
	method := "startBackgroundSync"

	return NewBaseRequest(method, data)
}

func StopBackgroundSyncMethod() BaseRequest[NoType] {
	method := "stopBackgroundSync"

	return NewBaseRequestNoData(method)
}

func UpdateNodeAuthMethod[T UpdateNodeAuthMethodData](data T) BaseRequest[T] {
	method := "updateNodeAuth"

	return NewBaseRequest(method, data)
}
