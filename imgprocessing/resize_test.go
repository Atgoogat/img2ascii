package imgprocessing

import (
	"image"
	"testing"
)

func TestResizeWithScale(t *testing.T) {
	img := image.NewGray(image.Rect(0, 0, 1, 1))
	resizedImg := ResizeScale(img, 1)

	if !img.Bounds().Eq(resizedImg.Bounds()) {
		t.Error("img resized to diffrent size", img, resizedImg)
	}
}
