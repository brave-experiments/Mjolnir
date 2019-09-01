package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	cli := New()
	assert.Equal(t, cli.Name, CliName)
	assert.Equal(t, cli.Commands, RegisteredCommands)
}
