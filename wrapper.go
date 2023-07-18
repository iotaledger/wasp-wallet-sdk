package wasp_wallet_sdk

import (
	"errors"

	"github.com/awnumar/memguard"
	"github.com/goccy/go-json"

	"github.com/iotaledger/wasp-wallet-sdk/lib_loader"
	"github.com/iotaledger/wasp-wallet-sdk/types"

	"github.com/ebitengine/purego"
)

type (
	IotaClientPtr        uintptr
	IotaWalletPtr        uintptr
	IotaSecretManagerPtr uintptr

	ProtectedStringPtr *byte
)

type IOTASDK struct {
	handle uintptr

	libInitLogger func(*byte) bool

	libCreateClient        func(*byte) IotaClientPtr
	libCreateWallet        func(*byte) IotaWalletPtr
	libCreateSecretManager func(*byte) IotaSecretManagerPtr

	libDestroyClient        func(IotaClientPtr) bool
	libDestroyWallet        func(IotaWalletPtr) bool
	libDestroySecretManager func(IotaSecretManagerPtr) bool
	libDestroyString        func(uintptr) bool

	libGetClientFromWallet        func(IotaWalletPtr) IotaClientPtr
	libGetSecretManagerFromWallet func(IotaWalletPtr) IotaSecretManagerPtr

	libCallClientMethod        func(IotaClientPtr, *byte) uintptr
	libCallWalletMethod        func(IotaWalletPtr, *byte) uintptr
	libCallSecretManagerMethod func(IotaSecretManagerPtr, *byte) uintptr
	libCallUtilsMethod         func(*byte) uintptr

	libGetLastError func() string
}

// Encodes an object into a JSON string protected by memguard
// Uses an alternative JSON library to mitigate hidden copies of the serialized message.
func SerializeGuarded(obj any) (ProtectedStringPtr, func(), error) {
	stream := memguard.NewStream()

	err := json.NewEncoder(stream).Encode(obj)
	if err != nil {
		return nil, func() {}, err
	}

	preBuffer, err := stream.Flush()
	if err != nil {
		return nil, func() {}, err
	}

	enclave := preBuffer.Seal()
	buf, err := enclave.Open()
	if err != nil {
		return nil, func() {}, err
	}

	strPointer, cfree := CStringGo(buf.Data())

	return strPointer, func() {
		buf.Destroy()
		cfree()
	}, nil
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
	purego.RegisterLibFunc(&iotaSDKNative.libCallUtilsMethod, iotaSDK, "call_utils_method")

	purego.RegisterLibFunc(&iotaSDKNative.libGetLastError, iotaSDK, "binding_get_last_error")

	return &iotaSDKNative, nil
}

func (i *IOTASDK) Utils() *Utils {
	return &Utils{sdk: i}
}

func (i *IOTASDK) Destroy() {
	_ = lib_loader.UnloadLibrary(i.handle)
}

func (i *IOTASDK) GetLastError() error {
	result := i.libGetLastError()
	if result == "" {
		return nil
	}

	return errors.New(result)
}

func (i *IOTASDK) InitLogger(loggerConfig types.ILoggerConfig) (bool, error) {
	msg, free, err := SerializeGuarded(loggerConfig)
	defer free()
	if err != nil {
		return false, err
	}

	if !i.libInitLogger(msg) {
		return false, i.GetLastError()
	}

	return true, nil
}

func (i *IOTASDK) CreateClient(clientOptions types.ClientOptions) (clientPtr IotaClientPtr, err error) {
	msg, free, err := SerializeGuarded(clientOptions)
	defer free()
	if err != nil {
		return 0, err
	}

	if clientPtr = i.libCreateClient(msg); clientPtr == 0 {
		return 0, i.GetLastError()
	}

	return clientPtr, nil
}

func (i *IOTASDK) CreateSecretManager(secretManagerOptions types.WalletOptionsSecretManager) (secretManagerPtr IotaSecretManagerPtr, err error) {
	bytes, free, err := SerializeGuarded(secretManagerOptions)
	defer free()
	if err != nil {
		return 0, err
	}

	if secretManagerPtr = i.libCreateSecretManager(bytes); secretManagerPtr == 0 {
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

func (i *IOTASDK) CallUtilsMethod(method any) (response []byte, free func(), err error) {
	msg, free, err := SerializeGuarded(method)
	defer free()
	if err != nil {
		return nil, func() {}, err
	}

	var responsePtr uintptr
	if responsePtr = i.libCallUtilsMethod(msg); responsePtr == 0 {
		return nil, func() {}, i.GetLastError()
	}

	return i.CopyAndDestroyOriginalStringPtr(responsePtr)
}

func (i *IOTASDK) CallClientMethod(iotaClientPtr IotaClientPtr, method any) (response []byte, free func(), err error) {
	msg, free, err := SerializeGuarded(method)
	defer free()
	if err != nil {
		return nil, func() {}, err
	}

	var responsePtr uintptr
	if responsePtr = i.libCallClientMethod(iotaClientPtr, msg); responsePtr == 0 {
		return nil, func() {}, i.GetLastError()
	}

	return i.CopyAndDestroyOriginalStringPtr(responsePtr)
}

func (i *IOTASDK) CallWalletMethod(iotaWalletPtr IotaWalletPtr, method any) ([]byte, func(), error) {
	msg, free, err := SerializeGuarded(method)
	defer free()
	if err != nil {
		return nil, func() {}, err
	}

	var responsePtr uintptr
	if responsePtr = i.libCallWalletMethod(iotaWalletPtr, msg); responsePtr == 0 {
		return nil, func() {}, i.GetLastError()
	}

	return i.CopyAndDestroyOriginalStringPtr(responsePtr)
}

func (i *IOTASDK) CallSecretManagerMethod(iotaSecretManagerPtr IotaSecretManagerPtr, method any) ([]byte, func(), error) {
	msg, free, err := SerializeGuarded(method)
	defer free()
	if err != nil {
		return nil, func() {}, err
	}

	var responsePtr uintptr
	if responsePtr = i.libCallSecretManagerMethod(iotaSecretManagerPtr, msg); responsePtr == 0 {
		return nil, func() {}, i.GetLastError()
	}

	return i.CopyAndDestroyOriginalStringPtr(responsePtr)
}

func (i *IOTASDK) DestroyClient(client IotaClientPtr) (err error) {
	if client == 0 {
		return nil
	}

	if success := i.libDestroyClient(client); !success {
		return i.GetLastError()
	}

	return nil
}

func (i *IOTASDK) DestroyWallet(wallet IotaWalletPtr) (err error) {
	if wallet == 0 {
		return nil
	}

	if success := i.libDestroyWallet(wallet); !success {
		return i.GetLastError()
	}

	return nil
}

func (i *IOTASDK) DestroySecretManager(secretManager IotaSecretManagerPtr) (err error) {
	if secretManager == 0 {
		return nil
	}

	if success := i.libDestroySecretManager(secretManager); !success {
		return i.GetLastError()
	}

	return nil
}

func (i *IOTASDK) CopyAndDestroyOriginalStringPtr(response uintptr) ([]byte, func(), error) {
	goString, free := GoString(response)

	if err := i.DestroyString(response); err != nil {
		return nil, free, err
	}

	return goString, free, nil
}

func (i *IOTASDK) DestroyString(ptr uintptr) (err error) {
	if success := i.libDestroyString(ptr); !success {
		return i.GetLastError()
	}

	return nil
}
