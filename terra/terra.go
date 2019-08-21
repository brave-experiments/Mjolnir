package terra

import (
	"github.com/hashicorp/terraform/terraform"
	"github.com/johandry/terranova"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

type Client struct {
	platform *terranova.Platform
	state    *StateFile
}

func (client *Client) Apply(file File, destroy bool) (err error) {
	if nil != client.PreparePlatform(file) {
		return err
	}

	err = client.platform.Apply(destroy)

	if nil != err {
		return err
	}

	err = client.WriteStateToFile()

	if nil != err {
		return err
	}

	// Synchronize state from platform and state object
	err = client.state.ReadFile()

	return err
}

func (client *Client) PreparePlatform(file File) (err error) {
	err = file.ReadFile()

	if nil != err {
		return err
	}

	err = client.assignVariables(file)

	if nil != err {
		return err
	}

	client.platform.Code = file.Body

	return err
}

func (client *Client) WriteStateToFile() (err error) {
	err = client.guard()

	if nil != err {
		return err
	}

	client.platform, err = client.platform.WriteStateToFile(client.state.Location)

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
	client.platform = &terranova.Platform{
		Providers: make(map[string]terraform.ResourceProvider),
	}
	client.platform.AddProvider(DefaultProvider("aws"))

	client.state, err = DefaultStateFile()

	if nil != err {
		return err
	}

	err = client.assignStateFile()

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

	if nil == client.state {
		return ClientError{"No state file found"}
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

func (client *Client) assignStateFile() (err error) {
	err = client.guard()

	if nil != err {
		return err
	}

	client.platform, err = client.platform.ReadStateFromFile(client.state.Location)

	return err
}
