package fog

import (
	"encoding/xml"
	"hgo"
	"io/ioutil"
	"os"
)

type TDBGenApp struct {
	DB *TDBMan

	BugsEnabled        bool
	AttachmentsEnabled bool
}

func (this *TDBGenApp) Run() {
	this.DB = (&TDBMan{}).Create()
	this.DB.Start()
	if this.BugsEnabled {
		this.LoadBugs()
	}
	if this.AttachmentsEnabled {
		this.LoadAttachments()
	}
	this.DB.WriteToFile("db-shrink.bolt")
	this.DB.Stop()
	var oldSize = hgo.GetFileSize(this.DB.FilePath)
	os.Remove(this.DB.FilePath)
	os.Rename("db-shrink.bolt", this.DB.FilePath)
	var newSize = hgo.GetFileSize(this.DB.FilePath)
	WriteLog("DB shrinked: " + hgo.Int64ToStr(oldSize) + " -> " + hgo.Int64ToStr(newSize))
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
	for itemIndex, item := range bugListData.Cases.Cases {
		WriteLog("Writing item " + IntToStr(itemIndex))
		var data = this.ReadBugData(item.IxBug)
		if data != nil && len(data.Cases.Cases) > 0 {
			this.ClearEventsHtml(&data.Cases.Cases[0])
			this.DB.WriteBugData(&data.Cases.Cases[0])
		}
	}
}

func (this *TDBGenApp) ClearEventsHtml(data *TBugCaseData) {
	for i := range data.Events.Events {
		data.Events.Events[i].SHTML.Text = ""
	}
}

func (this *TDBGenApp) LoadAttachments() {
	var attachmentDB = (&TDBMan{}).Create()
	attachmentDB.FilePath = "data/db-attachments.bolt"
	attachmentDB.ReadOnly = true
	attachmentDB.Start()
	func() {
		var tx = attachmentDB.StartTx(false)
		defer tx.Commit()
		WriteLog("sourceAttachments=" + IntToStr(attachmentDB.GetCountOfAttachments(tx)))
	}()
	func() {
		var tx = this.DB.StartTx(true)
		defer tx.Commit()
		this.DB.ClearAttachments(tx)
	}()
	func() {
		var tx = this.DB.StartTx(true)
		defer tx.Commit()
		this.DB.CopyAttachments(tx, attachmentDB)
	}()
	func() {
		var tx = this.DB.StartTx(false)
		defer tx.Commit()
		WriteLog("destinationAttachments=" + IntToStr(this.DB.GetCountOfAttachments(tx)))
	}()
	func() {
		var tx = this.DB.StartTx(true)
		defer tx.Commit()
		var types = this.DB.DetectImageTypes(tx)
		for key, value := range types {
			WriteLog(key + "=" + IntToStr(value))
		}
	}()
	attachmentDB.Stop()
}
