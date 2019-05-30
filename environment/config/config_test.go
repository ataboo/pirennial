package config

import (
	"testing"

	"github.com/ataboo/pirennial/environment/filestorage"
)

func TestAppRoot(t *testing.T) {
	path, err := filestorage.FindAssetPath()
	if err != nil {
		t.Error("failed to load cfg: ", err)
	}

	if !filestorage.FileExists(path + "assets") {
		t.Error("should find assets directory in path")
	}
}
