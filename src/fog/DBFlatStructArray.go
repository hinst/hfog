package fog

import (
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
