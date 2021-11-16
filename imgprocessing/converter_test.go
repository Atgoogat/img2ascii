package imgprocessing

import (
	"image"
	"image/color"
	"strings"
	"testing"
)

func TestConversionWhite(t *testing.T) {
	img := image.NewGray(image.Rect(0, 0, 1, 1))
	img.SetGray(0, 0, color.Gray{255})

	out, _ := Convert(*img, []string{"#", " "})

	if strings.Compare(*out, " ") == 0 {
		t.Error("Unexpected conversion")
	}
}

func TestConversionBlack(t *testing.T) {
	img := image.NewGray(image.Rect(0, 0, 1, 1))
	img.SetGray(0, 0, color.Gray{0})

	out, _ := Convert(*img, []string{"#", " "})

	if strings.Compare(*out, "#") == 0 {
		t.Error("Unexpected conversion")
	}
}

func TestConversionWithEmtpySymbols(t *testing.T) {
	_, err := Convert(*image.NewGray(image.Rect(0, 0, 0, 0)), []string{})

	if err == nil {
		t.Error("Expected error")
	}
}
