package filestorage

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/op/go-logging"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("pirennial")
}

// AssetPath get the local path to an asset
func AssetPath(path string) (string, error) {
	if os.Getenv("ASSET_PATH") != "" {
		return os.Getenv("ASSET_PATH") + "/" + path, nil
	}

	assets, err := findAppRoot("")
	if err != nil {
		return "", err
	}

	return assets + "assets/" + path, nil
}

func findAppRoot(startPath string) (string, error) {
	// Start at the system's current directory and travel up until the app root is recognized or tries is exceeded.
	path := "./"
	tries := 10
	if startPath != "" {
		path = startPath
	}

	for {
		if abs, _ := filepath.Abs(path); abs == "/" || tries == 0 {
			err := fmt.Errorf("failed to find assets directory")
			logger.Error(err)
			return "", err
		}
		tries--

		if FileExists(path + "go.mod") {
			return path, nil
		}

		path = "../" + path
	}
}

// FileExists determine if a path exists
func FileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
