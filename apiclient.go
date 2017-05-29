package apiclient

import (
	"github.com/golang/protobuf/proto"

	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path"
	"strconv"
	"strings"
)

type APIError struct {
	StatusCode int
	Body       string
}

func (e APIError) Error() string {
	return "web server error " + strconv.FormatInt(int64(e.StatusCode), 10)
}

type APIClient struct {
	// Base URL of the GRR AdminUI server, including the username and
	// password for HTTP Basic Authentication, e.g.
	// url.Parse("http://admin:admin@localhost:8000")
	BaseURL *url.URL
	// Client to use for web requests. If nil, it will be filled
	// on-demand with a copy http.DefaultClient to which a
	// "net/http/cookiejar".Jar has been added. The cookie jar is
	// needed for handling the csrf token mechanism.
	Client    *http.Client
	csrftoken string
}

func (c *APIClient) client() *http.Client {
	if c.Client == nil {
		c.Client = new(http.Client)
		*c.Client = *http.DefaultClient
		c.Client.Jar, _ = cookiejar.New(nil)
	}
	return c.Client
}

func (c *APIClient) getCSRFToken() string {
	if c.csrftoken == "" {
		rq, _ := http.NewRequest("GET", c.BaseURL.String(), nil)
		c.client().Do(rq)
		for _, cookie := range c.client().Jar.Cookies(c.BaseURL) {
			if cookie.Name == "csrftoken" {
				c.csrftoken = cookie.Value
				break
			}
		}
	}
	return c.csrftoken
}

func (c *APIClient) dohttprequest(rq *http.Request) (*http.Response, error) {
	rq.Header.Set("x-csrftoken", c.getCSRFToken())
	rs, err := c.client().Do(rq)
	if err != nil {
		return nil, err
	}
	if strings.HasPrefix(rs.Header.Get("Content-Type"), "application/json") {
		// skip XSS header
		var buf [5]byte
		rs.Body.Read(buf[:])
	}
	if rs.StatusCode >= 400 {
		buf := &bytes.Buffer{}
		io.Copy(buf, rs.Body)
		return nil, APIError{
			StatusCode: rs.StatusCode,
			Body:       buf.String(),
		}
	}
	return rs, nil
}

// GET and POST calls where both the request and the response contain
// a JSON body
func (c *APIClient) do(method, apipath string, rqm, rsm proto.Message) error {
	u := *c.BaseURL
	u.Path = path.Join(u.Path, apipath)
	u.RawQuery = "strip_type_info=1"
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(rqm); err != nil {
		return err
	}
	rq, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return err
	}
	rs, err := c.dohttprequest(rq)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(rs.Body).Decode(rsm); err != nil {
		return err
	}
	return nil
}

// Simple GET calls where the response contains a JSON body
func (c *APIClient) get(apipath string, values url.Values, rsm proto.Message) error {
	u := *c.BaseURL
	u.Path = path.Join(u.Path, apipath)
	if values == nil {
		values = make(url.Values)
	}
	values.Set("strip_type_info", "1")
	u.RawQuery = values.Encode()
	rq, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}
	rs, err := c.dohttprequest(rq)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(rs.Body).Decode(rsm); err != nil {
		return err
	}
	return nil
}

// Simple POST calls where only the request contains a meaningful JSON body
func (c *APIClient) post(apipath string, rqm proto.Message) error {
	u := *c.BaseURL
	u.Path = path.Join(u.Path, apipath)
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(rqm); err != nil {
		return err
	}
	rq, err := http.NewRequest("POST", u.String(), buf)
	if err != nil {
		return err
	}
	if _, err = c.dohttprequest(rq); err != nil {
		return err
	}
	return nil
}
