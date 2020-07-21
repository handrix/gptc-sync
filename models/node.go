package models

type Block struct {
	Model

	Miner            string `json:"miner"`
	Hash             string `json:"hash"`
	Number           string `json:"number"`
	Difficulty       string `json:"difficulty"`
	ExtraData        string `json:"extra_data"`
	GasLimit         string `json:"gas_limit"`
	GasUsed          string `json:"gas_used"`
	MixHash          string `json:"mix_hash"`
	Nonce            string `json:"nonce"`
	NumberInt        int64  `json:"number_int"`
	ParentHash       string `json:"parent_hash"`
	ReceiptsRoot     string `json:"receipts_root"`
	Size             string `json:"size"`
	StateRoot        string `json:"state_root"`
	Timestamp        string `json:"timestamp"`
	TotalDifficulty  string `json:"total_difficulty"`
	TransactionsRoot string `json:"transactions_root"`
}

func AddBlock(data map[string]interface{}) error {
	db.Create(&Block{
		Miner:            data["miner"].(string),
		Hash:             data["hash"].(string),
		Number:           data["number"].(string),
		Difficulty:       data["difficulty"].(string),
		ExtraData:        data["extraData"].(string),
		GasLimit:         data["gasLimit"].(string),
		GasUsed:          data["gasUsed"].(string),
		MixHash:          data["mixHash"].(string),
		Nonce:            data["nonce"].(string),
		NumberInt:        data["numberInt"].(int64),
		ParentHash:       data["parentHash"].(string),
		ReceiptsRoot:     data["receiptsRoot"].(string),
		Size:             data["size"].(string),
		StateRoot:        data["stateRoot"].(string),
		Timestamp:        data["timestamp"].(string),
		TotalDifficulty:  data["totalDifficulty"].(string),
		TransactionsRoot: data["transactionsRoot"].(string),
	})
	return nil
}

func MaxIndex() int64 {
	var count int64
	db.Model(&Block{}).Count(&count)
	return count
}

func GetBlock(number int64) (block []Block) {
	db.Where("number_int = ?", number).First(&block)
	return
}
