package cli

import (
	"flag"
	"strings"
)

type CliArguments struct {
	Filepaths      []string
	OutputFilename *string
	Scale          *float64
	ScaleX         *float64
	ScaleY         *float64
}

func ParseArguments() CliArguments {
	outputFilename := flag.String("o", "", "output file (if blank printed to stdout)")

	scale, scaleX, scaleY := NewFloat64Flag(), NewFloat64Flag(), NewFloat64Flag()

	flag.Var(scale, "scale", "factor at which the image(s) are scaled")
	flag.Var(scaleX, "scaleX", "factor at which the image(s) are scaled on the x-axis")
	flag.Var(scaleY, "scaleY", "factor at which the image(s) are scaled on the y-axis")

	flag.Parse()

	outfilename := outputFilename
	if strings.Compare(*outfilename, "") == 0 {
		outfilename = nil
	}
	return CliArguments{
		Filepaths:      flag.Args(),
		OutputFilename: outfilename,
		Scale:          scale.value,
		ScaleX:         scaleX.value,
		ScaleY:         scaleY.value,
	}
}
