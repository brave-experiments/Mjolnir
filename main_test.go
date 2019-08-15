package main

import (
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

func TestNewApp(t *testing.T) {
	app := NewApp()
	assert.Equal(t, app.Name, os.Getenv("CLI_NAME"))
	assert.Equal(t, app.Version, os.Getenv("CLI_VERSION"))
	assert.Equal(t, app.Description, os.Getenv("CLI_DESCRIPTION"))

	// Test that not set env variables wont fail initialization
	assert.Equal(t, app.Author, "")
}
