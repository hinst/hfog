package fog

import (
	"hgo"
	"os"

	"net/url"

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

func (this *TDBMan) Stop() {
	if this.db != nil {
		this.db.Close()
	}
}

var DBManTitlesBucketKey = []byte("Titles")

func DBManGetTitlesBucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket(DBManTitlesBucketKey)
}

var DBManEventsBucketKey = []byte("Events")

func DBManGetEventsBucket(tx *bolt.Tx) *bolt.Bucket {
	return tx.Bucket(DBManEventsBucketKey)
}

func GetDBManKey(ids []string) (result []byte) {
	var text string
	for _, id := range ids {
		text += "/" + url.PathEscape(id)
	}
	result = []byte(text)
	return
}
