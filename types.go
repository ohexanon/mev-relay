package main

type Payload struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
}

type SendBundlePayload struct {
	JSONRPC string       `json:"jsonrpc"`
	ID      int          `json:"id"`
	Method  string       `json:"method"`
	Params  BundleParams `json:"params"`
}

type BundleParams struct {
	TXs               []string `json:"txs"`
	BlockNumber       string   `json:"blockNumber"`
	MinTimestamp      uint64   `json:"minTimestamp"`
	MaxTimestamp      uint64   `json:"maxTimestamp"`
	RevertingTXHashes []string `json:"revertingTxHashes"`
}
