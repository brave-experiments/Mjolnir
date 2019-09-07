package terra

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDefaultStateFileWithCreationOfFile(t *testing.T) {
	// Change default state for test purposes
	StateFileName = "dummy.tfstate"
	StateFileBody = ProperOutputFixture
	defaultStateFile, err := DefaultStateFile()
	assert.Nil(t, err)
	assert.IsType(t, &StateFile{}, defaultStateFile)
	assert.Equal(t, len(StateFileBody), len(defaultStateFile.Body))

	restoreFilesAndState(t)
}

func TestDefaultStateFileWithoutCreationOfFile(t *testing.T) {
	// Change default state for test purposes and create file
	StateFileName = "dummy.tfstate"
	fileBodyToTest := "Simple Dummy File Body"
	fileBody, err := os.Create(StateFileName)
	assert.Nil(t, err)
	bytesLen, err := fileBody.Write([]byte(fileBodyToTest))
	assert.Equal(t, len([]byte(fileBodyToTest)), bytesLen)
	assert.Nil(t, err)

	defaultStateFile, err := DefaultStateFile()
	assert.Nil(t, err)
	assert.IsType(t, &StateFile{}, defaultStateFile)

	err = defaultStateFile.ReadFile()
	assert.Empty(t, err)
	assert.Equal(t, fileBodyToTest, defaultStateFile.Body)

	restoreFilesAndState(t)
}

func TestDefaultStateFileFailure(t *testing.T) {
	StateFileName = "/some/dir/that/does/not/exists/dummy.tfstate"
	defaultStateFile, err := DefaultStateFile()
	assert.Error(t, err)
	assert.Equal(
		t,
		err.Error(),
		fmt.Sprintf("open %s: no such file or directory", StateFileName),
	)
	assert.IsType(t, &StateFile{}, defaultStateFile)
	StateFileName = DefaulStateFileName
}

func restoreFilesAndState(t *testing.T) {
	err := os.Remove(StateFileName)
	assert.Nil(t, err)
	StateFileName = DefaulStateFileName
	StateFileBody = DefaultStateFileBody
}
