package terra

import (
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"github.com/terraform-providers/terraform-provider-local/local"
	"github.com/terraform-providers/terraform-provider-random/random"
)

func DefaultProvider(key string) (returnKey string, provider terraform.ResourceProvider) {
	provider = aws.Provider()

	return key, provider
}

func RandomProvider(key string) (returnKey string, provider terraform.ResourceProvider) {
	provider = random.Provider()

	return key, provider
}

func LocalProvider(key string) (returnKey string, provider terraform.ResourceProvider) {
	provider = local.Provider()

	return key, provider
}
