package test

import (
	"bytes"
	"encoding/json"
	"os"
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/require"

	wasp_wallet_sdk "github.com/iotaledger/wasp-wallet-sdk"
)

/**
	(Run these test one-by-one, all at once will break the tests due to the nature of memory leaks and counting them)

	To validate that the serialization and freeing of memory works, this test does the following:

Step 1:
	* Serialize the TestObject (with original golang/json marshal and goccy/go-json +(memguard) as SerializeGuarded)
		* For the SerializeGuarded call `free()`
	* Dump and save the process memory

Step 2:
	* Read process memory
    * Scan memory for the serialized json message

Reason:

The original Golang JSON marshal function leaks memory and currently leaves **one** copy of the serialized message in memory.

There is a second copy of data, which is the returned data that does not automatically free after using.
While this return value can be zeroed, there is no option to free the copy created by json.marshal internally.

This is a problem when memory scanners read the memory of an application handling secrets/private keys/seeds.

The SerializeGuarded function uses `go-json` and `memguard`. The go-json marshal function streams the message into a memguard enclave which is internally encrypted and protected
It does not create a copy of the serialized message internally.
The usage of SerializeGuarded requires `free` to get called eventually, to get rid of all copies.

If `free` is called, the hits in memory should therefore be **0**, if free is not called, the hits should be **1**
*/

func serializeUnsafe[T any](t *testing.T, obj T) []byte {
	str, err := json.Marshal(obj)
	require.NoError(t, err)
	return str
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

type TestObject struct {
	FINDME int
}

var testObject = TestObject{
	FINDME: 9876431264789234,
}

var HEAPDUMP_PATH = ""

func findMemoryLeaks(t *testing.T) int {
	require.Greater(t, len(HEAPDUMP_PATH), 0)
	buffer := readHeapDump(t, HEAPDUMP_PATH)

	target := serializeUnsafe(t, testObject)
	hits := bytes.Count(buffer, target)

	return hits
}

func testMemoryLeakUnsafeAction(t *testing.T) {
	_ = serializeUnsafe(t, testObject)
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

func TestMemoryLeakUnsafeJSONNoMemguard(t *testing.T) {
	t.Run("Run unsafe serialization", testMemoryLeakUnsafeAction)
	t.Run("Find memory leaks", func(t *testing.T) {
		require.Equal(t, findMemoryLeaks(t), 2)
	})
}

func TestMemoryLeakWithoutFree(t *testing.T) {
	t.Run("Run guarded serialization without calling free", testMemoryLeakSafeAction(false))
	t.Run("Find memory leaks", func(t *testing.T) {
		require.Equal(t, findMemoryLeaks(t), 1)
	})
}

func TestMemoryLeakWithFree(t *testing.T) {
	t.Run("Run guarded serialization with calling free", testMemoryLeakSafeAction(true))
	t.Run("Find memory leaks", func(t *testing.T) {
		require.Equal(t, findMemoryLeaks(t), 0)
	})
}
