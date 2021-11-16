package imgprocessing

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func ImgDecode(f *os.File) (image.Image, error) {
	img, _, err := image.Decode(f)
	return img, err
}

func Img2Gray(img image.Image) image.Gray {
	grayImg := image.NewGray(img.Bounds())
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			grayImg.Set(x, y, img.At(x, y))
		}
	}
	return *grayImg
}
