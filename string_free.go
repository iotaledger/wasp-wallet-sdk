package wasp_wallet_sdk

import (
	"bytes"
	"unsafe"

	"github.com/awnumar/memguard"
)

/**
	The purego library can accept proper strings as return values and handle them automatically.
	The invoked Rust function needs to hold any allocated string in memory though,
	so it doesn't get freed after the function finishes. Otherwise, Go would receive null pointers.

	As purego does not provide an option to call a freeing function after the string has been handled, it is duplicated here,
    and instead of strings, uintptrs are used and supplied to the sdks free function `destroy_string`.
*/

//go:noinline
func CStringGo(str []byte) (*byte, func()) {
	var b []byte
	if len(str) == 0 {
		b = []byte{0x0}
	} else if bytes.HasSuffix(str, []byte{0x0}) {
		b = make([]byte, len(str))
		copy(b, str)
	} else {
		b = make([]byte, len(str)+1)
		copy(b, str)
	}
	return &b[0], func() {
		memguard.ScrambleBytes(b)
		memguard.WipeBytes(b)
	}
}

// Copied from https://github.com/ebitengine/purego/blob/main/internal/strings/strings.go
// copies a char* to a Go string.
func GoString(c uintptr) ([]byte, func()) {
	// We take the address and then dereference it to trick go vet from creating a possible misuse of unsafe.Pointer
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&c))
	if ptr == nil {
		return nil, func() {

		}
	}
	var length int
	for {
		if *(*byte)(unsafe.Add(ptr, uintptr(length))) == '\x00' {
			break
		}
		length++
	}
	byteSlice := unsafe.Slice((*byte)(ptr), length)

	// Until now, any data referenced is a pointer to the allocation inside the Rust library
	// To make use of the data, do a copy and make it possible to free it properly afterwords
	// It is mandatory to call the returned `free()` function to properly dispose it.
	goString := make([]byte, length)
	copy(goString, byteSlice)

	return goString, func() {
		memguard.ScrambleBytes(goString)
	}
}
