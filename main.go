package main

import (
	"fmt"
	"os"

	"github.com/atgoogat/img2ascii/cli"
	"github.com/atgoogat/img2ascii/imgprocessing"
)

func main() {
	empty := ""
	args := cli.ParseArguments(cli.CliArguments{
		Filepaths:      []string{},
		OutputFilename: &empty,
		Scale:          1,
	})

	err := validateArguments(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	var output *os.File
	if args.OutputFilename == nil {
		output = os.Stdout
	} else {
		output, err = os.Create(*args.OutputFilename)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Could not create output file")
		}
		defer output.Sync()
		defer output.Close()
	}

	for _, filepath := range args.Filepaths {
		f, err := os.Open(filepath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not read file: %s\n", filepath)
			return
		}
		defer f.Close()
		img, err := imgprocessing.ImgDecode(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not decode image: %s\n", filepath)
		}
		img = imgprocessing.ResizeScale(img, args.Scale)
		grayImg := imgprocessing.Img2Gray(img)

		ascii, _ := imgprocessing.Convert(grayImg, imgprocessing.DefaultSymbols())
		output.WriteString(*ascii)
		output.WriteString(fmt.Sprintln())
	}
}

func validateArguments(args cli.CliArguments) error {
	if len(args.Filepaths) == 0 {
		return fmt.Errorf("no input files provided")
	}
	if args.Scale <= 0 {
		return fmt.Errorf("scale must not be less or equal to zero")
	}
	return nil
}
