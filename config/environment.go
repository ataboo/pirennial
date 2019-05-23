package config

import (
	"os"
	"path/filepath"
	"runtime"
)

// GPIOActive if the arch is arm (pi)
var GPIOActive bool

func init() {
	GPIOActive = runtime.GOARCH == "arm"
}

// AssetPath get the local path to an asset
func AssetPath(path string) string {
	if os.Getenv("ASSET_PATH") != "" {
		return os.Getenv("ASSET_PATH") + "/" + path
	}

	return findAssetPath() + "assets/" + path
}

// findAssetPath keep stepping to parent directory until `assets` dir is present
func findAssetPath() string {
	path := "./"

	for {
		if abs, _ := filepath.Abs(path); abs == "/" || fileExists(path+"ataboo") {
			logger.Error("failed to find app root")
			return ""
		}

		if fileExists(path + "assets") {
			return path
		}

		path = "../" + path
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
