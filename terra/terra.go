package terra

import (
	"github.com/hashicorp/terraform/terraform"
	"github.com/johandry/terranova"
	"github.com/terraform-providers/terraform-provider-aws/aws"
)

type Client struct {
	Recipes  Recipes
	Provider terraform.ResourceProvider
	platform terranova.Platform
	state    terranova.State
}

func (client *Client) DefaultClient() {
	client.Recipes = Recipes{}
	client.Recipes.CreateWithDefaults()
	client.platform = terranova.Platform{
		Providers: make(map[string]terraform.ResourceProvider),
	}
	client.platform.AddProvider(DefaultProvider("aws"))
}

func DefaultProvider(key string) (returnKey string, provider terraform.ResourceProvider) {
	provider = aws.Provider()

	return key, provider
}
