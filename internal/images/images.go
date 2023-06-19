package images

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mensatt/backend/pkg/imageuploader"
	"net/http"
	"os"
	"strconv"
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
		imageUUID, err := uuid.Parse(uuidString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"}) // return 400 error if uuid is of an invalid format
			return
		}

		// default to 80% quality if no quality is specified
		quality := 80

		queryWidth := c.Query("width")
		queryHeight := c.Query("height")
		queryQuality := c.Query("quality")

		// check and set quality if specified and valid
		if queryQuality != "" {
			quality, err = strconv.Atoi(queryQuality)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid quality: must be an integer"})
				return
			}
			if quality < 0 || quality > 100 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid quality: must be between 0 and 100"})
				return
			}
		}

		path := ip.GetImagePath(imageUUID)

		// check if file exists and send error if it does not
		if _, err := os.Stat(path); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
			return
		}

		img, err := vips.NewImageFromFile(path)
		defer img.Close()

		// default to original image size if no width or height is specified
		width := img.Width()
		height := img.Height()

		if queryWidth != "" {
			width, err = strconv.Atoi(queryWidth)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid width: must be an integer"})
				return
			}
			if width < 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid width: must be greater than 0"})
				return
			}
		}

		if queryHeight != "" {
			height, err = strconv.Atoi(queryHeight)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid height: must be an integer"})
				return
			}
			if height < 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid height: must be greater than 0"})
				return
			}
		}

		// resize image
		err = img.ThumbnailWithSize(width, height, vips.InterestingAll, vips.SizeDown)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error resizing image"})
			return
		}

		data, _, err := img.ExportWebp(&vips.WebpExportParams{Quality: quality})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error exporting image"})
			return
		}

		c.Data(http.StatusOK, "image/webp", data)
		return
	}
}
