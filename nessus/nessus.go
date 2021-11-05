package nessus

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"reflect"
	"strings"
)

// ClientWithResponses builds on ClientInterface to offer response payloads
type APIClient struct {
	*ClientWithResponses

	Vulnerabilities     VulnerabilitiesInterface
	AgentExclusions     AgentExclusionsInterface
	Assets              AssetsInterface
	Policies            PoliciesInterface
	AuditLog            AuditLogInterface
	Editor              EditorInterface
	IoFilters           IoFiltersInterface
	IoPlugins           IoPluginsInterface
	IoV2                IoV2Interface
	AgentConfig         AgentConfigInterface
	AgentGroup          AgentGroupInterface
	IoScans             IoScansInterface
	IoV1                IoV1Interface
	Exclusions          ExclusionsInterface
	IoScanner           IoScannerInterface
	IoAgent             IoAgentInterface
	Scanners            ScannersInterface
	Scans               ScansInterface
	TargetGroups        TargetGroupsInterface
	Workbenches         WorkbenchesInterface
	ExportsVulns        ExportsVulnsInterface
	Folders             FoldersInterface
	Tags                TagsInterface
	Permissions         PermissionsInterface
	Scanner             ScannerInterface
	Bulk                BulkInterface
	IoNetworks          IoNetworksInterface
	Networks            NetworksInterface
	Credentials         CredentialsInterface
	IoExportsCompliance IoExportsComplianceInterface
}

// do execute and evaluate the request
func (c *Client) do(ctx context.Context, req *http.Request) error {
	// Headers for all request
	req.Header.Set("User-Agent", c.userAgent)
	if c.token != "" {
		req.Header.Set("X-Cookie", "token="+c.token)
	}
	return nil
}

// NewAPIClient Create a new API (yes, naming is awkward)
func NewAPIClient(baseURL string, opts ...ClientOption) (*APIClient, error) {
	cl, err := NewClient(baseURL, opts...)
	cl.RequestEditors = append(cl.RequestEditors, cl.do)
	if err != nil {
		return nil, err
	}
	clientWithResponses := &ClientWithResponses{cl}
	return &APIClient{
		ClientWithResponses: clientWithResponses,
		Credentials:         &Credentials{clientWithResponses},
		IoExportsCompliance: &IoExportsCompliance{clientWithResponses},
		IoNetworks:          &IoNetworks{clientWithResponses},
		Networks:            &Networks{clientWithResponses},
		AgentExclusions:     &AgentExclusions{clientWithResponses},
		Assets:              &Assets{clientWithResponses},
		Vulnerabilities:     &Vulnerabilities{clientWithResponses},
		Policies:            &Policies{clientWithResponses},
		AgentConfig:         &AgentConfig{clientWithResponses},
		AgentGroup:          &AgentGroup{clientWithResponses},
		AuditLog:            &AuditLog{clientWithResponses},
		Editor:              &Editor{clientWithResponses},
		IoFilters:           &IoFilters{clientWithResponses},
		IoPlugins:           &IoPlugins{clientWithResponses},
		IoV2:                &IoV2{clientWithResponses},
		Exclusions:          &Exclusions{clientWithResponses},
		IoScanner:           &IoScanner{clientWithResponses},
		IoScans:             &IoScans{clientWithResponses},
		IoV1:                &IoV1{clientWithResponses},
		ExportsVulns:        &ExportsVulns{clientWithResponses},
		Folders:             &Folders{clientWithResponses},
		IoAgent:             &IoAgent{clientWithResponses},
		Scanners:            &Scanners{clientWithResponses},
		Scans:               &Scans{clientWithResponses},
		TargetGroups:        &TargetGroups{clientWithResponses},
		Workbenches:         &Workbenches{clientWithResponses},
		Permissions:         &Permissions{clientWithResponses},
		Scanner:             &Scanner{clientWithResponses},
		Tags:                &Tags{clientWithResponses},
		Bulk:                &Bulk{clientWithResponses},
	}, nil
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	innerClient HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn

	ctx       context.Context
	userAgent string
	token     string
	user      string
	password  string
	insecure  bool
	trace     bool
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// NewClientFromEnvironment creates a new client from default environment variables
func NewClientFromEnvironment(opts ...ClientOption) (*APIClient, error) {
	baseURL := os.Getenv("NESSUS_URL")
	user := os.Getenv("NESSUS_USER")
	password := os.Getenv("NESSUS_PASSWORD")
	if os.Getenv("NESSUS_INSECURE") == "true" {
		opts = append(opts, WithInsecure(true))
	}

	opts = append(opts, WithLogin(user, password))

	c, err := NewAPIClient(baseURL, opts...)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	c := Client{
		Server:    server,
		userAgent: "go-nessus",
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&c); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(c.Server, "/") {
		c.Server += "/"
	}
	if c.ctx == nil {
		c.ctx = context.Background()
	}

	// create httpClient, if not already present
	if c.Client == nil {
		c.Client = &http.Client{}
	}

	// create httpClient, if not already present
	// c.Client = c
	c.innerClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: c.insecure, // test server certificate is not trusted.
			},
		},
	}
	return &c, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithTrace write all requests to the log
func WithTrace(trace bool) ClientOption {
	return func(c *Client) error {
		c.trace = trace
		return nil
	}
}

// WithUserAgent specify a user agent string to identify the client
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.userAgent = userAgent
		return nil
	}
}

// WithLogin specifies the credentials for
func WithLogin(user string, password string) ClientOption {
	return func(c *Client) error {
		c.user = user
		c.password = password
		return nil
	}
}

// WithContext specifies the credentials for
func WithContext(ctx context.Context) ClientOption {
	return func(c *Client) error {
		c.ctx = ctx
		return nil
	}
}

// WithInsecure accept all certificates
func WithInsecure(insecure bool) ClientOption {
	return func(c *Client) error {
		c.insecure = insecure
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// Do wrap the doer for tracing
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	r, e := c.innerClient.Do(req)
	if c.trace {
		var reqStr = ""
		dump, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			reqStr = strings.ReplaceAll(strings.TrimRight(string(dump), "\r\n"), "\n", "\n                            ")
		}
		if r == nil {
			dump = nil
			err = nil
		} else {
			dump, err = httputil.DumpResponse(r, true)
		}
		if err == nil {
			c.Tracef("%s\n\n                            %s\n", reqStr, strings.ReplaceAll(strings.TrimRight(string(dump), "\r\n"), "\n", "\n                            "))
		}
	}
	return r, e
}

// Errorf logs errors
func (c *Client) Errorf(format string, v ...interface{}) {
	log.Printf("[ERROR] %s", fmt.Sprintf(format, v...))
}

// Warnf logs warings
func (c *Client) Warnf(format string, v ...interface{}) {
	log.Printf("[WARN] %s", fmt.Sprintf(format, v...))
}

// Debugf logs debug info
func (c *Client) Debugf(format string, v ...interface{}) {
	log.Printf("[DEBUG] %s", fmt.Sprintf(format, v...))
}

// Tracef logs trace info
func (c *Client) Tracef(format string, v ...interface{}) {
	log.Printf("[TRACE] %s", fmt.Sprintf(format, v...))
}

// RawAPIResponse generic response wrapper
type RawAPIResponse interface {
	Status() string
	StatusCode() int
}

func getResponseObject(sr RawAPIResponse) (interface{}, error) {
	fldForCode := fmt.Sprintf("JSON%d", sr.StatusCode())
	v := reflect.ValueOf(sr).Elem()
	if _, ok := v.Type().FieldByName(fldForCode); ok {
		s := v.FieldByName(fldForCode).Interface()

		v := reflect.ValueOf(s).Elem()
		if _, ok := v.Type().FieldByName("Data"); ok {
			d := v.FieldByName("Data").Interface()
			return d, nil
		}

		if sr.StatusCode() > 399 {
			return s, errors.New(sr.Status())
		}
		return s, nil

	}
	if sr.StatusCode() > 399 {
		return sr, errors.New(sr.Status())
	}
	return sr, nil
}

func (c *ClientWithResponses) Authenticated() bool {
	return c.ClientInterface.(*Client).token != ""
}

type AuthResponse struct {
	MD5SumWizardTemplates string "json:\"md5sum_wizard_templates,omitempty\""
	Token                 string "json:\"token,omitempty\""
	MD5SumTenableLinks    string "json:\"md5sum_tenable_links,omitempty\""
}

//Authenticate login using basic auth to optain a token
func (c *ClientWithResponses) Authenticate() error {
	// Authenticate
	c.ClientInterface.(*Client).token = ""

	var authJson = []byte(fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\" }", c.ClientInterface.(*Client).user, c.ClientInterface.(*Client).password))
	req, err := http.NewRequest("POST", c.ClientInterface.(*Client).Server+"session", bytes.NewBuffer(authJson))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	res, err := c.ClientInterface.(*Client).Do(req)
	if err != nil {
		return err
	}
	if res == nil {
		return fmt.Errorf("authentication failed")
	}
	if res.StatusCode > 399 {
		return fmt.Errorf("%s returned %s", c.ClientInterface.(*Client).Server, res.Status)
	}
	if (res.Body == nil) {
		return fmt.Errorf("%s returned empty result", c.ClientInterface.(*Client).Server)
	}
	decoder := json.NewDecoder(res.Body)
	var authResponse AuthResponse
	err = decoder.Decode(&authResponse)
	if res == nil {
		return fmt.Errorf("%s returned currupt result", c.ClientInterface.(*Client).Server)
	}
	c.ClientInterface.(*Client).token = authResponse.Token

	return nil
}

func (c *ClientWithResponses) evaluateResponse(response RawAPIResponse, err error) (interface{}, error) {
	if err != nil {
		return nil, err
	}

	return getResponseObject(response)
}
