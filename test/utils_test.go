package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateMnemonic(t *testing.T) {
	sdk := GetOrInitTest(t)

	mnemonic, err := sdk.Utils().GenerateMnemonic()
	require.NoError(t, err)
	require.NotNil(t, mnemonic)
	defer sdk.Destroy()

	buffer, err := mnemonic.Open()
	require.NoError(t, err)

	mmnemonicStr := buffer.String()
	fmt.Println("STR: " + mmnemonicStr)
	require.Len(t, strings.Split(mmnemonicStr, " "), 24)
}
