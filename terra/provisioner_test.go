package terra

import (
	"github.com/hashicorp/terraform/builtin/provisioners/remote-exec"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoteProvisioner(t *testing.T) {
	expectedKey := "remote-exec"
	key, provisioner := RemoteProvisioner(expectedKey)
	assert.Equal(t, expectedKey, key)
	assert.IsType(t, remoteexec.Provisioner(), provisioner)
}
