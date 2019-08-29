package terra

import (
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"github.com/terraform-providers/terraform-provider-local/local"
	"github.com/terraform-providers/terraform-provider-null/null"
	"github.com/terraform-providers/terraform-provider-random/random"
	"github.com/terraform-providers/terraform-provider-tls/tls"
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

func NullProvider(key string) (returnKey string, provider terraform.ResourceProvider) {
	provider = null.Provider()

	return key, provider
}

func TlsProvider(key string) (returnKey string, provider terraform.ResourceProvider) {
	provider = tls.Provider()

	return key, provider
}
