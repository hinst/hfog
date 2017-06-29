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
		var bucket = tx.Bucket(DBManTitlesBucketKey)
		bucket.Put([]byte(bug.IxBug), []byte(bug.STitle.Text))
		return nil
	})
}

func (this *TDBMan) Stop() {
	if this.db != nil {
		this.db.Close()
	}
}

var DBManTitlesBucketKey = []byte("Titles")
var DBManEventsBucketKey = []byte("Events")

func GetDBManKey(ids ...string) (result string) {
	for _, id := range ids {
		result += url.PathEscape(id) + "/"
	}
	return
}
