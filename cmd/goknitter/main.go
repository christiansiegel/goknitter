package main

import (
	"os"

	"github.com/golang/glog"

	"github.com/christiansiegel/goknitter/pkg/cmd/goknitter"
)

func main() {
	defer glog.Flush()

	inputImage := os.Args[1]

	goknitter.Process(inputImage)
}
