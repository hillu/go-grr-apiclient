package apiclient

import (
	"github.com/golang/protobuf/proto"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
)

type examples map[string][]struct {
	Method               string
	URL                  string
	RequestPayload       json.RawMessage
	Response             json.RawMessage
	TestClass            string
	TypeStrippedResponse json.RawMessage `json:"type_stripped_response"`
}

type methodpath struct{ method, path string }

type response struct {
	URL                            *url.URL
	Response, TypeStrippedResponse json.RawMessage
}

type handler struct {
	t        *testing.T
	testdata map[methodpath][]response
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	_, useStripped := query["strip_type_info"]
	for _, s := range h.testdata[methodpath{r.Method, r.URL.Path}] {
		delete(query, "strip_type_info")
		if !reflect.DeepEqual(query, s.URL.Query()) {
			continue
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(")]}'\n"))
		if useStripped {
			h.t.Logf("URL: %s, Output: %s", s.URL, string(s.TypeStrippedResponse))
			w.Write(s.TypeStrippedResponse)
			return
		} else {
			h.t.Logf("URL: %s, Output: %s", s.URL, string(s.Response))
			w.Write(s.Response)
			return
		}
	}
}

// newHandler reads test_data/api-docs-examples.json (taken from GRR)
// and returns a http.Handler that acts according to the examples.
func newHandler(t *testing.T) http.Handler {
	f, err := os.Open("test_data/api-docs-examples.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	e := make(examples)
	h := handler{
		t:        t,
		testdata: make(map[methodpath][]response),
	}
	d := json.NewDecoder(f)
	if err := d.Decode(&e); err != nil {
		t.Fatal(err)
	}
	for _, samples := range e {
		for _, s := range samples {
			u, _ := url.Parse(s.URL)
			mp := methodpath{s.Method, u.Path}
			h.testdata[mp] = append(h.testdata[mp],
				response{u, s.Response, s.TypeStrippedResponse})
		}
	}
	return h
}

func TestApiExamples(t *testing.T) {
	s := httptest.NewServer(newHandler(t))
	defer s.Close()
	u, _ := url.Parse(s.URL)
	c := &APIClient{BaseURL: u}
	for i, test := range []struct {
		f        func() (proto.Message, error)
		expected proto.Message
	}{
		{
			func() (proto.Message, error) {
				return c.SearchClients(ApiSearchClientsArgs{Query: proto.String("C.1000000000000000")})
			},
			&ApiSearchClientsResult{
				Items: []*ApiClient{
					{
						AgentInfo: &ClientInformation{
							ClientName: proto.String("GRR Monitor"),
						},
						HardwareInfo: &HardwareInfo{},
						LastSeenAt:   proto.Uint64(42000000),
						OsInfo: &Uname{
							Fqdn:    proto.String("Host-0.example.com"),
							Node:    proto.String("Host-0"),
							Version: proto.String(""),
						},
						Urn:   proto.String("aff4:/C.1000000000000000"),
						Users: nil,
					},
				},
			}},
	} {
		got, err := test.f()
		if err != nil {
			t.Errorf("test %d: %s", i, err)
		}
		if proto.MarshalTextString(got) != proto.MarshalTextString(test.expected) {
			t.Errorf("Unexpected result\nGot:\n%s\nExpected:\n%s",
				proto.MarshalTextString(got),
				proto.MarshalTextString(test.expected))
		}
	}
}
