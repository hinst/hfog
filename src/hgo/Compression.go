package hgo

import (
	"bytes"
	"compress/flate"
	"io/ioutil"
)

func CompressBytes(a []byte) (result []byte) {
	var buffer bytes.Buffer
	var writer, writerResult = flate.NewWriter(&buffer, flate.DefaultCompression)
	AssertResult(writerResult)
	var _, writeResult = writer.Write(a)
	AssertResult(writeResult)
	var flushResult = writer.Flush()
	AssertResult(flushResult)
	var closeResult = writer.Close()
	AssertResult(closeResult)
	result = buffer.Bytes()
	return
}

func DecompressBytes(a []byte) []byte {
	var reader = bytes.NewReader(a)
	var deflater = flate.NewReader(reader)
	var data, readResult = ioutil.ReadAll(deflater)
	AssertResult(readResult)
	return data
}
