package utils

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/discord/lilliput"
)

const (
	outputFormat             = ".webp"
	thumbnailImageFilename   = "thumbnail" + outputFormat
	originalResImageFilename = "original_res" + outputFormat
)

var encodeOption = map[int]int{lilliput.WebpQuality: 85}

type ImageProcessorConfig struct {
	ImageDirectory  string
	MaxOutputSizeMB int32
	MaxResolution   int32
}

type ImageProcessor struct {
	imageDirectory string
	maxOutputSize  int
	maxResolution  int
}

func NewImageProcessor(params ImageProcessorConfig) *ImageProcessor {
	return &ImageProcessor{
		imageDirectory: params.ImageDirectory,
		maxOutputSize:  int(params.MaxOutputSizeMB) * 1024 * 1024,
		maxResolution:  int(params.MaxResolution),
	}
}

func (ip *ImageProcessor) getImageBasePath(imageStoreID string) string {
	return filepath.Join(ip.imageDirectory, imageStoreID, thumbnailImageFilename)
}

func (ip *ImageProcessor) GetThumbnailImagePath(imageStoreID string) string {
	return filepath.Join(ip.getImageBasePath(imageStoreID), thumbnailImageFilename)
}

func (ip *ImageProcessor) GetOriginalResImagePath(imageStoreID string) string {
	return filepath.Join(ip.getImageBasePath(imageStoreID), originalResImageFilename)
}

func (ip *ImageProcessor) StoreImage(image []byte) (string, error) {
	imageHash := fmt.Sprintf("%x", sha256.Sum256(image))

	imageDir := ip.getImageBasePath(imageHash)
	thumbnailImageFilename := ip.GetThumbnailImagePath(imageHash)
	originalResImageFilename := ip.GetOriginalResImagePath(imageHash)

	if FileOrFirExists(imageDir) {
		if FileOrFirExists(thumbnailImageFilename) && FileOrFirExists(originalResImageFilename) {
			// image has already been processed
			return imageHash, nil
		}

		// image has not been processed completely
		ip.RemoveImage(imageHash)
	}

	decoder, err := lilliput.NewDecoder(image)
	// this error reflects very basic checks, mostly just for the magic bytes of the file to match known image formats
	if err != nil {
		return "", fmt.Errorf("error decoding image: %v", err)
	}
	defer decoder.Close()

	header, err := decoder.Header()
	// this error is much more comprehensive and reflects format errors
	if err != nil {
		return "", fmt.Errorf("error decoding image: %v", err)
	}

	// print some basic info about the image
	fmt.Printf("file type: %s\n", decoder.Description())
	fmt.Printf("%dpx x %dpx\n", header.Width(), header.Height())

	if decoder.Duration() != 0 {
		return "", fmt.Errorf("image is duration is not 0 (duration: %.2f s)", float64(decoder.Duration())/float64(time.Second))
	}

	frame := lilliput.NewFramebuffer(ip.maxResolution, ip.maxResolution)
	defer frame.Close()

	err = decoder.DecodeTo(frame)
	if err != nil {
		return "", fmt.Errorf("error decoding image: %v", err)
	}

	// This will undo JPEG EXIF-based orientation.
	frame.OrientationTransform(header.Orientation())

	// create a buffer to store the output image
	outputImageBuffer := make([]byte, ip.maxOutputSize)

	// create and store thumbnail image
	encOpts := encodeOptions{
		Width:        400,
		Height:       400,
		ResizeMethod: lilliput.ImageOpsResize,
	}
	err = ip.createEncoded(frame, outputImageBuffer, &encOpts)
	if err != nil {
		return "", fmt.Errorf("error creating thumbnail image: %v", err)
	}
	err = ip.storeEncodedBuffer(thumbnailImageFilename, outputImageBuffer)
	if err != nil {
		return "", fmt.Errorf("error storing thumbnail image: %v", err)
	}

	// create and store originalRes image
	err = ip.createEncoded(frame, outputImageBuffer, nil)
	if err != nil {
		return "", fmt.Errorf("error creating originalRes image: %v", err)
	}
	err = ip.storeEncodedBuffer(originalResImageFilename, outputImageBuffer)
	if err != nil {
		return "", fmt.Errorf("error storing originalRes image: %v", err)
	}

	return imageHash, nil
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

func (ip *ImageProcessor) RemoveImage(imageStoreID string) error {
	imageDir := ip.getImageBasePath(imageStoreID)
	// maybe we should remove all images and the image directory separately, removeAll seems dangerous
	return os.RemoveAll(imageDir)
}
