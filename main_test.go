package main

import (
	"github.com/brave-experiments/apollo-devops/terra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	cli := New()
	assert.Equal(t, cli.Name, terra.StaticCliCliName)
	assert.Equal(t, cli.Version, terra.StaticCliCliVersion)
	assert.Equal(t, cli.Commands, RegisteredCommands)
}
