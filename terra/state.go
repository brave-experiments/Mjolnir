package terra

import "os"

const (
	DefaulStateFileName  = "terraform.tfstate"
	DefaultStateFileBody = `{
    "version": 3,
    "terraform_version": "0.11.13",
    "serial": 1,
    "outputs": {},
    "resources": []
}`
)

var (
	StateFileName = DefaulStateFileName
)

type StateFile struct {
	File
}

func DefaultStateFile() (stateFile *StateFile, err error) {
	defaultStateFile := new(StateFile)
	defaultStateFile.Location = StateFileName

	err = defaultStateFile.ReadFile()

	if nil == err && len(defaultStateFile.Body) > 0 {
		return defaultStateFile, err
	}

	fileBody, err := os.Create(StateFileName)

	if nil != err {
		return stateFile, err
	}

	// Write default state file if current is empty
	if len(defaultStateFile.Body) == 0 {
		_, err = fileBody.Write([]byte(DefaultStateFileBody))
	}

	if nil != err {
		return defaultStateFile, err
	}

	err = defaultStateFile.ReadFile()

	return defaultStateFile, err
}
