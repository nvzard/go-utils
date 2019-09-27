package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/keptn/go-utils/pkg/models"
)

// ConfigService represents the interface for accessing the configuration service
type ConfigService interface {
	getBaseURL() string
	getAuthToken() string
	getAuthHeader() string
	getHTTPClient() *http.Client
}

func post(uri string, data []byte, c ConfigService) (*models.Error, error) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	addAuthHeader(req, c)

	resp, err := c.getHTTPClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respErr models.Error
	err = json.Unmarshal(body, &respErr)
	if err != nil {
		return nil, err
	}

	return &respErr, nil
}

func delete(uri string, c ConfigService) (*models.Error, error) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("DELETE", uri, nil)
	req.Header.Set("Content-Type", "application/json")
	addAuthHeader(req, c)

	resp, err := c.getHTTPClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respErr models.Error
	err = json.Unmarshal(body, &respErr)
	if err != nil {
		return nil, err
	}

	return &respErr, nil
}

func get(uri string, c ConfigService) (*models.Project, error) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Set("Content-Type", "application/json")
	addAuthHeader(req, c)
	fmt.Println(uri)
	resp, err := c.getHTTPClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil, nil
	}
	fmt.Println("read all")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("marshal")
	var respProject models.Project
	err = json.Unmarshal(body, &respProject)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println("done")
	return &respProject, nil
}

func addAuthHeader(req *http.Request, c ConfigService) {
	if c.getAuthHeader() != "" && c.getAuthToken() != "" {
		req.Header.Set(c.getAuthHeader(), c.getAuthToken())
	}
}
