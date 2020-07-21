package syncer

import (
	"encoding/json"
	"fmt"
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

type BlockMethod struct{}

func (b *BlockMethod) CurrentBlock() uint32 {
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

func (b *BlockMethod) BlockDetail(number int64) map[string]interface{} {
	var dat map[string]interface{}

	client := &http.Client{}
	// format payload
	formatPayload := setting.BlockDetailParams
	hex := strconv.FormatInt(number, 16)
	formatPayload = fmt.Sprintf(formatPayload, "0x"+hex)
	fmt.Println(hex)
	fmt.Println(formatPayload)
	payload := strings.NewReader(formatPayload)

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

	result := dat["result"].(map[string]interface{})
	result["numberInt"] = number
	return result

}
