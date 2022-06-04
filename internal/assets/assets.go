package assets

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type AssetParams struct {
	AssetDirectory string
}

func Run(g *gin.RouterGroup, params *AssetParams) error {
	stat, err := os.Stat(params.AssetDirectory)
	if os.IsNotExist(err) {
		return fmt.Errorf("asset directory does not exist - path: %s", params.AssetDirectory)
	} else if err != nil {
		return fmt.Errorf("error checking asset directory - path: %s, error: %s", params.AssetDirectory, err)
	} else if !stat.IsDir() {
		return fmt.Errorf("asset directory is not a directory - path: %s", params.AssetDirectory)
	}

	g.Static("", params.AssetDirectory)
	return nil
}
