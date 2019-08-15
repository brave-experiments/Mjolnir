package terraform

import (
    "testing"
)

func TestTerraform_New(t *testing.T) {
    terraform := Terraform{}.New()

    if nil == terraform.Ui {
        t.Error("Ui in terraform is nil")
    }
}