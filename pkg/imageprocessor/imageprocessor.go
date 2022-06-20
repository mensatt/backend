package imageprocessor

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"log"

	"github.com/mensatt/backend/pkg/utils"
)

type ImageParams struct {
	ImageDirectory string
}

const (
	imageExtension = ".webp"
)

var (
	isValidImageStoreID      = regexp.MustCompile(`^[0-9a-f]{64}$`)
	ErrInvalidImageStoreID   = fmt.Errorf("invalid image store id")
	ErrOriginalImageNotFound = fmt.Errorf("original image not found")
)

type ImageProcessorConfig struct {
	ImageDirectory string
	MaxImageSizeMB int32
	MaxResolution  int32
}

type ImageProcessor struct {
	originalDirectory string
	resizedDirectory  string
}

func NewImageProcessor(params ImageProcessorConfig) (*ImageProcessor, error) {
	stat, err := os.Stat(params.ImageDirectory)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("image directory does not exist - path: %s", params.ImageDirectory)
	} else if err != nil {
		return nil, fmt.Errorf("error checking image directory - path: %s, error: %s", params.ImageDirectory, err)
	} else if !stat.IsDir() {
		return nil, fmt.Errorf("image directory is not a directory - path: %s", params.ImageDirectory)
	}

	originalDirectory := filepath.Join(params.ImageDirectory, "original")
	resizedDirectory := filepath.Join(params.ImageDirectory, "resized")

	err = utils.CreateDirIfNotExists(originalDirectory)
	if err != nil {
		return nil, err
	}
	err = utils.CreateDirIfNotExists(resizedDirectory)
	if err != nil {
		return nil, err
	}

	return &ImageProcessor{
		originalDirectory: originalDirectory,
		resizedDirectory:  resizedDirectory,
	}, nil
}

func (ih *ImageProcessor) StoreImage() {

}

func (ih *ImageProcessor) RemoveImage(imageStoreID string) error {
	err := os.Remove(ih.getOriginalFilepath(imageStoreID))
	if err != nil {
		return err
	}
	files, err := filepath.Glob(fmt.Sprintf("%s/%s_*%s", ih.resizedDirectory, imageStoreID, imageExtension))
	if err != nil {
		return err
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			log.Printf("error removing resized image file: %s - %s\n", f, err)
		}
	}
	return nil
}

func (ih *ImageProcessor) GetOriginal(imageStoreID string) (string, error) {
	// if !IsImageStoreIDValid(imageStoreID) {
	// 	return "", ErrInvalidImageStoreID
	// }

	return ih.getOriginalFilepath(imageStoreID), nil
}

func (ih *ImageProcessor) GetResized(imageStoreID string, width, height int) (string, error) {
	// if !IsImageStoreIDValid(imageStoreID) {
	// 	return "", ErrInvalidImageStoreID
	// }

	// TODO - check if resized image exists

	return ih.getResizedFilepath(imageStoreID, width, height), nil
}

func (ih *ImageProcessor) getOriginalFilepath(imageStoreID string) string {
	return filepath.Join(ih.originalDirectory, imageStoreID+imageExtension)
}

func (ih *ImageProcessor) getResizedFilepath(imageStoreID string, width, height int) string {
	filename := fmt.Sprintf("%s_%dx%d%s", imageStoreID, width, height, imageExtension)
	return filepath.Join(ih.originalDirectory, filename)
}

func IsImageStoreIDValid(imageStoreID string) bool {
	return isValidImageStoreID.MatchString(imageStoreID)
}
