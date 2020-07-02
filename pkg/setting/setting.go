package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string

	// gtpc settings
	Domain             string
	CurrentBlockParams string
	BlockDetailParams  string
)

func init() {
	var err error
	Cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadGPTC()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadGPTC() {
	sec, err := Cfg.GetSection("gptc")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	Domain = sec.Key("DOMAIN").MustString("https://www.gptcscan.org/rpc")
	CurrentBlockParams = sec.Key("CURRENT_BLOCK_PARAMS").MustString("{\"jsonrpc\":\"2.0\",\"method\":\"eth_blockNumber\",\"params\":[],\"id\":1}")
	BlockDetailParams = sec.Key("BLOCK_DETAIL_PARAMS").MustString("{\"jsonrpc\": \"2.0\",\"id\": 45,\"method\": \"eth_getBlockByNumber\",\"params\": [\"0x1\", false]}")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
