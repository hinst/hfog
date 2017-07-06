package fog

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"image/jpeg"

	"github.com/nfnt/resize"
)

type TCompressImage struct {
	TargetWidth int
}

func (this TCompressImage) Go(data []byte) (result []byte) {
	var img, typeStr, decodeResult = image.Decode(bytes.NewReader(data))
	if decodeResult == nil && img != nil {
		if img.Bounds().Dx() <= this.TargetWidth && typeStr == "jpeg" {
			result = data
		}
		if img.Bounds().Dx() <= this.TargetWidth && typeStr != "jpeg" {
			result = JpegEncode(img)
		}
		if this.TargetWidth < img.Bounds().Dx() {
			var resizedImg = resize.Resize(uint(this.TargetWidth), 0, img, resize.Lanczos3)
			result = JpegEncode(resizedImg)
		}
	}
	return
}

func JpegEncode(img image.Image) (result []byte) {
	var compressedData bytes.Buffer
	jpeg.Encode(&compressedData, img, &jpeg.Options{Quality: 50})
	result = compressedData.Bytes()
	return
}

var ImageFileNameSuffixes = []string{".png", ".gif", ".jpg", ".jpeg"}
