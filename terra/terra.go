package terra

import (
	"github.com/brave-experiments/apollo-devops/src/github.com/hashicorp/terraform/terraform"
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
}

func defaultProvider() terraform.ResourceProvider {
	provider := aws.Provider()

	return provider
	//https://github.com/johandry/terranova
}
