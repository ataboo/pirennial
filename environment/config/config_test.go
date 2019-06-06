package config

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseTOML(t *testing.T) {
	outStruct := struct {
		Str   string
		Array []string
	}{}

	rawToml := `str = "str value"
	array = [ "first", "second", "third" ]`

	err := parseTOML([]byte(rawToml), &outStruct)

	if err != nil {
		t.Error("failed to parse", err)
	}

	if outStruct.Array[1] != "second" {
		t.Error(`second item in Array should be "second"`, outStruct)
	}

	if outStruct.Str != "str value" {
		t.Error("unexpected value in struct")
	}
}

func TestLoadTOMLFile(t *testing.T) {
	outStruct := struct {
		Str   string
		Array []string
	}{}

	rawToml := `str = "str value"
	array = [ "first", "second", "third" ]`

	ioutil.WriteFile("/tmp/pirennial_toml_test.toml", []byte(rawToml), os.ModePerm)

	os.Setenv("ASSET_PATH", "/tmp")

	err := LoadTOMLFile("pirennial_toml_test.toml", &outStruct)

	if err != nil {
		t.Error("failed to load test toml file", err)
	}

	if outStruct.Array[1] != "second" {
		t.Error(`second item in Array should be "second"`, outStruct)
	}

	if outStruct.Str != "str value" {
		t.Error("unexpected value in struct")
	}

	os.Unsetenv("ASSET_PATH")
}

func TestLoadNonExistantFile(t *testing.T) {
	err := LoadTOMLFile("file_should_not_exist.toml", &struct{}{})

	if err == nil {
		t.Error("should throw error")
	}
}
