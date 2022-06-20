package imageprocessor

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

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

	maxImageSize  int64
	maxResolution int
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
		maxImageSize:      int64(params.MaxImageSizeMB) * 1024 * 1024,
		maxResolution:     int(params.MaxResolution),
	}, nil
}

func (ip *ImageProcessor) StoreImage(image []byte) (string, error) {
	imageHash := fmt.Sprintf("%x", sha256.Sum256(image))

	path := ip.getOriginalFilepath(imageHash)

	if _, err := os.Stat(path); err == nil {
		// an image with the same hash has already been processed
		return imageHash, nil
	}

	err := ip.createAndStoreEncoded(image, path, nil)
	return imageHash, err
}

func (ip *ImageProcessor) RemoveImage(imageStoreID string) error {
	err := os.Remove(ip.getOriginalFilepath(imageStoreID))
	if err != nil {
		return err
	}

	files, err := filepath.Glob(fmt.Sprintf("%s/%s_*%s", ip.resizedDirectory, imageStoreID, imageExtension))
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

func (ip *ImageProcessor) GetOriginal(imageStoreID string) (string, error) {
	return ip.getOriginalFilepath(imageStoreID), nil
}

func (ip *ImageProcessor) GetResized(imageStoreID string, width, height int) (string, error) {
	originalPath := ip.getOriginalFilepath(imageStoreID)

	original, err := ioutil.ReadFile(originalPath)
	if err == os.ErrNotExist {
		return "", ErrOriginalImageNotFound
	} else if err != nil {
		return "", err
	}

	resizedPath := ip.getResizedFilepath(imageStoreID, width, height)
	if _, err := os.Stat(resizedPath); err == nil {
		return resizedPath, nil
	}

	err = ip.createAndStoreEncoded(original, resizedPath, &resizeOptions{
		Width:  width,
		Height: height,
	})

	return resizedPath, err
}

func (ip *ImageProcessor) getOriginalFilepath(imageStoreID string) string {
	return filepath.Join(ip.originalDirectory, imageStoreID+imageExtension)
}

func (ip *ImageProcessor) getResizedFilepath(imageStoreID string, width, height int) string {
	filename := fmt.Sprintf("%s_%dx%d%s", imageStoreID, width, height, imageExtension)
	return filepath.Join(ip.originalDirectory, filename)
}

func (ip *ImageProcessor) GetMaxImageSize() int64 {
	return ip.maxImageSize
}

func IsImageStoreIDValid(imageStoreID string) bool {
	return isValidImageStoreID.MatchString(imageStoreID)
}
