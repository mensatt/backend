package images

import (
	"github.com/gin-gonic/gin"
	"github.com/mensatt/backend/pkg/imageuploader"
	"net/http"
	"os"
)

type ImageParams struct {
	ImageUploader *imageuploader.ImageUploader
}

func Run(g *gin.RouterGroup, params *ImageParams) error {
	relativePath := "/:imageHash"
	handler := imageHandler(params.ImageUploader)
	g.GET(relativePath, handler)
	g.HEAD(relativePath, handler)
	return nil
}

func imageHandler(ip *imageuploader.ImageUploader) gin.HandlerFunc {
	return func(c *gin.Context) {
		imageHash := c.Param("imageHash")
		print(imageHash)
		if !imageuploader.IsImageHashValid(imageHash) {
			// return 400 error if image hash format is invalid
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid image hash"})
			return
		}

		path := ip.GetImagePath(imageHash)
		print(path)

		// check if file exists and send it if it does
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			c.File(path)
			return
		} else if err != nil {
			// return 404 error if file does not exist
			c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
			return
		}
	}
}
