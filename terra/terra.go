package terra

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/terraform"
	"github.com/johandry/terranova"
	"os"
)

const (
	CombinedRecipeDefaultFileName = "temp.tf"
	LastExecutedVariablesFileName = "variables.log"
	TempDirPath                   = ".apollo"
)

var (
	LastExecutedFileName = LastExecutedVariablesFileName
	TempDirPathLocation  = TempDirPath
)

type Client struct {
	platform *terranova.Platform
	state    *StateFile
}

func (client *Client) CreateDirInTemp(dirName string) (location string, err error) {
	fullDirPath := fmt.Sprintf("%s/%s", TempDirPathLocation, dirName)
	err = os.MkdirAll(fullDirPath, 0777)

	if nil != err {
		return "", err
	}

	return fullDirPath, err
}

func (client *Client) ApplyCombined(recipe CombinedRecipe, destroy bool) (err error) {
	if len(recipe.Body) < 1 {
		err = recipe.ParseBody()
	}

	if nil != err {
		return err
	}

	if nil != &recipe.Location {
		recipe.Location = CombinedRecipeDefaultFileName
	}

	err = recipe.WriteFile()

	if nil != err {
		return err
	}

	file := File{
		Location:  recipe.Location,
		Body:      recipe.Body,
		Variables: recipe.Variables,
	}

	err = client.Apply(file, destroy)

	return err
}

func (client *Client) Apply(file File, destroy bool) (err error) {
	if nil != client.PreparePlatform(file) {
		return err
	}

	executableFile := File{
		Location: LastExecutedFileName,
		Body:     fmt.Sprintf("Last executed variables in recipe: \n%s", file.Variables),
	}

	err = executableFile.WriteFile()

	if nil != err {
		return err
	}

	err = client.platform.Apply(destroy)

	if nil != err {
		_ = client.WriteStateToFiles()

		return err
	}

	err = client.WriteStateToFiles()

	if nil != err {
		return err
	}

	// Synchronize state from platform and state object
	err = client.state.ReadFile()

	consoleOutputsFromTerraState := client.outputsAsString(true)
	fmt.Println(consoleOutputsFromTerraState)

	return err
}

func (client *Client) PreparePlatform(file File) (err error) {
	err = file.ReadFile()

	if nil != err {
		return err
	}

	err = client.assignVariables(file)

	if nil != err {
		return err
	}

	client.platform.Code = file.Body

	return err
}

func (client *Client) WriteStateToFiles() (err error) {
	err = client.guard()

	if nil != err {
		return err
	}

	client.platform, err = client.platform.WriteStateToFile(client.state.Location)

	if nil != err {
		return err
	}

	currentKeyPair := keyPair{}
	jsonBytes, err := json.Marshal(client.platform.State)
	currentKeyPair.FromJson(string(jsonBytes))
	err = currentKeyPair.Save()

	if nil != err {
		fmt.Println(err.Error())

		return err
	}

	return
}

func (client *Client) DumpVariables() (vars map[string]interface{}, err error) {
	err = client.guard()

	if nil != err {
		return nil, err
	}

	return client.platform.Vars, err
}

func (client *Client) DefaultClient() (err error) {
	client.addDependencies()

	client.state, err = DefaultStateFile()

	if nil != err {
		return err
	}

	err = client.assignStateFile()

	return err
}

func (client *Client) guard() (err error) {
	if nil == client.platform {
		return ClientError{"Platform is not initialized"}
	}

	if nil == client.state {
		return ClientError{"No state file found"}
	}

	return nil
}

func (client *Client) addDependencies() {
	client.platform = &terranova.Platform{
		Providers:    make(map[string]terraform.ResourceProvider),
		Provisioners: make(map[string]terraform.ResourceProvisioner),
	}
	client.platform.AddProvider(DefaultProvider("aws"))
	client.platform.AddProvider(LocalProvider("local"))
	client.platform.AddProvider(NullProvider("null"))
	client.platform.AddProvider(RandomProvider("random"))
	client.platform.AddProvider(TemplateProvider("template"))
	client.platform.AddProvider(TlsProvider("tls"))

	client.platform.AddProvisioner(LocalProvisioner("local-exec"))
	client.platform.AddProvisioner(RemoteProvisioner("remote-exec"))
}

func (client *Client) assignVariables(file File) (err error) {
	err = client.guard()

	if nil != err {
		return err
	}

	for key, value := range file.Variables {
		client.platform.Var(key, value)
	}

	return nil
}

func (client *Client) assignStateFile() (err error) {
	err = client.guard()

	if nil != err {
		return err
	}

	client.platform, err = client.platform.ReadStateFromFile(client.state.Location)

	return err
}

func (client *Client) outputsAsString(includeHeader bool) string {
	outputRecords := OutputRecords{}
	stateAsString := client.platform.State.String()

	return outputRecords.FromJsonAsString(stateAsString, true)
}
