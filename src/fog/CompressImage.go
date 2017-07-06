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

func (this TCompressImage) CompressImage(data []byte) (result []byte) {
	var img, _, decodeResult = image.Decode(bytes.NewReader(data))
	if decodeResult == nil && img != nil {
		var compressedImg = img
		if img.Bounds().Dx() <= this.TargetWidth {
		} else {
			compressedImg = resize.Resize(uint(this.TargetWidth), 0, img, resize.Lanczos3)
		}
		var compressedData bytes.Buffer
		jpeg.Encode(&compressedData, compressedImg, nil)
		result = compressedData.Bytes()
	}
	return
}
