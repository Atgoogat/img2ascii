package imgprocessing

import (
	"bytes"
	"fmt"
	"image"
	"math"
)

func DefaultSymbols() []string {
	return []string{
		"@",
		"#",
		"+",
		"*",
		"-",
		",",
		".",
		" ",
	}
}

func Convert(img image.Gray, symbols []string) (*string, error) {
	if len(symbols) == 0 {
		return nil, fmt.Errorf("symbols must not be empty")
	}
	bounds := img.Bounds()
	dx, dy := bounds.Dx(), bounds.Dy()

	symbolRelPos := float64(len(symbols) - 1)
	outBuffer := bytes.NewBuffer([]byte{})
	outBuffer.Grow(dx * dy)
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			grayColor := img.GrayAt(x, y)
			grayScale := float64(grayColor.Y) / 255

			symbolIndex := int(math.Round(grayScale * symbolRelPos))
			outBuffer.WriteString(symbols[symbolIndex])
		}
		outBuffer.WriteString(fmt.Sprintln())
	}

	out := outBuffer.String()
	return &out, nil
}
