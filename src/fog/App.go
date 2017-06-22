package fog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const FogApiUrl = "/api.asp"

type TApp struct {
	Config *TConfig
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
	this.Config = config
	WriteLog("RootURL: " + this.Config.RootURL)
}

func (this *TApp) Login() {
	var url = "?cmd=logon" +
		"&email=" + url.QueryEscape(this.Config.Email) +
		"&password=" + url.QueryEscape(this.Config.Password)
	var resp = this.Get(url)
}

func (this *TApp) GetURL() string {
	return this.Config.RootURL + FogApiUrl
}

func (this *TApp) Get(url string) *http.Response {
	WriteLog("Get " + url)
	var response, responseResult = http.Get(url)
	AssertResult(responseResult)
	return response
}
