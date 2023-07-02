package test

import (
	"testing"

	"github.com/iotaledger/wasp-wallet-sdk/types"

	"github.com/stretchr/testify/require"
)

func TestFirstCon(t *testing.T) {
	sdk := GetOrInitTest(t)

	clientPtr, err := sdk.CreateClient(types.ClientOptions{
		PrimaryNode: ShimmerNetworkAPI,
		Nodes:       []interface{}{ShimmerNetworkAPI},
	})
	require.NoError(t, err)
	require.NotNil(t, clientPtr)

	response, err := sdk.CallClientMethod(clientPtr, "RUBBISH")
	require.Empty(t, response)
	require.Error(t, err)

	t.Log(clientPtr)
}
