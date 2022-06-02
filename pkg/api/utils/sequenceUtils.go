package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	v2 "github.com/keptn/go-utils/pkg/api/utils/v2"
	"github.com/keptn/go-utils/pkg/common/httputils"
)

const v1SequenceControlPath = "/v1/sequence/%s/%s/control"

type SequencesV1Interface interface {
	ControlSequence(params SequenceControlParams) error
}

type SequenceControlHandler struct {
	sequenceControlHandler v2.SequenceControlHandler
	BaseURL                string
	AuthToken              string
	AuthHeader             string
	HTTPClient             *http.Client
	Scheme                 string
}

type SequenceControlParams struct {
	Project      string `json:"project"`
	KeptnContext string `json:"keptnContext"`
	Stage        string `json:"stage"`
	State        string `json:"state"`
}

func (s *SequenceControlParams) Validate() error {
	var errMsg []string
	if s.Project == "" {
		errMsg = append(errMsg, "project parameter not set")
	}
	if s.KeptnContext == "" {
		errMsg = append(errMsg, "keptn context parameter not set")
	}
	if s.State == "" {
		errMsg = append(errMsg, "sequence state parameter not set")
	}
	errStr := strings.Join(errMsg, ",")

	if len(errMsg) > 0 {
		return fmt.Errorf("failed to validate sequence control parameters: %s", errStr)
	}
	return nil
}

type SequenceControlBody struct {
	Stage string `json:"stage"`
	State string `json:"state"`
}

// Converts object to JSON string
func (s *SequenceControlBody) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

// FromJSON converts JSON string to object
func (s *SequenceControlBody) FromJSON(b []byte) error {
	var res SequenceControlBody
	if err := json.Unmarshal(b, &res); err != nil {
		return err
	}
	*s = res
	return nil
}

// NewSequenceControlHandler returns a new SequenceControlHandler
func NewSequenceControlHandler(baseURL string) *SequenceControlHandler {
	return NewSequenceControlHandlerWithHTTPClient(baseURL, &http.Client{Transport: wrapOtelTransport(getClientTransport(nil))})
}

// NewSequenceControlHandlerWithHTTPClient returns a new SequenceControlHandler using the specified http.Client
func NewSequenceControlHandlerWithHTTPClient(baseURL string, httpClient *http.Client) *SequenceControlHandler {
	return createSequenceControlHandler(baseURL, "", "", httpClient, "http")
}

// NewAuthenticatedSequenceControlHandler returns a new SequenceControlHandler that authenticates at the api via the provided token
// Deprecated: use APISet instead
func NewAuthenticatedSequenceControlHandler(baseURL string, authToken string, authHeader string, httpClient *http.Client, scheme string) *SequenceControlHandler {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	httpClient.Transport = wrapOtelTransport(getClientTransport(httpClient.Transport))
	return createAuthenticatedSequenceControlHandler(baseURL, authToken, authHeader, httpClient, scheme)
}

func createAuthenticatedSequenceControlHandler(baseURL string, authToken string, authHeader string, httpClient *http.Client, scheme string) *SequenceControlHandler {
	baseURL = strings.TrimRight(baseURL, "/")
	if !strings.HasSuffix(baseURL, shipyardControllerBaseURL) {
		baseURL += "/" + shipyardControllerBaseURL
	}

	return createSequenceControlHandler(baseURL, authToken, authHeader, httpClient, scheme)
}

func createSequenceControlHandler(baseURL string, authToken string, authHeader string, httpClient *http.Client, scheme string) *SequenceControlHandler {
	baseURL = httputils.TrimHTTPScheme(baseURL)
	return &SequenceControlHandler{
		BaseURL:    baseURL,
		AuthHeader: authHeader,
		AuthToken:  authToken,
		HTTPClient: httpClient,
		Scheme:     scheme,

		sequenceControlHandler: v2.SequenceControlHandler{
			BaseURL:    baseURL,
			AuthHeader: authHeader,
			AuthToken:  authToken,
			HTTPClient: httpClient,
			Scheme:     scheme,
		},
	}
}

func (s *SequenceControlHandler) getBaseURL() string {
	return s.BaseURL
}

func (s *SequenceControlHandler) getAuthToken() string {
	return s.AuthToken
}

func (s *SequenceControlHandler) getAuthHeader() string {
	return s.AuthHeader
}

func (s *SequenceControlHandler) getHTTPClient() *http.Client {
	return s.HTTPClient
}

func (s *SequenceControlHandler) ControlSequence(params SequenceControlParams) error {
	return s.sequenceControlHandler.ControlSequence(
		context.TODO(),
		v2.SequenceControlParams{
			Project:      params.Project,
			KeptnContext: params.KeptnContext,
			Stage:        params.Stage,
			State:        params.State,
		},
		v2.SequencesControlSequenceOptions{})
}
