package mbconnect

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// PlainResponse is a helper for receiving blank HTTP
// envelop responses without any payloads.
type PlainResponse struct {
	Code    int    `json:"code"`
	Message string `json:"string"`
}

// Client represents interface for Moneybots Connect client.
type Client struct {
	userId     string
	enctoken   string
	debug      bool
	baseURI    string
	httpClient HTTPClient
}

const (
	name           string        = "mbconnect"
	version        string        = "1.0.0"
	requestTimeout time.Duration = 7000 * time.Millisecond
	baseURI        string        = "https://api.moneybots.app"
)

// Useful public constants
const (
	// Details
	DetailsInstruments          = "i"
	DetailsInstrumentsWithToken = "it"
	DetailsInstrumentToken      = "t"
)

// API endpoints
const (
	// session
	URISessionLogin  string = "/session/token"
	URISessionLogout string = "/session/token"
	URISessionTotp   string = "/session/totp"
	URISessionValid  string = "/session/valid"

	// instruments
	URIInstrumentsInfo               string = "/instruments/info"
	URIInstrumentsQuery              string = "/instruments/query"
	URIInstrumentsOptionchain        string = "/instruments/fno/optionchain"
	URIInstrumentsFNOSegmentExpiries string = "/instruments/fno/segment_expiries/%s"
	URIInstrumentsFNOSegmentNames    string = "/instruments/fno/segment_names/%s"

	// indices
	URIIndicesAll              string = "/indices/all"
	URIIndicesByExchange       string = "/indices/%s/info"
	URIIndicesIndexInstruments string = "/indices/%s/%s/instruments"
)

// New creates a new client.
func New(userId string) *Client {
	client := &Client{
		userId:  userId,
		baseURI: baseURI,
	}

	// Create a default http handler with default timeout.
	client.SetHTTPClient(&http.Client{
		Timeout: requestTimeout,
	})

	return client
}

// SetHTTPClient overrides default http handler with a custom one.
// This can be used to set custom timeouts and transport.
func (c *Client) SetHTTPClient(h *http.Client) {
	c.httpClient = NewHTTPClient(h, nil, c.debug)
}

// SetDebug sets debug mode to enable HTTP logs.
func (c *Client) SetDebug(debug bool) {
	c.debug = debug
	c.httpClient.GetClient().debug = debug
}

// SetBaseURI overrides the base Moneybots API endpoint with custom url.
func (c *Client) SetBaseURI(baseURI string) {
	c.baseURI = baseURI
}

// SetTimeout sets request timeout for default http client.
func (c *Client) SetTimeout(timeout time.Duration) {
	hClient := c.httpClient.GetClient().client
	hClient.Timeout = timeout
}

// SetEnctoken sets the enctoken to the instance.
func (c *Client) SetEnctoken(enctoken string) {
	c.enctoken = enctoken
}

func (c *Client) doEnvelope(method, uri string, params url.Values, headers http.Header, v interface{}) error {
	if params == nil {
		params = url.Values{}
	}
	headers = c.getHeaders(headers)
	return c.httpClient.DoEnvelope(method, c.baseURI+uri, params, headers, v)
}

func (c *Client) do(method, uri string, params url.Values, headers http.Header) (HTTPResponse, error) {
	if params == nil {
		params = url.Values{}
	}
	headers = c.getHeaders(headers)
	return c.httpClient.Do(method, c.baseURI+uri, params, headers)
}

func (c *Client) doRaw(method, uri string, reqBody []byte, headers http.Header) (HTTPResponse, error) {
	headers = c.getHeaders(headers)
	return c.httpClient.DoRaw(method, c.baseURI+uri, reqBody, headers)
}

func (c *Client) getHeaders(headers http.Header) http.Header {
	if headers == nil {
		headers = map[string][]string{}
	}
	headers.Add("User-Agent", fmt.Sprintf("%s/%s", name, version))
	if c.userId != "" && c.enctoken != "" {
		headers.Add("Authorization", fmt.Sprintf("%s:%s", c.userId, c.enctoken))
	}
	return headers
}
