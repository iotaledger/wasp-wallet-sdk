package wasp_wallet_sdk

import (
	"encoding/json"
	"errors"
	"strings"
	"unsafe"

	"github.com/iotaledger/wasp_wallet_sdk/lib_loader"
	"github.com/iotaledger/wasp_wallet_sdk/types"

	"github.com/ebitengine/purego"
)

func serialize[T any](obj T) (string, error) {
	jsonMessage, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(jsonMessage), nil
}

type (
	IotaClientPtr        uintptr
	IotaWalletPtr        uintptr
	IotaSecretManagerPtr uintptr
)

type IOTASDK struct {
	handle uintptr

	libInitLogger func(string) bool

	libCreateClient        func(string) IotaClientPtr
	libCreateWallet        func(string) IotaWalletPtr
	libCreateSecretManager func(string) IotaSecretManagerPtr

	libDestroyClient        func(ptr IotaClientPtr) bool
	libDestroyWallet        func(ptr IotaWalletPtr) bool
	libDestroySecretManager func(ptr IotaSecretManagerPtr) bool
	libDestroyString        func(unsafe.Pointer) bool

	libGetClientFromWallet        func(ptr IotaWalletPtr) IotaClientPtr
	libGetSecretManagerFromWallet func(ptr IotaWalletPtr) IotaSecretManagerPtr

	libCallClientMethod        func(IotaClientPtr, string) string
	libCallWalletMethod        func(IotaWalletPtr, string) string
	libCallSecretManagerMethod func(IotaSecretManagerPtr, string) string

	libGetLastError func() string
}

func NewIotaSDK(libPath string) (*IOTASDK, error) {
	iotaSDK, err := lib_loader.LoadLibrary(libPath)
	if err != nil {
		return nil, err
	}

	iotaSDKNative := IOTASDK{
		handle: iotaSDK,
	}

	purego.RegisterLibFunc(&iotaSDKNative.libInitLogger, iotaSDK, "init_logger")

	purego.RegisterLibFunc(&iotaSDKNative.libCreateClient, iotaSDK, "create_client")
	purego.RegisterLibFunc(&iotaSDKNative.libCreateWallet, iotaSDK, "create_wallet")
	purego.RegisterLibFunc(&iotaSDKNative.libCreateSecretManager, iotaSDK, "create_secret_manager")

	purego.RegisterLibFunc(&iotaSDKNative.libDestroyClient, iotaSDK, "destroy_client")
	purego.RegisterLibFunc(&iotaSDKNative.libDestroyWallet, iotaSDK, "destroy_wallet")
	purego.RegisterLibFunc(&iotaSDKNative.libDestroySecretManager, iotaSDK, "destroy_secret_manager")
	purego.RegisterLibFunc(&iotaSDKNative.libDestroyString, iotaSDK, "destroy_string")

	purego.RegisterLibFunc(&iotaSDKNative.libGetClientFromWallet, iotaSDK, "get_client_from_wallet")
	purego.RegisterLibFunc(&iotaSDKNative.libGetSecretManagerFromWallet, iotaSDK, "get_secret_manager_from_wallet")

	purego.RegisterLibFunc(&iotaSDKNative.libCallClientMethod, iotaSDK, "call_client_method")
	purego.RegisterLibFunc(&iotaSDKNative.libCallWalletMethod, iotaSDK, "call_wallet_method")
	purego.RegisterLibFunc(&iotaSDKNative.libCallSecretManagerMethod, iotaSDK, "call_secret_manager_method")

	purego.RegisterLibFunc(&iotaSDKNative.libGetLastError, iotaSDK, "binding_get_last_error")

	return &iotaSDKNative, nil
}

func (i *IOTASDK) Destroy() error {
	return lib_loader.UnloadLibrary(i.handle)
}

func (i *IOTASDK) GetLastError() error {
	result := i.libGetLastError()
	if result == "" {
		return nil
	}

	return errors.New(result)
}

func (i *IOTASDK) InitLogger(loggerConfig types.ILoggerConfig) (bool, error) {
	msg, err := serialize(loggerConfig)
	if err != nil {
		return false, err
	}

	if !i.libInitLogger(msg) {
		return false, i.GetLastError()
	}

	return true, nil
}

func (i *IOTASDK) CreateClient(clientOptions types.ClientOptions) (clientPtr IotaClientPtr, err error) {
	msg, err := serialize(clientOptions)
	if err != nil {
		return 0, err
	}

	if clientPtr = i.libCreateClient(msg); clientPtr == 0 {
		return 0, i.GetLastError()
	}

	return clientPtr, nil
}

func (i *IOTASDK) CreateWallet(walletOptions types.WalletOptions) (wallet *Wallet, err error) {
	msg, err := serialize(walletOptions)
	if err != nil {
		return nil, err
	}

	var walletPtr IotaWalletPtr
	if walletPtr = i.libCreateWallet(msg); walletPtr == 0 {
		return nil, i.GetLastError()
	}

	clientPtr, err := i.GetClientFromWallet(walletPtr)
	if err != nil {
		return nil, err
	}

	secretManagerPtr, err := i.GetSecretManagerFromWallet(walletPtr)
	if err != nil {
		return nil, err
	}

	return NewWallet(i, walletPtr, clientPtr, secretManagerPtr), nil
}

func (i *IOTASDK) CreateSecretManager(secretManagerOptions types.WalletOptionsSecretManager) (secretManagerPtr IotaSecretManagerPtr, err error) {
	msg, err := serialize(secretManagerOptions)
	if err != nil {
		return 0, err
	}

	if secretManagerPtr = i.libCreateSecretManager(msg); secretManagerPtr == 0 {
		return 0, i.GetLastError()
	}

	return secretManagerPtr, nil
}

func (i *IOTASDK) GetClientFromWallet(iotaWalletPtr IotaWalletPtr) (clientPtr IotaClientPtr, err error) {
	if clientPtr = i.libGetClientFromWallet(iotaWalletPtr); clientPtr == 0 {
		return 0, i.GetLastError()
	}

	return clientPtr, nil
}

func (i *IOTASDK) GetSecretManagerFromWallet(iotaWalletPtr IotaWalletPtr) (secretManagerPtr IotaSecretManagerPtr, err error) {
	if secretManagerPtr = i.libGetSecretManagerFromWallet(iotaWalletPtr); secretManagerPtr == 0 {
		return 0, i.GetLastError()
	}

	return secretManagerPtr, nil
}

func (i *IOTASDK) CallClientMethod(iotaClientPtr IotaClientPtr, method any) (response string, err error) {
	msg, err := serialize(method)
	if err != nil {
		return "", err
	}

	if response = i.libCallClientMethod(iotaClientPtr, msg); response == "" {
		return "", i.GetLastError()
	}

	return i.CopyAndDestroyString(&response)
}

func (i *IOTASDK) CallWalletMethod(iotaWalletPtr IotaWalletPtr, method any) (response string, err error) {
	msg, err := serialize(method)
	if err != nil {
		return "", err
	}

	if response = i.libCallWalletMethod(iotaWalletPtr, msg); response == "" {
		return "", i.GetLastError()
	}

	return i.CopyAndDestroyString(&response)
}

func (i *IOTASDK) CallSecretManagerMethod(iotaSecretManagerPtr IotaSecretManagerPtr, method any) (response string, err error) {
	msg, err := serialize(method)
	if err != nil {
		return "", err
	}

	if response = i.libCallSecretManagerMethod(iotaSecretManagerPtr, msg); response == "" {
		return "", i.GetLastError()
	}

	return i.CopyAndDestroyString(&response)
}

func (i *IOTASDK) DestroyClient(client IotaClientPtr) (err error) {
	if success := i.libDestroyClient(client); !success {
		return i.GetLastError()
	}

	return nil
}

func (i *IOTASDK) DestroyWallet(client IotaWalletPtr) (err error) {
	if success := i.libDestroyWallet(client); !success {
		return i.GetLastError()
	}

	return nil
}

func (i *IOTASDK) DestroySecretManager(client IotaSecretManagerPtr) (err error) {
	if success := i.libDestroySecretManager(client); !success {
		return i.GetLastError()
	}

	return nil
}

func (i *IOTASDK) CopyAndDestroyString(response *string) (string, error) {
	clonedResponse := strings.Clone(*response)

	if err := i.DestroyString(unsafe.Pointer(response)); err != nil {
		return "", err
	}

	return clonedResponse, nil
}

func (i *IOTASDK) DestroyString(ptr unsafe.Pointer) (err error) {
	// TODO: Properly free strings received from the native lib
	// As this lib is currently only used in the cli, it does not cause a bigger build up, it's not great nonetheless.
	// No security related strings (seed, etc) are ever reaching this library.
	// return nil

	if success := i.libDestroyString(ptr); !success {
		return i.GetLastError()
	}

	return nil
}
