package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	displayWindow := gocv.NewWindow("Find a face")

}
