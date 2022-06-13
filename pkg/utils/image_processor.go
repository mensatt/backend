package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/discord/lilliput"
)

const (
	outputFormat           = ".webp"
	thumbnailImageFilename = "thumbnail" + outputFormat
	originalImageFilename  = "original" + outputFormat
)

var encodeOption = map[int]int{lilliput.WebpQuality: 85}

type ImageProcessor struct {
	imageDirectory  string
	maxOutputSizeMB int
	maxResolution   int
}

func NewImageProcessor(dir string) *ImageProcessor {
	return &ImageProcessor{
		imageDirectory: dir,
	}
}

func (ip *ImageProcessor) getImageBasePath(imageID string) string {
	return filepath.Join(ip.imageDirectory, imageID, thumbnailImageFilename)
}

func (ip *ImageProcessor) GetThumbnailImagePath(imageID string) string {
	return filepath.Join(ip.getImageBasePath(imageID), thumbnailImageFilename)
}

func (ip *ImageProcessor) GetLargeImagePath(imageID string) string {
	return filepath.Join(ip.getImageBasePath(imageID), originalImageFilename)
}

func (ip *ImageProcessor) StoreImage(imageID string, image []byte) error {
	imageDir := ip.getImageBasePath(imageID)
	if _, err := os.Stat(imageDir); !os.IsNotExist(err) {
		fmt.Printf("output filename %s exists, quitting\n", imageDir)
		os.Exit(1)
	}

	decoder, err := lilliput.NewDecoder(image)
	// this error reflects very basic checks, mostly just for the magic bytes of the file to match known image formats
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}
	defer decoder.Close()

	header, err := decoder.Header()
	// this error is much more comprehensive and reflects format errors
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}

	// print some basic info about the image
	fmt.Printf("file type: %s\n", decoder.Description())
	fmt.Printf("%dpx x %dpx\n", header.Width(), header.Height())

	if decoder.Duration() != 0 {
		return fmt.Errorf("image is duration is not 0 (duration: %.2f s)", float64(decoder.Duration())/float64(time.Second))
	}

	frame := lilliput.NewFramebuffer(ip.maxResolution, ip.maxResolution)
	defer frame.Close()

	err = decoder.DecodeTo(frame)
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}

	frame.OrientationTransform(header.Orientation())

	// create a buffer to store the output image, 50MB in this case
	outputImageBuffer := make([]byte, ip.maxOutputSizeMB*1024*1024)

	// create and store thumbnail image
	encOpts := encodeOptions{
		Width:        400,
		Height:       400,
		ResizeMethod: lilliput.ImageOpsResize,
	}
	err = ip.createEncoded(frame, outputImageBuffer, &encOpts)
	if err != nil {
		return fmt.Errorf("error creating thumbnail: %v", err)
	}
	err = ip.storeEncodedBuffer(ip.GetThumbnailImagePath(imageID), outputImageBuffer)
	if err != nil {
		return fmt.Errorf("error storing thumbnail: %v", err)
	}

	// create and store original image
	err = ip.createEncoded(frame, outputImageBuffer, nil)
	if err != nil {
		return fmt.Errorf("error creating original: %v", err)
	}
	err = ip.storeEncodedBuffer(ip.GetThumbnailImagePath(imageID), outputImageBuffer)
	if err != nil {
		return fmt.Errorf("error storing original: %v", err)
	}

	// thumbnailImageFilename := ip.GetThumbnailImagePath(imageID)
	// err = ioutil.WriteFile(thumbnailImageFilename, outputImg, 0400)
	// if err != nil {
	// 	return fmt.Errorf("error writing out resized image, %s", err)
	// }

	// // get ready to resize image,
	// // using 8192x8192 maximum resize buffer size
	// ops := lilliput.NewImageOps(8192)
	// defer ops.Close()

	// // create a buffer to store the output image, 50MB in this case
	// outputImg := make([]byte, 50*1024*1024)

	// opts := &lilliput.ImageOptions{
	// 	FileType:             outputFormat,
	// 	Width:                header.Width(),
	// 	Height:               header.Height(),
	// 	NormalizeOrientation: true,
	// 	EncodeOptions:        encodeOption,
	// }

	// // resize and transcode image
	// outputImg, err = ops.Transform(decoder, opts, outputImg)
	// if err != nil {
	// 	fmt.Printf("error transforming image, %s\n", err)
	// 	return errors.New("error transforming image")
	// }

	// err = ioutil.WriteFile("outputFilename", outputImg, 0400)
	// if err != nil {
	// 	fmt.Printf("error writing out resized image, %s\n", err)
	// 	return errors.New("error writing out resized image")
	// }

	// // fmt.Printf("image written to %s\n", outputFilename)

	return nil
}

type encodeOptions struct {
	// Width controls the width of the output image
	Width int

	// Height controls the height of the output image
	Height int

	// ResizeMethod controls how the image will be transformed to
	// its output size. Notably, ImageOpsFit will do a cropping
	// resize, while ImageOpsResize will stretch the image.
	ResizeMethod lilliput.ImageOpsSizeMethod
}

func (ip *ImageProcessor) createEncoded(frame *lilliput.Framebuffer, outputBuf []byte, opts *encodeOptions) error {
	enc, err := lilliput.NewEncoder(outputFormat, nil, outputBuf)
	if err != nil {
		return err
	}
	defer enc.Close()

	if opts == nil || opts.ResizeMethod != lilliput.ImageOpsNoResize {
		resizedFrame := lilliput.NewFramebuffer(ip.maxResolution, ip.maxResolution)
		defer resizedFrame.Close()

		if opts.ResizeMethod == lilliput.ImageOpsFit {
			err = frame.Fit(opts.Width, opts.Height, resizedFrame)
		} else if opts.ResizeMethod == lilliput.ImageOpsResize {
			err = frame.ResizeTo(opts.Width, opts.Height, resizedFrame)
		} else {
			err = fmt.Errorf("invalid resize method: %v", opts.ResizeMethod)
		}
		if err != nil {
			return err
		}

		frame = resizedFrame
	}

	_, err = enc.Encode(frame, encodeOption)
	return err
}

func (ip *ImageProcessor) storeEncodedBuffer(filename string, imageBuffer []byte) error {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return fmt.Errorf("file already exists: %s", filename)
	}

	return ioutil.WriteFile(filename, imageBuffer, 0400)
}

func (ip *ImageProcessor) RemoveImage(imageID string) error {
	imageDir := ip.getImageBasePath(imageID)
	// maybe we should remove all images and the image directory separately, removeAll seems dangerous
	return os.RemoveAll(imageDir)
}
