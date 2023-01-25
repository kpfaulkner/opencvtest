package main

import (
	"fmt"

	"gocv.io/x/gocv"
)


func main() {
	fileName := "./test4.png"
	src := gocv.IMRead(fileName, gocv.IMReadColor)

	if src.Empty() {
		fmt.Printf("Invalid read of Source Mat in TestInpaint test\n")
		return
	}
	defer src.Close()

	img := gocv.NewMat()
	gocv.CvtColor(src, &img, gocv.ColorBGRToGray)

	points := gocv.FindContours(img, gocv.RetrievalCComp, gocv.ChainApproxNone)

	pointVectors := points.ToPoints()
	for _, outer := range pointVectors {
		for _, inner := range outer {
			fmt.Println(inner)
		}
	}
}
