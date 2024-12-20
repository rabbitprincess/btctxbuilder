package transaction

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/rabbitprincess/btctxbuilder/client"
	"github.com/rabbitprincess/btctxbuilder/types"
	"github.com/stretchr/testify/require"
)

func TestTransfer(t *testing.T) {
	fromPrivKeyHex := "a6018c89646f3c7596516544602283135e8d6e5b31421e335b91b86ae9c76409"
	fromPrivKey, _ := hex.DecodeString(fromPrivKeyHex)
	fromPubKey := "0248d7c76f23e387bb151e6094590eb8f7777a8efbea9d0a5ddd1ea1833fa3925c"
	_ = fromPubKey
	fromAddress := "n368zCWREFiRRX7icJRBb6n8nMsjJjNVK8"

	toAddress := "miNCoTMLbp5Y6eQJ9ycH3xDic3oiKqUtv9"
	var toAmount int64 = 1000

	net := types.BTC_Signet
	btcclient := client.NewClient(net)
	psbtPacket, err := NewTransferTx(btcclient, fromAddress, map[string]int64{toAddress: toAmount}, fromAddress)
	require.NoError(t, err)

	var buf bytes.Buffer
	err = psbtPacket.Serialize(&buf)
	require.NoError(t, err)
	psbtRaw := buf.Bytes()

	signedTxRaw, err := SignTx(net, psbtRaw, fromPrivKey)
	require.NoError(t, err)

	signedTxHex := hex.EncodeToString(signedTxRaw)
	fmt.Println(signedTxHex)

	// txid, err := btcclient.BroadcastTx(signedTxHex)
	// require.NoError(t, err)
	// fmt.Println("txid:", txid)

	newTx, err := client.DecodeRawTransaction(signedTxRaw)
	require.NoError(t, err)

	jsonNewTx, _ := json.MarshalIndent(newTx, "", "\t")
	fmt.Println(string(jsonNewTx))
}
