package terra

import (
	"github.com/hashicorp/terraform/builtin/provisioners/remote-exec"
	"github.com/hashicorp/terraform/terraform"
)

func RemoteProvisioner(key string) (returnKey string, provisioner terraform.ResourceProvisioner) {
	provisioner = remoteexec.Provisioner()

	return key, provisioner
}
