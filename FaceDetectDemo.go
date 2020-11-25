package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"golang.org/x/image/colornames"
	"image"
	"log"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	displayWindow := gocv.NewWindow("Find a face")
	classifier := gocv.NewCascadeClassifier()
	success := classifier.Load("haarcascade_frontalface_default.xml")
	if !success {
		log.Fatal("Failed to load classifier - can't continue")
	}
	defer classifier.Close()
	FindFaces(webcam, displayWindow, classifier)
}

func FindFaces(camera *gocv.VideoCapture, window *gocv.Window, faceFiindingNet gocv.CascadeClassifier) {
	img := gocv.NewMat()
	defer img.Close()

	for {
		if ok := camera.Read(&img); !ok {
			fmt.Printf("cannot read from camera!")
			continue
		}
		if img.Empty() {
			continue
		}

		potentialFaces := faceFiindingNet.DetectMultiScale(img)

		for _, rectangle := range potentialFaces {
			faceRegion := img.Region(rectangle)
			gocv.GaussianBlur(faceRegion, &faceRegion, image.Pt(55, 95), 0, 0, gocv.BorderDefault)
			faceRegion.Close()
			gocv.Rectangle(&img, rectangle, colornames.Darkkhaki, 3)
		}
		window.IMShow(img)
		if window.WaitKey(10) >= 0 {
			break
		}

	}

}
