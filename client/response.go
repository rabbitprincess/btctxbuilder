package client

type Block struct {
	ID                string  `json:"id"`
	Height            int     `json:"height"`
	Version           int     `json:"version"`
	Timestamp         int     `json:"timestamp"`
	TxCount           int     `json:"tx_count"`
	Size              int     `json:"size"`
	Weight            int     `json:"weight"`
	MerkleRoot        string  `json:"merkle_root"`
	Previousblockhash string  `json:"previousblockhash"`
	Mediantime        int     `json:"mediantime"`
	Nonce             int     `json:"nonce"`
	Bits              int     `json:"bits"`
	Difficulty        float64 `json:"difficulty"`
}

type Transaction struct {
	Txid     string      `json:"txid"`
	Version  int         `json:"version"`
	Locktime int         `json:"locktime"`
	Vin      []Vin       `json:"vin"`
	Vout     []Vout      `json:"vout"`
	Size     int         `json:"size"`
	Weight   int         `json:"weight"`
	Fee      int         `json:"fee"`
	Status   BlockStatus `json:"status"`
}

type Vin struct {
	Txid         string   `json:"txid"`
	Vout         int64    `json:"vout"`
	Prevout      any      `json:"prevout"`
	Scriptsig    string   `json:"scriptsig"`
	ScriptsigAsm string   `json:"scriptsig_asm"`
	Witness      []string `json:"witness"`
	IsCoinbase   bool     `json:"is_coinbase"`
	Sequence     int64    `json:"sequence"`
}

type Vout struct {
	Scriptpubkey        string `json:"scriptpubkey"`
	ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
	ScriptpubkeyType    string `json:"scriptpubkey_type"`
	ScriptpubkeyAddress string `json:"scriptpubkey_address,omitempty"`
	Value               int    `json:"value"`
}

type Utxo struct {
	Txid   string      `json:"txid"`
	Vout   int         `json:"vout"`
	Status BlockStatus `json:"status"`
	Value  int         `json:"value"`
}

type BlockStatus struct {
	Confirmed   bool   `json:"confirmed"`
	BlockHeight int    `json:"block_height"`
	BlockHash   string `json:"block_hash"`
	BlockTime   int    `json:"block_time"`
}
