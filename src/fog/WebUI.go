package fog

import (
	"encoding/xml"
	"hgo"
	"io/ioutil"
	"net/http"
)

type TWebUI struct {
}

func (this *TWebUI) Create() *TWebUI {
	return this
}

func (this *TWebUI) Start() {
	this.LoadBugs()
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
	var bugs []TBugHeaderWebStruct
	var bugListData = ReadBugsFromFile(hgo.AppDir + "/data/bugs.xml")
	for _, item := range bugListData.Cases.Cases {
		var data = this.ReadBugData(item.IxBug)
		if data != nil && len(data.Cases.Cases) > 0 {
			bugs = append(bugs,
				TBugHeaderWebStruct{
					Number: StrToInt0(item.IxBug),
					Title:  data.Cases.Cases[0].STitle,
				})
			WriteLog(bugs[len(bugs)-1].Title)
		}
	}
}

func (this *TWebUI) GetBugs(response http.ResponseWriter, request *http.Request) {
}
