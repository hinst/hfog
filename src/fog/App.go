package fog

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)

type TApp struct {
	RootURL string
}

func (this *TApp) Create() *TApp {
	return this
}

func (this *TApp) Run() {
	this.ReadConfig()
}

func (this *TApp) ReadConfig() {
	var config = (&TConfig{}).Create()
	var data, readFileResult = ioutil.ReadFile("config.json")
	AssertResult(readFileResult)
	var decodeResult = json.Unmarshal(data, config)
	AssertResult(decodeResult)
	this.RootURL = config.RootURL
	WriteLog("RootURL: " + this.RootURL)
}

func (this *TApp) Login() {
	var url = url.QueryEscape(cmd=logon&email=xxx@example.com&password=BigMac	
}
