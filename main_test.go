package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	cli := New()
	assert.Equal(t, cli.Name, os.Getenv("CLI_NAME"))
	assert.Equal(t, cli.Commands, RegisteredCommands)
}
