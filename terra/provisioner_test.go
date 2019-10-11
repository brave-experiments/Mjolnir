package terra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoteProvisioner(t *testing.T) {
	expectedKey := "remote-exec"
	key, provisioner := RemoteProvisioner(expectedKey)
	assert.Equal(t, expectedKey, key)
	assert.IsType(t, localexec.Provisioner(), provisioner)
}
