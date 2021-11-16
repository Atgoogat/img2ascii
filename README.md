# Img2Ascii

Simple command-line-tool to convert an png or jpeg image to ascii-art

## Run

    ./img2ascii --help // for help
    ./img2ascii picture.png // converts and prints picture.png to stdout

### Options

* -o - Output file (default stdout)
* -scale - Scales the input images

## Build

    go get 
    go build 

## Test

    go test ./...
