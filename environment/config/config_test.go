package config

import "testing"

func TestAppRoot(t *testing.T) {
	path, err := findAssetPath()
	if err != nil {
		t.Error("failed to load cfg: ", err)
	}

	if !fileExists(path + "assets") {
		t.Error("should find assets directory in path")
	}
}
