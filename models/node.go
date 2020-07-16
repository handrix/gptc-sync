package models

type Node struct {
	ID               int    `gorm:"primary_key" json:"id"`
	Miner            string `json:"miner"`
	Hash             string `json:"hash"`
	Number           string `json:"number"`
	Difficulty       string `json:"difficulty"`
	ExtraData        string `json:"extra_data"`
	GasLimit         string `json:"gas_limit"`
	GasUsed          string `json:"gas_used"`
	MixHash          string `json:"mix_hash"`
	Nonce            string `json:"nonce"`
	NumberInt        string `json:"number_int"`
	ParentHash       string `json:"parent_hash"`
	ReceiptsRoot     string `json:"receipts_root"`
	Sha3Uncles       string `json:"sha_3_uncles"`
	Size             string `json:"size"`
	StateRoot        string `json:"state_root"`
	Timestamp        string `json:"timestamp"`
	TotalDifficulty  string `json:"total_difficulty"`
	TransactionsRoot string `json:"transactions_root"`
}

func AddBlock(data map[string]interface{}) error {
	db.Create(&Node{
		Miner:            data["miner"].(string),
		Hash:             data["hash"].(string),
		Number:           data["number"].(string),
		Difficulty:       data["difficulty"].(string),
		ExtraData:        data["extra_data"].(string),
		GasLimit:         data["gas_limit"].(string),
		GasUsed:          data["gas_used"].(string),
		MixHash:          data["gas_used"].(string),
		Nonce:            data["nonce"].(string),
		NumberInt:        data["number_int"].(string),
		ParentHash:       data["parent_hash"].(string),
		ReceiptsRoot:     data["receipts_root"].(string),
		Sha3Uncles:       data["sha_3_uncles"].(string),
		Size:             data["size"].(string),
		StateRoot:        data["state_root"].(string),
		Timestamp:        data["timestamp"].(string),
		TotalDifficulty:  data["total_difficulty"].(string),
		TransactionsRoot: data["transactions_root"].(string),
	})
	return nil
}
