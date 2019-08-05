package spi

import "testing"

func TestNewInstance(t *testing.T) {
	app, err := NewInstance("../config/app.toml")
	if err != nil {
		t.Fail()
	}
	t.Log(app.DB)
}
