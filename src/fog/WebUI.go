package fog

import (
	"encoding/json"
	"encoding/xml"
	"hgo"
	"io/ioutil"
	"net/http"
)

type TWebUI struct {
	Bugs []TBugCaseData
	URL  string
}

func (this *TWebUI) Create() *TWebUI {
	this.URL = "/FogBugzBackup"
	return this
}

func (this *TWebUI) Start() {
	this.LoadBugs()
	this.AddRequestHandler("/bugs", this.GetBugs)
}

func (this *TWebUI) ReadBugList() {
	ReadBugsFromFile("data/bugs.xml")
}

func (this *TWebUI) ReadBugData(id string) (result *TBugData) {
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

func (this *TWebUI) LoadBugs() {
	var bugListData = ReadBugsFromFile(hgo.AppDir + "/data/bugs.xml")
	for _, item := range bugListData.Cases.Cases {
		var data = this.ReadBugData(item.IxBug)
		if data != nil && len(data.Cases.Cases) > 0 {
			this.Bugs = append(this.Bugs, data.Cases.Cases[0])
		}
	}
	WriteLog("Loaded bugs: " + IntToStr(len(this.Bugs)))
}

func (this *TWebUI) GetBugs(response http.ResponseWriter, request *http.Request) {
	var bugs []TBugHeaderWebStruct
	for _, item := range this.Bugs {
		bugs = append(bugs,
			TBugHeaderWebStruct{
				Number: StrToInt0(item.IxBug),
				Title:  item.STitle.Text,
			})
	}
	var data, marshalResult = json.Marshal(&bugs)
	WriteLogResult(marshalResult)
	if marshalResult == nil {
		response.Header().Set("Content-Type", "application/json")
		response.Write(data)
	}
}

func (this *TWebUI) AddRequestHandler(url string, f func(response http.ResponseWriter, request *http.Request)) {
	var wrappedFunc = func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		f(response, request)
	}
	http.HandleFunc(this.URL+url, wrappedFunc)
}
