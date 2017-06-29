package fog

import (
	"encoding/xml"
	"hgo"
	"io/ioutil"
)

type TDBGenApp struct {
	Bugs []TBugCaseData
}

func (this *TDBGenApp) Run() {

}

func (this *TDBGenApp) ReadBugData(id string) (result *TBugData) {
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

func (this *TDBGenApp) LoadBugs() {
	var bugListData = ReadBugsFromFile(hgo.AppDir + "/data/bugs.xml")
	for _, item := range bugListData.Cases.Cases {
		var data = this.ReadBugData(item.IxBug)
		if data != nil && len(data.Cases.Cases) > 0 {
			this.Bugs = append(this.Bugs, data.Cases.Cases[0])
		}
	}
	WriteLog("Loaded bugs: " + IntToStr(len(this.Bugs)))
}
