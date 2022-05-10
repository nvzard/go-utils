package v2

import (
	"context"
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"

	"github.com/keptn/go-utils/pkg/common/httputils"

	"github.com/keptn/go-utils/pkg/api/models"
)

const v1ProjectPath = "/v1/project"

type ProjectsV1Interface interface {
	// CreateProject creates a new project.
	CreateProject(project models.Project) (*models.EventContext, *models.Error)

	// CreateProjectWithContext creates a new project.
	CreateProjectWithContext(ctx context.Context, project models.Project) (*models.EventContext, *models.Error)

	// DeleteProject deletes a project.
	DeleteProject(project models.Project) (*models.EventContext, *models.Error)

	// DeleteProjectWithContext deletes a project.
	DeleteProjectWithContext(ctx context.Context, project models.Project) (*models.EventContext, *models.Error)

	// GetProject returns a project.
	GetProject(project models.Project) (*models.Project, *models.Error)

	// GetProjectWithContext returns a project.
	GetProjectWithContext(ctx context.Context, project models.Project) (*models.Project, *models.Error)

	// GetAllProjects returns all projects.
	GetAllProjects() ([]*models.Project, error)

	// GetAllProjectsWithContext returns all projects.
	GetAllProjectsWithContext(ctx context.Context) ([]*models.Project, error)

	// UpdateConfigurationServiceProject updates a configuration service project.
	UpdateConfigurationServiceProject(project models.Project) (*models.EventContext, *models.Error)

	// UpdateConfigurationServiceProjectWithContext updates a configuration service project.
	UpdateConfigurationServiceProjectWithContext(ctx context.Context, project models.Project) (*models.EventContext, *models.Error)
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

// CreateProject creates a new project.
func (p *ProjectHandler) CreateProject(project models.Project) (*models.EventContext, *models.Error) {
	return p.CreateProjectWithContext(context.TODO(), project)
}

// CreateProjectWithContext creates a new project.
func (p *ProjectHandler) CreateProjectWithContext(ctx context.Context, project models.Project) (*models.EventContext, *models.Error) {
	bodyStr, err := project.ToJSON()
	if err != nil {
		return nil, buildErrorResponse(err.Error())
	}
	return postWithEventContext(ctx, p.Scheme+"://"+p.getBaseURL()+v1ProjectPath, bodyStr, p)
}

// DeleteProject deletes a project.
func (p *ProjectHandler) DeleteProject(project models.Project) (*models.EventContext, *models.Error) {
	return p.DeleteProjectWithContext(context.TODO(), project)
}

// DeleteProjectWithContext deletes a project.
func (p *ProjectHandler) DeleteProjectWithContext(ctx context.Context, project models.Project) (*models.EventContext, *models.Error) {
	return deleteWithEventContext(ctx, p.Scheme+"://"+p.getBaseURL()+v1ProjectPath+"/"+project.ProjectName, p)
}

// GetProject returns a project.
func (p *ProjectHandler) GetProject(project models.Project) (*models.Project, *models.Error) {
	return p.GetProjectWithContext(context.TODO(), project)
}

// GetProjectWithContext returns a project.
func (p *ProjectHandler) GetProjectWithContext(ctx context.Context, project models.Project) (*models.Project, *models.Error) {
	body, mErr := getAndExpectSuccess(ctx, p.Scheme+"://"+p.getBaseURL()+v1ProjectPath+"/"+project.ProjectName, p)
	if mErr != nil {
		return nil, mErr
	}

	respProject := &models.Project{}
	if err := respProject.FromJSON(body); err != nil {
		return nil, buildErrorResponse(err.Error())
	}

	return respProject, nil
}

// GetAllProjects returns all projects.
func (p *ProjectHandler) GetAllProjects() ([]*models.Project, error) {
	return p.GetAllProjectsWithContext(context.TODO())
}

// GetAllProjectsWithContext returns all projects.
func (p *ProjectHandler) GetAllProjectsWithContext(ctx context.Context) ([]*models.Project, error) {
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

		body, mErr := getAndExpectOK(ctx, url.String(), p)
		if mErr != nil {
			return nil, mErr.ToError()
		}

		received := &models.Projects{}
		if err = received.FromJSON(body); err != nil {
			return nil, err
		}
		projects = append(projects, received.Projects...)

		if received.NextPageKey == "" || received.NextPageKey == "0" {
			break
		}
		nextPageKey = received.NextPageKey
	}

	return projects, nil
}

// UpdateConfigurationServiceProject updates a configuration service project.
func (p *ProjectHandler) UpdateConfigurationServiceProject(project models.Project) (*models.EventContext, *models.Error) {
	return p.UpdateConfigurationServiceProjectWithContext(context.TODO(), project)
}

// UpdateConfigurationServiceProjectWithContext updates a configuration service project.
func (p *ProjectHandler) UpdateConfigurationServiceProjectWithContext(ctx context.Context, project models.Project) (*models.EventContext, *models.Error) {
	bodyStr, err := project.ToJSON()
	if err != nil {
		return nil, buildErrorResponse(err.Error())
	}
	return putWithEventContext(ctx, p.Scheme+"://"+p.getBaseURL()+v1ProjectPath+"/"+project.ProjectName, bodyStr, p)
}
