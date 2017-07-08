package fog

import (
	"bytes"

	"github.com/boltdb/bolt"
)

type TDBFlatStructArray []TDBFlatStruct

func (this TDBFlatStructArray) SaveStrings(a map[string]string) (result TDBFlatStructArray) {
	result = make(TDBFlatStructArray, len(a))
	for key, value := range a {
		var item = TDBFlatStruct{Key: key, Data: value}
		result = append(result, item)
	}
	return
}

func (this TDBFlatStructArray) WriteToBucket(bucket *bolt.Bucket, key []string) {
	for _, item := range this {
		bucket.Put(GetDBManKey(append(key, item.Key)), []byte(item.Data))
	}
}

func (this TDBFlatStructArray) ReadStrings() (result map[string]string) {
	result = make(map[string]string)
	for _, item := range this {
		result[item.Key] = item.Data
	}
	return
}

func (this TDBFlatStructArray) ReadFromBucket(bucket *bolt.Bucket, key []string) (result TDBFlatStructArray) {
	var prefix = GetDBManKey(key)
	var cursor = bucket.Cursor()
	for subKey, value := cursor.Seek(prefix); subKey != nil && bytes.HasPrefix(subKey, prefix); subKey, value = cursor.Next() {
		var item TDBFlatStruct
		item.Key = UnpackDBManKey(subKey)[len(key)]
		item.Data = string(value)
		result = append(result, item)
	}
	return
}
