package connect

import (
	"fmt"
	"github.com/brave-experiments/apollo-devops/terra"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const (
	ValidKeyPath = "dummy_id_rsa"
)

func TestSshClient_NewFailure(t *testing.T) {
	client := SshClient{}
	invalidKeyPath := "invalidKeyPath"
	user := "dummyUser"
	host := "dummyHost"
	err := client.New(user, host, invalidKeyPath)
	assert.Error(t, err)
	assert.Equal(t, fmt.Sprintf("open %s: no such file or directory", invalidKeyPath), err.Error())
}

func TestSshClient_New(t *testing.T) {
	client := SshClient{}
	user := "dummyUser"
	host := "dummyHost"
	createPrivKey(t, ProperPrivateKey)
	err := client.New(user, host, ValidKeyPath)
	assert.Nil(t, err)
	deletePrivKey(t)
}

func TestSshClient_DialFailure(t *testing.T) {
	terra.TempDirPathLocation = "dummyPath"
	client := SshClient{}
	dialClient, err := client.Dial()
	assert.Error(t, err)
	assert.Nil(t, dialClient)
	terra.TempDirPathLocation = terra.TempDirPath
}

func TestDialError_Error(t *testing.T) {
	message := "dummyMessage"
	dialError := DialError{Message: message}
	assert.Equal(t, message, dialError.Error())
}

func createPrivKey(t *testing.T, keyBody string) {
	err := ioutil.WriteFile(ValidKeyPath, []byte(keyBody), os.FileMode(0400))
	assert.Nil(t, err)
}

func deletePrivKey(t *testing.T) {
	err := os.RemoveAll(ValidKeyPath)
	assert.Nil(t, err)
}
