package terra

import (
	"github.com/stretchr/testify/assert"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"github.com/terraform-providers/terraform-provider-random/random"
	"testing"
)

func TestDefaultProvider(t *testing.T) {
	keyToTest := "dummy"
	key, provider := DefaultProvider(keyToTest)
	assert.Equal(t, keyToTest, key)
	assert.IsType(t, aws.Provider(), provider)
}

func TestRandomProvider(t *testing.T) {
	keyToTest := "dummy"
	key, provider := RandomProvider(keyToTest)
	assert.Equal(t, keyToTest, key)
	assert.IsType(t, random.Provider(), provider)
}

func TestLocalProvider(t *testing.T) {
	keyToTest := "dummy"
	key, provider := LocalProvider(keyToTest)
	assert.Equal(t, keyToTest, key)
	assert.IsType(t, aws.Provider(), provider)
}

func TestNullProvider(t *testing.T) {
	keyToTest := "dummy"
	key, provider := NullProvider(keyToTest)
	assert.Equal(t, keyToTest, key)
	assert.IsType(t, aws.Provider(), provider)
}

func TestTlsProvider(t *testing.T) {
	keyToTest := "dummy"
	key, provider := TlsProvider(keyToTest)
	assert.Equal(t, keyToTest, key)
	assert.IsType(t, aws.Provider(), provider)
}
