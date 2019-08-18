package terra

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_DefaultClient(t *testing.T) {
	client := Client{}
	client.DefaultClient()
	assert.Greater(t, len(client.Recipes.Elements), 0)
}
