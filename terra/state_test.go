package terra

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDefaultStateFileWithCreationOfFile(t *testing.T) {
	// Change default state for test purposes
	StateFileName = "dummy.tfstate"
	defaultStateFile, err := DefaultStateFile()
	assert.Nil(t, err)
	assert.IsType(t, &StateFile{}, defaultStateFile)

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

func restoreFilesAndState(t *testing.T) {
	err := os.Remove(StateFileName)
	assert.Nil(t, err)
	StateFileName = DefaulStateFileName
}
