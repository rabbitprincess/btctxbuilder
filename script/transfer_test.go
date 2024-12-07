package script

import (
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/rabbitprincess/btctxbuilder/types"
	"github.com/stretchr/testify/require"
)

func TestEncodeTransferScript(t *testing.T) {
	for _, test := range []struct {
		addrType types.AddrType
		network  types.Network
	}{
		// from public key
		{types.P2PKH, types.BTC_Testnet3},
		{types.P2WPKH, types.BTC_Testnet3},
		{types.P2WPKH_NESTED, types.BTC_Testnet3},
		{types.TAPROOT, types.BTC_Testnet3},
	} {
		// make params and address
		param := types.GetParams(test.network)
		privKey, err := btcec.NewPrivateKey()
		require.NoError(t, err)
		pubKey := privKey.PubKey().SerializeUncompressed()
		address, err := types.PubKeyToAddr(pubKey, test.addrType, test.network)
		require.NoError(t, err)

		// encode address to script
		script, err := EncodeTransferScript(address)
		require.NoError(t, err)

		// decode script to address
		decodedAddress, err := DecodeTransferScript(script, param)
		require.NoError(t, err)
		require.Equal(t, address.EncodeAddress(), decodedAddress.EncodeAddress())
	}
}
