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
		ExtraData:        data["difficulty"].(string),
		GasLimit:         data["difficulty"].(string),
		GasUsed:          data["difficulty"].(string),
		MixHash:          data["difficulty"].(string),
		Nonce:            data["difficulty"].(string),
		NumberInt:        data["difficulty"].(string),
		ParentHash:       data["difficulty"].(string),
		ReceiptsRoot:     data["difficulty"].(string),
		Sha3Uncles:       data["difficulty"].(string),
		Size:             data["difficulty"].(string),
		StateRoot:        data["difficulty"].(string),
		Timestamp:        data["difficulty"].(string),
		TotalDifficulty:  data["difficulty"].(string),
		TransactionsRoot: data["difficulty"].(string),
	})
	return nil
}