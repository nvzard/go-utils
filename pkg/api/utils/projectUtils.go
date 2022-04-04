package api

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/keptn/go-utils/pkg/common/httputils"

	"github.com/keptn/go-utils/pkg/api/models"
)

const v1ProjectPath = "/v1/project"

type ProjectsV1Interface interface {
	CreateProject(project models.Project) (*models.EventContext, *models.Error)
	DeleteProject(project models.Project) (*models.EventContext, *models.Error)
	GetProject(project models.Project) (*models.Project, *models.Error)
	GetAllProjects() ([]*models.Project, error)
	UpdateConfigurationServiceProject(project models.Project) (*models.EventContext, *models.Error)
}

// ProjectHandler handles projects
type ProjectHandler struct {
	BaseURL    string
	AuthToken  string
	AuthHeader string
	HTTPClient *http.Client
	Scheme     string
}

// NewProjectHandler returns a new ProjectHandler which sends all requests directly to the configuration-service
func NewProjectHandler(baseURL string) *ProjectHandler {
	baseURL = httputils.TrimHTTPScheme(baseURL)
	return &ProjectHandler{
		BaseURL:    baseURL,
		AuthHeader: "",
		AuthToken:  "",
		HTTPClient: &http.Client{Transport: wrapOtelTransport(getClientTransport(nil))},
		Scheme:     "http",
	}
}

// NewAuthenticatedProjectHandler returns a new ProjectHandler that authenticates at the api via the provided token
// and sends all requests directly to the configuration-service
// Deprecated: use APISet instead
func NewAuthenticatedProjectHandler(baseURL string, authToken string, authHeader string, httpClient *http.Client, scheme string) *ProjectHandler {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	httpClient.Transport = wrapOtelTransport(getClientTransport(httpClient.Transport))
	return createAuthProjectHandler(baseURL, authToken, authHeader, httpClient, scheme)
}

func createAuthProjectHandler(baseURL string, authToken string, authHeader string, httpClient *http.Client, scheme string) *ProjectHandler {
	baseURL = strings.TrimPrefix(baseURL, "http://")
	baseURL = strings.TrimPrefix(baseURL, "https://")
	baseURL = strings.TrimRight(baseURL, "/")

	if !strings.HasSuffix(baseURL, shipyardControllerBaseURL) {
		baseURL += "/" + shipyardControllerBaseURL
	}

	return &ProjectHandler{
		BaseURL:    baseURL,
		AuthHeader: authHeader,
		AuthToken:  authToken,
		HTTPClient: httpClient,
		Scheme:     scheme,
	}
}

func (p *ProjectHandler) getBaseURL() string {
	return p.BaseURL
}

func (p *ProjectHandler) getAuthToken() string {
	return p.AuthToken
}

func (p *ProjectHandler) getAuthHeader() string {
	return p.AuthHeader
}

func (p *ProjectHandler) getHTTPClient() *http.Client {
	return p.HTTPClient
}

// CreateProject creates a new project
func (p *ProjectHandler) CreateProject(project models.Project) (*models.EventContext, *models.Error) {
	bodyStr, err := project.ToJSON()
	if err != nil {
		return nil, buildErrorResponse(err.Error())
	}
	return postWithEventContext(context.TODO(), p.Scheme+"://"+p.getBaseURL()+v1ProjectPath, bodyStr, p)
}

// DeleteProject deletes a project
func (p *ProjectHandler) DeleteProject(project models.Project) (*models.EventContext, *models.Error) {
	return deleteWithEventContext(context.TODO(), p.Scheme+"://"+p.getBaseURL()+v1ProjectPath+"/"+project.ProjectName, p)
}

// GetProject returns a project
func (p *ProjectHandler) GetProject(project models.Project) (*models.Project, *models.Error) {
	return getProject(p.Scheme+"://"+p.getBaseURL()+v1ProjectPath+"/"+project.ProjectName, p)
}

// GetProjects returns a project
func (p *ProjectHandler) GetAllProjects() ([]*models.Project, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	projects := []*models.Project{}

	nextPageKey := ""

	for {
		url, err := url.Parse(p.Scheme + "://" + p.getBaseURL() + v1ProjectPath)
		if err != nil {
			return nil, err
		}
		q := url.Query()
		if nextPageKey != "" {
			q.Set("nextPageKey", nextPageKey)
			url.RawQuery = q.Encode()
		}
		req, err := http.NewRequest("GET", url.String(), nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		addAuthHeader(req, p)

		resp, err := p.HTTPClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode == 200 {

			received := &models.Projects{}
			if err = received.FromJSON(body); err != nil {
				return nil, err
			}
			projects = append(projects, received.Projects...)

			if received.NextPageKey == "" || received.NextPageKey == "0" {
				break
			}
			nextPageKey = received.NextPageKey
		} else {
			return nil, handleErrStatusCode(resp.StatusCode, body).ToError()
		}
	}

	return projects, nil
}

func getProject(uri string, api APIService) (*models.Project, *models.Error) {

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, buildErrorResponse(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	addAuthHeader(req, api)

	resp, err := api.getHTTPClient().Do(req)
	if err != nil {
		return nil, buildErrorResponse(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, buildErrorResponse(err.Error())
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if len(body) > 0 {
			respProject := &models.Project{}
			if err = respProject.FromJSON(body); err != nil {
				return nil, buildErrorResponse(err.Error())
			}

			return respProject, nil
		}
		return nil, nil
	}

	return nil, handleErrStatusCode(resp.StatusCode, body)
}

func (p *ProjectHandler) UpdateConfigurationServiceProject(project models.Project) (*models.EventContext, *models.Error) {
	bodyStr, err := project.ToJSON()
	if err != nil {
		return nil, buildErrorResponse(err.Error())
	}
	return putWithEventContext(context.TODO(), p.Scheme+"://"+p.getBaseURL()+v1ProjectPath+"/"+project.ProjectName, bodyStr, p)
}
