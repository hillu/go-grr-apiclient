package apiclient

import (
	"github.com/golang/protobuf/proto"

	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type APIError struct {
	StatusCode int
	Body       string
}

func (e APIError) Error() string {
	return "web server error " + strconv.FormatInt(int64(e.StatusCode), 10)
}

type APIClient struct {
	BaseURL    *url.URL
	User, Pass string
	Client     *http.Client
}

func (c APIClient) client() *http.Client {
	if c.Client == nil {
		return http.DefaultClient
	}
	return c.Client
}

func (c APIClient) do(method, apipath string, rqm, rsm proto.Message) error {
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
	rq.SetBasicAuth(c.User, c.Pass)
	rs, err := c.client().Do(rq)
	if err != nil {
		return err
	}
	if rs.StatusCode >= 400 {
		buf := &bytes.Buffer{}
		io.Copy(buf, rs.Body)
		return APIError{
			StatusCode: rs.StatusCode,
			Body:       buf.String(),
		}
	}
	skipXSS(rs.Body)
	if err := json.NewDecoder(rs.Body).Decode(rsm); err != nil {
		return err
	}
	return nil
}

func (c APIClient) get(apipath string, values url.Values, rsm proto.Message) error {
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
	rq.SetBasicAuth(c.User, c.Pass)
	rs, err := c.client().Do(rq)
	if err != nil {
		return err
	}
	if rs.StatusCode >= 400 {
		buf := &bytes.Buffer{}
		io.Copy(buf, rs.Body)
		return APIError{
			StatusCode: rs.StatusCode,
			Body:       buf.String(),
		}
	}
	skipXSS(rs.Body)
	if err := json.NewDecoder(rs.Body).Decode(rsm); err != nil {
		return err
	}
	return nil
}

func (c APIClient) post(apipath string, rqm proto.Message) error {
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
	rq.SetBasicAuth(c.User, c.Pass)
	rs, err := c.client().Do(rq)
	if err != nil {
		return err
	}
	if rs.StatusCode >= 400 {
		buf := &bytes.Buffer{}
		io.Copy(buf, rs.Body)
		return APIError{
			StatusCode: rs.StatusCode,
			Body:       buf.String(),
		}
	}
	return nil
}

func skipXSS(r io.Reader) {
	var buf [5]byte
	r.Read(buf[:])
}
