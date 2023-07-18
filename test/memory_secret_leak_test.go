package test

import (
	"bytes"
	"encoding/json"
	"os"
	"runtime/debug"
	"testing"

	"github.com/awnumar/memguard"
	"github.com/stretchr/testify/require"

	wasp_wallet_sdk "github.com/iotaledger/wasp-wallet-sdk"
)

/**
	(Run these test one-by-one, all at once will break the tests due to the nature of memory leaks and counting them)

	This test validates protected JSON serialization and freeing of memory, to make sure that no secret is left in memory.
	This test does the following:

Step 1:
	* Serialize the TestObject (1: with original golang/json marshal and 2: goccy/go-json + memguard that is used in the custom method SerializeGuarded)
		* For the SerializeGuarded function call `free()`
		* For the original marshal function call `memguard.ScrambleBytes()` to invalidate the returned bytes
	* Dump and save the process memory

Step 2:
	* Marshal the TestObject again
	* Read process memory dump from Step 1
    * Scan memory and count the occurrences of the serialized json message

Reason:

The original Golang JSON marshal function leaks memory and allocates the marshaled result twice. One for the returned result, one for the internal state pool.
While the return value can be zeroed/scrambled after it has been used, there is no option to free the copy in the state pool.

This is a problem when memory scanners read the memory of an application that handles secrets/private keys/seeds.

The SerializeGuarded function used in this wrapper uses `go-json` and `memguard`. The go-json marshal function streams the message into a memguard enclave which is internally encrypted and protected
The go-json marshal function does not create a copy of the serialized message internally as there is no state pool.
The usage of SerializeGuarded requires `free` to be called eventually, to get rid of the memguard allocation.

If `free` is called, the actual hits in memory should therefore be **0**, if free is not called, the hits should be exactly **1**
*/

type TestObject struct {
	FINDME int
}

var HEAPDUMP_PATH = ""

var testObject = TestObject{
	FINDME: 9876431264789234,
}

func createHeapDump(t *testing.T, name string) string {
	f, err := os.CreateTemp("", name)
	defer f.Close()
	require.NoError(t, err)
	debug.WriteHeapDump(f.Fd())
	return f.Name()
}

func readHeapDump(t *testing.T, path string) []byte {
	buffer, err := os.ReadFile(path)
	require.NoError(t, err)
	return buffer
}

func serializeUnsafe[T any](t *testing.T, obj T) []byte {
	str, err := json.Marshal(obj)
	require.NoError(t, err)
	return str
}

//nolint:gocritic
func testMemoryLeakUnsafeAction(t *testing.T) {
	serialized := serializeUnsafe(t, testObject)

	// Randomize serialized buffer immediately
	memguard.ScrambleBytes(serialized)

	HEAPDUMP_PATH = createHeapDump(t, "unsafe_test")
}

func testMemoryLeakSafeAction(callFree bool) func(t *testing.T) {
	return func(t *testing.T) {
		_, free, err := wasp_wallet_sdk.SerializeGuarded(testObject)
		require.NoError(t, err)

		if callFree {
			free()
		}

		HEAPDUMP_PATH = createHeapDump(t, "safe_test")
	}
}

func findMemoryLeaks(t *testing.T) int {
	require.Greater(t, len(HEAPDUMP_PATH), 0)
	buffer := readHeapDump(t, HEAPDUMP_PATH)

	// Recreate the serialized message created a step before
	target := serializeUnsafe(t, testObject)

	// Count the occurrences in the memory dump
	hits := bytes.Count(buffer, target)

	return hits
}

func TestMemoryLeakUnsafe(t *testing.T) {
	t.Run("Run unsafe serialization", testMemoryLeakUnsafeAction)
	t.Run("Find memory leaks", func(t *testing.T) {
		require.Equal(t, 1, findMemoryLeaks(t))
	})
}

func TestMemoryLeakProtectionWithoutFree(t *testing.T) {
	t.Run("Run guarded serialization without calling free", testMemoryLeakSafeAction(false))
	t.Run("Find memory leaks", func(t *testing.T) {
		require.Equal(t, 1, findMemoryLeaks(t))
	})
}

func TestMemoryLeakProtectionWithFree(t *testing.T) {
	t.Run("Run guarded serialization with calling free", testMemoryLeakSafeAction(true))
	t.Run("Find memory leaks", func(t *testing.T) {
		require.Equal(t, 0, findMemoryLeaks(t))
	})
}
