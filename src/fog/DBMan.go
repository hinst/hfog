package fog

import (
	"hgo"
	"os"

	"strings"

	"github.com/boltdb/bolt"
)

type TDBMan struct {
	FilePath string
	ReadOnly bool
	db       *bolt.DB
}

func (this *TDBMan) Create() *TDBMan {
	this.FilePath = hgo.AppDir + "/data/db.bolt"
	return this
}

func (this *TDBMan) Start() {
	var o = bolt.DefaultOptions
	o.ReadOnly = this.ReadOnly
	var db, dbOpenResult = bolt.Open(this.FilePath, os.ModePerm, nil)
	AssertResult(dbOpenResult)
	this.db = db
	this.db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(DBManTitlesBucketKey)
		tx.CreateBucketIfNotExists(DBManEventsBucketKey)
		tx.CreateBucketIfNotExists(DBManAttachmentsBucketKey)
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
	var file, fileResult = os.Create(filePath)
	AssertResult(fileResult)
	this.db.View(func(tx *bolt.Tx) error {
		tx.WriteTo(file)
		file.Close()
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
	for eventIndex := range bug.Events.Events {
		var dbStruct = TDBFlatStructArray{}.ReadFromBucket(eventsBucket, []string{bug.IxBug, IntToStr(eventIndex)})
		bug.Events.Events[eventIndex].LoadDBStruct(dbStruct)
	}
}

func (this *TDBMan) StartTx(canWrite bool) *bolt.Tx {
	var transaction, transactionResult = this.db.Begin(canWrite)
	AssertResult(transactionResult)
	return transaction
}

func (this *TDBMan) CopyAttachments(tx *bolt.Tx, db *TDBMan) {
	var dbTx = db.StartTx(false)
	dbTx.Bucket(DBManAttachmentsBucketKey).ForEach(
		func(key, value []byte) error {
			tx.Bucket(DBManAttachmentsBucketKey).Put(key, value)
			return nil
		})
}

func (this *TDBMan) ClearAttachments(tx *bolt.Tx) {
	var keys [][]byte
	tx.Bucket(DBManAttachmentsBucketKey).ForEach(
		func(key, value []byte) error {
			keys = append(keys, key)
			return nil
		})
	for _, key := range keys {
		tx.Bucket(DBManAttachmentsBucketKey).Delete(key)
	}
}

func (this *TDBMan) GetCountOfAttachments(tx *bolt.Tx) (result int) {
	var op TDBAttachmentOp
	op.Tx = tx
	op.ForEach(func() {
		result++
	})
	return
}

func (this *TDBMan) DetectImageTypes(tx *bolt.Tx) (types map[string]int) {
	types = make(map[string]int)
	var op TDBAttachmentOp
	op.Tx = tx
	op.ForEach(func() {
		op.DetectImageType()
		op.WriteImageType()
		var count, exists = types[op.ImageType]
		if false == exists {
			count = 0
		}
		count++
		types[op.ImageType] = count
	})
	return
}

func (this *TDBMan) LoadAttachment(tx *bolt.Tx, key string) (data []byte) {
	var op TDBAttachmentOp
	op.Tx = tx
	op.Key = key
	op.Read()
	return op.Data
}
