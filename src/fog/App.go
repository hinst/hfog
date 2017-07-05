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
	"strings"
	"sync/atomic"
	"time"
)

const FogApiUrl = "/api.asp"

type TApp struct {
	Config *TConfig
	Active int32
	DB     *TDBMan

	LoadBugsModeEnabled       bool
	AttachmentsModeEnabled    bool
	AttachmentTestModeEnabled bool
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
	if this.AttachmentsModeEnabled {
		this.RunAttachmentsMode()
	}
	if this.AttachmentTestModeEnabled {
		this.RunAttachmentTestMode()
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
	WriteLog(IntToStr(len(data)))
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

func (this *TApp) LoadAttachment(aurl string) []byte {
	aurl = this.Config.RootURL + "/" + strings.Replace(aurl, "&amp;", "&", -1) +
		"&token=" + url.QueryEscape(this.Config.Token)
	return this.Get(aurl)
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

func (this *TApp) RunAttachmentsMode() {
	this.DB = (&TDBMan{}).Create()
	this.DB.FilePath = "data/db-attachments.bolt"
	this.DB.Start()
	var bugListData = ReadBugsFromFile(hgo.AppDir + "/data/bugs.xml")
	var countOfAttachments = 0
	for _, item := range bugListData.Cases.Cases {
		if false == this.CheckActive() {
			break
		}
		var data = this.ReadBugData(item.IxBug)
		for _, event := range data.Cases.Cases[0].Events.Events {
			for _, attachment := range event.RGAttachments.Attachments {
				if strings.HasSuffix(attachment.SFileName.Text, ".doc") || strings.HasSuffix(attachment.SFileName.Text, ".docx") {
					countOfAttachments++
					if this.GrabAttachmentIfNecess(attachment) {
						WriteLog(IntToStr(countOfAttachments) + " " + attachment.SFileName.Text + " " + attachment.SURL.Text)
						time.Sleep(3 * time.Second)
					}
				}
			}
		}
	}
	this.DB.Stop()
}

func (this *TApp) ReadBugData(id string) (result *TBugData) {
	var data, readFileResult = ioutil.ReadFile(hgo.AppDir + "/data/" + id + ".xml")
	WriteLogResult(readFileResult)
	if readFileResult == nil {
		result = &TBugData{}
		var unmarshalResult = xml.Unmarshal(data, result)
		WriteLogResult(unmarshalResult)
		if nil != unmarshalResult {
			result = nil
		}
	}
	return
}

func (this *TApp) GrabAttachmentIfNecess(a TAttachment) bool {
	var op TDBAttachmentOp
	op.Tx = this.DB.StartTx(true)
	defer op.Tx.Commit()
	op.Key = a.SURL.Text
	if false == op.CheckExists() {
		op.FileName = a.SFileName.Text
		var data = this.LoadAttachment(a.SURL.Text)
		op.Allowed = len(data) < 1024*1024
		if op.Allowed {
			op.Data = data
		}
		op.Write()
		return true
	} else {
		return false
	}
}

func (this *TApp) RunAttachmentTestMode() {
	this.DB = (&TDBMan{}).Create()
	this.DB.FilePath = "data/db-attachments.bolt"
	this.DB.Start()
	var op TDBAttachmentOp
	op.Tx = this.DB.StartTx(false)
	op.ForEach(func() {
		WriteLog(op.FileName + " allowed=" + hgo.BoolToStr(op.Allowed))
		WriteLog(op.Key)
		WriteLog(IntToStr(len(op.Data)) + " " + strconv.FormatFloat(float64(op.CompressionRate), 'f', 2, 64))
		ioutil.WriteFile("data/attachments/"+op.FileName, op.Data, os.ModePerm)
	})
	defer op.Tx.Commit()
	this.DB.Stop()
}
