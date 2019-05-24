package config

import "testing"

func TestAppRoot(t *testing.T) {
	t.Logf("root: %s", findAssetPath())
}

func TestConfigLoad(t *testing.T) {
	t.Log(cfg.Pumps)
}
