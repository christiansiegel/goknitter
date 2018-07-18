package goknitter

import (
	"image"
	"log"
	"os"

	"image/jpeg"
	_ "image/png" // register png for decoding
)

func Process(inputImage string) {
	infile, err := os.Open(inputImage)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer infile.Close()

	log.Println(infile)
	img, _, err := image.Decode(infile)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	outfile, err := os.Create("out.jpg")
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer outfile.Close()

	err = jpeg.Encode(outfile, img, &jpeg.Options{Quality: 25})
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
