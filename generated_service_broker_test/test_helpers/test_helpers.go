package test_helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func ReadJsonTestFixtures(packageName, fileName string) ([]byte, error) {
	wd, _ := os.Getwd()
	return ioutil.ReadFile(filepath.Join(wd, "..", "test_fixtures", packageName, fileName))
}
