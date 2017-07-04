package fog

import (
	"hgo"
	"os"

	"strings"

	"github.com/boltdb/bolt"
)

type TDBMan struct {
	FilePath string
	db       *bolt.DB
}

func (this *TDBMan) Create() *TDBMan {
	this.FilePath = hgo.AppDir + "/data/db.bolt"
	return this
}

func (this *TDBMan) Start() {
	var db, dbOpenResult = bolt.Open(this.FilePath, os.ModePerm, nil)
	AssertResult(dbOpenResult)
	this.db = db
	this.db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(DBManTitlesBucketKey)
		tx.CreateBucketIfNotExists(DBManEventsBucketKey)
		tx.CreateBucketIfNotExists(DBManEventsBucketKey)
		return nil
	})
}

func (this *TDBMan) Stop() {
	if this.db != nil {
		this.db.Close()
	}
}

func (this *TDBMan) WriteBugData(bug *TBugCaseData) {
	this.db.Update(func(tx *bolt.Tx) error {
		DBManGetTitlesBucket(tx).Put(
			[]byte(bug.IxBug),
			[]byte(bug.STitle.Text))
		this.WriteBugDataEvents(bug, tx)
		return nil
	})
}

func (this *TDBMan) WriteBugDataEvents(bug *TBugCaseData, tx *bolt.Tx) {
	var eventsBucket = DBManGetEventsBucket(tx)
	eventsBucket.Put(
		GetDBManKey([]string{bug.IxBug, "n"}),
		[]byte(IntToStr(len(bug.Events.Events))))
	for eventIndex, event := range bug.Events.Events {
		event.ToDBStruct().WriteToBucket(eventsBucket, []string{bug.IxBug, IntToStr(eventIndex)})
	}
}

func (this *TDBMan) GetTitles() (result map[int]string) {
	result = make(map[int]string)
	this.db.View(func(tx *bolt.Tx) error {
		var bucket = DBManGetTitlesBucket(tx)
		var cursor = bucket.Cursor()
		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			var stringKey = string(key)
			var intKey = StrToInt0(stringKey)
			result[intKey] = string(value)
		}
		return nil
	})
	return
}

func (this *TDBMan) GetTitlesFiltered(filterString string) (result map[int]TRankedTitle) {
	result = make(map[int]TRankedTitle)
	var filterWords = strings.Split(filterString, " ")
	this.db.View(func(tx *bolt.Tx) error {
		var bucket = DBManGetTitlesBucket(tx)
		var cursor = bucket.Cursor()
		for key, value := cursor.First(); key != nil; key, value = cursor.Next() {
			var stringKey = string(key)
			var intKey = StrToInt0(stringKey)
			var rank = this.CheckBugFits(tx, stringKey, filterWords)
			if rank > 0 {
				result[intKey] = TRankedTitle{string(value), rank}
			}
		}
		return nil
	})
	return
}

func (this *TDBMan) CheckBugFits(tx *bolt.Tx, bugId string, filterWords []string) int {
	var title = string(DBManGetTitlesBucket(tx).Get([]byte(bugId)))
	//var bug = this.ReadBugData(tx, bugId)
	return CountStringContainsFromArray(title, filterWords) + CountStringContainsFromArray(bugId, filterWords)
}

func (this *TDBMan) WriteToFile(filePath string) {
	this.db.View(func(tx *bolt.Tx) error {
		tx.CopyFile(filePath, os.ModePerm)
		return nil
	})
}

func (this *TDBMan) LoadBugData(bugId string) (result *TBugCaseData) {
	this.db.Update(func(tx *bolt.Tx) error {
		result = this.ReadBugData(tx, bugId)
		return nil
	})
	return
}

func (this *TDBMan) ReadBugData(tx *bolt.Tx, bugId string) (result *TBugCaseData) {
	var titleData = DBManGetTitlesBucket(tx).Get([]byte(bugId))
	if titleData != nil {
		result = &TBugCaseData{}
		result.IxBug = bugId
		result.STitle.Text = string(titleData)
		this.ReadBugEvents(tx, result)
	}
	return
}

func (this *TDBMan) ReadBugEvents(tx *bolt.Tx, bug *TBugCaseData) {
	var eventsBucket = DBManGetEventsBucket(tx)
	var countOfEvents = StrToInt0(string(
		eventsBucket.Get(GetDBManKey([]string{bug.IxBug, "n"})),
	))
	bug.Events.Events = make([]TBugCaseEventData, countOfEvents)
	for eventIndex, event := range bug.Events.Events {
		var dbStruct = event.ToDBStruct().ReadFromBucket(eventsBucket, []string{bug.IxBug, IntToStr(eventIndex)})
		bug.Events.Events[eventIndex].LoadDBStruct(dbStruct)
	}
}

func (this *TDBMan) WriteAttachment(fogUrlKey string, data []byte) {

}
