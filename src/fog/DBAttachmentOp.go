package fog

import (
	"bytes"
	"hgo"
	"image"

	"github.com/boltdb/bolt"
)

type TDBAttachmentOp struct {
	Tx        *bolt.Tx
	Key       string
	Allowed   bool
	Data      []byte
	FileName  string
	ImageType string

	CompressionRate float32
	HeadMode        bool
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
	bucket.Put(
		GetDBManKey([]string{this.Key, "ImageType"}),
		[]byte(this.ImageType),
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
	this.Allowed = BoolFromData(
		bucket.Get(
			GetDBManKey([]string{this.Key, "Allowed"}),
		),
	)
	if false == this.HeadMode && this.Allowed {
		var compressedData = bucket.Get(
			GetDBManKey([]string{this.Key, "Data"}),
		)
		this.Data = hgo.DecompressBytes(compressedData)
		this.CompressionRate = float32(len(compressedData)) / float32(len(this.Data))
	}
	this.FileName = string(
		bucket.Get(
			GetDBManKey([]string{this.Key, "FileName"}),
		),
	)
	this.FileName = string(
		bucket.Get(
			GetDBManKey([]string{this.Key, "ImageType"}),
		),
	)
}

func (this *TDBAttachmentOp) Delete() {
	var bucket = this.Tx.Bucket(DBManAttachmentsBucketKey)
	bucket.Delete(GetDBManKey([]string{this.Key, "Data"}))
	bucket.Delete(GetDBManKey([]string{this.Key, "Allowed"}))
	bucket.Delete(GetDBManKey([]string{this.Key, "FileName"}))
}

func (this *TDBAttachmentOp) DetectImageType() {
	var _, typeStr, decodeResult = image.Decode(bytes.NewReader(this.Data))
	if decodeResult == nil {
		this.ImageType = typeStr
	}
}

func (this *TDBAttachmentOp) WriteImageType() {
	var bucket = this.Tx.Bucket(DBManAttachmentsBucketKey)
	bucket.Put(
		GetDBManKey([]string{this.Key, "ImageType"}),
		[]byte(this.ImageType),
	)
}
