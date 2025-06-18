package helpers

import (
	"encoding/json"
	"os"
)

func Read(inputFile string) ([]byte, error) {
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return nil, err
	}

	fileData, err := os.ReadFile(inputFile)
	if err != nil {
		return nil, err
	}

	return fileData, nil
}

func ReadAsType[T any](inputFile string) (T, error) {
	fileData, err := Read(inputFile)
	var returnObject T
	if err != nil {
		return returnObject, err
	}
	err = json.Unmarshal(fileData, &returnObject)
	if err != nil {
		return returnObject, err
	}
	return returnObject, nil
}
