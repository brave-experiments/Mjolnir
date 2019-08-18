package terra

import (
	"github.com/stretchr/testify/assert"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"testing"
)

func TestClient_DefaultClient(t *testing.T) {
	client := Client{}
	client.DefaultClient()
	assert.Greater(t, len(client.Recipes.Elements), 0)
}

func TestDefaultProvider(t *testing.T) {
	keyToTest := "dummy"
	key, provider := DefaultProvider(keyToTest)
	assert.Equal(t, keyToTest, key)
	assert.IsType(t, aws.Provider(), provider)
}
