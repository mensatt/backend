package imageuploader

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/mensatt/backend/pkg/utils"
)

var (
	isValidImageHash = regexp.MustCompile(`^[0-9a-f]{64}$`)

	// https://www.garykessler.net/library/file_sigs.html
	magicNumbers = map[string][]byte{
		"jpg":  {0xff, 0xd8, 0xff},
		"jpeg": {0xff, 0xd8, 0xff},
		"png":  {0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a},
		"webp": {0x52, 0x49, 0x46, 0x46},
		"heif": {0x00, 0x00, 0x00, 0x18, 0x66, 0x74, 0x79, 0x70, 0x68, 0x65, 0x69, 0x63},
	}
)

type Config struct {
	ImageDirectory string
	MaxImageSizeMB int32
	MaxResolution  int32
}

type ImageUploader struct {
	imageDirectory string
	maxImageSize   int
	maxResolution  int
}

func NewImageUploader(params Config) (*ImageUploader, error) {
	stat, err := os.Stat(params.ImageDirectory)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("image directory does not exist - path: %s", params.ImageDirectory)
	} else if err != nil {
		return nil, fmt.Errorf("error checking image directory - path: %s, error: %s", params.ImageDirectory, err)
	} else if !stat.IsDir() {
		return nil, fmt.Errorf("image directory is not a directory - path: %s", params.ImageDirectory)
	}

	err = utils.CreateDirIfNotExists(params.ImageDirectory)
	if err != nil {
		return nil, err
	}

	return &ImageUploader{
		imageDirectory: params.ImageDirectory,
		maxImageSize:   int(params.MaxImageSizeMB) * 1024 * 1024,
		maxResolution:  int(params.MaxResolution),
	}, nil
}

func (iu *ImageUploader) ValidateAndStoreImage(image []byte) (string, error) {
	if len(image) > iu.maxImageSize {
		return "", fmt.Errorf("image is too large - max size: %d, actual size: %d", iu.maxImageSize, len(image))
	}

	if !IsImageValid(image) {
		return "", fmt.Errorf("image is invalid or format not accepted")
	}

	// todo: check if image dimensions are too large

	imageHash := fmt.Sprintf("%x", sha256.Sum256(image))

	path := iu.GetImagePath(imageHash)

	if _, err := os.Stat(path); err == nil {
		// an image with the same hash has already been processed
		// todo: check this behaviour
		return imageHash, nil
	}

	err := os.WriteFile(path, image, 0644)
	return imageHash, err
}

func (iu *ImageUploader) RemoveImage(imageHash string) error {
	err := os.Remove(iu.GetImagePath(imageHash))
	if err != nil {
		return err
	}

	files, err := filepath.Glob(fmt.Sprintf("%s/%s_*", iu.imageDirectory, imageHash))
	if err != nil {
		return err
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			log.Printf("error removing image file: %s - %s\n", f, err)
		}
	}
	return nil
}

func (iu *ImageUploader) GetImagePath(imageHash string) string {
	return filepath.Join(iu.imageDirectory, imageHash)
}

func (iu *ImageUploader) GetMaxImageSize() int {
	return iu.maxImageSize
}

func IsImageHashValid(imageHash string) bool {
	return isValidImageHash.MatchString(imageHash)
}

func IsImageValid(image []byte) bool {
	for _, magicNumber := range magicNumbers {
		if isMagicNumberValid(image, magicNumber) {
			return true
		}
	}
	return false
}

func isMagicNumberValid(image []byte, magicNumber []byte) bool {
	if len(image) < len(magicNumber) {
		return false
	}

	for i := range magicNumber {
		if image[i] != magicNumber[i] {
			return false
		}
	}
	return true
}
