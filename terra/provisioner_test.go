package terra

import (
	"github.com/brave-experiments/apollo-devops/src/github.com/stretchr/testify/assert"
	"github.com/hashicorp/terraform/builtin/provisioners/local-exec"
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
