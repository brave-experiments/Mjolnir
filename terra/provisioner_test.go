package terra

import (
	"github.com/hashicorp/terraform/builtin/provisioners/local-exec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalProvisioner(t *testing.T) {
	expectedKey := "local-exec"
	key, provisioner := LocalProvisioner(expectedKey)
	assert.Equal(t, expectedKey, key)
	assert.IsType(t, localexec.Provisioner(), provisioner)
}

func TestRemoteProvisioner(t *testing.T) {
	expectedKey := "remote-exec"
	key, provisioner := RemoteProvisioner(expectedKey)
	assert.Equal(t, expectedKey, key)
	assert.IsType(t, localexec.Provisioner(), provisioner)
}
