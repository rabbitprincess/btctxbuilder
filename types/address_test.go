package types

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPubKeyToAddr(t *testing.T) {
	network := BTC_Testnet3
	pubKeyHex := "0357bbb2d4a9cb8a2357633f201b9c518c2795ded682b7913c6beef3fe23bd6d2f"
	publicKey, err := hex.DecodeString(pubKeyHex)
	assert.NoError(t, err)

	p2pkh, err := PubKeyToAddr(publicKey, P2PKH, network)
	assert.NoError(t, err)
	assert.Equal(t, "mouQtmBWDS7JnT65Grj2tPzdSmGKJgRMhE", p2pkh.EncodeAddress())

	p2wpkh, err := PubKeyToAddr(publicKey, P2WPKH, network)
	assert.NoError(t, err)
	assert.Equal(t, "tb1qtsq9c4fje6qsmheql8gajwtrrdrs38kdzeersc", p2wpkh.EncodeAddress())

	p2sh, err := PubKeyToAddr(publicKey, P2WPKH_NESTED, network)
	assert.NoError(t, err)
	assert.Equal(t, "2NF33rckfiQTiE5Guk5ufUdwms8PgmtnEdc", p2sh.EncodeAddress())

	p2tr, err := PubKeyToAddr(publicKey, TAPROOT, network)
	assert.NoError(t, err)
	assert.Equal(t, "tb1pklh8lqax5l7m2ycypptv2emc4gata2dy28svnwcp9u32wlkenvsspcvhsr", p2tr.EncodeAddress())
}
