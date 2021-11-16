package imgprocessing

import (
	"image"

	"github.com/nfnt/resize"
)

func ResizeScale(img image.Image, scale float64) image.Image {
	bounds := img.Bounds()
	w, h := float64(bounds.Dx()), float64(bounds.Dy())
	return ResizeWidthHeight(img, uint(w*scale), uint(h*scale))
}

func ResizeWidthHeight(img image.Image, width, height uint) image.Image {
	return resize.Resize(width, height, img, resize.Bilinear)
}
