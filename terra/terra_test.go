package terra

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/terraform-providers/terraform-provider-aws/aws"
	"os"
	"testing"
)

func TestClient_DefaultClient(t *testing.T) {
	client := Client{}
	err := client.DefaultClient()
	assert.Greater(t, len(client.Recipes.Elements), 0)
	assert.Nil(t, err)
}

func TestDefaultProvider(t *testing.T) {
	keyToTest := "dummy"
	key, provider := DefaultProvider(keyToTest)
	assert.Equal(t, keyToTest, key)
	assert.IsType(t, aws.Provider(), provider)
}

func TestClient_RunPlatformFailure_RecipeDoesNotExist(t *testing.T) {
	fileName := "dummy.tf"
	client := Client{}
	file := File{
		Location: fileName,
	}
	err := client.RunPlatform(file)
	assert.Error(t, err)
	assert.IsType(t, &os.PathError{}, err)
	assert.Equal(
		t,
		err.Error(),
		fmt.Sprintf("open %s: no such file or directory", fileName),
	)
}

func TestClient_RunPlatformFailure_PlatformIsNotInitialized(t *testing.T) {
	fileName := "dummyRecipe.tf"
	fileBody := "dummy file body"
	PrepareDummyFile(t, fileName, fileBody)
	client := Client{}
	file := File{
		Location: fileName,
	}
	err := client.RunPlatform(file)
	assert.Error(t, err)
	assert.IsType(t, ClientError{}, err)
	assert.Equal(t, "Platform is not initialized", err.Error())
	RemoveDummyFile(t, fileName)
}

//func TestClient_RunPlatformWithVariables(t *testing.T) {
//	client := Client{}
//	err := client.DefaultClient()
//	assert.Nil(t, err)
//}

//func TestClient_InitializePlatform(t *testing.T) {
//	fileBody := `
//		variable "count"    { default = 2 }
//  		variable "key_name" {}
//        provider "aws" {
//            region        = "us-west-2"
//        }
//        resource "aws_instance" "server" {
//            instance_type = "t2.micro"
//            ami           = "ami-6e1a0117"
//            count         = "${var.count}"
//            key_name      = "${var.key_name}"
//        }
//    `
//
//	fileName := "dummyFileUniqueName.tf"
//
//	err := ioutil.WriteFile(fileName, []byte(fileBody), 0644)
//	assert.Nil(t, err)
//
//	file := File{
//		Location: fileName,
//	}
//}
