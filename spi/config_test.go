package spi

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	dir, err := os.Getwd()
	t.Log(dir)
	config, err := Load("../config/app.toml.example")
	if err != nil {
		t.Fail()
	}
	t.Log(config)
}
