package images

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mensatt/backend/pkg/imageuploader"
	"net/http"
	"os"
)

type ImageParams struct {
	ImageUploader *imageuploader.ImageUploader
}

func Run(g *gin.RouterGroup, params *ImageParams) error {
	relativePath := "/:imageUUID"
	handler := imageHandler(params.ImageUploader)
	g.GET(relativePath, handler)
	g.HEAD(relativePath, handler)
	return nil
}

func imageHandler(ip *imageuploader.ImageUploader) gin.HandlerFunc {
	return func(c *gin.Context) {
		uuidString := c.Param("imageUUID")
		print(uuidString)
		imageUUID, err := uuid.Parse(uuidString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"}) // return 400 error if uuid is of an invalid format
			return
		}

		path := ip.GetImagePath(imageUUID)

		// check if file exists and send it if it does
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			c.File(path)
			return
		} else {
			// return 404 error if file does not exist
			c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
			return
		}
	}
}
