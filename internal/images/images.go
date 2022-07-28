package images

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mensatt/backend/pkg/imageprocessor"
)

type ImageParams struct {
	ImageProcessor *imageprocessor.ImageProcessor
}

func Run(g *gin.RouterGroup, params *ImageParams) error {
	relativePath := "/:imageStoreID"
	handler := imageHandler(params.ImageProcessor)
	g.GET(relativePath, handler)
	g.HEAD(relativePath, handler)
	return nil
}

func imageHandler(ip *imageprocessor.ImageProcessor) gin.HandlerFunc {
	return func(c *gin.Context) {
		imageStoreID := c.Param("imageStoreID")
		if !imageprocessor.IsImageStoreIDValid(imageStoreID) {
			c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid imageStoreID: %s", imageStoreID))
		}

		widthStr := c.Query("width")
		heightStr := c.Query("height")

		var path string
		var err error
		if widthStr == "" && heightStr == "" {
			path, err = ip.GetOriginal(imageStoreID)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
		} else {
			width, err := strconv.Atoi(widthStr)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid width: %s", widthStr))
				return
			}
			height, err := strconv.Atoi(heightStr)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid height: %s", heightStr))
				return
			}

			path, err = ip.GetResized(imageStoreID, width, height)
			if err == imageprocessor.ErrOriginalImageNotFound {
				c.AbortWithError(http.StatusNotFound, err)
				return
			} else if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
		}
		c.File(path)
	}
}
