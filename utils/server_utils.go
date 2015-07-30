package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"text/template"
)

const (
	SERVER_URL         = "localhost:8001"
	TEMPLATE_ROOT_PATH = "templates"
	NON_VERBOSE        = "NON_VERBOSE"
)

type HttpClient interface {
	DoRawHttpRequest(path string, requestType string, requestBody *bytes.Buffer) ([]byte, error)
	GenerateRequestBody(templateData interface{}) (*bytes.Buffer, error)
	HasErrors(body map[string]interface{}) error
	CheckForHttpResponseErrors(data []byte) error
}

type httpClient struct {
	username     string
	apiKey       string
	templatePath string

	httpClient *http.Client

	nonVerbose bool
}

func NewHttpClient(username, apiKey string) *httpClient {
	pwd, _ := os.Getwd()
	httpc := &httpClient{
		username: username,
		apiKey:   apiKey,

		templatePath: filepath.Join(pwd, TEMPLATE_ROOT_PATH),

		httpClient: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		},

		nonVerbose: checkNonVerbose(),
	}

	return httpc
}

func (httpc *httpClient) DoRawHttpRequest(path string, requestType string, requestBody *bytes.Buffer) ([]byte, error) {
	url := fmt.Sprintf("http://%s/%s", SERVER_URL, path)
	return httpc.makeHttpRequest(url, requestType, requestBody)
}

func (httpc *httpClient) GenerateRequestBody(templateData interface{}) (*bytes.Buffer, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	bodyTemplate := template.Must(template.ParseFiles(filepath.Join(cwd, httpc.templatePath)))
	body := new(bytes.Buffer)
	bodyTemplate.Execute(body, templateData)

	return body, nil
}

func (httpc *httpClient) HasErrors(body map[string]interface{}) error {
	if errString, ok := body["error"]; !ok {
		return nil
	} else {
		return errors.New(errString.(string))
	}
}

func (httpc *httpClient) CheckForHttpResponseErrors(data []byte) error {
	var decodedResponse map[string]interface{}
	err := json.Unmarshal(data, &decodedResponse)
	if err != nil {
		return err
	}

	if err := httpc.HasErrors(decodedResponse); err != nil {
		return err
	}

	return nil
}

//Private methods

func (httpc *httpClient) makeHttpRequest(url string, requestType string, requestBody *bytes.Buffer) ([]byte, error) {
	req, err := http.NewRequest(requestType, url, requestBody)
	if err != nil {
		return nil, err
	}

	bs, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, err
	}

	if !httpc.nonVerbose {
		fmt.Fprintf(os.Stderr, "\n---\n[http-go] Request:\n%s\n", string(bs))
	}

	resp, err := httpc.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bs, err = httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}

	if !httpc.nonVerbose {
		fmt.Fprintf(os.Stderr, "[http-go] Response:\n%s\n", string(bs))
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func checkNonVerbose() bool {
	nonVerbose := os.Getenv(NON_VERBOSE)
	switch nonVerbose {
	case "yes":
		return true
	case "YES":
		return true
	case "true":
		return true
	case "TRUE":
		return true
	}

	return false
}
