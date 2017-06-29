package fog

import (
	"hgo"
	"os"

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
