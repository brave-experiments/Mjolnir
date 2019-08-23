package terra

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-aws/aws"
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
	return key, &schema.Provider{
		Schema: map[string]*schema.Schema{},
		ResourcesMap: map[string]*schema.Resource{
			"local_file": resourceLocalFile(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"local_file": dataSourceLocalFile(),
		},
	}
}
