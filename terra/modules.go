package terra

import (
	"github.com/hashicorp/terraform/config/module"
	"github.com/hashicorp/terraform/svchost/disco"
)

var (
	ModulesToRegister = map[string]string{
		"vpc": "terraform-aws-modules/vpc/aws",
	}
)

func (client *Client) DefaultModules(modulesDir string) (err error) {
	discovery := disco.New()

	if nil == client.storage {
		client.storage = module.NewStorage(modulesDir, discovery)
	}

	for key, modulePath := range ModulesToRegister {
		err := client.storage.GetModule(key, modulePath)

		if nil != err {
			return err
		}
	}

	return nil
}
