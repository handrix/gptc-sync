package syncer

import (
	"encoding/json"
	"gptc-sync/pkg/setting"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var (
	url = setting.Domain
)

type NodeDetail struct {
	miner  string
	hash   string
	number string
}

func CurrentBlock() uint32 {
	var dat map[string]interface{}

	client := &http.Client{}
	payload := strings.NewReader(setting.CurrentBlockParams)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		panic("[getBlock] Do request error")
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	// json解析
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal([]byte(body), &dat)
	// 16进制转换
	hex := dat["result"].(string)
	val := hex[2:]
	n, err := strconv.ParseUint(val, 16, 32)

	if err != nil {
		panic("[getBlock] Parser error")
	}

	currentBlock := uint32(n)
	return currentBlock
}

func BlockDetail() *NodeDetail {
	var dat map[string]interface{}

	client := &http.Client{}
	payload := strings.NewReader(setting.BlockDetailParams)
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	// json解析
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal([]byte(body), &dat)
	// 提取数据
	result := dat["result"].(map[string]interface{})
	// data
	miner := result["miner"].(string)
	hash := result["hash"].(string)
	number := result["number"].(string)

	nodeDetail := &NodeDetail{miner: miner, hash: hash, number: number}
	return nodeDetail

}
