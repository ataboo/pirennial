package filestorage

import (
	"os"
	"strings"
	"testing"
)

func TestAppRoot(t *testing.T) {
	path, err := findAppRoot("")
	if err != nil {
		t.Error("failed to find app root", err)
	}

	if !FileExists(path + "assets") {
		t.Error("should find assets directory in root path")
	}

	if !FileExists(path + ".git") {
		t.Error(".git should be in app root")
	}

	if !FileExists(path + "go.mod") {
		t.Error("go.mod should be in app root")
	}
}

func TestAppRootGivesUpAtSystemRoot(t *testing.T) {
	_, err := findAppRoot("/")

	if err == nil {
		t.Error("app root should have failed")
	}
}

func TestAssetPath(t *testing.T) {
	path, err := AssetPath("")
	if err != nil {
		t.Error("should have found asset path", err)
	}

	if !strings.HasSuffix(path, "/assets/") {
		t.Error("asset path unnexpected", path)
	}

	if !FileExists(path) {
		t.Error("asset path should exist", path)
	}
}

func TestAssetPathCanBeSetInEnv(t *testing.T) {
	os.Setenv("ASSET_PATH", "/tmp")

	path, err := AssetPath("")
	if err != nil {
		t.Error("unexpected err", err)
	}

	if path != "/tmp/" {
		t.Error("asset path should be overridden", path)
	}

	os.Unsetenv("ASSET_PATH")
}
