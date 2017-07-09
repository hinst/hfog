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

	LoadBugsModeEnabled             bool
	AttachmentsModeEnabled          bool
	AttachmentTestModeEnabled       bool
	EnumAttachmentsModeEnabled      bool
	ImageCompressionTestModeEnabled bool
	RunAllowImagesMode              bool

	AttachmentFilter []string
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
	if this.EnumAttachmentsModeEnabled {
		this.RunEnumAttachmentsMode()
	}
	if this.ImageCompressionTestModeEnabled {
		this.RunImageCompressionTest()
	}
	if this.RunAllowImagesMode {
		this.RunAllowImages()
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
				if this.CheckAttachmentFilterPass(attachment.SFileName.Text) {
					countOfAttachments++
					if this.GrabAttachmentIfNecess(attachment) {
						WriteLog(IntToStr(countOfAttachments) + " " + attachment.SFileName.Text + " " + attachment.SURL.Text)
						time.Sleep(5 * time.Second)
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
		op.Allowed = true //len(data) < 1024*1024
		if op.Allowed {
			if CheckStringHasSuffixes(op.FileName, ImageFileNameSuffixes) {
				var imageData = TCompressImage{TargetWidth: 1000}.Go(data)
				WriteLog("Image " + IntToStr(len(data)) + " -> " + IntToStr(len(imageData)))
				data = imageData
			}
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
	var totalCount = 0
	var allowedCount = 0
	op.ForEach(func() {
		totalCount++
		WriteLog(op.FileName + " allowed=" + hgo.BoolToStr(op.Allowed))
		if op.Allowed {
			allowedCount++
		}
		WriteLog(op.Key)
		WriteLog(IntToStr(len(op.Data)) + " " + strconv.FormatFloat(float64(op.CompressionRate), 'f', 2, 64))
		var writeFileResult = ioutil.WriteFile("data/attachments/"+op.FileName, op.Data, os.ModePerm)
		if false {
			AssertResult(writeFileResult)
		}
	})
	defer op.Tx.Commit()
	WriteLog("total=" + IntToStr(totalCount) + " allowed=" + IntToStr(allowedCount))
	this.DB.Stop()
}

func (this *TApp) RunEnumAttachmentsMode() {
	var bugListData = ReadBugsFromFile(hgo.AppDir + "/data/bugs.xml")
	var countOfAttachments = 0
	for _, item := range bugListData.Cases.Cases {
		if false == this.CheckActive() {
			break
		}
		var data = this.ReadBugData(item.IxBug)
		for _, event := range data.Cases.Cases[0].Events.Events {
			for _, attachment := range event.RGAttachments.Attachments {
				if this.CheckAttachmentFilterPass(attachment.SFileName.Text) {
					countOfAttachments++
				}
			}
		}
	}
	WriteLog("countOfAttachments=" + IntToStr(countOfAttachments))
}

func (this *TApp) CheckAttachmentFilterPass(x string) (result bool) {
	for _, suffix := range this.AttachmentFilter {
		if strings.HasSuffix(x, suffix) {
			result = true
			break
		}
	}
	return
}

func (this *TApp) RunImageCompressionTest() {
	var compressImage = TCompressImage{TargetWidth: 1000}

	var png1024, _ = ioutil.ReadFile("testData/1024.png")
	var png1024c = compressImage.Go(png1024)
	ioutil.WriteFile("testData/1024.png.jpg", png1024c, os.ModePerm)

	var png900, _ = ioutil.ReadFile("testData/900.png")
	var png900c = compressImage.Go(png900)
	ioutil.WriteFile("testData/900.png.jpg", png900c, os.ModePerm)

	var jpg900, _ = ioutil.ReadFile("testData/900.jpg")
	var jpg900c = compressImage.Go(jpg900)
	ioutil.WriteFile("testData/900.jpg.jpg", jpg900c, os.ModePerm)

	var jpg1024, _ = ioutil.ReadFile("testData/1024.jpg")
	var jpg1024c = compressImage.Go(jpg1024)
	ioutil.WriteFile("testData/1024.jpg.jpg", jpg1024c, os.ModePerm)
}

func (this *TApp) RunAllowImages() {
	this.DB = (&TDBMan{}).Create()
	this.DB.FilePath = "data/db-attachments.bolt"
	this.DB.Start()
	var op TDBAttachmentOp
	op.HeadMode = true
	op.Tx = this.DB.StartTx(true)
	var totalCount = 0
	var keysToRemove []string
	op.ForEach(func() {
		if false && CheckStringHasSuffixes(op.FileName, ImageFileNameSuffixes) && false == op.Allowed ||
			CheckStringHasSuffixes(op.FileName, []string{".png", ".gif"}) {
			keysToRemove = append(keysToRemove, op.Key)
		}
	})
	for _, key := range keysToRemove {
		op.Key = key
		op.Delete()
		totalCount++
	}
	WriteLog("reset=" + IntToStr(totalCount))
	op.Tx.Commit()
	this.DB.Stop()
}
