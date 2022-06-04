package utils

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

type ImageProcessor struct {
	imageDirectory string
}

func NewImageProcessor(dir string) *ImageProcessor {
	return &ImageProcessor{
		imageDirectory: dir,
	}
}

func (ip *ImageProcessor) StoreImage(image *graphql.Upload) error {
	// TODO: remove graphql type here

	// imageFileName :=
	// dst, err := os.Create(filepath.Join(ip.imageDirectory, filepath.Base(file.Filename))) // dir is directory where you want to save file.
	// if err != nil {
	// 	checkErr(err)
	// 	return
	// }

	return nil
}

func (ip *ImageProcessor) RemoveImage(image uuid.UUID) error {
	// TODO: remove graphql type here

	// imageFileName :=
	// dst, err := os.Create(filepath.Join(ip.imageDirectory, filepath.Base(file.Filename))) // dir is directory where you want to save file.
	// if err != nil {
	// 	checkErr(err)
	// 	return
	// }

	return nil
}
