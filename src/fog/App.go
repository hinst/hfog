package fog

import (
	"encoding/json"
	"encoding/xml"
	"hgo"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

const FogApiUrl = "/api.asp"

type TApp struct {
	Config              *TConfig
	Active              int32
	LoadBugsModeEnabled bool
}

func (this *TApp) Create() *TApp {
	this.Active = 1
	return this
}

func (this *TApp) Run() {
	this.Prepare()
	this.ReadConfig()
	if this.LoadBugsModeEnabled {
		this.RunLoadBugsMode()
	}
}

func (this *TApp) RunLoadBugsMode() {
	var bugs = this.ReadBugs()
	var remainingBugs = this.RemoveExistingBugs(bugs)
	WriteLog("Need bugs: " + strconv.Itoa(len(remainingBugs.Cases.Cases)) + " of " + strconv.Itoa(len(bugs.Cases.Cases)))
	for bugIndex, bug := range remainingBugs.Cases.Cases {
		var data = this.LoadBug(&bug)
		var filePath = this.GetBugFilePath(&bug)
		WriteLog("Now writing " +
			strconv.Itoa(bugIndex) + "/" + strconv.Itoa(len(remainingBugs.Cases.Cases)) +
			"//" + strconv.Itoa(len(bugs.Cases.Cases)) +
			" " + filePath)
		var writeFileResult = ioutil.WriteFile(filePath, data, os.ModePerm)
		AssertResult(writeFileResult)
		time.Sleep(3000 * time.Millisecond)
		if false == this.CheckActive() {
			break
		}
	}
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
	if false {
		WriteLog("Get " + url)
	}
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

func (this *TApp) ReadBugs() *TBugList {
	var bugs TBugList
	var data, readResult = ioutil.ReadFile("data/bugs.xml")
	AssertResult(readResult)
	var decodeResult = xml.Unmarshal(data, &bugs)
	AssertResult(decodeResult)
	return &bugs
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

func (this *TApp) RemoveExistingBugs(a *TBugList) *TBugList {
	var result TBugList
	for _, bug := range a.Cases.Cases {
		if false == this.CheckHasBug(&bug) {
			result.Cases.Cases = append(result.Cases.Cases, bug)
		}
	}
	return &result
}

func (this *TApp) CheckHasBug(bug *TBugListCase) bool {
	var bugFilePath = this.GetBugFilePath(bug)
	if false {
		WriteLog(bugFilePath)
	}
	var _, err = os.Stat(bugFilePath)
	var exists = false == os.IsNotExist(err)
	return exists
}

func (this *TApp) GetBugFilePath(bug *TBugListCase) string {
	return "data/" + bug.IxBug + ".xml"
}

func (this *TApp) LoadBug(bug *TBugListCase) []byte {
	var url = this.GetURL() +
		"?token=" + url.QueryEscape(this.Config.Token) +
		"&cmd=search" +
		"&q=" + url.QueryEscape(bug.IxBug) +
		"&cols=events,sTitle"
	return this.Get(url)
}

func (this *TApp) CheckActive() bool {
	return atomic.LoadInt32(&this.Active) > 0
}

func (this *TApp) SetActive(v bool) {
	if v {
		atomic.StoreInt32(&this.Active, 1)
	} else {
		atomic.StoreInt32(&this.Active, 0)
	}
}

func (this *TApp) Prepare() {
	hgo.InstallShutdownReceiver(
		func() {
			this.SetActive(false)
		})
}
