package terra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_DefaultModules(t *testing.T) {
	client := Client{}
	err := client.DefaultModules(FetchedModules)
	assert.Nil(t, err)
}
