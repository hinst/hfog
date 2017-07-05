package fog

import (
	"net/url"

	"strings"

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

func UnpackDBManKey(key []byte) (result []string) {
	if len(key) > 0 {
		key = key[1:]
		var parts = strings.Split(string(key), "/")
		for _, part := range parts {
			var unescapedPart, unespaceResult = url.PathUnescape(part)
			if unespaceResult != nil {
				unescapedPart = ""
			}
			result = append(result,
				unescapedPart,
			)
		}
	}
	return
}

var DBManAttachmentsBucketKey = []byte("Attachments")
