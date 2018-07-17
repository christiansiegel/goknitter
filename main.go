package main

import (
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
)

const inputImageName = "/home/cs/Dev/knitter/go/in.jpg"
const outputImageName = "/home/cs/Dev/knitter/go/out.jpg"

func main() {
	infile, err := os.Open(inputImageName)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer infile.Close()

	img, _, err := image.Decode(infile)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	outfile, err := os.Create(outputImageName)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer outfile.Close()

	err = jpeg.Encode(outfile, img, &jpeg.Options{Quality: 25})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
