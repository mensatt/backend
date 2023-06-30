package imageuploader

import (
	"crypto/sha256"
	"fmt"
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/google/uuid"
	"github.com/mensatt/backend/pkg/utils"
	"os"
	"path/filepath"
)

var (
	// https://www.garykessler.net/library/file_sigs.html
	magicNumbers = map[string][]byte{
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

	vips.Startup(nil)

	return &ImageUploader{
		imageDirectory: params.ImageDirectory,
		maxImageSize:   int(params.MaxImageSizeMB) * 1024 * 1024,
		maxResolution:  int(params.MaxResolution),
	}, nil
}

func (iu *ImageUploader) ValidateAndStoreImage(image []byte, rotation *int) (uuid.UUID, string, error) {
	if len(image) > iu.maxImageSize {
		return uuid.Nil, "", fmt.Errorf("image is too large - max size: %d, actual size: %d", iu.maxImageSize, len(image))
	}

	if !isImageValid(image) {
		return uuid.Nil, "", fmt.Errorf("image is invalid or format not accepted")
	}

	// check if rotation angle is valid if provided
	if rotation != nil && !(*rotation == 0 || *rotation == 90 || *rotation == 180 || *rotation == 270) {
		return uuid.Nil, "", fmt.Errorf("invalid rotation value - value: %d, allowed values: 90, 180, 270", *rotation)
	}

	// convert integer rotation to vips.Angle
	var vipsAngle vips.Angle
	if rotation != nil {
		switch *rotation {
		case 0:
			vipsAngle = vips.Angle0
		case 90:
			vipsAngle = vips.Angle90
		case 180:
			vipsAngle = vips.Angle180
		case 270:
			vipsAngle = vips.Angle270
		}
	}

	vipsImage, err := vips.NewImageFromBuffer(image)
	defer vipsImage.Close()
	if err != nil {
		return uuid.Nil, "", err
	}

	err = vipsImage.Rotate(vipsAngle)
	if err != nil {
		return uuid.Nil, "", err
	}

	webpImage, _, err := vipsImage.ExportWebp(nil)
	if err != nil {
		return uuid.Nil, "", err
	}

	imageHash := fmt.Sprintf("%x", sha256.Sum256(webpImage))
	imageUUID := uuid.New()

	// generate new uuids until we find one that is not already in use
	for {
		path := iu.GetImagePath(imageUUID)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			break
		}
		imageUUID = uuid.New()
	}

	path := iu.GetImagePath(imageUUID)

	// owner: read, write; group: read; others: read
	err = os.WriteFile(path, webpImage, 0644)

	return imageUUID, imageHash, err
}

func (iu *ImageUploader) RemoveImage(uuid uuid.UUID) error {
	err := os.Remove(iu.GetImagePath(uuid))
	if err != nil {
		return err
	}
	return nil
}

func (iu *ImageUploader) GetImagePath(uuid uuid.UUID) string {
	return filepath.Join(iu.imageDirectory, uuid.String())
}

func (iu *ImageUploader) GetMaxImageSize() int {
	return iu.maxImageSize
}

func isImageValid(image []byte) bool {
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
