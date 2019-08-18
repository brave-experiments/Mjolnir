package terra

import "os"

const (
	DefaulStateFileName = "default.tfstate"
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

	if nil != err {
		_, err = os.Create(StateFileName)
		err = defaultStateFile.ReadFile()
	}

	return defaultStateFile, err
}
