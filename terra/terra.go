package terra

import (
	"github.com/hashicorp/terraform/terraform"
	"github.com/johandry/terranova"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

type Client struct {
	Recipes  Recipes
	Provider terraform.ResourceProvider
	state    *StateFile
	platform *terranova.Platform
}

func (client *Client) RunPlatform(file File) (err error) {
	err = file.ReadFile()

	if nil != err {
		return err
	}

	err = client.assignVariables(file)

	return err
}

func (client *Client) DumpVariables() (vars map[string]interface{}, err error) {
	err = client.guard()

	if nil != err {
		return nil, err
	}

	return client.platform.Vars, err
}

func (client *Client) DefaultClient() (err error) {
	// Maybe recipes should not be registered within client, but CLI ?
	client.Recipes = Recipes{}
	client.Recipes.CreateWithDefaults()
	client.platform = &terranova.Platform{
		Providers: make(map[string]terraform.ResourceProvider),
	}
	client.platform.AddProvider(DefaultProvider("aws"))

	state, err := DefaultStateFile()
	client.state = state

	return err
}

func DefaultProvider(key string) (returnKey string, provider terraform.ResourceProvider) {
	provider = aws.Provider()

	return key, provider
}

func (client *Client) guard() (err error) {
	if nil == client.platform {
		return ClientError{"Platform is not initialized"}
	}

	return nil
}

func (client *Client) assignVariables(file File) (err error) {
	err = client.guard()

	if nil != err {
		return err
	}

	for key, value := range file.Variables {
		client.platform.Var(key, value)
	}

	return nil
}
