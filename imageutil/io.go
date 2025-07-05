package imageutil

import (
	"image"
	"image/jpeg"
	"os"

	"golang.org/x/image/tiff"
)

func ReadRawImage(filename string) image.Image {
	imageReader, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer imageReader.Close()

	image, err := tiff.Decode(imageReader)
	if err != nil {
		panic(err)
	}

	return image
}

func ReadImage(filename string) image.Image {
	imageReader, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer imageReader.Close()

	image, _, err := image.Decode(imageReader)
	if err != nil {
		panic(err)
	}

	return image
}

func WriteImage(filename string, image image.Image) {
	imageWriter, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer imageWriter.Close()

	err = jpeg.Encode(imageWriter, image, nil)
	if err != nil {
		panic(err)
	}
}
