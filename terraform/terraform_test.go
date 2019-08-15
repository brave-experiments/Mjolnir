package terraform

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTerraform_New(t *testing.T) {
	terraform, err := Terraform{}.New()

	if nil == terraform.Ui {
		t.Error("Ui in terraform is nil")
	}

	if nil == terraform.CliRunner {
		t.Error("CliRunner in terraform is nil")
	}

	assert.Nil(t, err)
}
