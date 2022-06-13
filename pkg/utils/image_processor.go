package utils

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/discord/lilliput"
	"github.com/google/uuid"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type ImageProcessor struct {
	imageDirectory string
}

func NewImageProcessor(dir string) *ImageProcessor {
	return &ImageProcessor{
		imageDirectory: dir,
	}
}

var EncodeOptions = map[string]map[int]int{
	".jpeg": {lilliput.JpegQuality: 85},
	".png":  {lilliput.PngCompression: 7},
	".webp": {lilliput.WebpQuality: 85},
}

func (ip *ImageProcessor) StoreImage(image *graphql.Upload) error {
	// TODO: remove graphql type here

	assetDir := MustGet("ASSETS_DIR")

	originalDir := filepath.Join(assetDir, "original")
	compressedDir := filepath.Join(assetDir, "compressed")

	if err := os.MkdirAll(originalDir, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(compressedDir, 0755); err != nil {
		return err
	}

	// decoder wants []byte, so read the whole file into a buffer
	originalImgBuf, err := ioutil.ReadAll(image.File)
	if err != nil {
		return err
	}

	// copy buffer for a compressed version
	compressedImgBuf := make([]byte, len(originalImgBuf))
	copy(compressedImgBuf, originalImgBuf)

	// --------------------------------------------------

	originalDecoder, err := lilliput.NewDecoder(originalImgBuf)
	// this error reflects very basic checks,
	// mostly just for the magic bytes of the file to match known image formats
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}
	defer originalDecoder.Close()

	originalHeader, err := originalDecoder.Header()
	// this error is much more comprehensive and reflects format errors
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}

	if originalDecoder.Duration() != 0 { // todo: check if this is correct
		return fmt.Errorf("animated images are not supported")
		// fmt.Printf("duration: %.2f s\n", float64(decoder.Duration())/float64(time.Second))
	}

	// --------------------------------------------------

	compressedDecoder, err := lilliput.NewDecoder(compressedImgBuf)
	// this error reflects very basic checks,
	// mostly just for the magic bytes of the file to match known image formats
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}
	defer compressedDecoder.Close()

	compressedHeader, err := compressedDecoder.Header()
	// this error is much more comprehensive and reflects format errors
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err)
	}

	if compressedDecoder.Duration() != 0 { // todo: check if this is correct
		return fmt.Errorf("animated images are not supported")
		// fmt.Printf("duration: %.2f s\n", float64(decoder.Duration())/float64(time.Second))
	}

	// ------------------------------------------------------------

	// get ready to resize image,
	// using 8192x8192 maximum resize buffer size
	ops := lilliput.NewImageOps(8192)
	defer ops.Close()

	// create a buffer to store the output image, 50MB in this case
	originalOutputImg := make([]byte, 50*1024*1024)
	compressedOutputImg := make([]byte, 50*1024*1024)

	originalOpts := &lilliput.ImageOptions{
		FileType:             ".jpeg",
		Width:                originalHeader.Width(),
		Height:               originalHeader.Height(),
		ResizeMethod:         lilliput.ImageOpsNoResize,
		NormalizeOrientation: true,
	}

	compressedOpts := &lilliput.ImageOptions{
		FileType:             ".webp",
		Width:                compressedHeader.Width(),
		Height:               compressedHeader.Height(),
		ResizeMethod:         lilliput.ImageOpsNoResize,
		NormalizeOrientation: true,
		EncodeOptions:        EncodeOptions[".webp"],
	}

	// ------------------------------------------------------------

	// resize and transcode image
	originalOutputImg, err = ops.Transform(originalDecoder, originalOpts, originalOutputImg)
	if err != nil {
		return fmt.Errorf("error transforming image: %v", err)
	}

	// write image to disk
	originalImageName := strings.Replace(path.Join(originalDir, image.Filename), filepath.Ext(image.Filename), ".jpeg", -1) // todo, name after uuid

	if _, err := os.Stat(originalImageName); !os.IsNotExist(err) {
		return fmt.Errorf("file already exists: %s", originalImageName)
	}

	err = ioutil.WriteFile(originalImageName, originalOutputImg, 0400)
	if err != nil {
		return fmt.Errorf("error writing image: %v", err)
	}

	fmt.Printf("image written to %s\n", originalImageName)

	// --------------------------------------------------------------------------------

	// resize and transcode image
	compressedOutputImg, err = ops.Transform(compressedDecoder, compressedOpts, compressedOutputImg)
	if err != nil {
		return fmt.Errorf("error transforming image: %v", err)
	}

	// write image to disk
	compressedImageName := strings.Replace(path.Join(compressedDir, image.Filename), filepath.Ext(image.Filename), ".webp", -1) // todo, name after uuid

	if _, err := os.Stat(compressedImageName); !os.IsNotExist(err) {
		return fmt.Errorf("file already exists: %s", compressedImageName)
	}

	err = ioutil.WriteFile(compressedImageName, compressedOutputImg, 0400)
	if err != nil {
		return fmt.Errorf("error writing image: %v", err)
	}

	fmt.Printf("image written to %s\n", compressedImageName)

	return nil
}

func (ip *ImageProcessor) RemoveImage(image uuid.UUID) error {
	// TODO: remove graphql type here

	return nil
}
