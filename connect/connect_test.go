package connect

import (
	"github.com/brave-experiments/apollo-devops/terra"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	ValidKeyPath = "dummy_id_rsa"
)

func TestSshClient_New(t *testing.T) {
	client := SshClient{}
	user := "dummyUser"
	host := "dummyHost"
	client.New(user, host, ValidKeyPath)
	assert.IsType(t, SshClient{}, client)
}

func TestSshClient_DialFailure(t *testing.T) {
	terra.TempDirPathLocation = "dummyPath"
	client := SshClient{}
	err := client.Dial()
	assert.Error(t, err)
	terra.TempDirPathLocation = terra.TempDirPath
}

func TestDialError_Error(t *testing.T) {
	message := "dummyMessage"
	dialError := DialError{Message: message}
	assert.Equal(t, message, dialError.Error())
}
