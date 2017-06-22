package fog

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
	this.ReadBugs()
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

func (this *TApp) Login() bool {
	var url = this.GetURL() +
		"?cmd=logon" +
		"&email=" + url.QueryEscape(this.Config.Email) +
		"&password=" + url.QueryEscape(this.Config.Password)
	var resp = this.Get(url)
	var respObj TLoginResponse
	WriteLog(string(resp))
	xml.Unmarshal(resp, &respObj)
	if respObj.Response.Error == "" {
		WriteLog("Logged in successfully; token = " + respObj.Response.Token)
	} else {
		WriteLog("Login failed; error = " + respObj.Response.Error)
	}
	return respObj.Response.Error == ""
}

func (this *TApp) GetURL() string {
	return this.Config.RootURL + FogApiUrl
}

func (this *TApp) GetResponse(url string) *http.Response {
	WriteLog("Get " + url)
	var response, responseResult = http.Get(url)
	AssertResult(responseResult)
	return response
}

func (this *TApp) Get(url string) []byte {
	var resp = this.GetResponse(url)
	var data, readResult = ioutil.ReadAll(resp.Body)
	AssertResult(readResult)
	return data
}

func (this *TApp) ReadBugs() {
	var bugs TBugList
	var data, readResult = ioutil.ReadFile("data/bugs.xml")
	AssertResult(readResult)
	var decodeResult = xml.Unmarshal(data, &bugs)
	AssertResult(decodeResult)
	WriteLog("Number of bugs: " + strconv.Itoa(len(bugs.Cases.Cases)))
}

func (this *TApp) WriteSampleBugs() {
	var bugs = TBugList{
		Cases: TBugListCases{
			Cases: []TBugListCase{
				TBugListCase{IxBug: "1"},
				TBugListCase{IxBug: "2"},
			},
		},
	}
	var data, encodeResult = xml.Marshal(&bugs)
	AssertResult(encodeResult)
	var writeResult = ioutil.WriteFile("bugs.xml", data, os.ModePerm)
	AssertResult(writeResult)
	Unuse(bugs)
}
