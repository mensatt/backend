package imageprocessor

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/discord/lilliput"
)

var encodeOption = map[int]int{lilliput.WebpQuality: 85}

type resizeOptions struct {
	// Width controls the width of the output image
	Width int

	// Height controls the height of the output image
	Height int
}

func (ip *ImageProcessor) createAndStoreEncoded(image []byte, outputFilename string, resize *resizeOptions) error {

	decoder, err := lilliput.NewDecoder(image)
	// this error reflects very basic checks,
	// mostly just for the magic bytes of the file to match known image formats
	if err != nil {
		return fmt.Errorf("error decoding image, %s", err)
	}
	defer decoder.Close()

	header, err := decoder.Header()
	// this error is much more comprehensive and reflects format errors
	if err != nil {
		return fmt.Errorf("error decoding image, %s", err)
	}

	if decoder.Duration() != 0 {
		return fmt.Errorf("image is duration is not 0 (duration: %.2f s)", float64(decoder.Duration())/float64(time.Second))
	}

	ops := lilliput.NewImageOps(ip.maxResolution)
	defer ops.Close()

	outputImg := make([]byte, ip.maxImageSize)

	resizeMethod := lilliput.ImageOpsNoResize
	outputWidth := header.Width()
	outputHeight := header.Height()
	if resize != nil && (resize.Width != outputWidth || resize.Height != outputHeight) {
		resizeMethod = lilliput.ImageOpsFit
		outputWidth = resize.Width
		outputHeight = resize.Height
	}

	opts := &lilliput.ImageOptions{
		FileType:             imageExtension,
		Width:                outputWidth,
		Height:               outputHeight,
		ResizeMethod:         resizeMethod,
		NormalizeOrientation: true,
		EncodeOptions:        encodeOption,
	}

	// resize and transcode image
	outputImg, err = ops.Transform(decoder, opts, outputImg)
	if err != nil {
		return fmt.Errorf("error transforming image, %s", err)
	}

	return ioutil.WriteFile(outputFilename, outputImg, 0400)
}
