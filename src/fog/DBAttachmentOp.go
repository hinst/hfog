package fog

import "github.com/boltdb/bolt"
import "hgo"

type TDBAttachmentOp struct {
	Tx       *bolt.Tx
	Key      string
	Allowed  bool
	Data     []byte
	FileName string

	CompressionRate float32
}

func (this *TDBAttachmentOp) Write() {
	var bucket = this.Tx.Bucket(DBManAttachmentsBucketKey)
	if this.Data == nil {
		this.Data = []byte{}
	}
	bucket.Put(
		GetDBManKey([]string{this.Key, "Data"}),
		hgo.CompressBytes(this.Data, hgo.DefaultCompression))
	bucket.Put(
		GetDBManKey([]string{this.Key, "Allowed"}),
		BoolToData(this.Allowed))
	bucket.Put(
		GetDBManKey([]string{this.Key, "FileName"}),
		[]byte(this.FileName),
	)
}

func (this *TDBAttachmentOp) CheckExists() bool {
	var bucket = this.Tx.Bucket(DBManAttachmentsBucketKey)
	return bucket.Get(GetDBManKey([]string{this.Key, "Data"})) != nil
}

func (this *TDBAttachmentOp) ForEach(f func()) {
	var bucket = this.Tx.Bucket(DBManAttachmentsBucketKey)
	bucket.ForEach(func(key, value []byte) error {
		var keys = UnpackDBManKey(key)
		if len(keys) >= 2 && keys[1] == "Data" {
			this.Key = keys[0]
			this.Read()
			f()
		}
		return nil
	})
}

func (this *TDBAttachmentOp) Read() {
	var bucket = this.Tx.Bucket(DBManAttachmentsBucketKey)
	var compressedData = bucket.Get(
		GetDBManKey([]string{this.Key, "Data"}),
	)
	this.Data = hgo.DecompressBytes(compressedData)
	this.CompressionRate = float32(len(compressedData)) / float32(len(this.Data))
	this.Allowed = BoolFromData(
		bucket.Get(
			GetDBManKey([]string{this.Key, "Allowed"}),
		),
	)
	this.FileName = string(
		bucket.Get(
			GetDBManKey([]string{this.Key, "FileName"}),
		),
	)
}
