package fog

import "github.com/boltdb/bolt"
import "hgo"

type TDBAttachmentOp struct {
	Tx       *bolt.Tx
	Key      string
	Allowed  bool
	Data     []byte
	FileName string
}

func (this *TDBAttachmentOp) Write() {
	var bucket = this.Tx.Bucket(DBManAttachmentsBucketKey)
	if this.Data == nil {
		this.Data = []byte{}
	}
	bucket.Put(
		GetDBManKey([]string{this.Key, "Data"}),
		hgo.CompressBytes(this.Data))
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

		return nil
	})
}
