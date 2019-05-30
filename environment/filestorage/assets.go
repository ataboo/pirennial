package filestorage

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/op/go-logging"
)

var logger *logging.Logger

// AssetPath get the local path to an asset
func AssetPath(path string) (string, error) {
	logger = logging.MustGetLogger("pirennial")

	if os.Getenv("ASSET_PATH") != "" {
		return os.Getenv("ASSET_PATH") + "/" + path, nil
	}

	assets, err := FindAssetPath()
	if err != nil {
		return "", err
	}

	return assets + "assets/" + path, nil
}

// FindAssetPath keep stepping to parent directory until `assets` dir is present
func FindAssetPath() (string, error) {
	path := "./"
	tries := 10

	for {
		if abs, _ := filepath.Abs(path); abs == "/" || FileExists(path+"pirennial") || tries == 0 {
			err := fmt.Errorf("failed to find assets directory")
			logger.Error(err)
			return "", err
		}
		tries--

		if FileExists(path + "assets") {
			return path, nil
		}

		path = "../" + path
	}
}

func FileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
