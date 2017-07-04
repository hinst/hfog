package fog

import "github.com/boltdb/bolt"
import "hgo"

type TDBAttachmentOp struct {
	Tx      *bolt.Tx
	Key     string
	Allowed bool
	Data    []byte
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
}
