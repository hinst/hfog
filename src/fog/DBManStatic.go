package fog

import (
	"net/url"

	"github.com/boltdb/bolt"
)

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

var DBManAttachmentsBucketKey = []byte("Attachments")
