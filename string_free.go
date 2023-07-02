package wasp_wallet_sdk

import "unsafe"

/**
	The purego library can accept proper strings as return values and handle them automatically.
	The invoked Rust function needs to hold any allocated string in memory though,
	so it doesn't get freed after the function finishes. Otherwise, Go would receive null pointers.

	As purego does not provide an option to call a freeing function after the string has been handled, it is duplicated here,
    and instead of strings, uintptrs are used and supplied to the sdks free function `destroy_string`.
*/

// Copied from https://github.com/ebitengine/purego/blob/main/internal/strings/strings.go
// copies a char* to a Go string.
func GoString(c uintptr) string {
	// We take the address and then dereference it to trick go vet from creating a possible misuse of unsafe.Pointer
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&c))
	if ptr == nil {
		return ""
	}
	var length int
	for {
		if *(*byte)(unsafe.Add(ptr, uintptr(length))) == '\x00' {
			break
		}
		length++
	}
	return string(unsafe.Slice((*byte)(ptr), length))
}
