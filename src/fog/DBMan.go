package fog

import (
	"os"

	"github.com/boltdb/bolt"
)

type TDBMan struct {
	FilePath string
	db       *bolt.DB
}

func (this *TDBMan) Start() {
	var db, dbOpenResult = bolt.Open(this.FilePath, os.ModePerm, nil)
	AssertResult(dbOpenResult)
	this.db = db
}

func (this *TDBMan) Stop() {
	if this.db != nil {
		this.db.Close()
	}
}
