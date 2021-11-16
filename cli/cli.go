package cli

import (
	"flag"
	"strings"
)

type CliArguments struct {
	Filepaths      []string
	OutputFilename *string
	Scale          float64
}

func ParseArguments(defaultArg CliArguments) CliArguments {
	outputFilename := flag.String("o", *defaultArg.OutputFilename, "output file (if blank printed to stdout)")
	scale := flag.Float64("scale", defaultArg.Scale, "factor at which the image(s) are scaled")
	flag.Parse()

	outfilename := outputFilename
	if strings.Compare(*outfilename, "") == 0 {
		outfilename = nil
	}
	return CliArguments{
		Filepaths:      flag.Args(),
		OutputFilename: outfilename,
		Scale:          *scale,
	}
}
