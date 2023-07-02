package test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateMnemonic(t *testing.T) {
	sdk := GetOrInitTest(t)

	mnemonic, err := sdk.Utils().GenerateMnemonic()
	defer sdk.Destroy()

	require.NoError(t, err)
	require.NotNil(t, mnemonic)
	require.Len(t, strings.Split(*mnemonic, " "), 24)
}
