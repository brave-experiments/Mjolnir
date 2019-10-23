package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"

	"github.com/brave-experiments/Mjolnir/connect"
	"github.com/brave-experiments/Mjolnir/terra"
	"github.com/mitchellh/cli"
)

const (
	ExitCodeSuccess               = 0
	ExitCodeInvalidNumOfArgs      = 1
	ExitCodeInvalidSetup          = 2
	ExitCodeInvalidArgument       = 3
	ExitCodeTerraformError        = 4
	ExitCodeYamlBindingError      = 5
	ExitCodeEnvUnbindingError     = 6
	ExitCodeNoTempDirectory       = 7
	ExitCodeNoBastionIp           = 8
	ExitCodeSshKeyNotPresent      = 9
	ExitCodeSshError              = 10
	ExitCodeSshDialError          = 11
	ExitCodeTerraformDestroyError = 12
)

var (
	RegisteredCommands = map[string]cli.CommandFactory{
		"apply":    ApplyCmdFactory,
		"destroy":  DestroyCmdFactory,
		"bastion":  SshCmdFactory,
		"node":     NodeSshCmdFactory,
		"nodeinfo": NodeInfoSshCmdFactory,
		"geth":     GethCmdFactory,
	}
)

type ApplyCmd struct {
	cli.Command
	Recipes terra.Recipes
}

type DestroyCmd struct {
	ApplyCmd
}

type SshCmd struct {
	cli.Command
}

type NodeSshCmd struct {
	SshCmd
}

type NodeInfoSshCmd struct {
	SshCmd
}

type GethCmd struct {
	SshCmd
}

func GethCmdFactory() (command cli.Command, err error) {
	return GethCmd{}, nil
}

func NodeSshCmdFactory() (command cli.Command, err error) {
	return NodeSshCmd{}, nil
}

func NodeInfoSshCmdFactory() (command cli.Command, err error) {
	return NodeInfoSshCmd{}, nil
}

func SshCmdFactory() (command cli.Command, err error) {
	return SshCmd{}, nil
}

func ApplyCmdFactory() (command cli.Command, err error) {
	recipes := terra.Recipes{}
	recipes.CreateWithDefaults()

	command = ApplyCmd{
		Recipes: recipes,
	}

	return command, err
}

func DestroyCmdFactory() (command cli.Command, err error) {
	recipes := terra.Recipes{}
	recipes.CreateWithDefaults()
	command = DestroyCmd{
		ApplyCmd: ApplyCmd{
			Recipes: recipes,
		},
	}

	return command, err
}

func (gethCmd GethCmd) Run(args []string) (exitCode int) {
	bastionScriptName := "Node"
	exitCode = gethCmd.runWithScriptLocation(bastionScriptName, args)

	return exitCode
}

func (nodeSshCmd NodeSshCmd) Run(args []string) (exitCode int) {
	bastionScriptName := "NodeSsh"
	exitCode = nodeSshCmd.runWithScriptLocation(bastionScriptName, args)

	return exitCode
}

func (nodeSshCmd NodeInfoSshCmd) Run(args []string) (exitCode int) {
	exitCode = nodeSshCmd.getNodeInfo()

	return exitCode
}

func (sshCmd SshCmd) Run(args []string) (exitCode int) {
	if runtime.GOOS == "windows" {
		panic("Cannot run ssh command on Windows")
	}

	desiredBastionKey := "bastion_host_ip"
	directoryLocator := terra.TempDirPathLocation
	directoriesList, err := ioutil.ReadDir(directoryLocator)

	if nil != err || len(directoriesList) < 1 {
		fmt.Printf("%s temp directory not present \n", directoryLocator)
		return ExitCodeNoTempDirectory
	}

	err, bastionIp := terra.ReadOutputLogVar(desiredBastionKey)

	if nil != err {
		fmt.Println(err)
		return ExitCodeNoBastionIp
	}

	sshClient := connect.SshClient{}
	adminUser := "admin"

	err, certLocation := terra.ReadSshLocation()

	if nil != err {
		fmt.Println(err)
		return ExitCodeSshKeyNotPresent
	}

	sshClient.New(adminUser, bastionIp, certLocation)

	if nil != err {
		fmt.Println(err)
		return ExitCodeSshError
	}

	err = sshClient.Dial(args)

	if nil != err {
		fmt.Println(err)
		return ExitCodeSshDialError
	}

	return ExitCodeSuccess
}

func (destroyCmd DestroyCmd) Run(args []string) (exitCode int) {
	expectedMinimumArguments := 1

	if len(args) < expectedMinimumArguments {
		fmt.Printf(
			"Not enough arguments, expected more than %v \n write `destroy --help` for more info\n",
			expectedMinimumArguments,
		)

		return ExitCodeInvalidNumOfArgs
	}

	yamlFilePath := args[0]
	recipe, exitCode := destroyCmd.getRecipe(terra.DestroyDefaultRecipeVar)

	if exitCode != ExitCodeSuccess {
		fmt.Printf("Exited with code: %v, recipe not found", exitCode)
	}

	err := recipe.BindYamlWithVars(yamlFilePath)

	if nil != err {
		fmt.Println(err)
		return ExitCodeYamlBindingError
	}

	err = destroyCmd.ApplyCmd.executeTerra(recipe, true)
	exitCode = ExitCodeSuccess

	if nil != err {
		fmt.Println(err)
		exitCode = ExitCodeTerraformDestroyError
	}

	fmt.Println("Restoring env variables.")
	err = destroyCmd.ApplyCmd.restoreEnvVariables(recipe)

	if nil != err {
		fmt.Printf("Error restoring variables: %s", err)
		exitCode = ExitCodeEnvUnbindingError
	}

	fmt.Printf("Destroy complete! Resources in network %s destroyed.\n", recipe.Variables["network_name"])

	return exitCode
}

func (applyCmd ApplyCmd) Run(args []string) (exitCode int) {
	expectedMinimumArguments := 2

	if len(args) < expectedMinimumArguments {
		fmt.Printf(
			"Not enough arguments, expected more than %v \n write `apply --help` for more info\n",
			expectedMinimumArguments,
		)

		return ExitCodeInvalidNumOfArgs
	}

	if nil == &applyCmd.Recipes || len(applyCmd.Recipes.Elements) < 1 {
		fmt.Println("Invalid recipes configuration")

		return ExitCodeInvalidSetup
	}

	recipeKey := args[0]
	recipe, exitCode := applyCmd.getRecipe(recipeKey)

	if exitCode > ExitCodeSuccess {
		return exitCode
	}

	yamlFilePath := args[1]
	err := recipe.BindYamlWithVars(yamlFilePath)

	if nil != err {
		fmt.Println(err)
		return ExitCodeYamlBindingError
	}

	fmt.Printf(
		"Executing %s \n",
		recipeKey,
	)

	err = applyCmd.executeTerra(recipe, false)
	exitCode = ExitCodeSuccess

	if nil != err {
		fmt.Println(err)
		exitCode = ExitCodeTerraformError
	}

	fmt.Println("Restoring env variables.")
	err = applyCmd.restoreEnvVariables(recipe)

	if nil != err {
		fmt.Printf("Error restoring variables: %s", err)
		exitCode = ExitCodeEnvUnbindingError
	}

	return exitCode
}

func (sshCmd SshCmd) Help() (helpMessage string) {
	return "Ssh into bastion. .mjolnir directory must be present (previous deploy must complete)"
}

func (applyCmd ApplyCmd) Help() (helpMessage string) {
	helpMessage = "\nThis is apply command. Usage: mjolnir apply [recipe] [yamlFilePath]\n"

	if nil != &applyCmd.Recipes {
		helpMessage = helpMessage + "\nAvailable Recipes: \n"

		for recipeKey := range applyCmd.Recipes.Elements {
			helpMessage = helpMessage + "\n \t" + recipeKey + "\n"
		}
	}

	helpMessage = helpMessage + "\n" + "Filepath must be valid yaml file with variables, like:" + "\n"
	helpMessage = helpMessage + "\n" + terra.SchemaV02 + "\n"

	return helpMessage
}

func (destroyCmd DestroyCmd) Help() (helpMessage string) {
	helpMessage = "\nThis is destroy command. Usage: mjolnir destroy [yamlFilePath]"
	helpMessage = helpMessage + "\nThere must be a valid terraform.tfstate file at root of execution"

	return helpMessage
}

func (nodeSshCmd NodeSshCmd) Help() (helpMessage string) {
	helpMessage = "\n This command let you attach via ssh to certain node\n"
	helpMessage = helpMessage + "You must provide node number as argument. If number is out of range, ssh will fail\n"
	helpMessage = helpMessage + "Example: mjolnir node 1"

	return helpMessage
}

func (nodeSshCmd NodeInfoSshCmd) Help() (helpMessage string) {
	helpMessage = "\n This command let you to get detailed info about runing nodes\n"
	helpMessage = helpMessage + "You must not provide any arguments. There are not arguments in this command\n"
	helpMessage = helpMessage + "Example: mjolnir nodeinfo"

	return helpMessage
}

func (gethCmd GethCmd) Help() (helpMessage string) {
	helpMessage = "\n This command let you attach via rpc (geth) to certain node\n"
	helpMessage = helpMessage + "You must provide node number as argument. If number is out of range, ssh will fail\n"
	helpMessage = helpMessage + "Example: mjolnir geth 1"

	return helpMessage
}

func (applyCmd ApplyCmd) Synopsis() (synopsis string) {
	synopsis = "apply [recipe] [yamlSchemaPath]"
	return synopsis
}

func (destroyCmd DestroyCmd) Synopsis() (synopsis string) {
	synopsis = "destroy [yamlSchemaPath]"
	return synopsis
}

func (sshCmd SshCmd) Synopsis() (synopsis string) {
	synopsis = "bastion [arguments]"
	return synopsis
}

func (nodeSshCmd NodeSshCmd) Synopsis() (synopsis string) {
	synopsis = "node [number]"
	return synopsis
}

func (nodeSshCmd NodeInfoSshCmd) Synopsis() (synopsis string) {
	synopsis = "nodeinfo"
	return synopsis
}

func (gethCmd GethCmd) Synopsis() (synopsis string) {
	synopsis = "geth [number]"
	return synopsis
}

func (applyCmd ApplyCmd) printRecipesKeys() {
	if nil == &applyCmd.Recipes {
		return
	}

	for key := range applyCmd.Recipes.Elements {
		fmt.Println(key)
	}
}

func (applyCmd *ApplyCmd) executeTerra(recipe terra.CombinedRecipe, destroy bool) (err error) {
	terraClient := terra.Client{}
	err = terraClient.DefaultClient()

	if nil != err {
		return err
	}

	err = terraClient.ApplyCombined(recipe, destroy)

	if true == destroy {
		err = os.RemoveAll(terra.TempDirPathLocation)
	}

	return err
}

func (applyCmd *ApplyCmd) getRecipe(recipeKey string) (recipe terra.CombinedRecipe, exitCode int) {
	recipes := applyCmd.Recipes.Elements
	recipe, contains := recipes[recipeKey]

	if false == contains {
		fmt.Printf(
			"Recipe %s not found within recipes, available are: \n",
			recipeKey,
		)
		applyCmd.printRecipesKeys()

		return terra.CombinedRecipe{}, ExitCodeInvalidArgument
	}

	return recipe, ExitCodeSuccess
}

func (applyCmd *ApplyCmd) restoreEnvVariables(recipe terra.CombinedRecipe) (err error) {
	return recipe.UnbindEnvVars()
}

func (sshCmd *SshCmd) runWithScriptLocation(scriptName string, args []string) (exitCode int) {
	desiredArgsLen := 1

	if len(args) < desiredArgsLen {
		fmt.Println("Please provide node number to attach to")
		return ExitCodeInvalidNumOfArgs
	}

	nodeNumber, err := strconv.ParseInt(args[0], 10, 64)

	if err != nil {
		fmt.Println(err)
		return ExitCodeInvalidArgument
	}

	bastionSshScriptLocator := fmt.Sprintf("/usr/local/bin/%s%v", scriptName, nodeNumber)
	additionalSshCmdArgs := []string{bastionSshScriptLocator}
	coreCmd, err := SshCmdFactory()

	if nil != err {
		fmt.Println(err)
		return ExitCodeInvalidSetup
	}

	exitCode = coreCmd.Run(additionalSshCmdArgs)

	return exitCode
}

func (sshCmd *SshCmd) getNodeInfo() (exitCode int) {
	bastionSshNodeInfoLocator := "cat /qdata/nodeinfo/ip_* | sort"
	additionalSshCmdArgs := []string{bastionSshNodeInfoLocator}
	coreCmd, err := SshCmdFactory()

	if nil != err {
		fmt.Println(err)
		return ExitCodeInvalidSetup
	}

	exitCode = coreCmd.Run(additionalSshCmdArgs)

	return exitCode
}
