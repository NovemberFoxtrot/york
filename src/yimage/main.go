package main

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"yeasy"
)

func identical(source, target image.Image) bool {
	sourcebounds := source.Bounds()
	targetbounds := target.Bounds()

	if (sourcebounds.Min.Y != targetbounds.Min.Y) || (sourcebounds.Min.X != targetbounds.Min.X) || (sourcebounds.Max.Y != targetbounds.Max.Y) || (sourcebounds.Max.X != targetbounds.Max.X) {
		return false
	}

	for y := sourcebounds.Min.Y; y < sourcebounds.Max.Y; y++ {
		for x := sourcebounds.Min.X; x < sourcebounds.Max.X; x++ {
			sr, sg, sb, sa := source.At(x, y).RGBA()
			tr, tg, tb, ta := target.At(x, y).RGBA()

			if (sr != tr) || (sg != tg) || (sb != tb) || (sa != ta) {
				return false
			}
		}
	}

	return true
}

func compare(source, target string) bool {
	sourcefile, err := os.Open(source)
	yeasy.CheckError(err)
	defer sourcefile.Close()

	targetfile, err := os.Open(target)
	yeasy.CheckError(err)
	defer targetfile.Close()

	sourceimage, _, err := image.Decode(sourcefile)
	targetimage, _, err := image.Decode(targetfile)

	return identical(sourceimage, targetimage)
}
