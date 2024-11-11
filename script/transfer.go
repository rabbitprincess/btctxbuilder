package script

import (
	"errors"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/txscript"
	"github.com/rabbitprincess/btctxbuilder/types"
)

// SingleSig
func EncodeTransferScript(network types.Network, address string) ([]byte, error) {
	addr, err := btcutil.DecodeAddress(address, types.GetParams(network))
	if err != nil {
		return nil, err
	}

	return txscript.PayToAddrScript(addr)
}

func DecodeTransferScript(script []byte) (btcutil.Address, error) {
	_, addresses, _, err := txscript.ExtractPkScriptAddrs(script, types.GetParams(types.BTC))
	if err != nil {
		return nil, err
	}
	// only one address is expected
	if len(addresses) != 1 {
		return nil, errors.New("invalid script address")
	}

	return addresses[0], nil
}
