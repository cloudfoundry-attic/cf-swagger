package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadTestFixtures(fileName string) (*bytes.Buffer, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(filepath.Join(wd, "../..", "test_fixtures", fileName))
	if err != nil {
		return nil, err
	}

	jsonParams, err := json.Marshal(file)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(jsonParams), nil
}
