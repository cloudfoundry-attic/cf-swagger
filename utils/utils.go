package utils

import (
	"bytes"
	"fmt"
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

	fmt.Println("sending back file content")
	return bytes.NewBuffer(file), nil
}
